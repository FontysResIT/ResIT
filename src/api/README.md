# ResIT API

<a href="https://github.com/RealSnowKid/ResIT">
    <img src="https://raw.githubusercontent.com/RealSnowKid/ResIT/master/img/logo_horizontal.png" alt="Logo" width="170" height="55">
  </a>

## Table of contents

1. How to run
2. Build with
3. Folder structure

## How to run

1. Make sure you have [Go](https://go.dev/) installed
2. Move to the current directory `cd src/api`
3. Run `go get`
4. Run `go run .`

## Build with

- Go
- Gin Framework

<br>
<a href="https://go.dev/">
<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1280px-Go_Logo_Blue.svg.png" alt="Logo"  height="55">
</a>
<a href="https://go.dev/">
<img src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png" alt="Logo"  height="55">
</a>

## Folder Structure

```c++
├───config          // Configuration files (development, production, testing etc.)
├───logic           // Business logic files
├───model           // Data models
├───repository      // Contains repositories
│   └───scylla      // Scylla database repository implementations
└───router          // Routers (Http, gRPC etc.)
    └───http        // Http router
        └───handler // Http router handler functions
```
