
# Weather CEP

Weather CEP is a simple API that provides weather information based on the Brazilian postal code (CEP).

## Prerequisites

- Go 1.23.6+
- Docker
- Make
- Cloud Run VSCode extension

## Setup

1. Clone the repository:
    ```sh
    git https://github.com/pimentafm/weatherapi
    cd weather-cep
    ```

2. Copy the example environment file and update it with your API key:
    ```sh
    cp env.example .env
    ```

3. Update the `.env` file with your project name and weather API key:
    ```plaintext
    PROJECT_NAME=fullcycle
    WEATHERAPI_API_KEY=your_weather_api_key_here
    ```

## Build and Run

### Using Makefile

- To build the project:
    ```sh
    make build
    ```

- To run the project:
    ```sh
    make run
    ```

- To run the tests:
    ```sh
    make test
    ```

- To build the Docker image:
    ```sh
    make weatherapi-build
    ```

- To run the Docker container:
    ```sh
    make weatherapi-run
    ```

- To delete the Docker container:
    ```sh
    make delete-container
    ```

- To clean up Docker resources:
    ```sh
    make docker-cleanup
    ```

## API Endpoints

### Get Temperature

- **URL:** `/temperature/{cep}`
- **Method:** `GET`
- **URL Params:**
    - `cep=[string]` (required) - Brazilian postal code (CEP)
- **Success Response:**
    - **Code:** 200
    - **Content:** `{ "temperature": 25.5 }`
- **Error Response:**
    - **Code:** 400 BAD REQUEST
    - **Content:** `{ "error": "Invalid CEP format" }`

## Testing

You can use the provided HTTP file to test the API endpoints:

1. Open the `api/temperature.http` file in your preferred HTTP client (e.g., VSCode REST Client).
2. Execute the requests to test the API.

```http
### Variables
@baseUrl = https://cloudrun-goexpert-weather-59578092225.us-central1.run.app

### Get Temperature
GET {{baseUrl}}/temperature/35780000
Content-Type: application/json

### Get Temperature - invalid CEP
GET {{baseUrl}}/temperature/35780 00
Content-Type: application/json

### Get Temperature - invalid CEP correct format
GET {{baseUrl}}/temperature/35780300
Content-Type: application/json
```

## License

This project is licensed under the MIT License.
