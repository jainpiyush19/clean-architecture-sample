package design

import (
	. "goa.design/goa/v3/dsl"
)

const Transfer = "transfer"

//TransferRequestBody request body for transfer API
var TransferRequestBody = Type("TransferRequestBody", func() {
	Description("payload for transfer api")

	// Token field for implementation of JWT auth
	Token("token", String, "JWT token used to perform authorization")

	// Body
	Attribute("receiverID", Int64, func() {
		Minimum(0)
	})

	Attribute("amount", Float64, func() {
		Minimum(0)
	})

	Required("receiverID", "amount")
})
