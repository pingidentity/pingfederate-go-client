/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// OauthTokenExchangeTokenGeneratorMappingsAPIService OauthTokenExchangeTokenGeneratorMappingsAPI service
type OauthTokenExchangeTokenGeneratorMappingsAPIService service

type ApiCreateTokenGeneratorMappingRequest struct {
	ctx                       context.Context
	ApiService                *OauthTokenExchangeTokenGeneratorMappingsAPIService
	body                      *ProcessorPolicyToGeneratorMapping
	xBypassExternalValidation *bool
}

// Configuration for a new Token Exchange Processor policy to Token Generator Mapping.
func (r ApiCreateTokenGeneratorMappingRequest) Body(body ProcessorPolicyToGeneratorMapping) ApiCreateTokenGeneratorMappingRequest {
	r.body = &body
	return r
}

// External validation will be bypassed when set to true. Default to false.
func (r ApiCreateTokenGeneratorMappingRequest) XBypassExternalValidation(xBypassExternalValidation bool) ApiCreateTokenGeneratorMappingRequest {
	r.xBypassExternalValidation = &xBypassExternalValidation
	return r
}

func (r ApiCreateTokenGeneratorMappingRequest) Execute() (*ProcessorPolicyToGeneratorMapping, *http.Response, error) {
	return r.ApiService.CreateTokenGeneratorMappingExecute(r)
}

