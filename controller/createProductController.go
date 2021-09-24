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
	"strconv"
	//"fmt"
	"encoding/base64"
	"strings"
)

func CreateProduct(res http.ResponseWriter, req *http.Request) {
	var mysqlDatabase = database.NewMysqlDatabase()
	var uc usecase.ProductUsecase = usecase.NewProductUsecase(mysqlDatabase)
	var request entityself.Create
	var err error
	var price int
	isValidImage := false
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

	imageB64Object := strings.Split(request.Image, ",")

	if len(imageB64Object) == 2 && (imageB64Object[0] == "data:image/jpeg;base64" || imageB64Object[0] == "data:image/png;base64") {
		isValidImage = true
	}

	if !isValidImage {
		responseObject.Code = "VE"
		responseObject.Desc = "Validation Error"
		responseObject.Data = "File uploaded expected as an image jpeg or png"
		utility.SetHttpResponse(res, req, responseObject)
		return
	}

	imageByte, err := base64.StdEncoding.DecodeString(imageB64Object[1])
	if err != nil {
		responseObject.Code = "VE"
		responseObject.Desc = "Validation Error"
		responseObject.Data = "File uploaded expected as an image jpeg or png"
		utility.SetHttpResponse(res, req, responseObject)
		return
	}

	ctx = context.WithValue(req.Context(), constant.CtxKeyRefnum, request.RequestRefnum)

	price, err = strconv.Atoi(request.Price)
	code, err, desc := uc.CreateProduct(ctx, request.Name, request.Desc, price, imageByte)

	responseObject.Code = code
	responseObject.Desc = desc
	//responseObject.Data = nil

	utility.SetHttpResponse(res, req, responseObject)
	return
}
