# Instagram-like Backend API


## Contact

- **Full Name:** [Ayush Katiyar]
- **GitHub:** [https://github.com/aerick8602/Codesmith](https://github.com/aerick8602/Codesmith)

## Overview

A backend API for an Instagram-like app, built with Go and PostgreSQL. It supports user and post management.

## API Endpoints

### User Endpoints

- **Create a User**
  - **URL:** `/users`
  - **Method:** `POST`
  - **Body:** `{"name": "John Doe", "email": "johndoe@example.com", "password": "password123"}`
  - **Responses:** `201 Created`, `400 Bad Request`, `500 Internal Server Error`

- **Get a User by ID**
  - **URL:** `/users/{id}`
  - **Method:** `GET`
  - **Responses:** `200 OK`, `400 Bad Request`, `404 Not Found`

### Post Endpoints

- **Create a Post**
  - **URL:** `/posts`
  - **Method:** `POST`
  - **Body:** `{"caption": "Caption", "imageUrl": "https://example.com/image.jpg", "postedTimestamp": "2024-09-18T12:00:00Z", "userId": 1}`
  - **Responses:** `201 Created`, `400 Bad Request`, `500 Internal Server Error`

- **Get a Post by ID**
  - **URL:** `/posts/{id}`
  - **Method:** `GET`
  - **Responses:** `200 OK`, `400 Bad Request`, `404 Not Found`

- **List Posts by User**
  - **URL:** `/posts/users/{userId}`
  - **Method:** `GET`
  - **Responses:** `200 OK`, `400 Bad Request`, `404 Not Found`

## Setup

1. Clone the repo: `git clone https://github.com/aerick8602/Codesmith.git`
2. Navigate to the directory: `cd Codesmith`
3. Install dependencies: `go mod tidy`
4. Set up PostgreSQL and run SQL commands to create tables.
5. Start the server: `go run main.go`



