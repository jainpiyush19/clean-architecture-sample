package design

import (
	. "goa.design/goa/v3/dsl"
)

const AdminWallet = "admin/wallets"

//AdminWalletRequestBody
var AdminWalletRequestBody = Type("AdminWalletRequestBody", func() {
	Description("payload for admin wallet api")

	// Token field for implementation of JWT auth
	Token("token", String, "JWT token used to perform authorization")
})

//AdminWalletResponseBody to return api response
var AdminWalletResponseBody = Type("AdminWalletResponseBody", func() {
	Description("response for admin wallet api")

	Attribute("data", ArrayOf(Wallet))
})

var Wallet = Type("Wallet", func() {
	Attribute("userID", Int64)
	Attribute("Amount", Float64)
	Attribute("UpdatedAt", String)
})
