# Free Contest API

A REST API for managing officer tests and contests built with Go Gin and Swagger.

## Features

- RESTful API with Gin framework
- Swagger documentation
- Officer test management
- Subject-based testing system
- Random question selection from chapters
- Test caching per officer-subject combination

## API Endpoints

### GET /api/v1/tests/officer-subject

Get a test for an officer in a specific subject with randomly selected questions.

**Parameters:**
- `officerID` (query, required): Officer ID (integer)
- `subjectID` (query, required): Subject ID (integer)

**Response:**
- `200 OK`: Returns a Test object with random questions
- `400 Bad Request`: Invalid or missing parameters
- `404 Not Found`: Officer or subject not found
- `500 Internal Server Error`: Server error

### POST /api/v1/tests/start

Start a test for an officer by setting the start time.

**Parameters:**
- `officerID` (query, required): Officer ID (integer)
- `testID` (query, required): Test ID (integer)

**Response:**
- `200 OK`: Returns the started Test object with updated start time
- `400 Bad Request`: Invalid or missing parameters
- `404 Not Found`: Test or officer not found
- `409 Conflict`: Test already started
- `500 Internal Server Error`: Server error

### GET /api/v1/units

Get all units in the system.

**Response:**
- `200 OK`: Returns a list of all units

**Example Response:**
```json
{
  "count": 9,
  "data": [
    {
      "id": 1,
      "name": "Phòng Tham mưu"
    },
    {
      "id": 2,
      "name": "Phòng Chính trị"
    }
  ]
}
```

### GET /api/v1/officers

Get all officers in the system with their unit information.

**Response:**
- `200 OK`: Returns a list of all officers with their unit details

**Example Response:**
```json
{
  "count": 3,
  "data": [
    {
      "id": 1,
      "name": "Nguyễn Văn A",
      "unit_id": 1,
      "unit": {
        "id": 1,
        "name": "Phòng Tham mưu"
      },
      "score": 0,
      "rank": "",
      "position": ""
    }
  ]
}
```

**Example Request:**
```bash
# Get a test for officer
curl "http://localhost:8080/api/v1/tests/officer-subject?officerID=1&subjectID=1"

# Start a test
curl -X POST "http://localhost:8080/api/v1/tests/start?officerID=1&testID=123"

# Get all units
curl "http://localhost:8080/api/v1/units"

# Get all officers
curl "http://localhost:8080/api/v1/officers"
```

**Example Response:**
```json
{
  "id": 707,
  "name": "",
  "contest_id": "",
  "duration": 3600,
  "subject": {
    "id": 1,
    "name": "Đề thi HCKT",
    "description": "Đề thi HCKT",
    "contest_id": 1,
    "chapters": [...],
    "num_question_test": 13,
    "test_time": 60,
    "folder_path": "..."
  },
  "officer": {
    "id": 1,
    "name": "Nguyễn Văn A",
    "unit_id": 1,
    "score": 0,
    "rank": "",
    "position": ""
  },
  "questions": [...],
  "remaining_time": 3600
}
```

### GET /health

Health check endpoint to verify API status.

**Response:**
```json
{
  "status": "OK",
  "message": "Free Contest API is running"
}
```

## Running the API

1. **Build the application:**
   ```bash
   go build -o bin/api cmd/api/main.go
   ```

2. **Run the server:**
   ```bash
   ./bin/api
   ```

3. **Access the API:**
   - API Base URL: `http://localhost:8080/api/v1`
   - Swagger Documentation: `http://localhost:8080/swagger/index.html`
   - Health Check: `http://localhost:8080/health`

## Configuration

The API uses a `config.json` file in the root directory with the following structure:

```json
{
  "list_unit": [
    {
      "id": 1,
      "name": "Phòng Tham mưu"
    }
  ],
  "list_officer": [
    {
      "id": 1,
      "name": "Nguyễn Văn A",
      "unit_id": 1,
      "score": 0,
      "rank": "",
      "position": ""
    }
  ],
  "contest_path": "/path/to/contest/data"
}
```

## Development

### Prerequisites
- Go 1.24.1 or later
- Contest data in Excel format (questions stored in `.xlsx` files)

### Dependencies
- `github.com/gin-gonic/gin` - HTTP web framework
- `github.com/swaggo/gin-swagger` - Swagger documentation
- `github.com/swaggo/files` - Swagger file server
- `github.com/xuri/excelize/v2` - Excel file processing

### Generating Swagger Documentation
```bash
swag init -g cmd/api/main.go
```

## Data Structure

The API works with contest data organized in folders:
- Contest folder contains subjects
- Each subject folder contains chapters
- Each chapter folder contains an Excel file with questions
- Questions are randomly selected based on chapter requirements

## Test Caching

Once a test is generated for an officer-subject combination, it's cached and subsequent requests return the same test. This ensures consistency during the testing process.
