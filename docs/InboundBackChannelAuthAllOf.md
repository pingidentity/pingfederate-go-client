# InboundBackChannelAuthAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerificationSubjectDN** | Pointer to **string** | If this property is set, the verification trust model is Anchored. The verification certificate must be signed by a trusted CA and included in the incoming message, and the subject DN of the expected certificate is specified in this property. If this property is not set, then a primary verification certificate must be specified in the certs array. | [optional] 
**VerificationIssuerDN** | Pointer to **string** | If a verification Subject DN is provided, you can optionally restrict the issuer to a specific trusted CA by specifying its DN in this field. | [optional] 
**Certs** | Pointer to [**[]ConnectionCert**](ConnectionCert.md) | The certificate used for signature verification and XML encryption. | [optional] 
**RequireSsl** | Pointer to **bool** | Incoming HTTP transmissions must use a secure channel. | [optional] 

## Methods

### NewInboundBackChannelAuthAllOf

`func NewInboundBackChannelAuthAllOf() *InboundBackChannelAuthAllOf`

NewInboundBackChannelAuthAllOf instantiates a new InboundBackChannelAuthAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboundBackChannelAuthAllOfWithDefaults

`func NewInboundBackChannelAuthAllOfWithDefaults() *InboundBackChannelAuthAllOf`

NewInboundBackChannelAuthAllOfWithDefaults instantiates a new InboundBackChannelAuthAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerificationSubjectDN

`func (o *InboundBackChannelAuthAllOf) GetVerificationSubjectDN() string`

GetVerificationSubjectDN returns the VerificationSubjectDN field if non-nil, zero value otherwise.

### GetVerificationSubjectDNOk

`func (o *InboundBackChannelAuthAllOf) GetVerificationSubjectDNOk() (*string, bool)`

GetVerificationSubjectDNOk returns a tuple with the VerificationSubjectDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationSubjectDN

`func (o *InboundBackChannelAuthAllOf) SetVerificationSubjectDN(v string)`

SetVerificationSubjectDN sets VerificationSubjectDN field to given value.

### HasVerificationSubjectDN

`func (o *InboundBackChannelAuthAllOf) HasVerificationSubjectDN() bool`

HasVerificationSubjectDN returns a boolean if a field has been set.

### GetVerificationIssuerDN

`func (o *InboundBackChannelAuthAllOf) GetVerificationIssuerDN() string`

GetVerificationIssuerDN returns the VerificationIssuerDN field if non-nil, zero value otherwise.

### GetVerificationIssuerDNOk

`func (o *InboundBackChannelAuthAllOf) GetVerificationIssuerDNOk() (*string, bool)`

GetVerificationIssuerDNOk returns a tuple with the VerificationIssuerDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationIssuerDN

`func (o *InboundBackChannelAuthAllOf) SetVerificationIssuerDN(v string)`

SetVerificationIssuerDN sets VerificationIssuerDN field to given value.

### HasVerificationIssuerDN

`func (o *InboundBackChannelAuthAllOf) HasVerificationIssuerDN() bool`

HasVerificationIssuerDN returns a boolean if a field has been set.

### GetCerts

`func (o *InboundBackChannelAuthAllOf) GetCerts() []ConnectionCert`

GetCerts returns the Certs field if non-nil, zero value otherwise.

### GetCertsOk

`func (o *InboundBackChannelAuthAllOf) GetCertsOk() (*[]ConnectionCert, bool)`

GetCertsOk returns a tuple with the Certs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCerts

`func (o *InboundBackChannelAuthAllOf) SetCerts(v []ConnectionCert)`

SetCerts sets Certs field to given value.

### HasCerts

`func (o *InboundBackChannelAuthAllOf) HasCerts() bool`

HasCerts returns a boolean if a field has been set.

### GetRequireSsl

`func (o *InboundBackChannelAuthAllOf) GetRequireSsl() bool`

GetRequireSsl returns the RequireSsl field if non-nil, zero value otherwise.

### GetRequireSslOk

`func (o *InboundBackChannelAuthAllOf) GetRequireSslOk() (*bool, bool)`

GetRequireSslOk returns a tuple with the RequireSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSsl

`func (o *InboundBackChannelAuthAllOf) SetRequireSsl(v bool)`

SetRequireSsl sets RequireSsl field to given value.

### HasRequireSsl

`func (o *InboundBackChannelAuthAllOf) HasRequireSsl() bool`

HasRequireSsl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


