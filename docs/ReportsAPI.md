# \ReportsAPI

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateDownloadUrl**](ReportsAPI.md#GenerateDownloadUrl) | **Post** /reports/{report_id}/executions/{report_execution_id}/url | Generate report download URL
[**GetReport**](ReportsAPI.md#GetReport) | **Get** /reports/{report_id} | Get report
[**GetReportExecution**](ReportsAPI.md#GetReportExecution) | **Get** /report-executions/{report_execution_id} | Get report execution
[**ListAllReportExecutions**](ReportsAPI.md#ListAllReportExecutions) | **Get** /report-executions | List all report executions
[**ListReportExecutions**](ReportsAPI.md#ListReportExecutions) | **Get** /reports/{report_id}/executions | List executions for report
[**ListReports**](ReportsAPI.md#ListReports) | **Get** /reports | List reports
[**NewReport**](ReportsAPI.md#NewReport) | **Post** /reports | New report
[**UpdateReport**](ReportsAPI.md#UpdateReport) | **Put** /reports/{report_id} | Update report



## GenerateDownloadUrl

> ReportExecutionUrl GenerateDownloadUrl(ctx, reportId, reportExecutionId).Execute()

Generate report download URL



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
	reportId := "8724fd24-5489-4a5d-90fd-0604df7d3b83" // string | The unique ID for a report.
	reportExecutionId := "8724fd24-5489-4a5d-90fd-0604df7d3b83" // string | The unique ID for a report execution.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GenerateDownloadUrl(context.Background(), reportId, reportExecutionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GenerateDownloadUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateDownloadUrl`: ReportExecutionUrl
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GenerateDownloadUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** | The unique ID for a report. | 
**reportExecutionId** | **string** | The unique ID for a report execution. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateDownloadUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ReportExecutionUrl**](ReportExecutionUrl.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReport

> Report GetReport(ctx, reportId).Execute()

Get report



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
	reportId := "8724fd24-5489-4a5d-90fd-0604df7d3b83" // string | The unique ID for a report.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetReport(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReport`: Report
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** | The unique ID for a report. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Report**](Report.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportExecution

> ReportExecution GetReportExecution(ctx, reportExecutionId).Execute()

Get report execution



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
	reportExecutionId := "8724fd24-5489-4a5d-90fd-0604df7d3b83" // string | The unique ID for a report execution.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.GetReportExecution(context.Background(), reportExecutionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportExecution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportExecution`: ReportExecution
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportExecution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportExecutionId** | **string** | The unique ID for a report execution. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportExecutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReportExecution**](ReportExecution.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllReportExecutions

> ReportExecutions ListAllReportExecutions(ctx).Cursor(cursor).Limit(limit).CreatedAtGte(createdAtGte).CreatedAtLte(createdAtLte).ReportName(reportName).Status(status).CreatorId(creatorId).Execute()

List all report executions



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
	cursor := "ZXhhbXBsZTE" // string | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the `next_cursor` field. Similarly the `previous_cursor` can be used to reverse backwards in the list. (optional)
	limit := int32(1) // int32 | Defines the maximum number of items to return for this request. (optional) (default to 20)
	createdAtGte := "2022-01-01T12:00:00+08:00" // string | Filters the results to report executions created after this ISO date-time string. The time zone must be included.  Ensure that the date-time string is URL encoded, e.g. `2022-01-01T12:00:00+08:00` must be encoded as `2022-01-01T12%3A00%3A00%2B08%3A00`. (optional)
	createdAtLte := "2022-01-01T12:00:00+08:00" // string | Filters the results to report executions created before this ISO date-time string. The time zone must be included.  Ensure that the date-time string is URL encoded, e.g. `2022-01-01T12:00:00+08:00` must be encoded as `2022-01-01T12%3A00%3A00%2B08%3A00`. (optional)
	reportName := "Failed+Authorizations+042022" // string | Filters for executions of reports that have a matching `name` value. This filter is case-insensitive.  Ensure that when necessary, the value you pass for this filter is URL encoded. (optional)
	status := []string{"succeeded"} // []string | Filters for report executions that have a matching `status` value.  This filter accepts multiple values. (optional)
	creatorId := []string{"Inner_example"} // []string | Filters the results to only match the reports that their `creator_id` matches with any of the provided creator IDs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ListAllReportExecutions(context.Background()).Cursor(cursor).Limit(limit).CreatedAtGte(createdAtGte).CreatedAtLte(createdAtLte).ReportName(reportName).Status(status).CreatorId(creatorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ListAllReportExecutions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllReportExecutions`: ReportExecutions
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ListAllReportExecutions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAllReportExecutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the &#x60;next_cursor&#x60; field. Similarly the &#x60;previous_cursor&#x60; can be used to reverse backwards in the list. | 
 **limit** | **int32** | Defines the maximum number of items to return for this request. | [default to 20]
 **createdAtGte** | **string** | Filters the results to report executions created after this ISO date-time string. The time zone must be included.  Ensure that the date-time string is URL encoded, e.g. &#x60;2022-01-01T12:00:00+08:00&#x60; must be encoded as &#x60;2022-01-01T12%3A00%3A00%2B08%3A00&#x60;. | 
 **createdAtLte** | **string** | Filters the results to report executions created before this ISO date-time string. The time zone must be included.  Ensure that the date-time string is URL encoded, e.g. &#x60;2022-01-01T12:00:00+08:00&#x60; must be encoded as &#x60;2022-01-01T12%3A00%3A00%2B08%3A00&#x60;. | 
 **reportName** | **string** | Filters for executions of reports that have a matching &#x60;name&#x60; value. This filter is case-insensitive.  Ensure that when necessary, the value you pass for this filter is URL encoded. | 
 **status** | **[]string** | Filters for report executions that have a matching &#x60;status&#x60; value.  This filter accepts multiple values. | 
 **creatorId** | **[]string** | Filters the results to only match the reports that their &#x60;creator_id&#x60; matches with any of the provided creator IDs. | 

### Return type

[**ReportExecutions**](ReportExecutions.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListReportExecutions

> ReportExecutions ListReportExecutions(ctx, reportId).Cursor(cursor).Limit(limit).Execute()

List executions for report



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
	reportId := "8724fd24-5489-4a5d-90fd-0604df7d3b83" // string | The unique ID for a report.
	cursor := "ZXhhbXBsZTE" // string | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the `next_cursor` field. Similarly the `previous_cursor` can be used to reverse backwards in the list. (optional)
	limit := int32(1) // int32 | Defines the maximum number of items to return for this request. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ListReportExecutions(context.Background(), reportId).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ListReportExecutions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListReportExecutions`: ReportExecutions
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ListReportExecutions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** | The unique ID for a report. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListReportExecutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the &#x60;next_cursor&#x60; field. Similarly the &#x60;previous_cursor&#x60; can be used to reverse backwards in the list. | 
 **limit** | **int32** | Defines the maximum number of items to return for this request. | [default to 20]

### Return type

[**ReportExecutions**](ReportExecutions.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListReports

> Reports ListReports(ctx).Cursor(cursor).Limit(limit).Name(name).Schedule(schedule).ScheduleEnabled(scheduleEnabled).Execute()

List reports



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
	cursor := "ZXhhbXBsZTE" // string | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the `next_cursor` field. Similarly the `previous_cursor` can be used to reverse backwards in the list. (optional)
	limit := int32(1) // int32 | Defines the maximum number of items to return for this request. (optional) (default to 20)
	name := "Failed+Authorizations+042022" // string | Filters for reports that have a matching `name` value. This filter is case-insensitive.  Ensure that when necessary, the value you pass for this filter is URL encoded. (optional)
	schedule := []string{"Schedule_example"} // []string | Filters for reports that have matching `schedule` values. (optional)
	scheduleEnabled := true // bool | Filters for reports that have a matching `schedule_enabled` value. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ListReports(context.Background()).Cursor(cursor).Limit(limit).Name(name).Schedule(schedule).ScheduleEnabled(scheduleEnabled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ListReports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListReports`: Reports
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ListReports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the &#x60;next_cursor&#x60; field. Similarly the &#x60;previous_cursor&#x60; can be used to reverse backwards in the list. | 
 **limit** | **int32** | Defines the maximum number of items to return for this request. | [default to 20]
 **name** | **string** | Filters for reports that have a matching &#x60;name&#x60; value. This filter is case-insensitive.  Ensure that when necessary, the value you pass for this filter is URL encoded. | 
 **schedule** | **[]string** | Filters for reports that have matching &#x60;schedule&#x60; values. | 
 **scheduleEnabled** | **bool** | Filters for reports that have a matching &#x60;schedule_enabled&#x60; value. | 

### Return type

[**Reports**](Reports.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewReport

> Report NewReport(ctx).ReportCreate(reportCreate).Execute()

New report



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
	reportCreate := *openapiclient.NewReportCreate("Failed Authorizations 042022", *openapiclient.NewReportSpec("transactions", map[string]interface{}(123))) // ReportCreate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.NewReport(context.Background()).ReportCreate(reportCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.NewReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NewReport`: Report
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.NewReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportCreate** | [**ReportCreate**](ReportCreate.md) |  | 

### Return type

[**Report**](Report.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReport

> Report UpdateReport(ctx, reportId).ReportUpdate(reportUpdate).Execute()

Update report



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
	reportId := "8724fd24-5489-4a5d-90fd-0604df7d3b83" // string | The unique ID for a report.
	reportUpdate := *openapiclient.NewReportUpdate() // ReportUpdate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.UpdateReport(context.Background(), reportId).ReportUpdate(reportUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.UpdateReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateReport`: Report
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.UpdateReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** | The unique ID for a report. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportUpdate** | [**ReportUpdate**](ReportUpdate.md) |  | 

### Return type

[**Report**](Report.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

