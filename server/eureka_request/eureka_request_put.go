package eureka_request

import (
	"github.com/go-martini/martini"
	"net/http"
	"strings"
	"github.com/martini-contrib/render"
	"github.com/ArthurHlt/microcos/eureka_client"
)

type EurekaRequestPut struct {
	EurekaRequest
}

func NewEurekaRequestPut(server *martini.ClassicMartini, eurekaClient *eureka_client.EurekaClient) *EurekaRequestPut {
	eurekaRequest := &EurekaRequestPut{}
	eurekaRequest.eurekaClient = eurekaClient
	eurekaRequest.server = server
	return eurekaRequest
}

func (this *EurekaRequestPut) requestHeartBeat(r render.Render, resp http.ResponseWriter, req *http.Request, params martini.Params) {
	values := []string{"apps", this.getAppId(params["appId"]), params["instanceId"]}
	path := strings.Join(values, "/")
	clientResp, err := this.eurekaClient.Put(path, nil)
	if err != nil {
		this.showError(err, r)
		return
	}
	resp.WriteHeader(clientResp.StatusCode)
	resp.Write(clientResp.Body)

}

func (this *EurekaRequestPut) SetRoutes(r martini.Router) {
	r.Put("/apps/:appId/:instanceId", this.requestHeartBeat)
}