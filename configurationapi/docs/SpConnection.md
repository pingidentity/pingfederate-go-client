# SpConnection

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
**SpBrowserSso** | Pointer to [**SpBrowserSso**](SpBrowserSso.md) |  | [optional] 
**AttributeQuery** | Pointer to [**SpAttributeQuery**](SpAttributeQuery.md) |  | [optional] 
**WsTrust** | Pointer to [**SpWsTrust**](SpWsTrust.md) |  | [optional] 
**ApplicationName** | Pointer to **string** | The application name. | [optional] 
**ApplicationIconUrl** | Pointer to **string** | The application icon url. | [optional] 
**OutboundProvision** | Pointer to [**OutboundProvision**](OutboundProvision.md) |  | [optional] 
**ConnectionTargetType** | Pointer to **string** | The connection target type. This field is intended for bulk import/export usage. Changing its value may result in unexpected behavior. | [optional] 

## Methods

### NewSpConnection

`func NewSpConnection(entityId string, name string, ) *SpConnection`

NewSpConnection instantiates a new SpConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpConnectionWithDefaults

`func NewSpConnectionWithDefaults() *SpConnection`

NewSpConnectionWithDefaults instantiates a new SpConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SpConnection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SpConnection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SpConnection) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SpConnection) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *SpConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SpConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SpConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEntityId

`func (o *SpConnection) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *SpConnection) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *SpConnection) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetName

`func (o *SpConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpConnection) SetName(v string)`

SetName sets Name field to given value.


### GetModificationDate

`func (o *SpConnection) GetModificationDate() time.Time`

GetModificationDate returns the ModificationDate field if non-nil, zero value otherwise.

### GetModificationDateOk

`func (o *SpConnection) GetModificationDateOk() (*time.Time, bool)`

GetModificationDateOk returns a tuple with the ModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationDate

`func (o *SpConnection) SetModificationDate(v time.Time)`

SetModificationDate sets ModificationDate field to given value.

### HasModificationDate

`func (o *SpConnection) HasModificationDate() bool`

HasModificationDate returns a boolean if a field has been set.

### GetCreationDate

`func (o *SpConnection) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *SpConnection) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *SpConnection) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *SpConnection) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetActive

`func (o *SpConnection) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SpConnection) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SpConnection) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *SpConnection) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetBaseUrl

`func (o *SpConnection) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *SpConnection) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *SpConnection) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *SpConnection) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetDefaultVirtualEntityId

`func (o *SpConnection) GetDefaultVirtualEntityId() string`

GetDefaultVirtualEntityId returns the DefaultVirtualEntityId field if non-nil, zero value otherwise.

### GetDefaultVirtualEntityIdOk

`func (o *SpConnection) GetDefaultVirtualEntityIdOk() (*string, bool)`

GetDefaultVirtualEntityIdOk returns a tuple with the DefaultVirtualEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVirtualEntityId

`func (o *SpConnection) SetDefaultVirtualEntityId(v string)`

SetDefaultVirtualEntityId sets DefaultVirtualEntityId field to given value.

### HasDefaultVirtualEntityId

`func (o *SpConnection) HasDefaultVirtualEntityId() bool`

HasDefaultVirtualEntityId returns a boolean if a field has been set.

### GetVirtualEntityIds

`func (o *SpConnection) GetVirtualEntityIds() []string`

GetVirtualEntityIds returns the VirtualEntityIds field if non-nil, zero value otherwise.

### GetVirtualEntityIdsOk

`func (o *SpConnection) GetVirtualEntityIdsOk() (*[]string, bool)`

GetVirtualEntityIdsOk returns a tuple with the VirtualEntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualEntityIds

`func (o *SpConnection) SetVirtualEntityIds(v []string)`

SetVirtualEntityIds sets VirtualEntityIds field to given value.

### HasVirtualEntityIds

`func (o *SpConnection) HasVirtualEntityIds() bool`

HasVirtualEntityIds returns a boolean if a field has been set.

### GetMetadataReloadSettings

`func (o *SpConnection) GetMetadataReloadSettings() ConnectionMetadataUrl`

GetMetadataReloadSettings returns the MetadataReloadSettings field if non-nil, zero value otherwise.

### GetMetadataReloadSettingsOk

`func (o *SpConnection) GetMetadataReloadSettingsOk() (*ConnectionMetadataUrl, bool)`

GetMetadataReloadSettingsOk returns a tuple with the MetadataReloadSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataReloadSettings

`func (o *SpConnection) SetMetadataReloadSettings(v ConnectionMetadataUrl)`

SetMetadataReloadSettings sets MetadataReloadSettings field to given value.

