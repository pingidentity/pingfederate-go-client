# TextFieldDescriptorAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Encrypted** | Pointer to **bool** | Determines whether the field value should be masked in the UI and encrypted on disk. | [optional] 
**Size** | Pointer to **int64** | The size of the text field. | [optional] 

## Methods

### NewTextFieldDescriptorAllOf

`func NewTextFieldDescriptorAllOf() *TextFieldDescriptorAllOf`

NewTextFieldDescriptorAllOf instantiates a new TextFieldDescriptorAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextFieldDescriptorAllOfWithDefaults

`func NewTextFieldDescriptorAllOfWithDefaults() *TextFieldDescriptorAllOf`

NewTextFieldDescriptorAllOfWithDefaults instantiates a new TextFieldDescriptorAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncrypted

`func (o *TextFieldDescriptorAllOf) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *TextFieldDescriptorAllOf) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *TextFieldDescriptorAllOf) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *TextFieldDescriptorAllOf) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### GetSize

`func (o *TextFieldDescriptorAllOf) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TextFieldDescriptorAllOf) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TextFieldDescriptorAllOf) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *TextFieldDescriptorAllOf) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


