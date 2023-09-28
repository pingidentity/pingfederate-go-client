# TextFieldDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Encrypted** | Pointer to **bool** | Determines whether the field value should be masked in the UI and encrypted on disk. | [optional] 
**Size** | Pointer to **int64** | The size of the text field. | [optional] 

## Methods

### NewTextFieldDescriptor

`func NewTextFieldDescriptor() *TextFieldDescriptor`

NewTextFieldDescriptor instantiates a new TextFieldDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextFieldDescriptorWithDefaults

`func NewTextFieldDescriptorWithDefaults() *TextFieldDescriptor`

NewTextFieldDescriptorWithDefaults instantiates a new TextFieldDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncrypted

`func (o *TextFieldDescriptor) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *TextFieldDescriptor) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *TextFieldDescriptor) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *TextFieldDescriptor) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### GetSize

`func (o *TextFieldDescriptor) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TextFieldDescriptor) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TextFieldDescriptor) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *TextFieldDescriptor) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


