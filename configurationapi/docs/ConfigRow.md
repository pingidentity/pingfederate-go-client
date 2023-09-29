# ConfigRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | [**[]ConfigField**](ConfigField.md) | The configuration fields in the row. | 
**DefaultRow** | Pointer to **bool** | Whether this row is the default. | [optional] 

## Methods

### NewConfigRow

`func NewConfigRow(fields []ConfigField, ) *ConfigRow`

NewConfigRow instantiates a new ConfigRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigRowWithDefaults

`func NewConfigRowWithDefaults() *ConfigRow`

NewConfigRowWithDefaults instantiates a new ConfigRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *ConfigRow) GetFields() []ConfigField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ConfigRow) GetFieldsOk() (*[]ConfigField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ConfigRow) SetFields(v []ConfigField)`

SetFields sets Fields field to given value.


### GetDefaultRow

`func (o *ConfigRow) GetDefaultRow() bool`

GetDefaultRow returns the DefaultRow field if non-nil, zero value otherwise.

### GetDefaultRowOk

`func (o *ConfigRow) GetDefaultRowOk() (*bool, bool)`

GetDefaultRowOk returns a tuple with the DefaultRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRow

`func (o *ConfigRow) SetDefaultRow(v bool)`

SetDefaultRow sets DefaultRow field to given value.

### HasDefaultRow

`func (o *ConfigRow) HasDefaultRow() bool`

HasDefaultRow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


