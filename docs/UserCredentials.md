# UserCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPassword** | Pointer to **string** | Current password. Required only during Password Change and not applicable for Password Reset. | [optional] 
**NewPassword** | **string** | A new password. | 

## Methods

### NewUserCredentials

`func NewUserCredentials(newPassword string, ) *UserCredentials`

NewUserCredentials instantiates a new UserCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCredentialsWithDefaults

`func NewUserCredentialsWithDefaults() *UserCredentials`

NewUserCredentialsWithDefaults instantiates a new UserCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPassword

`func (o *UserCredentials) GetCurrentPassword() string`

GetCurrentPassword returns the CurrentPassword field if non-nil, zero value otherwise.

### GetCurrentPasswordOk

`func (o *UserCredentials) GetCurrentPasswordOk() (*string, bool)`

GetCurrentPasswordOk returns a tuple with the CurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPassword

`func (o *UserCredentials) SetCurrentPassword(v string)`

SetCurrentPassword sets CurrentPassword field to given value.

### HasCurrentPassword

`func (o *UserCredentials) HasCurrentPassword() bool`

HasCurrentPassword returns a boolean if a field has been set.

### GetNewPassword

`func (o *UserCredentials) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *UserCredentials) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *UserCredentials) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


