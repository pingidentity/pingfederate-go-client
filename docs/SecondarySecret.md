# SecondarySecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secret** | Pointer to **string** | Secondary client secret for Basic Authentication.  To update the secondary client secret, specify the plaintext value in this field.  This field will not be populated for GET requests. | [optional] 
**EncryptedSecret** | Pointer to **string** | For GET requests, this field contains the encrypted secondary client secret, if one exists.  For POST and PUT requests, if you wish to reuse the existing secret, this field should be passed back unchanged. | [optional] 
**ExpiryTime** | Pointer to **time.Time** | The expiry time of the secondary secret. | [optional] 

## Methods

### NewSecondarySecret

`func NewSecondarySecret() *SecondarySecret`

NewSecondarySecret instantiates a new SecondarySecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecondarySecretWithDefaults

`func NewSecondarySecretWithDefaults() *SecondarySecret`

NewSecondarySecretWithDefaults instantiates a new SecondarySecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecret

`func (o *SecondarySecret) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *SecondarySecret) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *SecondarySecret) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *SecondarySecret) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetEncryptedSecret

`func (o *SecondarySecret) GetEncryptedSecret() string`

GetEncryptedSecret returns the EncryptedSecret field if non-nil, zero value otherwise.

### GetEncryptedSecretOk

`func (o *SecondarySecret) GetEncryptedSecretOk() (*string, bool)`

GetEncryptedSecretOk returns a tuple with the EncryptedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedSecret

`func (o *SecondarySecret) SetEncryptedSecret(v string)`

SetEncryptedSecret sets EncryptedSecret field to given value.

### HasEncryptedSecret

`func (o *SecondarySecret) HasEncryptedSecret() bool`

HasEncryptedSecret returns a boolean if a field has been set.

### GetExpiryTime

`func (o *SecondarySecret) GetExpiryTime() time.Time`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *SecondarySecret) GetExpiryTimeOk() (*time.Time, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *SecondarySecret) SetExpiryTime(v time.Time)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *SecondarySecret) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


