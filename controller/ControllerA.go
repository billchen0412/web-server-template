package controller

import (
  "fmt"
  "net/http"
  "bytes"
  "strings"
	"web-server/service/serviceA"
)

func MyHttpHandlerA(w http.ResponseWriter, r *http.Request) {
  fmt.Println("method:", r.Method) // get the request method

	buf := new(bytes.Buffer) // allocate memory for buf
  buf.ReadFrom(r.Body) // read from the body which has io.ReadCloser type
  body := buf.String() // convert buf to string
  fmt.Println("body:", body) // get the request body
  fmt.Println("Url = ", r.URL.Path)

  cookie, err := r.Cookie("myCookie") // get the request body
	if err == nil {
		fmt.Println("Domain:", cookie.Domain)
		fmt.Println("Expires:", cookie.Expires)
		fmt.Println("Name:", cookie.Name)
		fmt.Println("Value:", cookie.Value)
	}

  // process the service based on URL
  slashIndex := strings.LastIndex(r.URL.Path, "/")
  if slashIndex != -1 {
    serviceName := r.URL.Path[slashIndex + 1:len(r.URL.Path)]
    fmt.Println("serviceName:", serviceName)
    if (serviceName == "service1") {
      serviceA.ProcessService1(w)
    } else if (serviceName == "service2") {
      serviceA.ProcessService2(w)
    }
  }
}
