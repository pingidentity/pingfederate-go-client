# IdpConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of this connection. Default is &#39;IDP&#39;. | [optional] 
**Id** | Pointer to **string** | The persistent, unique ID for the connection. It can be any combination of [a-zA-Z0-9._-]. This property is system-assigned if not specified. | [optional] 
**EntityId** | **string** | The partner&#39;s entity ID (connection ID) or issuer value (for OIDC Connections). | 
**Name** | **string** | The connection name. | 
**ModificationDate** | Pointer to **time.Time** | The time at which the connection was last changed. This property is read only and is ignored on PUT and POST requests. | [optional] 
**CreationDate** | Pointer to **time.Time** | The time at which the connection was created. This property is read only and is ignored on PUT and POST requests. | [optional] 
**Active** | Pointer to **bool** | Specifies whether the connection is active and ready to process incoming requests. The default value is false. | [optional] 
**BaseUrl** | Pointer to **string** | The fully-qualified hostname and port on which your partner&#39;s federation deployment runs. | [optional] 
**DefaultVirtualEntityId** | Pointer to **string** | The default alternate entity ID that identifies the local server to this partner. It is required when virtualEntityIds is not empty and must be included in that list. | [optional] 
**VirtualEntityIds** | Pointer to **[]string** | List of alternate entity IDs that identifies the local server to this partner. | [optional] 
**MetadataReloadSettings** | Pointer to [**ConnectionMetadataUrl**](ConnectionMetadataUrl.md) |  | [optional] 
**Credentials** | Pointer to [**ConnectionCredentials**](ConnectionCredentials.md) |  | [optional] 
**ContactInfo** | Pointer to [**ContactInfo**](ContactInfo.md) |  | [optional] 
**LicenseConnectionGroup** | Pointer to **string** | The license connection group. If your PingFederate license is based on connection groups, each connection must be assigned to a group before it can be used. | [optional] 
**LoggingMode** | Pointer to **string** | The level of transaction logging applicable for this connection. Default is STANDARD. | [optional] 
**AdditionalAllowedEntitiesConfiguration** | Pointer to [**AdditionalAllowedEntitiesConfiguration**](AdditionalAllowedEntitiesConfiguration.md) |  | [optional] 
**ExtendedProperties** | Pointer to [**map[string]ParameterValues**](ParameterValues.md) | Extended Properties allows to store additional information for IdP/SP Connections. The names of these extended properties should be defined in /extendedProperties. | [optional] 
**ReplicationStatus** | Pointer to **string** | This status indicates whether the connection has been replicated to the cluster. This property only applies when automatic replication of connections is enabled. It is read only and is ignored on PUT and POST requests. | [optional] 
**OidcClientCredentials** | Pointer to [**OIDCClientCredentials**](OIDCClientCredentials.md) |  | [optional] 
**IdpBrowserSso** | Pointer to [**IdpBrowserSso**](IdpBrowserSso.md) |  | [optional] 
**AttributeQuery** | Pointer to [**IdpAttributeQuery**](IdpAttributeQuery.md) |  | [optional] 
**IdpOAuthGrantAttributeMapping** | Pointer to [**IdpOAuthGrantAttributeMapping**](IdpOAuthGrantAttributeMapping.md) |  | [optional] 
**WsTrust** | Pointer to [**IdpWsTrust**](IdpWsTrust.md) |  | [optional] 
**InboundProvisioning** | Pointer to [**IdpInboundProvisioning**](IdpInboundProvisioning.md) |  | [optional] 
**ErrorPageMsgId** | Pointer to **string** | Identifier that specifies the message displayed on a user-facing error page. | [optional] 

## Methods

### NewIdpConnection

`func NewIdpConnection(entityId string, name string, ) *IdpConnection`

NewIdpConnection instantiates a new IdpConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpConnectionWithDefaults

`func NewIdpConnectionWithDefaults() *IdpConnection`

NewIdpConnectionWithDefaults instantiates a new IdpConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IdpConnection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdpConnection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdpConnection) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IdpConnection) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *IdpConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdpConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdpConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdpConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEntityId

