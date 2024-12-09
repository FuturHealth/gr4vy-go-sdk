# \AuditLogsAPI

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAuditLogs**](AuditLogsAPI.md#ListAuditLogs) | **Get** /audit-logs | List audit logs



## ListAuditLogs

> AuditLogs ListAuditLogs(ctx).Limit(limit).Cursor(cursor).UserId(userId).Action(action).ResourceType(resourceType).Execute()

List audit logs



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
	userId := "8724fd24-5489-4a5d-90fd-0604df7d3b83" // string | Filters the results to only the items for which the `user` has an `id` that matches this value. (optional)
	action := "created" // string | Filters the results to only the items for which the `audit-log` has an `action` that matches this value. (optional)
	resourceType := "buyer" // string | Filters the results to only the items for which the `audit-log` has a `resource` that matches this type value. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditLogsAPI.ListAuditLogs(context.Background()).Limit(limit).Cursor(cursor).UserId(userId).Action(action).ResourceType(resourceType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsAPI.ListAuditLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAuditLogs`: AuditLogs
	fmt.Fprintf(os.Stdout, "Response from `AuditLogsAPI.ListAuditLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAuditLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Defines the maximum number of items to return for this request. | [default to 20]
 **cursor** | **string** | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the &#x60;next_cursor&#x60; field. Similarly the &#x60;previous_cursor&#x60; can be used to reverse backwards in the list. | 
 **userId** | **string** | Filters the results to only the items for which the &#x60;user&#x60; has an &#x60;id&#x60; that matches this value. | 
 **action** | **string** | Filters the results to only the items for which the &#x60;audit-log&#x60; has an &#x60;action&#x60; that matches this value. | 
 **resourceType** | **string** | Filters the results to only the items for which the &#x60;audit-log&#x60; has a &#x60;resource&#x60; that matches this type value. | 

### Return type

[**AuditLogs**](AuditLogs.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

