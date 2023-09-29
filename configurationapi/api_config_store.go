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

// ConfigStoreAPIService ConfigStoreAPI service
type ConfigStoreAPIService service

type ApiDeleteConfigStoreSettingRequest struct {
	ctx        context.Context
	ApiService *ConfigStoreAPIService
	bundle     string
	id         string
}

func (r ApiDeleteConfigStoreSettingRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteConfigStoreSettingExecute(r)
}

/*
DeleteConfigStoreSetting Delete a setting.

Delete a setting. This is an advanced operation with minimal validation. Incorrect use of this operation can harm the integrity of your PingFederate configuration. Please ensure you have specified the correct bundle name and setting ID before invoking this operation.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param bundle This field represents a configuration file that contains a bundle of settings.
	@param id ID of setting to delete.
	@return ApiDeleteConfigStoreSettingRequest
*/
func (a *ConfigStoreAPIService) DeleteConfigStoreSetting(ctx context.Context, bundle string, id string) ApiDeleteConfigStoreSettingRequest {
	return ApiDeleteConfigStoreSettingRequest{
		ApiService: a,
		ctx:        ctx,
		bundle:     bundle,
		id:         id,
	}
}

// Execute executes the request
func (a *ConfigStoreAPIService) DeleteConfigStoreSettingExecute(r ApiDeleteConfigStoreSettingRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigStoreAPIService.DeleteConfigStoreSetting")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/configStore/{bundle}/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"bundle"+"}", url.PathEscape(parameterValueToString(r.bundle, "bundle")), -1)
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

type ApiGetConfigStoreSettingRequest struct {
	ctx        context.Context
	ApiService *ConfigStoreAPIService
	bundle     string
	id         string
}

func (r ApiGetConfigStoreSettingRequest) Execute() (*ConfigStoreSetting, *http.Response, error) {
	return r.ApiService.GetConfigStoreSettingExecute(r)
}

/*
GetConfigStoreSetting Get a single setting from a bundle.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param bundle This field represents a configuration file that contains a bundle of settings.
	@param id ID of setting to retrieve.
	@return ApiGetConfigStoreSettingRequest
*/
func (a *ConfigStoreAPIService) GetConfigStoreSetting(ctx context.Context, bundle string, id string) ApiGetConfigStoreSettingRequest {
	return ApiGetConfigStoreSettingRequest{
		ApiService: a,
		ctx:        ctx,
		bundle:     bundle,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ConfigStoreSetting
func (a *ConfigStoreAPIService) GetConfigStoreSettingExecute(r ApiGetConfigStoreSettingRequest) (*ConfigStoreSetting, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ConfigStoreSetting
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigStoreAPIService.GetConfigStoreSetting")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/configStore/{bundle}/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"bundle"+"}", url.PathEscape(parameterValueToString(r.bundle, "bundle")), -1)
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

type ApiGetConfigStoreSettingsRequest struct {
	ctx        context.Context
	ApiService *ConfigStoreAPIService
	bundle     string
}

func (r ApiGetConfigStoreSettingsRequest) Execute() (*ConfigStoreBundle, *http.Response, error) {
	return r.ApiService.GetConfigStoreSettingsExecute(r)
}

/*
GetConfigStoreSettings Get all settings from a bundle.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param bundle This field represents a configuration file that contains a bundle of settings.
	@return ApiGetConfigStoreSettingsRequest
*/
func (a *ConfigStoreAPIService) GetConfigStoreSettings(ctx context.Context, bundle string) ApiGetConfigStoreSettingsRequest {
	return ApiGetConfigStoreSettingsRequest{
		ApiService: a,
		ctx:        ctx,
		bundle:     bundle,
	}
}

// Execute executes the request
//
//	@return ConfigStoreBundle
func (a *ConfigStoreAPIService) GetConfigStoreSettingsExecute(r ApiGetConfigStoreSettingsRequest) (*ConfigStoreBundle, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ConfigStoreBundle
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigStoreAPIService.GetConfigStoreSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/configStore/{bundle}"
	localVarPath = strings.Replace(localVarPath, "{"+"bundle"+"}", url.PathEscape(parameterValueToString(r.bundle, "bundle")), -1)

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

type ApiUpdateConfigStoreSettingRequest struct {
	ctx        context.Context
	ApiService *ConfigStoreAPIService
	bundle     string
	id         string
	body       *ConfigStoreSetting
}

// Configuration setting.
func (r ApiUpdateConfigStoreSettingRequest) Body(body ConfigStoreSetting) ApiUpdateConfigStoreSettingRequest {
	r.body = &body
	return r
}

func (r ApiUpdateConfigStoreSettingRequest) Execute() (*ConfigStoreSetting, *http.Response, error) {
	return r.ApiService.UpdateConfigStoreSettingExecute(r)
}

/*
UpdateConfigStoreSetting Create or update a setting/bundle.

Create or update a setting/bundle. This is an advanced operation with minimal validation. Incorrect use of this operation can harm the integrity of your PingFederate configuration. Please ensure you have specified the correct bundle name, setting ID, and setting value before invoking this operation.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param bundle This field represents a configuration file that contains a bundle of settings.
	@param id ID of setting to create/update.
	@return ApiUpdateConfigStoreSettingRequest
*/
func (a *ConfigStoreAPIService) UpdateConfigStoreSetting(ctx context.Context, bundle string, id string) ApiUpdateConfigStoreSettingRequest {
	return ApiUpdateConfigStoreSettingRequest{
		ApiService: a,
		ctx:        ctx,
		bundle:     bundle,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ConfigStoreSetting
func (a *ConfigStoreAPIService) UpdateConfigStoreSettingExecute(r ApiUpdateConfigStoreSettingRequest) (*ConfigStoreSetting, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ConfigStoreSetting
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigStoreAPIService.UpdateConfigStoreSetting")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/configStore/{bundle}/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"bundle"+"}", url.PathEscape(parameterValueToString(r.bundle, "bundle")), -1)
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
