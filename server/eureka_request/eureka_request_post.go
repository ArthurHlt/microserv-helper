package eureka_request

import (
	"github.com/go-martini/martini"
	"net/http"
	"io/ioutil"
	"strings"
	"github.com/martini-contrib/render"
	"github.com/ArthurHlt/microcos/eureka_client"
)

type EurekaRequestPost struct {
	EurekaRequest
}

func NewEurekaRequestPost(server *martini.ClassicMartini, eurekaClient *eureka_client.EurekaClient) *EurekaRequestPost {
	eurekaRequest := &EurekaRequestPost{}
	eurekaRequest.eurekaClient = eurekaClient
	eurekaRequest.server = server
	return eurekaRequest
}

func (this *EurekaRequestPost) requestRegisterApp(r render.Render, resp http.ResponseWriter, req *http.Request, params martini.Params) {
	values := []string{"apps", this.getAppId(params["appId"])}
	path := strings.Join(values, "/")
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		this.showError(err, r)
		return
	}
	clientResp, err := this.eurekaClient.Post(path, body)
	if err != nil {
		this.showError(err, r)
		return
	}
	resp.WriteHeader(clientResp.StatusCode)
	resp.Write(clientResp.Body)

}

func (this *EurekaRequestPost) SetRoutes(r martini.Router) {
	r.Post("/apps/:appId", this.requestRegisterApp)
}