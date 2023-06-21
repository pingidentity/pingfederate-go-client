# ServiceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id of the service. | [optional] 
**SharedSecret** | Pointer to **string** | Shared secret for the service. | [optional] 
**EncryptedSharedSecret** | Pointer to **string** | Encrypted shared secret for the service. | [optional] 

## Methods

### NewServiceModel

`func NewServiceModel() *ServiceModel`

NewServiceModel instantiates a new ServiceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceModelWithDefaults

`func NewServiceModelWithDefaults() *ServiceModel`

NewServiceModelWithDefaults instantiates a new ServiceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSharedSecret

`func (o *ServiceModel) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *ServiceModel) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *ServiceModel) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *ServiceModel) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### GetEncryptedSharedSecret

`func (o *ServiceModel) GetEncryptedSharedSecret() string`

GetEncryptedSharedSecret returns the EncryptedSharedSecret field if non-nil, zero value otherwise.

### GetEncryptedSharedSecretOk

`func (o *ServiceModel) GetEncryptedSharedSecretOk() (*string, bool)`

GetEncryptedSharedSecretOk returns a tuple with the EncryptedSharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedSharedSecret

`func (o *ServiceModel) SetEncryptedSharedSecret(v string)`

SetEncryptedSharedSecret sets EncryptedSharedSecret field to given value.

### HasEncryptedSharedSecret

`func (o *ServiceModel) HasEncryptedSharedSecret() bool`

HasEncryptedSharedSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


