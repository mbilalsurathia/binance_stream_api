# binance_stream_api

Binance WebSocket Data Fetcher
This Go application allows you to fetch real-time streaming data and historical data from Binance's WebSocket endpoint.

Getting Started
Follow the steps below to get started with using this application:

Prerequisites
Go installed on your machine. You can download and install Go from the official website: https://golang.org/
Ensure you have a stable internet connection to fetch data from Binance's WebSocket endpoint.
Installation
Clone the repository to your local machine:
```
git clone https://github.com/your-username/binance-websocket-data-fetcher.git
```
Navigate to the project directory:
cd binance-websocket-data-fetcher
Run go mod tidy to ensure all dependencies are downloaded:
```
go mod tidy
```
Usage
Run the main Go file to start the application:
```
go run main.go
```
Once the application is running, you can access the following endpoints:
```
/stream: This endpoint provides real-time streaming data from Binance's WebSocket endpoint.
/historical_data: This endpoint allows you to fetch past historical data for your chart.
```
To access the streaming data, make a GET request to the /stream endpoint. The application will continuously stream real-time data from Binance.
To fetch historical data for your chart, make a GET request to the /historical_data endpoint. This will retrieve past data that you can use to visualize historical trends.
Example Usage
To access streaming data, open your web browser and navigate to http://localhost:8888/stream.
To fetch historical data, make a GET request to http://localhost:8888/historical_data using a tool like cURL or Postman.
