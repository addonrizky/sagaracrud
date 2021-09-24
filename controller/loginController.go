package controller

import (
	"context"
	"encoding/json"
	"github.com/addonrizky/sagaracrud/constant"
	"github.com/addonrizky/sagaracrud/entity/entityself"
	"github.com/addonrizky/sagaracrud/repository/database"
	"github.com/addonrizky/sagaracrud/usecase"
	"github.com/addonrizky/sagaracrud/utility"
	"github.com/addonrizky/sagaracrud/validator"
	"net/http"
	//"fmt"
)

func Login(res http.ResponseWriter, req *http.Request) {
	mysqlDatabase := database.NewMysqlDatabase()
	ucl := usecase.NewLoginUsecase(mysqlDatabase)
	var request entityself.Login
	var err error
	ctx := req.Context()
	requestID := ctx.Value(constant.CtxKeyIdTransaction)
	requestBody := ctx.Value(constant.CtxKeyRequestBody)
	responseObject := entityself.NewResponse(requestID.(string))

	_ = json.Unmarshal([]byte(requestBody.(string)), &request)
	err = validator.ValidateRequest(req, request)

	if err != nil {
		responseObject.Code = "VE"
		responseObject.Desc = "Validation Error"
		responseObject.Data = err.Error()
		utility.SetHttpResponse(res, req, responseObject)
		return
	}

	ctx = context.WithValue(req.Context(), constant.CtxKeyRefnum, request.RequestRefnum)

	code, desc, jwtToken := ucl.Login(ctx, request.Username, request.Password)

	responseObject.Code = code
	responseObject.Desc = desc
	responseObject.Data = jwtToken

	utility.SetHttpResponse(res, req, responseObject)
	return
}
