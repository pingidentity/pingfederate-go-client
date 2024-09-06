# LocalIdentityField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the local identity field. | 
**Id** | **string** | Id of the local identity field. | 
**Label** | **string** | Label of the local identity field. | 
**RegistrationPageField** | Pointer to **bool** | Whether this is a registration page field or not. | [optional] 
**ProfilePageField** | Pointer to **bool** | Whether this is a profile page field or not. | [optional] 
**Attributes** | Pointer to **map[string]bool** | Attributes of the local identity field. | [optional] 
**Options** | Pointer to **[]string** | The list of options for this selection field. | [optional] 
**DefaultValue** | Pointer to **string** | The default value for this field. | [optional] 

## Methods

### NewLocalIdentityField

`func NewLocalIdentityField(type_ string, id string, label string, ) *LocalIdentityField`

NewLocalIdentityField instantiates a new LocalIdentityField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalIdentityFieldWithDefaults

`func NewLocalIdentityFieldWithDefaults() *LocalIdentityField`

NewLocalIdentityFieldWithDefaults instantiates a new LocalIdentityField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LocalIdentityField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LocalIdentityField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LocalIdentityField) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *LocalIdentityField) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LocalIdentityField) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LocalIdentityField) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *LocalIdentityField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LocalIdentityField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LocalIdentityField) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetRegistrationPageField

`func (o *LocalIdentityField) GetRegistrationPageField() bool`

GetRegistrationPageField returns the RegistrationPageField field if non-nil, zero value otherwise.

### GetRegistrationPageFieldOk

`func (o *LocalIdentityField) GetRegistrationPageFieldOk() (*bool, bool)`

GetRegistrationPageFieldOk returns a tuple with the RegistrationPageField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationPageField

`func (o *LocalIdentityField) SetRegistrationPageField(v bool)`

SetRegistrationPageField sets RegistrationPageField field to given value.

### HasRegistrationPageField

`func (o *LocalIdentityField) HasRegistrationPageField() bool`

HasRegistrationPageField returns a boolean if a field has been set.

### GetProfilePageField

`func (o *LocalIdentityField) GetProfilePageField() bool`

GetProfilePageField returns the ProfilePageField field if non-nil, zero value otherwise.

### GetProfilePageFieldOk

`func (o *LocalIdentityField) GetProfilePageFieldOk() (*bool, bool)`

GetProfilePageFieldOk returns a tuple with the ProfilePageField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePageField

`func (o *LocalIdentityField) SetProfilePageField(v bool)`

SetProfilePageField sets ProfilePageField field to given value.

### HasProfilePageField

`func (o *LocalIdentityField) HasProfilePageField() bool`

HasProfilePageField returns a boolean if a field has been set.

### GetAttributes

`func (o *LocalIdentityField) GetAttributes() map[string]bool`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *LocalIdentityField) GetAttributesOk() (*map[string]bool, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *LocalIdentityField) SetAttributes(v map[string]bool)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *LocalIdentityField) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetOptions

`func (o *LocalIdentityField) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *LocalIdentityField) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *LocalIdentityField) SetOptions(v []string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *LocalIdentityField) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetDefaultValue

`func (o *LocalIdentityField) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *LocalIdentityField) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *LocalIdentityField) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *LocalIdentityField) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


