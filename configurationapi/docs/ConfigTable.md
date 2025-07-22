# ConfigTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the table. | 
**Rows** | Pointer to [**[]ConfigRow**](ConfigRow.md) | List of table rows. | [optional] 

## Methods

### NewConfigTable

`func NewConfigTable(name string, ) *ConfigTable`

NewConfigTable instantiates a new ConfigTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigTableWithDefaults

`func NewConfigTableWithDefaults() *ConfigTable`

NewConfigTableWithDefaults instantiates a new ConfigTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConfigTable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigTable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigTable) SetName(v string)`

SetName sets Name field to given value.


### GetRows

`func (o *ConfigTable) GetRows() []ConfigRow`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *ConfigTable) GetRowsOk() (*[]ConfigRow, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *ConfigTable) SetRows(v []ConfigRow)`

SetRows sets Rows field to given value.

### HasRows

`func (o *ConfigTable) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


