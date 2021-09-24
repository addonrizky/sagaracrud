package database

import (
	"github.com/addonrizky/sagaracrud/entity/entitydatabase"
)

type Database interface {
	GetUserByUsername(username string) (entitydatabase.User, error, string)
	AddProduct(name string, desc string, price int, image string) (string, error, string)
	EditProduct(name string, desc string, price int, image string, id string) (string, error, string)
	SelectProduct(id string) (entitydatabase.Product, error, string)
	DeleteProduct(id string) (string, error, string)
}