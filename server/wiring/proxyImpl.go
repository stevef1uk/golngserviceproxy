package wiring



import (



	"github.com/go-openapi/runtime/middleware"
	//"github.com/tylerb/graceful"
	"goproxy/client/client/operations"
	server "goproxy/server/restapi/operations"
	"goproxy/server/models"
	//"time"
	apiclient "goproxy/client/client"

	"time"
)



func DoIt(params server.GetAccounts2Params) middleware.Responder {
	//myparams := &operationsc.GetAccountsParams{ID: 1}
	ret := make(models.GetAccounts2OKBody,1)
	myparams := operations.NewGetAccounts2Params()
	myparams.ID = params.ID
	myparams.WithTimeout(time.Second * 2)
	//fmt.Printf("I am doing it")

	resp, err := apiclient.Default.Operations.GetAccounts2(myparams)
	if err != nil {
		//log.Fatal(err)
		return server.NewGetAccounts2BadRequest()
	}
	ret = resp.Payload

	return server.NewGetAccounts2OK().WithPayload(ret)

}