### HasMetadataReloadSettings

`func (o *SpConnection) HasMetadataReloadSettings() bool`

HasMetadataReloadSettings returns a boolean if a field has been set.

### GetCredentials

`func (o *SpConnection) GetCredentials() ConnectionCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *SpConnection) GetCredentialsOk() (*ConnectionCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *SpConnection) SetCredentials(v ConnectionCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *SpConnection) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetContactInfo

`func (o *SpConnection) GetContactInfo() ContactInfo`

GetContactInfo returns the ContactInfo field if non-nil, zero value otherwise.

### GetContactInfoOk

`func (o *SpConnection) GetContactInfoOk() (*ContactInfo, bool)`

GetContactInfoOk returns a tuple with the ContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactInfo

`func (o *SpConnection) SetContactInfo(v ContactInfo)`

SetContactInfo sets ContactInfo field to given value.

### HasContactInfo

`func (o *SpConnection) HasContactInfo() bool`

HasContactInfo returns a boolean if a field has been set.

### GetLicenseConnectionGroup

`func (o *SpConnection) GetLicenseConnectionGroup() string`

GetLicenseConnectionGroup returns the LicenseConnectionGroup field if non-nil, zero value otherwise.

### GetLicenseConnectionGroupOk

`func (o *SpConnection) GetLicenseConnectionGroupOk() (*string, bool)`

GetLicenseConnectionGroupOk returns a tuple with the LicenseConnectionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseConnectionGroup

`func (o *SpConnection) SetLicenseConnectionGroup(v string)`

SetLicenseConnectionGroup sets LicenseConnectionGroup field to given value.

### HasLicenseConnectionGroup

`func (o *SpConnection) HasLicenseConnectionGroup() bool`

HasLicenseConnectionGroup returns a boolean if a field has been set.

### GetLoggingMode

`func (o *SpConnection) GetLoggingMode() string`

GetLoggingMode returns the LoggingMode field if non-nil, zero value otherwise.

### GetLoggingModeOk

`func (o *SpConnection) GetLoggingModeOk() (*string, bool)`

GetLoggingModeOk returns a tuple with the LoggingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingMode

`func (o *SpConnection) SetLoggingMode(v string)`

SetLoggingMode sets LoggingMode field to given value.

### HasLoggingMode

`func (o *SpConnection) HasLoggingMode() bool`

HasLoggingMode returns a boolean if a field has been set.

### GetAdditionalAllowedEntitiesConfiguration

`func (o *SpConnection) GetAdditionalAllowedEntitiesConfiguration() AdditionalAllowedEntitiesConfiguration`

GetAdditionalAllowedEntitiesConfiguration returns the AdditionalAllowedEntitiesConfiguration field if non-nil, zero value otherwise.

### GetAdditionalAllowedEntitiesConfigurationOk

`func (o *SpConnection) GetAdditionalAllowedEntitiesConfigurationOk() (*AdditionalAllowedEntitiesConfiguration, bool)`

GetAdditionalAllowedEntitiesConfigurationOk returns a tuple with the AdditionalAllowedEntitiesConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAllowedEntitiesConfiguration

`func (o *SpConnection) SetAdditionalAllowedEntitiesConfiguration(v AdditionalAllowedEntitiesConfiguration)`

SetAdditionalAllowedEntitiesConfiguration sets AdditionalAllowedEntitiesConfiguration field to given value.

### HasAdditionalAllowedEntitiesConfiguration

`func (o *SpConnection) HasAdditionalAllowedEntitiesConfiguration() bool`

HasAdditionalAllowedEntitiesConfiguration returns a boolean if a field has been set.

### GetExtendedProperties

`func (o *SpConnection) GetExtendedProperties() map[string]ParameterValues`

GetExtendedProperties returns the ExtendedProperties field if non-nil, zero value otherwise.

### GetExtendedPropertiesOk

`func (o *SpConnection) GetExtendedPropertiesOk() (*map[string]ParameterValues, bool)`

GetExtendedPropertiesOk returns a tuple with the ExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedProperties

`func (o *SpConnection) SetExtendedProperties(v map[string]ParameterValues)`

SetExtendedProperties sets ExtendedProperties field to given value.

### HasExtendedProperties

`func (o *SpConnection) HasExtendedProperties() bool`

HasExtendedProperties returns a boolean if a field has been set.

### GetReplicationStatus

`func (o *SpConnection) GetReplicationStatus() string`

GetReplicationStatus returns the ReplicationStatus field if non-nil, zero value otherwise.

### GetReplicationStatusOk

`func (o *SpConnection) GetReplicationStatusOk() (*string, bool)`

GetReplicationStatusOk returns a tuple with the ReplicationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStatus

`func (o *SpConnection) SetReplicationStatus(v string)`

SetReplicationStatus sets ReplicationStatus field to given value.

### HasReplicationStatus

`func (o *SpConnection) HasReplicationStatus() bool`

HasReplicationStatus returns a boolean if a field has been set.

### GetSpBrowserSso

`func (o *SpConnection) GetSpBrowserSso() SpBrowserSso`

GetSpBrowserSso returns the SpBrowserSso field if non-nil, zero value otherwise.

### GetSpBrowserSsoOk

`func (o *SpConnection) GetSpBrowserSsoOk() (*SpBrowserSso, bool)`

GetSpBrowserSsoOk returns a tuple with the SpBrowserSso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpBrowserSso

`func (o *SpConnection) SetSpBrowserSso(v SpBrowserSso)`

SetSpBrowserSso sets SpBrowserSso field to given value.

### HasSpBrowserSso

`func (o *SpConnection) HasSpBrowserSso() bool`

HasSpBrowserSso returns a boolean if a field has been set.

### GetAttributeQuery

`func (o *SpConnection) GetAttributeQuery() SpAttributeQuery`

GetAttributeQuery returns the AttributeQuery field if non-nil, zero value otherwise.

### GetAttributeQueryOk

`func (o *SpConnection) GetAttributeQueryOk() (*SpAttributeQuery, bool)`

GetAttributeQueryOk returns a tuple with the AttributeQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeQuery

`func (o *SpConnection) SetAttributeQuery(v SpAttributeQuery)`

SetAttributeQuery sets AttributeQuery field to given value.

### HasAttributeQuery

`func (o *SpConnection) HasAttributeQuery() bool`

HasAttributeQuery returns a boolean if a field has been set.

### GetWsTrust

`func (o *SpConnection) GetWsTrust() SpWsTrust`

GetWsTrust returns the WsTrust field if non-nil, zero value otherwise.

### GetWsTrustOk

`func (o *SpConnection) GetWsTrustOk() (*SpWsTrust, bool)`

GetWsTrustOk returns a tuple with the WsTrust field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWsTrust

`func (o *SpConnection) SetWsTrust(v SpWsTrust)`

SetWsTrust sets WsTrust field to given value.

### HasWsTrust

`func (o *SpConnection) HasWsTrust() bool`

HasWsTrust returns a boolean if a field has been set.

### GetApplicationName

`func (o *SpConnection) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *SpConnection) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *SpConnection) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *SpConnection) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetApplicationIconUrl

