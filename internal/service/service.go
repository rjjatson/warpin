package service

import (
	"net/http"
	"time"
	"warpin/internal/dao"
	"warpin/internal/model"

	"github.com/emicklei/go-restful"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// Service is handler for the API
type Service struct {
	notifDAO dao.NotificationDAO
	inbound  chan []byte
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // todo : remove debug
}

// New create new notification DAO
func New(dao dao.NotificationDAO) *Service {
	return &Service{
		notifDAO: dao,
	}
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
	response.WriteHeaderAndEntity(http.StatusOK, model.GetAllNotifResponse{Notifications: respMessages})
}

// Connect upgrade connection to websocket connection
func (svc *Service) Connect(request *restful.Request, response *restful.Response) {
	conn, err := upgrader.Upgrade(response.ResponseWriter, request.Request, nil)
	if err != nil {
		// todo log
		logrus.Error("unable to upgrade connection ", err)
		return
	}

	conn.WriteMessage(websocket.TextMessage, []byte("ok"))

	// todo build the client
}
