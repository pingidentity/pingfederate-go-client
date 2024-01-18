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
)

// RedirectValidationAPIService RedirectValidationAPI service
type RedirectValidationAPIService service

type ApiGetRedirectValidationSettingsRequest struct {
	ctx        context.Context
	ApiService *RedirectValidationAPIService
}

func (r ApiGetRedirectValidationSettingsRequest) Execute() (*RedirectValidationSettings, *http.Response, error) {
	return r.ApiService.GetRedirectValidationSettingsExecute(r)
}

/*
GetRedirectValidationSettings Retrieve redirect validation settings.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetRedirectValidationSettingsRequest
*/
func (a *RedirectValidationAPIService) GetRedirectValidationSettings(ctx context.Context) ApiGetRedirectValidationSettingsRequest {
	return ApiGetRedirectValidationSettingsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RedirectValidationSettings
func (a *RedirectValidationAPIService) GetRedirectValidationSettingsExecute(r ApiGetRedirectValidationSettingsRequest) (*RedirectValidationSettings, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *RedirectValidationSettings
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalGetRedirectValidationSettingsExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *RedirectValidationAPIService) internalGetRedirectValidationSettingsExecute(r ApiGetRedirectValidationSettingsRequest) (*RedirectValidationSettings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RedirectValidationSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RedirectValidationAPIService.GetRedirectValidationSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/redirectValidation"

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

type ApiUpdateRedirectValidationSettingsRequest struct {
	ctx        context.Context
	ApiService *RedirectValidationAPIService
	body       *RedirectValidationSettings
}

// Redirect validation settings.
func (r ApiUpdateRedirectValidationSettingsRequest) Body(body RedirectValidationSettings) ApiUpdateRedirectValidationSettingsRequest {
	r.body = &body
	return r
}

func (r ApiUpdateRedirectValidationSettingsRequest) Execute() (*RedirectValidationSettings, *http.Response, error) {
	return r.ApiService.UpdateRedirectValidationSettingsExecute(r)
}

/*
UpdateRedirectValidationSettings Update redirect validation settings.

<b>Note: </b>Ensure IdP Discovery and/or WS-Federation is enabled for redirect validation to function for IdP Discovery and/or wreply for SLO respectively.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiUpdateRedirectValidationSettingsRequest
*/
func (a *RedirectValidationAPIService) UpdateRedirectValidationSettings(ctx context.Context) ApiUpdateRedirectValidationSettingsRequest {
	return ApiUpdateRedirectValidationSettingsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RedirectValidationSettings
func (a *RedirectValidationAPIService) UpdateRedirectValidationSettingsExecute(r ApiUpdateRedirectValidationSettingsRequest) (*RedirectValidationSettings, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *RedirectValidationSettings
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalUpdateRedirectValidationSettingsExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *RedirectValidationAPIService) internalUpdateRedirectValidationSettingsExecute(r ApiUpdateRedirectValidationSettingsRequest) (*RedirectValidationSettings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RedirectValidationSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RedirectValidationAPIService.UpdateRedirectValidationSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/redirectValidation"

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
