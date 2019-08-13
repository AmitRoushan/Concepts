package main

import (
	"bufio"
	"fmt"
	"k8s.io/apimachinery/pkg/util/json"
	"log"
	"net/http"
	"os"

	"github.com/emicklei/go-restful"
)

type GVK struct {
	ApiVersion string `json:"apiVersion"`
	Kind string `json:"kind"`
}

type TokenReviewRequest struct {
	GVK
	Spec struct {
		Token string `json:"token"`
	} `json:"spec"`
}

type TokenReviewResponse struct {
	GVK
	Status struct {
		Authenticated bool `json:"authenticated"`
		User UserInfo `json:"user,omitempty"`
	} `json:"status"`
}

type UserInfo struct {
	Username string `json:"username"`
	UID string `json:"uid"`
	Groups []string `json:"groups"`
}

type TokenInfo struct {
	Token string `json:"token"`
	User *UserInfo `json:"user"`
}

type WebHook struct {}

type WebHookResource struct {
	// normally one would use DAO (data access object)
	webhooks map[string]WebHook
}

func (u WebHookResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/web-hook").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML) // you can specify this per route as well

	ws.Route(ws.POST("/authenticate").To(u.authenticate))
	ws.Route(ws.POST("/authorize").To(u.authorize))
	ws.Route(ws.GET("listAuthenticator").To(u.listAuthenticator))
	ws.Route(ws.GET("listAuthorizer").To(u.listAuthorizer))

	container.Add(ws)
}

func (u WebHookResource) authenticate(request *restful.Request, response *restful.Response) {
	fmt.Println("Received token review request", )
	tokenReview := &TokenReviewRequest{}
	err := request.ReadEntity(&tokenReview)
	if err != nil {
		response.WriteHeaderAndEntity(http.StatusAccepted, getTokenReviewErroRes())
	}

	fmt.Println("Received token review request as", tokenReview)

	tokenInfos, err := getTokenList()
	if err != nil {
		fmt.Println("failed in serialization")
		response.WriteHeaderAndEntity(http.StatusAccepted, getTokenReviewErroRes())
		return
	}

	for _, tokenInfo := range tokenInfos {
		if tokenInfo.Token == tokenReview.Spec.Token {
			fmt.Println(fmt.Sprintf("Send success token response as %+v", getTokenSuccessRes(tokenInfo.User)))
			response.WriteHeaderAndEntity(http.StatusCreated, getTokenSuccessRes(tokenInfo.User))
			return
		}
	}

	fmt.Println(fmt.Sprintf("Send failure token response as %v", getTokenReviewErroRes()))
	response.WriteHeaderAndEntity(http.StatusAccepted, getTokenReviewErroRes())
}

func getTokenReviewErroRes() TokenReviewResponse {
	tokenResponse := TokenReviewResponse{}
	tokenResponse.ApiVersion = "authentication.k8s.io/v1beta1"
	tokenResponse.Kind = "TokenReview"
	tokenResponse.Status.Authenticated=false
	return tokenResponse
}

func getTokenSuccessRes (user *UserInfo) TokenReviewResponse {
	tokenResponse := TokenReviewResponse{}
	tokenResponse.ApiVersion = "authentication.k8s.io/v1beta1"
	tokenResponse.Kind = "TokenReview"
	tokenResponse.Status.Authenticated=true
	tokenResponse.Status.User = *user
	return tokenResponse
}

func getTokenList() ([]TokenInfo, error)  {
	file, err := os.Open(`tokenFile`)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	tokenList := make([]TokenInfo, 0)

	i := 0
	var tokenInfo *TokenInfo
	for scanner.Scan() {
		i++
		b := scanner.Bytes()
		tokenInfo = new(TokenInfo)
		err := json.Unmarshal(b, tokenInfo)
		if err != nil {
			fmt.Println("failed to unmarshal", err)
			continue
		}
		tokenList = append(tokenList, *tokenInfo)
	}

	return tokenList, nil
}

func (u WebHookResource) authorize(request *restful.Request, response *restful.Response) {
}

func (u WebHookResource) listAuthenticator(request *restful.Request, response *restful.Response) {
	fmt.Println("Received listAuthenticator request", )
	tokenList, err := getTokenList()
	if err != nil {
		response.WriteHeaderAndEntity(http.StatusBadRequest, "failed to list")
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, tokenList)
}

func (u WebHookResource) listAuthorizer(request *restful.Request, response *restful.Response) {

}


func main() {
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	u := WebHookResource{map[string]WebHook{}}
	u.Register(wsContainer)

	log.Print("start listening on 192.168.56.50:8080")

	log.Fatal(http.ListenAndServeTLS(`192.168.56.50:8080`, `/etc/kubernetes/pki/apiserver.crt`,
		`/etc/kubernetes/pki/apiserver.key`, wsContainer))
}