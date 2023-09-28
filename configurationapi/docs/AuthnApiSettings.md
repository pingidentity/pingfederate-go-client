# AuthnApiSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiEnabled** | Pointer to **bool** | Specifies whether the authentication API is enabled. The default value is false. | [optional] 
**DefaultApplicationRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**EnableApiDescriptions** | Pointer to **bool** | Enable the API Descriptions endpoint. | [optional] 
**RestrictAccessToRedirectlessMode** | Pointer to **bool** | Determines whether access to the authentication API redirectless mode is restricted to specified applications. | [optional] 
**IncludeRequestContext** | Pointer to **bool** | Determines whether the request context parameters are included in response for authentication API. The default value is false. | [optional] 

## Methods

### NewAuthnApiSettings

`func NewAuthnApiSettings() *AuthnApiSettings`

NewAuthnApiSettings instantiates a new AuthnApiSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthnApiSettingsWithDefaults

`func NewAuthnApiSettingsWithDefaults() *AuthnApiSettings`

NewAuthnApiSettingsWithDefaults instantiates a new AuthnApiSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiEnabled

`func (o *AuthnApiSettings) GetApiEnabled() bool`

GetApiEnabled returns the ApiEnabled field if non-nil, zero value otherwise.

### GetApiEnabledOk

`func (o *AuthnApiSettings) GetApiEnabledOk() (*bool, bool)`

GetApiEnabledOk returns a tuple with the ApiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiEnabled

`func (o *AuthnApiSettings) SetApiEnabled(v bool)`

SetApiEnabled sets ApiEnabled field to given value.

### HasApiEnabled

`func (o *AuthnApiSettings) HasApiEnabled() bool`

HasApiEnabled returns a boolean if a field has been set.

### GetDefaultApplicationRef

`func (o *AuthnApiSettings) GetDefaultApplicationRef() ResourceLink`

GetDefaultApplicationRef returns the DefaultApplicationRef field if non-nil, zero value otherwise.

### GetDefaultApplicationRefOk

`func (o *AuthnApiSettings) GetDefaultApplicationRefOk() (*ResourceLink, bool)`

GetDefaultApplicationRefOk returns a tuple with the DefaultApplicationRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultApplicationRef

`func (o *AuthnApiSettings) SetDefaultApplicationRef(v ResourceLink)`

SetDefaultApplicationRef sets DefaultApplicationRef field to given value.

### HasDefaultApplicationRef

`func (o *AuthnApiSettings) HasDefaultApplicationRef() bool`

HasDefaultApplicationRef returns a boolean if a field has been set.

### GetEnableApiDescriptions

`func (o *AuthnApiSettings) GetEnableApiDescriptions() bool`

GetEnableApiDescriptions returns the EnableApiDescriptions field if non-nil, zero value otherwise.

### GetEnableApiDescriptionsOk

`func (o *AuthnApiSettings) GetEnableApiDescriptionsOk() (*bool, bool)`

GetEnableApiDescriptionsOk returns a tuple with the EnableApiDescriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableApiDescriptions

`func (o *AuthnApiSettings) SetEnableApiDescriptions(v bool)`

SetEnableApiDescriptions sets EnableApiDescriptions field to given value.

### HasEnableApiDescriptions

`func (o *AuthnApiSettings) HasEnableApiDescriptions() bool`

HasEnableApiDescriptions returns a boolean if a field has been set.

### GetRestrictAccessToRedirectlessMode

`func (o *AuthnApiSettings) GetRestrictAccessToRedirectlessMode() bool`

GetRestrictAccessToRedirectlessMode returns the RestrictAccessToRedirectlessMode field if non-nil, zero value otherwise.

### GetRestrictAccessToRedirectlessModeOk

`func (o *AuthnApiSettings) GetRestrictAccessToRedirectlessModeOk() (*bool, bool)`

GetRestrictAccessToRedirectlessModeOk returns a tuple with the RestrictAccessToRedirectlessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictAccessToRedirectlessMode

`func (o *AuthnApiSettings) SetRestrictAccessToRedirectlessMode(v bool)`

SetRestrictAccessToRedirectlessMode sets RestrictAccessToRedirectlessMode field to given value.

### HasRestrictAccessToRedirectlessMode

`func (o *AuthnApiSettings) HasRestrictAccessToRedirectlessMode() bool`

HasRestrictAccessToRedirectlessMode returns a boolean if a field has been set.

### GetIncludeRequestContext

`func (o *AuthnApiSettings) GetIncludeRequestContext() bool`

GetIncludeRequestContext returns the IncludeRequestContext field if non-nil, zero value otherwise.

### GetIncludeRequestContextOk

`func (o *AuthnApiSettings) GetIncludeRequestContextOk() (*bool, bool)`

GetIncludeRequestContextOk returns a tuple with the IncludeRequestContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRequestContext

`func (o *AuthnApiSettings) SetIncludeRequestContext(v bool)`

SetIncludeRequestContext sets IncludeRequestContext field to given value.

### HasIncludeRequestContext

`func (o *AuthnApiSettings) HasIncludeRequestContext() bool`

HasIncludeRequestContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


