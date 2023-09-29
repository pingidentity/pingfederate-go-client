# ExtendedProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The property name. | [optional] 
**Description** | Pointer to **string** | The property description. | [optional] 
**MultiValued** | Pointer to **bool** | Indicates whether the property should allow multiple values. | [optional] 

## Methods

### NewExtendedProperty

`func NewExtendedProperty() *ExtendedProperty`

NewExtendedProperty instantiates a new ExtendedProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedPropertyWithDefaults

`func NewExtendedPropertyWithDefaults() *ExtendedProperty`

NewExtendedPropertyWithDefaults instantiates a new ExtendedProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExtendedProperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtendedProperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtendedProperty) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExtendedProperty) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ExtendedProperty) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExtendedProperty) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExtendedProperty) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExtendedProperty) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMultiValued

`func (o *ExtendedProperty) GetMultiValued() bool`

GetMultiValued returns the MultiValued field if non-nil, zero value otherwise.

### GetMultiValuedOk

`func (o *ExtendedProperty) GetMultiValuedOk() (*bool, bool)`

GetMultiValuedOk returns a tuple with the MultiValued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValued

`func (o *ExtendedProperty) SetMultiValued(v bool)`

SetMultiValued sets MultiValued field to given value.

### HasMultiValued

`func (o *ExtendedProperty) HasMultiValued() bool`

HasMultiValued returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


