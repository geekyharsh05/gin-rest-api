# CRUD API with Gin

This is a simple CRUD API built using the Go programming language and the Gin web framework. The API allows you to manage tags through basic operations like creating, reading, updating, and deleting.

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