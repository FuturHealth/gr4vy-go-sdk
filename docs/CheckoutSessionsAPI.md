# \CheckoutSessionsAPI

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCheckoutSession**](CheckoutSessionsAPI.md#DeleteCheckoutSession) | **Delete** /checkout/sessions/{checkout_session_id} | Delete checkout session
[**GetCheckoutSession**](CheckoutSessionsAPI.md#GetCheckoutSession) | **Get** /checkout/sessions/{checkout_session_id} | Get checkout session
[**NewCheckoutSession**](CheckoutSessionsAPI.md#NewCheckoutSession) | **Post** /checkout/sessions | New checkout session
[**UpdateCheckoutSession**](CheckoutSessionsAPI.md#UpdateCheckoutSession) | **Put** /checkout/sessions/{checkout_session_id} | Update checkout session



## DeleteCheckoutSession

> DeleteCheckoutSession(ctx, checkoutSessionId).Execute()

Delete checkout session



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
	checkoutSessionId := "8724fd24-5489-4a5d-90fd-0604df7d3b83" // string | The unique ID for a Checkout Session.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CheckoutSessionsAPI.DeleteCheckoutSession(context.Background(), checkoutSessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckoutSessionsAPI.DeleteCheckoutSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkoutSessionId** | **string** | The unique ID for a Checkout Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCheckoutSessionRequest struct via the builder pattern


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


## GetCheckoutSession

> CheckoutSession GetCheckoutSession(ctx, checkoutSessionId).Execute()

Get checkout session



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
	checkoutSessionId := "8724fd24-5489-4a5d-90fd-0604df7d3b83" // string | The unique ID for a Checkout Session.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckoutSessionsAPI.GetCheckoutSession(context.Background(), checkoutSessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckoutSessionsAPI.GetCheckoutSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCheckoutSession`: CheckoutSession
	fmt.Fprintf(os.Stdout, "Response from `CheckoutSessionsAPI.GetCheckoutSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkoutSessionId** | **string** | The unique ID for a Checkout Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCheckoutSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CheckoutSession**](CheckoutSession.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewCheckoutSession

> CheckoutSession NewCheckoutSession(ctx).CheckoutSessionCreateRequest(checkoutSessionCreateRequest).Execute()

New checkout session



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
	checkoutSessionCreateRequest := *openapiclient.NewCheckoutSessionCreateRequest() // CheckoutSessionCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckoutSessionsAPI.NewCheckoutSession(context.Background()).CheckoutSessionCreateRequest(checkoutSessionCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckoutSessionsAPI.NewCheckoutSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NewCheckoutSession`: CheckoutSession
	fmt.Fprintf(os.Stdout, "Response from `CheckoutSessionsAPI.NewCheckoutSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewCheckoutSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkoutSessionCreateRequest** | [**CheckoutSessionCreateRequest**](CheckoutSessionCreateRequest.md) |  | 

### Return type

[**CheckoutSession**](CheckoutSession.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCheckoutSession

> CheckoutSession UpdateCheckoutSession(ctx, checkoutSessionId).CheckoutSessionUpdateRequest(checkoutSessionUpdateRequest).Execute()

Update checkout session



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
	checkoutSessionId := "8724fd24-5489-4a5d-90fd-0604df7d3b83" // string | The unique ID for a Checkout Session.
	checkoutSessionUpdateRequest := *openapiclient.NewCheckoutSessionUpdateRequest() // CheckoutSessionUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckoutSessionsAPI.UpdateCheckoutSession(context.Background(), checkoutSessionId).CheckoutSessionUpdateRequest(checkoutSessionUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckoutSessionsAPI.UpdateCheckoutSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCheckoutSession`: CheckoutSession
	fmt.Fprintf(os.Stdout, "Response from `CheckoutSessionsAPI.UpdateCheckoutSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkoutSessionId** | **string** | The unique ID for a Checkout Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCheckoutSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **checkoutSessionUpdateRequest** | [**CheckoutSessionUpdateRequest**](CheckoutSessionUpdateRequest.md) |  | 

### Return type

[**CheckoutSession**](CheckoutSession.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

