# \VaultForwardAPI

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MakeVaultForward**](VaultForwardAPI.md#MakeVaultForward) | **Post** /vault-forward | Forward PCI data



## MakeVaultForward

> string MakeVaultForward(ctx).XVaultForwardPaymentMethods(xVaultForwardPaymentMethods).XVaultForwardUrl(xVaultForwardUrl).XVaultForwardHttpMethod(xVaultForwardHttpMethod).XVaultForwardAuthentications(xVaultForwardAuthentications).XVaultForwardHeaderHEADERNAME(xVaultForwardHeaderHEADERNAME).XVaultForwardTimeout(xVaultForwardTimeout).Body(body).Execute()

Forward PCI data



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
	xVaultForwardPaymentMethods := "faaad066-30b4-4997-a438-242b0752d7e1,faaad066-30b4-4997-a438-242b0752d7e2" // string | A comma-separated list of Payment Method IDs that can be used to fill in the request template. At least 1 must be given, and a maximum of 100 are accepted.
	xVaultForwardUrl := "https://api.amadeus.com/booking" // string | The URL to forward card data to.
	xVaultForwardHttpMethod := "POST" // string | The HTTP method that is used when forwarding the request to the `x-vault-forward-url`.
	xVaultForwardAuthentications := "faaad066-30b4-4997-a438-242b0752d7e1,faaad066-30b4-4997-a438-242b0752d7e2" // string | A comma-separated list of IDs for the authentication methods that will be applied to a Vault Forward request. (optional)
	xVaultForwardHeaderHEADERNAME := "x-vault-forward-header-x-frame-options" // string | A header that is forwarded to the `x-vault-forward-url`. The header will be forwarded without the `x-vault-forward-header` part. For example, `x-vault-forward-header-x-frame-options: SAMEORIGIN` is forwarded as `x-frame-options: SAMEORIGIN`. (optional)
	xVaultForwardTimeout := int32(10) // int32 | The number of seconds to wait before timing out when forwarding the request. (optional) (default to 30)
	body := "{
  "number":"{{CARD_NUMBER_1}}",
  "expiry":"{{CARD_EXPIRATION_DATE_1}}"
}
" // string | Payload to forward in the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VaultForwardAPI.MakeVaultForward(context.Background()).XVaultForwardPaymentMethods(xVaultForwardPaymentMethods).XVaultForwardUrl(xVaultForwardUrl).XVaultForwardHttpMethod(xVaultForwardHttpMethod).XVaultForwardAuthentications(xVaultForwardAuthentications).XVaultForwardHeaderHEADERNAME(xVaultForwardHeaderHEADERNAME).XVaultForwardTimeout(xVaultForwardTimeout).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VaultForwardAPI.MakeVaultForward``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MakeVaultForward`: string
	fmt.Fprintf(os.Stdout, "Response from `VaultForwardAPI.MakeVaultForward`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMakeVaultForwardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xVaultForwardPaymentMethods** | **string** | A comma-separated list of Payment Method IDs that can be used to fill in the request template. At least 1 must be given, and a maximum of 100 are accepted. | 
 **xVaultForwardUrl** | **string** | The URL to forward card data to. | 
 **xVaultForwardHttpMethod** | **string** | The HTTP method that is used when forwarding the request to the &#x60;x-vault-forward-url&#x60;. | 
 **xVaultForwardAuthentications** | **string** | A comma-separated list of IDs for the authentication methods that will be applied to a Vault Forward request. | 
 **xVaultForwardHeaderHEADERNAME** | **string** | A header that is forwarded to the &#x60;x-vault-forward-url&#x60;. The header will be forwarded without the &#x60;x-vault-forward-header&#x60; part. For example, &#x60;x-vault-forward-header-x-frame-options: SAMEORIGIN&#x60; is forwarded as &#x60;x-frame-options: SAMEORIGIN&#x60;. | 
 **xVaultForwardTimeout** | **int32** | The number of seconds to wait before timing out when forwarding the request. | [default to 30]
 **body** | **string** | Payload to forward in the request. | 

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

