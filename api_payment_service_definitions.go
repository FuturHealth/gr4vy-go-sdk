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
	"strings"
)


// PaymentServiceDefinitionsAPIService PaymentServiceDefinitionsAPI service
type PaymentServiceDefinitionsAPIService service

type ApiCreatePaymentServiceDefinitionSessionRequest struct {
	ctx context.Context
	ApiService *PaymentServiceDefinitionsAPIService
	paymentServiceDefinitionId string
	requestBody *map[string]interface{}
}

func (r ApiCreatePaymentServiceDefinitionSessionRequest) RequestBody(requestBody map[string]interface{}) ApiCreatePaymentServiceDefinitionSessionRequest {
	r.requestBody = &requestBody
	return r
}

func (r ApiCreatePaymentServiceDefinitionSessionRequest) Execute() (*PaymentServiceSession, *http.Response, error) {
	return r.ApiService.CreatePaymentServiceDefinitionSessionExecute(r)
}

/*
CreatePaymentServiceDefinitionSession Create a session for a payment service

Creates a session for a payment service. This endpoint directly
passes the request through to the relevant payment service for processing,
and so the schema will differ based on the service used.

If the downstream service returns an error, this API will return a successful response
with the status code in the response.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param paymentServiceDefinitionId The unique ID of the payment service definition.
 @return ApiCreatePaymentServiceDefinitionSessionRequest
*/
func (a *PaymentServiceDefinitionsAPIService) CreatePaymentServiceDefinitionSession(ctx context.Context, paymentServiceDefinitionId string) ApiCreatePaymentServiceDefinitionSessionRequest {
	return ApiCreatePaymentServiceDefinitionSessionRequest{
		ApiService: a,
		ctx: ctx,
		paymentServiceDefinitionId: paymentServiceDefinitionId,
	}
}

// Execute executes the request
//  @return PaymentServiceSession
func (a *PaymentServiceDefinitionsAPIService) CreatePaymentServiceDefinitionSessionExecute(r ApiCreatePaymentServiceDefinitionSessionRequest) (*PaymentServiceSession, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaymentServiceSession
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentServiceDefinitionsAPIService.CreatePaymentServiceDefinitionSession")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment-service-definitions/{payment_service_definition_id}/sessions"
	localVarPath = strings.Replace(localVarPath, "{"+"payment_service_definition_id"+"}", url.PathEscape(parameterValueToString(r.paymentServiceDefinitionId, "paymentServiceDefinitionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.requestBody
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error400BadRequest
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error404NotFound
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

type ApiGetPaymentServiceDefinitionRequest struct {
	ctx context.Context
	ApiService *PaymentServiceDefinitionsAPIService
	paymentServiceDefinitionId string
}

func (r ApiGetPaymentServiceDefinitionRequest) Execute() (*PaymentServiceDefinition, *http.Response, error) {
	return r.ApiService.GetPaymentServiceDefinitionExecute(r)
}

/*
GetPaymentServiceDefinition Get payment service definition

Gets the definition for a payment service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param paymentServiceDefinitionId The unique ID of the payment service definition.
 @return ApiGetPaymentServiceDefinitionRequest
*/
func (a *PaymentServiceDefinitionsAPIService) GetPaymentServiceDefinition(ctx context.Context, paymentServiceDefinitionId string) ApiGetPaymentServiceDefinitionRequest {
	return ApiGetPaymentServiceDefinitionRequest{
		ApiService: a,
		ctx: ctx,
		paymentServiceDefinitionId: paymentServiceDefinitionId,
	}
}

// Execute executes the request
//  @return PaymentServiceDefinition
func (a *PaymentServiceDefinitionsAPIService) GetPaymentServiceDefinitionExecute(r ApiGetPaymentServiceDefinitionRequest) (*PaymentServiceDefinition, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaymentServiceDefinition
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentServiceDefinitionsAPIService.GetPaymentServiceDefinition")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment-service-definitions/{payment_service_definition_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"payment_service_definition_id"+"}", url.PathEscape(parameterValueToString(r.paymentServiceDefinitionId, "paymentServiceDefinitionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error404NotFound
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

type ApiListPaymentServiceDefinitionsRequest struct {
	ctx context.Context
	ApiService *PaymentServiceDefinitionsAPIService
	limit *int32
	cursor *string
}

// Defines the maximum number of items to return for this request.
func (r ApiListPaymentServiceDefinitionsRequest) Limit(limit int32) ApiListPaymentServiceDefinitionsRequest {
	r.limit = &limit
	return r
}

// A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the &#x60;next_cursor&#x60; field. Similarly the &#x60;previous_cursor&#x60; can be used to reverse backwards in the list.
func (r ApiListPaymentServiceDefinitionsRequest) Cursor(cursor string) ApiListPaymentServiceDefinitionsRequest {
	r.cursor = &cursor
	return r
}

func (r ApiListPaymentServiceDefinitionsRequest) Execute() (*PaymentServiceDefinitions, *http.Response, error) {
	return r.ApiService.ListPaymentServiceDefinitionsExecute(r)
}

/*
ListPaymentServiceDefinitions List payment service definitions

Returns a list of all available payment service definitions.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListPaymentServiceDefinitionsRequest
*/
func (a *PaymentServiceDefinitionsAPIService) ListPaymentServiceDefinitions(ctx context.Context) ApiListPaymentServiceDefinitionsRequest {
	return ApiListPaymentServiceDefinitionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaymentServiceDefinitions
func (a *PaymentServiceDefinitionsAPIService) ListPaymentServiceDefinitionsExecute(r ApiListPaymentServiceDefinitionsRequest) (*PaymentServiceDefinitions, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaymentServiceDefinitions
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentServiceDefinitionsAPIService.ListPaymentServiceDefinitions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment-service-definitions"

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
