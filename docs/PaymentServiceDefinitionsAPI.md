# \PaymentServiceDefinitionsAPI

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePaymentServiceDefinitionSession**](PaymentServiceDefinitionsAPI.md#CreatePaymentServiceDefinitionSession) | **Post** /payment-service-definitions/{payment_service_definition_id}/sessions | Create a session for a payment service
[**GetPaymentServiceDefinition**](PaymentServiceDefinitionsAPI.md#GetPaymentServiceDefinition) | **Get** /payment-service-definitions/{payment_service_definition_id} | Get payment service definition
[**ListPaymentServiceDefinitions**](PaymentServiceDefinitionsAPI.md#ListPaymentServiceDefinitions) | **Get** /payment-service-definitions | List payment service definitions



## CreatePaymentServiceDefinitionSession

> PaymentServiceSession CreatePaymentServiceDefinitionSession(ctx, paymentServiceDefinitionId).RequestBody(requestBody).Execute()

Create a session for a payment service



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/FuturHealth/gr4vy-go-sdk"
)

func main() {
	paymentServiceDefinitionId := "stripe-card" // string | The unique ID of the payment service definition.
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentServiceDefinitionsAPI.CreatePaymentServiceDefinitionSession(context.Background(), paymentServiceDefinitionId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentServiceDefinitionsAPI.CreatePaymentServiceDefinitionSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePaymentServiceDefinitionSession`: PaymentServiceSession
	fmt.Fprintf(os.Stdout, "Response from `PaymentServiceDefinitionsAPI.CreatePaymentServiceDefinitionSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentServiceDefinitionId** | **string** | The unique ID of the payment service definition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentServiceDefinitionSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]interface{}** |  | 

### Return type

[**PaymentServiceSession**](PaymentServiceSession.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentServiceDefinition

> PaymentServiceDefinition GetPaymentServiceDefinition(ctx, paymentServiceDefinitionId).Execute()

Get payment service definition



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/FuturHealth/gr4vy-go-sdk"
)

func main() {
	paymentServiceDefinitionId := "stripe-card" // string | The unique ID of the payment service definition.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentServiceDefinitionsAPI.GetPaymentServiceDefinition(context.Background(), paymentServiceDefinitionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentServiceDefinitionsAPI.GetPaymentServiceDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPaymentServiceDefinition`: PaymentServiceDefinition
	fmt.Fprintf(os.Stdout, "Response from `PaymentServiceDefinitionsAPI.GetPaymentServiceDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentServiceDefinitionId** | **string** | The unique ID of the payment service definition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentServiceDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentServiceDefinition**](PaymentServiceDefinition.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPaymentServiceDefinitions

> PaymentServiceDefinitions ListPaymentServiceDefinitions(ctx).Limit(limit).Cursor(cursor).Execute()

List payment service definitions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/FuturHealth/gr4vy-go-sdk"
)

func main() {
	limit := int32(1) // int32 | Defines the maximum number of items to return for this request. (optional) (default to 20)
	cursor := "ZXhhbXBsZTE" // string | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the `next_cursor` field. Similarly the `previous_cursor` can be used to reverse backwards in the list. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentServiceDefinitionsAPI.ListPaymentServiceDefinitions(context.Background()).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentServiceDefinitionsAPI.ListPaymentServiceDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPaymentServiceDefinitions`: PaymentServiceDefinitions
	fmt.Fprintf(os.Stdout, "Response from `PaymentServiceDefinitionsAPI.ListPaymentServiceDefinitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPaymentServiceDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Defines the maximum number of items to return for this request. | [default to 20]
 **cursor** | **string** | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the &#x60;next_cursor&#x60; field. Similarly the &#x60;previous_cursor&#x60; can be used to reverse backwards in the list. | 

### Return type

[**PaymentServiceDefinitions**](PaymentServiceDefinitions.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

