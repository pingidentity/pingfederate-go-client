# OpenIdConnectSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultPolicyRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**SessionSettings** | Pointer to [**OIDCSessionSettings**](OIDCSessionSettings.md) |  | [optional] 

## Methods

### NewOpenIdConnectSettings

`func NewOpenIdConnectSettings() *OpenIdConnectSettings`

NewOpenIdConnectSettings instantiates a new OpenIdConnectSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenIdConnectSettingsWithDefaults

`func NewOpenIdConnectSettingsWithDefaults() *OpenIdConnectSettings`

NewOpenIdConnectSettingsWithDefaults instantiates a new OpenIdConnectSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultPolicyRef

`func (o *OpenIdConnectSettings) GetDefaultPolicyRef() ResourceLink`

GetDefaultPolicyRef returns the DefaultPolicyRef field if non-nil, zero value otherwise.

### GetDefaultPolicyRefOk

`func (o *OpenIdConnectSettings) GetDefaultPolicyRefOk() (*ResourceLink, bool)`

GetDefaultPolicyRefOk returns a tuple with the DefaultPolicyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPolicyRef

`func (o *OpenIdConnectSettings) SetDefaultPolicyRef(v ResourceLink)`

SetDefaultPolicyRef sets DefaultPolicyRef field to given value.

### HasDefaultPolicyRef

`func (o *OpenIdConnectSettings) HasDefaultPolicyRef() bool`

HasDefaultPolicyRef returns a boolean if a field has been set.

### GetSessionSettings

`func (o *OpenIdConnectSettings) GetSessionSettings() OIDCSessionSettings`

GetSessionSettings returns the SessionSettings field if non-nil, zero value otherwise.

### GetSessionSettingsOk

`func (o *OpenIdConnectSettings) GetSessionSettingsOk() (*OIDCSessionSettings, bool)`

GetSessionSettingsOk returns a tuple with the SessionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSettings

`func (o *OpenIdConnectSettings) SetSessionSettings(v OIDCSessionSettings)`

SetSessionSettings sets SessionSettings field to given value.

### HasSessionSettings

`func (o *OpenIdConnectSettings) HasSessionSettings() bool`

HasSessionSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


