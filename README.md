# Ecomili - Inspiration API

This repository contains the backend/API server for Ecomili's "Inspiration" feature, a part of the Ecomili ecommerce platform. The Inspiration feature serves as a source of inspiration for users, showcasing trending or popular News, Infobites etc.

## Technologies Used

- **Go:** The primary language used for the API.
- **GORM:** An Object-Relational Mapping (ORM) library for Golang, used for interacting with the MySQL database.
- **Go Fiber:** A web framework used to build the API endpoints and manage HTTP requests.
- **MySQL:** The chosen relational database for storing data.
- **JWT (JSON Web Tokens):** Used for token-based authentication and authorization.

## Getting Started

### Prerequisites

- Go installed on your machine.
- MySQL database set up and running.
- Dependencies fetched using `go mod`.

### Installation

1. Clone the repository.
2. Install dependencies using `go mod`.
3. Set up the MySQL database and configure the connection in the `.env` file.
4. Run the application.

### Configuration

Make sure to set up your `.env` file with the necessary configurations:

```plaintext
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=password
DB_NAME=ecomili_inspiration
JWT_SECRET=mySecretKey
```
