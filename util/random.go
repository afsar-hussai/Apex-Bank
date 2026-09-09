package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomFullName generates a random full name for users
func RandomFullName() string {
	return fmt.Sprintf("%s %s", RandomString(5), RandomString(7))
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@apexbank.com", RandomString(6))
}

// RandomPhone generates a random 10-digit phone number
func RandomPhone() string {
	return fmt.Sprintf("98%d", RandomInt(10000000, 99999999))
}

// RandomMoney generates a random amount of money in paise
func RandomMoney() int64 {
	return RandomInt(100, 100000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{"INR", "USD", "EUR"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// RandomAccountType generates a random account type
func RandomAccountType() string {
	types := []string{"savings", "current", "fixed_deposit"}
	n := len(types)
	return types[rand.Intn(n)]
}

// RandomRole generates a random user role
func RandomRole() string {
	roles := []string{"customer", "teller", "admin"}
	return roles[rand.Intn(len(roles))]
}

// RandomKycStatus generates a random KYC status
func RandomKycStatus() string {
	statuses := []string{"pending", "approved", "rejected"}
	return statuses[rand.Intn(len(statuses))]
}

// RandomTransactionType generates a random transaction type
func RandomTransactionType() string {
	types := []string{"NEFT", "RTGS", "IMPS"}
	return types[rand.Intn(len(types))]
}