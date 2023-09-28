# Connection

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

## Methods

### NewConnection

`func NewConnection(entityId string, name string, ) *Connection`

NewConnection instantiates a new Connection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionWithDefaults

`func NewConnectionWithDefaults() *Connection`

NewConnectionWithDefaults instantiates a new Connection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Connection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Connection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Connection) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Connection) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *Connection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Connection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Connection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Connection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEntityId

`func (o *Connection) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *Connection) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *Connection) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetName

`func (o *Connection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Connection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Connection) SetName(v string)`

SetName sets Name field to given value.


### GetModificationDate

`func (o *Connection) GetModificationDate() time.Time`

GetModificationDate returns the ModificationDate field if non-nil, zero value otherwise.

### GetModificationDateOk

`func (o *Connection) GetModificationDateOk() (*time.Time, bool)`

GetModificationDateOk returns a tuple with the ModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationDate

`func (o *Connection) SetModificationDate(v time.Time)`

SetModificationDate sets ModificationDate field to given value.

### HasModificationDate

`func (o *Connection) HasModificationDate() bool`

HasModificationDate returns a boolean if a field has been set.

### GetCreationDate

`func (o *Connection) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *Connection) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *Connection) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *Connection) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetActive

`func (o *Connection) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Connection) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Connection) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Connection) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetBaseUrl

`func (o *Connection) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *Connection) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *Connection) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *Connection) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetDefaultVirtualEntityId

`func (o *Connection) GetDefaultVirtualEntityId() string`

GetDefaultVirtualEntityId returns the DefaultVirtualEntityId field if non-nil, zero value otherwise.

### GetDefaultVirtualEntityIdOk

`func (o *Connection) GetDefaultVirtualEntityIdOk() (*string, bool)`

GetDefaultVirtualEntityIdOk returns a tuple with the DefaultVirtualEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVirtualEntityId

`func (o *Connection) SetDefaultVirtualEntityId(v string)`

SetDefaultVirtualEntityId sets DefaultVirtualEntityId field to given value.

### HasDefaultVirtualEntityId

`func (o *Connection) HasDefaultVirtualEntityId() bool`

HasDefaultVirtualEntityId returns a boolean if a field has been set.

### GetVirtualEntityIds

`func (o *Connection) GetVirtualEntityIds() []string`

GetVirtualEntityIds returns the VirtualEntityIds field if non-nil, zero value otherwise.

### GetVirtualEntityIdsOk

`func (o *Connection) GetVirtualEntityIdsOk() (*[]string, bool)`

GetVirtualEntityIdsOk returns a tuple with the VirtualEntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualEntityIds

`func (o *Connection) SetVirtualEntityIds(v []string)`

SetVirtualEntityIds sets VirtualEntityIds field to given value.

### HasVirtualEntityIds

`func (o *Connection) HasVirtualEntityIds() bool`

HasVirtualEntityIds returns a boolean if a field has been set.

### GetMetadataReloadSettings

`func (o *Connection) GetMetadataReloadSettings() ConnectionMetadataUrl`

GetMetadataReloadSettings returns the MetadataReloadSettings field if non-nil, zero value otherwise.

### GetMetadataReloadSettingsOk

`func (o *Connection) GetMetadataReloadSettingsOk() (*ConnectionMetadataUrl, bool)`

GetMetadataReloadSettingsOk returns a tuple with the MetadataReloadSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataReloadSettings

`func (o *Connection) SetMetadataReloadSettings(v ConnectionMetadataUrl)`

SetMetadataReloadSettings sets MetadataReloadSettings field to given value.

### HasMetadataReloadSettings

`func (o *Connection) HasMetadataReloadSettings() bool`

HasMetadataReloadSettings returns a boolean if a field has been set.

### GetCredentials

`func (o *Connection) GetCredentials() ConnectionCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *Connection) GetCredentialsOk() (*ConnectionCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *Connection) SetCredentials(v ConnectionCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *Connection) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetContactInfo

`func (o *Connection) GetContactInfo() ContactInfo`

GetContactInfo returns the ContactInfo field if non-nil, zero value otherwise.

### GetContactInfoOk

`func (o *Connection) GetContactInfoOk() (*ContactInfo, bool)`

GetContactInfoOk returns a tuple with the ContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactInfo

`func (o *Connection) SetContactInfo(v ContactInfo)`

SetContactInfo sets ContactInfo field to given value.

### HasContactInfo

`func (o *Connection) HasContactInfo() bool`

HasContactInfo returns a boolean if a field has been set.

### GetLicenseConnectionGroup

`func (o *Connection) GetLicenseConnectionGroup() string`

GetLicenseConnectionGroup returns the LicenseConnectionGroup field if non-nil, zero value otherwise.

### GetLicenseConnectionGroupOk

`func (o *Connection) GetLicenseConnectionGroupOk() (*string, bool)`

GetLicenseConnectionGroupOk returns a tuple with the LicenseConnectionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseConnectionGroup

`func (o *Connection) SetLicenseConnectionGroup(v string)`

SetLicenseConnectionGroup sets LicenseConnectionGroup field to given value.

### HasLicenseConnectionGroup

`func (o *Connection) HasLicenseConnectionGroup() bool`

HasLicenseConnectionGroup returns a boolean if a field has been set.

### GetLoggingMode

`func (o *Connection) GetLoggingMode() string`

GetLoggingMode returns the LoggingMode field if non-nil, zero value otherwise.

### GetLoggingModeOk

`func (o *Connection) GetLoggingModeOk() (*string, bool)`

GetLoggingModeOk returns a tuple with the LoggingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingMode

`func (o *Connection) SetLoggingMode(v string)`

SetLoggingMode sets LoggingMode field to given value.

### HasLoggingMode

`func (o *Connection) HasLoggingMode() bool`

HasLoggingMode returns a boolean if a field has been set.

### GetAdditionalAllowedEntitiesConfiguration

`func (o *Connection) GetAdditionalAllowedEntitiesConfiguration() AdditionalAllowedEntitiesConfiguration`

GetAdditionalAllowedEntitiesConfiguration returns the AdditionalAllowedEntitiesConfiguration field if non-nil, zero value otherwise.

### GetAdditionalAllowedEntitiesConfigurationOk

`func (o *Connection) GetAdditionalAllowedEntitiesConfigurationOk() (*AdditionalAllowedEntitiesConfiguration, bool)`

GetAdditionalAllowedEntitiesConfigurationOk returns a tuple with the AdditionalAllowedEntitiesConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAllowedEntitiesConfiguration

`func (o *Connection) SetAdditionalAllowedEntitiesConfiguration(v AdditionalAllowedEntitiesConfiguration)`

SetAdditionalAllowedEntitiesConfiguration sets AdditionalAllowedEntitiesConfiguration field to given value.

### HasAdditionalAllowedEntitiesConfiguration

`func (o *Connection) HasAdditionalAllowedEntitiesConfiguration() bool`

HasAdditionalAllowedEntitiesConfiguration returns a boolean if a field has been set.

### GetExtendedProperties

`func (o *Connection) GetExtendedProperties() map[string]ParameterValues`

GetExtendedProperties returns the ExtendedProperties field if non-nil, zero value otherwise.

### GetExtendedPropertiesOk

`func (o *Connection) GetExtendedPropertiesOk() (*map[string]ParameterValues, bool)`

GetExtendedPropertiesOk returns a tuple with the ExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedProperties

`func (o *Connection) SetExtendedProperties(v map[string]ParameterValues)`

SetExtendedProperties sets ExtendedProperties field to given value.

### HasExtendedProperties

`func (o *Connection) HasExtendedProperties() bool`

HasExtendedProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


