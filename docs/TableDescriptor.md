# TableDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the table. | [optional] 
**Description** | Pointer to **string** | Description for the table. | [optional] 
**Columns** | Pointer to [**[]FieldDescriptor**](FieldDescriptor.md) | Get the columns in the table. | [optional] 
**Label** | Pointer to **string** | Label for the table to be displayed in the administrative console. | [optional] 
**RequireDefaultRow** | Pointer to **bool** | Configure whether this table requires default row to be set. | [optional] 

## Methods

### NewTableDescriptor

`func NewTableDescriptor() *TableDescriptor`

NewTableDescriptor instantiates a new TableDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableDescriptorWithDefaults

`func NewTableDescriptorWithDefaults() *TableDescriptor`

NewTableDescriptorWithDefaults instantiates a new TableDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TableDescriptor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TableDescriptor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TableDescriptor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TableDescriptor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *TableDescriptor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TableDescriptor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TableDescriptor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TableDescriptor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetColumns

`func (o *TableDescriptor) GetColumns() []FieldDescriptor`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *TableDescriptor) GetColumnsOk() (*[]FieldDescriptor, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *TableDescriptor) SetColumns(v []FieldDescriptor)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *TableDescriptor) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetLabel

`func (o *TableDescriptor) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TableDescriptor) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TableDescriptor) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *TableDescriptor) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRequireDefaultRow

`func (o *TableDescriptor) GetRequireDefaultRow() bool`

GetRequireDefaultRow returns the RequireDefaultRow field if non-nil, zero value otherwise.

### GetRequireDefaultRowOk

`func (o *TableDescriptor) GetRequireDefaultRowOk() (*bool, bool)`

GetRequireDefaultRowOk returns a tuple with the RequireDefaultRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireDefaultRow

`func (o *TableDescriptor) SetRequireDefaultRow(v bool)`

SetRequireDefaultRow sets RequireDefaultRow field to given value.

### HasRequireDefaultRow

`func (o *TableDescriptor) HasRequireDefaultRow() bool`

HasRequireDefaultRow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


