# ResIT Reservation API

![GitHub Workflow Status (branch)](https://img.shields.io/github/workflow/status/FontysResIT/ResIt/CI/development?style=for-the-badge)
![Codecov branch](https://img.shields.io/codecov/c/github/FontysResIT/ResIT/development?style=for-the-badge)

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

## Update API Docs

1. Run `swag init --parseDependency --parseInternal --ot go,yaml` in this current folder
2. New swagger documentation can now be found in the `/docs` folder

_Read [swaggo](https://github.com/swaggo/swag) docs for more info_

## Folder Structure

```c++
├───config          // Configuration files (development, production, testing etc.)
├───logic           // Business logic files
├───model           // Data models
├───repository      // Contains repositories
│   └───mongodb     // MongoDB database repository implementations
└───router          // Routers (Http, gRPC etc.)
    └───http        // Http router
        └───handler // Http router handler functions
```
