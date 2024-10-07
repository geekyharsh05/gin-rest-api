# CRUD API with Gin

This is a simple CRUD API built using the Go programming language and the Gin web framework. The API allows you to manage tags through basic operations like creating, reading, updating, and deleting.

## Installation and Setup

### With Docker

<details>
<summary>
Click Me
</summary>

1. Clone the repository
2. Use docker-compose to run the server

```bash
git clone git@github.com:geekyharsh05/gin-rest-api.git
cd gin-rest-api
docker-compose up -d
```

</details>

### Without Docker

<details>
<summary>
Click Me
</summary>

1. Clone the repository
2. Install the dependencies
3. Run the server

```bash
git clone git@github.com:geekyharsh05/gin-rest-api.git
cd cd gin-rest-api
go mod tidy
go run main.go
```

</details>

### NOTE: Setup environment variable in .env from .env.example

You can use PostgresSQL as the database for this project. To set up the database, follow these steps:

```bash
docker run --name postgres -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 postgres:alpine
```

## API Endpoints

The following endpoints are available in the application:

- **GET /api/tags**: Retrieve all tags.
- **GET /api/tags/:tagId**: Retrieve a tag by ID.
- **POST /api/tags**: Create a new tag.
- **PATCH /api/tags/:tagId**: Update an existing tag by ID.
- **DELETE /api/tags/:tagId**: Delete a tag by ID.

## Usage

To use the API, you can make HTTP requests to the endpoints. Here's an example of how to create a new tag:

```bash
curl --request POST \
  --url http://localhost:8888/api/tags \
  --header 'Content-Type: application/json' \
  --data '{
    "name": "Golang"
}
```

### Swagger Documentation

The API also provides a Swagger documentation endpoint. You can access it by making a GET request to the following URL:

```bash
http://localhost:8888/docs/index.html
```

This will display the Swagger documentation page, which includes information about the available endpoints, their parameters, and their responses.

## Author

**Author Name** &nbsp; : &nbsp; Harsh Vardhan Pandey <br>
**Author URI** &nbsp; &nbsp; &nbsp; : &nbsp; [www.aboutharsh.vercel.app](https://aboutharsh.vercel.app/) <br>
**GitHub URI** &nbsp; &nbsp; &nbsp; : &nbsp; [geekyharsh05](https://github.com/geekyharsh05)

## License

[![License: MIT](https://img.shields.io/badge/License-MIT-red.svg)](https://opensource.org/licenses/MIT)