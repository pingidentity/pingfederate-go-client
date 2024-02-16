# JdbcAttributeSource

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

## Methods

### NewJdbcAttributeSource

`func NewJdbcAttributeSource(type_ string, dataStoreRef ResourceLink, table string, filter string, ) *JdbcAttributeSource`

NewJdbcAttributeSource instantiates a new JdbcAttributeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJdbcAttributeSourceWithDefaults

`func NewJdbcAttributeSourceWithDefaults() *JdbcAttributeSource`

NewJdbcAttributeSourceWithDefaults instantiates a new JdbcAttributeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *JdbcAttributeSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JdbcAttributeSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JdbcAttributeSource) SetType(v string)`

SetType sets Type field to given value.


### GetDataStoreRef

`func (o *JdbcAttributeSource) GetDataStoreRef() ResourceLink`

GetDataStoreRef returns the DataStoreRef field if non-nil, zero value otherwise.

### GetDataStoreRefOk

`func (o *JdbcAttributeSource) GetDataStoreRefOk() (*ResourceLink, bool)`

GetDataStoreRefOk returns a tuple with the DataStoreRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStoreRef

`func (o *JdbcAttributeSource) SetDataStoreRef(v ResourceLink)`

SetDataStoreRef sets DataStoreRef field to given value.


### GetId

`func (o *JdbcAttributeSource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JdbcAttributeSource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JdbcAttributeSource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JdbcAttributeSource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *JdbcAttributeSource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JdbcAttributeSource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JdbcAttributeSource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JdbcAttributeSource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAttributeContractFulfillment

`func (o *JdbcAttributeSource) GetAttributeContractFulfillment() map[string]AttributeFulfillmentValue`

GetAttributeContractFulfillment returns the AttributeContractFulfillment field if non-nil, zero value otherwise.

### GetAttributeContractFulfillmentOk

`func (o *JdbcAttributeSource) GetAttributeContractFulfillmentOk() (*map[string]AttributeFulfillmentValue, bool)`

GetAttributeContractFulfillmentOk returns a tuple with the AttributeContractFulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContractFulfillment

`func (o *JdbcAttributeSource) SetAttributeContractFulfillment(v map[string]AttributeFulfillmentValue)`

SetAttributeContractFulfillment sets AttributeContractFulfillment field to given value.

### HasAttributeContractFulfillment

`func (o *JdbcAttributeSource) HasAttributeContractFulfillment() bool`

HasAttributeContractFulfillment returns a boolean if a field has been set.

### GetSchema

`func (o *JdbcAttributeSource) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *JdbcAttributeSource) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *JdbcAttributeSource) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *JdbcAttributeSource) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetTable

`func (o *JdbcAttributeSource) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *JdbcAttributeSource) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *JdbcAttributeSource) SetTable(v string)`

SetTable sets Table field to given value.


### GetColumnNames

`func (o *JdbcAttributeSource) GetColumnNames() []string`

GetColumnNames returns the ColumnNames field if non-nil, zero value otherwise.

### GetColumnNamesOk

`func (o *JdbcAttributeSource) GetColumnNamesOk() (*[]string, bool)`

GetColumnNamesOk returns a tuple with the ColumnNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnNames

`func (o *JdbcAttributeSource) SetColumnNames(v []string)`

SetColumnNames sets ColumnNames field to given value.

### HasColumnNames

`func (o *JdbcAttributeSource) HasColumnNames() bool`

HasColumnNames returns a boolean if a field has been set.

### GetFilter

`func (o *JdbcAttributeSource) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *JdbcAttributeSource) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *JdbcAttributeSource) SetFilter(v string)`

SetFilter sets Filter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


