# ConfigField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the configuration field. | 
**Value** | Pointer to **string** | The value for the configuration field. For encrypted or hashed fields, GETs will not return this attribute. To update an encrypted or hashed field, specify the new value in this attribute. | [optional] 
**EncryptedValue** | Pointer to **string** | For encrypted or hashed fields, this attribute contains the encrypted representation of the field&#39;s value, if a value is defined. If you do not want to update the stored value, this attribute should be passed back unchanged. | [optional] 
**Inherited** | Pointer to **bool** | Whether this field is inherited from its parent instance. If true, the value/encrypted value properties become read-only. The default value is false. | [optional] 

## Methods

### NewConfigField

`func NewConfigField(name string, ) *ConfigField`

NewConfigField instantiates a new ConfigField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigFieldWithDefaults

`func NewConfigFieldWithDefaults() *ConfigField`

NewConfigFieldWithDefaults instantiates a new ConfigField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConfigField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigField) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *ConfigField) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConfigField) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConfigField) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ConfigField) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetEncryptedValue

`func (o *ConfigField) GetEncryptedValue() string`

GetEncryptedValue returns the EncryptedValue field if non-nil, zero value otherwise.

### GetEncryptedValueOk

`func (o *ConfigField) GetEncryptedValueOk() (*string, bool)`

GetEncryptedValueOk returns a tuple with the EncryptedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedValue

`func (o *ConfigField) SetEncryptedValue(v string)`

SetEncryptedValue sets EncryptedValue field to given value.

### HasEncryptedValue

`func (o *ConfigField) HasEncryptedValue() bool`

HasEncryptedValue returns a boolean if a field has been set.

### GetInherited

`func (o *ConfigField) GetInherited() bool`

GetInherited returns the Inherited field if non-nil, zero value otherwise.

### GetInheritedOk

`func (o *ConfigField) GetInheritedOk() (*bool, bool)`

GetInheritedOk returns a tuple with the Inherited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInherited

`func (o *ConfigField) SetInherited(v bool)`

SetInherited sets Inherited field to given value.

### HasInherited

`func (o *ConfigField) HasInherited() bool`

HasInherited returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


