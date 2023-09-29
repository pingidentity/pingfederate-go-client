# PersistentGrantContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoreAttributes** | [**[]PersistentGrantAttribute**](PersistentGrantAttribute.md) | This is a read-only list of persistent grant attributes and includes USER_KEY and USER_NAME. Changes to this field will be ignored. | 
**ExtendedAttributes** | Pointer to [**[]PersistentGrantAttribute**](PersistentGrantAttribute.md) | A list of additional attributes for the persistent grant contract. | [optional] 

## Methods

### NewPersistentGrantContract

`func NewPersistentGrantContract(coreAttributes []PersistentGrantAttribute, ) *PersistentGrantContract`

NewPersistentGrantContract instantiates a new PersistentGrantContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersistentGrantContractWithDefaults

`func NewPersistentGrantContractWithDefaults() *PersistentGrantContract`

NewPersistentGrantContractWithDefaults instantiates a new PersistentGrantContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoreAttributes

`func (o *PersistentGrantContract) GetCoreAttributes() []PersistentGrantAttribute`

GetCoreAttributes returns the CoreAttributes field if non-nil, zero value otherwise.

### GetCoreAttributesOk

`func (o *PersistentGrantContract) GetCoreAttributesOk() (*[]PersistentGrantAttribute, bool)`

GetCoreAttributesOk returns a tuple with the CoreAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreAttributes

`func (o *PersistentGrantContract) SetCoreAttributes(v []PersistentGrantAttribute)`

SetCoreAttributes sets CoreAttributes field to given value.


### GetExtendedAttributes

`func (o *PersistentGrantContract) GetExtendedAttributes() []PersistentGrantAttribute`

GetExtendedAttributes returns the ExtendedAttributes field if non-nil, zero value otherwise.

### GetExtendedAttributesOk

`func (o *PersistentGrantContract) GetExtendedAttributesOk() (*[]PersistentGrantAttribute, bool)`

GetExtendedAttributesOk returns a tuple with the ExtendedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedAttributes

`func (o *PersistentGrantContract) SetExtendedAttributes(v []PersistentGrantAttribute)`

SetExtendedAttributes sets ExtendedAttributes field to given value.

### HasExtendedAttributes

`func (o *PersistentGrantContract) HasExtendedAttributes() bool`

HasExtendedAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


