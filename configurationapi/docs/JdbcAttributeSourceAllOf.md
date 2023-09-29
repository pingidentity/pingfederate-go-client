# JdbcAttributeSourceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | Lists the table structure that stores information within a database. Some databases, such as Oracle, require a schema for a JDBC query. Other databases, such as MySQL, do not require a schema. | [optional] 
**Table** | **string** | The name of the database table. The name is used to construct the SQL query to retrieve data from the data store. | 
**ColumnNames** | Pointer to **[]string** | A list of column names used to construct the SQL query to retrieve data from the specified table in the datastore. | [optional] 
**Filter** | **string** | The JDBC WHERE clause used to query your data store to locate a user record. | 

## Methods

### NewJdbcAttributeSourceAllOf

`func NewJdbcAttributeSourceAllOf(table string, filter string, ) *JdbcAttributeSourceAllOf`

NewJdbcAttributeSourceAllOf instantiates a new JdbcAttributeSourceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJdbcAttributeSourceAllOfWithDefaults

`func NewJdbcAttributeSourceAllOfWithDefaults() *JdbcAttributeSourceAllOf`

NewJdbcAttributeSourceAllOfWithDefaults instantiates a new JdbcAttributeSourceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *JdbcAttributeSourceAllOf) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *JdbcAttributeSourceAllOf) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *JdbcAttributeSourceAllOf) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *JdbcAttributeSourceAllOf) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetTable

`func (o *JdbcAttributeSourceAllOf) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *JdbcAttributeSourceAllOf) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *JdbcAttributeSourceAllOf) SetTable(v string)`

SetTable sets Table field to given value.


### GetColumnNames

`func (o *JdbcAttributeSourceAllOf) GetColumnNames() []string`

GetColumnNames returns the ColumnNames field if non-nil, zero value otherwise.

### GetColumnNamesOk

`func (o *JdbcAttributeSourceAllOf) GetColumnNamesOk() (*[]string, bool)`

GetColumnNamesOk returns a tuple with the ColumnNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnNames

`func (o *JdbcAttributeSourceAllOf) SetColumnNames(v []string)`

SetColumnNames sets ColumnNames field to given value.

### HasColumnNames

`func (o *JdbcAttributeSourceAllOf) HasColumnNames() bool`

HasColumnNames returns a boolean if a field has been set.

### GetFilter

`func (o *JdbcAttributeSourceAllOf) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *JdbcAttributeSourceAllOf) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *JdbcAttributeSourceAllOf) SetFilter(v string)`

SetFilter sets Filter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


