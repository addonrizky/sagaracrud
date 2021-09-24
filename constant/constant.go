package constant

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	//context key
	CtxKeyStartTransaction = "startTransaction"
	CtxKeyIdTransaction    = "idTransaction"
	CtxKeyResponseDefault  = "responseDefault"
	CtxKeyRequestBody      = "requestBody"
	CtxKeyRefnum           = "requestRefnum"
	CtxUsername            = "username"
	CtxEmail               = "email"
	CtxTypeUser            = "typeUser"

	//Server
	ServerPort         = ":4000"
	ServerReadTimeout  = 20 * time.Second
	ServerWriteTimeout = 20 * time.Second
	ServiceName        = "sagara"

	//Log
	LogTimeFormat  = "[2006-01-02 15:04:05] "
	LogTimeElastic = "2006-01-02T15:04:05-0700"
	LogFilename    = "log-2006-01-02.log"
	LogLevel       = logrus.DebugLevel

	//Path
	PathApiVersion = "/api/v1"

	//Response Object
	RespContentType    = "application/json"
	InvalidMethodResp  = "Invalid Method for this request path"
	InvalidHeaderResp  = "content type not application/json, request fail"
	EmptyHeaderResp    = "content type not sent, request fail"
	InquirySuccess     = "Inquiry Successful"
	PaymentSuccessResp = "Transfer Request Successful"
	HealthyResp        = "Alhamdulillah sehat"
	BankAdmin          = "0"
	RealBankAdmin      = 250
	EmptyDataResp      struct{}

	//Response Code
	RCValidationError     = "VE"
	RCValidationErrorDesc = "Validation Error"
	RCDatabaseError       = "DE"
	RCDatabaseErrorDesc   = "Database Error"
	RCSuccess             = "00"
	RCSuccessDesc         = "Sukses"
	RCAlreadyExist        = "AE"
	RCAlreadyExistDesc    = "Data Already Exist"
	RCTimeout             = "TE"
	RCTimeoutDesc         = "Timeout ke ESB"
	RCTimeoutVikendi      = "TV"
	RCLoadParamFail       = "LP"
	RCLoadParamDesc       = "Load Parameter Fail"
	RCAccount             = "NA"
	RCAccDesc             = "Load Account Fail"
	RCAccNotMatchDesc     = "Account Not Match"
	RCInvalidLimit        = "IL"
	RCLimitDesc           = "Invalid Limit"
	RCDataNotFound        = "NF"
	RCDataNotFoundDesc    = "Data Not Found"
	RCWrongAuth           = "WA"
	RCWrongAuthDesc       = "Invalid auth user / password"
	RCImageFail0          = "D0"
	RCImageFail1          = "D1"
	RCImageFail2          = "D2"
	RCImageFailDesc       = "Decode Image Fail"

	//url path
	LoginPath    = "/login"
	CreatePath   = "/product/create"
	UpdatePath   = "/product/update"
	RetrievePath = "/product"

	//jwt part
	APPLICATION_NAME          = "My Simple JWT App"
	LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
	JWT_SIGNING_METHOD        = jwt.SigningMethodHS256
	JWT_SIGNATURE_KEY         = []byte("wakanda forever nippon paint")
)
