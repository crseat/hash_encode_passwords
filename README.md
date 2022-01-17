# Hash and Encode a Password String

## Overview
HTTP server that listens on a given port following Hexagonal architecture (*ports & adapters*)
It supports multiple connections simultaneously, and provides the following endpoints:

 - `/hash`
    - input: form field named password using POST to provide the value to convert to a base64 encoded string of the SHA512 hash after a 5 second delay. Password must be between 8 and 50 characters, and include a capital letter.
    - output: An incrementing identifier returns immediately.
 
  - `/hash/id`
    - input: id of previously hashed password using GET
    - output: base64 encoded string of the SHA512 hash of the identified password.

 - `/stats`
    - input: None. GET
    - output: JSON like `{"total": "1", "averagfe": "123"}`. Total being the number of POST requests, and average being the average time taken to process those in microseconds.
 
 - `/shutdown`
    - input: None. 
    - output: Server will gracefully shutdown after waiting for all active requests to complete.
    

## Installation
### Setup

Clone this repo and launch server in a terminal with go run main.go port#
  - ex: go run main.go 8000. The server automatically starts on localhost

In a new terminal send your POST and GET requests.
  - curl —data “password=angryMonkey” http://localhost:8080/hash
  - curl http://localhost:8080/hash/1
  - curl http://localhost:8080/stats
  - curl http://localhost:8080/shutdown

## Architecture

Code architecture follows hexagonal architecture principles, also known as *ports and adapters*.

This architecture is divided in three main layers:

- **Application**:  The outer layer. Handlers and all I/O related stuff (web framework, DB, ...). Anything that can change by an "external" cause (not by your decision), is in this layer. 

- **Service**: Use cases. Actions triggered by API calls, represented by application services. It includes repositories specific interfaces, known as *adapters*.

- **Domain**: Inner layer. Business logic and rules goes here. Repositories Interfaces, known as *ports*, belongs to this layer.

I have also included an data transfer object (dto) abstraction layer. This allows us to control exactly what is passed back to the client.

## Testing 

Tests have been built with golangs built in testing package. In root directory of repo run `go test ./...`

## Considerations made

- Because we are dealing with passwords I was thinking of maybe doing some TLS (transport layer security) but since it's all localhost it didn't seem necessary and I thought that would overcomplicate things. 
- I decided to include the 5 second delay in the stats average time calculation because in my opinion that represents processing time.
- I considered several approaches to the routing including some regex strategies, but in the end I thought a "no router" approach was the most straight forward.
- Because only the standard library was allowed I did not hook up a databsase so the hashes are lost once the server is shutdown. However, because of the hexagonal architecture, hooking up a DB would be trivial. I could have saved the hashes to a file and read them back in when the server started, but I could not in good conscience save password hashes to a local file even if just for an example program. 
- If I was able to use 3rd party libraries I would have included more in depth testing of the services using gomock.
- If I was to iterate on this I would add more debug logging functionality, and more error checking. 
