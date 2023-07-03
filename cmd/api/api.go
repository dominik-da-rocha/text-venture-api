package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	v1 "rochatronic.net/text-venture-service/cmd/api/v1"
	"rochatronic.net/text-venture-service/cmd/model"
)

func Start(txt model.ServiceConfig) {
   router := mux.NewRouter()
   v1.Version(router);
   v1.GuestBook(router, txt.Database);
   
   log.Printf("Start listing on port %d", txt.Port)
   http.Handle("/", router)
   http.ListenAndServe(":" + strconv.Itoa(txt.Port), nil)
}

