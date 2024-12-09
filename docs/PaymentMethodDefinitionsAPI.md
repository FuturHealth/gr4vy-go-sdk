# \PaymentMethodDefinitionsAPI

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPaymentMethodDefinitions**](PaymentMethodDefinitionsAPI.md#ListPaymentMethodDefinitions) | **Get** /payment-method-definitions | List payment method definitions



## ListPaymentMethodDefinitions

> PaymentMethodDefinitions ListPaymentMethodDefinitions(ctx).Execute()

List payment method definitions



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentMethodDefinitionsAPI.ListPaymentMethodDefinitions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodDefinitionsAPI.ListPaymentMethodDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPaymentMethodDefinitions`: PaymentMethodDefinitions
	fmt.Fprintf(os.Stdout, "Response from `PaymentMethodDefinitionsAPI.ListPaymentMethodDefinitions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPaymentMethodDefinitionsRequest struct via the builder pattern


### Return type

[**PaymentMethodDefinitions**](PaymentMethodDefinitions.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