/*
CreateTokenGeneratorMapping Create a new Token Exchange Processor policy to Token Generator Mapping.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateTokenGeneratorMappingRequest
*/
func (a *OauthTokenExchangeTokenGeneratorMappingsAPIService) CreateTokenGeneratorMapping(ctx context.Context) ApiCreateTokenGeneratorMappingRequest {
	return ApiCreateTokenGeneratorMappingRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ProcessorPolicyToGeneratorMapping
func (a *OauthTokenExchangeTokenGeneratorMappingsAPIService) CreateTokenGeneratorMappingExecute(r ApiCreateTokenGeneratorMappingRequest) (*ProcessorPolicyToGeneratorMapping, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *ProcessorPolicyToGeneratorMapping
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalCreateTokenGeneratorMappingExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *OauthTokenExchangeTokenGeneratorMappingsAPIService) internalCreateTokenGeneratorMappingExecute(r ApiCreateTokenGeneratorMappingRequest) (*ProcessorPolicyToGeneratorMapping, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProcessorPolicyToGeneratorMapping
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthTokenExchangeTokenGeneratorMappingsAPIService.CreateTokenGeneratorMapping")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/tokenExchange/tokenGeneratorMappings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

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
	if r.xBypassExternalValidation != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-BypassExternalValidation", r.xBypassExternalValidation, "")
	}
	// body params
	localVarPostBody = r.body
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
		if localVarHTTPResponse.StatusCode == 422 {
			var v ApiResult
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

type ApiDeleteTokenGeneratorMappingByIdRequest struct {
	ctx        context.Context
	ApiService *OauthTokenExchangeTokenGeneratorMappingsAPIService
	id         string
}

func (r ApiDeleteTokenGeneratorMappingByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteTokenGeneratorMappingByIdExecute(r)
}

/*
DeleteTokenGeneratorMappingById Delete a Token Exchange Processor policy to Token Generator Mapping.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ID of Token Exchange Processor policy to Token Generator Mapping to delete.
	@return ApiDeleteTokenGeneratorMappingByIdRequest
*/
func (a *OauthTokenExchangeTokenGeneratorMappingsAPIService) DeleteTokenGeneratorMappingById(ctx context.Context, id string) ApiDeleteTokenGeneratorMappingByIdRequest {
	return ApiDeleteTokenGeneratorMappingByIdRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *OauthTokenExchangeTokenGeneratorMappingsAPIService) DeleteTokenGeneratorMappingByIdExecute(r ApiDeleteTokenGeneratorMappingByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthTokenExchangeTokenGeneratorMappingsAPIService.DeleteTokenGeneratorMappingById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/tokenExchange/tokenGeneratorMappings/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiResult
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetTokenGeneratorMappingByIdRequest struct {
	ctx        context.Context
	ApiService *OauthTokenExchangeTokenGeneratorMappingsAPIService
	id         string
}

func (r ApiGetTokenGeneratorMappingByIdRequest) Execute() (*ProcessorPolicyToGeneratorMapping, *http.Response, error) {
	return r.ApiService.GetTokenGeneratorMappingByIdExecute(r)
}

/*
GetTokenGeneratorMappingById Get a Token Exchange Processor policy to Token Generator Mapping.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ID of Token Exchange Processor policy to Token Generator Mapping to fetch.
	@return ApiGetTokenGeneratorMappingByIdRequest
*/
func (a *OauthTokenExchangeTokenGeneratorMappingsAPIService) GetTokenGeneratorMappingById(ctx context.Context, id string) ApiGetTokenGeneratorMappingByIdRequest {
	return ApiGetTokenGeneratorMappingByIdRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ProcessorPolicyToGeneratorMapping
func (a *OauthTokenExchangeTokenGeneratorMappingsAPIService) GetTokenGeneratorMappingByIdExecute(r ApiGetTokenGeneratorMappingByIdRequest) (*ProcessorPolicyToGeneratorMapping, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *ProcessorPolicyToGeneratorMapping
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalGetTokenGeneratorMappingByIdExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *OauthTokenExchangeTokenGeneratorMappingsAPIService) internalGetTokenGeneratorMappingByIdExecute(r ApiGetTokenGeneratorMappingByIdRequest) (*ProcessorPolicyToGeneratorMapping, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProcessorPolicyToGeneratorMapping
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthTokenExchangeTokenGeneratorMappingsAPIService.GetTokenGeneratorMappingById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/tokenExchange/tokenGeneratorMappings/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiResult
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

type ApiGetTokenGeneratorMappingsRequest struct {
	ctx        context.Context
	ApiService *OauthTokenExchangeTokenGeneratorMappingsAPIService
}

func (r ApiGetTokenGeneratorMappingsRequest) Execute() (*ProcessorPolicyToGeneratorMappings, *http.Response, error) {
	return r.ApiService.GetTokenGeneratorMappingsExecute(r)
}

/*
GetTokenGeneratorMappings Get the list of Token Exchange Processor policy to Token Generator Mappings.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetTokenGeneratorMappingsRequest
*/
func (a *OauthTokenExchangeTokenGeneratorMappingsAPIService) GetTokenGeneratorMappings(ctx context.Context) ApiGetTokenGeneratorMappingsRequest {
	return ApiGetTokenGeneratorMappingsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ProcessorPolicyToGeneratorMappings
func (a *OauthTokenExchangeTokenGeneratorMappingsAPIService) GetTokenGeneratorMappingsExecute(r ApiGetTokenGeneratorMappingsRequest) (*ProcessorPolicyToGeneratorMappings, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *ProcessorPolicyToGeneratorMappings
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalGetTokenGeneratorMappingsExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *OauthTokenExchangeTokenGeneratorMappingsAPIService) internalGetTokenGeneratorMappingsExecute(r ApiGetTokenGeneratorMappingsRequest) (*ProcessorPolicyToGeneratorMappings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProcessorPolicyToGeneratorMappings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthTokenExchangeTokenGeneratorMappingsAPIService.GetTokenGeneratorMappings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/tokenExchange/tokenGeneratorMappings"

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

type ApiUpdateTokenGeneratorMappingByIdRequest struct {
	ctx                       context.Context
	ApiService                *OauthTokenExchangeTokenGeneratorMappingsAPIService
	id                        string
	body                      *ProcessorPolicyToGeneratorMapping
	xBypassExternalValidation *bool
}

// Configuration for updated Token Exchange Processor policy to Token Generator Mapping.
func (r ApiUpdateTokenGeneratorMappingByIdRequest) Body(body ProcessorPolicyToGeneratorMapping) ApiUpdateTokenGeneratorMappingByIdRequest {
	r.body = &body
	return r
}

// External validation will be bypassed when set to true. Default to false.
func (r ApiUpdateTokenGeneratorMappingByIdRequest) XBypassExternalValidation(xBypassExternalValidation bool) ApiUpdateTokenGeneratorMappingByIdRequest {
	r.xBypassExternalValidation = &xBypassExternalValidation
	return r
}

func (r ApiUpdateTokenGeneratorMappingByIdRequest) Execute() (*ProcessorPolicyToGeneratorMapping, *http.Response, error) {
	return r.ApiService.UpdateTokenGeneratorMappingByIdExecute(r)
}

/*
UpdateTokenGeneratorMappingById Update a Token Exchange Processor policy to Token Generator Mapping.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ID of Token Exchange Processor policy to Token Generator Mapping to update.
	@return ApiUpdateTokenGeneratorMappingByIdRequest
*/
func (a *OauthTokenExchangeTokenGeneratorMappingsAPIService) UpdateTokenGeneratorMappingById(ctx context.Context, id string) ApiUpdateTokenGeneratorMappingByIdRequest {
	return ApiUpdateTokenGeneratorMappingByIdRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ProcessorPolicyToGeneratorMapping
func (a *OauthTokenExchangeTokenGeneratorMappingsAPIService) UpdateTokenGeneratorMappingByIdExecute(r ApiUpdateTokenGeneratorMappingByIdRequest) (*ProcessorPolicyToGeneratorMapping, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *ProcessorPolicyToGeneratorMapping
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalUpdateTokenGeneratorMappingByIdExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *OauthTokenExchangeTokenGeneratorMappingsAPIService) internalUpdateTokenGeneratorMappingByIdExecute(r ApiUpdateTokenGeneratorMappingByIdRequest) (*ProcessorPolicyToGeneratorMapping, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProcessorPolicyToGeneratorMapping
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthTokenExchangeTokenGeneratorMappingsAPIService.UpdateTokenGeneratorMappingById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/tokenExchange/tokenGeneratorMappings/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	if r.xBypassExternalValidation != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-BypassExternalValidation", r.xBypassExternalValidation, "")
	}
	// body params
	localVarPostBody = r.body
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiResult
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ApiResult
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
