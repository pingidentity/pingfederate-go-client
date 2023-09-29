# Table

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | **string** | Lists the table structure that stores information within a database. | 
**TableName** | **string** | The name of the database table. | 
**UniqueIdColumn** | **string** | The database column that uniquely identifies the provisioned user on the SP side. | 

## Methods

### NewTable

`func NewTable(schema string, tableName string, uniqueIdColumn string, ) *Table`

NewTable instantiates a new Table object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableWithDefaults

`func NewTableWithDefaults() *Table`

NewTableWithDefaults instantiates a new Table object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *Table) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *Table) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *Table) SetSchema(v string)`

SetSchema sets Schema field to given value.


### GetTableName

`func (o *Table) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *Table) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *Table) SetTableName(v string)`

SetTableName sets TableName field to given value.


### GetUniqueIdColumn

`func (o *Table) GetUniqueIdColumn() string`

GetUniqueIdColumn returns the UniqueIdColumn field if non-nil, zero value otherwise.

### GetUniqueIdColumnOk

`func (o *Table) GetUniqueIdColumnOk() (*string, bool)`

GetUniqueIdColumnOk returns a tuple with the UniqueIdColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdColumn

`func (o *Table) SetUniqueIdColumn(v string)`

SetUniqueIdColumn sets UniqueIdColumn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


