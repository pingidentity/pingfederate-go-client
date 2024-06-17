# PasswordCredentialValidatorAttributeContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoreAttributes** | Pointer to [**[]PasswordCredentialValidatorAttribute**](PasswordCredentialValidatorAttribute.md) | A list of read-only attributes that are automatically populated by the password credential validator descriptor. | [optional] 
**ExtendedAttributes** | Pointer to [**[]PasswordCredentialValidatorAttribute**](PasswordCredentialValidatorAttribute.md) | A list of additional attributes that can be returned by the password credential validator. The extended attributes are only used if the adapter supports them. | [optional] 

## Methods

### NewPasswordCredentialValidatorAttributeContract

`func NewPasswordCredentialValidatorAttributeContract() *PasswordCredentialValidatorAttributeContract`

NewPasswordCredentialValidatorAttributeContract instantiates a new PasswordCredentialValidatorAttributeContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordCredentialValidatorAttributeContractWithDefaults

`func NewPasswordCredentialValidatorAttributeContractWithDefaults() *PasswordCredentialValidatorAttributeContract`

NewPasswordCredentialValidatorAttributeContractWithDefaults instantiates a new PasswordCredentialValidatorAttributeContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoreAttributes

`func (o *PasswordCredentialValidatorAttributeContract) GetCoreAttributes() []PasswordCredentialValidatorAttribute`

GetCoreAttributes returns the CoreAttributes field if non-nil, zero value otherwise.

### GetCoreAttributesOk

`func (o *PasswordCredentialValidatorAttributeContract) GetCoreAttributesOk() (*[]PasswordCredentialValidatorAttribute, bool)`

GetCoreAttributesOk returns a tuple with the CoreAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreAttributes

`func (o *PasswordCredentialValidatorAttributeContract) SetCoreAttributes(v []PasswordCredentialValidatorAttribute)`

SetCoreAttributes sets CoreAttributes field to given value.

### HasCoreAttributes

`func (o *PasswordCredentialValidatorAttributeContract) HasCoreAttributes() bool`

HasCoreAttributes returns a boolean if a field has been set.

### GetExtendedAttributes

`func (o *PasswordCredentialValidatorAttributeContract) GetExtendedAttributes() []PasswordCredentialValidatorAttribute`

GetExtendedAttributes returns the ExtendedAttributes field if non-nil, zero value otherwise.

### GetExtendedAttributesOk

`func (o *PasswordCredentialValidatorAttributeContract) GetExtendedAttributesOk() (*[]PasswordCredentialValidatorAttribute, bool)`

GetExtendedAttributesOk returns a tuple with the ExtendedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedAttributes

`func (o *PasswordCredentialValidatorAttributeContract) SetExtendedAttributes(v []PasswordCredentialValidatorAttribute)`

SetExtendedAttributes sets ExtendedAttributes field to given value.

### HasExtendedAttributes

`func (o *PasswordCredentialValidatorAttributeContract) HasExtendedAttributes() bool`

HasExtendedAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


