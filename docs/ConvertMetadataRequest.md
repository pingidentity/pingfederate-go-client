# ConvertMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | **string** | The expected connection type to convert. | 
**ExpectedProtocol** | **string** | The expected browser-based SSO protocol to convert. In this case the protocol is restricted to SAML. | 
**ExpectedEntityId** | Pointer to **string** | The entity ID of the connection to be obtained from the input SAML Metadata. Required if the SAML Metadata has more than one connection defined. | [optional] 
**SamlMetadata** | **string** | The base-64 encoded XML SAML metadata. | 
**VerificationCertificate** | Pointer to **string** | The certificate to validate the metadata signature against. The certificate can be in PEM format or base-64 encoded DER format. | [optional] 
**TemplateConnection** | Pointer to [**Connection**](Connection.md) |  | [optional] 

## Methods

### NewConvertMetadataRequest

`func NewConvertMetadataRequest(connectionType string, expectedProtocol string, samlMetadata string, ) *ConvertMetadataRequest`

NewConvertMetadataRequest instantiates a new ConvertMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvertMetadataRequestWithDefaults

`func NewConvertMetadataRequestWithDefaults() *ConvertMetadataRequest`

NewConvertMetadataRequestWithDefaults instantiates a new ConvertMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionType

`func (o *ConvertMetadataRequest) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *ConvertMetadataRequest) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *ConvertMetadataRequest) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetExpectedProtocol

`func (o *ConvertMetadataRequest) GetExpectedProtocol() string`

GetExpectedProtocol returns the ExpectedProtocol field if non-nil, zero value otherwise.

### GetExpectedProtocolOk

`func (o *ConvertMetadataRequest) GetExpectedProtocolOk() (*string, bool)`

GetExpectedProtocolOk returns a tuple with the ExpectedProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedProtocol

`func (o *ConvertMetadataRequest) SetExpectedProtocol(v string)`

SetExpectedProtocol sets ExpectedProtocol field to given value.


### GetExpectedEntityId

`func (o *ConvertMetadataRequest) GetExpectedEntityId() string`

GetExpectedEntityId returns the ExpectedEntityId field if non-nil, zero value otherwise.

### GetExpectedEntityIdOk

`func (o *ConvertMetadataRequest) GetExpectedEntityIdOk() (*string, bool)`

GetExpectedEntityIdOk returns a tuple with the ExpectedEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedEntityId

`func (o *ConvertMetadataRequest) SetExpectedEntityId(v string)`

SetExpectedEntityId sets ExpectedEntityId field to given value.

### HasExpectedEntityId

`func (o *ConvertMetadataRequest) HasExpectedEntityId() bool`

HasExpectedEntityId returns a boolean if a field has been set.

### GetSamlMetadata

`func (o *ConvertMetadataRequest) GetSamlMetadata() string`

GetSamlMetadata returns the SamlMetadata field if non-nil, zero value otherwise.

### GetSamlMetadataOk

`func (o *ConvertMetadataRequest) GetSamlMetadataOk() (*string, bool)`

GetSamlMetadataOk returns a tuple with the SamlMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlMetadata

`func (o *ConvertMetadataRequest) SetSamlMetadata(v string)`

SetSamlMetadata sets SamlMetadata field to given value.


### GetVerificationCertificate

`func (o *ConvertMetadataRequest) GetVerificationCertificate() string`

GetVerificationCertificate returns the VerificationCertificate field if non-nil, zero value otherwise.

### GetVerificationCertificateOk

`func (o *ConvertMetadataRequest) GetVerificationCertificateOk() (*string, bool)`

GetVerificationCertificateOk returns a tuple with the VerificationCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationCertificate

`func (o *ConvertMetadataRequest) SetVerificationCertificate(v string)`

SetVerificationCertificate sets VerificationCertificate field to given value.

### HasVerificationCertificate

`func (o *ConvertMetadataRequest) HasVerificationCertificate() bool`

HasVerificationCertificate returns a boolean if a field has been set.

### GetTemplateConnection

`func (o *ConvertMetadataRequest) GetTemplateConnection() Connection`

GetTemplateConnection returns the TemplateConnection field if non-nil, zero value otherwise.

### GetTemplateConnectionOk

`func (o *ConvertMetadataRequest) GetTemplateConnectionOk() (*Connection, bool)`

GetTemplateConnectionOk returns a tuple with the TemplateConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateConnection

`func (o *ConvertMetadataRequest) SetTemplateConnection(v Connection)`

SetTemplateConnection sets TemplateConnection field to given value.

### HasTemplateConnection

`func (o *ConvertMetadataRequest) HasTemplateConnection() bool`

HasTemplateConnection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


