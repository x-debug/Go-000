package api

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/x-debug/Go-000/Week02/service"
	"net/http"
	"strconv"
)

func UserHandler(resp http.ResponseWriter, req *http.Request) {
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

		resp.WriteHeader(http.StatusOK)
		_, _ = resp.Write([]byte(fmt.Sprintf("%s", user)))
		return
	}

	resp.WriteHeader(http.StatusNotImplemented)
}
