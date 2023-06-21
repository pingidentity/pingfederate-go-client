# ConnectionCredentialsOutboundBackChannelAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SslAuthKeyPairRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**ValidatePartnerCert** | Pointer to **interface{}** | Validate the partner server certificate. Default is true. | [optional] 

## Methods

### NewConnectionCredentialsOutboundBackChannelAuth

`func NewConnectionCredentialsOutboundBackChannelAuth() *ConnectionCredentialsOutboundBackChannelAuth`

NewConnectionCredentialsOutboundBackChannelAuth instantiates a new ConnectionCredentialsOutboundBackChannelAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionCredentialsOutboundBackChannelAuthWithDefaults

`func NewConnectionCredentialsOutboundBackChannelAuthWithDefaults() *ConnectionCredentialsOutboundBackChannelAuth`

NewConnectionCredentialsOutboundBackChannelAuthWithDefaults instantiates a new ConnectionCredentialsOutboundBackChannelAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSslAuthKeyPairRef

`func (o *ConnectionCredentialsOutboundBackChannelAuth) GetSslAuthKeyPairRef() ResourceLink`

GetSslAuthKeyPairRef returns the SslAuthKeyPairRef field if non-nil, zero value otherwise.

### GetSslAuthKeyPairRefOk

`func (o *ConnectionCredentialsOutboundBackChannelAuth) GetSslAuthKeyPairRefOk() (*ResourceLink, bool)`

GetSslAuthKeyPairRefOk returns a tuple with the SslAuthKeyPairRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslAuthKeyPairRef

`func (o *ConnectionCredentialsOutboundBackChannelAuth) SetSslAuthKeyPairRef(v ResourceLink)`

SetSslAuthKeyPairRef sets SslAuthKeyPairRef field to given value.

### HasSslAuthKeyPairRef

`func (o *ConnectionCredentialsOutboundBackChannelAuth) HasSslAuthKeyPairRef() bool`

HasSslAuthKeyPairRef returns a boolean if a field has been set.

### GetValidatePartnerCert

`func (o *ConnectionCredentialsOutboundBackChannelAuth) GetValidatePartnerCert() interface{}`

GetValidatePartnerCert returns the ValidatePartnerCert field if non-nil, zero value otherwise.

### GetValidatePartnerCertOk

`func (o *ConnectionCredentialsOutboundBackChannelAuth) GetValidatePartnerCertOk() (*interface{}, bool)`

GetValidatePartnerCertOk returns a tuple with the ValidatePartnerCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatePartnerCert

`func (o *ConnectionCredentialsOutboundBackChannelAuth) SetValidatePartnerCert(v interface{})`

SetValidatePartnerCert sets ValidatePartnerCert field to given value.

### HasValidatePartnerCert

`func (o *ConnectionCredentialsOutboundBackChannelAuth) HasValidatePartnerCert() bool`

HasValidatePartnerCert returns a boolean if a field has been set.

### SetValidatePartnerCertNil

`func (o *ConnectionCredentialsOutboundBackChannelAuth) SetValidatePartnerCertNil(b bool)`

 SetValidatePartnerCertNil sets the value for ValidatePartnerCert to be an explicit nil

### UnsetValidatePartnerCert
`func (o *ConnectionCredentialsOutboundBackChannelAuth) UnsetValidatePartnerCert()`

UnsetValidatePartnerCert ensures that no value is present for ValidatePartnerCert, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


