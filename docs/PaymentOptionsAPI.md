# \PaymentOptionsAPI

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPaymentOptions**](PaymentOptionsAPI.md#ListPaymentOptions) | **Get** /payment-options | List payment options
[**PostListPaymentOptions**](PaymentOptionsAPI.md#PostListPaymentOptions) | **Post** /payment-options | List payment options with POST



## ListPaymentOptions

> PaymentOptions ListPaymentOptions(ctx).Country(country).Currency(currency).Amount(amount).Metadata(metadata).Locale(locale).Execute()

List payment options



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
	country := "US" // string | Filters the results to only the items which support this country code. A country is formatted as 2-letter ISO country code. (optional)
	currency := "USD" // string | Filters the results to only the items which support this currency code. A currency is formatted as 3-letter ISO currency code. (optional)
	amount := int32(500) // int32 | Used by the Flow engine to filter the results based on the transaction amount. (optional)
	metadata := "{"restricted_items": "True"}" // string | Used by the Flow engine to filter available options based on various client-defined parameters. If present, this must be a string representing a valid JSON dictionary. (optional)
	locale := "en-US" // string | An ISO 639-1 Language Code and optional ISO 3166 Country Code. This locale determines the language for the labels returned for every payment option. (optional) (default to "en")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentOptionsAPI.ListPaymentOptions(context.Background()).Country(country).Currency(currency).Amount(amount).Metadata(metadata).Locale(locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentOptionsAPI.ListPaymentOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPaymentOptions`: PaymentOptions
	fmt.Fprintf(os.Stdout, "Response from `PaymentOptionsAPI.ListPaymentOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPaymentOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **string** | Filters the results to only the items which support this country code. A country is formatted as 2-letter ISO country code. | 
 **currency** | **string** | Filters the results to only the items which support this currency code. A currency is formatted as 3-letter ISO currency code. | 
 **amount** | **int32** | Used by the Flow engine to filter the results based on the transaction amount. | 
 **metadata** | **string** | Used by the Flow engine to filter available options based on various client-defined parameters. If present, this must be a string representing a valid JSON dictionary. | 
 **locale** | **string** | An ISO 639-1 Language Code and optional ISO 3166 Country Code. This locale determines the language for the labels returned for every payment option. | [default to &quot;en&quot;]

### Return type

[**PaymentOptions**](PaymentOptions.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostListPaymentOptions

> PaymentOptions PostListPaymentOptions(ctx).PaymentOptionsRequest(paymentOptionsRequest).Execute()

List payment options with POST



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
	paymentOptionsRequest := *openapiclient.NewPaymentOptionsRequest() // PaymentOptionsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentOptionsAPI.PostListPaymentOptions(context.Background()).PaymentOptionsRequest(paymentOptionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentOptionsAPI.PostListPaymentOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostListPaymentOptions`: PaymentOptions
	fmt.Fprintf(os.Stdout, "Response from `PaymentOptionsAPI.PostListPaymentOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostListPaymentOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentOptionsRequest** | [**PaymentOptionsRequest**](PaymentOptionsRequest.md) |  | 

### Return type

[**PaymentOptions**](PaymentOptions.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

