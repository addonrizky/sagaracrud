package usecase

import (
	"context"
	"github.com/addonrizky/sagaracrud/entity/entitydatabase"
	"github.com/addonrizky/sagaracrud/repository/database"
	//"github.com/addonrizky/sagaracrud/utility"
	"fmt"
	"github.com/addonrizky/sagaracrud/constant"
	//"encoding/base64"
	"os"
)

var (
	dbProduct database.Database
)

type ProductUsecase interface {
	CreateProduct(ctx context.Context, name string, desc string, price int, image []byte) (string, error, string)
	UpdateProduct(ctx context.Context, name string, desc string, price int, image []byte, id string) (string, error, string)
	RetrieveProduct(ctx context.Context, id string) (string, error, string, entitydatabase.Product)
	DeleteProduct(ctx context.Context, id string) (string, error, string)
	//UpdateProduct(ctx context.Context, username string) (string, string, string)
	//DeleteProduct(ctx context.Context, username string) (string, string, string)
}

type usecaseProd struct{}

func NewProductUsecase(database database.Database) ProductUsecase {
	dbProduct = database
	return &usecaseProd{}
}

func (*usecaseProd) CreateProduct(ctx context.Context, name string, desc string, price int, image []byte) (string, error, string) {
	var (
		err  error
		code string
	)

	//upload product first    ***************************************//
	filename := name + ".jpg"

	f, err := os.Create("image_product/" + filename)
	if err != nil {
		fmt.Println(err)
		return constant.RCImageFail0, nil, constant.RCImageFailDesc
	}
	defer f.Close()

	if _, err := f.Write(image); err != nil {
		return constant.RCImageFail1, nil, constant.RCImageFailDesc
	}
	if err := f.Sync(); err != nil {
		return constant.RCImageFail2, nil, constant.RCImageFailDesc
	}
	// *********************************************************** //

	_, err, code = dbProduct.AddProduct(name, desc, price, filename)

	if err != nil {
		fmt.Println(err)
		return code, err, ""
	}

	return constant.RCSuccess, nil, constant.RCSuccessDesc
}

func (*usecaseProd) UpdateProduct(ctx context.Context, name string, desc string, price int, image []byte, id string) (string, error, string) {
	var (
		err  error
		code string
	)

	filename := ""

	/*upload product first    ***************************************
		dec, err := base64.StdEncoding.DecodeString(image)
	    if err != nil {
	        panic(err)
	    }

		filename := name+".jpg"

	    f, err := os.Create("image_product/"+filename)
	    if err != nil {
	        panic(err)
	    }
	    defer f.Close()

	    if _, err := f.Write(dec); err != nil {
	        panic(err)
	    }
	    if err := f.Sync(); err != nil {
	        panic(err)
	    }
		// *********************************************************** */

	_, err, code = dbProduct.EditProduct(name, desc, price, filename, id)

	if err != nil {
		fmt.Println(err)
		return code, err, ""
	}

	return constant.RCSuccess, nil, constant.RCSuccessDesc
}

func (*usecaseProd) RetrieveProduct(ctx context.Context, id string) (string, error, string, entitydatabase.Product) {
	product, err, code := dbProduct.SelectProduct(id)

	if err != nil {
		fmt.Println(err)
		return code, err, "", product
	}

	if code != "00" {
		return code, nil, "Retrieval Not Success", product
	}

	return constant.RCSuccess, nil, constant.RCSuccessDesc, product
}

func (*usecaseProd) DeleteProduct(ctx context.Context, id string) (string, error, string) {
	var (
		err  error
		code string
	)

	_, err, code = dbProduct.DeleteProduct(id)

	if err != nil {
		fmt.Println(err)
		return code, err, ""
	}

	if code != "00" {
		return code, nil, "Deletion Not Success"
	}

	return constant.RCSuccess, nil, constant.RCSuccessDesc
}
