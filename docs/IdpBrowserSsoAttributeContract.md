# IdpBrowserSsoAttributeContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoreAttributes** | Pointer to [**[]IdpBrowserSsoAttribute**](IdpBrowserSsoAttribute.md) | A list of read-only assertion attributes that are automatically populated by PingFederate. | [optional] 
**ExtendedAttributes** | Pointer to [**[]IdpBrowserSsoAttribute**](IdpBrowserSsoAttribute.md) | A list of additional attributes that are present in the incoming assertion. | [optional] 

## Methods

### NewIdpBrowserSsoAttributeContract

`func NewIdpBrowserSsoAttributeContract() *IdpBrowserSsoAttributeContract`

NewIdpBrowserSsoAttributeContract instantiates a new IdpBrowserSsoAttributeContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpBrowserSsoAttributeContractWithDefaults

`func NewIdpBrowserSsoAttributeContractWithDefaults() *IdpBrowserSsoAttributeContract`

NewIdpBrowserSsoAttributeContractWithDefaults instantiates a new IdpBrowserSsoAttributeContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoreAttributes

`func (o *IdpBrowserSsoAttributeContract) GetCoreAttributes() []IdpBrowserSsoAttribute`

GetCoreAttributes returns the CoreAttributes field if non-nil, zero value otherwise.

### GetCoreAttributesOk

`func (o *IdpBrowserSsoAttributeContract) GetCoreAttributesOk() (*[]IdpBrowserSsoAttribute, bool)`

GetCoreAttributesOk returns a tuple with the CoreAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreAttributes

`func (o *IdpBrowserSsoAttributeContract) SetCoreAttributes(v []IdpBrowserSsoAttribute)`

SetCoreAttributes sets CoreAttributes field to given value.

### HasCoreAttributes

`func (o *IdpBrowserSsoAttributeContract) HasCoreAttributes() bool`

HasCoreAttributes returns a boolean if a field has been set.

### GetExtendedAttributes

`func (o *IdpBrowserSsoAttributeContract) GetExtendedAttributes() []IdpBrowserSsoAttribute`

GetExtendedAttributes returns the ExtendedAttributes field if non-nil, zero value otherwise.

### GetExtendedAttributesOk

`func (o *IdpBrowserSsoAttributeContract) GetExtendedAttributesOk() (*[]IdpBrowserSsoAttribute, bool)`

GetExtendedAttributesOk returns a tuple with the ExtendedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedAttributes

`func (o *IdpBrowserSsoAttributeContract) SetExtendedAttributes(v []IdpBrowserSsoAttribute)`

SetExtendedAttributes sets ExtendedAttributes field to given value.

### HasExtendedAttributes

`func (o *IdpBrowserSsoAttributeContract) HasExtendedAttributes() bool`

HasExtendedAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


