package entitydatabase

import (
	"database/sql"
)

type Referral struct {
	Username       sql.NullString
	ReferralCode   sql.NullString
	Status         sql.NullString
	CreatedDate    sql.NullString
	ExpirationDate sql.NullString
}
