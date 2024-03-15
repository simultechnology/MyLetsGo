# MyProgLog

## Overview

MyProgLog is a Go-based server application that provides a simple yet powerful logging mechanism along with an HTTP interface for creating and reading log records. Leveraging the `gorilla/mux` router, it offers structured HTTP routing for efficient data handling. This project is ideal for applications requiring fast, in-memory logging with the capability to interact through a RESTful API.

## Installation

To install MyProgLog, ensure you have Go installed on your system (version 1.21.6 or higher is recommended). Follow these steps:

1. Clone the repository to your local machine:

    ```bash
    git clone https://github.com/simultechnology/myproglog.git
    cd myproglog
    ```

2. Build the application:

    ```bash
    go build -o myproglog .
    ```

## Running the Server

To start the server on the default port (`58888`), run:

```bash
./myproglog
```

The server will now be listening for HTTP requests to produce and consume log records.


## Usage

### Producing a Record

To add a new record:

```
curl -X POST localhost:58888 -d '{"record": {"value": "your_log_message_here"}}'
```

This will append a new record to the log and return the offset of the newly added record.

### Consuming a Record

To read a record by its offset:

```bash
curl -X GET localhost:58888?offset=0
```
Replace 0 with the desired offset. The response will include the record's value and its offset.

### Result

```
{
  "record": {
    "value": "Now",
    "offset": 2
  },
  "value": "Now"
}
```

## Features

- Efficient In-Memory Logging: Fast access and append operations with concurrency control.
- HTTP API Interface: RESTful endpoints for creating and reading log records.
- Flexible and Simple: Easy to integrate with other applications or extend for various logging needs.


