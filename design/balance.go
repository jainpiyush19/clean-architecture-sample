package design

import (
	. "goa.design/goa/v3/dsl"
)

const Balance = "balance"

//BalanceParams temporary since JWT is not implemented
var BalanceParams = func() {
	Param("userID", Int64, "user ID", func() {
		Minimum(0)
	})
}

//BalanceRequestBody
var BalanceRequestBody = Type("BalanceRequestBody", func() {
	Description("payload for balance api")

	// Token field for implementation of JWT auth
	Token("token", String, "JWT token used to perform authorization")

	// temporary since JWT is not implemented
	Attribute("userID", Int64, "user ID")
	Required("userID")
})

//BalanceResponseBody to return api response
var BalanceResponseBody = Type("BalanceResponseBody", func() {
	Description("response for balance api")

	Attribute("Amount", Float64)
	Attribute("UpdatedAt", String)

	Required("Amount", "UpdatedAt")
})
