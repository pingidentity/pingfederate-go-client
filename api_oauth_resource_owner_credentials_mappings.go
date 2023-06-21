/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 0.0.1
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

// OauthResourceOwnerCredentialsMappingsApiService OauthResourceOwnerCredentialsMappingsApi service
type OauthResourceOwnerCredentialsMappingsApiService service

type ApiCreateResourceOwnerCredentialsMappingRequest struct {
	ctx                       context.Context
	ApiService                *OauthResourceOwnerCredentialsMappingsApiService
	body                      *ResourceOwnerCredentialsMapping
	xBypassExternalValidation *bool
}

// Configuration for Resource Owner Credentials mapping.
func (r ApiCreateResourceOwnerCredentialsMappingRequest) Body(body ResourceOwnerCredentialsMapping) ApiCreateResourceOwnerCredentialsMappingRequest {
	r.body = &body
	return r
}

// External validation will be bypassed when set to true. Default to false.
func (r ApiCreateResourceOwnerCredentialsMappingRequest) XBypassExternalValidation(xBypassExternalValidation bool) ApiCreateResourceOwnerCredentialsMappingRequest {
	r.xBypassExternalValidation = &xBypassExternalValidation
	return r
}

func (r ApiCreateResourceOwnerCredentialsMappingRequest) Execute() (*ResourceOwnerCredentialsMapping, *http.Response, error) {
	return r.ApiService.CreateResourceOwnerCredentialsMappingExecute(r)
}