`func (o *SpConnection) GetApplicationIconUrl() string`

GetApplicationIconUrl returns the ApplicationIconUrl field if non-nil, zero value otherwise.

### GetApplicationIconUrlOk

`func (o *SpConnection) GetApplicationIconUrlOk() (*string, bool)`

GetApplicationIconUrlOk returns a tuple with the ApplicationIconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIconUrl

`func (o *SpConnection) SetApplicationIconUrl(v string)`

SetApplicationIconUrl sets ApplicationIconUrl field to given value.

### HasApplicationIconUrl

`func (o *SpConnection) HasApplicationIconUrl() bool`

HasApplicationIconUrl returns a boolean if a field has been set.

### GetOutboundProvision

`func (o *SpConnection) GetOutboundProvision() OutboundProvision`

GetOutboundProvision returns the OutboundProvision field if non-nil, zero value otherwise.

### GetOutboundProvisionOk

`func (o *SpConnection) GetOutboundProvisionOk() (*OutboundProvision, bool)`

GetOutboundProvisionOk returns a tuple with the OutboundProvision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundProvision

`func (o *SpConnection) SetOutboundProvision(v OutboundProvision)`

SetOutboundProvision sets OutboundProvision field to given value.

### HasOutboundProvision

`func (o *SpConnection) HasOutboundProvision() bool`

HasOutboundProvision returns a boolean if a field has been set.

### GetConnectionTargetType

`func (o *SpConnection) GetConnectionTargetType() string`

GetConnectionTargetType returns the ConnectionTargetType field if non-nil, zero value otherwise.

### GetConnectionTargetTypeOk

`func (o *SpConnection) GetConnectionTargetTypeOk() (*string, bool)`

GetConnectionTargetTypeOk returns a tuple with the ConnectionTargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTargetType

`func (o *SpConnection) SetConnectionTargetType(v string)`

SetConnectionTargetType sets ConnectionTargetType field to given value.

### HasConnectionTargetType

`func (o *SpConnection) HasConnectionTargetType() bool`

HasConnectionTargetType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


