# DecryptionKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryKeyRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**SecondaryKeyPairRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 

## Methods

### NewDecryptionKeys

`func NewDecryptionKeys() *DecryptionKeys`

NewDecryptionKeys instantiates a new DecryptionKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecryptionKeysWithDefaults

`func NewDecryptionKeysWithDefaults() *DecryptionKeys`

NewDecryptionKeysWithDefaults instantiates a new DecryptionKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimaryKeyRef

`func (o *DecryptionKeys) GetPrimaryKeyRef() ResourceLink`

GetPrimaryKeyRef returns the PrimaryKeyRef field if non-nil, zero value otherwise.

### GetPrimaryKeyRefOk

`func (o *DecryptionKeys) GetPrimaryKeyRefOk() (*ResourceLink, bool)`

GetPrimaryKeyRefOk returns a tuple with the PrimaryKeyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryKeyRef

`func (o *DecryptionKeys) SetPrimaryKeyRef(v ResourceLink)`

SetPrimaryKeyRef sets PrimaryKeyRef field to given value.

### HasPrimaryKeyRef

`func (o *DecryptionKeys) HasPrimaryKeyRef() bool`

HasPrimaryKeyRef returns a boolean if a field has been set.

### GetSecondaryKeyPairRef

`func (o *DecryptionKeys) GetSecondaryKeyPairRef() ResourceLink`

GetSecondaryKeyPairRef returns the SecondaryKeyPairRef field if non-nil, zero value otherwise.

### GetSecondaryKeyPairRefOk

`func (o *DecryptionKeys) GetSecondaryKeyPairRefOk() (*ResourceLink, bool)`

GetSecondaryKeyPairRefOk returns a tuple with the SecondaryKeyPairRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryKeyPairRef

`func (o *DecryptionKeys) SetSecondaryKeyPairRef(v ResourceLink)`

SetSecondaryKeyPairRef sets SecondaryKeyPairRef field to given value.

### HasSecondaryKeyPairRef

`func (o *DecryptionKeys) HasSecondaryKeyPairRef() bool`

HasSecondaryKeyPairRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


