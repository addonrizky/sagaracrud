package usecase

import (
	"github.com/addonrizky/sagaracrud/entity/entitydatabase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"crypto/sha256"
	"encoding/hex"
	"testing"
)
type mockDatabase struct {
	mock.Mock
}

func (mock *mockDatabase) GetUserByUsername(username string) (entitydatabase.User, error, string) {
	args := mock.Called(username)
	result:= args.Get(0)

	return result.(entitydatabase.User), args.Error(1), args.String(2)
}

func (mock *mockDatabase) AddProduct(name string, desc string, price int, image string) (string, error, string) {
	args := mock.Called(name,desc,price,image)
	return args.String(0), args.Error(1), args.String(2)
}

func (mock *mockDatabase) EditProduct(name string, desc string, price int, image string, id string) (string, error, string) {
	return "",nil,""
}

func (mock *mockDatabase) SelectProduct(id string) (entitydatabase.Product, error, string) {
	args := mock.Called(id)
	result := args.Get(0)
	return result.(entitydatabase.Product), args.Error(1), args.String(2)
}

func (mock *mockDatabase) DeleteProduct(id string) (string, error, string) {
	return "",nil,""
}

func TestShouldSuccessLogin(t *testing.T) {
	mockRepo := new(mockDatabase)
	loginUsecase := NewLoginUsecase(mockRepo)

	usernameInput := "joni jondoe"
	passwordInput := "jonijondoe123"

	//encrypted pass
	hasher := sha256.New()
	userpass := usernameInput + passwordInput
	hasher.Write([]byte(userpass))
	encryptedPassword := hex.EncodeToString(hasher.Sum(nil))

	user := entitydatabase.User{
		Username:		usernameInput,
		Password:		encryptedPassword,     
		FullName:		"Joni Depth",     
		TypeUser:		"1",    
		EmailAddress:	"adhon.rizky@gmail.com",
	}
	mockRepo.On("GetUserByUsername", usernameInput).Return(user, nil, "00")
	code,_,_ := loginUsecase.Login(nil, usernameInput, passwordInput)


	assert.NotNil(t, code)
	assert.Equal(t, code, "00", "they should be equal")
}

func TestShouldFailWrongPassword(t *testing.T) {
	mockRepo := new(mockDatabase)
	loginUsecase := NewLoginUsecase(mockRepo)

	usernameInput := "joni jondoe"
	passwordInput := "jonijondoe123"
	indbPassword := "9jfefj2e09j2093j209j2930gj9023jg90g3"

	//encrypted pass
	/*hasher := sha256.New()
	userpass := usernameInput + passwordInput
	hasher.Write([]byte(userpass))
	encryptedPassword := hex.EncodeToString(hasher.Sum(nil))*/

	user := entitydatabase.User{
		Username:		usernameInput,
		Password:		indbPassword,     
		FullName:		"Joni Depth",     
		TypeUser:		"1",    
		EmailAddress:	"adhon.rizky@gmail.com",
	}
	mockRepo.On("GetUserByUsername", usernameInput).Return(user, nil, "00")
	code,_,_ := loginUsecase.Login(nil, usernameInput, passwordInput)


	assert.NotNil(t, code)
	assert.Equal(t, code, "WA", "they should be equal")
}