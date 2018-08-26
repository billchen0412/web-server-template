package serviceB

import (
    "net/http"
)


func ProcessService1(w http.ResponseWriter) {
  w.Write([]byte("You are calling service 1 in B controller"))
}

func ProcessService2(w http.ResponseWriter) {
  w.Write([]byte("You are calling service 2 in B controller"))
}
