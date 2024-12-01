package data

type OTPData struct {
	PhoneNumber string `json:"phoneNumber" validate:"required"`
}

type VerifyData struct {
	User *OTPData `json:"user" validate:"required"`
	Code string   `json:"code" validate:"required"`
}
