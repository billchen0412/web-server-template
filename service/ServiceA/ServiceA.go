package serviceA

import (
    "net/http"
  	"web-server/dao"
)

func ProcessService1(w http.ResponseWriter) {
  var u dao.UserInfo
  w.Write([]byte("You are calling service 1 in A controller"))
}

func ProcessService2(w http.ResponseWriter) {
  w.Write([]byte("You are calling service 2 in A controller"))
}
