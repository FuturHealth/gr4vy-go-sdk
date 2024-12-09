/*
Gr4vy API

Welcome to the Gr4vy API reference documentation. Our API is still very much a work in product and subject to change.

API version: 1.1.0-beta
Contact: code@gr4vy.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// AuditLogsAPIService AuditLogsAPI service
type AuditLogsAPIService service

type ApiListAuditLogsRequest struct {
	ctx context.Context
	ApiService *AuditLogsAPIService
	limit *int32
	cursor *string
	userId *string
	action *string
	resourceType *string
}

// Defines the maximum number of items to return for this request.
func (r ApiListAuditLogsRequest) Limit(limit int32) ApiListAuditLogsRequest {
	r.limit = &limit
	return r
}

// A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the &#x60;next_cursor&#x60; field. Similarly the &#x60;previous_cursor&#x60; can be used to reverse backwards in the list.
func (r ApiListAuditLogsRequest) Cursor(cursor string) ApiListAuditLogsRequest {
	r.cursor = &cursor
	return r
}

// Filters the results to only the items for which the &#x60;user&#x60; has an &#x60;id&#x60; that matches this value.
func (r ApiListAuditLogsRequest) UserId(userId string) ApiListAuditLogsRequest {
	r.userId = &userId
	return r
}

// Filters the results to only the items for which the &#x60;audit-log&#x60; has an &#x60;action&#x60; that matches this value.
func (r ApiListAuditLogsRequest) Action(action string) ApiListAuditLogsRequest {
	r.action = &action
	return r
}

// Filters the results to only the items for which the &#x60;audit-log&#x60; has a &#x60;resource&#x60; that matches this type value.
func (r ApiListAuditLogsRequest) ResourceType(resourceType string) ApiListAuditLogsRequest {
	r.resourceType = &resourceType
	return r
}

func (r ApiListAuditLogsRequest) Execute() (*AuditLogs, *http.Response, error) {
	return r.ApiService.ListAuditLogsExecute(r)
}

/*
ListAuditLogs List audit logs

Returns a list of audit logs.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListAuditLogsRequest
*/
func (a *AuditLogsAPIService) ListAuditLogs(ctx context.Context) ApiListAuditLogsRequest {
	return ApiListAuditLogsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AuditLogs
func (a *AuditLogsAPIService) ListAuditLogsExecute(r ApiListAuditLogsRequest) (*AuditLogs, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuditLogs
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditLogsAPIService.ListAuditLogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/audit-logs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	} else {
		var defaultValue int32 = 20
		r.limit = &defaultValue
	}
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "form", "")
	}
	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "user_id", r.userId, "form", "")
	}
	if r.action != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "action", r.action, "form", "")
	}
	if r.resourceType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "resource_type", r.resourceType, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error401Unauthorized
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
