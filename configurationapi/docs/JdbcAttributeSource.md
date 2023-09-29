# JdbcAttributeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | Lists the table structure that stores information within a database. Some databases, such as Oracle, require a schema for a JDBC query. Other databases, such as MySQL, do not require a schema. | [optional] 
**Table** | **string** | The name of the database table. The name is used to construct the SQL query to retrieve data from the data store. | 
**ColumnNames** | Pointer to **[]string** | A list of column names used to construct the SQL query to retrieve data from the specified table in the datastore. | [optional] 
**Filter** | **string** | The JDBC WHERE clause used to query your data store to locate a user record. | 

## Methods

### NewJdbcAttributeSource

`func NewJdbcAttributeSource(table string, filter string, ) *JdbcAttributeSource`

NewJdbcAttributeSource instantiates a new JdbcAttributeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJdbcAttributeSourceWithDefaults

`func NewJdbcAttributeSourceWithDefaults() *JdbcAttributeSource`

NewJdbcAttributeSourceWithDefaults instantiates a new JdbcAttributeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


