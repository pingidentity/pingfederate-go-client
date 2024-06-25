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

// TokenProcessorToTokenGeneratorMappingsAPIService TokenProcessorToTokenGeneratorMappingsAPI service
type TokenProcessorToTokenGeneratorMappingsAPIService service

type ApiCreateTokenToTokenMappingRequest struct {
	ctx                       context.Context
	ApiService                *TokenProcessorToTokenGeneratorMappingsAPIService
	body                      *TokenToTokenMapping
	xBypassExternalValidation *bool
}

// Configuration for a new Token Processor to Token Generator Mapping.
func (r ApiCreateTokenToTokenMappingRequest) Body(body TokenToTokenMapping) ApiCreateTokenToTokenMappingRequest {
	r.body = &body
	return r
}

// External validation will be bypassed when set to true. Default to false.
func (r ApiCreateTokenToTokenMappingRequest) XBypassExternalValidation(xBypassExternalValidation bool) ApiCreateTokenToTokenMappingRequest {
	r.xBypassExternalValidation = &xBypassExternalValidation
	return r
}

func (r ApiCreateTokenToTokenMappingRequest) Execute() (*TokenToTokenMapping, *http.Response, error) {
	return r.ApiService.CreateTokenToTokenMappingExecute(r)
}

/*
CreateTokenToTokenMapping Create a new Token Processor to Token Generator Mapping.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateTokenToTokenMappingRequest
*/
func (a *TokenProcessorToTokenGeneratorMappingsAPIService) CreateTokenToTokenMapping(ctx context.Context) ApiCreateTokenToTokenMappingRequest {
	return ApiCreateTokenToTokenMappingRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TokenToTokenMapping
func (a *TokenProcessorToTokenGeneratorMappingsAPIService) CreateTokenToTokenMappingExecute(r ApiCreateTokenToTokenMappingRequest) (*TokenToTokenMapping, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *TokenToTokenMapping
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalCreateTokenToTokenMappingExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *TokenProcessorToTokenGeneratorMappingsAPIService) internalCreateTokenToTokenMappingExecute(r ApiCreateTokenToTokenMappingRequest) (*TokenToTokenMapping, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TokenToTokenMapping
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TokenProcessorToTokenGeneratorMappingsAPIService.CreateTokenToTokenMapping")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tokenProcessorToTokenGeneratorMappings"

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

type ApiDeleteTokenToTokenMappingByIdRequest struct {
	ctx        context.Context
	ApiService *TokenProcessorToTokenGeneratorMappingsAPIService
	id         string
}

func (r ApiDeleteTokenToTokenMappingByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteTokenToTokenMappingByIdExecute(r)
}

/*
DeleteTokenToTokenMappingById Delete a Token Processor to Token Generator Mapping.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ID of Token Processor to Token Generator Mapping to delete.
	@return ApiDeleteTokenToTokenMappingByIdRequest
*/
func (a *TokenProcessorToTokenGeneratorMappingsAPIService) DeleteTokenToTokenMappingById(ctx context.Context, id string) ApiDeleteTokenToTokenMappingByIdRequest {
	return ApiDeleteTokenToTokenMappingByIdRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *TokenProcessorToTokenGeneratorMappingsAPIService) DeleteTokenToTokenMappingByIdExecute(r ApiDeleteTokenToTokenMappingByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TokenProcessorToTokenGeneratorMappingsAPIService.DeleteTokenToTokenMappingById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tokenProcessorToTokenGeneratorMappings/{id}"
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

type ApiGetTokenToTokenMappingByIdRequest struct {
	ctx        context.Context
	ApiService *TokenProcessorToTokenGeneratorMappingsAPIService
	id         string
}

func (r ApiGetTokenToTokenMappingByIdRequest) Execute() (*TokenToTokenMapping, *http.Response, error) {
	return r.ApiService.GetTokenToTokenMappingByIdExecute(r)
}

/*
GetTokenToTokenMappingById Get a Token Processor to Token Generator Mapping.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ID of Token Processor to Token Generator Mapping to fetch.
	@return ApiGetTokenToTokenMappingByIdRequest
*/
func (a *TokenProcessorToTokenGeneratorMappingsAPIService) GetTokenToTokenMappingById(ctx context.Context, id string) ApiGetTokenToTokenMappingByIdRequest {
	return ApiGetTokenToTokenMappingByIdRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return TokenToTokenMapping
func (a *TokenProcessorToTokenGeneratorMappingsAPIService) GetTokenToTokenMappingByIdExecute(r ApiGetTokenToTokenMappingByIdRequest) (*TokenToTokenMapping, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *TokenToTokenMapping
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalGetTokenToTokenMappingByIdExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *TokenProcessorToTokenGeneratorMappingsAPIService) internalGetTokenToTokenMappingByIdExecute(r ApiGetTokenToTokenMappingByIdRequest) (*TokenToTokenMapping, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TokenToTokenMapping
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TokenProcessorToTokenGeneratorMappingsAPIService.GetTokenToTokenMappingById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tokenProcessorToTokenGeneratorMappings/{id}"
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

type ApiGetTokenToTokenMappingsRequest struct {
	ctx        context.Context
	ApiService *TokenProcessorToTokenGeneratorMappingsAPIService
}

func (r ApiGetTokenToTokenMappingsRequest) Execute() (*TokenToTokenMappings, *http.Response, error) {
	return r.ApiService.GetTokenToTokenMappingsExecute(r)
}

/*
GetTokenToTokenMappings Get the list of Token Processor to Token Generator Mappings.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetTokenToTokenMappingsRequest
*/
func (a *TokenProcessorToTokenGeneratorMappingsAPIService) GetTokenToTokenMappings(ctx context.Context) ApiGetTokenToTokenMappingsRequest {
	return ApiGetTokenToTokenMappingsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TokenToTokenMappings
func (a *TokenProcessorToTokenGeneratorMappingsAPIService) GetTokenToTokenMappingsExecute(r ApiGetTokenToTokenMappingsRequest) (*TokenToTokenMappings, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *TokenToTokenMappings
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalGetTokenToTokenMappingsExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *TokenProcessorToTokenGeneratorMappingsAPIService) internalGetTokenToTokenMappingsExecute(r ApiGetTokenToTokenMappingsRequest) (*TokenToTokenMappings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TokenToTokenMappings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TokenProcessorToTokenGeneratorMappingsAPIService.GetTokenToTokenMappings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tokenProcessorToTokenGeneratorMappings"

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

type ApiUpdateTokenToTokenMappingByIdRequest struct {
	ctx                       context.Context
	ApiService                *TokenProcessorToTokenGeneratorMappingsAPIService
	id                        string
	body                      *TokenToTokenMapping
	xBypassExternalValidation *bool
}

// Configuration for updated Token Processor to Token Generator Mapping.
func (r ApiUpdateTokenToTokenMappingByIdRequest) Body(body TokenToTokenMapping) ApiUpdateTokenToTokenMappingByIdRequest {
	r.body = &body
	return r
}

// External validation will be bypassed when set to true. Default to false.
func (r ApiUpdateTokenToTokenMappingByIdRequest) XBypassExternalValidation(xBypassExternalValidation bool) ApiUpdateTokenToTokenMappingByIdRequest {
	r.xBypassExternalValidation = &xBypassExternalValidation
	return r
}

func (r ApiUpdateTokenToTokenMappingByIdRequest) Execute() (*TokenToTokenMapping, *http.Response, error) {
	return r.ApiService.UpdateTokenToTokenMappingByIdExecute(r)
}

/*
UpdateTokenToTokenMappingById Update a Token Processor to Token Generator Mapping.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ID of Token Processor to Token Generator Mapping to update.
	@return ApiUpdateTokenToTokenMappingByIdRequest
*/
func (a *TokenProcessorToTokenGeneratorMappingsAPIService) UpdateTokenToTokenMappingById(ctx context.Context, id string) ApiUpdateTokenToTokenMappingByIdRequest {
	return ApiUpdateTokenToTokenMappingByIdRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return TokenToTokenMapping
func (a *TokenProcessorToTokenGeneratorMappingsAPIService) UpdateTokenToTokenMappingByIdExecute(r ApiUpdateTokenToTokenMappingByIdRequest) (*TokenToTokenMapping, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *TokenToTokenMapping
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalUpdateTokenToTokenMappingByIdExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *TokenProcessorToTokenGeneratorMappingsAPIService) internalUpdateTokenToTokenMappingByIdExecute(r ApiUpdateTokenToTokenMappingByIdRequest) (*TokenToTokenMapping, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TokenToTokenMapping
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TokenProcessorToTokenGeneratorMappingsAPIService.UpdateTokenToTokenMappingById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tokenProcessorToTokenGeneratorMappings/{id}"
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
