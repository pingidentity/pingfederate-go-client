# RedirectValidationSettingsWhitelistEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetResourceSSO** | Pointer to **bool** | Enable this target resource for SSO redirect validation. | [optional] 
**TargetResourceSLO** | Pointer to **bool** | Enable this target resource for SLO redirect validation. | [optional] 
**InErrorResource** | Pointer to **bool** | Enable this target resource for in error resource validation. | [optional] 
**IdpDiscovery** | Pointer to **bool** | Enable this target resource for IdP discovery validation. | [optional] 
**ValidDomain** | **string** | Domain of a valid resource. | 
**ValidPath** | Pointer to **string** | Path of a valid resource. | [optional] 
**AllowQueryAndFragment** | Pointer to **bool** | Allow any query parameters and fragment in the resource. | [optional] 
**RequireHttps** | Pointer to **bool** | Require HTTPS for accessing this resource. | [optional] 

## Methods

### NewRedirectValidationSettingsWhitelistEntry

`func NewRedirectValidationSettingsWhitelistEntry(validDomain string, ) *RedirectValidationSettingsWhitelistEntry`

NewRedirectValidationSettingsWhitelistEntry instantiates a new RedirectValidationSettingsWhitelistEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedirectValidationSettingsWhitelistEntryWithDefaults

`func NewRedirectValidationSettingsWhitelistEntryWithDefaults() *RedirectValidationSettingsWhitelistEntry`

NewRedirectValidationSettingsWhitelistEntryWithDefaults instantiates a new RedirectValidationSettingsWhitelistEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetResourceSSO

`func (o *RedirectValidationSettingsWhitelistEntry) GetTargetResourceSSO() bool`

GetTargetResourceSSO returns the TargetResourceSSO field if non-nil, zero value otherwise.

### GetTargetResourceSSOOk

`func (o *RedirectValidationSettingsWhitelistEntry) GetTargetResourceSSOOk() (*bool, bool)`

GetTargetResourceSSOOk returns a tuple with the TargetResourceSSO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResourceSSO

`func (o *RedirectValidationSettingsWhitelistEntry) SetTargetResourceSSO(v bool)`

SetTargetResourceSSO sets TargetResourceSSO field to given value.

### HasTargetResourceSSO

`func (o *RedirectValidationSettingsWhitelistEntry) HasTargetResourceSSO() bool`

HasTargetResourceSSO returns a boolean if a field has been set.

### GetTargetResourceSLO

`func (o *RedirectValidationSettingsWhitelistEntry) GetTargetResourceSLO() bool`

GetTargetResourceSLO returns the TargetResourceSLO field if non-nil, zero value otherwise.

### GetTargetResourceSLOOk

`func (o *RedirectValidationSettingsWhitelistEntry) GetTargetResourceSLOOk() (*bool, bool)`

GetTargetResourceSLOOk returns a tuple with the TargetResourceSLO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResourceSLO

`func (o *RedirectValidationSettingsWhitelistEntry) SetTargetResourceSLO(v bool)`

SetTargetResourceSLO sets TargetResourceSLO field to given value.

### HasTargetResourceSLO

`func (o *RedirectValidationSettingsWhitelistEntry) HasTargetResourceSLO() bool`

HasTargetResourceSLO returns a boolean if a field has been set.

### GetInErrorResource

`func (o *RedirectValidationSettingsWhitelistEntry) GetInErrorResource() bool`

GetInErrorResource returns the InErrorResource field if non-nil, zero value otherwise.

### GetInErrorResourceOk

`func (o *RedirectValidationSettingsWhitelistEntry) GetInErrorResourceOk() (*bool, bool)`

GetInErrorResourceOk returns a tuple with the InErrorResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInErrorResource

`func (o *RedirectValidationSettingsWhitelistEntry) SetInErrorResource(v bool)`

SetInErrorResource sets InErrorResource field to given value.

### HasInErrorResource

