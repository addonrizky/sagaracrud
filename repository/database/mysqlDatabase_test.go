package database

import (
    "testing"
	"github.com/DATA-DOG/go-sqlmock"
	//"fmt"
	"errors"
	"github.com/addonrizky/sagaracrud/entity/entitydatabase"
	"github.com/stretchr/testify/assert"
	"strconv"
)

var (
	mock sqlmock.Sqlmock
)

var product = &entitydatabase.Product{
	Id:    1,
	Name:  "Momo",
	Desc: "Momoroo baby product",
	Price: 200000,
	Image : "momoroo.jpg",
}

func initMock() {
	dbms, mock, err = sqlmock.New()
	//fmt.Println(dbms,"",err)
}

func TestShouldInsertData(t *testing.T) {
	initMock()
	db := NewMysqlDatabase()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	//defer db.Close()

	mock.ExpectExec("insert into product").WithArgs("name", "description", 10000, "apa.jpg").WillReturnResult(sqlmock.NewResult(1, 1))

	// now we execute our method
	if _, err, _ = db.AddProduct("name", "description", 10000, "apa.jpg"); err != nil {
		//fmt.Println(err)
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldUpdateData(t *testing.T) {
	initMock()
	db := NewMysqlDatabase()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	//defer db.Close()

	//mock.ExpectExec("insert into product").WithArgs("name", "description", 10000, "apa.jpg").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("UPDATE product SET name = \\?, description = \\?, price = \\?, image = \\? WHERE id = \\?").WithArgs(product.Name, product.Desc, product.Price, product.Image, strconv.Itoa(int(product.Id))).WillReturnResult(sqlmock.NewResult(0, 1))
	
	// now we execute our method
	if _, err, _ = db.EditProduct(product.Name, product.Desc, product.Price, product.Image, strconv.Itoa(int(product.Id))); err != nil {
		//fmt.Println(err)
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldRetrieveData(t *testing.T) {
	initMock()
	db := NewMysqlDatabase()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	//defer db.Close()

	query := "SELECT id,name,description,price,image FROM product WHERE id = \\? AND status = 1"

   	rows := mock.NewRows([]string{"id", "name", "description", "price", "image"}).
    AddRow(product.Id, product.Name, product.Desc, product.Price, product.Image)

  	mock.ExpectQuery(query).WithArgs(strconv.Itoa(int(product.Id))).WillReturnRows(rows)

	result, err, _ := db.SelectProduct(strconv.Itoa(int(product.Id)))

	assert.NotNil(t, result)
	assert.NoError(t, err)
}

func TestShouldRetrieveDataError(t *testing.T) {
	initMock()
	db := NewMysqlDatabase()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	//defer db.Close()

	query := "SELECT id,name,description,price,image FROM product WHERE id = \\? AND status = 1"

   	rows := mock.NewRows([]string{"id", "name", "description", "price", "image"})

  	mock.ExpectQuery(query).WithArgs(strconv.Itoa(int(product.Id))).WillReturnRows(rows)

	result, err, _ := db.SelectProduct(strconv.Itoa(int(product.Id)))

	assert.Empty(t, result)
	assert.NoError(t, err)
}

func TestShouldHandleErrorGenericOnRetrieve(t *testing.T) {
	initMock()
	db := NewMysqlDatabase()

	var result entitydatabase.Product

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	//defer db.Close()

	query := "SELECT id,name,description,price,image FROM product WHERE id = \\? AND status = 1"
   	//rows := mock.NewRows([]string{"id", "name", "description", "price", "image"})

  	mock.ExpectQuery(query).WithArgs(strconv.Itoa(int(product.Id))).WillReturnError(errors.New("EMPTY_BODY"))

	result, err, _ := db.SelectProduct(strconv.Itoa(int(product.Id)))

	assert.Error(t, err)
	assert.Empty(t, result)
}

func TestShouldHandleErrorGenericOnInsert(t *testing.T) {
	initMock()
	db := NewMysqlDatabase()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	//defer db.Close()

	//func (e *ExpectedExec) WillReturnError(err error) *ExpectedExec
	mock.ExpectExec("insert into product").WithArgs("name", "description", 900000, "apa.jpg").WillReturnError(errors.New("EMPTY_BODY"))

	// now we execute our method
	if _, err, _ = db.AddProduct("name", "description", 900000, "apa.jpg"); err == nil {
		//fmt.Println(err)
		t.Errorf("error expected while inserting data product: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldHandleErrorGenericOnUpdate(t *testing.T) {
	initMock()
	db := NewMysqlDatabase()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	//defer db.Close()

	//mock.ExpectExec("insert into product").WithArgs("name", "description", 10000, "apa.jpg").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("UPDATE product SET name = \\?, description = \\?, price = \\?, image = \\? WHERE id = \\?").WithArgs(product.Name, product.Desc, product.Price, product.Image, strconv.Itoa(int(product.Id))).WillReturnError(errors.New("EMPTY_BODY"))
	
	// now we execute our method
	if _, err, _ = db.EditProduct(product.Name, product.Desc, product.Price, product.Image, strconv.Itoa(int(product.Id))); err == nil {
		//fmt.Println(err)
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}