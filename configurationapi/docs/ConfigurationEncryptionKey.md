# ConfigurationEncryptionKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyId** | Pointer to **string** | The id of the key. | [optional] 
**CreationDate** | Pointer to **time.Time** | The creation date of the key. | [optional] 

## Methods

### NewConfigurationEncryptionKey

`func NewConfigurationEncryptionKey() *ConfigurationEncryptionKey`

NewConfigurationEncryptionKey instantiates a new ConfigurationEncryptionKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationEncryptionKeyWithDefaults

`func NewConfigurationEncryptionKeyWithDefaults() *ConfigurationEncryptionKey`

NewConfigurationEncryptionKeyWithDefaults instantiates a new ConfigurationEncryptionKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyId

`func (o *ConfigurationEncryptionKey) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *ConfigurationEncryptionKey) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *ConfigurationEncryptionKey) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *ConfigurationEncryptionKey) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetCreationDate

`func (o *ConfigurationEncryptionKey) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ConfigurationEncryptionKey) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ConfigurationEncryptionKey) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *ConfigurationEncryptionKey) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