`func (o *IdpConnection) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *IdpConnection) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *IdpConnection) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetName

`func (o *IdpConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdpConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdpConnection) SetName(v string)`

SetName sets Name field to given value.


### GetModificationDate

`func (o *IdpConnection) GetModificationDate() time.Time`

GetModificationDate returns the ModificationDate field if non-nil, zero value otherwise.

### GetModificationDateOk

`func (o *IdpConnection) GetModificationDateOk() (*time.Time, bool)`

GetModificationDateOk returns a tuple with the ModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationDate

`func (o *IdpConnection) SetModificationDate(v time.Time)`

SetModificationDate sets ModificationDate field to given value.

### HasModificationDate

`func (o *IdpConnection) HasModificationDate() bool`

HasModificationDate returns a boolean if a field has been set.

### GetCreationDate

`func (o *IdpConnection) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *IdpConnection) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *IdpConnection) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *IdpConnection) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetActive

`func (o *IdpConnection) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *IdpConnection) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *IdpConnection) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *IdpConnection) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetBaseUrl

`func (o *IdpConnection) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *IdpConnection) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *IdpConnection) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *IdpConnection) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetDefaultVirtualEntityId

`func (o *IdpConnection) GetDefaultVirtualEntityId() string`

GetDefaultVirtualEntityId returns the DefaultVirtualEntityId field if non-nil, zero value otherwise.

### GetDefaultVirtualEntityIdOk

`func (o *IdpConnection) GetDefaultVirtualEntityIdOk() (*string, bool)`

GetDefaultVirtualEntityIdOk returns a tuple with the DefaultVirtualEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVirtualEntityId

`func (o *IdpConnection) SetDefaultVirtualEntityId(v string)`

SetDefaultVirtualEntityId sets DefaultVirtualEntityId field to given value.

### HasDefaultVirtualEntityId

`func (o *IdpConnection) HasDefaultVirtualEntityId() bool`

HasDefaultVirtualEntityId returns a boolean if a field has been set.

### GetVirtualEntityIds

`func (o *IdpConnection) GetVirtualEntityIds() []string`

GetVirtualEntityIds returns the VirtualEntityIds field if non-nil, zero value otherwise.

### GetVirtualEntityIdsOk

`func (o *IdpConnection) GetVirtualEntityIdsOk() (*[]string, bool)`

GetVirtualEntityIdsOk returns a tuple with the VirtualEntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualEntityIds

`func (o *IdpConnection) SetVirtualEntityIds(v []string)`

SetVirtualEntityIds sets VirtualEntityIds field to given value.

### HasVirtualEntityIds

`func (o *IdpConnection) HasVirtualEntityIds() bool`

HasVirtualEntityIds returns a boolean if a field has been set.

### GetMetadataReloadSettings

`func (o *IdpConnection) GetMetadataReloadSettings() ConnectionMetadataUrl`

GetMetadataReloadSettings returns the MetadataReloadSettings field if non-nil, zero value otherwise.

### GetMetadataReloadSettingsOk

`func (o *IdpConnection) GetMetadataReloadSettingsOk() (*ConnectionMetadataUrl, bool)`

GetMetadataReloadSettingsOk returns a tuple with the MetadataReloadSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataReloadSettings

`func (o *IdpConnection) SetMetadataReloadSettings(v ConnectionMetadataUrl)`

SetMetadataReloadSettings sets MetadataReloadSettings field to given value.

### HasMetadataReloadSettings

`func (o *IdpConnection) HasMetadataReloadSettings() bool`

HasMetadataReloadSettings returns a boolean if a field has been set.

### GetCredentials

