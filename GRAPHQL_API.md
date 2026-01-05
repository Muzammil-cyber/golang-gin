# GraphQL API Documentation

This project now includes a fully functional GraphQL API for video management alongside the existing REST API.

## Endpoints

- **GraphQL Playground**: `http://localhost:5000/` - Interactive GraphQL IDE
- **GraphQL Query Endpoint**: `http://localhost:5000/query`

## Authentication

GraphQL mutations (create, update, delete) require JWT authentication. Include your JWT token in the Authorization header:

```
Authorization: Bearer <your-jwt-token>
```

To get a JWT token, use the REST login endpoint:
```bash
POST http://localhost:5000/api/login
```

## Schema

### Types

```graphql
type Person {
  id: ID!
  name: String!
  age: Int!
  email: String!
  createdAt: String!
  updatedAt: String!
}

type Video {
  id: ID!
  title: String!
  description: String!
  url: String!
  author: Person!
  createdAt: String!
  updatedAt: String!
}
```

### Queries

#### Get all videos
```graphql
query {
  videos {
    id
    title
    description
    url
    author {
      id
      name
      age
      email
    }
    createdAt
    updatedAt
  }
}
```

#### Get a single video by ID
```graphql
query {
  video(id: "123e4567-e89b-12d3-a456-426614174001") {
    id
    title
    description
    url
    author {
      name
      email
    }
  }
}
```

### Mutations (Requires JWT Authentication)

#### Create a new video
```graphql
mutation {
  createVideo(input: {
    title: "Introduction to GraphQL"
    description: "Learn GraphQL basics"
    url: "https://www.youtube.com/watch?v=abc"
    author: {
      name: "John Doe"
      age: 30
      email: "john.doe@example.com"
    }
  }) {
    id
    title
    description
    url
    author {
      id
      name
      email
    }
    createdAt
  }
}
```

#### Update an existing video
```graphql
mutation {
  updateVideo(
    id: "123e4567-e89b-12d3-a456-426614174001"
    input: {
      title: "Updated Video Title"
      description: "Updated description"
    }
  ) {
    id
    title
    description
    updatedAt
  }
}
```

#### Delete a video
```graphql
mutation {
  deleteVideo(id: "123e4567-e89b-12d3-a456-426614174001")
}
```

## Testing with cURL

### Get all videos (no auth required)
```bash
curl -X POST http://localhost:5000/query \
  -H "Content-Type: application/json" \
  -d '{"query":"{ videos { id title description url author { name email } } }"}'
```

### Create a video (requires JWT)
```bash
curl -X POST http://localhost:5000/query \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "query": "mutation($input: CreateVideoInput!) { createVideo(input: $input) { id title } }",
    "variables": {
      "input": {
        "title": "Test Video",
        "description": "Test description",
        "url": "https://youtube.com/watch?v=test",
        "author": {
          "name": "Test Author",
          "age": 25,
          "email": "test@example.com"
        }
      }
    }
  }'
```

## Using the GraphQL Playground

1. Start the server: `go run main.go`
2. Open your browser to `http://localhost:5000/`
3. For mutations, add the Authorization header in the "HTTP HEADERS" section at the bottom:
```json
{
  "Authorization": "Bearer YOUR_JWT_TOKEN"
}
```

## Features

- ✅ Full CRUD operations for videos
- ✅ JWT authentication for mutations
- ✅ GraphQL Playground for interactive testing
- ✅ Nested queries for author information
- ✅ Input validation
- ✅ Error handling
- ✅ Integration with existing video service and repository
