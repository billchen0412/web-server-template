package main

import (
	"net/http"
	"web-server/controller"
	"log"
)

func main() {
  http.HandleFunc("/A/", controller.MyHttpHandlerA) // set the route path
	http.HandleFunc("/B/", controller.MyHttpHandlerB) // set the route path
  err := http.ListenAndServe(":9090", nil) // set the listening port
  if err != nil {
      log.Fatal("ListenAndServe: ", err)
  }
}
