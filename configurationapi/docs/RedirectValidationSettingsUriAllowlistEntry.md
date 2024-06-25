# RedirectValidationSettingsUriAllowlistEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetResourceSSO** | Pointer to **bool** | Enable this target resource for SSO redirect validation. | [optional] 
**TargetResourceSLO** | Pointer to **bool** | Enable this target resource for SLO redirect validation. | [optional] 
**InErrorResource** | Pointer to **bool** | Enable this target resource for in error resource validation. | [optional] 
**IdpDiscovery** | Pointer to **bool** | Enable this target resource for IdP discovery validation. | [optional] 
**AllowQueryAndFragment** | Pointer to **bool** | Allow any query parameters and fragment in the resource. | [optional] 
**ValidUri** | **string** | URI of a valid resource. | 

## Methods

### NewRedirectValidationSettingsUriAllowlistEntry

`func NewRedirectValidationSettingsUriAllowlistEntry(validUri string, ) *RedirectValidationSettingsUriAllowlistEntry`

NewRedirectValidationSettingsUriAllowlistEntry instantiates a new RedirectValidationSettingsUriAllowlistEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedirectValidationSettingsUriAllowlistEntryWithDefaults

`func NewRedirectValidationSettingsUriAllowlistEntryWithDefaults() *RedirectValidationSettingsUriAllowlistEntry`

NewRedirectValidationSettingsUriAllowlistEntryWithDefaults instantiates a new RedirectValidationSettingsUriAllowlistEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetResourceSSO

`func (o *RedirectValidationSettingsUriAllowlistEntry) GetTargetResourceSSO() bool`

GetTargetResourceSSO returns the TargetResourceSSO field if non-nil, zero value otherwise.

### GetTargetResourceSSOOk

`func (o *RedirectValidationSettingsUriAllowlistEntry) GetTargetResourceSSOOk() (*bool, bool)`

GetTargetResourceSSOOk returns a tuple with the TargetResourceSSO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResourceSSO

`func (o *RedirectValidationSettingsUriAllowlistEntry) SetTargetResourceSSO(v bool)`

SetTargetResourceSSO sets TargetResourceSSO field to given value.

### HasTargetResourceSSO

`func (o *RedirectValidationSettingsUriAllowlistEntry) HasTargetResourceSSO() bool`

HasTargetResourceSSO returns a boolean if a field has been set.

### GetTargetResourceSLO

`func (o *RedirectValidationSettingsUriAllowlistEntry) GetTargetResourceSLO() bool`

GetTargetResourceSLO returns the TargetResourceSLO field if non-nil, zero value otherwise.

### GetTargetResourceSLOOk

`func (o *RedirectValidationSettingsUriAllowlistEntry) GetTargetResourceSLOOk() (*bool, bool)`

GetTargetResourceSLOOk returns a tuple with the TargetResourceSLO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResourceSLO

`func (o *RedirectValidationSettingsUriAllowlistEntry) SetTargetResourceSLO(v bool)`

SetTargetResourceSLO sets TargetResourceSLO field to given value.

### HasTargetResourceSLO

`func (o *RedirectValidationSettingsUriAllowlistEntry) HasTargetResourceSLO() bool`

HasTargetResourceSLO returns a boolean if a field has been set.

### GetInErrorResource

`func (o *RedirectValidationSettingsUriAllowlistEntry) GetInErrorResource() bool`

GetInErrorResource returns the InErrorResource field if non-nil, zero value otherwise.

### GetInErrorResourceOk

`func (o *RedirectValidationSettingsUriAllowlistEntry) GetInErrorResourceOk() (*bool, bool)`

GetInErrorResourceOk returns a tuple with the InErrorResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInErrorResource

`func (o *RedirectValidationSettingsUriAllowlistEntry) SetInErrorResource(v bool)`

SetInErrorResource sets InErrorResource field to given value.

### HasInErrorResource

`func (o *RedirectValidationSettingsUriAllowlistEntry) HasInErrorResource() bool`

HasInErrorResource returns a boolean if a field has been set.

### GetIdpDiscovery

`func (o *RedirectValidationSettingsUriAllowlistEntry) GetIdpDiscovery() bool`

GetIdpDiscovery returns the IdpDiscovery field if non-nil, zero value otherwise.

### GetIdpDiscoveryOk

`func (o *RedirectValidationSettingsUriAllowlistEntry) GetIdpDiscoveryOk() (*bool, bool)`

GetIdpDiscoveryOk returns a tuple with the IdpDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpDiscovery

`func (o *RedirectValidationSettingsUriAllowlistEntry) SetIdpDiscovery(v bool)`

SetIdpDiscovery sets IdpDiscovery field to given value.

### HasIdpDiscovery

`func (o *RedirectValidationSettingsUriAllowlistEntry) HasIdpDiscovery() bool`

HasIdpDiscovery returns a boolean if a field has been set.

### GetAllowQueryAndFragment

`func (o *RedirectValidationSettingsUriAllowlistEntry) GetAllowQueryAndFragment() bool`

GetAllowQueryAndFragment returns the AllowQueryAndFragment field if non-nil, zero value otherwise.

### GetAllowQueryAndFragmentOk

`func (o *RedirectValidationSettingsUriAllowlistEntry) GetAllowQueryAndFragmentOk() (*bool, bool)`

GetAllowQueryAndFragmentOk returns a tuple with the AllowQueryAndFragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowQueryAndFragment

`func (o *RedirectValidationSettingsUriAllowlistEntry) SetAllowQueryAndFragment(v bool)`

SetAllowQueryAndFragment sets AllowQueryAndFragment field to given value.

### HasAllowQueryAndFragment

`func (o *RedirectValidationSettingsUriAllowlistEntry) HasAllowQueryAndFragment() bool`

HasAllowQueryAndFragment returns a boolean if a field has been set.

### GetValidUri

`func (o *RedirectValidationSettingsUriAllowlistEntry) GetValidUri() string`

GetValidUri returns the ValidUri field if non-nil, zero value otherwise.

### GetValidUriOk

`func (o *RedirectValidationSettingsUriAllowlistEntry) GetValidUriOk() (*string, bool)`

GetValidUriOk returns a tuple with the ValidUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUri

`func (o *RedirectValidationSettingsUriAllowlistEntry) SetValidUri(v string)`

SetValidUri sets ValidUri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


