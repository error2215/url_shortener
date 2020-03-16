# simple_mongodb

You can run this app by 2 ways:
  1) Run MongoDB separately in Docker container and app by yourself
  2) Run MongoDB and app by docker-compose
 
## 1) Run MongoDB separately in Docker container and app by yourself
  * Run mongo using command `make mongo_docker`
  * Run app by command `go run main.go`
    
## 2) Run all in Docker containers
   * Run command `make all_docker`
   
## After application run you can get access to it's API by address `localhost:3034`
   1) POST `/api/v1/url_shortener` (`{"url": "https://google.com"}`) - you will receive json response with alias like `5d2399ef96fb765873a24bae`, you can use it for redirect
   2) GET `/rd/{alias}` - use this endpoint to redirect 
