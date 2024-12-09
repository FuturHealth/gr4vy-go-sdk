# \GiftCardServiceDefinitionsAPI

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGiftCardServiceDefinition**](GiftCardServiceDefinitionsAPI.md#GetGiftCardServiceDefinition) | **Get** /gift-card-service-definitions/{gift_card_service_definition_id} | Get gift card service definition



## GetGiftCardServiceDefinition

> GiftCardServiceDefinition GetGiftCardServiceDefinition(ctx, giftCardServiceDefinitionId).Execute()

Get gift card service definition



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
	giftCardServiceDefinitionId := "qwikcilver-gift-card" // string | The unique ID of the gift card service definition.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GiftCardServiceDefinitionsAPI.GetGiftCardServiceDefinition(context.Background(), giftCardServiceDefinitionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GiftCardServiceDefinitionsAPI.GetGiftCardServiceDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGiftCardServiceDefinition`: GiftCardServiceDefinition
	fmt.Fprintf(os.Stdout, "Response from `GiftCardServiceDefinitionsAPI.GetGiftCardServiceDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**giftCardServiceDefinitionId** | **string** | The unique ID of the gift card service definition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGiftCardServiceDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GiftCardServiceDefinition**](GiftCardServiceDefinition.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

