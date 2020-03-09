# Warpin Service
## Feature
- Send notification 

    Send string message notification. 
    POST `{base url}/notifications` <br>
    Request body: 
    ```
    {"message":"string notification message"}
    ```

- Get all notification
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
    
- Real time notification
    initiate ws connection to `{base url}/connect` then listen to the connection for notification update. 

    I use [this](https://chrome.google.com/webstore/detail/simple-websocket-client/pfdhoblngboilpfeibdedpjgfnlcodoo) simple websocket client extension on chrome as websocket client.



## Build


## Deployment