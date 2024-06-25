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
	"os"
)

// ConfigArchiveAPIService ConfigArchiveAPI service
type ConfigArchiveAPIService service

type ApiExportConfigArchiveRequest struct {
	ctx        context.Context
	ApiService *ConfigArchiveAPIService
}

func (r ApiExportConfigArchiveRequest) Execute() (*http.Response, error) {
	return r.ApiService.ExportConfigArchiveExecute(r)
}

/*
ExportConfigArchive Export a configuration archive.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiExportConfigArchiveRequest
*/
func (a *ConfigArchiveAPIService) ExportConfigArchive(ctx context.Context) ApiExportConfigArchiveRequest {
	return ApiExportConfigArchiveRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ConfigArchiveAPIService) ExportConfigArchiveExecute(r ApiExportConfigArchiveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigArchiveAPIService.ExportConfigArchive")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/configArchive/export"

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
	localVarHTTPHeaderAccepts := []string{}

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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiImportConfigArchiveRequest struct {
	ctx                    context.Context
	ApiService             *ConfigArchiveAPIService
	forceImport            *bool
	forceUnsupportedImport *bool
	reencryptData          *bool
	file                   *os.File
}

func (r ApiImportConfigArchiveRequest) ForceImport(forceImport bool) ApiImportConfigArchiveRequest {
	r.forceImport = &forceImport
	return r
}

// Force import of unsupported versions.
func (r ApiImportConfigArchiveRequest) ForceUnsupportedImport(forceUnsupportedImport bool) ApiImportConfigArchiveRequest {
	r.forceUnsupportedImport = &forceUnsupportedImport
	return r
}

// Reencrypt configuration archive data with the current deployment&#39;s encryption key.
func (r ApiImportConfigArchiveRequest) ReencryptData(reencryptData bool) ApiImportConfigArchiveRequest {
	r.reencryptData = &reencryptData
	return r
}

func (r ApiImportConfigArchiveRequest) File(file *os.File) ApiImportConfigArchiveRequest {
	r.file = file
	return r
}

func (r ApiImportConfigArchiveRequest) Execute() (*ApiResult, *http.Response, error) {
	return r.ApiService.ImportConfigArchiveExecute(r)
}

/*
ImportConfigArchive Import a configuration archive.

If there are missing components or license inconsistencies, the import is halted by default to allow you to install the necessary components or license. However, you can choose to force the deployment by setting 'forceImport' to true and then install the necessary files later.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiImportConfigArchiveRequest
*/
func (a *ConfigArchiveAPIService) ImportConfigArchive(ctx context.Context) ApiImportConfigArchiveRequest {
	return ApiImportConfigArchiveRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ApiResult
func (a *ConfigArchiveAPIService) ImportConfigArchiveExecute(r ApiImportConfigArchiveRequest) (*ApiResult, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *ApiResult
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalImportConfigArchiveExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *ConfigArchiveAPIService) internalImportConfigArchiveExecute(r ApiImportConfigArchiveRequest) (*ApiResult, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ApiResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigArchiveAPIService.ImportConfigArchive")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/configArchive/import"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.forceImport != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "forceImport", r.forceImport, "")
	}
	if r.forceUnsupportedImport != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "forceUnsupportedImport", r.forceUnsupportedImport, "")
	} else {
		var defaultValue bool = false
		r.forceUnsupportedImport = &defaultValue
	}
	if r.reencryptData != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "reencryptData", r.reencryptData, "")
	} else {
		var defaultValue bool = false
		r.reencryptData = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	var fileLocalVarFormFileName string
	var fileLocalVarFileName string
	var fileLocalVarFileBytes []byte

	fileLocalVarFormFileName = "file"

	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
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
