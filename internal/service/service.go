package service

import (
	"net/http"
	"warpin/internal/dao"

	"github.com/emicklei/go-restful"
)

// Service is handler for the API
type Service struct {
	notifDAO *dao.NotifMessage
	inbound  chan []byte
}

// HandleGetAll is the handler of get all notif API
func (svc *Service) HandleGetAll(request *restful.Request, response *restful.Response) {

}

// HandleStore is the handler of store notif API
func (svc *Service) HandleStore(request *restful.Request, response *restful.Response) {

	response.WriteHeader(http.StatusOK)
}
