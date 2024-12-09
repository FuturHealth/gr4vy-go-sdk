# \CardSchemeDefinitionsAPI

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCardSchemeDefinitions**](CardSchemeDefinitionsAPI.md#ListCardSchemeDefinitions) | **Get** /card-scheme-definitions | List card scheme definitions



## ListCardSchemeDefinitions

> CardSchemeDefinitions ListCardSchemeDefinitions(ctx).Execute()

List card scheme definitions



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
	resp, r, err := apiClient.CardSchemeDefinitionsAPI.ListCardSchemeDefinitions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CardSchemeDefinitionsAPI.ListCardSchemeDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCardSchemeDefinitions`: CardSchemeDefinitions
	fmt.Fprintf(os.Stdout, "Response from `CardSchemeDefinitionsAPI.ListCardSchemeDefinitions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCardSchemeDefinitionsRequest struct via the builder pattern


### Return type

[**CardSchemeDefinitions**](CardSchemeDefinitions.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

