# ValidationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorId** | Pointer to **string** | Error identifier. | [optional] 
**Message** | Pointer to **string** | User-friendly error description. | [optional] 
**DeveloperMessage** | Pointer to **string** | Developer-oriented error message, if available. | [optional] 
**FieldPath** | Pointer to **string** | The path to the model field to which the error relates, if one exists. | [optional] 

## Methods

### NewValidationError

`func NewValidationError() *ValidationError`

NewValidationError instantiates a new ValidationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationErrorWithDefaults

`func NewValidationErrorWithDefaults() *ValidationError`

NewValidationErrorWithDefaults instantiates a new ValidationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorId

`func (o *ValidationError) GetErrorId() string`

GetErrorId returns the ErrorId field if non-nil, zero value otherwise.

### GetErrorIdOk

`func (o *ValidationError) GetErrorIdOk() (*string, bool)`

GetErrorIdOk returns a tuple with the ErrorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorId

`func (o *ValidationError) SetErrorId(v string)`

SetErrorId sets ErrorId field to given value.

### HasErrorId

`func (o *ValidationError) HasErrorId() bool`

HasErrorId returns a boolean if a field has been set.

### GetMessage

`func (o *ValidationError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ValidationError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ValidationError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ValidationError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDeveloperMessage

`func (o *ValidationError) GetDeveloperMessage() string`

GetDeveloperMessage returns the DeveloperMessage field if non-nil, zero value otherwise.

### GetDeveloperMessageOk

`func (o *ValidationError) GetDeveloperMessageOk() (*string, bool)`

GetDeveloperMessageOk returns a tuple with the DeveloperMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperMessage

`func (o *ValidationError) SetDeveloperMessage(v string)`

SetDeveloperMessage sets DeveloperMessage field to given value.

### HasDeveloperMessage

`func (o *ValidationError) HasDeveloperMessage() bool`

HasDeveloperMessage returns a boolean if a field has been set.

### GetFieldPath

`func (o *ValidationError) GetFieldPath() string`

GetFieldPath returns the FieldPath field if non-nil, zero value otherwise.

### GetFieldPathOk

`func (o *ValidationError) GetFieldPathOk() (*string, bool)`

GetFieldPathOk returns a tuple with the FieldPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldPath

`func (o *ValidationError) SetFieldPath(v string)`

SetFieldPath sets FieldPath field to given value.

### HasFieldPath

`func (o *ValidationError) HasFieldPath() bool`

HasFieldPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


