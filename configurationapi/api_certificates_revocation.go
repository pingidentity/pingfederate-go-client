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

// CertificatesRevocationAPIService CertificatesRevocationAPI service
type CertificatesRevocationAPIService service

type ApiDeleteOcspCertificateByIdRequest struct {
	ctx        context.Context
	ApiService *CertificatesRevocationAPIService
	id         string
}

func (r ApiDeleteOcspCertificateByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteOcspCertificateByIdExecute(r)
}

/*
DeleteOcspCertificateById Delete an OCSP responder signature verification certificate by ID.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Certificate ID.
	@return ApiDeleteOcspCertificateByIdRequest
*/
func (a *CertificatesRevocationAPIService) DeleteOcspCertificateById(ctx context.Context, id string) ApiDeleteOcspCertificateByIdRequest {
	return ApiDeleteOcspCertificateByIdRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *CertificatesRevocationAPIService) DeleteOcspCertificateByIdExecute(r ApiDeleteOcspCertificateByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificatesRevocationAPIService.DeleteOcspCertificateById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/certificates/revocation/ocspCertificates/{id}"
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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetOcspCertificateByIdRequest struct {
	ctx        context.Context
	ApiService *CertificatesRevocationAPIService
	id         string
}

func (r ApiGetOcspCertificateByIdRequest) Execute() (*CertView, *http.Response, error) {
	return r.ApiService.GetOcspCertificateByIdExecute(r)
}

/*
GetOcspCertificateById Get an OCSP responder signature verification certificate by ID.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Certificate ID.
	@return ApiGetOcspCertificateByIdRequest
*/
func (a *CertificatesRevocationAPIService) GetOcspCertificateById(ctx context.Context, id string) ApiGetOcspCertificateByIdRequest {
	return ApiGetOcspCertificateByIdRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return CertView
func (a *CertificatesRevocationAPIService) GetOcspCertificateByIdExecute(r ApiGetOcspCertificateByIdRequest) (*CertView, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *CertView
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalGetOcspCertificateByIdExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *CertificatesRevocationAPIService) internalGetOcspCertificateByIdExecute(r ApiGetOcspCertificateByIdRequest) (*CertView, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CertView
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificatesRevocationAPIService.GetOcspCertificateById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/certificates/revocation/ocspCertificates/{id}"
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

type ApiGetOcspCertificatesRequest struct {
	ctx        context.Context
	ApiService *CertificatesRevocationAPIService
}

func (r ApiGetOcspCertificatesRequest) Execute() (*CertViews, *http.Response, error) {
	return r.ApiService.GetOcspCertificatesExecute(r)
}

/*
GetOcspCertificates Get the list of available OCSP responder signature verification certificates.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetOcspCertificatesRequest
*/
func (a *CertificatesRevocationAPIService) GetOcspCertificates(ctx context.Context) ApiGetOcspCertificatesRequest {
	return ApiGetOcspCertificatesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CertViews
func (a *CertificatesRevocationAPIService) GetOcspCertificatesExecute(r ApiGetOcspCertificatesRequest) (*CertViews, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *CertViews
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalGetOcspCertificatesExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *CertificatesRevocationAPIService) internalGetOcspCertificatesExecute(r ApiGetOcspCertificatesRequest) (*CertViews, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CertViews
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificatesRevocationAPIService.GetOcspCertificates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/certificates/revocation/ocspCertificates"

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

type ApiGetRevocationSettingsRequest struct {
	ctx        context.Context
	ApiService *CertificatesRevocationAPIService
}

func (r ApiGetRevocationSettingsRequest) Execute() (*CertificateRevocationSettings, *http.Response, error) {
	return r.ApiService.GetRevocationSettingsExecute(r)
}

/*
GetRevocationSettings Get certificate revocation settings.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetRevocationSettingsRequest
*/
func (a *CertificatesRevocationAPIService) GetRevocationSettings(ctx context.Context) ApiGetRevocationSettingsRequest {
	return ApiGetRevocationSettingsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CertificateRevocationSettings
func (a *CertificatesRevocationAPIService) GetRevocationSettingsExecute(r ApiGetRevocationSettingsRequest) (*CertificateRevocationSettings, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *CertificateRevocationSettings
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalGetRevocationSettingsExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *CertificatesRevocationAPIService) internalGetRevocationSettingsExecute(r ApiGetRevocationSettingsRequest) (*CertificateRevocationSettings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CertificateRevocationSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificatesRevocationAPIService.GetRevocationSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/certificates/revocation/settings"

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

type ApiImportOcspCertificateRequest struct {
	ctx        context.Context
	ApiService *CertificatesRevocationAPIService
	body       *X509File
}

// File to import.
func (r ApiImportOcspCertificateRequest) Body(body X509File) ApiImportOcspCertificateRequest {
	r.body = &body
	return r
}

func (r ApiImportOcspCertificateRequest) Execute() (*CertView, *http.Response, error) {
	return r.ApiService.ImportOcspCertificateExecute(r)
}

/*
ImportOcspCertificate Import an OCSP responder signature verification certificate.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiImportOcspCertificateRequest
*/
func (a *CertificatesRevocationAPIService) ImportOcspCertificate(ctx context.Context) ApiImportOcspCertificateRequest {
	return ApiImportOcspCertificateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CertView
func (a *CertificatesRevocationAPIService) ImportOcspCertificateExecute(r ApiImportOcspCertificateRequest) (*CertView, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *CertView
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalImportOcspCertificateExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *CertificatesRevocationAPIService) internalImportOcspCertificateExecute(r ApiImportOcspCertificateRequest) (*CertView, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CertView
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificatesRevocationAPIService.ImportOcspCertificate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/certificates/revocation/ocspCertificates"

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

type ApiUpdateRevocationSettingsRequest struct {
	ctx        context.Context
	ApiService *CertificatesRevocationAPIService
	body       *CertificateRevocationSettings
}

// Certificate revocation settings.
func (r ApiUpdateRevocationSettingsRequest) Body(body CertificateRevocationSettings) ApiUpdateRevocationSettingsRequest {
	r.body = &body
	return r
}

func (r ApiUpdateRevocationSettingsRequest) Execute() (*CertificateRevocationSettings, *http.Response, error) {
	return r.ApiService.UpdateRevocationSettingsExecute(r)
}

/*
UpdateRevocationSettings Update certificate revocation settings.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiUpdateRevocationSettingsRequest
*/
func (a *CertificatesRevocationAPIService) UpdateRevocationSettings(ctx context.Context) ApiUpdateRevocationSettingsRequest {
	return ApiUpdateRevocationSettingsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CertificateRevocationSettings
func (a *CertificatesRevocationAPIService) UpdateRevocationSettingsExecute(r ApiUpdateRevocationSettingsRequest) (*CertificateRevocationSettings, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *CertificateRevocationSettings
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalUpdateRevocationSettingsExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *CertificatesRevocationAPIService) internalUpdateRevocationSettingsExecute(r ApiUpdateRevocationSettingsRequest) (*CertificateRevocationSettings, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CertificateRevocationSettings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificatesRevocationAPIService.UpdateRevocationSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/certificates/revocation/settings"

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
