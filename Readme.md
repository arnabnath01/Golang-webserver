# Go Server

This is a simple HTTP server written in Go. It serves static files and handles specific routes for form submission and a greeting message.

## Features

- Serves static files from the `./static` directory.
- Handles a `/form` route for form submission.
- Handles a `/hello` route for a greeting message.

## How to Run

1. Ensure you have Go installed on your machine. This project uses Go version 1.22.0.
2. Navigate to the project directory.
3. Run the command `go run main.go` to start the server.
4. Open your web browser and visit `localhost:3000` to view the served static files.
5. Visit `localhost:3000/form` to view and submit the form.
6. Visit `localhost:3000/hello` to view the greeting message.

## Project Structure

- `main.go`: This is the main file where the server is started and routes are defined.
- `static/`: This directory contains the static files served by the server.

## Routes

- `/`: Serves the static files in the `./static` directory.
- `/form`: Handles form submission. It accepts POST requests and logs the submitted form values to the server console.
- `/hello`: Returns a greeting message.

**Note:** This is a simple server for demonstration purposes. It does not include any form of validation or error handling for the form.
