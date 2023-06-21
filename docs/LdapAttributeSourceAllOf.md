# LdapAttributeSourceAllOf

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

### NewLdapAttributeSourceAllOf

`func NewLdapAttributeSourceAllOf(searchScope string, searchFilter string, ) *LdapAttributeSourceAllOf`

NewLdapAttributeSourceAllOf instantiates a new LdapAttributeSourceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapAttributeSourceAllOfWithDefaults

`func NewLdapAttributeSourceAllOfWithDefaults() *LdapAttributeSourceAllOf`

NewLdapAttributeSourceAllOfWithDefaults instantiates a new LdapAttributeSourceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseDn

`func (o *LdapAttributeSourceAllOf) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *LdapAttributeSourceAllOf) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *LdapAttributeSourceAllOf) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.

### HasBaseDn

`func (o *LdapAttributeSourceAllOf) HasBaseDn() bool`

HasBaseDn returns a boolean if a field has been set.

### GetSearchScope

`func (o *LdapAttributeSourceAllOf) GetSearchScope() string`

GetSearchScope returns the SearchScope field if non-nil, zero value otherwise.

### GetSearchScopeOk

`func (o *LdapAttributeSourceAllOf) GetSearchScopeOk() (*string, bool)`

GetSearchScopeOk returns a tuple with the SearchScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchScope

`func (o *LdapAttributeSourceAllOf) SetSearchScope(v string)`

SetSearchScope sets SearchScope field to given value.


### GetSearchFilter

`func (o *LdapAttributeSourceAllOf) GetSearchFilter() string`

GetSearchFilter returns the SearchFilter field if non-nil, zero value otherwise.

### GetSearchFilterOk

`func (o *LdapAttributeSourceAllOf) GetSearchFilterOk() (*string, bool)`

GetSearchFilterOk returns a tuple with the SearchFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilter

`func (o *LdapAttributeSourceAllOf) SetSearchFilter(v string)`

SetSearchFilter sets SearchFilter field to given value.


### GetSearchAttributes

`func (o *LdapAttributeSourceAllOf) GetSearchAttributes() []string`

GetSearchAttributes returns the SearchAttributes field if non-nil, zero value otherwise.

### GetSearchAttributesOk

`func (o *LdapAttributeSourceAllOf) GetSearchAttributesOk() (*[]string, bool)`

GetSearchAttributesOk returns a tuple with the SearchAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAttributes

`func (o *LdapAttributeSourceAllOf) SetSearchAttributes(v []string)`

SetSearchAttributes sets SearchAttributes field to given value.

### HasSearchAttributes

`func (o *LdapAttributeSourceAllOf) HasSearchAttributes() bool`

HasSearchAttributes returns a boolean if a field has been set.

### GetBinaryAttributeSettings

`func (o *LdapAttributeSourceAllOf) GetBinaryAttributeSettings() map[string]BinaryLdapAttributeSettings`

GetBinaryAttributeSettings returns the BinaryAttributeSettings field if non-nil, zero value otherwise.

### GetBinaryAttributeSettingsOk

`func (o *LdapAttributeSourceAllOf) GetBinaryAttributeSettingsOk() (*map[string]BinaryLdapAttributeSettings, bool)`

GetBinaryAttributeSettingsOk returns a tuple with the BinaryAttributeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryAttributeSettings

`func (o *LdapAttributeSourceAllOf) SetBinaryAttributeSettings(v map[string]BinaryLdapAttributeSettings)`

SetBinaryAttributeSettings sets BinaryAttributeSettings field to given value.

### HasBinaryAttributeSettings

`func (o *LdapAttributeSourceAllOf) HasBinaryAttributeSettings() bool`

HasBinaryAttributeSettings returns a boolean if a field has been set.

### GetMemberOfNestedGroup

`func (o *LdapAttributeSourceAllOf) GetMemberOfNestedGroup() bool`

GetMemberOfNestedGroup returns the MemberOfNestedGroup field if non-nil, zero value otherwise.

### GetMemberOfNestedGroupOk

`func (o *LdapAttributeSourceAllOf) GetMemberOfNestedGroupOk() (*bool, bool)`

GetMemberOfNestedGroupOk returns a tuple with the MemberOfNestedGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOfNestedGroup

`func (o *LdapAttributeSourceAllOf) SetMemberOfNestedGroup(v bool)`

SetMemberOfNestedGroup sets MemberOfNestedGroup field to given value.

### HasMemberOfNestedGroup

`func (o *LdapAttributeSourceAllOf) HasMemberOfNestedGroup() bool`

HasMemberOfNestedGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


