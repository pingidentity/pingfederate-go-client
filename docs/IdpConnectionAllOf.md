# IdpConnectionAllOf

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

### NewIdpConnectionAllOf

`func NewIdpConnectionAllOf() *IdpConnectionAllOf`

NewIdpConnectionAllOf instantiates a new IdpConnectionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpConnectionAllOfWithDefaults

`func NewIdpConnectionAllOfWithDefaults() *IdpConnectionAllOf`

NewIdpConnectionAllOfWithDefaults instantiates a new IdpConnectionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOidcClientCredentials

`func (o *IdpConnectionAllOf) GetOidcClientCredentials() OIDCClientCredentials`

GetOidcClientCredentials returns the OidcClientCredentials field if non-nil, zero value otherwise.

### GetOidcClientCredentialsOk

`func (o *IdpConnectionAllOf) GetOidcClientCredentialsOk() (*OIDCClientCredentials, bool)`

GetOidcClientCredentialsOk returns a tuple with the OidcClientCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcClientCredentials

`func (o *IdpConnectionAllOf) SetOidcClientCredentials(v OIDCClientCredentials)`

SetOidcClientCredentials sets OidcClientCredentials field to given value.

### HasOidcClientCredentials

`func (o *IdpConnectionAllOf) HasOidcClientCredentials() bool`

HasOidcClientCredentials returns a boolean if a field has been set.

### GetIdpBrowserSso

`func (o *IdpConnectionAllOf) GetIdpBrowserSso() IdpBrowserSso`

GetIdpBrowserSso returns the IdpBrowserSso field if non-nil, zero value otherwise.

### GetIdpBrowserSsoOk

`func (o *IdpConnectionAllOf) GetIdpBrowserSsoOk() (*IdpBrowserSso, bool)`

GetIdpBrowserSsoOk returns a tuple with the IdpBrowserSso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpBrowserSso

`func (o *IdpConnectionAllOf) SetIdpBrowserSso(v IdpBrowserSso)`

SetIdpBrowserSso sets IdpBrowserSso field to given value.

### HasIdpBrowserSso

`func (o *IdpConnectionAllOf) HasIdpBrowserSso() bool`

HasIdpBrowserSso returns a boolean if a field has been set.

### GetAttributeQuery

`func (o *IdpConnectionAllOf) GetAttributeQuery() IdpAttributeQuery`

GetAttributeQuery returns the AttributeQuery field if non-nil, zero value otherwise.

### GetAttributeQueryOk

`func (o *IdpConnectionAllOf) GetAttributeQueryOk() (*IdpAttributeQuery, bool)`

GetAttributeQueryOk returns a tuple with the AttributeQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeQuery

`func (o *IdpConnectionAllOf) SetAttributeQuery(v IdpAttributeQuery)`

SetAttributeQuery sets AttributeQuery field to given value.

### HasAttributeQuery

`func (o *IdpConnectionAllOf) HasAttributeQuery() bool`

HasAttributeQuery returns a boolean if a field has been set.

### GetIdpOAuthGrantAttributeMapping

`func (o *IdpConnectionAllOf) GetIdpOAuthGrantAttributeMapping() IdpOAuthGrantAttributeMapping`

GetIdpOAuthGrantAttributeMapping returns the IdpOAuthGrantAttributeMapping field if non-nil, zero value otherwise.

### GetIdpOAuthGrantAttributeMappingOk

`func (o *IdpConnectionAllOf) GetIdpOAuthGrantAttributeMappingOk() (*IdpOAuthGrantAttributeMapping, bool)`

GetIdpOAuthGrantAttributeMappingOk returns a tuple with the IdpOAuthGrantAttributeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpOAuthGrantAttributeMapping

`func (o *IdpConnectionAllOf) SetIdpOAuthGrantAttributeMapping(v IdpOAuthGrantAttributeMapping)`

SetIdpOAuthGrantAttributeMapping sets IdpOAuthGrantAttributeMapping field to given value.

### HasIdpOAuthGrantAttributeMapping

`func (o *IdpConnectionAllOf) HasIdpOAuthGrantAttributeMapping() bool`

HasIdpOAuthGrantAttributeMapping returns a boolean if a field has been set.

### GetWsTrust

`func (o *IdpConnectionAllOf) GetWsTrust() IdpWsTrust`

GetWsTrust returns the WsTrust field if non-nil, zero value otherwise.

### GetWsTrustOk

`func (o *IdpConnectionAllOf) GetWsTrustOk() (*IdpWsTrust, bool)`

GetWsTrustOk returns a tuple with the WsTrust field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWsTrust

`func (o *IdpConnectionAllOf) SetWsTrust(v IdpWsTrust)`

SetWsTrust sets WsTrust field to given value.

### HasWsTrust

`func (o *IdpConnectionAllOf) HasWsTrust() bool`

HasWsTrust returns a boolean if a field has been set.

### GetInboundProvisioning

`func (o *IdpConnectionAllOf) GetInboundProvisioning() IdpInboundProvisioning`

GetInboundProvisioning returns the InboundProvisioning field if non-nil, zero value otherwise.

### GetInboundProvisioningOk

`func (o *IdpConnectionAllOf) GetInboundProvisioningOk() (*IdpInboundProvisioning, bool)`

GetInboundProvisioningOk returns a tuple with the InboundProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundProvisioning

`func (o *IdpConnectionAllOf) SetInboundProvisioning(v IdpInboundProvisioning)`

SetInboundProvisioning sets InboundProvisioning field to given value.

### HasInboundProvisioning

`func (o *IdpConnectionAllOf) HasInboundProvisioning() bool`

HasInboundProvisioning returns a boolean if a field has been set.

### GetErrorPageMsgId

`func (o *IdpConnectionAllOf) GetErrorPageMsgId() string`

GetErrorPageMsgId returns the ErrorPageMsgId field if non-nil, zero value otherwise.

### GetErrorPageMsgIdOk

`func (o *IdpConnectionAllOf) GetErrorPageMsgIdOk() (*string, bool)`

GetErrorPageMsgIdOk returns a tuple with the ErrorPageMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorPageMsgId

`func (o *IdpConnectionAllOf) SetErrorPageMsgId(v string)`

SetErrorPageMsgId sets ErrorPageMsgId field to given value.

### HasErrorPageMsgId

`func (o *IdpConnectionAllOf) HasErrorPageMsgId() bool`

HasErrorPageMsgId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


