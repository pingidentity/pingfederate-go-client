# FieldConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | Pointer to [**[]LocalIdentityField**](LocalIdentityField.md) | The field configuration for the local identity profile. | [optional] 
**StripSpaceFromUniqueField** | Pointer to **bool** | Strip leading/trailing spaces from unique ID field. Default is true. | [optional] 

## Methods

### NewFieldConfig

`func NewFieldConfig() *FieldConfig`

NewFieldConfig instantiates a new FieldConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldConfigWithDefaults

`func NewFieldConfigWithDefaults() *FieldConfig`

NewFieldConfigWithDefaults instantiates a new FieldConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *FieldConfig) GetFields() []LocalIdentityField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *FieldConfig) GetFieldsOk() (*[]LocalIdentityField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *FieldConfig) SetFields(v []LocalIdentityField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *FieldConfig) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetStripSpaceFromUniqueField

`func (o *FieldConfig) GetStripSpaceFromUniqueField() bool`

GetStripSpaceFromUniqueField returns the StripSpaceFromUniqueField field if non-nil, zero value otherwise.

### GetStripSpaceFromUniqueFieldOk

`func (o *FieldConfig) GetStripSpaceFromUniqueFieldOk() (*bool, bool)`

GetStripSpaceFromUniqueFieldOk returns a tuple with the StripSpaceFromUniqueField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripSpaceFromUniqueField

`func (o *FieldConfig) SetStripSpaceFromUniqueField(v bool)`

SetStripSpaceFromUniqueField sets StripSpaceFromUniqueField field to given value.

### HasStripSpaceFromUniqueField

`func (o *FieldConfig) HasStripSpaceFromUniqueField() bool`

HasStripSpaceFromUniqueField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


