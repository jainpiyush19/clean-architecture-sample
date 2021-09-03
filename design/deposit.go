package design

import (
	. "goa.design/goa/v3/dsl"
)

const Deposit = "deposit"

//DepositRequestBody request body for deposit API
var DepositRequestBody = Type("DepositRequestBody", func() {
	Description("payload for deposit api")

	// Token field for implementation of JWT auth
	Token("token", String, "JWT token used to perform authorization")

	// Body
	Attribute("amount", Float64, func() {
		Minimum(0)
	})
	Required("amount")
})
