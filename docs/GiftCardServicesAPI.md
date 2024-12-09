# \GiftCardServicesAPI

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteGiftCardService**](GiftCardServicesAPI.md#DeleteGiftCardService) | **Delete** /gift-card-services/{gift_card_service_id} | Delete gift card service
[**GetGiftCardService**](GiftCardServicesAPI.md#GetGiftCardService) | **Get** /gift-card-services/{gift_card_service_id} | Get gift card service
[**NewGiftCardService**](GiftCardServicesAPI.md#NewGiftCardService) | **Post** /gift-card-services | New gift card service
[**UpdateGiftCardService**](GiftCardServicesAPI.md#UpdateGiftCardService) | **Put** /gift-card-services/{gift_card_service_id} | Update gift card service
[**VerifyGiftCardService**](GiftCardServicesAPI.md#VerifyGiftCardService) | **Post** /gift-card-services/verify | Verify gift card service credentials



## DeleteGiftCardService

> DeleteGiftCardService(ctx, giftCardServiceId).Execute()

Delete gift card service



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
	giftCardServiceId := "541b126f-44c5-4c5e-a06b-d0e0d54c7d3f" // string | The unique ID of the gift card service.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GiftCardServicesAPI.DeleteGiftCardService(context.Background(), giftCardServiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GiftCardServicesAPI.DeleteGiftCardService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**giftCardServiceId** | **string** | The unique ID of the gift card service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGiftCardServiceRequest struct via the builder pattern


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


## GetGiftCardService

> GiftCardService GetGiftCardService(ctx, giftCardServiceId).Execute()

Get gift card service



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
	giftCardServiceId := "541b126f-44c5-4c5e-a06b-d0e0d54c7d3f" // string | The unique ID of the gift card service.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GiftCardServicesAPI.GetGiftCardService(context.Background(), giftCardServiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GiftCardServicesAPI.GetGiftCardService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGiftCardService`: GiftCardService
	fmt.Fprintf(os.Stdout, "Response from `GiftCardServicesAPI.GetGiftCardService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**giftCardServiceId** | **string** | The unique ID of the gift card service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGiftCardServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GiftCardService**](GiftCardService.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewGiftCardService

> GiftCardService NewGiftCardService(ctx).GiftCardServiceCreateRequest(giftCardServiceCreateRequest).Execute()

New gift card service



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
	giftCardServiceCreateRequest := *openapiclient.NewGiftCardServiceCreateRequest("qwikcilver-gift-card", "Qwikcilver UK", []openapiclient.GiftCardServiceCreateRequestFieldsInner{*openapiclient.NewGiftCardServiceCreateRequestFieldsInner("private_key", "pk_26PHem9AhJZvU623DfE1x4sd")}) // GiftCardServiceCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GiftCardServicesAPI.NewGiftCardService(context.Background()).GiftCardServiceCreateRequest(giftCardServiceCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GiftCardServicesAPI.NewGiftCardService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NewGiftCardService`: GiftCardService
	fmt.Fprintf(os.Stdout, "Response from `GiftCardServicesAPI.NewGiftCardService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewGiftCardServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **giftCardServiceCreateRequest** | [**GiftCardServiceCreateRequest**](GiftCardServiceCreateRequest.md) |  | 

### Return type

[**GiftCardService**](GiftCardService.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGiftCardService

> GiftCardService UpdateGiftCardService(ctx, giftCardServiceId).GiftCardServiceUpdateRequest(giftCardServiceUpdateRequest).Execute()

Update gift card service



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
	giftCardServiceId := "541b126f-44c5-4c5e-a06b-d0e0d54c7d3f" // string | The unique ID of the gift card service.
	giftCardServiceUpdateRequest := *openapiclient.NewGiftCardServiceUpdateRequest() // GiftCardServiceUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GiftCardServicesAPI.UpdateGiftCardService(context.Background(), giftCardServiceId).GiftCardServiceUpdateRequest(giftCardServiceUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GiftCardServicesAPI.UpdateGiftCardService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGiftCardService`: GiftCardService
	fmt.Fprintf(os.Stdout, "Response from `GiftCardServicesAPI.UpdateGiftCardService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**giftCardServiceId** | **string** | The unique ID of the gift card service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGiftCardServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **giftCardServiceUpdateRequest** | [**GiftCardServiceUpdateRequest**](GiftCardServiceUpdateRequest.md) |  | 

### Return type

[**GiftCardService**](GiftCardService.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyGiftCardService

> VerifyGiftCardService(ctx).GiftCardServiceVerifyRequest(giftCardServiceVerifyRequest).Execute()

Verify gift card service credentials



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
	giftCardServiceVerifyRequest := *openapiclient.NewGiftCardServiceVerifyRequest("qwikcilver-gift-card", []openapiclient.GiftCardServiceVerifyRequestFieldsInner{*openapiclient.NewGiftCardServiceVerifyRequestFieldsInner("private_key", "pk_26PHem9AhJZvU623DfE1x4sd")}) // GiftCardServiceVerifyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GiftCardServicesAPI.VerifyGiftCardService(context.Background()).GiftCardServiceVerifyRequest(giftCardServiceVerifyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GiftCardServicesAPI.VerifyGiftCardService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyGiftCardServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **giftCardServiceVerifyRequest** | [**GiftCardServiceVerifyRequest**](GiftCardServiceVerifyRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

