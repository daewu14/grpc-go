# Order Service Testing

This project contains unit tests for the **Order Service** in a gRPC-based application.It uses a mock implementation of the OrderService client for testing purposes.

## Run Application
1. Open 2 terminal windows.
2. In the first terminal, run the gRPC server:
    ```bash
    go run main.go grpc
    ```
    this app running on port 9000.
3. In the second terminal, run the HTTP server:
    ```bash
    go run main.go http
    ```
   this app running on port 8000.

## API Testing Example (HTTP Test)
The following command demonstrates test execution and displays detailed results:

```curl
curl --location 'http://localhost:8000/orders' \
--header 'Content-Type: application/json' \
--data '{
    "CustomerID": 42,
    "ProductID": 32,
    "Quantity": 3
}'
```

## Test Cases

Each test case provides:
- **Name:** A unique identifier for the test.
- **Mock Response:** The predefined gRPC response.
- **Mock Error:** The predefined gRPC error returned by the mock client.
- **Expected Status Code:** The HTTP response code expected from the handler.
- **Request Body:** The input JSON test payload.

## Structure of a Test Case

```go
{
	name: "success",
	mockResponse: &orders.CreateOrderResponse{
		Status: "success",
		Order: &orders.Order{
			OrderID:	1,
			CustomerID: 3,
			ProductID:  4,
			Quantity:   6,
		},
	},
	requestBody:		`{"CustomerID":42, "ProductID": 32, "Quantity": 3}`,
	mockError:		  nil,
	expectedStatusCode: http.StatusOK,
}
```

## Testing Frameworks

- [**Testify**](https://github.com/stretchr/testify): Used for assertions to validate test results.

## How to Run Tests

Run the following command in the terminal to execute the tests:

```bash
go test -v ./test
```

## Testing Example (Unit Test)

The following command demonstrates test execution and displays detailed results:

```bash
go test -v ./test
```

Example output:
```aiignore
=== RUN   TestNewCreateOrderHandler
=== RUN   TestNewCreateOrderHandler/success
=== RUN   TestNewCreateOrderHandler/validation_error
=== RUN   TestNewCreateOrderHandler/grpc_error
2025/04/20 19:27:25 client error: grpc error
=== RUN   TestNewCreateOrderHandler/invalid_grpc_response
2025/04/20 19:27:25 client error: data is nil
--- PASS: TestNewCreateOrderHandler (0.00s)
    --- PASS: TestNewCreateOrderHandler/success (0.00s)
    --- PASS: TestNewCreateOrderHandler/validation_error (0.00s)
    --- PASS: TestNewCreateOrderHandler/grpc_error (0.00s)
    --- PASS: TestNewCreateOrderHandler/invalid_grpc_response (0.00s)
PASS
ok      grpc/test       (cached)
```
