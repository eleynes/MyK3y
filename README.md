# Go REST API Boilerplate
Golang API boilerplate using GoFiber, PostgreSQL and Redis

## Folder Structure

- `/api/v1`
    - `/routes` - All API routes are defined here
    - `/controllers` - For validating requests and calling services
    - `/services` - For business logic, database calls and other services
    - `/middlewares` - For authentication, logging, rate limiting etc.

- `/build` - Contains built binary, gitignore'd
- `/cmd` - Initializes the fiber app and basic middlewares configuration
- `/config` - For handling configuration/env variables
- `/db` - For handling database connections 
- `/handlers` - For handling responses and db transactions
- `/models` - GORM Models
- `/types` - For defining custom types that can be used across the app

- `main.go` - Entrypoint of the app


## How to use
1. The repo contains product API implementation for reference.

2. Clone the repo and rename the folder to your project name.

3. Search for `eleynes/MyK3y` in the project and replace it with `<your-github-id/project-name>`.

4. Run `go mod tidy` to install all the dependencies.

5. Copy `.env.example` to `.env` and change the values as per your configuration.

6. Run `go build -o ./build/main && ./build/main` to build and run the app.


## Optional
1. Run `go get github.com/cosmtrek/air` to install air for hot reloading.

2. Run `air` to start the app with hot reloading.


## Notes
- `Success` Handler for successful requests

- `BuildError` Handler for build errors

- `/api/v1` is the base path for all routes except `/` for health check