`func (o *IdpConnection) GetCredentials() ConnectionCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *IdpConnection) GetCredentialsOk() (*ConnectionCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *IdpConnection) SetCredentials(v ConnectionCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *IdpConnection) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetContactInfo

`func (o *IdpConnection) GetContactInfo() ContactInfo`

GetContactInfo returns the ContactInfo field if non-nil, zero value otherwise.

### GetContactInfoOk

`func (o *IdpConnection) GetContactInfoOk() (*ContactInfo, bool)`

GetContactInfoOk returns a tuple with the ContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactInfo

`func (o *IdpConnection) SetContactInfo(v ContactInfo)`

SetContactInfo sets ContactInfo field to given value.

### HasContactInfo

`func (o *IdpConnection) HasContactInfo() bool`

HasContactInfo returns a boolean if a field has been set.

### GetLicenseConnectionGroup

`func (o *IdpConnection) GetLicenseConnectionGroup() string`

GetLicenseConnectionGroup returns the LicenseConnectionGroup field if non-nil, zero value otherwise.

### GetLicenseConnectionGroupOk

`func (o *IdpConnection) GetLicenseConnectionGroupOk() (*string, bool)`

GetLicenseConnectionGroupOk returns a tuple with the LicenseConnectionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseConnectionGroup

`func (o *IdpConnection) SetLicenseConnectionGroup(v string)`

SetLicenseConnectionGroup sets LicenseConnectionGroup field to given value.

### HasLicenseConnectionGroup

`func (o *IdpConnection) HasLicenseConnectionGroup() bool`

HasLicenseConnectionGroup returns a boolean if a field has been set.

### GetLoggingMode

`func (o *IdpConnection) GetLoggingMode() string`

GetLoggingMode returns the LoggingMode field if non-nil, zero value otherwise.

### GetLoggingModeOk

`func (o *IdpConnection) GetLoggingModeOk() (*string, bool)`

GetLoggingModeOk returns a tuple with the LoggingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingMode

`func (o *IdpConnection) SetLoggingMode(v string)`

SetLoggingMode sets LoggingMode field to given value.

### HasLoggingMode

`func (o *IdpConnection) HasLoggingMode() bool`

HasLoggingMode returns a boolean if a field has been set.

### GetAdditionalAllowedEntitiesConfiguration

`func (o *IdpConnection) GetAdditionalAllowedEntitiesConfiguration() AdditionalAllowedEntitiesConfiguration`

GetAdditionalAllowedEntitiesConfiguration returns the AdditionalAllowedEntitiesConfiguration field if non-nil, zero value otherwise.

### GetAdditionalAllowedEntitiesConfigurationOk

`func (o *IdpConnection) GetAdditionalAllowedEntitiesConfigurationOk() (*AdditionalAllowedEntitiesConfiguration, bool)`

GetAdditionalAllowedEntitiesConfigurationOk returns a tuple with the AdditionalAllowedEntitiesConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAllowedEntitiesConfiguration

`func (o *IdpConnection) SetAdditionalAllowedEntitiesConfiguration(v AdditionalAllowedEntitiesConfiguration)`

SetAdditionalAllowedEntitiesConfiguration sets AdditionalAllowedEntitiesConfiguration field to given value.

### HasAdditionalAllowedEntitiesConfiguration

`func (o *IdpConnection) HasAdditionalAllowedEntitiesConfiguration() bool`

HasAdditionalAllowedEntitiesConfiguration returns a boolean if a field has been set.

### GetExtendedProperties

`func (o *IdpConnection) GetExtendedProperties() map[string]ParameterValues`

GetExtendedProperties returns the ExtendedProperties field if non-nil, zero value otherwise.

### GetExtendedPropertiesOk

`func (o *IdpConnection) GetExtendedPropertiesOk() (*map[string]ParameterValues, bool)`

GetExtendedPropertiesOk returns a tuple with the ExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedProperties

`func (o *IdpConnection) SetExtendedProperties(v map[string]ParameterValues)`

SetExtendedProperties sets ExtendedProperties field to given value.

### HasExtendedProperties

`func (o *IdpConnection) HasExtendedProperties() bool`

HasExtendedProperties returns a boolean if a field has been set.

### GetReplicationStatus

`func (o *IdpConnection) GetReplicationStatus() string`

GetReplicationStatus returns the ReplicationStatus field if non-nil, zero value otherwise.

### GetReplicationStatusOk

`func (o *IdpConnection) GetReplicationStatusOk() (*string, bool)`

GetReplicationStatusOk returns a tuple with the ReplicationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStatus

`func (o *IdpConnection) SetReplicationStatus(v string)`

SetReplicationStatus sets ReplicationStatus field to given value.

### HasReplicationStatus

`func (o *IdpConnection) HasReplicationStatus() bool`

HasReplicationStatus returns a boolean if a field has been set.

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


