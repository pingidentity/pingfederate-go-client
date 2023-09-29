# KerberosKeySet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptedKeySet** | **string** | The encrypted key set. | 
**DeactivatedAt** | Pointer to **time.Time** | Time at which the key set was deactivated due to password change. This field is not populated if the key set is active. | [optional] 

## Methods

### NewKerberosKeySet

`func NewKerberosKeySet(encryptedKeySet string, ) *KerberosKeySet`

NewKerberosKeySet instantiates a new KerberosKeySet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKerberosKeySetWithDefaults

`func NewKerberosKeySetWithDefaults() *KerberosKeySet`

NewKerberosKeySetWithDefaults instantiates a new KerberosKeySet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptedKeySet

`func (o *KerberosKeySet) GetEncryptedKeySet() string`

GetEncryptedKeySet returns the EncryptedKeySet field if non-nil, zero value otherwise.

### GetEncryptedKeySetOk

`func (o *KerberosKeySet) GetEncryptedKeySetOk() (*string, bool)`

GetEncryptedKeySetOk returns a tuple with the EncryptedKeySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedKeySet

`func (o *KerberosKeySet) SetEncryptedKeySet(v string)`

SetEncryptedKeySet sets EncryptedKeySet field to given value.


### GetDeactivatedAt

`func (o *KerberosKeySet) GetDeactivatedAt() time.Time`

GetDeactivatedAt returns the DeactivatedAt field if non-nil, zero value otherwise.

### GetDeactivatedAtOk

`func (o *KerberosKeySet) GetDeactivatedAtOk() (*time.Time, bool)`

GetDeactivatedAtOk returns a tuple with the DeactivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivatedAt

`func (o *KerberosKeySet) SetDeactivatedAt(v time.Time)`

SetDeactivatedAt sets DeactivatedAt field to given value.

### HasDeactivatedAt

`func (o *KerberosKeySet) HasDeactivatedAt() bool`

HasDeactivatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


