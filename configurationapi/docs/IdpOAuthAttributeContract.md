# IdpOAuthAttributeContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoreAttributes** | Pointer to [**[]IdpBrowserSsoAttribute**](IdpBrowserSsoAttribute.md) | A list of read-only assertion attributes that are automatically populated by PingFederate. | [optional] 
**ExtendedAttributes** | Pointer to [**[]IdpBrowserSsoAttribute**](IdpBrowserSsoAttribute.md) | A list of additional attributes that are present in the incoming assertion. | [optional] 

## Methods

### NewIdpOAuthAttributeContract

`func NewIdpOAuthAttributeContract() *IdpOAuthAttributeContract`

NewIdpOAuthAttributeContract instantiates a new IdpOAuthAttributeContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpOAuthAttributeContractWithDefaults

`func NewIdpOAuthAttributeContractWithDefaults() *IdpOAuthAttributeContract`

NewIdpOAuthAttributeContractWithDefaults instantiates a new IdpOAuthAttributeContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoreAttributes

`func (o *IdpOAuthAttributeContract) GetCoreAttributes() []IdpBrowserSsoAttribute`

GetCoreAttributes returns the CoreAttributes field if non-nil, zero value otherwise.

### GetCoreAttributesOk

`func (o *IdpOAuthAttributeContract) GetCoreAttributesOk() (*[]IdpBrowserSsoAttribute, bool)`

GetCoreAttributesOk returns a tuple with the CoreAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreAttributes

`func (o *IdpOAuthAttributeContract) SetCoreAttributes(v []IdpBrowserSsoAttribute)`

SetCoreAttributes sets CoreAttributes field to given value.

### HasCoreAttributes

`func (o *IdpOAuthAttributeContract) HasCoreAttributes() bool`

HasCoreAttributes returns a boolean if a field has been set.

### GetExtendedAttributes

`func (o *IdpOAuthAttributeContract) GetExtendedAttributes() []IdpBrowserSsoAttribute`

GetExtendedAttributes returns the ExtendedAttributes field if non-nil, zero value otherwise.

### GetExtendedAttributesOk

`func (o *IdpOAuthAttributeContract) GetExtendedAttributesOk() (*[]IdpBrowserSsoAttribute, bool)`

GetExtendedAttributesOk returns a tuple with the ExtendedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedAttributes

`func (o *IdpOAuthAttributeContract) SetExtendedAttributes(v []IdpBrowserSsoAttribute)`

SetExtendedAttributes sets ExtendedAttributes field to given value.

### HasExtendedAttributes

`func (o *IdpOAuthAttributeContract) HasExtendedAttributes() bool`

HasExtendedAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


