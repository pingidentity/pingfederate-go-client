# UsernamePasswordCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | The username. | [optional] 
**Password** | Pointer to **string** | User password.  To update the password, specify the plaintext value in this field.  This field will not be populated for GET requests. | [optional] 
**EncryptedPassword** | Pointer to **string** | For GET requests, this field contains the encrypted password, if one exists.  For POST and PUT requests, if you wish to reuse the existing password, this field should be passed back unchanged. | [optional] 

## Methods

### NewUsernamePasswordCredentials

`func NewUsernamePasswordCredentials() *UsernamePasswordCredentials`

NewUsernamePasswordCredentials instantiates a new UsernamePasswordCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsernamePasswordCredentialsWithDefaults

`func NewUsernamePasswordCredentialsWithDefaults() *UsernamePasswordCredentials`

NewUsernamePasswordCredentialsWithDefaults instantiates a new UsernamePasswordCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UsernamePasswordCredentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UsernamePasswordCredentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UsernamePasswordCredentials) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UsernamePasswordCredentials) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *UsernamePasswordCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UsernamePasswordCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UsernamePasswordCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UsernamePasswordCredentials) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetEncryptedPassword

`func (o *UsernamePasswordCredentials) GetEncryptedPassword() string`

GetEncryptedPassword returns the EncryptedPassword field if non-nil, zero value otherwise.

### GetEncryptedPasswordOk

`func (o *UsernamePasswordCredentials) GetEncryptedPasswordOk() (*string, bool)`

GetEncryptedPasswordOk returns a tuple with the EncryptedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedPassword

`func (o *UsernamePasswordCredentials) SetEncryptedPassword(v string)`

SetEncryptedPassword sets EncryptedPassword field to given value.

### HasEncryptedPassword

`func (o *UsernamePasswordCredentials) HasEncryptedPassword() bool`

HasEncryptedPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


