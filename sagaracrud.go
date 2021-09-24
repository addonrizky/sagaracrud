package main

import (
	"github.com/addonrizky/sagaracrud/constant"
	"github.com/addonrizky/sagaracrud/controller"
	"github.com/addonrizky/sagaracrud/log"
	"github.com/addonrizky/sagaracrud/middleware"
	"github.com/addonrizky/sagaracrud/repository/database"
	"github.com/addonrizky/sagaracrud/validator"
	"github.com/gorilla/mux"
	"net/http"
)

/*var (
	dbengine = database.NewMysqlDatabase()
)*/

func init() {
	/*log.Init()
	validator.Init()*/
	//opentracing.SetGlobalTracer(apmot.New())
	validator.Init()
	database.Init()
	log.Init()
}

func main() {
	router := mux.NewRouter()
	// library.SetupRedisData()

	//declare middleware
	router.Use(middleware.AuthorizationJwt)
	router.Use(middleware.InitializeContext)
	router.Use(middleware.ReadRequest)
	//router.Use(middleware.ValidateRequestBody)
	//router.Use(middleware.BodyValidationMiddleware)

	//declare routing
	router.HandleFunc(constant.LoginPath, controller.Login).Methods("POST").Schemes("http").Headers("Content-type", "application/json")
	router.HandleFunc(constant.CreatePath, controller.CreateProduct).Methods("POST").Headers("Content-type", "application/json")
	router.HandleFunc(constant.UpdatePath, controller.UpdateProduct).Methods("POST").Headers("Content-type", "application/json")
	router.HandleFunc(constant.RetrievePath, controller.RetrieveProduct).Methods("GET").Headers("Content-type", "application/json")
	router.HandleFunc(constant.RetrievePath, controller.DeleteProduct).Methods("DELETE").Headers("Content-type", "application/json")
	//router.HandleFunc(constant.ReferralRegistrationPath, controller.ReferralRegistration).Methods("POST").Schemes("http").Headers("Content-type", "application/json")
	//router.HandleFunc(constant.ContributionClickPath, controller.ContributionClick).Methods("GET").Schemes("http")

	//fmt.Println("cekidot")
	//server up and listen to port xxxx
	http.ListenAndServe(constant.ServerPort, router)
}

/*func cekFirst(w http.ResponseWriter, r *http.Request) {

	//Debugging
	//cekstart := r.Context().Value("startTransaction")
	//idTransaction := r.Context().Value("idTransaction")
	//requestBody := r.Context().Value("requestBody")

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "sayonara im the main controller")
}*/
