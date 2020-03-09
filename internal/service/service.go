package service

import (
	"net/http"
	"time"
	"warpin/internal/dao"
	"warpin/internal/model"

	"github.com/emicklei/go-restful"
)

// Service is handler for the API
type Service struct {
	notifDAO *dao.NotifMessage
	inbound  chan []byte
}

// HandleStore is the handler of store notif API
func (svc *Service) HandleStore(request *restful.Request, response *restful.Response) {
	reqMessage := new(model.SendNotifRequest)
	err := request.ReadEntity(&reqMessage)
	if err != nil {
		// todo log
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	err = svc.notifDAO.Store(reqMessage.Message, time.Now().UTC())
	if err != nil {
		// todo log
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	// todo log
	response.WriteHeader(http.StatusOK)
}

// HandleGetAll is the handler of get all notif API
func (svc *Service) HandleGetAll(request *restful.Request, response *restful.Response) {
	storedMessages, err := svc.notifDAO.GetAll()
	if err != nil {
		// todo log
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	respMessages := make([]model.Notification, 0)

	for _, sm := range storedMessages {
		respMessages = append(respMessages, model.Notification{
			Message: sm.Message,
			Time:    sm.Time.Format(time.RFC3339),
		})
	}

	// todo log
	response.WriteHeaderAndEntity(http.StatusOK, respMessages)
}
