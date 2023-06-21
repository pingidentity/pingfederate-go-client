# FederationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseUrl** | Pointer to **string** | The fully qualified host name, port, and path (if applicable) on which the PingFederate server runs. | [optional] 
**Saml2EntityId** | Pointer to **string** | This ID defines your organization as the entity operating the server for SAML 2.0 transactions. It is usually defined as an organization&#39;s URL or a DNS address; for example: pingidentity.com. The SAML SourceID used for artifact resolution is derived from this ID using SHA1. | [optional] 
**AutoConnectEntityId** | Pointer to **string** | This property has been deprecated and no longer used | [optional] 
**Saml1xIssuerId** | Pointer to **string** | This ID identifies your federation server for SAML 1.x transactions. As with SAML 2.0, it is usually defined as an organization&#39;s URL or a DNS address. The SourceID used for artifact resolution is derived from this ID using SHA1. | [optional] 
**Saml1xSourceId** | Pointer to **string** | If supplied, the Source ID value entered here is used for SAML 1.x, instead of being derived from the SAML 1.x Issuer/Audience. | [optional] 
**WsfedRealm** | Pointer to **string** | The URI of the realm associated with the PingFederate server. A realm represents a single unit of security administration or trust. | [optional] 

## Methods

### NewFederationInfo

`func NewFederationInfo() *FederationInfo`

NewFederationInfo instantiates a new FederationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFederationInfoWithDefaults

`func NewFederationInfoWithDefaults() *FederationInfo`

NewFederationInfoWithDefaults instantiates a new FederationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseUrl

`func (o *FederationInfo) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *FederationInfo) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *FederationInfo) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *FederationInfo) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetSaml2EntityId

`func (o *FederationInfo) GetSaml2EntityId() string`

GetSaml2EntityId returns the Saml2EntityId field if non-nil, zero value otherwise.

### GetSaml2EntityIdOk

`func (o *FederationInfo) GetSaml2EntityIdOk() (*string, bool)`

GetSaml2EntityIdOk returns a tuple with the Saml2EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml2EntityId

`func (o *FederationInfo) SetSaml2EntityId(v string)`

SetSaml2EntityId sets Saml2EntityId field to given value.

### HasSaml2EntityId

`func (o *FederationInfo) HasSaml2EntityId() bool`

HasSaml2EntityId returns a boolean if a field has been set.

### GetAutoConnectEntityId

`func (o *FederationInfo) GetAutoConnectEntityId() string`

GetAutoConnectEntityId returns the AutoConnectEntityId field if non-nil, zero value otherwise.

### GetAutoConnectEntityIdOk

`func (o *FederationInfo) GetAutoConnectEntityIdOk() (*string, bool)`

GetAutoConnectEntityIdOk returns a tuple with the AutoConnectEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoConnectEntityId

`func (o *FederationInfo) SetAutoConnectEntityId(v string)`

SetAutoConnectEntityId sets AutoConnectEntityId field to given value.

### HasAutoConnectEntityId

`func (o *FederationInfo) HasAutoConnectEntityId() bool`

HasAutoConnectEntityId returns a boolean if a field has been set.

### GetSaml1xIssuerId

`func (o *FederationInfo) GetSaml1xIssuerId() string`

GetSaml1xIssuerId returns the Saml1xIssuerId field if non-nil, zero value otherwise.

### GetSaml1xIssuerIdOk

`func (o *FederationInfo) GetSaml1xIssuerIdOk() (*string, bool)`

GetSaml1xIssuerIdOk returns a tuple with the Saml1xIssuerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml1xIssuerId

`func (o *FederationInfo) SetSaml1xIssuerId(v string)`

SetSaml1xIssuerId sets Saml1xIssuerId field to given value.

### HasSaml1xIssuerId

`func (o *FederationInfo) HasSaml1xIssuerId() bool`

HasSaml1xIssuerId returns a boolean if a field has been set.

### GetSaml1xSourceId

`func (o *FederationInfo) GetSaml1xSourceId() string`

GetSaml1xSourceId returns the Saml1xSourceId field if non-nil, zero value otherwise.

### GetSaml1xSourceIdOk

`func (o *FederationInfo) GetSaml1xSourceIdOk() (*string, bool)`

GetSaml1xSourceIdOk returns a tuple with the Saml1xSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml1xSourceId

`func (o *FederationInfo) SetSaml1xSourceId(v string)`

SetSaml1xSourceId sets Saml1xSourceId field to given value.

### HasSaml1xSourceId

`func (o *FederationInfo) HasSaml1xSourceId() bool`

HasSaml1xSourceId returns a boolean if a field has been set.

### GetWsfedRealm

`func (o *FederationInfo) GetWsfedRealm() string`

GetWsfedRealm returns the WsfedRealm field if non-nil, zero value otherwise.

### GetWsfedRealmOk

`func (o *FederationInfo) GetWsfedRealmOk() (*string, bool)`

GetWsfedRealmOk returns a tuple with the WsfedRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWsfedRealm

`func (o *FederationInfo) SetWsfedRealm(v string)`

SetWsfedRealm sets WsfedRealm field to given value.

### HasWsfedRealm

`func (o *FederationInfo) HasWsfedRealm() bool`

HasWsfedRealm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


