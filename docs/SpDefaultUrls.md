# SpDefaultUrls

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SsoSuccessUrl** | Pointer to **string** | Provide the default URL you would like to send the user to when Single Sign On (SSO) has succeeded. | [optional] 
**ConfirmSlo** | Pointer to **bool** | Determines whether the user is prompted to confirm Single Logout (SLO). The default is false. | [optional] 
**SloSuccessUrl** | Pointer to **string** | Provide the default URL you would like to send the user to when Single Logout (SLO) has succeeded. | [optional] 

## Methods

### NewSpDefaultUrls

`func NewSpDefaultUrls() *SpDefaultUrls`

NewSpDefaultUrls instantiates a new SpDefaultUrls object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpDefaultUrlsWithDefaults

`func NewSpDefaultUrlsWithDefaults() *SpDefaultUrls`

NewSpDefaultUrlsWithDefaults instantiates a new SpDefaultUrls object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsoSuccessUrl

`func (o *SpDefaultUrls) GetSsoSuccessUrl() string`

GetSsoSuccessUrl returns the SsoSuccessUrl field if non-nil, zero value otherwise.

### GetSsoSuccessUrlOk

`func (o *SpDefaultUrls) GetSsoSuccessUrlOk() (*string, bool)`

GetSsoSuccessUrlOk returns a tuple with the SsoSuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoSuccessUrl

`func (o *SpDefaultUrls) SetSsoSuccessUrl(v string)`

SetSsoSuccessUrl sets SsoSuccessUrl field to given value.

### HasSsoSuccessUrl

`func (o *SpDefaultUrls) HasSsoSuccessUrl() bool`

HasSsoSuccessUrl returns a boolean if a field has been set.

### GetConfirmSlo

`func (o *SpDefaultUrls) GetConfirmSlo() bool`

GetConfirmSlo returns the ConfirmSlo field if non-nil, zero value otherwise.

### GetConfirmSloOk

`func (o *SpDefaultUrls) GetConfirmSloOk() (*bool, bool)`

GetConfirmSloOk returns a tuple with the ConfirmSlo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmSlo

`func (o *SpDefaultUrls) SetConfirmSlo(v bool)`

SetConfirmSlo sets ConfirmSlo field to given value.

### HasConfirmSlo

`func (o *SpDefaultUrls) HasConfirmSlo() bool`

HasConfirmSlo returns a boolean if a field has been set.

### GetSloSuccessUrl

`func (o *SpDefaultUrls) GetSloSuccessUrl() string`

GetSloSuccessUrl returns the SloSuccessUrl field if non-nil, zero value otherwise.

### GetSloSuccessUrlOk

`func (o *SpDefaultUrls) GetSloSuccessUrlOk() (*string, bool)`

GetSloSuccessUrlOk returns a tuple with the SloSuccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloSuccessUrl

`func (o *SpDefaultUrls) SetSloSuccessUrl(v string)`

SetSloSuccessUrl sets SloSuccessUrl field to given value.

### HasSloSuccessUrl

`func (o *SpDefaultUrls) HasSloSuccessUrl() bool`

HasSloSuccessUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


