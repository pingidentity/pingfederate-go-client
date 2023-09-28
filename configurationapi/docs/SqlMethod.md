# SqlMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Table** | Pointer to [**Table**](Table.md) |  | [optional] 
**StoredProcedure** | Pointer to [**StoredProcedure**](StoredProcedure.md) |  | [optional] 

## Methods

### NewSqlMethod

`func NewSqlMethod() *SqlMethod`

NewSqlMethod instantiates a new SqlMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlMethodWithDefaults

`func NewSqlMethodWithDefaults() *SqlMethod`

NewSqlMethodWithDefaults instantiates a new SqlMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTable

`func (o *SqlMethod) GetTable() Table`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *SqlMethod) GetTableOk() (*Table, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *SqlMethod) SetTable(v Table)`

SetTable sets Table field to given value.

### HasTable

`func (o *SqlMethod) HasTable() bool`

HasTable returns a boolean if a field has been set.

### GetStoredProcedure

`func (o *SqlMethod) GetStoredProcedure() StoredProcedure`

GetStoredProcedure returns the StoredProcedure field if non-nil, zero value otherwise.

### GetStoredProcedureOk

`func (o *SqlMethod) GetStoredProcedureOk() (*StoredProcedure, bool)`

GetStoredProcedureOk returns a tuple with the StoredProcedure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredProcedure

`func (o *SqlMethod) SetStoredProcedure(v StoredProcedure)`

SetStoredProcedure sets StoredProcedure field to given value.

### HasStoredProcedure

`func (o *SqlMethod) HasStoredProcedure() bool`

HasStoredProcedure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


