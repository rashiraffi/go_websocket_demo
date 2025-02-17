# README

## Backend Service (BE)

This backend service follows the clean architecture pattern, separating concerns into handler, service, and model layers. Additionally, it includes local packages for authentication, middleware, and logging setup.

### Execution Flow

The execution begins from `cmd/server.go`, where the application:

- Loads environment variables from a `.env` file.
- Sets up the logger.
- Initializes the layers of the clean architecture.
- Calls the router layer to obtain the Fiber application instance.
- Starts listening for requests.

### Layers

#### Handler Layer

The handler layer is responsible for:

- Receiving and validating incoming requests.
- Invoking the service layer for processing.
- Handling WebSocket connections using channels. The service layer passes a channel, and the output from this channel is sent to the client via WebSocket.

#### Service Layer

The service layer implements the business logic:

- For login operations, it queries the model layer for user details.
- For stock price retrieval, it connects to an external API WebSocket and streams price updates through a channel to the handler, which then sends them to the client via WebSocket.

#### Model Layer

Currently, the model layer is a placeholder. There is no database connection; instead, user details are stored in an in-memory map.

### External Packages Used

This project utilizes several external libraries:

- `github.com/gofiber`:
  - `jwt` for authentication.
  - `websocket` for WebSocket implementation.
- `github.com/gorilla/websocket`: For WebSocket connections with external services.
- `github.com/uber-go/zap`: For structured logging.
- `github.com/joho/godotenv`: For loading environment variables from a `.env` file.

## External Backend Service

The external backend service provides stock market data to this backend service. It shares a similar architecture.

### Local Packages

#### StockExchange

The `StockExchange` package simulates stock market price updates:

- A predefined list of stocks with an opening price is maintained.
- A `*time.Ticker` generates price variations every second.
- Price updates are sent to the service layer through a channel.

## References Used

The following documentation sources were referenced during development:

- [Fiber WebSocket Documentation](https://docs.gofiber.io/contrib/websocket/)
- [Joho Godotenv](https://github.com/joho/godotenv)
- [Uber Go Zap](https://github.com/uber-go/zap)
- [Gorilla WebSocket](https://github.com/gorilla/websocket)
