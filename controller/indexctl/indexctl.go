package indexctl

import (
    "net/http"
)

func GetIndexPage(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hello world from apisvr!\n"))
}
