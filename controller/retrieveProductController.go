package controller

import (
	"context"
	"github.com/addonrizky/sagaracrud/constant"
	"github.com/addonrizky/sagaracrud/entity/entityself"
	"github.com/addonrizky/sagaracrud/repository/database"
	"github.com/addonrizky/sagaracrud/usecase"
	"github.com/addonrizky/sagaracrud/utility"
	"net/http"
	//"fmt"
)

func RetrieveProduct(res http.ResponseWriter, req *http.Request) {
	var mysqlDatabase = database.NewMysqlDatabase()
	var uc usecase.ProductUsecase = usecase.NewProductUsecase(mysqlDatabase)
	ctx := req.Context()
	requestID := ctx.Value(constant.CtxKeyIdTransaction)
	responseObject := entityself.NewResponse(requestID.(string))

	token := req.URL.Query().Get("tok")
	productID := req.URL.Query().Get("product_id")

	if token == "" {
		responseObject.Code = "VE"
		responseObject.Desc = "Validation Error, token not exist"
		utility.SetHttpResponse(res, req, responseObject)
		return
	}
	if productID == "" {
		responseObject.Code = "VE"
		responseObject.Desc = "Validation Error, product id not exist"
		utility.SetHttpResponse(res, req, responseObject)
		return
	}

	request := &entityself.Get{Token: token, ProductID: productID}

	ctx = context.WithValue(req.Context(), constant.CtxKeyRefnum, "")

	code, err, desc, product := uc.RetrieveProduct(ctx, request.ProductID)

	if err != nil {
		responseObject.Code = "GE"
		responseObject.Desc = "General Error"
		utility.SetHttpResponse(res, req, responseObject)
		return
	}

	responseObject.Code = code
	responseObject.Desc = desc

	if code == "00" {
		responseObject.Data = product
	}

	utility.SetHttpResponse(res, req, responseObject)
	return
}
