package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API("wallet", func() {
	Title("Wallet Service")
	Description("Service for APIs for wallet transactions")
	Server("wallet", func() {
		Host("localhost", func() { URI("http://localhost:8088") })
	})
})

var JWTAuth = JWTSecurity("jwt")

var _ = Service("v1/wallet", func() {
	Description("wallet service contains APIs for wallet transactions")

	cors.Origin("*")

	// Health check API sor the server and downstream services
	Method("health", func() {
		Description("This API checks for status 200 for downstream services")
		HTTP(func() {
			GET("health")
			Response(StatusOK)
		})
	})

	Method(Deposit, func() {
		Description("This API deposit money in wallet")
		// JWT Auth for all the methods in this service
		Security(JWTAuth)
		Payload(DepositRequestBody)
		HTTP(func() {
			POST(Deposit)
			Response(StatusCreated)
		})
	})

	Method(Withdraw, func() {
		Description("This API withdraw money from wallet")
		// JWT Auth for all the methods in this service
		Security(JWTAuth)
		Payload(WithdrawRequestBody)
		HTTP(func() {
			POST(Withdraw)
			Response(StatusCreated)
		})
	})

	Method(Transfer, func() {
		Description("This API transfer money from one wallet to another")
		// JWT Auth for all the methods in this service
		Security(JWTAuth)
		Payload(TransferRequestBody)
		HTTP(func() {
			POST(Transfer)
			Response(StatusCreated)
		})
	})

	Method(Balance, func() {
		Description("This API checks balance in wallet")
		// JWT Auth for all the methods in this service
		Security(JWTAuth)
		Payload(BalanceRequestBody)
		Result(BalanceResponseBody)
		HTTP(func() {
			GET(Balance)
			Params(BalanceParams)
			Response(StatusOK)
		})
	})

	Method(AdminWallet, func() {
		Description("This API returns all wallets")
		// JWT Auth for all the methods in this service
		Security(JWTAuth)
		Payload(AdminWalletRequestBody)
		Result(AdminWalletResponseBody)
		HTTP(func() {
			GET(AdminWallet)
			Response(StatusOK)
		})
	})
})
