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

// BulkAPIService BulkAPI service
type BulkAPIService service

type ApiExportConfigurationRequest struct {
	ctx                      context.Context
	ApiService               *BulkAPIService
	includeExternalResources *bool
}

// Include external resources like OAuth clients stored outside of PingFederate.
func (r ApiExportConfigurationRequest) IncludeExternalResources(includeExternalResources bool) ApiExportConfigurationRequest {
	r.includeExternalResources = &includeExternalResources
	return r
}

func (r ApiExportConfigurationRequest) Execute() (*BulkConfig, *http.Response, error) {
	return r.ApiService.ExportConfigurationExecute(r)
}

/*
ExportConfiguration Export all API resources to a JSON file.

For the /configStore resource type, only the settings that are different from the defaults for this version of PingFederate are included in the export.<br><br>Only resource types currently supported by the Administrative API are included in the exported data. Resources not yet supported include:<br><br>- SMS Provider Settings<br>

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiExportConfigurationRequest
*/
func (a *BulkAPIService) ExportConfiguration(ctx context.Context) ApiExportConfigurationRequest {
	return ApiExportConfigurationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BulkConfig
func (a *BulkAPIService) ExportConfigurationExecute(r ApiExportConfigurationRequest) (*BulkConfig, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *BulkConfig
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalExportConfigurationExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *BulkAPIService) internalExportConfigurationExecute(r ApiExportConfigurationRequest) (*BulkConfig, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BulkConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BulkAPIService.ExportConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bulk/export"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includeExternalResources != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeExternalResources", r.includeExternalResources, "")
	} else {
		var defaultValue bool = false
		r.includeExternalResources = &defaultValue
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

type ApiImportConfigurationRequest struct {
	ctx                       context.Context
	ApiService                *BulkAPIService
	body                      *BulkConfig
	failFast                  *bool
	xBypassExternalValidation *bool
}

// Configuration to import.
func (r ApiImportConfigurationRequest) Body(body BulkConfig) ApiImportConfigurationRequest {
	r.body = &body
	return r
}

// When set to true, stops the import as soon as any validation errors are encountered. When false, import will continue to validate configuration after the first failure to identify all validation errors. If any validation errors are present PingFederate will roll back to the state prior to the import attempt.
func (r ApiImportConfigurationRequest) FailFast(failFast bool) ApiImportConfigurationRequest {
	r.failFast = &failFast
	return r
}

// External validation will be bypassed when set to true. Default to false.
func (r ApiImportConfigurationRequest) XBypassExternalValidation(xBypassExternalValidation bool) ApiImportConfigurationRequest {
	r.xBypassExternalValidation = &xBypassExternalValidation
	return r
}

func (r ApiImportConfigurationRequest) Execute() (*http.Response, error) {
	return r.ApiService.ImportConfigurationExecute(r)
}

/*
ImportConfiguration Import configuration for a PingFederate deployment from a JSON file.

All existing configuration will be wiped before the import begins. If any validation errors are found, PingFederate will roll back to the previous configuration. The master key set in pf.jwk must include the key in use when the JSON configuration was originally exported.

JSON configurations exported from earlier versions of PingFederate can be imported. However, this is not a supported way of upgrading from an earlier version. Instead, you should run the upgrade utility and then reexport to get an updated configuration JSON.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiImportConfigurationRequest
*/
func (a *BulkAPIService) ImportConfiguration(ctx context.Context) ApiImportConfigurationRequest {
	return ApiImportConfigurationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *BulkAPIService) ImportConfigurationExecute(r ApiImportConfigurationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BulkAPIService.ImportConfiguration")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bulk/import"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return nil, reportError("body is required and must be specified")
	}

	if r.failFast != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "failFast", r.failFast, "")
	} else {
		var defaultValue bool = true
		r.failFast = &defaultValue
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
