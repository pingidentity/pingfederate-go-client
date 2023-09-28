# SchemaAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the attribute. | [optional] 
**MultiValued** | Pointer to **bool** | Indicates whether the attribute is multi-valued. | [optional] 
**Types** | Pointer to **[]string** | Represents the name of each attribute type in case of multi-valued attribute. | [optional] 
**SubAttributes** | Pointer to **[]string** | List of sub-attributes for an attribute. | [optional] 

## Methods

### NewSchemaAttribute

`func NewSchemaAttribute() *SchemaAttribute`

NewSchemaAttribute instantiates a new SchemaAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaAttributeWithDefaults

`func NewSchemaAttributeWithDefaults() *SchemaAttribute`

NewSchemaAttributeWithDefaults instantiates a new SchemaAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SchemaAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SchemaAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SchemaAttribute) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SchemaAttribute) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMultiValued

`func (o *SchemaAttribute) GetMultiValued() bool`

GetMultiValued returns the MultiValued field if non-nil, zero value otherwise.

### GetMultiValuedOk

`func (o *SchemaAttribute) GetMultiValuedOk() (*bool, bool)`

GetMultiValuedOk returns a tuple with the MultiValued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValued

`func (o *SchemaAttribute) SetMultiValued(v bool)`

SetMultiValued sets MultiValued field to given value.

### HasMultiValued

`func (o *SchemaAttribute) HasMultiValued() bool`

HasMultiValued returns a boolean if a field has been set.

### GetTypes

`func (o *SchemaAttribute) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *SchemaAttribute) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *SchemaAttribute) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *SchemaAttribute) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetSubAttributes

`func (o *SchemaAttribute) GetSubAttributes() []string`

GetSubAttributes returns the SubAttributes field if non-nil, zero value otherwise.

### GetSubAttributesOk

`func (o *SchemaAttribute) GetSubAttributesOk() (*[]string, bool)`

GetSubAttributesOk returns a tuple with the SubAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAttributes

`func (o *SchemaAttribute) SetSubAttributes(v []string)`

SetSubAttributes sets SubAttributes field to given value.

### HasSubAttributes

`func (o *SchemaAttribute) HasSubAttributes() bool`

HasSubAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


