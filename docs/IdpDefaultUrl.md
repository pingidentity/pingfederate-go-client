# IdpDefaultUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfirmIdpSlo** | Pointer to **bool** | Prompt user to confirm Single Logout (SLO). | [optional] 
**IdpSloSuccessUrl** | Pointer to **string** | Provide the default URL you would like to send the user to when Single Logout has succeeded. | [optional] 
**IdpErrorMsg** | **string** | Provide the error text displayed in a user&#39;s browser when an SSO operation fails. | 

## Methods

### NewIdpDefaultUrl

`func NewIdpDefaultUrl(idpErrorMsg string, ) *IdpDefaultUrl`

NewIdpDefaultUrl instantiates a new IdpDefaultUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpDefaultUrlWithDefaults

`func NewIdpDefaultUrlWithDefaults() *IdpDefaultUrl`

NewIdpDefaultUrlWithDefaults instantiates a new IdpDefaultUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfirmIdpSlo

`func (o *IdpDefaultUrl) GetConfirmIdpSlo() bool`

GetConfirmIdpSlo returns the ConfirmIdpSlo field if non-nil, zero value otherwise.

### GetConfirmIdpSloOk

`func (o *IdpDefaultUrl) GetConfirmIdpSloOk() (*bool, bool)`

GetConfirmIdpSloOk returns a tuple with the ConfirmIdpSlo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmIdpSlo

`func (o *IdpDefaultUrl) SetConfirmIdpSlo(v bool)`

SetConfirmIdpSlo sets ConfirmIdpSlo field to given value.

### HasConfirmIdpSlo

`func (o *IdpDefaultUrl) HasConfirmIdpSlo() bool`

HasConfirmIdpSlo returns a boolean if a field has been set.

### GetIdpSloSuccessUrl

`func (o *IdpDefaultUrl) GetIdpSloSuccessUrl() string`

GetIdpSloSuccessUrl returns the IdpSloSuccessUrl field if non-nil, zero value otherwise.

### GetIdpSloSuccessUrlOk

`func (o *IdpDefaultUrl) GetIdpSloSuccessUrlOk() (*string, bool)`

GetIdpSloSuccessUrlOk returns a tuple with the IdpSloSuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpSloSuccessUrl

`func (o *IdpDefaultUrl) SetIdpSloSuccessUrl(v string)`

SetIdpSloSuccessUrl sets IdpSloSuccessUrl field to given value.

### HasIdpSloSuccessUrl

`func (o *IdpDefaultUrl) HasIdpSloSuccessUrl() bool`

HasIdpSloSuccessUrl returns a boolean if a field has been set.

### GetIdpErrorMsg

`func (o *IdpDefaultUrl) GetIdpErrorMsg() string`

GetIdpErrorMsg returns the IdpErrorMsg field if non-nil, zero value otherwise.

### GetIdpErrorMsgOk

`func (o *IdpDefaultUrl) GetIdpErrorMsgOk() (*string, bool)`

GetIdpErrorMsgOk returns a tuple with the IdpErrorMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpErrorMsg

`func (o *IdpDefaultUrl) SetIdpErrorMsg(v string)`

SetIdpErrorMsg sets IdpErrorMsg field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


