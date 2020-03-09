package api

import (
	"net/http"
	"warpin/internal/model"
	"warpin/internal/service"

	"github.com/emicklei/go-restful"
)

// API manages service API
type API struct {
	svc *service.Service
}

// New creates mew API
func New(svc *service.Service) *API {
	return &API{
		svc: svc,
	}
}

// NewWebservice create new webservice
func (api *API) NewWebservice() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/notifications").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

	ws.Route(ws.POST("").
		To(api.svc.HandleStore).
		Reads(model.SendNotifRequest{}).
		Returns(http.StatusOK, "", nil))

	ws.Route(ws.GET("").
		To(api.svc.HandleGetAll).
		Writes(model.GetAllNotifResponse{}).
		Returns(http.StatusOK, "", nil))

	return ws
}
