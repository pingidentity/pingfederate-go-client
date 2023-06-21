# StoredProcedure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | **string** | Lists the table structure that stores information within a database. | 
**StoredProcedure** | **string** | The name of the database stored procedure. | 

## Methods

### NewStoredProcedure

`func NewStoredProcedure(schema string, storedProcedure string, ) *StoredProcedure`

NewStoredProcedure instantiates a new StoredProcedure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoredProcedureWithDefaults

`func NewStoredProcedureWithDefaults() *StoredProcedure`

NewStoredProcedureWithDefaults instantiates a new StoredProcedure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *StoredProcedure) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *StoredProcedure) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *StoredProcedure) SetSchema(v string)`

SetSchema sets Schema field to given value.


### GetStoredProcedure

`func (o *StoredProcedure) GetStoredProcedure() string`

GetStoredProcedure returns the StoredProcedure field if non-nil, zero value otherwise.

### GetStoredProcedureOk

`func (o *StoredProcedure) GetStoredProcedureOk() (*string, bool)`

GetStoredProcedureOk returns a tuple with the StoredProcedure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoredProcedure

`func (o *StoredProcedure) SetStoredProcedure(v string)`

SetStoredProcedure sets StoredProcedure field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


