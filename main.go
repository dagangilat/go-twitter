package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/mux"
)

var (
	DEBUG  = false
	tweets Tweets
)

// Tweet represents a tweet
type Tweet struct {
	ID   string `json:"id,omitempty"`
	User string `json:"user,omitempty"`
	Text string `json:"text,omitempty"`
}

// Tweets represents a collection of tweets
type Tweets struct {
	sync.Mutex
	Tweets []Tweet `json:"tweets"`
}

func initTweets() error {
	file, err := os.Open("tweets.json")
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &tweets)
	if err != nil {
		return err
	}

	return nil
}

func saveTweets() error {
	data, err := json.Marshal(tweets)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("tweets.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func logMessage(message string) {
	if DEBUG {
		log.Println(message)
	}
}

func GetTweets(w http.ResponseWriter, r *http.Request) {
	tweets.Lock()
	defer tweets.Unlock()

	json.NewEncoder(w).Encode(tweets.Tweets)
	logMessage("Listed all tweets")
}

func GetTweet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tweetID := params["id"]

	tweets.Lock()
	defer tweets.Unlock()

	for _, tweet := range tweets.Tweets {
		if tweet.ID == tweetID {
			json.NewEncoder(w).Encode(tweet)
			logMessage(fmt.Sprintf("Retrieved tweet with ID: %s", tweetID))
			return
		}
	}

	http.NotFound(w, r)
	logMessage(fmt.Sprintf("Tweet with ID: %s not found", tweetID))
}

func CreateTweet(w http.ResponseWriter, r *http.Request) {
	var tweet Tweet
	_ = json.NewDecoder(r.Body).Decode(&tweet)

	tweets.Lock()
	defer tweets.Unlock()

	tweets.Tweets = append(tweets.Tweets, tweet)

	err := saveTweets()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tweet)
	logMessage("Created a new tweet")
}

func UpdateTweet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tweetID := params["id"]

	tweets.Lock()
	defer tweets.Unlock()

	for index, tweet := range tweets.Tweets {
		if tweet.ID == tweetID {
			var updatedTweet Tweet
			_ = json.NewDecoder(r.Body).Decode(&updatedTweet)

			tweets.Tweets[index].User = updatedTweet.User
			tweets.Tweets[index].Text = updatedTweet.Text

			err := saveTweets()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			json.NewEncoder(w).Encode(tweets.Tweets[index])
			logMessage(fmt.Sprintf("Updated tweet with ID: %s", tweetID))
			return
		}
	}

	http.NotFound(w, r)
	logMessage(fmt.Sprintf("Tweet with ID: %s not found", tweetID))
}

func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tweetID := params["id"]

	tweets.Lock()
	defer tweets.Unlock()

	for index, tweet := range tweets.Tweets {
		if tweet.ID == tweetID {
			tweets.Tweets = append(tweets.Tweets[:index], tweets.Tweets[index+1:]...)

			err := saveTweets()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusNoContent)
			logMessage(fmt.Sprintf("Deleted tweet with ID: %s", tweetID))
			return
		}
	}

	http.NotFound(w, r)
	logMessage(fmt.Sprintf("Tweet with ID: %s not found", tweetID))
}

func main() {
	err := initTweets()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/tweets", GetTweets).Methods("GET")
	r.HandleFunc("/tweets/{id}", GetTweet).Methods("GET")
	r.HandleFunc("/tweets", CreateTweet).Methods("POST")
	r.HandleFunc("/tweets/{id}", UpdateTweet).Methods("PUT")
	r.HandleFunc("/tweets/{id}", DeleteTweet).Methods("DELETE")

	fmt.Println("Server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
