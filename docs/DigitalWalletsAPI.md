# \DigitalWalletsAPI

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDigitalWalletDomainName**](DigitalWalletsAPI.md#AddDigitalWalletDomainName) | **Post** /digital-wallets/{digital_wallet_id}/domains | Add digital wallet domain name
[**DeleteDigitalWallet**](DigitalWalletsAPI.md#DeleteDigitalWallet) | **Delete** /digital-wallets/{digital_wallet_id} | De-register digital wallet
[**DeleteDigitalWalletDomainName**](DigitalWalletsAPI.md#DeleteDigitalWalletDomainName) | **Delete** /digital-wallets/{digital_wallet_id}/domains | Remove digital wallet domain name
[**GetDigitalWallet**](DigitalWalletsAPI.md#GetDigitalWallet) | **Get** /digital-wallets/{digital_wallet_id} | Get digital wallet
[**ListDigitalWallets**](DigitalWalletsAPI.md#ListDigitalWallets) | **Get** /digital-wallets | List digital wallets
[**NewDigitalWallet**](DigitalWalletsAPI.md#NewDigitalWallet) | **Post** /digital-wallets | Register digital wallet
[**UpdateDigitalWallet**](DigitalWalletsAPI.md#UpdateDigitalWallet) | **Put** /digital-wallets/{digital_wallet_id} | Update digital wallet



## AddDigitalWalletDomainName

> AddDigitalWalletDomainName(ctx, digitalWalletId).DigitalWalletDomain(digitalWalletDomain).Execute()

Add digital wallet domain name



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
	digitalWalletId := "fe26475d-ec3e-4884-9553-f7356683f7f9" // string | The ID of the registered digital wallet.
	digitalWalletDomain := *openapiclient.NewDigitalWalletDomain("example.com") // DigitalWalletDomain |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DigitalWalletsAPI.AddDigitalWalletDomainName(context.Background(), digitalWalletId).DigitalWalletDomain(digitalWalletDomain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletsAPI.AddDigitalWalletDomainName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**digitalWalletId** | **string** | The ID of the registered digital wallet. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDigitalWalletDomainNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **digitalWalletDomain** | [**DigitalWalletDomain**](DigitalWalletDomain.md) |  | 

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


## DeleteDigitalWallet

> DeleteDigitalWallet(ctx, digitalWalletId).Execute()

De-register digital wallet



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
	digitalWalletId := "fe26475d-ec3e-4884-9553-f7356683f7f9" // string | The ID of the registered digital wallet.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DigitalWalletsAPI.DeleteDigitalWallet(context.Background(), digitalWalletId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletsAPI.DeleteDigitalWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**digitalWalletId** | **string** | The ID of the registered digital wallet. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDigitalWalletRequest struct via the builder pattern


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


## DeleteDigitalWalletDomainName

> DeleteDigitalWalletDomainName(ctx, digitalWalletId).DigitalWalletDomain(digitalWalletDomain).Execute()

Remove digital wallet domain name



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
	digitalWalletId := "fe26475d-ec3e-4884-9553-f7356683f7f9" // string | The ID of the registered digital wallet.
	digitalWalletDomain := *openapiclient.NewDigitalWalletDomain("example.com") // DigitalWalletDomain |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DigitalWalletsAPI.DeleteDigitalWalletDomainName(context.Background(), digitalWalletId).DigitalWalletDomain(digitalWalletDomain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletsAPI.DeleteDigitalWalletDomainName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**digitalWalletId** | **string** | The ID of the registered digital wallet. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDigitalWalletDomainNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **digitalWalletDomain** | [**DigitalWalletDomain**](DigitalWalletDomain.md) |  | 

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


## GetDigitalWallet

> DigitalWallet GetDigitalWallet(ctx, digitalWalletId).Execute()

Get digital wallet



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
	digitalWalletId := "fe26475d-ec3e-4884-9553-f7356683f7f9" // string | The ID of the registered digital wallet.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DigitalWalletsAPI.GetDigitalWallet(context.Background(), digitalWalletId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletsAPI.GetDigitalWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDigitalWallet`: DigitalWallet
	fmt.Fprintf(os.Stdout, "Response from `DigitalWalletsAPI.GetDigitalWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**digitalWalletId** | **string** | The ID of the registered digital wallet. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDigitalWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DigitalWallet**](DigitalWallet.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDigitalWallets

> DigitalWallets ListDigitalWallets(ctx).Execute()

List digital wallets



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DigitalWalletsAPI.ListDigitalWallets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletsAPI.ListDigitalWallets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDigitalWallets`: DigitalWallets
	fmt.Fprintf(os.Stdout, "Response from `DigitalWalletsAPI.ListDigitalWallets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDigitalWalletsRequest struct via the builder pattern


### Return type

[**DigitalWallets**](DigitalWallets.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewDigitalWallet

> DigitalWallet NewDigitalWallet(ctx).DigitalWalletRequest(digitalWalletRequest).Execute()

Register digital wallet



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
	digitalWalletRequest := *openapiclient.NewDigitalWalletRequest("apple", "Gr4vy", []string{"DomainNames_example"}, true) // DigitalWalletRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DigitalWalletsAPI.NewDigitalWallet(context.Background()).DigitalWalletRequest(digitalWalletRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletsAPI.NewDigitalWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NewDigitalWallet`: DigitalWallet
	fmt.Fprintf(os.Stdout, "Response from `DigitalWalletsAPI.NewDigitalWallet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewDigitalWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **digitalWalletRequest** | [**DigitalWalletRequest**](DigitalWalletRequest.md) |  | 

### Return type

[**DigitalWallet**](DigitalWallet.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDigitalWallet

> DigitalWallet UpdateDigitalWallet(ctx, digitalWalletId).DigitalWalletUpdate(digitalWalletUpdate).Execute()

Update digital wallet



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
	digitalWalletId := "fe26475d-ec3e-4884-9553-f7356683f7f9" // string | The ID of the registered digital wallet.
	digitalWalletUpdate := *openapiclient.NewDigitalWalletUpdate() // DigitalWalletUpdate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DigitalWalletsAPI.UpdateDigitalWallet(context.Background(), digitalWalletId).DigitalWalletUpdate(digitalWalletUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletsAPI.UpdateDigitalWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDigitalWallet`: DigitalWallet
	fmt.Fprintf(os.Stdout, "Response from `DigitalWalletsAPI.UpdateDigitalWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**digitalWalletId** | **string** | The ID of the registered digital wallet. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDigitalWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **digitalWalletUpdate** | [**DigitalWalletUpdate**](DigitalWalletUpdate.md) |  | 

### Return type

[**DigitalWallet**](DigitalWallet.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

