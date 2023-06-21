# RedirectValidationLocalSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableTargetResourceValidationForSSO** | Pointer to **bool** | Enable target resource validation for SSO. | [optional] 
**EnableTargetResourceValidationForSLO** | Pointer to **bool** | Enable target resource validation for SLO. | [optional] 
**EnableTargetResourceValidationForIdpDiscovery** | Pointer to **bool** | Enable target resource validation for IdP discovery. | [optional] 
**EnableInErrorResourceValidation** | Pointer to **bool** | Enable validation for error resource. | [optional] 
**WhiteList** | Pointer to [**[]RedirectValidationSettingsWhitelistEntry**](RedirectValidationSettingsWhitelistEntry.md) | List of URLs that are designated as valid target resources. | [optional] 

## Methods

### NewRedirectValidationLocalSettings

`func NewRedirectValidationLocalSettings() *RedirectValidationLocalSettings`

NewRedirectValidationLocalSettings instantiates a new RedirectValidationLocalSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedirectValidationLocalSettingsWithDefaults

`func NewRedirectValidationLocalSettingsWithDefaults() *RedirectValidationLocalSettings`

NewRedirectValidationLocalSettingsWithDefaults instantiates a new RedirectValidationLocalSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableTargetResourceValidationForSSO

`func (o *RedirectValidationLocalSettings) GetEnableTargetResourceValidationForSSO() bool`

GetEnableTargetResourceValidationForSSO returns the EnableTargetResourceValidationForSSO field if non-nil, zero value otherwise.

### GetEnableTargetResourceValidationForSSOOk

`func (o *RedirectValidationLocalSettings) GetEnableTargetResourceValidationForSSOOk() (*bool, bool)`

GetEnableTargetResourceValidationForSSOOk returns a tuple with the EnableTargetResourceValidationForSSO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTargetResourceValidationForSSO

`func (o *RedirectValidationLocalSettings) SetEnableTargetResourceValidationForSSO(v bool)`

SetEnableTargetResourceValidationForSSO sets EnableTargetResourceValidationForSSO field to given value.

### HasEnableTargetResourceValidationForSSO

`func (o *RedirectValidationLocalSettings) HasEnableTargetResourceValidationForSSO() bool`

HasEnableTargetResourceValidationForSSO returns a boolean if a field has been set.

### GetEnableTargetResourceValidationForSLO

`func (o *RedirectValidationLocalSettings) GetEnableTargetResourceValidationForSLO() bool`

GetEnableTargetResourceValidationForSLO returns the EnableTargetResourceValidationForSLO field if non-nil, zero value otherwise.

### GetEnableTargetResourceValidationForSLOOk

`func (o *RedirectValidationLocalSettings) GetEnableTargetResourceValidationForSLOOk() (*bool, bool)`

GetEnableTargetResourceValidationForSLOOk returns a tuple with the EnableTargetResourceValidationForSLO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTargetResourceValidationForSLO

`func (o *RedirectValidationLocalSettings) SetEnableTargetResourceValidationForSLO(v bool)`

SetEnableTargetResourceValidationForSLO sets EnableTargetResourceValidationForSLO field to given value.

### HasEnableTargetResourceValidationForSLO

`func (o *RedirectValidationLocalSettings) HasEnableTargetResourceValidationForSLO() bool`

HasEnableTargetResourceValidationForSLO returns a boolean if a field has been set.

### GetEnableTargetResourceValidationForIdpDiscovery

`func (o *RedirectValidationLocalSettings) GetEnableTargetResourceValidationForIdpDiscovery() bool`

GetEnableTargetResourceValidationForIdpDiscovery returns the EnableTargetResourceValidationForIdpDiscovery field if non-nil, zero value otherwise.

### GetEnableTargetResourceValidationForIdpDiscoveryOk

`func (o *RedirectValidationLocalSettings) GetEnableTargetResourceValidationForIdpDiscoveryOk() (*bool, bool)`

GetEnableTargetResourceValidationForIdpDiscoveryOk returns a tuple with the EnableTargetResourceValidationForIdpDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTargetResourceValidationForIdpDiscovery

`func (o *RedirectValidationLocalSettings) SetEnableTargetResourceValidationForIdpDiscovery(v bool)`

SetEnableTargetResourceValidationForIdpDiscovery sets EnableTargetResourceValidationForIdpDiscovery field to given value.

### HasEnableTargetResourceValidationForIdpDiscovery

`func (o *RedirectValidationLocalSettings) HasEnableTargetResourceValidationForIdpDiscovery() bool`

HasEnableTargetResourceValidationForIdpDiscovery returns a boolean if a field has been set.

### GetEnableInErrorResourceValidation

`func (o *RedirectValidationLocalSettings) GetEnableInErrorResourceValidation() bool`

GetEnableInErrorResourceValidation returns the EnableInErrorResourceValidation field if non-nil, zero value otherwise.

### GetEnableInErrorResourceValidationOk

`func (o *RedirectValidationLocalSettings) GetEnableInErrorResourceValidationOk() (*bool, bool)`

GetEnableInErrorResourceValidationOk returns a tuple with the EnableInErrorResourceValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInErrorResourceValidation

`func (o *RedirectValidationLocalSettings) SetEnableInErrorResourceValidation(v bool)`

SetEnableInErrorResourceValidation sets EnableInErrorResourceValidation field to given value.

### HasEnableInErrorResourceValidation

`func (o *RedirectValidationLocalSettings) HasEnableInErrorResourceValidation() bool`

HasEnableInErrorResourceValidation returns a boolean if a field has been set.

### GetWhiteList

`func (o *RedirectValidationLocalSettings) GetWhiteList() []RedirectValidationSettingsWhitelistEntry`

GetWhiteList returns the WhiteList field if non-nil, zero value otherwise.

### GetWhiteListOk

`func (o *RedirectValidationLocalSettings) GetWhiteListOk() (*[]RedirectValidationSettingsWhitelistEntry, bool)`

GetWhiteListOk returns a tuple with the WhiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhiteList

`func (o *RedirectValidationLocalSettings) SetWhiteList(v []RedirectValidationSettingsWhitelistEntry)`

SetWhiteList sets WhiteList field to given value.

### HasWhiteList

`func (o *RedirectValidationLocalSettings) HasWhiteList() bool`

HasWhiteList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


