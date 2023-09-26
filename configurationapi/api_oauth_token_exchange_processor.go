/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 1.0.0
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

// OauthTokenExchangeProcessorApiService OauthTokenExchangeProcessorApi service
type OauthTokenExchangeProcessorApiService service

type ApiCreateOauthTokenExchangeProcessorPolicyRequest struct {
	ctx                      context.Context
	ApiService               *OauthTokenExchangeProcessorApiService
	body                     *TokenExchangeProcessorPolicy
	bypassExternalValidation *bool
}

// Configuration for new OAuth 2.0 Token Exchange Processor.
func (r ApiCreateOauthTokenExchangeProcessorPolicyRequest) Body(body TokenExchangeProcessorPolicy) ApiCreateOauthTokenExchangeProcessorPolicyRequest {
	r.body = &body
	return r
}

// External validation will be bypassed when set to true. Default to false.
func (r ApiCreateOauthTokenExchangeProcessorPolicyRequest) BypassExternalValidation(bypassExternalValidation bool) ApiCreateOauthTokenExchangeProcessorPolicyRequest {
	r.bypassExternalValidation = &bypassExternalValidation
	return r
}

func (r ApiCreateOauthTokenExchangeProcessorPolicyRequest) Execute() (*TokenExchangeProcessorPolicy, *http.Response, error) {
	return r.ApiService.CreateOauthTokenExchangeProcessorPolicyExecute(r)
}

