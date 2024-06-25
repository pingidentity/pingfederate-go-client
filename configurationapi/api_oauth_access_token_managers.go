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

// OauthAccessTokenManagersAPIService OauthAccessTokenManagersAPI service
type OauthAccessTokenManagersAPIService service

type ApiCreateTokenManagerRequest struct {
	ctx        context.Context
	ApiService *OauthAccessTokenManagersAPIService
	body       *AccessTokenManager
}

// Configuration for plugin instance.
func (r ApiCreateTokenManagerRequest) Body(body AccessTokenManager) ApiCreateTokenManagerRequest {
	r.body = &body
	return r
}

func (r ApiCreateTokenManagerRequest) Execute() (*AccessTokenManager, *http.Response, error) {
	return r.ApiService.CreateTokenManagerExecute(r)
}

/*
CreateTokenManager Create a token management plugin instance.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateTokenManagerRequest
*/
func (a *OauthAccessTokenManagersAPIService) CreateTokenManager(ctx context.Context) ApiCreateTokenManagerRequest {
	return ApiCreateTokenManagerRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AccessTokenManager
func (a *OauthAccessTokenManagersAPIService) CreateTokenManagerExecute(r ApiCreateTokenManagerRequest) (*AccessTokenManager, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *AccessTokenManager
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalCreateTokenManagerExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *OauthAccessTokenManagersAPIService) internalCreateTokenManagerExecute(r ApiCreateTokenManagerRequest) (*AccessTokenManager, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AccessTokenManager
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthAccessTokenManagersAPIService.CreateTokenManager")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/accessTokenManagers"

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

type ApiDeleteTokenManagerRequest struct {
	ctx        context.Context
	ApiService *OauthAccessTokenManagersAPIService
	id         string
}

func (r ApiDeleteTokenManagerRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteTokenManagerExecute(r)
}

/*
DeleteTokenManager Delete a token management plugin instance.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ID of token management plugin instance.
	@return ApiDeleteTokenManagerRequest
*/
func (a *OauthAccessTokenManagersAPIService) DeleteTokenManager(ctx context.Context, id string) ApiDeleteTokenManagerRequest {
	return ApiDeleteTokenManagerRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *OauthAccessTokenManagersAPIService) DeleteTokenManagerExecute(r ApiDeleteTokenManagerRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthAccessTokenManagersAPIService.DeleteTokenManager")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/accessTokenManagers/{id}"
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

type ApiGetOauthAccessTokenManagersSettingsRequest struct {
	ctx        context.Context
	ApiService *OauthAccessTokenManagersAPIService
}

func (r ApiGetOauthAccessTokenManagersSettingsRequest) Execute() (*AccessTokenManagementSettings, *http.Response, error) {
	return r.ApiService.GetOauthAccessTokenManagersSettingsExecute(r)
}

/*
GetOauthAccessTokenManagersSettings Get general access token management settings.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetOauthAccessTokenManagersSettingsRequest
*/
func (a *OauthAccessTokenManagersAPIService) GetOauthAccessTokenManagersSettings(ctx context.Context) ApiGetOauthAccessTokenManagersSettingsRequest {
	return ApiGetOauthAccessTokenManagersSettingsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AccessTokenManagementSettings
func (a *OauthAccessTokenManagersAPIService) GetOauthAccessTokenManagersSettingsExecute(r ApiGetOauthAccessTokenManagersSettingsRequest) (*AccessTokenManagementSettings, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *AccessTokenManagementSettings
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalGetOauthAccessTokenManagersSettingsExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *OauthAccessTokenManagersAPIService) internalGetOauthAccessTokenManagersSettingsExecute(r ApiGetOauthAccessTokenManagersSettingsRequest) (*AccessTokenManagementSettings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AccessTokenManagementSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthAccessTokenManagersAPIService.GetOauthAccessTokenManagersSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/accessTokenManagers/settings"

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

type ApiGetTokenManagerRequest struct {
	ctx        context.Context
	ApiService *OauthAccessTokenManagersAPIService
	id         string
}

func (r ApiGetTokenManagerRequest) Execute() (*AccessTokenManager, *http.Response, error) {
	return r.ApiService.GetTokenManagerExecute(r)
}

/*
GetTokenManager Get a specific token management plugin instance.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ID of token management plugin instance.
	@return ApiGetTokenManagerRequest
*/
func (a *OauthAccessTokenManagersAPIService) GetTokenManager(ctx context.Context, id string) ApiGetTokenManagerRequest {
	return ApiGetTokenManagerRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return AccessTokenManager
func (a *OauthAccessTokenManagersAPIService) GetTokenManagerExecute(r ApiGetTokenManagerRequest) (*AccessTokenManager, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *AccessTokenManager
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalGetTokenManagerExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *OauthAccessTokenManagersAPIService) internalGetTokenManagerExecute(r ApiGetTokenManagerRequest) (*AccessTokenManager, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AccessTokenManager
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthAccessTokenManagersAPIService.GetTokenManager")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/accessTokenManagers/{id}"
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

type ApiGetTokenManagerDescriptorRequest struct {
	ctx        context.Context
	ApiService *OauthAccessTokenManagersAPIService
	id         string
}

func (r ApiGetTokenManagerDescriptorRequest) Execute() (*AccessTokenManagerDescriptor, *http.Response, error) {
	return r.ApiService.GetTokenManagerDescriptorExecute(r)
}

/*
GetTokenManagerDescriptor Get the description of a token management plugin descriptor.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ID of token management plugin descriptor.
	@return ApiGetTokenManagerDescriptorRequest
*/
func (a *OauthAccessTokenManagersAPIService) GetTokenManagerDescriptor(ctx context.Context, id string) ApiGetTokenManagerDescriptorRequest {
	return ApiGetTokenManagerDescriptorRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return AccessTokenManagerDescriptor
func (a *OauthAccessTokenManagersAPIService) GetTokenManagerDescriptorExecute(r ApiGetTokenManagerDescriptorRequest) (*AccessTokenManagerDescriptor, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *AccessTokenManagerDescriptor
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalGetTokenManagerDescriptorExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *OauthAccessTokenManagersAPIService) internalGetTokenManagerDescriptorExecute(r ApiGetTokenManagerDescriptorRequest) (*AccessTokenManagerDescriptor, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AccessTokenManagerDescriptor
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthAccessTokenManagersAPIService.GetTokenManagerDescriptor")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/accessTokenManagers/descriptors/{id}"
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

type ApiGetTokenManagerDescriptorsRequest struct {
	ctx        context.Context
	ApiService *OauthAccessTokenManagersAPIService
}

func (r ApiGetTokenManagerDescriptorsRequest) Execute() (*AccessTokenManagerDescriptors, *http.Response, error) {
	return r.ApiService.GetTokenManagerDescriptorsExecute(r)
}

/*
GetTokenManagerDescriptors Get the list of available token management plugin descriptors.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetTokenManagerDescriptorsRequest
*/
func (a *OauthAccessTokenManagersAPIService) GetTokenManagerDescriptors(ctx context.Context) ApiGetTokenManagerDescriptorsRequest {
	return ApiGetTokenManagerDescriptorsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AccessTokenManagerDescriptors
func (a *OauthAccessTokenManagersAPIService) GetTokenManagerDescriptorsExecute(r ApiGetTokenManagerDescriptorsRequest) (*AccessTokenManagerDescriptors, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *AccessTokenManagerDescriptors
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalGetTokenManagerDescriptorsExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *OauthAccessTokenManagersAPIService) internalGetTokenManagerDescriptorsExecute(r ApiGetTokenManagerDescriptorsRequest) (*AccessTokenManagerDescriptors, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AccessTokenManagerDescriptors
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthAccessTokenManagersAPIService.GetTokenManagerDescriptors")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/accessTokenManagers/descriptors"

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

type ApiGetTokenManagersRequest struct {
	ctx        context.Context
	ApiService *OauthAccessTokenManagersAPIService
}

func (r ApiGetTokenManagersRequest) Execute() (*AccessTokenManagers, *http.Response, error) {
	return r.ApiService.GetTokenManagersExecute(r)
}

/*
GetTokenManagers Get a list of all token management plugin instances.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetTokenManagersRequest
*/
func (a *OauthAccessTokenManagersAPIService) GetTokenManagers(ctx context.Context) ApiGetTokenManagersRequest {
	return ApiGetTokenManagersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AccessTokenManagers
func (a *OauthAccessTokenManagersAPIService) GetTokenManagersExecute(r ApiGetTokenManagersRequest) (*AccessTokenManagers, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *AccessTokenManagers
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalGetTokenManagersExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *OauthAccessTokenManagersAPIService) internalGetTokenManagersExecute(r ApiGetTokenManagersRequest) (*AccessTokenManagers, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AccessTokenManagers
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthAccessTokenManagersAPIService.GetTokenManagers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/accessTokenManagers"

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

type ApiUpdateOauthAccessTokenManagersSettingsRequest struct {
	ctx        context.Context
	ApiService *OauthAccessTokenManagersAPIService
	body       *AccessTokenManagementSettings
}

// Access token management settings.
func (r ApiUpdateOauthAccessTokenManagersSettingsRequest) Body(body AccessTokenManagementSettings) ApiUpdateOauthAccessTokenManagersSettingsRequest {
	r.body = &body
	return r
}

func (r ApiUpdateOauthAccessTokenManagersSettingsRequest) Execute() (*AccessTokenManagementSettings, *http.Response, error) {
	return r.ApiService.UpdateOauthAccessTokenManagersSettingsExecute(r)
}

/*
UpdateOauthAccessTokenManagersSettings Update general access token management settings.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiUpdateOauthAccessTokenManagersSettingsRequest
*/
func (a *OauthAccessTokenManagersAPIService) UpdateOauthAccessTokenManagersSettings(ctx context.Context) ApiUpdateOauthAccessTokenManagersSettingsRequest {
	return ApiUpdateOauthAccessTokenManagersSettingsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AccessTokenManagementSettings
func (a *OauthAccessTokenManagersAPIService) UpdateOauthAccessTokenManagersSettingsExecute(r ApiUpdateOauthAccessTokenManagersSettingsRequest) (*AccessTokenManagementSettings, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *AccessTokenManagementSettings
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalUpdateOauthAccessTokenManagersSettingsExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *OauthAccessTokenManagersAPIService) internalUpdateOauthAccessTokenManagersSettingsExecute(r ApiUpdateOauthAccessTokenManagersSettingsRequest) (*AccessTokenManagementSettings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AccessTokenManagementSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthAccessTokenManagersAPIService.UpdateOauthAccessTokenManagersSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/accessTokenManagers/settings"

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

type ApiUpdateTokenManagerRequest struct {
	ctx        context.Context
	ApiService *OauthAccessTokenManagersAPIService
	id         string
	body       *AccessTokenManager
}

// Configuration for token management plugin instance.
func (r ApiUpdateTokenManagerRequest) Body(body AccessTokenManager) ApiUpdateTokenManagerRequest {
	r.body = &body
	return r
}

func (r ApiUpdateTokenManagerRequest) Execute() (*AccessTokenManager, *http.Response, error) {
	return r.ApiService.UpdateTokenManagerExecute(r)
}

/*
UpdateTokenManager Update a token management plugin instance.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ID of token management plugin instance.
	@return ApiUpdateTokenManagerRequest
*/
func (a *OauthAccessTokenManagersAPIService) UpdateTokenManager(ctx context.Context, id string) ApiUpdateTokenManagerRequest {
	return ApiUpdateTokenManagerRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return AccessTokenManager
func (a *OauthAccessTokenManagersAPIService) UpdateTokenManagerExecute(r ApiUpdateTokenManagerRequest) (*AccessTokenManager, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *AccessTokenManager
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalUpdateTokenManagerExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *OauthAccessTokenManagersAPIService) internalUpdateTokenManagerExecute(r ApiUpdateTokenManagerRequest) (*AccessTokenManager, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AccessTokenManager
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthAccessTokenManagersAPIService.UpdateTokenManager")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/accessTokenManagers/{id}"
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
