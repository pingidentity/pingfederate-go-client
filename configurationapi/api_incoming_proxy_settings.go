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
)

// IncomingProxySettingsAPIService IncomingProxySettingsAPI service
type IncomingProxySettingsAPIService service

type ApiGetIncomingProxySettingsRequest struct {
	ctx        context.Context
	ApiService *IncomingProxySettingsAPIService
}

func (r ApiGetIncomingProxySettingsRequest) Execute() (*IncomingProxySettings, *http.Response, error) {
	return r.ApiService.GetIncomingProxySettingsExecute(r)
}

/*
GetIncomingProxySettings Get incoming proxy settings.

When PingFederate is deployed behind a proxy server or load balancer, use information in HTTP headers added by the proxy server to construct correct responses.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetIncomingProxySettingsRequest
*/
func (a *IncomingProxySettingsAPIService) GetIncomingProxySettings(ctx context.Context) ApiGetIncomingProxySettingsRequest {
	return ApiGetIncomingProxySettingsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return IncomingProxySettings
func (a *IncomingProxySettingsAPIService) GetIncomingProxySettingsExecute(r ApiGetIncomingProxySettingsRequest) (*IncomingProxySettings, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *IncomingProxySettings
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalGetIncomingProxySettingsExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *IncomingProxySettingsAPIService) internalGetIncomingProxySettingsExecute(r ApiGetIncomingProxySettingsRequest) (*IncomingProxySettings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *IncomingProxySettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IncomingProxySettingsAPIService.GetIncomingProxySettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/incomingProxySettings"

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

type ApiUpdateIncomingProxySettingsRequest struct {
	ctx        context.Context
	ApiService *IncomingProxySettingsAPIService
	body       *IncomingProxySettings
}

// Incoming proxy settings.
func (r ApiUpdateIncomingProxySettingsRequest) Body(body IncomingProxySettings) ApiUpdateIncomingProxySettingsRequest {
	r.body = &body
	return r
}

func (r ApiUpdateIncomingProxySettingsRequest) Execute() (*IncomingProxySettings, *http.Response, error) {
	return r.ApiService.UpdateIncomingProxySettingsExecute(r)
}

/*
UpdateIncomingProxySettings Update incoming proxy settings.

When PingFederate is deployed behind a proxy server or load balancer, use information in HTTP headers added by the proxy server to construct correct responses.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiUpdateIncomingProxySettingsRequest
*/
func (a *IncomingProxySettingsAPIService) UpdateIncomingProxySettings(ctx context.Context) ApiUpdateIncomingProxySettingsRequest {
	return ApiUpdateIncomingProxySettingsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return IncomingProxySettings
func (a *IncomingProxySettingsAPIService) UpdateIncomingProxySettingsExecute(r ApiUpdateIncomingProxySettingsRequest) (*IncomingProxySettings, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *IncomingProxySettings
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalUpdateIncomingProxySettingsExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *IncomingProxySettingsAPIService) internalUpdateIncomingProxySettingsExecute(r ApiUpdateIncomingProxySettingsRequest) (*IncomingProxySettings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *IncomingProxySettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IncomingProxySettingsAPIService.UpdateIncomingProxySettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/incomingProxySettings"

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
