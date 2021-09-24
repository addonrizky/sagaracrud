package entityself

type Login struct {
	ClientID      string `json:"client" validate:"required"`
	RequestRefnum string `json:"request_refnum" validate:"required,numeric,len=12"`
	Username      string `json:"username" validate:"required"`
	Password      string `json:"password" validate:"required"`
	Timestamp     string `json:"timestamp" validate:"required,numeric"`
}

type Create struct {
	ClientID      string `json:"client" validate:"required"`
	RequestRefnum string `json:"request_refnum" validate:"required,numeric,len=12"`
	Name          string `json:"name" validate:"required"`
	Desc          string `json:"desc" validate:"required"`
	Price         string `json:"price" validate:"required,numeric"`
	Image         string `json:"image" validate:"required"`
	Timestamp     string `json:"timestamp" validate:"required,numeric"`
}

type Update struct {
	ClientID      string `json:"client" validate:"required"`
	RequestRefnum string `json:"request_refnum" validate:"required,numeric,len=12"`
	ProductID     string `json:"product_id" validate:"required,numeric"`
	Name          string `json:"name"`
	Desc          string `json:"desc"`
	Price         string `json:"price" validate:"required,numeric"`
	Image         string `json:"image"`
	Timestamp     string `json:"timestamp" validate:"required,numeric"`
}

type Get struct {
	Token     string `json:"tok" validate:"required"`
	ProductID string `json:"product_id" validate:"required,numeric"`
}
