# AttributeSourceAggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The data store type of this attribute source. | 
**DataStoreRef** | [**ResourceLink**](ResourceLink.md) |  | 
**Id** | Pointer to **string** | The ID that defines this attribute source. Only alphanumeric characters allowed.&lt;br&gt;Note: Required for OpenID Connect policy attribute sources, OAuth IdP adapter mappings, OAuth access token mappings and APC-to-SP Adapter Mappings. IdP Connections will ignore this property since it only allows one attribute source to be defined per mapping. IdP-to-SP Adapter Mappings can contain multiple attribute sources. | [optional] 
**Description** | Pointer to **string** | The description of this attribute source. The description needs to be unique amongst the attribute sources for the mapping.&lt;br&gt;Note: Required for APC-to-SP Adapter Mappings | [optional] 
**AttributeContractFulfillment** | Pointer to [**map[string]AttributeFulfillmentValue**](AttributeFulfillmentValue.md) | A list of mappings from attribute names to their fulfillment values. This field is only valid for the SP Connection&#39;s Browser SSO mappings | [optional] 
**Schema** | Pointer to **string** | Lists the table structure that stores information within a database. Some databases, such as Oracle, require a schema for a JDBC query. Other databases, such as MySQL, do not require a schema. | [optional] 
**Table** | **string** | The name of the database table. The name is used to construct the SQL query to retrieve data from the data store. | 
**ColumnNames** | Pointer to **[]string** | A list of column names used to construct the SQL query to retrieve data from the specified table in the datastore. | [optional] 
**Filter** | **string** | The JDBC WHERE clause used to query your data store to locate a user record. | 
**BaseDn** | Pointer to **string** | The base DN to search from. If not specified, the search will start at the LDAP&#39;s root. | [optional] 
**SearchScope** | **string** | Determines the node depth of the query. | 
**SearchFilter** | **string** | The LDAP filter that will be used to lookup the objects from the directory. | 
**SearchAttributes** | Pointer to **[]string** | A list of LDAP attributes returned from search and available for mapping. | [optional] 
**BinaryAttributeSettings** | Pointer to [**map[string]BinaryLdapAttributeSettings**](BinaryLdapAttributeSettings.md) | The advanced settings for binary LDAP attributes. | [optional] 
**MemberOfNestedGroup** | Pointer to **bool** | Set this to true to return transitive group memberships for the &#39;memberOf&#39; attribute.  This only applies for Active Directory data sources.  All other data sources will be set to false. | [optional] 

## Methods

### NewAttributeSourceAggregation

`func NewAttributeSourceAggregation(type_ string, dataStoreRef ResourceLink, table string, filter string, searchScope string, searchFilter string, ) *AttributeSourceAggregation`

NewAttributeSourceAggregation instantiates a new AttributeSourceAggregation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeSourceAggregationWithDefaults

`func NewAttributeSourceAggregationWithDefaults() *AttributeSourceAggregation`

NewAttributeSourceAggregationWithDefaults instantiates a new AttributeSourceAggregation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AttributeSourceAggregation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AttributeSourceAggregation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AttributeSourceAggregation) SetType(v string)`

SetType sets Type field to given value.


### GetDataStoreRef

`func (o *AttributeSourceAggregation) GetDataStoreRef() ResourceLink`

GetDataStoreRef returns the DataStoreRef field if non-nil, zero value otherwise.

### GetDataStoreRefOk

`func (o *AttributeSourceAggregation) GetDataStoreRefOk() (*ResourceLink, bool)`

GetDataStoreRefOk returns a tuple with the DataStoreRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStoreRef

`func (o *AttributeSourceAggregation) SetDataStoreRef(v ResourceLink)`

SetDataStoreRef sets DataStoreRef field to given value.


### GetId

`func (o *AttributeSourceAggregation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttributeSourceAggregation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttributeSourceAggregation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AttributeSourceAggregation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *AttributeSourceAggregation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AttributeSourceAggregation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AttributeSourceAggregation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AttributeSourceAggregation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAttributeContractFulfillment

`func (o *AttributeSourceAggregation) GetAttributeContractFulfillment() map[string]AttributeFulfillmentValue`

GetAttributeContractFulfillment returns the AttributeContractFulfillment field if non-nil, zero value otherwise.

### GetAttributeContractFulfillmentOk

`func (o *AttributeSourceAggregation) GetAttributeContractFulfillmentOk() (*map[string]AttributeFulfillmentValue, bool)`

GetAttributeContractFulfillmentOk returns a tuple with the AttributeContractFulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContractFulfillment

`func (o *AttributeSourceAggregation) SetAttributeContractFulfillment(v map[string]AttributeFulfillmentValue)`

SetAttributeContractFulfillment sets AttributeContractFulfillment field to given value.

### HasAttributeContractFulfillment

`func (o *AttributeSourceAggregation) HasAttributeContractFulfillment() bool`

HasAttributeContractFulfillment returns a boolean if a field has been set.

### GetSchema

`func (o *AttributeSourceAggregation) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *AttributeSourceAggregation) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *AttributeSourceAggregation) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *AttributeSourceAggregation) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetTable

`func (o *AttributeSourceAggregation) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *AttributeSourceAggregation) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *AttributeSourceAggregation) SetTable(v string)`

SetTable sets Table field to given value.


### GetColumnNames

`func (o *AttributeSourceAggregation) GetColumnNames() []string`

GetColumnNames returns the ColumnNames field if non-nil, zero value otherwise.

### GetColumnNamesOk

`func (o *AttributeSourceAggregation) GetColumnNamesOk() (*[]string, bool)`

GetColumnNamesOk returns a tuple with the ColumnNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnNames

`func (o *AttributeSourceAggregation) SetColumnNames(v []string)`

SetColumnNames sets ColumnNames field to given value.

### HasColumnNames

`func (o *AttributeSourceAggregation) HasColumnNames() bool`

HasColumnNames returns a boolean if a field has been set.

### GetFilter

`func (o *AttributeSourceAggregation) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AttributeSourceAggregation) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AttributeSourceAggregation) SetFilter(v string)`

SetFilter sets Filter field to given value.


### GetBaseDn

`func (o *AttributeSourceAggregation) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *AttributeSourceAggregation) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *AttributeSourceAggregation) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.

