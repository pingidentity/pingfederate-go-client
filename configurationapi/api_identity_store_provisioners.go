/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.0.0.9
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

// IdentityStoreProvisionersAPIService IdentityStoreProvisionersAPI service
type IdentityStoreProvisionersAPIService service

type ApiCreateIdentityStoreProvisionerRequest struct {
	ctx        context.Context
	ApiService *IdentityStoreProvisionersAPIService
	body       *IdentityStoreProvisioner
}

// Configuration for the identity store provisioner instance.
func (r ApiCreateIdentityStoreProvisionerRequest) Body(body IdentityStoreProvisioner) ApiCreateIdentityStoreProvisionerRequest {
	r.body = &body
	return r
}

func (r ApiCreateIdentityStoreProvisionerRequest) Execute() (*IdentityStoreProvisioner, *http.Response, error) {
	return r.ApiService.CreateIdentityStoreProvisionerExecute(r)
}

/*
CreateIdentityStoreProvisioner Create a new identity store provisioner instance.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateIdentityStoreProvisionerRequest
*/
func (a *IdentityStoreProvisionersAPIService) CreateIdentityStoreProvisioner(ctx context.Context) ApiCreateIdentityStoreProvisionerRequest {
	return ApiCreateIdentityStoreProvisionerRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return IdentityStoreProvisioner
func (a *IdentityStoreProvisionersAPIService) CreateIdentityStoreProvisionerExecute(r ApiCreateIdentityStoreProvisionerRequest) (*IdentityStoreProvisioner, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *IdentityStoreProvisioner
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalCreateIdentityStoreProvisionerExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *IdentityStoreProvisionersAPIService) internalCreateIdentityStoreProvisionerExecute(r ApiCreateIdentityStoreProvisionerRequest) (*IdentityStoreProvisioner, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *IdentityStoreProvisioner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityStoreProvisionersAPIService.CreateIdentityStoreProvisioner")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/identityStoreProvisioners"

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

type ApiDeleteIdentityStoreProvisionerRequest struct {
	ctx        context.Context
	ApiService *IdentityStoreProvisionersAPIService
	id         string
}

func (r ApiDeleteIdentityStoreProvisionerRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteIdentityStoreProvisionerExecute(r)
}

/*
DeleteIdentityStoreProvisioner Delete an identity store provisioner instance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ID of identity store provisioner instance
	@return ApiDeleteIdentityStoreProvisionerRequest
*/
func (a *IdentityStoreProvisionersAPIService) DeleteIdentityStoreProvisioner(ctx context.Context, id string) ApiDeleteIdentityStoreProvisionerRequest {
	return ApiDeleteIdentityStoreProvisionerRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *IdentityStoreProvisionersAPIService) DeleteIdentityStoreProvisionerExecute(r ApiDeleteIdentityStoreProvisionerRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityStoreProvisionersAPIService.DeleteIdentityStoreProvisioner")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/identityStoreProvisioners/{id}"
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
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
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

type ApiGetIdentityStoreProvisionerRequest struct {
	ctx        context.Context
	ApiService *IdentityStoreProvisionersAPIService
	id         string
}

func (r ApiGetIdentityStoreProvisionerRequest) Execute() (*IdentityStoreProvisioner, *http.Response, error) {
	return r.ApiService.GetIdentityStoreProvisionerExecute(r)
}

/*
GetIdentityStoreProvisioner Get an identity store provisioner by ID.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ID of identity store provisioner instance
	@return ApiGetIdentityStoreProvisionerRequest
*/
func (a *IdentityStoreProvisionersAPIService) GetIdentityStoreProvisioner(ctx context.Context, id string) ApiGetIdentityStoreProvisionerRequest {
	return ApiGetIdentityStoreProvisionerRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return IdentityStoreProvisioner
func (a *IdentityStoreProvisionersAPIService) GetIdentityStoreProvisionerExecute(r ApiGetIdentityStoreProvisionerRequest) (*IdentityStoreProvisioner, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *IdentityStoreProvisioner
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalGetIdentityStoreProvisionerExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *IdentityStoreProvisionersAPIService) internalGetIdentityStoreProvisionerExecute(r ApiGetIdentityStoreProvisionerRequest) (*IdentityStoreProvisioner, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *IdentityStoreProvisioner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityStoreProvisionersAPIService.GetIdentityStoreProvisioner")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/identityStoreProvisioners/{id}"
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

type ApiGetIdentityStoreProvisionerDescriptorByIdRequest struct {
	ctx        context.Context
	ApiService *IdentityStoreProvisionersAPIService
	id         string
}

func (r ApiGetIdentityStoreProvisionerDescriptorByIdRequest) Execute() (*IdentityStoreProvisionerDescriptor, *http.Response, error) {
	return r.ApiService.GetIdentityStoreProvisionerDescriptorByIdExecute(r)
}

/*
GetIdentityStoreProvisionerDescriptorById Get the descriptor of an identity store provisioner by ID.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ID of identity store provisioner descriptor
	@return ApiGetIdentityStoreProvisionerDescriptorByIdRequest
*/
func (a *IdentityStoreProvisionersAPIService) GetIdentityStoreProvisionerDescriptorById(ctx context.Context, id string) ApiGetIdentityStoreProvisionerDescriptorByIdRequest {
	return ApiGetIdentityStoreProvisionerDescriptorByIdRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return IdentityStoreProvisionerDescriptor
func (a *IdentityStoreProvisionersAPIService) GetIdentityStoreProvisionerDescriptorByIdExecute(r ApiGetIdentityStoreProvisionerDescriptorByIdRequest) (*IdentityStoreProvisionerDescriptor, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *IdentityStoreProvisionerDescriptor
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalGetIdentityStoreProvisionerDescriptorByIdExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *IdentityStoreProvisionersAPIService) internalGetIdentityStoreProvisionerDescriptorByIdExecute(r ApiGetIdentityStoreProvisionerDescriptorByIdRequest) (*IdentityStoreProvisionerDescriptor, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *IdentityStoreProvisionerDescriptor
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityStoreProvisionersAPIService.GetIdentityStoreProvisionerDescriptorById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/identityStoreProvisioners/descriptors/{id}"
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

type ApiGetIdentityStoreProvisionerDescriptorsRequest struct {
	ctx        context.Context
	ApiService *IdentityStoreProvisionersAPIService
}

func (r ApiGetIdentityStoreProvisionerDescriptorsRequest) Execute() (*IdentityStoreProvisionerDescriptors, *http.Response, error) {
	return r.ApiService.GetIdentityStoreProvisionerDescriptorsExecute(r)
}

/*
GetIdentityStoreProvisionerDescriptors Get the list of available identity store provisioner descriptors.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetIdentityStoreProvisionerDescriptorsRequest
*/
func (a *IdentityStoreProvisionersAPIService) GetIdentityStoreProvisionerDescriptors(ctx context.Context) ApiGetIdentityStoreProvisionerDescriptorsRequest {
	return ApiGetIdentityStoreProvisionerDescriptorsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return IdentityStoreProvisionerDescriptors
func (a *IdentityStoreProvisionersAPIService) GetIdentityStoreProvisionerDescriptorsExecute(r ApiGetIdentityStoreProvisionerDescriptorsRequest) (*IdentityStoreProvisionerDescriptors, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *IdentityStoreProvisionerDescriptors
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalGetIdentityStoreProvisionerDescriptorsExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *IdentityStoreProvisionersAPIService) internalGetIdentityStoreProvisionerDescriptorsExecute(r ApiGetIdentityStoreProvisionerDescriptorsRequest) (*IdentityStoreProvisionerDescriptors, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *IdentityStoreProvisionerDescriptors
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityStoreProvisionersAPIService.GetIdentityStoreProvisionerDescriptors")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/identityStoreProvisioners/descriptors"

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

type ApiGetIdentityStoreProvisionersRequest struct {
	ctx        context.Context
	ApiService *IdentityStoreProvisionersAPIService
}

func (r ApiGetIdentityStoreProvisionersRequest) Execute() (*IdentityStoreProvisioners, *http.Response, error) {
	return r.ApiService.GetIdentityStoreProvisionersExecute(r)
}

/*
GetIdentityStoreProvisioners Get the list of configured identity store provisioner instances.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetIdentityStoreProvisionersRequest
*/
func (a *IdentityStoreProvisionersAPIService) GetIdentityStoreProvisioners(ctx context.Context) ApiGetIdentityStoreProvisionersRequest {
	return ApiGetIdentityStoreProvisionersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return IdentityStoreProvisioners
func (a *IdentityStoreProvisionersAPIService) GetIdentityStoreProvisionersExecute(r ApiGetIdentityStoreProvisionersRequest) (*IdentityStoreProvisioners, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *IdentityStoreProvisioners
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalGetIdentityStoreProvisionersExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *IdentityStoreProvisionersAPIService) internalGetIdentityStoreProvisionersExecute(r ApiGetIdentityStoreProvisionersRequest) (*IdentityStoreProvisioners, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *IdentityStoreProvisioners
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityStoreProvisionersAPIService.GetIdentityStoreProvisioners")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/identityStoreProvisioners"

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

type ApiUpdateIdentityStoreProvisionerRequest struct {
	ctx        context.Context
	ApiService *IdentityStoreProvisionersAPIService
	id         string
	body       *IdentityStoreProvisioner
}

// Configuration for the identity store provisioner instance
func (r ApiUpdateIdentityStoreProvisionerRequest) Body(body IdentityStoreProvisioner) ApiUpdateIdentityStoreProvisionerRequest {
	r.body = &body
	return r
}

func (r ApiUpdateIdentityStoreProvisionerRequest) Execute() (*IdentityStoreProvisioner, *http.Response, error) {
	return r.ApiService.UpdateIdentityStoreProvisionerExecute(r)
}

/*
UpdateIdentityStoreProvisioner Update an identity store provisioner instance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ID of identity store provisioner instance
	@return ApiUpdateIdentityStoreProvisionerRequest
*/
func (a *IdentityStoreProvisionersAPIService) UpdateIdentityStoreProvisioner(ctx context.Context, id string) ApiUpdateIdentityStoreProvisionerRequest {
	return ApiUpdateIdentityStoreProvisionerRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return IdentityStoreProvisioner
func (a *IdentityStoreProvisionersAPIService) UpdateIdentityStoreProvisionerExecute(r ApiUpdateIdentityStoreProvisionerRequest) (*IdentityStoreProvisioner, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *IdentityStoreProvisioner
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalUpdateIdentityStoreProvisionerExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *IdentityStoreProvisionersAPIService) internalUpdateIdentityStoreProvisionerExecute(r ApiUpdateIdentityStoreProvisionerRequest) (*IdentityStoreProvisioner, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *IdentityStoreProvisioner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityStoreProvisionersAPIService.UpdateIdentityStoreProvisioner")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/identityStoreProvisioners/{id}"
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
