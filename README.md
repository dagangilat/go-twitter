# Twitter-Like App - Go Backend

This is a simple Twitter-like application backend built with Go. It provides a set of API endpoints to create, retrieve, update, and delete tweets. The tweets are stored in a JSON file for persistence.

## API Endpoints

- `GET /tweets`: Retrieve all tweets.
- `GET /tweets/{id}`: Retrieve a specific tweet by ID.
- `POST /tweets`: Create a new tweet.
- `PUT /tweets/{id}`: Update an existing tweet by ID.
- `DELETE /tweets/{id}`: Delete a tweet by ID.

## Installation

1. Clone the repository:

   ```bash
   git clone git@github.com:dagangilat/go-twitter.git
   ```

2. Change to the project directory:

   ```bash
   cd go-twitter
   ```

3. Build the Go backend:

   ```bash
   go build
   ```

4. Run the application:

   ```bash
   ./go-twitter
   ```

   The server will start on port 8000.

## Usage

You can use tools like cURL or Postman to interact with the API endpoints. Here are some examples using cURL:

- List all tweets:

  ```bash
  curl -i http://localhost:8000/tweets
  ```

- Retrieve a tweet by ID:

  ```bash
  curl -i http://localhost:8000/tweets/{id}
  ```

- Create a new tweet:

  ```bash
  curl -i -X POST -H "Content-Type: application/json" -d '{"user":"John","text":"Hello, Twitter!"}' http://localhost:8000/tweets
  ```

- Update a tweet by ID:

  ```bash
  curl -i -X PUT -H "Content-Type: application/json" -d '{"user":"John

 Doe","text":"Updated tweet"}' http://localhost:8000/tweets/{id}
  ```

- Delete a tweet by ID:

  ```bash
  curl -i -X DELETE http://localhost:8000/tweets/{id}
  ```

## Dependencies

- [github.com/gorilla/mux](https://github.com/gorilla/mux) - HTTP router and dispatcher for building Go web servers.
- [github.com/google/uuid](https://github.com/google/uuid) - Package for generating UUIDs.

