# PluginConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tables** | Pointer to [**[]ConfigTable**](ConfigTable.md) | List of configuration tables. | [optional] 
**Fields** | Pointer to [**[]ConfigField**](ConfigField.md) | List of configuration fields. | [optional] 

## Methods

### NewPluginConfiguration

`func NewPluginConfiguration() *PluginConfiguration`

NewPluginConfiguration instantiates a new PluginConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginConfigurationWithDefaults

`func NewPluginConfigurationWithDefaults() *PluginConfiguration`

NewPluginConfigurationWithDefaults instantiates a new PluginConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTables

`func (o *PluginConfiguration) GetTables() []ConfigTable`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *PluginConfiguration) GetTablesOk() (*[]ConfigTable, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *PluginConfiguration) SetTables(v []ConfigTable)`

SetTables sets Tables field to given value.

### HasTables

`func (o *PluginConfiguration) HasTables() bool`

HasTables returns a boolean if a field has been set.

### GetFields

`func (o *PluginConfiguration) GetFields() []ConfigField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *PluginConfiguration) GetFieldsOk() (*[]ConfigField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *PluginConfiguration) SetFields(v []ConfigField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *PluginConfiguration) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


