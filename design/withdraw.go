package design

import (
	. "goa.design/goa/v3/dsl"
)

const Withdraw = "withdraw"

//WithdrawRequestBody request body for withdraw API
var WithdrawRequestBody = Type("WithdrawRequestBody", func() {
	Description("payload for withdraw api")

	// Token field for implementation of JWT auth
	Token("token", String, "JWT token used to perform authorization")

	// Body
	Attribute("amount", Float64, func() {
		Minimum(0)
	})
	Required("amount")
})
