// Code generated by goa v3.4.3, DO NOT EDIT.
//
// wallet HTTP client CLI support package
//
// Command:
// $ goa gen github.com/jainpiyush19/cryptowallet/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	v1walletc "github.com/jainpiyush19/cryptowallet/gen/http/v1_wallet/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `v1-/wallet (health|deposit|withdraw|transfer|balance|admin-/wallets)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` v1-/wallet health` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		v1WalletFlags = flag.NewFlagSet("v1-/wallet", flag.ContinueOnError)

		v1WalletHealthFlags = flag.NewFlagSet("health", flag.ExitOnError)

		v1WalletDepositFlags     = flag.NewFlagSet("deposit", flag.ExitOnError)
		v1WalletDepositBodyFlag  = v1WalletDepositFlags.String("body", "REQUIRED", "")
		v1WalletDepositTokenFlag = v1WalletDepositFlags.String("token", "", "")

		v1WalletWithdrawFlags     = flag.NewFlagSet("withdraw", flag.ExitOnError)
		v1WalletWithdrawBodyFlag  = v1WalletWithdrawFlags.String("body", "REQUIRED", "")
		v1WalletWithdrawTokenFlag = v1WalletWithdrawFlags.String("token", "", "")

		v1WalletTransferFlags     = flag.NewFlagSet("transfer", flag.ExitOnError)
		v1WalletTransferBodyFlag  = v1WalletTransferFlags.String("body", "REQUIRED", "")
		v1WalletTransferTokenFlag = v1WalletTransferFlags.String("token", "", "")

		v1WalletBalanceFlags      = flag.NewFlagSet("balance", flag.ExitOnError)
		v1WalletBalanceUserIDFlag = v1WalletBalanceFlags.String("user-id", "REQUIRED", "")
		v1WalletBalanceTokenFlag  = v1WalletBalanceFlags.String("token", "", "")

		v1WalletAdminWalletsFlags     = flag.NewFlagSet("admin-/wallets", flag.ExitOnError)
		v1WalletAdminWalletsTokenFlag = v1WalletAdminWalletsFlags.String("token", "", "")
	)
	v1WalletFlags.Usage = v1WalletUsage
	v1WalletHealthFlags.Usage = v1WalletHealthUsage
	v1WalletDepositFlags.Usage = v1WalletDepositUsage
	v1WalletWithdrawFlags.Usage = v1WalletWithdrawUsage
	v1WalletTransferFlags.Usage = v1WalletTransferUsage
	v1WalletBalanceFlags.Usage = v1WalletBalanceUsage
	v1WalletAdminWalletsFlags.Usage = v1WalletAdminWalletsUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "v1-/wallet":
			svcf = v1WalletFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "v1-/wallet":
			switch epn {
			case "health":
				epf = v1WalletHealthFlags

			case "deposit":
				epf = v1WalletDepositFlags

			case "withdraw":
				epf = v1WalletWithdrawFlags

			case "transfer":
				epf = v1WalletTransferFlags

			case "balance":
				epf = v1WalletBalanceFlags

			case "admin-/wallets":
				epf = v1WalletAdminWalletsFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "v1-/wallet":
			c := v1walletc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "health":
				endpoint = c.Health()
				data = nil
			case "deposit":
				endpoint = c.Deposit()
				data, err = v1walletc.BuildDepositPayload(*v1WalletDepositBodyFlag, *v1WalletDepositTokenFlag)
			case "withdraw":
				endpoint = c.Withdraw()
				data, err = v1walletc.BuildWithdrawPayload(*v1WalletWithdrawBodyFlag, *v1WalletWithdrawTokenFlag)
			case "transfer":
				endpoint = c.Transfer()
				data, err = v1walletc.BuildTransferPayload(*v1WalletTransferBodyFlag, *v1WalletTransferTokenFlag)
			case "balance":
				endpoint = c.Balance()
				data, err = v1walletc.BuildBalancePayload(*v1WalletBalanceUserIDFlag, *v1WalletBalanceTokenFlag)
			case "admin-/wallets":
				endpoint = c.AdminWallets()
				data, err = v1walletc.BuildAdminWalletsPayload(*v1WalletAdminWalletsTokenFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// v1-/walletUsage displays the usage of the v1-/wallet command and its
// subcommands.
func v1WalletUsage() {
	fmt.Fprintf(os.Stderr, `wallet service contains APIs for wallet transactions
Usage:
    %s [globalflags] v1-/wallet COMMAND [flags]

COMMAND:
    health: This API checks for status 200 for downstream services
    deposit: This API deposit money in wallet
    withdraw: This API withdraw money from wallet
    transfer: This API transfer money from one wallet to another
    balance: This API checks balance in wallet
    admin-/wallets: This API returns all wallets

Additional help:
    %s v1-/wallet COMMAND --help
`, os.Args[0], os.Args[0])
}
func v1WalletHealthUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] v1-/wallet health

This API checks for status 200 for downstream services

Example:
    `+os.Args[0]+` v1-/wallet health
`, os.Args[0])
}

func v1WalletDepositUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] v1-/wallet deposit -body JSON -token STRING

This API deposit money in wallet
    -body JSON: 
    -token STRING: 

Example:
    `+os.Args[0]+` v1-/wallet deposit --body '{
      "amount": 0.7821871049830849
   }' --token "Provident doloremque et totam officiis molestiae vel."
`, os.Args[0])
}

func v1WalletWithdrawUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] v1-/wallet withdraw -body JSON -token STRING

This API withdraw money from wallet
    -body JSON: 
    -token STRING: 

Example:
    `+os.Args[0]+` v1-/wallet withdraw --body '{
      "amount": 0.7427372802563682
   }' --token "Ut accusamus iure."
`, os.Args[0])
}

func v1WalletTransferUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] v1-/wallet transfer -body JSON -token STRING

This API transfer money from one wallet to another
    -body JSON: 
    -token STRING: 

Example:
    `+os.Args[0]+` v1-/wallet transfer --body '{
      "amount": 0.3172865797517538,
      "receiverID": 294003090530668039
   }' --token "Qui voluptatem."
`, os.Args[0])
}

func v1WalletBalanceUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] v1-/wallet balance -user-id INT64 -token STRING

This API checks balance in wallet
    -user-id INT64: 
    -token STRING: 

Example:
    `+os.Args[0]+` v1-/wallet balance --user-id 5012556684053456084 --token "Repellat unde ad."
`, os.Args[0])
}

func v1WalletAdminWalletsUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] v1-/wallet admin-/wallets -token STRING

This API returns all wallets
    -token STRING: 

Example:
    `+os.Args[0]+` v1-/wallet admin-/wallets --token "Eum officiis laboriosam."
`, os.Args[0])
}
