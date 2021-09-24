package usecase

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/addonrizky/sagaracrud/constant"
	"github.com/addonrizky/sagaracrud/entity/entitydatabase"
	"github.com/addonrizky/sagaracrud/utility"
	"github.com/addonrizky/sagaracrud/repository/database"
)

/*
var(
	mail notification.EmailNotification = notification.NewEmailNotification()
)
*/

var (
	dbLogin database.Database
)

type LoginUsecase interface {
	Login(ctx context.Context, username string, password string) (string, string, string)
}

type usecaseLogin struct{}

func NewLoginUsecase(database database.Database) LoginUsecase {
	dbLogin = database
	return &usecaseLogin{}
}

func (*usecaseLogin) Login(ctx context.Context, username string, password string) (string, string, string) {
	var err error
	var user entitydatabase.User
	var code string

	hasher := sha256.New()
	userpass := username + password
	hasher.Write([]byte(userpass))
	encryptedPassword := hex.EncodeToString(hasher.Sum(nil))

	user, err, code = dbLogin.GetUserByUsername(username)

	if err != nil {
		return constant.RCDatabaseError, constant.RCDatabaseErrorDesc, ""
	}

	if code == constant.RCDataNotFound {
		return constant.RCDataNotFound, constant.RCDataNotFoundDesc, ""
	}

	if user.Password != encryptedPassword {
		return constant.RCWrongAuth, constant.RCWrongAuthDesc, ""
	}

	tokenJwt := utility.EncodeTokenJwt(username, user.EmailAddress, user.TypeUser)

	return constant.RCSuccess, constant.RCSuccessDesc, tokenJwt
}
