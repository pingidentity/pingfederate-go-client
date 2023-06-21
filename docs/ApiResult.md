# ApiResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultId** | Pointer to **string** | Result identifier. | [optional] 
**Message** | Pointer to **string** | Success or error message. | [optional] 
**DeveloperMessage** | Pointer to **string** | Developer-oriented error message, if available. | [optional] 
**ValidationErrors** | Pointer to [**[]ValidationError**](ValidationError.md) | List of validation errors, if any. | [optional] 

## Methods

### NewApiResult

`func NewApiResult() *ApiResult`

NewApiResult instantiates a new ApiResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiResultWithDefaults

`func NewApiResultWithDefaults() *ApiResult`

NewApiResultWithDefaults instantiates a new ApiResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultId

`func (o *ApiResult) GetResultId() string`

GetResultId returns the ResultId field if non-nil, zero value otherwise.

### GetResultIdOk

`func (o *ApiResult) GetResultIdOk() (*string, bool)`

GetResultIdOk returns a tuple with the ResultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultId

`func (o *ApiResult) SetResultId(v string)`

SetResultId sets ResultId field to given value.

### HasResultId

`func (o *ApiResult) HasResultId() bool`

HasResultId returns a boolean if a field has been set.

### GetMessage

`func (o *ApiResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDeveloperMessage

`func (o *ApiResult) GetDeveloperMessage() string`

GetDeveloperMessage returns the DeveloperMessage field if non-nil, zero value otherwise.

### GetDeveloperMessageOk

`func (o *ApiResult) GetDeveloperMessageOk() (*string, bool)`

GetDeveloperMessageOk returns a tuple with the DeveloperMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperMessage

`func (o *ApiResult) SetDeveloperMessage(v string)`

SetDeveloperMessage sets DeveloperMessage field to given value.

### HasDeveloperMessage

`func (o *ApiResult) HasDeveloperMessage() bool`

HasDeveloperMessage returns a boolean if a field has been set.

### GetValidationErrors

`func (o *ApiResult) GetValidationErrors() []ValidationError`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *ApiResult) GetValidationErrorsOk() (*[]ValidationError, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *ApiResult) SetValidationErrors(v []ValidationError)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *ApiResult) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


