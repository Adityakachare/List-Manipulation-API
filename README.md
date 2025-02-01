# List Manipulation API

This is a Go-based API that maintains a list of integers and manipulates it based on the sign of an input number. The program exposes an API that allows users to append or modify the list. The behavior is determined by the sign of the input number and the existing values in the list.

## Functionality

- **List Initialization**: The list starts as an empty list of integers.
- **API Endpoint**: The program exposes a single API endpoint to accept a number and manipulate the list.
- **Same Sign**: If the sign of the input number matches the sign of the integers in the list, the input number is appended to the list.
- **Opposite Sign**: If the sign of the input number is opposite to the sign of the integers in the list, the program reduces the values in the list in a FIFO (First In First Out) manner until the input number is exhausted.

### Example Workflow:

1. **Input**: 5 
   - **Updated List**: [5]

2. **Input**: 10
   - **Updated List**: [5, 10]

3. **Input**: -6
   - **Updated List**: [9] 
   (5 is removed and the remainder 1 is subtracted from 10)

4. **Input**: 100
   - **Updated List**: [9, 100]

5. **Input**: -90
   - **Updated List**: [19]
   (9 is removed, and 81 is subtracted from 100)

## How to Run

### Prerequisites

Make sure you have the following tools installed on your system:
- [Go](https://golang.org/dl/) version 1.18 or higher
- [Gin](https://github.com/gin-gonic/gin) (Used for setting up the API)
- [Testify](https://github.com/stretchr/testify) (Used for unit testing)

### 1. Clone the repository

Clone the project to your local machine:
```
https://github.com/Adityakachare/List-Manipulation-API.git
cd list-manipulation-api
```

### 2. Install Dependencies

Install the required Go modules:
```bash
go mod tidy
```

### 3. Run the Go Program

Start the server by running the `main.go` file:
```bash
go run main.go
```

The server will start and listen on port `8080`.

### 4. API Usage

- **Endpoint**: `/manipulate`
- **Method**: `GET`
- **Query Parameter**: `number` (integer to manipulate the list)

#### Example Requests:
- `GET http://localhost:8080/manipulate?number=5`
- `GET http://localhost:8080/manipulate?number=-6`

### 5. Running Unit Tests

To run the unit tests, use the following command:

```bash
go test -v
```

This will run all the tests and output detailed information for each test case.

## Code Overview

- **`main.go`**:
  - Implements the logic for manipulating the list of integers.
  - Sets up a Gin web server and exposes the API endpoint.
  - The `manipulateList` function handles appending or modifying the list based on the sign of the input number.
  
- **`main_test.go`**:
  - Contains unit tests using the Testify framework.
  - Tests various scenarios like appending, reducing the list based on input, and edge cases.

## Edge Cases Considered

- Input of `0`: No changes are made to the list.
- If the number to be reduced exceeds the list values, it will completely remove the older elements, and if any quantity is still left, it will be appended back to the list.

---

**Author**: Aditya Ashok Kachare  
**Date**: February 2025  
