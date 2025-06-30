# Go based MiniCloud REST API

MiniCloud is a lightweight __RESTful API__ written in __Go__ that simulates basic virtual machine lifecycle operations using _Docker containers_.

It is designed for local development, educational purposes, or prototyping infrastructure workflows before integrating with real cloud platforms.

The API provides endpoints for creating, retrieving, updating, and deleting virtual machines (simulated containers).

Each VM is represented by a record in a _SQLite database_ and is backed by a _Docker container_ that simulates the compute instance.

## Key Features

## Accounts

- Create account (username + password) to connect from Terraform Provider
- Get all accounts
- Delete accounts
- Update account
- Verify credential / password

### Virtual Machine (Docker simulation)

- __Create a VM__ – Launches a Docker container based on the provided image, environment variables, and ports, and stores metadata in the database.
- __List VMs__ – Retrieves all VM records stored in the system.
- __Get VM by ID__ – Fetches a single VM and its metadata using its unique identifier.
- __Update a VM__ – Modifies the metadata of an existing VM.
- __Delete a VM__ – Stops and removes the associated Docker container, then deletes the VM record from the database.

### Stroage (Docker volume simulation)

- __List storage__
- __Create storage__
- __Get a storage__
- __Delete a storage__
- __Update storage size__
- __Attach a storage to VM__
- __Deatch a storage from VM__

### Storage Account (MiniO simulation)

## Architecture Overview

- __models/__ – Defines data structures and database logic (SQLite)
- __handlers/__ – Contains HTTP request handlers mapped to API routes
- __services/__ – Implements Docker container management (start, stop, remove)
- __routes/__ – Registers all REST endpoints with the Gin engine
- __db/__ – Handles database initialization and table creation

This project serves as a foundational backend that can later be extended with a Terraform provider or integrated into more complex infrastructure tooling.

## Logical Flow

The API is structured to follow a clear layered approach:

```text
main.go > Routes > Handlers > Models
```

For example, calling GET /vms executes the following:

```text
main.go > GET /vms (Route) > ListVMs() (Handler) > GetAllVms() (Model)
```

This separation ensures a clean architecture where HTTP request handling, business logic, and persistence are properly decoupled.

## Prerequisites

To run this project successfully, the following tools and versions are required:

- __Go__ 1.20 or newer
- __Docker__ installed and available on the system path (used to simulate VMs) and proper permission of the runner user
- __SQLite3__ (used via Go’s github.com/mattn/go-sqlite3 driver)
- __Git__ (to clone the repository, if needed)
- __Postman__ or __REST Client__ extension in VSCode

## Running the Project

To start the MiniCloud REST API locally, follow these steps:

- Make sure prerequisites are installed (see Prerequisites).
- Clone the repository (if needed), then initialize and run the application:

```bash
go run .
```

This will start the HTTP server at __http://localhost:8080__.

## Testing the API

You can test the API using the pre-written __.http__ files found in the `rest-api-tests/` directory.

These files can be executed using a supported IDE/editor such as:

- Visual Studio Code with the REST Client extension.
- JetBrains IDEs like GoLand or IntelliJ, which support .http files natively.

### Available Test Files

#### VM

- __create-vm.http__ > Sends a POST /vms request to create a new VM
- __get-all-vms.http__ > Sends a GET /vms request to retrieve all VMs
- __get-vm-by-id.http__ > Sends a GET /vms/:id request for a specific VM
- __update-vm.http__ > Sends a PUT /vms/:id request to update an existing VM
- __delete-vm-by-id.http__ – Sends a DELETE /vms/:id request to delete a VM and its corresponding

## OpenAPI preparation

Install important packages to Swagger:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```

Steps:

- Give the `annotations` to the handlers.
- Initialize the Swag (`/docs`)
