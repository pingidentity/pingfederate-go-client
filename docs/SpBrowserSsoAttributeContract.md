# SpBrowserSsoAttributeContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoreAttributes** | Pointer to [**[]SpBrowserSsoAttribute**](SpBrowserSsoAttribute.md) | A list of read-only assertion attributes (for example, SAML_SUBJECT) that are automatically populated by PingFederate. | [optional] 
**ExtendedAttributes** | Pointer to [**[]SpBrowserSsoAttribute**](SpBrowserSsoAttribute.md) | A list of additional attributes that are added to the outgoing assertion. | [optional] 

## Methods

### NewSpBrowserSsoAttributeContract

`func NewSpBrowserSsoAttributeContract() *SpBrowserSsoAttributeContract`

NewSpBrowserSsoAttributeContract instantiates a new SpBrowserSsoAttributeContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpBrowserSsoAttributeContractWithDefaults

`func NewSpBrowserSsoAttributeContractWithDefaults() *SpBrowserSsoAttributeContract`

NewSpBrowserSsoAttributeContractWithDefaults instantiates a new SpBrowserSsoAttributeContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoreAttributes

`func (o *SpBrowserSsoAttributeContract) GetCoreAttributes() []SpBrowserSsoAttribute`

GetCoreAttributes returns the CoreAttributes field if non-nil, zero value otherwise.

### GetCoreAttributesOk

`func (o *SpBrowserSsoAttributeContract) GetCoreAttributesOk() (*[]SpBrowserSsoAttribute, bool)`

GetCoreAttributesOk returns a tuple with the CoreAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreAttributes

`func (o *SpBrowserSsoAttributeContract) SetCoreAttributes(v []SpBrowserSsoAttribute)`

SetCoreAttributes sets CoreAttributes field to given value.

### HasCoreAttributes

`func (o *SpBrowserSsoAttributeContract) HasCoreAttributes() bool`

HasCoreAttributes returns a boolean if a field has been set.

### GetExtendedAttributes

`func (o *SpBrowserSsoAttributeContract) GetExtendedAttributes() []SpBrowserSsoAttribute`

GetExtendedAttributes returns the ExtendedAttributes field if non-nil, zero value otherwise.

### GetExtendedAttributesOk

`func (o *SpBrowserSsoAttributeContract) GetExtendedAttributesOk() (*[]SpBrowserSsoAttribute, bool)`

GetExtendedAttributesOk returns a tuple with the ExtendedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedAttributes

`func (o *SpBrowserSsoAttributeContract) SetExtendedAttributes(v []SpBrowserSsoAttribute)`

SetExtendedAttributes sets ExtendedAttributes field to given value.

### HasExtendedAttributes

`func (o *SpBrowserSsoAttributeContract) HasExtendedAttributes() bool`

HasExtendedAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


