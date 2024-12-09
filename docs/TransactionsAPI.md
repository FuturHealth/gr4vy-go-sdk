# \TransactionsAPI

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CaptureTransaction**](TransactionsAPI.md#CaptureTransaction) | **Post** /transactions/{transaction_id}/capture | Capture transaction
[**GetRefund**](TransactionsAPI.md#GetRefund) | **Get** /transactions/{transaction_id}/refunds/{refund_id} | Get refund
[**GetSingleRefund**](TransactionsAPI.md#GetSingleRefund) | **Get** /refunds/{refund_id} | Get refund
[**GetTransaction**](TransactionsAPI.md#GetTransaction) | **Get** /transactions/{transaction_id} | Get transaction
[**GetTransactionSettlement**](TransactionsAPI.md#GetTransactionSettlement) | **Get** /transactions/{transaction_id}/settlements/{settlement_id} | Get settlement
[**GetTransactionSettlements**](TransactionsAPI.md#GetTransactionSettlements) | **Get** /transactions/{transaction_id}/settlements | List settlements for transaction
[**ListTransactionRefunds**](TransactionsAPI.md#ListTransactionRefunds) | **Get** /transactions/{transaction_id}/refunds | List refunds
[**ListTransactions**](TransactionsAPI.md#ListTransactions) | **Get** /transactions | List transactions
[**NewRefund**](TransactionsAPI.md#NewRefund) | **Post** /transactions/{transaction_id}/refunds | Refund transaction
[**NewTransaction**](TransactionsAPI.md#NewTransaction) | **Post** /transactions | New transaction
[**RefundAll**](TransactionsAPI.md#RefundAll) | **Post** /transactions/{transaction_id}/refunds/all | Refund all instruments in a transaction
[**SyncTransaction**](TransactionsAPI.md#SyncTransaction) | **Post** /transactions/{transaction_id}/sync | Sync transaction
[**VoidTransaction**](TransactionsAPI.md#VoidTransaction) | **Post** /transactions/{transaction_id}/void | Void transaction



## CaptureTransaction

> Transaction CaptureTransaction(ctx, transactionId).TransactionCaptureRequest(transactionCaptureRequest).Execute()

Capture transaction



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	transactionId := "fe26475d-ec3e-4884-9553-f7356683f7f9" // string | The ID for the transaction to get the information for.
	transactionCaptureRequest := *openapiclient.NewTransactionCaptureRequest() // TransactionCaptureRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.CaptureTransaction(context.Background(), transactionId).TransactionCaptureRequest(transactionCaptureRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.CaptureTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CaptureTransaction`: Transaction
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.CaptureTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | The ID for the transaction to get the information for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCaptureTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionCaptureRequest** | [**TransactionCaptureRequest**](TransactionCaptureRequest.md) |  | 

### Return type

[**Transaction**](Transaction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRefund

> Refund GetRefund(ctx, transactionId, refundId).Execute()

Get refund



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	transactionId := "fe26475d-ec3e-4884-9553-f7356683f7f9" // string | The ID for the transaction to get the information for.
	refundId := "8724fd24-5489-4a5d-90fd-0604df7d3b83" // string | The unique ID of the refund.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetRefund(context.Background(), transactionId, refundId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetRefund``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRefund`: Refund
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetRefund`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | The ID for the transaction to get the information for. | 
**refundId** | **string** | The unique ID of the refund. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRefundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Refund**](Refund.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSingleRefund

> Refund GetSingleRefund(ctx, refundId).Execute()

Get refund



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	refundId := "8724fd24-5489-4a5d-90fd-0604df7d3b83" // string | The unique ID of the refund.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetSingleRefund(context.Background(), refundId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetSingleRefund``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSingleRefund`: Refund
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetSingleRefund`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**refundId** | **string** | The unique ID of the refund. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleRefundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Refund**](Refund.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransaction

> Transaction GetTransaction(ctx, transactionId).Execute()

Get transaction



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	transactionId := "fe26475d-ec3e-4884-9553-f7356683f7f9" // string | The ID for the transaction to get the information for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetTransaction(context.Background(), transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransaction`: Transaction
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | The ID for the transaction to get the information for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Transaction**](Transaction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionSettlement

> Settlement GetTransactionSettlement(ctx, transactionId, settlementId).Execute()

Get settlement



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	transactionId := "fe26475d-ec3e-4884-9553-f7356683f7f9" // string | The ID for the transaction to get the information for.
	settlementId := "8724fd24-5489-4a5d-90fd-0604df7d3b83" // string | The unique ID of the settlement.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetTransactionSettlement(context.Background(), transactionId, settlementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransactionSettlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionSettlement`: Settlement
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransactionSettlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | The ID for the transaction to get the information for. | 
**settlementId** | **string** | The unique ID of the settlement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionSettlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Settlement**](Settlement.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionSettlements

> Settlements GetTransactionSettlements(ctx, transactionId).Execute()

List settlements for transaction



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	transactionId := "fe26475d-ec3e-4884-9553-f7356683f7f9" // string | The ID for the transaction to get the information for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.GetTransactionSettlements(context.Background(), transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransactionSettlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionSettlements`: Settlements
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransactionSettlements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | The ID for the transaction to get the information for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionSettlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Settlements**](Settlements.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactionRefunds

> Refunds ListTransactionRefunds(ctx, transactionId).Execute()

List refunds



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	transactionId := "fe26475d-ec3e-4884-9553-f7356683f7f9" // string | The ID for the transaction to get the information for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.ListTransactionRefunds(context.Background(), transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.ListTransactionRefunds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTransactionRefunds`: Refunds
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.ListTransactionRefunds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | The ID for the transaction to get the information for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionRefundsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Refunds**](Refunds.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactions

> Transactions ListTransactions(ctx).BuyerExternalIdentifier(buyerExternalIdentifier).BuyerId(buyerId).Cursor(cursor).Limit(limit).AmountEq(amountEq).AmountGte(amountGte).AmountLte(amountLte).CheckoutSessionId(checkoutSessionId).CreatedAtGte(createdAtGte).CreatedAtLte(createdAtLte).Currency(currency).ExternalIdentifier(externalIdentifier).GiftCardId(giftCardId).GiftCardLast4(giftCardLast4).HasGiftCardRedemptions(hasGiftCardRedemptions).HasRefunds(hasRefunds).HasSettlements(hasSettlements).Id(id).Metadata(metadata).Method(method).PaymentMethodId(paymentMethodId).PaymentMethodLabel(paymentMethodLabel).PaymentServiceId(paymentServiceId).PaymentServiceTransactionId(paymentServiceTransactionId).PendingReview(pendingReview).ReconciliationId(reconciliationId).Search(search).Status(status).UpdatedAtGte(updatedAtGte).UpdatedAtLte(updatedAtLte).IsSubsequentPayment(isSubsequentPayment).MerchantInitiated(merchantInitiated).PaymentMethodBin(paymentMethodBin).PaymentSource(paymentSource).Execute()

List transactions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	buyerExternalIdentifier := "user-12345" // string | Filters the results to only the items for which the `buyer` has an `external_identifier` that matches this value. (optional)
	buyerId := "8724fd24-5489-4a5d-90fd-0604df7d3b83" // string | Filters the results to only the items for which the `buyer` has an `id` that matches this value. (optional)
	cursor := "ZXhhbXBsZTE" // string | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the `next_cursor` field. Similarly the `previous_cursor` can be used to reverse backwards in the list. (optional)
	limit := int32(1) // int32 | Defines the maximum number of items to return for this request. (optional) (default to 20)
	amountEq := int32(500) // int32 | Filters for transactions that have an `amount` that is equal to the provided `amount_eq` value. (optional)
	amountGte := int32(500) // int32 | Filters for transactions that have an `amount` that is greater than or equal to the `amount_gte` value. (optional)
	amountLte := int32(500) // int32 | Filters for transactions that have an `amount` that is less than or equal to the `amount_lte` value. (optional)
	checkoutSessionId := "8724fd24-5489-4a5d-90fd-0604df7d3b83" // string | Filters for transactions that are linked to the unique ID for a Checkout Session. (optional)
	createdAtGte := time.Now() // time.Time | Filters the results to only transactions created after this ISO date-time string. The time zone must be included.  Ensure that the date-time string is URL encoded, e.g. `2022-01-01T12:00:00+08:00` must be encoded as `2022-01-01T12%3A00%3A00%2B08%3A00`. (optional)
	createdAtLte := time.Now() // time.Time | Filters the results to only transactions created before this ISO date-time string. The time zone must be included.  Ensure that the date-time string is URL encoded, e.g. `2022-01-01T12:00:00+08:00` must be encoded as `2022-01-01T12%3A00%3A00%2B08%3A00`. (optional)
	currency := []string{"Inner_example"} // []string | Filters for transactions that have matching `currency` values. The `currency` values provided must be formatted as 3-letter ISO currency code. (optional)
	externalIdentifier := "user-12345" // string | Filters the results to only the items for which the `external_identifier` matches this value. (optional)
	giftCardId := "be828248-56de-481e-a580-44b6e1d4df81" // string | Filters for transactions that have at least one gift card redemption with a matching `gift_card_id` value. (optional)
	giftCardLast4 := "7890" // string | Filters for transactions that have at least one gift card redemption where the last 4 digits of its gift card number matches exactly with the provided value. (optional)
	hasGiftCardRedemptions := true // bool | When set to `true`, filters for transactions that have at least one gift card redemption associated with it. When set to `false`, filter for transactions that have no gift card redemptions. (optional)
	hasRefunds := true // bool | When set to `true`, filter for transactions that have at least one completed refund (including gift card refunds) associated with it. When set to `false`, filter for transactions that have no completed refunds. (optional)
	hasSettlements := true // bool | When set to `true`, filter for transactions that have at least one settlement. When set to `false`, filter for transactions that have no settlements. (optional)
	id := "be828248-56de-481e-a580-44b6e1d4df81" // string | Filters for the transaction that has a matching `id` value. (optional)
	metadata := []string{"Inner_example"} // []string | Filters for transactions where their `metadata` values contain all of the provided `metadata` keys. The value sent for `metadata` must be formatted as a JSON string, and all keys and values must be strings. This value should also be URL encoded.  Duplicate keys are not supported. If a key is duplicated, only the last appearing value will be used. (optional)
	method := []string{"card"} // []string | Filters the results to only the items for which the `method` has been set to this value. (optional)
	paymentMethodId := "46973e9d-88a7-44a6-abfe-be4ff0134ff4" // string | Filters for transactions that have a payment method with an ID that matches exactly with the provided value. (optional)
	paymentMethodLabel := "1234" // string | Filters for transactions that have a payment method with a label that matches exactly with the provided value. (optional)
	paymentServiceId := []string{"Inner_example"} // []string | Filters for transactions that were processed by the provided `payment_service_id` values. (optional)
	paymentServiceTransactionId := "transaction_123" // string | Filters for transactions that have a matching `payment_service_transaction_id` value. The `payment_service_transaction_id` is the identifier of the transaction given by the payment service. (optional)
	pendingReview := true // bool | When set to `true`, filter for transactions that have a manual review pending. When set to `false`, filter for transactions that don't have a manual review pending. (optional)
	reconciliationId := "7EgeeeTX0DS45RBDNt4AEY" // string | Filters for transactions based on their transaction reconciliation identifier. (optional)
	search := "be828248-56de-481e-a580-44b6e1d4df81" // string | Filters for transactions that have one of the following fields match exactly with the provided `search` value.  * `buyer_external_identifier` * `buyer_id` * `external_identifier` * `id` * `payment_service_transaction_id`  The search will apply against all fields at the same time. Please do not use this query parameter in a production application, as this API call is very inefficient and may negatively impact transaction performance. (optional)
	status := []string{"Status_example"} // []string | Filters the results to only the transactions that have a `status` that matches with any of the provided status values. (optional)
	updatedAtGte := time.Now() // time.Time | Filters the results to only transactions last updated after this ISO date-time string. The time zone must be included.  Ensure that the date-time string is URL encoded, e.g. `2022-01-01T12:00:00+08:00` must be encoded as `2022-01-01T12%3A00%3A00%2B08%3A00`. (optional)
	updatedAtLte := time.Now() // time.Time | Filters the results to only transactions last updated before this ISO date-time string. The time zone must be included.  Ensure that the date-time string is URL encoded, e.g. `2022-01-01T12:00:00+08:00` must be encoded as `2022-01-01T12%3A00%3A00%2B08%3A00`. (optional)
	isSubsequentPayment := true // bool | When set to `true`, filter for transactions that are a subsequent payment. When set to `false`, filter for transactions that are not an initial payment. (optional)
	merchantInitiated := true // bool | When set to `true`, filter for transactions that are merchant initiated. When set to `false`, filter for transactions that are not merchant initiated. (optional)
	paymentMethodBin := "411111" // string | Filter for transactions that have a card with a BIN that matches exactly with the provided value. (optional)
	paymentSource := []string{"recurring"} // []string | Filters the results to only the transactions that have a payment source that matches with any of the provided values. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.ListTransactions(context.Background()).BuyerExternalIdentifier(buyerExternalIdentifier).BuyerId(buyerId).Cursor(cursor).Limit(limit).AmountEq(amountEq).AmountGte(amountGte).AmountLte(amountLte).CheckoutSessionId(checkoutSessionId).CreatedAtGte(createdAtGte).CreatedAtLte(createdAtLte).Currency(currency).ExternalIdentifier(externalIdentifier).GiftCardId(giftCardId).GiftCardLast4(giftCardLast4).HasGiftCardRedemptions(hasGiftCardRedemptions).HasRefunds(hasRefunds).HasSettlements(hasSettlements).Id(id).Metadata(metadata).Method(method).PaymentMethodId(paymentMethodId).PaymentMethodLabel(paymentMethodLabel).PaymentServiceId(paymentServiceId).PaymentServiceTransactionId(paymentServiceTransactionId).PendingReview(pendingReview).ReconciliationId(reconciliationId).Search(search).Status(status).UpdatedAtGte(updatedAtGte).UpdatedAtLte(updatedAtLte).IsSubsequentPayment(isSubsequentPayment).MerchantInitiated(merchantInitiated).PaymentMethodBin(paymentMethodBin).PaymentSource(paymentSource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.ListTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTransactions`: Transactions
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.ListTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buyerExternalIdentifier** | **string** | Filters the results to only the items for which the &#x60;buyer&#x60; has an &#x60;external_identifier&#x60; that matches this value. | 
 **buyerId** | **string** | Filters the results to only the items for which the &#x60;buyer&#x60; has an &#x60;id&#x60; that matches this value. | 
 **cursor** | **string** | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the &#x60;next_cursor&#x60; field. Similarly the &#x60;previous_cursor&#x60; can be used to reverse backwards in the list. | 
 **limit** | **int32** | Defines the maximum number of items to return for this request. | [default to 20]
 **amountEq** | **int32** | Filters for transactions that have an &#x60;amount&#x60; that is equal to the provided &#x60;amount_eq&#x60; value. | 
 **amountGte** | **int32** | Filters for transactions that have an &#x60;amount&#x60; that is greater than or equal to the &#x60;amount_gte&#x60; value. | 
 **amountLte** | **int32** | Filters for transactions that have an &#x60;amount&#x60; that is less than or equal to the &#x60;amount_lte&#x60; value. | 
 **checkoutSessionId** | **string** | Filters for transactions that are linked to the unique ID for a Checkout Session. | 
 **createdAtGte** | **time.Time** | Filters the results to only transactions created after this ISO date-time string. The time zone must be included.  Ensure that the date-time string is URL encoded, e.g. &#x60;2022-01-01T12:00:00+08:00&#x60; must be encoded as &#x60;2022-01-01T12%3A00%3A00%2B08%3A00&#x60;. | 
 **createdAtLte** | **time.Time** | Filters the results to only transactions created before this ISO date-time string. The time zone must be included.  Ensure that the date-time string is URL encoded, e.g. &#x60;2022-01-01T12:00:00+08:00&#x60; must be encoded as &#x60;2022-01-01T12%3A00%3A00%2B08%3A00&#x60;. | 
 **currency** | **[]string** | Filters for transactions that have matching &#x60;currency&#x60; values. The &#x60;currency&#x60; values provided must be formatted as 3-letter ISO currency code. | 
 **externalIdentifier** | **string** | Filters the results to only the items for which the &#x60;external_identifier&#x60; matches this value. | 
 **giftCardId** | **string** | Filters for transactions that have at least one gift card redemption with a matching &#x60;gift_card_id&#x60; value. | 
 **giftCardLast4** | **string** | Filters for transactions that have at least one gift card redemption where the last 4 digits of its gift card number matches exactly with the provided value. | 
 **hasGiftCardRedemptions** | **bool** | When set to &#x60;true&#x60;, filters for transactions that have at least one gift card redemption associated with it. When set to &#x60;false&#x60;, filter for transactions that have no gift card redemptions. | 
 **hasRefunds** | **bool** | When set to &#x60;true&#x60;, filter for transactions that have at least one completed refund (including gift card refunds) associated with it. When set to &#x60;false&#x60;, filter for transactions that have no completed refunds. | 
 **hasSettlements** | **bool** | When set to &#x60;true&#x60;, filter for transactions that have at least one settlement. When set to &#x60;false&#x60;, filter for transactions that have no settlements. | 
 **id** | **string** | Filters for the transaction that has a matching &#x60;id&#x60; value. | 
 **metadata** | **[]string** | Filters for transactions where their &#x60;metadata&#x60; values contain all of the provided &#x60;metadata&#x60; keys. The value sent for &#x60;metadata&#x60; must be formatted as a JSON string, and all keys and values must be strings. This value should also be URL encoded.  Duplicate keys are not supported. If a key is duplicated, only the last appearing value will be used. | 
 **method** | **[]string** | Filters the results to only the items for which the &#x60;method&#x60; has been set to this value. | 
 **paymentMethodId** | **string** | Filters for transactions that have a payment method with an ID that matches exactly with the provided value. | 
 **paymentMethodLabel** | **string** | Filters for transactions that have a payment method with a label that matches exactly with the provided value. | 
 **paymentServiceId** | **[]string** | Filters for transactions that were processed by the provided &#x60;payment_service_id&#x60; values. | 
 **paymentServiceTransactionId** | **string** | Filters for transactions that have a matching &#x60;payment_service_transaction_id&#x60; value. The &#x60;payment_service_transaction_id&#x60; is the identifier of the transaction given by the payment service. | 
 **pendingReview** | **bool** | When set to &#x60;true&#x60;, filter for transactions that have a manual review pending. When set to &#x60;false&#x60;, filter for transactions that don&#39;t have a manual review pending. | 
 **reconciliationId** | **string** | Filters for transactions based on their transaction reconciliation identifier. | 
 **search** | **string** | Filters for transactions that have one of the following fields match exactly with the provided &#x60;search&#x60; value.  * &#x60;buyer_external_identifier&#x60; * &#x60;buyer_id&#x60; * &#x60;external_identifier&#x60; * &#x60;id&#x60; * &#x60;payment_service_transaction_id&#x60;  The search will apply against all fields at the same time. Please do not use this query parameter in a production application, as this API call is very inefficient and may negatively impact transaction performance. | 
 **status** | **[]string** | Filters the results to only the transactions that have a &#x60;status&#x60; that matches with any of the provided status values. | 
 **updatedAtGte** | **time.Time** | Filters the results to only transactions last updated after this ISO date-time string. The time zone must be included.  Ensure that the date-time string is URL encoded, e.g. &#x60;2022-01-01T12:00:00+08:00&#x60; must be encoded as &#x60;2022-01-01T12%3A00%3A00%2B08%3A00&#x60;. | 
 **updatedAtLte** | **time.Time** | Filters the results to only transactions last updated before this ISO date-time string. The time zone must be included.  Ensure that the date-time string is URL encoded, e.g. &#x60;2022-01-01T12:00:00+08:00&#x60; must be encoded as &#x60;2022-01-01T12%3A00%3A00%2B08%3A00&#x60;. | 
 **isSubsequentPayment** | **bool** | When set to &#x60;true&#x60;, filter for transactions that are a subsequent payment. When set to &#x60;false&#x60;, filter for transactions that are not an initial payment. | 
 **merchantInitiated** | **bool** | When set to &#x60;true&#x60;, filter for transactions that are merchant initiated. When set to &#x60;false&#x60;, filter for transactions that are not merchant initiated. | 
 **paymentMethodBin** | **string** | Filter for transactions that have a card with a BIN that matches exactly with the provided value. | 
 **paymentSource** | **[]string** | Filters the results to only the transactions that have a payment source that matches with any of the provided values. | 

### Return type

[**Transactions**](Transactions.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewRefund

> Refund NewRefund(ctx, transactionId).TransactionRefundRequest(transactionRefundRequest).Execute()

Refund transaction



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	transactionId := "fe26475d-ec3e-4884-9553-f7356683f7f9" // string | The ID for the transaction to get the information for.
	transactionRefundRequest := *openapiclient.NewTransactionRefundRequest() // TransactionRefundRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.NewRefund(context.Background(), transactionId).TransactionRefundRequest(transactionRefundRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.NewRefund``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NewRefund`: Refund
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.NewRefund`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | The ID for the transaction to get the information for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNewRefundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionRefundRequest** | [**TransactionRefundRequest**](TransactionRefundRequest.md) |  | 

### Return type

[**Refund**](Refund.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewTransaction

> Transaction NewTransaction(ctx).IdempotencyKey(idempotencyKey).TransactionRequest(transactionRequest).Execute()

New transaction



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	idempotencyKey := "bffa9ce6-7a8a-449c-889a-65bd2ee86903" // string | A unique key that identifies this request. Providing this header will make this an idempotent request. We recommend using V4 UUIDs, or another random string with enough entropy to avoid collisions. (optional)
	transactionRequest := *openapiclient.NewTransactionRequest(int32(1299), "USD") // TransactionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.NewTransaction(context.Background()).IdempotencyKey(idempotencyKey).TransactionRequest(transactionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.NewTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NewTransaction`: Transaction
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.NewTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **idempotencyKey** | **string** | A unique key that identifies this request. Providing this header will make this an idempotent request. We recommend using V4 UUIDs, or another random string with enough entropy to avoid collisions. | 
 **transactionRequest** | [**TransactionRequest**](TransactionRequest.md) |  | 

### Return type

[**Transaction**](Transaction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefundAll

> Refunds RefundAll(ctx, transactionId).TransactionRefundAllRequest(transactionRefundAllRequest).Execute()

Refund all instruments in a transaction



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	transactionId := "fe26475d-ec3e-4884-9553-f7356683f7f9" // string | The ID for the transaction to get the information for.
	transactionRefundAllRequest := *openapiclient.NewTransactionRefundAllRequest() // TransactionRefundAllRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.RefundAll(context.Background(), transactionId).TransactionRefundAllRequest(transactionRefundAllRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.RefundAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefundAll`: Refunds
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.RefundAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | The ID for the transaction to get the information for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRefundAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionRefundAllRequest** | [**TransactionRefundAllRequest**](TransactionRefundAllRequest.md) |  | 

### Return type

[**Refunds**](Refunds.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncTransaction

> Transaction SyncTransaction(ctx, transactionId).Execute()

Sync transaction



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	transactionId := "fe26475d-ec3e-4884-9553-f7356683f7f9" // string | The ID for the transaction to get the information for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.SyncTransaction(context.Background(), transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.SyncTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncTransaction`: Transaction
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.SyncTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | The ID for the transaction to get the information for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Transaction**](Transaction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoidTransaction

> Transaction VoidTransaction(ctx, transactionId).Execute()

Void transaction



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	transactionId := "fe26475d-ec3e-4884-9553-f7356683f7f9" // string | The ID for the transaction to get the information for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsAPI.VoidTransaction(context.Background(), transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.VoidTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VoidTransaction`: Transaction
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.VoidTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | The ID for the transaction to get the information for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVoidTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Transaction**](Transaction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

