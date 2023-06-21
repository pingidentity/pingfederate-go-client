# CSRResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileData** | **string** | The CSR response file data in PKCS7 format or as an X.509 certificate. PEM encoding (with or without the header and footer lines) is required. New line characters should be omitted or encoded in this value. | 

## Methods

### NewCSRResponse

`func NewCSRResponse(fileData string, ) *CSRResponse`

NewCSRResponse instantiates a new CSRResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSRResponseWithDefaults

`func NewCSRResponseWithDefaults() *CSRResponse`

NewCSRResponseWithDefaults instantiates a new CSRResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileData

`func (o *CSRResponse) GetFileData() string`

GetFileData returns the FileData field if non-nil, zero value otherwise.

### GetFileDataOk

`func (o *CSRResponse) GetFileDataOk() (*string, bool)`

GetFileDataOk returns a tuple with the FileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileData

`func (o *CSRResponse) SetFileData(v string)`

SetFileData sets FileData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


