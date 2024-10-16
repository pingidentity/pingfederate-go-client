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

// VirtualHostNamesAPIService VirtualHostNamesAPI service
type VirtualHostNamesAPIService service

type ApiGetVirtualHostNamesSettingsRequest struct {
	ctx        context.Context
	ApiService *VirtualHostNamesAPIService
}

func (r ApiGetVirtualHostNamesSettingsRequest) Execute() (*VirtualHostNameSettings, *http.Response, error) {
	return r.ApiService.GetVirtualHostNamesSettingsExecute(r)
}

/*
GetVirtualHostNamesSettings Retrieve virtual host names settings.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetVirtualHostNamesSettingsRequest
*/
func (a *VirtualHostNamesAPIService) GetVirtualHostNamesSettings(ctx context.Context) ApiGetVirtualHostNamesSettingsRequest {
	return ApiGetVirtualHostNamesSettingsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return VirtualHostNameSettings
func (a *VirtualHostNamesAPIService) GetVirtualHostNamesSettingsExecute(r ApiGetVirtualHostNamesSettingsRequest) (*VirtualHostNameSettings, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *VirtualHostNameSettings
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalGetVirtualHostNamesSettingsExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *VirtualHostNamesAPIService) internalGetVirtualHostNamesSettingsExecute(r ApiGetVirtualHostNamesSettingsRequest) (*VirtualHostNameSettings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *VirtualHostNameSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VirtualHostNamesAPIService.GetVirtualHostNamesSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/virtualHostNames"

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

type ApiUpdateVirtualHostNamesSettingsRequest struct {
	ctx        context.Context
	ApiService *VirtualHostNamesAPIService
	body       *VirtualHostNameSettings
}

// Virtual host names settings.
func (r ApiUpdateVirtualHostNamesSettingsRequest) Body(body VirtualHostNameSettings) ApiUpdateVirtualHostNamesSettingsRequest {
	r.body = &body
	return r
}

func (r ApiUpdateVirtualHostNamesSettingsRequest) Execute() (*VirtualHostNameSettings, *http.Response, error) {
	return r.ApiService.UpdateVirtualHostNamesSettingsExecute(r)
}

/*
UpdateVirtualHostNamesSettings Update virtual host names settings.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiUpdateVirtualHostNamesSettingsRequest
*/
func (a *VirtualHostNamesAPIService) UpdateVirtualHostNamesSettings(ctx context.Context) ApiUpdateVirtualHostNamesSettingsRequest {
	return ApiUpdateVirtualHostNamesSettingsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return VirtualHostNameSettings
func (a *VirtualHostNamesAPIService) UpdateVirtualHostNamesSettingsExecute(r ApiUpdateVirtualHostNamesSettingsRequest) (*VirtualHostNameSettings, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *VirtualHostNameSettings
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalUpdateVirtualHostNamesSettingsExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *VirtualHostNamesAPIService) internalUpdateVirtualHostNamesSettingsExecute(r ApiUpdateVirtualHostNamesSettingsRequest) (*VirtualHostNameSettings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *VirtualHostNameSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VirtualHostNamesAPIService.UpdateVirtualHostNamesSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/virtualHostNames"

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
