# IdpConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OidcClientCredentials** | Pointer to [**OIDCClientCredentials**](OIDCClientCredentials.md) |  | [optional] 
**IdpBrowserSso** | Pointer to [**IdpBrowserSso**](IdpBrowserSso.md) |  | [optional] 
**AttributeQuery** | Pointer to [**IdpAttributeQuery**](IdpAttributeQuery.md) |  | [optional] 
**IdpOAuthGrantAttributeMapping** | Pointer to [**IdpOAuthGrantAttributeMapping**](IdpOAuthGrantAttributeMapping.md) |  | [optional] 
**WsTrust** | Pointer to [**IdpWsTrust**](IdpWsTrust.md) |  | [optional] 
**InboundProvisioning** | Pointer to [**IdpInboundProvisioning**](IdpInboundProvisioning.md) |  | [optional] 
**ErrorPageMsgId** | Pointer to **string** | Identifier that specifies the message displayed on a user-facing error page. | [optional] 

## Methods

### NewIdpConnection

`func NewIdpConnection() *IdpConnection`

NewIdpConnection instantiates a new IdpConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpConnectionWithDefaults

`func NewIdpConnectionWithDefaults() *IdpConnection`

NewIdpConnectionWithDefaults instantiates a new IdpConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOidcClientCredentials

`func (o *IdpConnection) GetOidcClientCredentials() OIDCClientCredentials`

GetOidcClientCredentials returns the OidcClientCredentials field if non-nil, zero value otherwise.

### GetOidcClientCredentialsOk

`func (o *IdpConnection) GetOidcClientCredentialsOk() (*OIDCClientCredentials, bool)`

GetOidcClientCredentialsOk returns a tuple with the OidcClientCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcClientCredentials

`func (o *IdpConnection) SetOidcClientCredentials(v OIDCClientCredentials)`

SetOidcClientCredentials sets OidcClientCredentials field to given value.

### HasOidcClientCredentials

`func (o *IdpConnection) HasOidcClientCredentials() bool`

HasOidcClientCredentials returns a boolean if a field has been set.

### GetIdpBrowserSso

`func (o *IdpConnection) GetIdpBrowserSso() IdpBrowserSso`

GetIdpBrowserSso returns the IdpBrowserSso field if non-nil, zero value otherwise.

### GetIdpBrowserSsoOk

`func (o *IdpConnection) GetIdpBrowserSsoOk() (*IdpBrowserSso, bool)`

GetIdpBrowserSsoOk returns a tuple with the IdpBrowserSso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpBrowserSso

`func (o *IdpConnection) SetIdpBrowserSso(v IdpBrowserSso)`

SetIdpBrowserSso sets IdpBrowserSso field to given value.

### HasIdpBrowserSso

`func (o *IdpConnection) HasIdpBrowserSso() bool`

HasIdpBrowserSso returns a boolean if a field has been set.

### GetAttributeQuery

`func (o *IdpConnection) GetAttributeQuery() IdpAttributeQuery`

GetAttributeQuery returns the AttributeQuery field if non-nil, zero value otherwise.

### GetAttributeQueryOk

`func (o *IdpConnection) GetAttributeQueryOk() (*IdpAttributeQuery, bool)`

GetAttributeQueryOk returns a tuple with the AttributeQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeQuery

`func (o *IdpConnection) SetAttributeQuery(v IdpAttributeQuery)`

SetAttributeQuery sets AttributeQuery field to given value.

### HasAttributeQuery

`func (o *IdpConnection) HasAttributeQuery() bool`

HasAttributeQuery returns a boolean if a field has been set.

### GetIdpOAuthGrantAttributeMapping

`func (o *IdpConnection) GetIdpOAuthGrantAttributeMapping() IdpOAuthGrantAttributeMapping`

GetIdpOAuthGrantAttributeMapping returns the IdpOAuthGrantAttributeMapping field if non-nil, zero value otherwise.

### GetIdpOAuthGrantAttributeMappingOk

`func (o *IdpConnection) GetIdpOAuthGrantAttributeMappingOk() (*IdpOAuthGrantAttributeMapping, bool)`

GetIdpOAuthGrantAttributeMappingOk returns a tuple with the IdpOAuthGrantAttributeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpOAuthGrantAttributeMapping

`func (o *IdpConnection) SetIdpOAuthGrantAttributeMapping(v IdpOAuthGrantAttributeMapping)`

SetIdpOAuthGrantAttributeMapping sets IdpOAuthGrantAttributeMapping field to given value.

### HasIdpOAuthGrantAttributeMapping

`func (o *IdpConnection) HasIdpOAuthGrantAttributeMapping() bool`

HasIdpOAuthGrantAttributeMapping returns a boolean if a field has been set.

### GetWsTrust

`func (o *IdpConnection) GetWsTrust() IdpWsTrust`

GetWsTrust returns the WsTrust field if non-nil, zero value otherwise.

### GetWsTrustOk

`func (o *IdpConnection) GetWsTrustOk() (*IdpWsTrust, bool)`

GetWsTrustOk returns a tuple with the WsTrust field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWsTrust

`func (o *IdpConnection) SetWsTrust(v IdpWsTrust)`

SetWsTrust sets WsTrust field to given value.

### HasWsTrust

`func (o *IdpConnection) HasWsTrust() bool`

HasWsTrust returns a boolean if a field has been set.

### GetInboundProvisioning

`func (o *IdpConnection) GetInboundProvisioning() IdpInboundProvisioning`

GetInboundProvisioning returns the InboundProvisioning field if non-nil, zero value otherwise.

### GetInboundProvisioningOk

`func (o *IdpConnection) GetInboundProvisioningOk() (*IdpInboundProvisioning, bool)`

GetInboundProvisioningOk returns a tuple with the InboundProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundProvisioning

`func (o *IdpConnection) SetInboundProvisioning(v IdpInboundProvisioning)`

SetInboundProvisioning sets InboundProvisioning field to given value.

### HasInboundProvisioning

`func (o *IdpConnection) HasInboundProvisioning() bool`

HasInboundProvisioning returns a boolean if a field has been set.

### GetErrorPageMsgId

`func (o *IdpConnection) GetErrorPageMsgId() string`

GetErrorPageMsgId returns the ErrorPageMsgId field if non-nil, zero value otherwise.

### GetErrorPageMsgIdOk

`func (o *IdpConnection) GetErrorPageMsgIdOk() (*string, bool)`

GetErrorPageMsgIdOk returns a tuple with the ErrorPageMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorPageMsgId

`func (o *IdpConnection) SetErrorPageMsgId(v string)`

SetErrorPageMsgId sets ErrorPageMsgId field to given value.

### HasErrorPageMsgId

`func (o *IdpConnection) HasErrorPageMsgId() bool`

HasErrorPageMsgId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


