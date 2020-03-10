package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"warpin/internal/api"
	"warpin/internal/dao/localstore"
	"warpin/internal/service"

	"github.com/emicklei/go-restful"
)

func main() {
	notifDAO := localstore.New()
	svc := service.New(notifDAO)
	api := api.New(svc)

	restful.DefaultContainer.Add(api.NewHTTPAPI())
	restful.DefaultContainer.Add(api.NewWebsocketAPI())

	portNum := os.Getenv("SERVICE_PORT")
	if portNum == "" {
		portNum = "8787"
	}
	fmt.Println("warpin")
	fmt.Println("listening to port " + portNum)
	log.Fatal(http.ListenAndServe(":"+portNum, nil))
}
