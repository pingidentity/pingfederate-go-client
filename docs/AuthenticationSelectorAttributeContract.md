# AuthenticationSelectorAttributeContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtendedAttributes** | Pointer to [**[]AuthenticationSelectorAttribute**](AuthenticationSelectorAttribute.md) | A list of additional attributes that can be returned by the Authentication Selector. The extended attributes are only used if the Authentication Selector supports them. | [optional] 

## Methods

### NewAuthenticationSelectorAttributeContract

`func NewAuthenticationSelectorAttributeContract() *AuthenticationSelectorAttributeContract`

NewAuthenticationSelectorAttributeContract instantiates a new AuthenticationSelectorAttributeContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationSelectorAttributeContractWithDefaults

`func NewAuthenticationSelectorAttributeContractWithDefaults() *AuthenticationSelectorAttributeContract`

NewAuthenticationSelectorAttributeContractWithDefaults instantiates a new AuthenticationSelectorAttributeContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtendedAttributes

`func (o *AuthenticationSelectorAttributeContract) GetExtendedAttributes() []AuthenticationSelectorAttribute`

GetExtendedAttributes returns the ExtendedAttributes field if non-nil, zero value otherwise.

### GetExtendedAttributesOk

`func (o *AuthenticationSelectorAttributeContract) GetExtendedAttributesOk() (*[]AuthenticationSelectorAttribute, bool)`

GetExtendedAttributesOk returns a tuple with the ExtendedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedAttributes

`func (o *AuthenticationSelectorAttributeContract) SetExtendedAttributes(v []AuthenticationSelectorAttribute)`

SetExtendedAttributes sets ExtendedAttributes field to given value.

### HasExtendedAttributes

`func (o *AuthenticationSelectorAttributeContract) HasExtendedAttributes() bool`

HasExtendedAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