`func (o *RedirectValidationSettingsWhitelistEntry) HasInErrorResource() bool`

HasInErrorResource returns a boolean if a field has been set.

### GetIdpDiscovery

`func (o *RedirectValidationSettingsWhitelistEntry) GetIdpDiscovery() bool`

GetIdpDiscovery returns the IdpDiscovery field if non-nil, zero value otherwise.

### GetIdpDiscoveryOk

`func (o *RedirectValidationSettingsWhitelistEntry) GetIdpDiscoveryOk() (*bool, bool)`

GetIdpDiscoveryOk returns a tuple with the IdpDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpDiscovery

`func (o *RedirectValidationSettingsWhitelistEntry) SetIdpDiscovery(v bool)`

SetIdpDiscovery sets IdpDiscovery field to given value.

### HasIdpDiscovery

`func (o *RedirectValidationSettingsWhitelistEntry) HasIdpDiscovery() bool`

HasIdpDiscovery returns a boolean if a field has been set.

### GetValidDomain

`func (o *RedirectValidationSettingsWhitelistEntry) GetValidDomain() string`

GetValidDomain returns the ValidDomain field if non-nil, zero value otherwise.

### GetValidDomainOk

`func (o *RedirectValidationSettingsWhitelistEntry) GetValidDomainOk() (*string, bool)`

GetValidDomainOk returns a tuple with the ValidDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidDomain

`func (o *RedirectValidationSettingsWhitelistEntry) SetValidDomain(v string)`

SetValidDomain sets ValidDomain field to given value.


### GetValidPath

`func (o *RedirectValidationSettingsWhitelistEntry) GetValidPath() string`

GetValidPath returns the ValidPath field if non-nil, zero value otherwise.

### GetValidPathOk

`func (o *RedirectValidationSettingsWhitelistEntry) GetValidPathOk() (*string, bool)`

GetValidPathOk returns a tuple with the ValidPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidPath

`func (o *RedirectValidationSettingsWhitelistEntry) SetValidPath(v string)`

SetValidPath sets ValidPath field to given value.

### HasValidPath

`func (o *RedirectValidationSettingsWhitelistEntry) HasValidPath() bool`

HasValidPath returns a boolean if a field has been set.

### GetAllowQueryAndFragment

`func (o *RedirectValidationSettingsWhitelistEntry) GetAllowQueryAndFragment() bool`

GetAllowQueryAndFragment returns the AllowQueryAndFragment field if non-nil, zero value otherwise.

### GetAllowQueryAndFragmentOk

`func (o *RedirectValidationSettingsWhitelistEntry) GetAllowQueryAndFragmentOk() (*bool, bool)`

GetAllowQueryAndFragmentOk returns a tuple with the AllowQueryAndFragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowQueryAndFragment

`func (o *RedirectValidationSettingsWhitelistEntry) SetAllowQueryAndFragment(v bool)`

SetAllowQueryAndFragment sets AllowQueryAndFragment field to given value.

### HasAllowQueryAndFragment

`func (o *RedirectValidationSettingsWhitelistEntry) HasAllowQueryAndFragment() bool`

HasAllowQueryAndFragment returns a boolean if a field has been set.

### GetRequireHttps

`func (o *RedirectValidationSettingsWhitelistEntry) GetRequireHttps() bool`

GetRequireHttps returns the RequireHttps field if non-nil, zero value otherwise.

### GetRequireHttpsOk

`func (o *RedirectValidationSettingsWhitelistEntry) GetRequireHttpsOk() (*bool, bool)`

GetRequireHttpsOk returns a tuple with the RequireHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireHttps

`func (o *RedirectValidationSettingsWhitelistEntry) SetRequireHttps(v bool)`

SetRequireHttps sets RequireHttps field to given value.

### HasRequireHttps

`func (o *RedirectValidationSettingsWhitelistEntry) HasRequireHttps() bool`

HasRequireHttps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


