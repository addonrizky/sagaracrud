package database

import (
	"database/sql"
	"fmt"
	"github.com/addonrizky/sagaracrud/config"
	"github.com/addonrizky/sagaracrud/constant"
	"github.com/addonrizky/sagaracrud/entity/entitydatabase"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/DATA-DOG/go-sqlmock"

)

var (
	dbms *sql.DB
	err  error
)

type db struct{}

func NewMysqlDatabase() Database {
	return &db{}
}

func Init() {
	user := config.GetString("MYSQL_USER")
	password := config.GetString("MYSQL_PASSWORD")
	host := config.GetString("MYSQL_HOST")
	port := config.GetString("MYSQL_PORT")
	dbName := config.GetString("MYSQL_DATABASE")
	connectionString := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName
	dbms, err = sql.Open("mysql", connectionString)

	fmt.Println(connectionString)

	if err != nil {
		fmt.Println(err)
		panic("please restart")
	}
	err := dbms.Ping()
	if err != nil {
		fmt.Println(err)
		panic("please restart")
	}
	dbms.SetMaxOpenConns(10)
	dbms.SetMaxIdleConns(5)

}

func (*db) AddProduct(name string, desc string, price int, image string) (string, error, string) {
	_, err := dbms.Exec("insert into product (name,description,price,image) values (?, ?, ?, ?)", name, desc, price, image)
	if err != nil {
		return "NOT OK", err, "AF"
	}
	return "INSERT OK", nil, constant.RCSuccess
}

func (*db) EditProduct(name string, desc string, price int, image string, id string) (string, error, string) {
	_, err := dbms.Exec("UPDATE product SET name = ?, description = ?, price = ?, image = ? WHERE id = ?", name, desc, price, image, id)
	if err != nil {
		return "UPDATE FAIL", err, "EF"
	}
	return "UPDATE OK", nil, constant.RCSuccess
}

func (*db) DeleteProduct(id string) (string, error, string) {
	res, err := dbms.Exec("UPDATE product SET status = 0 WHERE id = ?", id)
	if err != nil {
		return "UPDATE FAIL", err, "EF"
	}
	count, err2 := res.RowsAffected()

	if err2 != nil {
		return "DELETION GEN FAIL", err, "NE"
	}

	if count == 0 {
		return "DELETION GEN FAIL", err, "ND"
	}

	return "UPDATE OK", nil, constant.RCSuccess
}

func (*db) SelectProduct(id string) (entitydatabase.Product, error, string) {
	product := entitydatabase.Product{}
	rows, err := dbms.Query("SELECT id,name,description,price,image FROM product WHERE id = ? AND status = 1", id)
	if err != nil {
		return product, err, constant.RCDatabaseError
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(
			&product.Id,
			&product.Name,
			&product.Desc,
			&product.Price,
			&product.Image,
		)

		if err != nil {
			return product, err, constant.RCDatabaseError
		}
	} else {
		return product, err, constant.RCDataNotFound
	}
	return product, nil, constant.RCSuccess
}

func (*db) GetUserByUsername(username string) (entitydatabase.User, error, string) {
	user := entitydatabase.User{}
	rows, err := dbms.Query("SELECT username,password,full_name,type_user FROM user WHERE username = ?", username)
	if err != nil {
		return user, err, constant.RCDatabaseError
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(
			&user.Username,
			&user.Password,
			&user.FullName,
			&user.TypeUser,
		)

		if err != nil {
			return user, err, constant.RCDatabaseError
		}
	} else {
		return user, err, constant.RCDataNotFound
	}
	return user, nil, constant.RCSuccess
}
