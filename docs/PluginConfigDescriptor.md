# PluginConfigDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of this plugin. | [optional] 
**Fields** | Pointer to [**[]FieldDescriptor**](FieldDescriptor.md) | The configuration fields available for this plugin. | [optional] 
**Tables** | Pointer to [**[]TableDescriptor**](TableDescriptor.md) | Configuration tables available for this plugin. | [optional] 
**ActionDescriptors** | Pointer to [**[]ActionDescriptor**](ActionDescriptor.md) | The available actions for this plugin. | [optional] 

## Methods

### NewPluginConfigDescriptor

`func NewPluginConfigDescriptor() *PluginConfigDescriptor`

NewPluginConfigDescriptor instantiates a new PluginConfigDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginConfigDescriptorWithDefaults

`func NewPluginConfigDescriptorWithDefaults() *PluginConfigDescriptor`

NewPluginConfigDescriptorWithDefaults instantiates a new PluginConfigDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PluginConfigDescriptor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PluginConfigDescriptor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PluginConfigDescriptor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PluginConfigDescriptor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFields

`func (o *PluginConfigDescriptor) GetFields() []FieldDescriptor`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *PluginConfigDescriptor) GetFieldsOk() (*[]FieldDescriptor, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *PluginConfigDescriptor) SetFields(v []FieldDescriptor)`

SetFields sets Fields field to given value.

### HasFields

`func (o *PluginConfigDescriptor) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetTables

`func (o *PluginConfigDescriptor) GetTables() []TableDescriptor`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *PluginConfigDescriptor) GetTablesOk() (*[]TableDescriptor, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *PluginConfigDescriptor) SetTables(v []TableDescriptor)`

SetTables sets Tables field to given value.

### HasTables

`func (o *PluginConfigDescriptor) HasTables() bool`

HasTables returns a boolean if a field has been set.

### GetActionDescriptors

`func (o *PluginConfigDescriptor) GetActionDescriptors() []ActionDescriptor`

GetActionDescriptors returns the ActionDescriptors field if non-nil, zero value otherwise.

### GetActionDescriptorsOk

`func (o *PluginConfigDescriptor) GetActionDescriptorsOk() (*[]ActionDescriptor, bool)`

GetActionDescriptorsOk returns a tuple with the ActionDescriptors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDescriptors

`func (o *PluginConfigDescriptor) SetActionDescriptors(v []ActionDescriptor)`

SetActionDescriptors sets ActionDescriptors field to given value.

### HasActionDescriptors

`func (o *PluginConfigDescriptor) HasActionDescriptors() bool`

HasActionDescriptors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