/*
CreateResourceOwnerCredentialsMapping Create a new Resource Owner Credentials mapping.

Create a new Resource Owner Credentials mapping. If a Resource Owner Credentials mapping can't be created, a 422 status code is returned along with a list of validation errors that must be corrected.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateResourceOwnerCredentialsMappingRequest
*/
func (a *OauthResourceOwnerCredentialsMappingsApiService) CreateResourceOwnerCredentialsMapping(ctx context.Context) ApiCreateResourceOwnerCredentialsMappingRequest {
	return ApiCreateResourceOwnerCredentialsMappingRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ResourceOwnerCredentialsMapping
func (a *OauthResourceOwnerCredentialsMappingsApiService) CreateResourceOwnerCredentialsMappingExecute(r ApiCreateResourceOwnerCredentialsMappingRequest) (*ResourceOwnerCredentialsMapping, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ResourceOwnerCredentialsMapping
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthResourceOwnerCredentialsMappingsApiService.CreateResourceOwnerCredentialsMapping")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/resourceOwnerCredentialsMappings"

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

type ApiDeleteResourceOwnerCredentialsMappingRequest struct {
	ctx        context.Context
	ApiService *OauthResourceOwnerCredentialsMappingsApiService
	id         string
}

func (r ApiDeleteResourceOwnerCredentialsMappingRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteResourceOwnerCredentialsMappingExecute(r)
}

/*
DeleteResourceOwnerCredentialsMapping Delete a Resource Owner Credentials mapping.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ID of the Resource Owner Credentials mapping.
	@return ApiDeleteResourceOwnerCredentialsMappingRequest
*/
func (a *OauthResourceOwnerCredentialsMappingsApiService) DeleteResourceOwnerCredentialsMapping(ctx context.Context, id string) ApiDeleteResourceOwnerCredentialsMappingRequest {
	return ApiDeleteResourceOwnerCredentialsMappingRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *OauthResourceOwnerCredentialsMappingsApiService) DeleteResourceOwnerCredentialsMappingExecute(r ApiDeleteResourceOwnerCredentialsMappingRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthResourceOwnerCredentialsMappingsApiService.DeleteResourceOwnerCredentialsMapping")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/resourceOwnerCredentialsMappings/{id}"
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

type ApiGetResourceOwnerCredentialsMappingRequest struct {
	ctx        context.Context
	ApiService *OauthResourceOwnerCredentialsMappingsApiService
	id         string
}

func (r ApiGetResourceOwnerCredentialsMappingRequest) Execute() (*ResourceOwnerCredentialsMapping, *http.Response, error) {
	return r.ApiService.GetResourceOwnerCredentialsMappingExecute(r)
}

/*
GetResourceOwnerCredentialsMapping Find the Resource Owner Credentials mapping by the ID.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ID of the Resource Owner Credentials mapping.
	@return ApiGetResourceOwnerCredentialsMappingRequest
*/
func (a *OauthResourceOwnerCredentialsMappingsApiService) GetResourceOwnerCredentialsMapping(ctx context.Context, id string) ApiGetResourceOwnerCredentialsMappingRequest {
	return ApiGetResourceOwnerCredentialsMappingRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ResourceOwnerCredentialsMapping
func (a *OauthResourceOwnerCredentialsMappingsApiService) GetResourceOwnerCredentialsMappingExecute(r ApiGetResourceOwnerCredentialsMappingRequest) (*ResourceOwnerCredentialsMapping, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ResourceOwnerCredentialsMapping
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthResourceOwnerCredentialsMappingsApiService.GetResourceOwnerCredentialsMapping")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/resourceOwnerCredentialsMappings/{id}"
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

type ApiGetResourceOwnerCredentialsMappingsRequest struct {
	ctx        context.Context
	ApiService *OauthResourceOwnerCredentialsMappingsApiService
}

func (r ApiGetResourceOwnerCredentialsMappingsRequest) Execute() (*ResourceOwnerCredentialsMappings, *http.Response, error) {
	return r.ApiService.GetResourceOwnerCredentialsMappingsExecute(r)
}

/*
GetResourceOwnerCredentialsMappings Get the list of Resource Owner Credentials Grant Mapping.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetResourceOwnerCredentialsMappingsRequest
*/
func (a *OauthResourceOwnerCredentialsMappingsApiService) GetResourceOwnerCredentialsMappings(ctx context.Context) ApiGetResourceOwnerCredentialsMappingsRequest {
	return ApiGetResourceOwnerCredentialsMappingsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ResourceOwnerCredentialsMappings
func (a *OauthResourceOwnerCredentialsMappingsApiService) GetResourceOwnerCredentialsMappingsExecute(r ApiGetResourceOwnerCredentialsMappingsRequest) (*ResourceOwnerCredentialsMappings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ResourceOwnerCredentialsMappings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthResourceOwnerCredentialsMappingsApiService.GetResourceOwnerCredentialsMappings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/resourceOwnerCredentialsMappings"

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

type ApiUpdateResourceOwnerCredentialsMappingRequest struct {
	ctx                       context.Context
	ApiService                *OauthResourceOwnerCredentialsMappingsApiService
	id                        string
	body                      *ResourceOwnerCredentialsMapping
	xBypassExternalValidation *bool
}

// Configuration for Resource Owner Credentials mapping.
func (r ApiUpdateResourceOwnerCredentialsMappingRequest) Body(body ResourceOwnerCredentialsMapping) ApiUpdateResourceOwnerCredentialsMappingRequest {
	r.body = &body
	return r
}

// External validation will be bypassed when set to true. Default to false.
func (r ApiUpdateResourceOwnerCredentialsMappingRequest) XBypassExternalValidation(xBypassExternalValidation bool) ApiUpdateResourceOwnerCredentialsMappingRequest {
	r.xBypassExternalValidation = &xBypassExternalValidation
	return r
}

func (r ApiUpdateResourceOwnerCredentialsMappingRequest) Execute() (*ResourceOwnerCredentialsMapping, *http.Response, error) {
	return r.ApiService.UpdateResourceOwnerCredentialsMappingExecute(r)
}

/*
UpdateResourceOwnerCredentialsMapping Update a Resource Owner Credentials mapping.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ID of the Resource Owner Credentials mapping to update.
	@return ApiUpdateResourceOwnerCredentialsMappingRequest
*/
func (a *OauthResourceOwnerCredentialsMappingsApiService) UpdateResourceOwnerCredentialsMapping(ctx context.Context, id string) ApiUpdateResourceOwnerCredentialsMappingRequest {
	return ApiUpdateResourceOwnerCredentialsMappingRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ResourceOwnerCredentialsMapping
func (a *OauthResourceOwnerCredentialsMappingsApiService) UpdateResourceOwnerCredentialsMappingExecute(r ApiUpdateResourceOwnerCredentialsMappingRequest) (*ResourceOwnerCredentialsMapping, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ResourceOwnerCredentialsMapping
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthResourceOwnerCredentialsMappingsApiService.UpdateResourceOwnerCredentialsMapping")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/resourceOwnerCredentialsMappings/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
