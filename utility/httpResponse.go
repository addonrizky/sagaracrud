package utility

import (
	"encoding/json"
	"fmt"
	"github.com/addonrizky/sagaracrud/entity/entityself"
	"net/http"
)

func SetHttpResponse(res http.ResponseWriter, req *http.Request, response *entityself.Response) {
	fmt.Println(response)
	res.Header().Set("Content-Security-Policy", "default-src 'self'")
	res.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(res).Encode(response)

	//start := context.Get(req, "start").(time.Time)
	//end := time.Now()
	//log.LogDebug("#" + constant.ServerName + "_RESP_TIME: " + (end.Sub(start)).String())
}
