package service

import (
	"encoding/json"
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
	notifDAO  dao.NotificationDAO
	wsClients []*websocket.Conn
	broadcast chan []byte
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // todo : remove debug
}

// New create new notification DAO
func New(dao dao.NotificationDAO) *Service {
	svc := &Service{
		notifDAO:  dao,
		broadcast: make(chan []byte, 0),
	}
	go svc.runBroadcast()
	return svc
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
	timestamp := time.Now().UTC()
	err = svc.notifDAO.Store(reqMessage.Message, timestamp)
	if err != nil {
		// todo log
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	go func() {
		b, _ := json.Marshal(model.Notification{
			Message: reqMessage.Message,
			Time:    timestamp.Format(time.RFC3339),
		})
		svc.broadcast <- b
	}()

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

	conn.WriteMessage(websocket.TextMessage, []byte(`{"type":"connect_notif","message":"ok"}`))

	svc.wsClients = append(svc.wsClients, conn)
}

func (svc *Service) runBroadcast() {
	for {
		select {
		case msg := <-svc.broadcast:
			for _, con := range svc.wsClients {
				con.WriteMessage(websocket.TextMessage, msg)
			}
		}
	}
}
