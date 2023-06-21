# ConnectionCredentialsInboundBackChannelAuthAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerificationSubjectDN** | Pointer to **interface{}** | If this property is set, the verification trust model is Anchored. The verification certificate must be signed by a trusted CA and included in the incoming message, and the subject DN of the expected certificate is specified in this property. If this property is not set, then a primary verification certificate must be specified in the certs array. | [optional] 
**VerificationIssuerDN** | Pointer to **interface{}** | If a verification Subject DN is provided, you can optionally restrict the issuer to a specific trusted CA by specifying its DN in this field. | [optional] 
**Certs** | Pointer to [**[]ConnectionCert**](ConnectionCert.md) | The certificate used for signature verification and XML encryption. | [optional] 
**RequireSsl** | Pointer to **interface{}** | Incoming HTTP transmissions must use a secure channel. | [optional] 

## Methods

### NewConnectionCredentialsInboundBackChannelAuthAllOf

`func NewConnectionCredentialsInboundBackChannelAuthAllOf() *ConnectionCredentialsInboundBackChannelAuthAllOf`

NewConnectionCredentialsInboundBackChannelAuthAllOf instantiates a new ConnectionCredentialsInboundBackChannelAuthAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionCredentialsInboundBackChannelAuthAllOfWithDefaults

`func NewConnectionCredentialsInboundBackChannelAuthAllOfWithDefaults() *ConnectionCredentialsInboundBackChannelAuthAllOf`

NewConnectionCredentialsInboundBackChannelAuthAllOfWithDefaults instantiates a new ConnectionCredentialsInboundBackChannelAuthAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerificationSubjectDN

`func (o *ConnectionCredentialsInboundBackChannelAuthAllOf) GetVerificationSubjectDN() interface{}`

GetVerificationSubjectDN returns the VerificationSubjectDN field if non-nil, zero value otherwise.

### GetVerificationSubjectDNOk

`func (o *ConnectionCredentialsInboundBackChannelAuthAllOf) GetVerificationSubjectDNOk() (*interface{}, bool)`

GetVerificationSubjectDNOk returns a tuple with the VerificationSubjectDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationSubjectDN

`func (o *ConnectionCredentialsInboundBackChannelAuthAllOf) SetVerificationSubjectDN(v interface{})`

SetVerificationSubjectDN sets VerificationSubjectDN field to given value.

### HasVerificationSubjectDN

`func (o *ConnectionCredentialsInboundBackChannelAuthAllOf) HasVerificationSubjectDN() bool`

HasVerificationSubjectDN returns a boolean if a field has been set.

### SetVerificationSubjectDNNil

`func (o *ConnectionCredentialsInboundBackChannelAuthAllOf) SetVerificationSubjectDNNil(b bool)`

 SetVerificationSubjectDNNil sets the value for VerificationSubjectDN to be an explicit nil

### UnsetVerificationSubjectDN
`func (o *ConnectionCredentialsInboundBackChannelAuthAllOf) UnsetVerificationSubjectDN()`

UnsetVerificationSubjectDN ensures that no value is present for VerificationSubjectDN, not even an explicit nil
### GetVerificationIssuerDN

`func (o *ConnectionCredentialsInboundBackChannelAuthAllOf) GetVerificationIssuerDN() interface{}`

GetVerificationIssuerDN returns the VerificationIssuerDN field if non-nil, zero value otherwise.

### GetVerificationIssuerDNOk

`func (o *ConnectionCredentialsInboundBackChannelAuthAllOf) GetVerificationIssuerDNOk() (*interface{}, bool)`

GetVerificationIssuerDNOk returns a tuple with the VerificationIssuerDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationIssuerDN

`func (o *ConnectionCredentialsInboundBackChannelAuthAllOf) SetVerificationIssuerDN(v interface{})`

SetVerificationIssuerDN sets VerificationIssuerDN field to given value.

### HasVerificationIssuerDN

`func (o *ConnectionCredentialsInboundBackChannelAuthAllOf) HasVerificationIssuerDN() bool`

HasVerificationIssuerDN returns a boolean if a field has been set.

### SetVerificationIssuerDNNil

`func (o *ConnectionCredentialsInboundBackChannelAuthAllOf) SetVerificationIssuerDNNil(b bool)`

 SetVerificationIssuerDNNil sets the value for VerificationIssuerDN to be an explicit nil

### UnsetVerificationIssuerDN
`func (o *ConnectionCredentialsInboundBackChannelAuthAllOf) UnsetVerificationIssuerDN()`

UnsetVerificationIssuerDN ensures that no value is present for VerificationIssuerDN, not even an explicit nil
### GetCerts

`func (o *ConnectionCredentialsInboundBackChannelAuthAllOf) GetCerts() []ConnectionCert`

GetCerts returns the Certs field if non-nil, zero value otherwise.

### GetCertsOk

`func (o *ConnectionCredentialsInboundBackChannelAuthAllOf) GetCertsOk() (*[]ConnectionCert, bool)`

GetCertsOk returns a tuple with the Certs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCerts

`func (o *ConnectionCredentialsInboundBackChannelAuthAllOf) SetCerts(v []ConnectionCert)`

SetCerts sets Certs field to given value.

### HasCerts

`func (o *ConnectionCredentialsInboundBackChannelAuthAllOf) HasCerts() bool`

HasCerts returns a boolean if a field has been set.

### GetRequireSsl

`func (o *ConnectionCredentialsInboundBackChannelAuthAllOf) GetRequireSsl() interface{}`

GetRequireSsl returns the RequireSsl field if non-nil, zero value otherwise.

### GetRequireSslOk

`func (o *ConnectionCredentialsInboundBackChannelAuthAllOf) GetRequireSslOk() (*interface{}, bool)`

GetRequireSslOk returns a tuple with the RequireSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSsl

`func (o *ConnectionCredentialsInboundBackChannelAuthAllOf) SetRequireSsl(v interface{})`

SetRequireSsl sets RequireSsl field to given value.

### HasRequireSsl

`func (o *ConnectionCredentialsInboundBackChannelAuthAllOf) HasRequireSsl() bool`

HasRequireSsl returns a boolean if a field has been set.

### SetRequireSslNil

`func (o *ConnectionCredentialsInboundBackChannelAuthAllOf) SetRequireSslNil(b bool)`

 SetRequireSslNil sets the value for RequireSsl to be an explicit nil

### UnsetRequireSsl
`func (o *ConnectionCredentialsInboundBackChannelAuthAllOf) UnsetRequireSsl()`

UnsetRequireSsl ensures that no value is present for RequireSsl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


