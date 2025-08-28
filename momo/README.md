# Momo MTN Golang SDK

A Golang SDK for the Momo MTN REST API. It provides a simple interface over the `Collection`, `Disbursement` and `Remittance` products.

## Features

Fully-featured SDK for all API products:
1. Collection:
	- bc-authorize
	- CancelInvoice
	- CancelPreApproval
	- CreateAccessToken
	- CreateInvoice
	- CreateOauth2Token
	- CreatePayments
	- GetAccountBalance
	- GetAccountBalanceInSpecificCurrency
	- GetApprovedPreApprovals
	- GetBasicUserinfo
	- GetInvoiceStatus
	- GetPaymentStatus
	- GetPreApprovalStatus
	- GetUserInfoWithConsent
	- PreApproval
	- RequesttoPay
	- RequesttoPayDeliveryNotification
	- RequesttoPayTransactionStatus
	- RequestToWithdrawTransactionStatus
	- RequestToWithdraw-V1
	- RequestToWithdraw-V2
	- ValidateAccountHolderStatus

2. Disbursement:
	- bc-authorize
	- CreateAccessToken
	- CreateOauth2Token
	- GetBasicUserinfo
	- GetUserInfoWithConsent
	- ValidateAccountHolderStatus
	- GetAccountBalance
	- GetAccountBalanceInSpecificCurrency
	- Deposit-v1
	- Deposit-v2
	- GetDepositStatus
	- GetRefundStatus
	- GetTransferStatus
	- Refund-v1
	- Refund-v2
	- Transfer

3. Remittance
	- bc-authorize
	- CreateAccessToken
	- CreateOauth2Token
	- GetBasicUserinfo
	- GetUserInfoWithConsent
	- ValidateAccountHolderStatus
	- GetAccountBalance
	- GetAccountBalanceInSpecificCurrency
	- CashTransfer
	- Transfer
	- GetCashTransferStatus
	- GetTransferStatus
	- GetBasicUserinfo(clone)
	- GetBasicUserinfo-v3

## Pre-requisites

To use this SDK, you'll need:

1. Sign up with [Momo MTN](https://momodeveloper.mtn.com/api-documentation) to get your API credentials
2. You'll need the following credentials:
   - API Key and API Secret
   - Subscription keys for the products you'll be using
   - If handling callback, register your callback URL in the portal

## Installation

Install the library with the go get command:

```bash
go get github.com/nutcas3/payment-rails/momo
```

Then import it to your main package as:

```go
package main

import (
    "github.com/nutcas3/payment-rails/momo"
)
```

## Usage and Examples

Set the following environment variables:
``` bash

    export TARGET_ENVRIONMENT="" # sandbox or production
    export API_KEY=""
    export API_SECRET=""
    export COLLECTION_SUBSCRIPTION_KEY=""
    export DISBURSEMENT_SUBSCRIPTION_KEY=""
    export REMITTANCE_SUBSCRIPTION_KEY=""

```

API endpoints for Momo MTN are classified into product groups i.e. Collection, Disbursement and Remittance. When creating the sdk client you can instantiate a single product group or all three depending on your uses. Subscription keys are a must when interacting with the API. So you must pass at least one of `Collection or Disbursement or Remittance` subscription keys. If you pass all three then the SDK instantiates a client with all three product API available, otherwise it only instantiates the product whose subscription key was passed and the others are nil.

Example uses of the SDK can be found in the [momo SDK examples](/examples/momo/) folder.

- [General use](/examples/momo/main.go)
- [Collection examples](/examples/momo/collection/main.go)
- [Disbursement examples](/examples/momo/disbursement/main.go)
- [Remittance examples](/examples/momo/remittance/main.go)

## Testing

To run the tests:

```bash
go test ./...
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
