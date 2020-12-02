package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/x-debug/Go-000/Week02/service"
	"log"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/users", func(resp http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			sid := req.FormValue("uid")
			uid, err := strconv.Atoi(sid)
			if err != nil {
				resp.WriteHeader(http.StatusBadRequest)
				return
			}
			user, err := service.GetUser(uint64(uid))
			if err != nil {
				resp.WriteHeader(http.StatusInternalServerError)
				fmt.Printf("%+v", err)
				err := errors.Cause(err)
				_, _ = resp.Write([]byte(fmt.Sprintf("%s", err)))
				return
			}

			_, _ = resp.Write([]byte(fmt.Sprintf("%s", user)))
		}

		resp.WriteHeader(http.StatusNotImplemented)
	})

	log.Fatalln(http.ListenAndServe(":1234", nil))
}
