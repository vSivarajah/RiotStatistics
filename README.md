# RiotStatistics


# How to get started
1. Install Go
2. Download all dependencies for the project
````
go get -u
````
3. Export RIOTAPI_KEY
````
Linux/MacOS
export RIOTAPI_KEY=<your api key>
export KAFKA_BOOTSTRAP_SERVERS=localhost:9092 

Windows
set RIOTAPI_KEY=<your api key>
set KAFKA_BOOTSTRAP_SERVERS=localhost:9092

````
4. Start the http server
````
go run main.go
````