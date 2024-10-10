# Twitter-Like Distributed System

This project implements a distributed service-based architecture for a twitter-like application where users can:
- Post short messages (under 400 characters).
- Read messages as a feed (last 10 messages).
- Like or unlike messages.

The system is built with separate services that handle different business domains: **User Service**, **Message Service**, and **Like Service**. It uses REST APIs for communication between services.

## Business Requirements
- **Posting Messages**: Only registered users can post messages (up to 400 characters).
- **Feed**: Any user can read messages, which are displayed as the last 10 messages posted.
- **Liking Messages**: Logged-in users can like and unlike messages. Likes are toggled on the same action.
- **Registration**: Users register with only a username (no passwords required).

## Services

### 1. **User Service**
This service handles user registration and authentication. It ensures that users can register and log in using only a username.

- **Endpoint: POST `/users`**
  - Registers a new user or logs in an existing user based on the provided username.
  - Example request:
    ```json
    {
      "username": "john_doe"
    }
    ```

- **Endpoint: GET `/users/:id`**
  - Returns user information by ID.
  - Example response:
    ```json
    {
      "id": 1,
      "username": "john_doe"
    }
    ```

### 2. **Message Service**
This service allows users to post messages and view the message feed.

- **Endpoint: POST `/messages`**
  - Posts a new message. The user must be authenticated.
  - Example request:
    ```json
    {
      "content": "Hello world!"
    }
    ```

- **Endpoint: GET `/messages`**
  - Returns the last 10 messages as a feed.
  - Example response:
    ```json
    [
      {
        "id": 123,
        "user_id": 1,
        "content": "This is a test message",
        "created_at": "2024-10-10T10:30:00Z"
      }
    ]
    ```

### 3. **Like Service**
This service allows users to like or unlike messages. The same endpoint is used to toggle the like status.

- **Endpoint: POST `/likes`**
  - Toggles a like on a message (adds or removes a like).
  - Example request:
    ```json
    {
      "MessageID": 123
    }
    ```

- **Endpoint: GET `/message/:messageID`**
  - Returns the message details along with the number of likes and whether the user liked the message.
  - Example response:
    ```json
    {
      "message_id": 123,
      "likes": 10,
      "liked": true
    }
    ```

## Prerequisites

- **Node.js**
- **Multiple terminals** to run the services separately.
- **Go**

## Setup Instructions

### 1. Clone the Repository

```bash
git clone git@github.com:kkkaterinaaa/SoftArch-A5-Team12.git
cd SoftArch-A5-Team12
```

### 2. Install Dependencies and Run Each Service

Each service should be run in its own terminal:

#### User Service

```bash
cd backend/user-service
go run main.go
```

#### Message Service

```bash
cd backend/message-service
go run main.go
```

#### Like Service

```bash
cd backend/interction-service
go run main.go
```

#### Proxy

```bash
cd backend/proxy
go run proxy.go
```


#### Frontend Service (Optional)

To run the frontend:

```bash
cd frontend
npm i
npm start
```

## API Usage

1. **Register or Login** (POST `/users`):  
   Register or log in using a username. No password is required.
   
2. **Post Message** (POST `/messages`):  
   Authenticated users can post a message.

3. **Read Messages** (GET `/messages`):  
   Fetch the last 10 messages from the feed.

4. **Like/Unlike Message** (POST `/likes`):  
   Toggle like or unlike on a message.

5. **Get Message Info** (GET `/message/:messageID`):  
   Retrieve message details, including the number of likes and whether the current user liked it.

## Example Workflow

1. **Register or Log in a User**:
   - Send a POST request to `/users` with a `username` field.
   
2. **Post a Message**:
   - Once logged in, send a POST request to `/messages` with `content` to create a new message.
   
3. **Read Messages**:
   - Retrieve the latest 10 messages by sending a GET request to `/messages`.

4. **Like/Unlike a Message**:
   - Toggle a like or unlike by sending a POST request to `/likes` with `MessageID`.

## License

This project is licensed under the MIT License.
