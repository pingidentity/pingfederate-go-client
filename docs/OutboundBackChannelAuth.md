# OutboundBackChannelAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SslAuthKeyPairRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**ValidatePartnerCert** | Pointer to **bool** | Validate the partner server certificate. Default is true. | [optional] 

## Methods

### NewOutboundBackChannelAuth

`func NewOutboundBackChannelAuth() *OutboundBackChannelAuth`

NewOutboundBackChannelAuth instantiates a new OutboundBackChannelAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboundBackChannelAuthWithDefaults

`func NewOutboundBackChannelAuthWithDefaults() *OutboundBackChannelAuth`

NewOutboundBackChannelAuthWithDefaults instantiates a new OutboundBackChannelAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSslAuthKeyPairRef

`func (o *OutboundBackChannelAuth) GetSslAuthKeyPairRef() ResourceLink`

GetSslAuthKeyPairRef returns the SslAuthKeyPairRef field if non-nil, zero value otherwise.

### GetSslAuthKeyPairRefOk

`func (o *OutboundBackChannelAuth) GetSslAuthKeyPairRefOk() (*ResourceLink, bool)`

GetSslAuthKeyPairRefOk returns a tuple with the SslAuthKeyPairRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslAuthKeyPairRef

`func (o *OutboundBackChannelAuth) SetSslAuthKeyPairRef(v ResourceLink)`

SetSslAuthKeyPairRef sets SslAuthKeyPairRef field to given value.

### HasSslAuthKeyPairRef

`func (o *OutboundBackChannelAuth) HasSslAuthKeyPairRef() bool`

HasSslAuthKeyPairRef returns a boolean if a field has been set.

### GetValidatePartnerCert

`func (o *OutboundBackChannelAuth) GetValidatePartnerCert() bool`

GetValidatePartnerCert returns the ValidatePartnerCert field if non-nil, zero value otherwise.

### GetValidatePartnerCertOk

`func (o *OutboundBackChannelAuth) GetValidatePartnerCertOk() (*bool, bool)`

GetValidatePartnerCertOk returns a tuple with the ValidatePartnerCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatePartnerCert

`func (o *OutboundBackChannelAuth) SetValidatePartnerCert(v bool)`

SetValidatePartnerCert sets ValidatePartnerCert field to given value.

### HasValidatePartnerCert

`func (o *OutboundBackChannelAuth) HasValidatePartnerCert() bool`

HasValidatePartnerCert returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