/*
CreateOauthTokenExchangeProcessorPolicy Create a new OAuth 2.0 Token Exchange Processor policy.

Create a new OAuth 2.0 Token Exchange Processor policy. If the OAuth 2.0 Token Exchange Processor policy is not properly configured, a 422 status code is returned along with a list of validation errors that must be corrected.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateOauthTokenExchangeProcessorPolicyRequest
*/
func (a *OauthTokenExchangeProcessorApiService) CreateOauthTokenExchangeProcessorPolicy(ctx context.Context) ApiCreateOauthTokenExchangeProcessorPolicyRequest {
	return ApiCreateOauthTokenExchangeProcessorPolicyRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TokenExchangeProcessorPolicy
func (a *OauthTokenExchangeProcessorApiService) CreateOauthTokenExchangeProcessorPolicyExecute(r ApiCreateOauthTokenExchangeProcessorPolicyRequest) (*TokenExchangeProcessorPolicy, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TokenExchangeProcessorPolicy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthTokenExchangeProcessorApiService.CreateOauthTokenExchangeProcessorPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/tokenExchange/processor/policies"

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
	if r.bypassExternalValidation != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "bypassExternalValidation", r.bypassExternalValidation, "")
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

type ApiDeleteOauthTokenExchangeProcessorPolicyyRequest struct {
	ctx        context.Context
	ApiService *OauthTokenExchangeProcessorApiService
	id         string
}

func (r ApiDeleteOauthTokenExchangeProcessorPolicyyRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteOauthTokenExchangeProcessorPolicyyExecute(r)
}

/*
DeleteOauthTokenExchangeProcessorPolicyy Delete an OAuth 2.0 Token Exchange Processor policy.

Delete an OAuth 2.0 Token Exchange Processor policy with the specified ID. A 404 status code is returned for nonexistent IDs. Note: If the request succeeds, the response body is empty. If the request fails, an ApiResult is returned with details of the error.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ID of OAuth 2.0 Token Exchange Processor policy to delete.
	@return ApiDeleteOauthTokenExchangeProcessorPolicyyRequest
*/
func (a *OauthTokenExchangeProcessorApiService) DeleteOauthTokenExchangeProcessorPolicyy(ctx context.Context, id string) ApiDeleteOauthTokenExchangeProcessorPolicyyRequest {
	return ApiDeleteOauthTokenExchangeProcessorPolicyyRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *OauthTokenExchangeProcessorApiService) DeleteOauthTokenExchangeProcessorPolicyyExecute(r ApiDeleteOauthTokenExchangeProcessorPolicyyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthTokenExchangeProcessorApiService.DeleteOauthTokenExchangeProcessorPolicyy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/tokenExchange/processor/policies/{id}"
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

type ApiGetOauthTokenExchangeProcessorPolicyByIdRequest struct {
	ctx        context.Context
	ApiService *OauthTokenExchangeProcessorApiService
	id         string
}

func (r ApiGetOauthTokenExchangeProcessorPolicyByIdRequest) Execute() (*TokenExchangeProcessorPolicy, *http.Response, error) {
	return r.ApiService.GetOauthTokenExchangeProcessorPolicyByIdExecute(r)
}

/*
GetOauthTokenExchangeProcessorPolicyById Find an OAuth 2.0 Token Exchange Processor policy by ID.

Get an OAuth 2.0 Token Exchange Processor policy with the specified ID. A 404 status code is returned for nonexistent IDs.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ID of the OAuth 2.0 Token Exchange Processor policy to fetch.
	@return ApiGetOauthTokenExchangeProcessorPolicyByIdRequest
*/
func (a *OauthTokenExchangeProcessorApiService) GetOauthTokenExchangeProcessorPolicyById(ctx context.Context, id string) ApiGetOauthTokenExchangeProcessorPolicyByIdRequest {
	return ApiGetOauthTokenExchangeProcessorPolicyByIdRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return TokenExchangeProcessorPolicy
func (a *OauthTokenExchangeProcessorApiService) GetOauthTokenExchangeProcessorPolicyByIdExecute(r ApiGetOauthTokenExchangeProcessorPolicyByIdRequest) (*TokenExchangeProcessorPolicy, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TokenExchangeProcessorPolicy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthTokenExchangeProcessorApiService.GetOauthTokenExchangeProcessorPolicyById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/tokenExchange/processor/policies/{id}"
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

type ApiGetOauthTokenExchangeProcessorPolicyPoliciesRequest struct {
	ctx        context.Context
	ApiService *OauthTokenExchangeProcessorApiService
}

func (r ApiGetOauthTokenExchangeProcessorPolicyPoliciesRequest) Execute() (*TokenExchangeProcessorPolicies, *http.Response, error) {
	return r.ApiService.GetOauthTokenExchangeProcessorPolicyPoliciesExecute(r)
}

/*
GetOauthTokenExchangeProcessorPolicyPolicies Get list of OAuth 2.0 Token Exchange Processor policies.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetOauthTokenExchangeProcessorPolicyPoliciesRequest
*/
func (a *OauthTokenExchangeProcessorApiService) GetOauthTokenExchangeProcessorPolicyPolicies(ctx context.Context) ApiGetOauthTokenExchangeProcessorPolicyPoliciesRequest {
	return ApiGetOauthTokenExchangeProcessorPolicyPoliciesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TokenExchangeProcessorPolicies
func (a *OauthTokenExchangeProcessorApiService) GetOauthTokenExchangeProcessorPolicyPoliciesExecute(r ApiGetOauthTokenExchangeProcessorPolicyPoliciesRequest) (*TokenExchangeProcessorPolicies, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TokenExchangeProcessorPolicies
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthTokenExchangeProcessorApiService.GetOauthTokenExchangeProcessorPolicyPolicies")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/tokenExchange/processor/policies"

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

type ApiGetOauthTokenExchangeProcessorPolicySettingsRequest struct {
	ctx        context.Context
	ApiService *OauthTokenExchangeProcessorApiService
}

func (r ApiGetOauthTokenExchangeProcessorPolicySettingsRequest) Execute() (*TokenExchangeProcessorSettings, *http.Response, error) {
	return r.ApiService.GetOauthTokenExchangeProcessorPolicySettingsExecute(r)
}

/*
GetOauthTokenExchangeProcessorPolicySettings Get general OAuth 2.0 Token Exchange Processor settings.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetOauthTokenExchangeProcessorPolicySettingsRequest
*/
func (a *OauthTokenExchangeProcessorApiService) GetOauthTokenExchangeProcessorPolicySettings(ctx context.Context) ApiGetOauthTokenExchangeProcessorPolicySettingsRequest {
	return ApiGetOauthTokenExchangeProcessorPolicySettingsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TokenExchangeProcessorSettings
func (a *OauthTokenExchangeProcessorApiService) GetOauthTokenExchangeProcessorPolicySettingsExecute(r ApiGetOauthTokenExchangeProcessorPolicySettingsRequest) (*TokenExchangeProcessorSettings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TokenExchangeProcessorSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthTokenExchangeProcessorApiService.GetOauthTokenExchangeProcessorPolicySettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/tokenExchange/processor/settings"

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

type ApiUpdateOauthTokenExchangeProcessorPolicyRequest struct {
	ctx                      context.Context
	ApiService               *OauthTokenExchangeProcessorApiService
	id                       string
	body                     *TokenExchangeProcessorPolicy
	bypassExternalValidation *bool
}

// Configuration for updated OAuth 2.0 Token Exchange Processor policy.
func (r ApiUpdateOauthTokenExchangeProcessorPolicyRequest) Body(body TokenExchangeProcessorPolicy) ApiUpdateOauthTokenExchangeProcessorPolicyRequest {
	r.body = &body
	return r
}

// External validation will be bypassed when set to true. Default to false.
func (r ApiUpdateOauthTokenExchangeProcessorPolicyRequest) BypassExternalValidation(bypassExternalValidation bool) ApiUpdateOauthTokenExchangeProcessorPolicyRequest {
	r.bypassExternalValidation = &bypassExternalValidation
	return r
}

func (r ApiUpdateOauthTokenExchangeProcessorPolicyRequest) Execute() (*TokenExchangeProcessorPolicy, *http.Response, error) {
	return r.ApiService.UpdateOauthTokenExchangeProcessorPolicyExecute(r)
}

/*
UpdateOauthTokenExchangeProcessorPolicy Update an OAuth 2.0 Token Exchange Processor policy.

Update an OAuth 2.0 Token Exchange Processor policy with the matching ID. If the policy is not properly configured, a 422 status code is returned along with a list of validation errors that must be corrected. Note: A 404 status code is returned for nonexistent IDs.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ID of the OAuth 2.0 Token Exchange Processor policy to update.
	@return ApiUpdateOauthTokenExchangeProcessorPolicyRequest
*/
func (a *OauthTokenExchangeProcessorApiService) UpdateOauthTokenExchangeProcessorPolicy(ctx context.Context, id string) ApiUpdateOauthTokenExchangeProcessorPolicyRequest {
	return ApiUpdateOauthTokenExchangeProcessorPolicyRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return TokenExchangeProcessorPolicy
func (a *OauthTokenExchangeProcessorApiService) UpdateOauthTokenExchangeProcessorPolicyExecute(r ApiUpdateOauthTokenExchangeProcessorPolicyRequest) (*TokenExchangeProcessorPolicy, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TokenExchangeProcessorPolicy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthTokenExchangeProcessorApiService.UpdateOauthTokenExchangeProcessorPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/tokenExchange/processor/policies/{id}"
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
	if r.bypassExternalValidation != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "bypassExternalValidation", r.bypassExternalValidation, "")
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

type ApiUpdateOauthTokenExchangeProcessorPolicySettingsRequest struct {
	ctx                      context.Context
	ApiService               *OauthTokenExchangeProcessorApiService
	body                     *TokenExchangeProcessorSettings
	bypassExternalValidation *bool
}

// OAuth 2.0 Token Exchange Processor settings.
func (r ApiUpdateOauthTokenExchangeProcessorPolicySettingsRequest) Body(body TokenExchangeProcessorSettings) ApiUpdateOauthTokenExchangeProcessorPolicySettingsRequest {
	r.body = &body
	return r
}

// External validation will be bypassed when set to true. Default to false.
func (r ApiUpdateOauthTokenExchangeProcessorPolicySettingsRequest) BypassExternalValidation(bypassExternalValidation bool) ApiUpdateOauthTokenExchangeProcessorPolicySettingsRequest {
	r.bypassExternalValidation = &bypassExternalValidation
	return r
}

func (r ApiUpdateOauthTokenExchangeProcessorPolicySettingsRequest) Execute() (*TokenExchangeProcessorSettings, *http.Response, error) {
	return r.ApiService.UpdateOauthTokenExchangeProcessorPolicySettingsExecute(r)
}

/*
UpdateOauthTokenExchangeProcessorPolicySettings Update general OAuth 2.0 Token Exchange Processor settings.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiUpdateOauthTokenExchangeProcessorPolicySettingsRequest
*/
func (a *OauthTokenExchangeProcessorApiService) UpdateOauthTokenExchangeProcessorPolicySettings(ctx context.Context) ApiUpdateOauthTokenExchangeProcessorPolicySettingsRequest {
	return ApiUpdateOauthTokenExchangeProcessorPolicySettingsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return TokenExchangeProcessorSettings
func (a *OauthTokenExchangeProcessorApiService) UpdateOauthTokenExchangeProcessorPolicySettingsExecute(r ApiUpdateOauthTokenExchangeProcessorPolicySettingsRequest) (*TokenExchangeProcessorSettings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TokenExchangeProcessorSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthTokenExchangeProcessorApiService.UpdateOauthTokenExchangeProcessorPolicySettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/tokenExchange/processor/settings"

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
	if r.bypassExternalValidation != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "bypassExternalValidation", r.bypassExternalValidation, "")
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