### HasBaseDn

`func (o *AttributeSourceAggregation) HasBaseDn() bool`

HasBaseDn returns a boolean if a field has been set.

### GetSearchScope

`func (o *AttributeSourceAggregation) GetSearchScope() string`

GetSearchScope returns the SearchScope field if non-nil, zero value otherwise.

### GetSearchScopeOk

`func (o *AttributeSourceAggregation) GetSearchScopeOk() (*string, bool)`

GetSearchScopeOk returns a tuple with the SearchScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchScope

`func (o *AttributeSourceAggregation) SetSearchScope(v string)`

SetSearchScope sets SearchScope field to given value.


### GetSearchFilter

`func (o *AttributeSourceAggregation) GetSearchFilter() string`

GetSearchFilter returns the SearchFilter field if non-nil, zero value otherwise.

### GetSearchFilterOk

`func (o *AttributeSourceAggregation) GetSearchFilterOk() (*string, bool)`

GetSearchFilterOk returns a tuple with the SearchFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilter

`func (o *AttributeSourceAggregation) SetSearchFilter(v string)`

SetSearchFilter sets SearchFilter field to given value.


### GetSearchAttributes

`func (o *AttributeSourceAggregation) GetSearchAttributes() []string`

GetSearchAttributes returns the SearchAttributes field if non-nil, zero value otherwise.

### GetSearchAttributesOk

`func (o *AttributeSourceAggregation) GetSearchAttributesOk() (*[]string, bool)`

GetSearchAttributesOk returns a tuple with the SearchAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAttributes

`func (o *AttributeSourceAggregation) SetSearchAttributes(v []string)`

SetSearchAttributes sets SearchAttributes field to given value.

### HasSearchAttributes

`func (o *AttributeSourceAggregation) HasSearchAttributes() bool`

HasSearchAttributes returns a boolean if a field has been set.

### GetBinaryAttributeSettings

`func (o *AttributeSourceAggregation) GetBinaryAttributeSettings() map[string]BinaryLdapAttributeSettings`

GetBinaryAttributeSettings returns the BinaryAttributeSettings field if non-nil, zero value otherwise.

### GetBinaryAttributeSettingsOk

`func (o *AttributeSourceAggregation) GetBinaryAttributeSettingsOk() (*map[string]BinaryLdapAttributeSettings, bool)`

GetBinaryAttributeSettingsOk returns a tuple with the BinaryAttributeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryAttributeSettings

`func (o *AttributeSourceAggregation) SetBinaryAttributeSettings(v map[string]BinaryLdapAttributeSettings)`

SetBinaryAttributeSettings sets BinaryAttributeSettings field to given value.

### HasBinaryAttributeSettings

`func (o *AttributeSourceAggregation) HasBinaryAttributeSettings() bool`

HasBinaryAttributeSettings returns a boolean if a field has been set.

### GetMemberOfNestedGroup

`func (o *AttributeSourceAggregation) GetMemberOfNestedGroup() bool`

GetMemberOfNestedGroup returns the MemberOfNestedGroup field if non-nil, zero value otherwise.

### GetMemberOfNestedGroupOk

`func (o *AttributeSourceAggregation) GetMemberOfNestedGroupOk() (*bool, bool)`

GetMemberOfNestedGroupOk returns a tuple with the MemberOfNestedGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOfNestedGroup

`func (o *AttributeSourceAggregation) SetMemberOfNestedGroup(v bool)`

SetMemberOfNestedGroup sets MemberOfNestedGroup field to given value.

### HasMemberOfNestedGroup

`func (o *AttributeSourceAggregation) HasMemberOfNestedGroup() bool`

HasMemberOfNestedGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


