# SystemKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationDate** | Pointer to **time.Time** | Creation time of the key. | [optional] 
**EncryptedKeyData** | Pointer to **string** | The system key encrypted. | [optional] 
**KeyData** | Pointer to **string** | The clear text system key base 64 encoded. The system key must be 32 bytes before base 64 encoding. | [optional] 

## Methods

### NewSystemKey

`func NewSystemKey() *SystemKey`

NewSystemKey instantiates a new SystemKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemKeyWithDefaults

`func NewSystemKeyWithDefaults() *SystemKey`

NewSystemKeyWithDefaults instantiates a new SystemKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationDate

`func (o *SystemKey) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *SystemKey) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *SystemKey) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *SystemKey) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetEncryptedKeyData

`func (o *SystemKey) GetEncryptedKeyData() string`

GetEncryptedKeyData returns the EncryptedKeyData field if non-nil, zero value otherwise.

### GetEncryptedKeyDataOk

`func (o *SystemKey) GetEncryptedKeyDataOk() (*string, bool)`

GetEncryptedKeyDataOk returns a tuple with the EncryptedKeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedKeyData

`func (o *SystemKey) SetEncryptedKeyData(v string)`

SetEncryptedKeyData sets EncryptedKeyData field to given value.

### HasEncryptedKeyData

`func (o *SystemKey) HasEncryptedKeyData() bool`

HasEncryptedKeyData returns a boolean if a field has been set.

### GetKeyData

`func (o *SystemKey) GetKeyData() string`

GetKeyData returns the KeyData field if non-nil, zero value otherwise.

### GetKeyDataOk

`func (o *SystemKey) GetKeyDataOk() (*string, bool)`

GetKeyDataOk returns a tuple with the KeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyData

`func (o *SystemKey) SetKeyData(v string)`

SetKeyData sets KeyData field to given value.

### HasKeyData

`func (o *SystemKey) HasKeyData() bool`

HasKeyData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


