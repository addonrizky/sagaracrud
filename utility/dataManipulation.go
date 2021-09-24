package utility

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"
)

func IsNowGreaterThanDate(datestring string) bool {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, datestring)

	if err != nil {
		fmt.Println(err)
	}

	if time.Now().Unix() > t.Unix() {
		return true
	}

	return false
}

func GenerateReferralCode() string {
	refcode := uuid.Must(uuid.NewRandom())
	refcodeString := strings.Replace(fmt.Sprintf("%v", refcode), "-", "", -1)
	return refcodeString
}

func AddDaysBy(days int) string {
	expirationDate := time.Now().AddDate(0, 0, days)
	return expirationDate.Format("2006-01-02 15:04:05")
}
