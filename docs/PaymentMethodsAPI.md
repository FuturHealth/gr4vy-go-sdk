# \PaymentMethodsAPI

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePaymentMethod**](PaymentMethodsAPI.md#DeletePaymentMethod) | **Delete** /payment-methods/{payment_method_id} | Delete payment method
[**GetPaymentMethod**](PaymentMethodsAPI.md#GetPaymentMethod) | **Get** /payment-methods/{payment_method_id} | Get payment method
[**ListBuyerPaymentMethods**](PaymentMethodsAPI.md#ListBuyerPaymentMethods) | **Get** /buyers/payment-methods | List payment methods for buyer
[**ListPaymentMethods**](PaymentMethodsAPI.md#ListPaymentMethods) | **Get** /payment-methods | List payment methods
[**NewPaymentMethod**](PaymentMethodsAPI.md#NewPaymentMethod) | **Post** /payment-methods | New payment method



## DeletePaymentMethod

> DeletePaymentMethod(ctx, paymentMethodId).Execute()

Delete payment method



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
	paymentMethodId := "46973e9d-88a7-44a6-abfe-be4ff0134ff4" // string | The ID of the payment method.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PaymentMethodsAPI.DeletePaymentMethod(context.Background(), paymentMethodId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsAPI.DeletePaymentMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentMethodId** | **string** | The ID of the payment method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentMethod

> PaymentMethod GetPaymentMethod(ctx, paymentMethodId).Execute()

Get payment method



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
	paymentMethodId := "46973e9d-88a7-44a6-abfe-be4ff0134ff4" // string | The ID of the payment method.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentMethodsAPI.GetPaymentMethod(context.Background(), paymentMethodId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsAPI.GetPaymentMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPaymentMethod`: PaymentMethod
	fmt.Fprintf(os.Stdout, "Response from `PaymentMethodsAPI.GetPaymentMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentMethodId** | **string** | The ID of the payment method. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentMethod**](PaymentMethod.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBuyerPaymentMethods

> PaymentMethodsTokenized ListBuyerPaymentMethods(ctx).BuyerId(buyerId).BuyerExternalIdentifier(buyerExternalIdentifier).Country(country).Currency(currency).SortBy(sortBy).OrderBy(orderBy).Execute()

List payment methods for buyer



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
	buyerId := "8724fd24-5489-4a5d-90fd-0604df7d3b83" // string | Filters the results to only the items for which the `buyer` has an `id` that matches this value. (optional)
	buyerExternalIdentifier := "user-12345" // string | Filters the results to only the items for which the `buyer` has an `external_identifier` that matches this value. (optional)
	country := "US" // string | Filters the results to only the items which support this country code. A country is formatted as 2-letter ISO country code. (optional)
	currency := "USD" // string | Filters the results to only the items which support this currency code. A currency is formatted as 3-letter ISO currency code. (optional)
	sortBy := "last_used_at" // string | Used by the payment method filter to sort the results by an specific usage field. (optional)
	orderBy := "desc" // string | Used to show the results ascending or descending order. (optional) (default to "desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentMethodsAPI.ListBuyerPaymentMethods(context.Background()).BuyerId(buyerId).BuyerExternalIdentifier(buyerExternalIdentifier).Country(country).Currency(currency).SortBy(sortBy).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsAPI.ListBuyerPaymentMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBuyerPaymentMethods`: PaymentMethodsTokenized
	fmt.Fprintf(os.Stdout, "Response from `PaymentMethodsAPI.ListBuyerPaymentMethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerPaymentMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buyerId** | **string** | Filters the results to only the items for which the &#x60;buyer&#x60; has an &#x60;id&#x60; that matches this value. | 
 **buyerExternalIdentifier** | **string** | Filters the results to only the items for which the &#x60;buyer&#x60; has an &#x60;external_identifier&#x60; that matches this value. | 
 **country** | **string** | Filters the results to only the items which support this country code. A country is formatted as 2-letter ISO country code. | 
 **currency** | **string** | Filters the results to only the items which support this currency code. A currency is formatted as 3-letter ISO currency code. | 
 **sortBy** | **string** | Used by the payment method filter to sort the results by an specific usage field. | 
 **orderBy** | **string** | Used to show the results ascending or descending order. | [default to &quot;desc&quot;]

### Return type

[**PaymentMethodsTokenized**](PaymentMethodsTokenized.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPaymentMethods

> PaymentMethods ListPaymentMethods(ctx).BuyerId(buyerId).BuyerExternalIdentifier(buyerExternalIdentifier).Status(status).ExternalIdentifier(externalIdentifier).Limit(limit).Cursor(cursor).Execute()

List payment methods



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
	buyerId := "8724fd24-5489-4a5d-90fd-0604df7d3b83" // string | Filters the results to only the items for which the `buyer` has an `id` that matches this value. (optional)
	buyerExternalIdentifier := "user-12345" // string | Filters the results to only the items for which the `buyer` has an `external_identifier` that matches this value. (optional)
	status := []string{"Status_example"} // []string | Filters the results to only the payment methods for which the `status` matches with any of the provided status values. (optional)
	externalIdentifier := "xid123" // string | Filters the results to only the items for which the resource has an `external_identifier` that matches this value. (optional)
	limit := int32(1) // int32 | Defines the maximum number of items to return for this request. (optional) (default to 20)
	cursor := "ZXhhbXBsZTE" // string | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the `next_cursor` field. Similarly the `previous_cursor` can be used to reverse backwards in the list. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentMethodsAPI.ListPaymentMethods(context.Background()).BuyerId(buyerId).BuyerExternalIdentifier(buyerExternalIdentifier).Status(status).ExternalIdentifier(externalIdentifier).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsAPI.ListPaymentMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPaymentMethods`: PaymentMethods
	fmt.Fprintf(os.Stdout, "Response from `PaymentMethodsAPI.ListPaymentMethods`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPaymentMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buyerId** | **string** | Filters the results to only the items for which the &#x60;buyer&#x60; has an &#x60;id&#x60; that matches this value. | 
 **buyerExternalIdentifier** | **string** | Filters the results to only the items for which the &#x60;buyer&#x60; has an &#x60;external_identifier&#x60; that matches this value. | 
 **status** | **[]string** | Filters the results to only the payment methods for which the &#x60;status&#x60; matches with any of the provided status values. | 
 **externalIdentifier** | **string** | Filters the results to only the items for which the resource has an &#x60;external_identifier&#x60; that matches this value. | 
 **limit** | **int32** | Defines the maximum number of items to return for this request. | [default to 20]
 **cursor** | **string** | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the &#x60;next_cursor&#x60; field. Similarly the &#x60;previous_cursor&#x60; can be used to reverse backwards in the list. | 

### Return type

[**PaymentMethods**](PaymentMethods.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewPaymentMethod

> PaymentMethod NewPaymentMethod(ctx).PaymentMethodRequest(paymentMethodRequest).Execute()

New payment method



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
	paymentMethodRequest := *openapiclient.NewPaymentMethodRequest("card") // PaymentMethodRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentMethodsAPI.NewPaymentMethod(context.Background()).PaymentMethodRequest(paymentMethodRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsAPI.NewPaymentMethod``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NewPaymentMethod`: PaymentMethod
	fmt.Fprintf(os.Stdout, "Response from `PaymentMethodsAPI.NewPaymentMethod`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewPaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentMethodRequest** | [**PaymentMethodRequest**](PaymentMethodRequest.md) |  | 

### Return type

[**PaymentMethod**](PaymentMethod.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

