# Caching in Golang 🚀

This repository demonstrates in-memory caching in Golang using the Fiber framework and the `go-cache` library. It includes a simple API to fetch and cache data from an external source, showcasing how to implement caching to enhance application performance.

## Overview 📝

Caching is a technique that temporarily stores frequently accessed data to reduce the time required to fetch it from the source. It is crucial in modern applications as it helps improve performance, reduce server load, and enhance the user experience by minimizing delays. In-memory caching is one of the fastest methods, storing data directly in the server’s memory.

### Key Concepts 🔑

- **In-memory caching**: Stores data in memory, making it quickly accessible.
- **Cache eviction policies**: Determine when to remove data from the cache. In this example, cache entries have a lifetime of 10 minutes and are evicted after 20 minutes of inactivity.
- **Custom headers**: Adds `X-Cache-Status` headers to indicate whether the response was served from the cache (`HIT`) or fetched from the API (`MISS`).

## Prerequisites 📋

- [Golang](https://golang.org/) installed.
- Basic understanding of the Fiber framework.

## Installation ⚙️

1. Install Fiber:
   ```bash
   go get -u github.com/gofiber/fiber/v2
   ```

2. Install `go-cache`:
   ```bash
   go get github.com/patrickmn/go-cache
   ```

## Project Structure 📂

```plaintext
.
├── main.go
├── routes
│   └── routes.go
├── middleware
│   └── cache.go
├── controller
│   └── controller.go
└── structs
    └── struct.go
```

## Running the Project ▶️

1. Clone the repository:
   ```bash
   git clone https://github.com/ArtiomStartev/Go-Cache-API.git
   ```

2. Run the application:
   ```bash
   go run main.go
   ```

3. Access the API:
   ```
   GET http://localhost:8080/posts/:id
   ```

Replace `:id` with the desired post ID.

## Features ✨

- Fetches data from `https://jsonplaceholder.typicode.com/posts`.
- Implements in-memory caching with custom eviction policies.
- Returns cached responses with `X-Cache-Status` headers.

## References 📚

- [Fiber Framework](https://gofiber.io/)
- [go-cache Library](https://github.com/patrickmn/go-cache)

