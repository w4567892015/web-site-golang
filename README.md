# Web Site Golang

A web application built with Golang.

## Project Structure

This project uses a multi-app structure with Go workspaces.
-   `apps/`: Contains the main applications.
    -   `web/`: The main web application.
-   `libs/`: Contains shared libraries.
    -   `otel/`: Library for OpenTelemetry.
-   `go.work`: Go workspace file.
-   `Dockerfile`: For building and running the application in Docker.
-   `Makefile`: Contains helper commands.

## Prerequisites

-   Go
-   Docker
-   Make

## Development

To start the development server for the `web` application, run:

```sh
make start.dev name=web
```

## Building

### Using Makefile

To build the `web` application, run:

```sh
make release name=web
```
This will create a `main` executable in the project root.

### Using Docker

To build the Docker image, run the following command.

```sh
docker build --build-arg appName=web -t web-site-golang .
```

## Running

### Using the executable

After building with `make release`, you can run the application:

```sh
./main
```

### Using Docker

To run the application in a Docker container:

```sh
docker run -p 3000:3000 web-site-golang
```

The application will be available at `http://localhost:3000`.
