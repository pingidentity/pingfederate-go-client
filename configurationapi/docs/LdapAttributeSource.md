# LdapAttributeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The data store type of this attribute source. | 
**DataStoreRef** | [**ResourceLink**](ResourceLink.md) |  | 
**Id** | Pointer to **string** | The ID that defines this attribute source. Only alphanumeric characters allowed.&lt;br&gt;Note: Required for OpenID Connect policy attribute sources, OAuth IdP adapter mappings, OAuth access token mappings and APC-to-SP Adapter Mappings. IdP Connections will ignore this property since it only allows one attribute source to be defined per mapping. IdP-to-SP Adapter Mappings can contain multiple attribute sources. | [optional] 
**Description** | Pointer to **string** | The description of this attribute source. The description needs to be unique amongst the attribute sources for the mapping.&lt;br&gt;Note: Required for APC-to-SP Adapter Mappings | [optional] 
**AttributeContractFulfillment** | Pointer to [**map[string]AttributeFulfillmentValue**](AttributeFulfillmentValue.md) | A list of mappings from attribute names to their fulfillment values. This field is only valid for the SP Connection&#39;s Browser SSO mappings | [optional] 
**BaseDn** | Pointer to **string** | The base DN to search from. If not specified, the search will start at the LDAP&#39;s root. | [optional] 
**SearchScope** | **string** | Determines the node depth of the query. | 
**SearchFilter** | **string** | The LDAP filter that will be used to lookup the objects from the directory. | 
**SearchAttributes** | Pointer to **[]string** | A list of LDAP attributes returned from search and available for mapping. | [optional] 
**BinaryAttributeSettings** | Pointer to [**map[string]BinaryLdapAttributeSettings**](BinaryLdapAttributeSettings.md) | The advanced settings for binary LDAP attributes. | [optional] 
**MemberOfNestedGroup** | Pointer to **bool** | Set this to true to return transitive group memberships for the &#39;memberOf&#39; attribute.  This only applies for Active Directory data sources.  All other data sources will be set to false. | [optional] 

## Methods

### NewLdapAttributeSource

`func NewLdapAttributeSource(type_ string, dataStoreRef ResourceLink, searchScope string, searchFilter string, ) *LdapAttributeSource`

NewLdapAttributeSource instantiates a new LdapAttributeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapAttributeSourceWithDefaults

`func NewLdapAttributeSourceWithDefaults() *LdapAttributeSource`

NewLdapAttributeSourceWithDefaults instantiates a new LdapAttributeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LdapAttributeSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LdapAttributeSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LdapAttributeSource) SetType(v string)`

SetType sets Type field to given value.


### GetDataStoreRef

`func (o *LdapAttributeSource) GetDataStoreRef() ResourceLink`

GetDataStoreRef returns the DataStoreRef field if non-nil, zero value otherwise.

### GetDataStoreRefOk

`func (o *LdapAttributeSource) GetDataStoreRefOk() (*ResourceLink, bool)`

GetDataStoreRefOk returns a tuple with the DataStoreRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStoreRef

`func (o *LdapAttributeSource) SetDataStoreRef(v ResourceLink)`

SetDataStoreRef sets DataStoreRef field to given value.


### GetId

`func (o *LdapAttributeSource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LdapAttributeSource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LdapAttributeSource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LdapAttributeSource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *LdapAttributeSource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LdapAttributeSource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LdapAttributeSource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LdapAttributeSource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAttributeContractFulfillment

`func (o *LdapAttributeSource) GetAttributeContractFulfillment() map[string]AttributeFulfillmentValue`

GetAttributeContractFulfillment returns the AttributeContractFulfillment field if non-nil, zero value otherwise.

### GetAttributeContractFulfillmentOk

`func (o *LdapAttributeSource) GetAttributeContractFulfillmentOk() (*map[string]AttributeFulfillmentValue, bool)`

GetAttributeContractFulfillmentOk returns a tuple with the AttributeContractFulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContractFulfillment

`func (o *LdapAttributeSource) SetAttributeContractFulfillment(v map[string]AttributeFulfillmentValue)`

SetAttributeContractFulfillment sets AttributeContractFulfillment field to given value.

### HasAttributeContractFulfillment

`func (o *LdapAttributeSource) HasAttributeContractFulfillment() bool`

HasAttributeContractFulfillment returns a boolean if a field has been set.

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


