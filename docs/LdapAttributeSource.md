# LdapAttributeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseDn** | Pointer to **string** | The base DN to search from. If not specified, the search will start at the LDAP&#39;s root. | [optional] 
**SearchScope** | **string** | Determines the node depth of the query. | 
**SearchFilter** | **string** | The LDAP filter that will be used to lookup the objects from the directory. | 
**SearchAttributes** | Pointer to **[]string** | A list of LDAP attributes returned from search and available for mapping. | [optional] 
**BinaryAttributeSettings** | Pointer to [**map[string]BinaryLdapAttributeSettings**](BinaryLdapAttributeSettings.md) | The advanced settings for binary LDAP attributes. | [optional] 
**MemberOfNestedGroup** | Pointer to **bool** | Set this to true to return transitive group memberships for the &#39;memberOf&#39; attribute.  This only applies for Active Directory data sources.  All other data sources will be set to false. | [optional] 

## Methods

### NewLdapAttributeSource

`func NewLdapAttributeSource(searchScope string, searchFilter string, ) *LdapAttributeSource`

NewLdapAttributeSource instantiates a new LdapAttributeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapAttributeSourceWithDefaults

`func NewLdapAttributeSourceWithDefaults() *LdapAttributeSource`

NewLdapAttributeSourceWithDefaults instantiates a new LdapAttributeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseDn

`func (o *LdapAttributeSource) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *LdapAttributeSource) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *LdapAttributeSource) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.

### HasBaseDn

`func (o *LdapAttributeSource) HasBaseDn() bool`

HasBaseDn returns a boolean if a field has been set.

### GetSearchScope

`func (o *LdapAttributeSource) GetSearchScope() string`

GetSearchScope returns the SearchScope field if non-nil, zero value otherwise.

### GetSearchScopeOk

`func (o *LdapAttributeSource) GetSearchScopeOk() (*string, bool)`

GetSearchScopeOk returns a tuple with the SearchScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchScope

`func (o *LdapAttributeSource) SetSearchScope(v string)`

SetSearchScope sets SearchScope field to given value.


### GetSearchFilter

`func (o *LdapAttributeSource) GetSearchFilter() string`

GetSearchFilter returns the SearchFilter field if non-nil, zero value otherwise.

### GetSearchFilterOk

`func (o *LdapAttributeSource) GetSearchFilterOk() (*string, bool)`

GetSearchFilterOk returns a tuple with the SearchFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilter

`func (o *LdapAttributeSource) SetSearchFilter(v string)`

SetSearchFilter sets SearchFilter field to given value.


### GetSearchAttributes

`func (o *LdapAttributeSource) GetSearchAttributes() []string`

GetSearchAttributes returns the SearchAttributes field if non-nil, zero value otherwise.

### GetSearchAttributesOk

`func (o *LdapAttributeSource) GetSearchAttributesOk() (*[]string, bool)`

GetSearchAttributesOk returns a tuple with the SearchAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAttributes

`func (o *LdapAttributeSource) SetSearchAttributes(v []string)`

SetSearchAttributes sets SearchAttributes field to given value.

### HasSearchAttributes

`func (o *LdapAttributeSource) HasSearchAttributes() bool`

HasSearchAttributes returns a boolean if a field has been set.

### GetBinaryAttributeSettings

`func (o *LdapAttributeSource) GetBinaryAttributeSettings() map[string]BinaryLdapAttributeSettings`

GetBinaryAttributeSettings returns the BinaryAttributeSettings field if non-nil, zero value otherwise.

### GetBinaryAttributeSettingsOk

`func (o *LdapAttributeSource) GetBinaryAttributeSettingsOk() (*map[string]BinaryLdapAttributeSettings, bool)`

GetBinaryAttributeSettingsOk returns a tuple with the BinaryAttributeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryAttributeSettings

`func (o *LdapAttributeSource) SetBinaryAttributeSettings(v map[string]BinaryLdapAttributeSettings)`

SetBinaryAttributeSettings sets BinaryAttributeSettings field to given value.

### HasBinaryAttributeSettings

`func (o *LdapAttributeSource) HasBinaryAttributeSettings() bool`

HasBinaryAttributeSettings returns a boolean if a field has been set.

### GetMemberOfNestedGroup

`func (o *LdapAttributeSource) GetMemberOfNestedGroup() bool`

GetMemberOfNestedGroup returns the MemberOfNestedGroup field if non-nil, zero value otherwise.

### GetMemberOfNestedGroupOk

`func (o *LdapAttributeSource) GetMemberOfNestedGroupOk() (*bool, bool)`

GetMemberOfNestedGroupOk returns a tuple with the MemberOfNestedGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOfNestedGroup

`func (o *LdapAttributeSource) SetMemberOfNestedGroup(v bool)`

SetMemberOfNestedGroup sets MemberOfNestedGroup field to given value.

### HasMemberOfNestedGroup

`func (o *LdapAttributeSource) HasMemberOfNestedGroup() bool`

HasMemberOfNestedGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


