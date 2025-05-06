# Omnenest Backend

This README provides instructions for running the main service and executing tests with code coverage for the Omnenest Backend project.

## Prerequisites

Before running the backend services, make sure you have the following installed:

- Go programming language
- Docker (if you need containerization)
- [Swagger CLI (`swag`)](https://github.com/swaggo/swag) for generating API documentation

## Running the Service

To run the service for a specific module, follow these steps:

1. Open a terminal and navigate to the directory containing the main file for the desired service.
   > **Note:** Replace &lt;service&gt; with the actual service name
   ```bash
   cd path/to/omnenest-backend/src/app/<service>
   ```
   Execute the main service file.
   ```bash
   go run main.go
   ```
## Running Swagger

To generate Swagger documentation for this service, follow these steps:

1. Open a terminal and navigate to the service's directory:
    ```bash
    cd path/to/omnenest-backend/src/app/<service>
    ```
    Execute the main service file.
    ```bash
    swag init
    ```
## Running TestCode

To run tests for the Omnenest Backend project, follow these steps:

1. Open a terminal and navigate to the root directory of your project:

    > **Note:** Replace &lt;service&gt; with the actual service name
    ```bash
    cd path/to/omnenest-backend
    ```
    #### To run Tests from root folder run below command
    ```bash
    go test <service>/... -coverpkg=./... -coverprofile <service>.out -covermode count
    ```
2. Open a terminal and navigate to the app service directory of your project to run tests from service:

    > **Note:** Replace &lt;service&gt; with the actual service name
    ```bash
    cd path/to/omnenest-backend/src/app/<service>
    ```
    #### To Run Tests from service folder run below command
    ```bash
    go test ./... -coverpkg=./... -coverprofile <service>.out -covermode count
    ```
#### To check function level coverage run below command 
```bash
go tool cover -func ./<service>.out
```
#### To check HTML coverage run below command 
```bash
go tool cover -html ./<service>.out
```
## Automation using Makefile commands

To run Makefile commands open a Git Bash and navigate to the root directory of your project:

1. To generate swagger for all the Services run below command
    ```bash
    make generate-swagger
    ```
2. To Run Test codes for all the Services run below command
    ```bash
    make test-all
    ```
3. To Run Go mod tidy for all the Services run below command
    ```bash
    make go-mod-tidy
    ```
4. To Copy api.yml configuration in setup environment run below command
    ```bash
    make copy-api-yml
    ```
5. To Run all the commands mentioned above (generate swagger, run test codes and go mod tidy and copy api.yml for all the Services) run below command
    ```bash
    make all
    ```
6. To Run Test codes and go mod tidy for a specific Service run below command
    
    > **Note:** Replace &lt;service&gt; with the actual service name
    ```bash
    make <service>
    ```

7. To create new microservice structure run below command
    ```bash
     make create_service servicename=<service> port=<port>
    ```
## Deployment Steps

1. Navigate to the root directory of your project

    > **Note:** Replace &lt;envName&gt; with the required environment name. For example, dev/qa

2. Open the release notes file present as release-notes-&lt;envName&gt;.txt

3. Before overwriting previous notes with the new notes, add the previous one into the respective archive file i.e release-notes-archive-&lt;envName&gt;.txt (dev or qa)

4. Give the heading to the notes with current timestamp following below syntax
  
   "------ BFF build release notes - &lt;envName&gt;-release-yyyymmddhhmmss ------"

5. Mention the Service names as 
    > **Note:** Replace &lt;service&gt; with the actual service name
    
    "=>&lt;service&gt;"

6. List down the changes made in the respective micro-service below it.

7. At last Mention the assigned ticket numbers as "Tickets delivered: ticketNumber1,ticketNumber2..."

8. After saving the files, run below command to create new tag
   
   > **Note:** &lt;tag-name&gt; will be "&lt;envName&gt;-release-yyyymmddhhmmss" given in heading of respective release-notes-&lt;envName&gt;.txt file

   ```bash
   git tag <tag-name>
   ```
9. Push the tag by running below command
   
   ```bash
   git push <tag-name>
   ```