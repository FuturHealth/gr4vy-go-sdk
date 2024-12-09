# \BuyersAPI

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBuyer**](BuyersAPI.md#DeleteBuyer) | **Delete** /buyers/{buyer_id} | Delete buyer
[**DeleteBuyerShippingDetail**](BuyersAPI.md#DeleteBuyerShippingDetail) | **Delete** /buyers/{buyer_id}/shipping-details/{shipping_detail_id} | Delete buyer shipping detail
[**GetBuyer**](BuyersAPI.md#GetBuyer) | **Get** /buyers/{buyer_id} | Get buyer
[**ListBuyerShippingDetails**](BuyersAPI.md#ListBuyerShippingDetails) | **Get** /buyers/{buyer_id}/shipping-details | List buyer shipping details
[**ListBuyers**](BuyersAPI.md#ListBuyers) | **Get** /buyers | List buyers
[**NewBuyer**](BuyersAPI.md#NewBuyer) | **Post** /buyers | New buyer
[**NewBuyerShippingDetail**](BuyersAPI.md#NewBuyerShippingDetail) | **Post** /buyers/{buyer_id}/shipping-details | New buyer shipping detail
[**UpdateBuyer**](BuyersAPI.md#UpdateBuyer) | **Put** /buyers/{buyer_id} | Update buyer
[**UpdateBuyerShippingDetail**](BuyersAPI.md#UpdateBuyerShippingDetail) | **Put** /buyers/{buyer_id}/shipping-details/{shipping_detail_id} | Update buyer shipping details



## DeleteBuyer

> DeleteBuyer(ctx, buyerId).Execute()

Delete buyer



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
	buyerId := "8724fd24-5489-4a5d-90fd-0604df7d3b83" // string | The unique ID for a buyer.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BuyersAPI.DeleteBuyer(context.Background(), buyerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuyersAPI.DeleteBuyer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buyerId** | **string** | The unique ID for a buyer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBuyerRequest struct via the builder pattern


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


## DeleteBuyerShippingDetail

> DeleteBuyerShippingDetail(ctx, buyerId, shippingDetailId).Execute()

Delete buyer shipping detail



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
	buyerId := "8724fd24-5489-4a5d-90fd-0604df7d3b83" // string | The unique ID for a buyer.
	shippingDetailId := "8724fd24-5489-4a5d-90fd-0604df7d3b83" // string | The unique ID for a buyer's shipping detail.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BuyersAPI.DeleteBuyerShippingDetail(context.Background(), buyerId, shippingDetailId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuyersAPI.DeleteBuyerShippingDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buyerId** | **string** | The unique ID for a buyer. | 
**shippingDetailId** | **string** | The unique ID for a buyer&#39;s shipping detail. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBuyerShippingDetailRequest struct via the builder pattern


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


## GetBuyer

> Buyer GetBuyer(ctx, buyerId).Execute()

Get buyer



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
	buyerId := "8724fd24-5489-4a5d-90fd-0604df7d3b83" // string | The unique ID for a buyer.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuyersAPI.GetBuyer(context.Background(), buyerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuyersAPI.GetBuyer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBuyer`: Buyer
	fmt.Fprintf(os.Stdout, "Response from `BuyersAPI.GetBuyer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buyerId** | **string** | The unique ID for a buyer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Buyer**](Buyer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBuyerShippingDetails

> ShippingDetails ListBuyerShippingDetails(ctx, buyerId).Execute()

List buyer shipping details



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
	buyerId := "8724fd24-5489-4a5d-90fd-0604df7d3b83" // string | The unique ID for a buyer.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuyersAPI.ListBuyerShippingDetails(context.Background(), buyerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuyersAPI.ListBuyerShippingDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBuyerShippingDetails`: ShippingDetails
	fmt.Fprintf(os.Stdout, "Response from `BuyersAPI.ListBuyerShippingDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buyerId** | **string** | The unique ID for a buyer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerShippingDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ShippingDetails**](ShippingDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBuyers

> Buyers ListBuyers(ctx).Search(search).ExternalIdentifier(externalIdentifier).Limit(limit).Cursor(cursor).Execute()

List buyers



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
	search := "John" // string | ***Please do not use this query parameter in a production application, as this API call is very inefficient and may negatively impact transaction performance.***  Filters the results to only the buyers for which the `display_name`, `email_address` or `external_identifier` matches this value. This field allows for a case insensitive partial match, matching any buyer for which any of the fields partially or completely matches. (optional)
	externalIdentifier := "user-12345" // string | Filters the results to only the items for which the `buyer` has an `external_identifier` that exactly matches this value. (optional)
	limit := int32(1) // int32 | Defines the maximum number of items to return for this request. (optional) (default to 20)
	cursor := "ZXhhbXBsZTE" // string | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the `next_cursor` field. Similarly the `previous_cursor` can be used to reverse backwards in the list. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuyersAPI.ListBuyers(context.Background()).Search(search).ExternalIdentifier(externalIdentifier).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuyersAPI.ListBuyers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBuyers`: Buyers
	fmt.Fprintf(os.Stdout, "Response from `BuyersAPI.ListBuyers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBuyersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | ***Please do not use this query parameter in a production application, as this API call is very inefficient and may negatively impact transaction performance.***  Filters the results to only the buyers for which the &#x60;display_name&#x60;, &#x60;email_address&#x60; or &#x60;external_identifier&#x60; matches this value. This field allows for a case insensitive partial match, matching any buyer for which any of the fields partially or completely matches. | 
 **externalIdentifier** | **string** | Filters the results to only the items for which the &#x60;buyer&#x60; has an &#x60;external_identifier&#x60; that exactly matches this value. | 
 **limit** | **int32** | Defines the maximum number of items to return for this request. | [default to 20]
 **cursor** | **string** | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the &#x60;next_cursor&#x60; field. Similarly the &#x60;previous_cursor&#x60; can be used to reverse backwards in the list. | 

### Return type

[**Buyers**](Buyers.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewBuyer

> Buyer NewBuyer(ctx).BuyerRequest(buyerRequest).Execute()

New buyer



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
	buyerRequest := *openapiclient.NewBuyerRequest() // BuyerRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuyersAPI.NewBuyer(context.Background()).BuyerRequest(buyerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuyersAPI.NewBuyer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NewBuyer`: Buyer
	fmt.Fprintf(os.Stdout, "Response from `BuyersAPI.NewBuyer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewBuyerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buyerRequest** | [**BuyerRequest**](BuyerRequest.md) |  | 

### Return type

[**Buyer**](Buyer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewBuyerShippingDetail

> ShippingDetail NewBuyerShippingDetail(ctx, buyerId).ShippingDetailRequest(shippingDetailRequest).Execute()

New buyer shipping detail



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
	buyerId := "8724fd24-5489-4a5d-90fd-0604df7d3b83" // string | The unique ID for a buyer.
	shippingDetailRequest := *openapiclient.NewShippingDetailRequest() // ShippingDetailRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuyersAPI.NewBuyerShippingDetail(context.Background(), buyerId).ShippingDetailRequest(shippingDetailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuyersAPI.NewBuyerShippingDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NewBuyerShippingDetail`: ShippingDetail
	fmt.Fprintf(os.Stdout, "Response from `BuyersAPI.NewBuyerShippingDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buyerId** | **string** | The unique ID for a buyer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNewBuyerShippingDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shippingDetailRequest** | [**ShippingDetailRequest**](ShippingDetailRequest.md) |  | 

### Return type

[**ShippingDetail**](ShippingDetail.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBuyer

> Buyer UpdateBuyer(ctx, buyerId).BuyerUpdate(buyerUpdate).Execute()

Update buyer



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
	buyerId := "8724fd24-5489-4a5d-90fd-0604df7d3b83" // string | The unique ID for a buyer.
	buyerUpdate := *openapiclient.NewBuyerUpdate() // BuyerUpdate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuyersAPI.UpdateBuyer(context.Background(), buyerId).BuyerUpdate(buyerUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuyersAPI.UpdateBuyer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBuyer`: Buyer
	fmt.Fprintf(os.Stdout, "Response from `BuyersAPI.UpdateBuyer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buyerId** | **string** | The unique ID for a buyer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBuyerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **buyerUpdate** | [**BuyerUpdate**](BuyerUpdate.md) |  | 

### Return type

[**Buyer**](Buyer.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBuyerShippingDetail

> ShippingDetail UpdateBuyerShippingDetail(ctx, buyerId, shippingDetailId).ShippingDetailUpdateRequest(shippingDetailUpdateRequest).Execute()

Update buyer shipping details



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
	buyerId := "8724fd24-5489-4a5d-90fd-0604df7d3b83" // string | The unique ID for a buyer.
	shippingDetailId := "8724fd24-5489-4a5d-90fd-0604df7d3b83" // string | The unique ID for a buyer's shipping detail.
	shippingDetailUpdateRequest := *openapiclient.NewShippingDetailUpdateRequest() // ShippingDetailUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuyersAPI.UpdateBuyerShippingDetail(context.Background(), buyerId, shippingDetailId).ShippingDetailUpdateRequest(shippingDetailUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuyersAPI.UpdateBuyerShippingDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBuyerShippingDetail`: ShippingDetail
	fmt.Fprintf(os.Stdout, "Response from `BuyersAPI.UpdateBuyerShippingDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buyerId** | **string** | The unique ID for a buyer. | 
**shippingDetailId** | **string** | The unique ID for a buyer&#39;s shipping detail. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBuyerShippingDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **shippingDetailUpdateRequest** | [**ShippingDetailUpdateRequest**](ShippingDetailUpdateRequest.md) |  | 

### Return type

[**ShippingDetail**](ShippingDetail.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

