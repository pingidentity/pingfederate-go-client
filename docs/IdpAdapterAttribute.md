# IdpAdapterAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of this attribute. | 
**Pseudonym** | Pointer to **bool** | Specifies whether this attribute is used to construct a pseudonym for the SP. Defaults to false. | [optional] 
**Masked** | Pointer to **bool** | Specifies whether this attribute is masked in PingFederate logs. Defaults to false. | [optional] 

## Methods

### NewIdpAdapterAttribute

`func NewIdpAdapterAttribute(name string, ) *IdpAdapterAttribute`

NewIdpAdapterAttribute instantiates a new IdpAdapterAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpAdapterAttributeWithDefaults

`func NewIdpAdapterAttributeWithDefaults() *IdpAdapterAttribute`

NewIdpAdapterAttributeWithDefaults instantiates a new IdpAdapterAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IdpAdapterAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdpAdapterAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdpAdapterAttribute) SetName(v string)`

SetName sets Name field to given value.


### GetPseudonym

`func (o *IdpAdapterAttribute) GetPseudonym() bool`

GetPseudonym returns the Pseudonym field if non-nil, zero value otherwise.

### GetPseudonymOk

`func (o *IdpAdapterAttribute) GetPseudonymOk() (*bool, bool)`

GetPseudonymOk returns a tuple with the Pseudonym field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPseudonym

`func (o *IdpAdapterAttribute) SetPseudonym(v bool)`

SetPseudonym sets Pseudonym field to given value.

### HasPseudonym

`func (o *IdpAdapterAttribute) HasPseudonym() bool`

HasPseudonym returns a boolean if a field has been set.

### GetMasked

`func (o *IdpAdapterAttribute) GetMasked() bool`

GetMasked returns the Masked field if non-nil, zero value otherwise.

### GetMaskedOk

`func (o *IdpAdapterAttribute) GetMaskedOk() (*bool, bool)`

GetMaskedOk returns a tuple with the Masked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasked

`func (o *IdpAdapterAttribute) SetMasked(v bool)`

SetMasked sets Masked field to given value.

### HasMasked

`func (o *IdpAdapterAttribute) HasMasked() bool`

HasMasked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


