package routers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lairdnote/oais-web/handler/index"
)

func Start()  {
	r := mux.NewRouter()
	r.HandleFunc("/", index.Welcome)
	log.Fatal(http.ListenAndServe(":8000", r))
}
