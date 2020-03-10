# Warpin Service
## Feature
### 1. Send notification 

Send string message notification. 
POST `{base url}/notifications` <br>
Request body: 
```
{"message":"string notification message"}
```

cURL example
```
curl --location --request POST 'localhost:8787/notifications' \
--header 'Content-Type: application/json' \
--data-raw '{
	"message":"warpin warpin"
}'
```



### 2. Get all notification
Collect all sent notification
GET `{base url}/notifications` <br>
Response body:
``` 
{
         notifications": [{
   			"message": "string 1",
   			"time": "2020-03-09T10:10:19+00:00"
   		},
   		{
   			"message": "string 2",
   			"time": "2020-03-09T10:10:20+00:00"
   		}]
}
```
cURL example
```
curl --location --request GET 'localhost:8787/notifications'
```
    
### 3. Real time notification
initiate ws connection to `{base url}/connect` e.g.: ws://localhost:8787/connect then listen to the connection for notification update. you will get success message if connection established successfully. 

I use [this](https://chrome.google.com/webstore/detail/simple-websocket-client/pfdhoblngboilpfeibdedpjgfnlcodoo) simple websocket client extension on chrome as websocket client.



## Build
clone the root folder to your $GOPATH


`go build -o service`

`./service`

## Test

`go test ./...`

## Dockerized Build
