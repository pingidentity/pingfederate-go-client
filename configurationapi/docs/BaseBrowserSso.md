# BaseBrowserSso

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | **string** | The browser-based SSO protocol to use. | 
**EnabledProfiles** | Pointer to **[]string** | The profiles that are enabled for browser-based SSO. SAML 2.0 supports all profiles whereas SAML 1.x IdP connections support both IdP and SP (non-standard) initiated SSO. This is required for SAMLx.x Connections.  | [optional] 
**IncomingBindings** | Pointer to **[]string** | The SAML bindings that are enabled for browser-based SSO. This is required for SAML 2.0 connections when the enabled profiles contain the SP-initiated SSO profile or either SLO profile. For SAML 1.x based connections, it is not used for SP Connections and it is optional for IdP Connections. | [optional] 
**MessageCustomizations** | Pointer to [**[]ProtocolMessageCustomization**](ProtocolMessageCustomization.md) | The message customizations for browser-based SSO. Depending on server settings, connection type, and protocol this may or may not be supported. | [optional] 
**Artifact** | Pointer to [**ArtifactSettings**](ArtifactSettings.md) |  | [optional] 
**SloServiceEndpoints** | Pointer to [**[]SloServiceEndpoint**](SloServiceEndpoint.md) | A list of possible endpoints to send SLO requests and responses. | [optional] 
**UrlWhitelistEntries** | Pointer to [**[]UrlWhitelistEntry**](UrlWhitelistEntry.md) | For WS-Federation connections, a whitelist of additional allowed domains and paths used to validate wreply for SLO, if enabled. | [optional] 
**DefaultTargetUrl** | Pointer to **string** | Default Target URL for SAML1.x connections. For SP connections, this default URL represents the destination on the SP where the user will be directed. For IdP connections, entering a URL in the Default Target URL field overrides the SP Default URL SSO setting. | [optional] 
**AlwaysSignArtifactResponse** | Pointer to **bool** | Specify to always sign the SAML ArtifactResponse. | [optional] 
**SsoApplicationEndpoint** | Pointer to **string** | Application endpoint that can be used to invoke single sign-on (SSO) for the connection. This is a read-only parameter. | [optional] 

## Methods

### NewBaseBrowserSso

`func NewBaseBrowserSso(protocol string, ) *BaseBrowserSso`

NewBaseBrowserSso instantiates a new BaseBrowserSso object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseBrowserSsoWithDefaults

`func NewBaseBrowserSsoWithDefaults() *BaseBrowserSso`

NewBaseBrowserSsoWithDefaults instantiates a new BaseBrowserSso object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *BaseBrowserSso) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *BaseBrowserSso) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *BaseBrowserSso) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetEnabledProfiles

`func (o *BaseBrowserSso) GetEnabledProfiles() []string`

GetEnabledProfiles returns the EnabledProfiles field if non-nil, zero value otherwise.

### GetEnabledProfilesOk

`func (o *BaseBrowserSso) GetEnabledProfilesOk() (*[]string, bool)`

GetEnabledProfilesOk returns a tuple with the EnabledProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledProfiles

`func (o *BaseBrowserSso) SetEnabledProfiles(v []string)`

SetEnabledProfiles sets EnabledProfiles field to given value.

### HasEnabledProfiles

`func (o *BaseBrowserSso) HasEnabledProfiles() bool`

HasEnabledProfiles returns a boolean if a field has been set.

### GetIncomingBindings

`func (o *BaseBrowserSso) GetIncomingBindings() []string`

GetIncomingBindings returns the IncomingBindings field if non-nil, zero value otherwise.

### GetIncomingBindingsOk

`func (o *BaseBrowserSso) GetIncomingBindingsOk() (*[]string, bool)`

GetIncomingBindingsOk returns a tuple with the IncomingBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingBindings

`func (o *BaseBrowserSso) SetIncomingBindings(v []string)`

SetIncomingBindings sets IncomingBindings field to given value.

### HasIncomingBindings

`func (o *BaseBrowserSso) HasIncomingBindings() bool`

HasIncomingBindings returns a boolean if a field has been set.

### GetMessageCustomizations

`func (o *BaseBrowserSso) GetMessageCustomizations() []ProtocolMessageCustomization`

GetMessageCustomizations returns the MessageCustomizations field if non-nil, zero value otherwise.

### GetMessageCustomizationsOk

`func (o *BaseBrowserSso) GetMessageCustomizationsOk() (*[]ProtocolMessageCustomization, bool)`

GetMessageCustomizationsOk returns a tuple with the MessageCustomizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCustomizations

`func (o *BaseBrowserSso) SetMessageCustomizations(v []ProtocolMessageCustomization)`

SetMessageCustomizations sets MessageCustomizations field to given value.

### HasMessageCustomizations

`func (o *BaseBrowserSso) HasMessageCustomizations() bool`

HasMessageCustomizations returns a boolean if a field has been set.

### GetArtifact

`func (o *BaseBrowserSso) GetArtifact() ArtifactSettings`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *BaseBrowserSso) GetArtifactOk() (*ArtifactSettings, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *BaseBrowserSso) SetArtifact(v ArtifactSettings)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *BaseBrowserSso) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetSloServiceEndpoints

`func (o *BaseBrowserSso) GetSloServiceEndpoints() []SloServiceEndpoint`

GetSloServiceEndpoints returns the SloServiceEndpoints field if non-nil, zero value otherwise.

### GetSloServiceEndpointsOk

`func (o *BaseBrowserSso) GetSloServiceEndpointsOk() (*[]SloServiceEndpoint, bool)`

GetSloServiceEndpointsOk returns a tuple with the SloServiceEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloServiceEndpoints

`func (o *BaseBrowserSso) SetSloServiceEndpoints(v []SloServiceEndpoint)`

SetSloServiceEndpoints sets SloServiceEndpoints field to given value.

### HasSloServiceEndpoints

`func (o *BaseBrowserSso) HasSloServiceEndpoints() bool`

HasSloServiceEndpoints returns a boolean if a field has been set.

### GetUrlWhitelistEntries

`func (o *BaseBrowserSso) GetUrlWhitelistEntries() []UrlWhitelistEntry`

GetUrlWhitelistEntries returns the UrlWhitelistEntries field if non-nil, zero value otherwise.

### GetUrlWhitelistEntriesOk

`func (o *BaseBrowserSso) GetUrlWhitelistEntriesOk() (*[]UrlWhitelistEntry, bool)`

GetUrlWhitelistEntriesOk returns a tuple with the UrlWhitelistEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlWhitelistEntries

`func (o *BaseBrowserSso) SetUrlWhitelistEntries(v []UrlWhitelistEntry)`

SetUrlWhitelistEntries sets UrlWhitelistEntries field to given value.

### HasUrlWhitelistEntries

`func (o *BaseBrowserSso) HasUrlWhitelistEntries() bool`

HasUrlWhitelistEntries returns a boolean if a field has been set.

### GetDefaultTargetUrl

`func (o *BaseBrowserSso) GetDefaultTargetUrl() string`

GetDefaultTargetUrl returns the DefaultTargetUrl field if non-nil, zero value otherwise.

### GetDefaultTargetUrlOk

`func (o *BaseBrowserSso) GetDefaultTargetUrlOk() (*string, bool)`

GetDefaultTargetUrlOk returns a tuple with the DefaultTargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTargetUrl

`func (o *BaseBrowserSso) SetDefaultTargetUrl(v string)`

SetDefaultTargetUrl sets DefaultTargetUrl field to given value.

### HasDefaultTargetUrl

`func (o *BaseBrowserSso) HasDefaultTargetUrl() bool`

HasDefaultTargetUrl returns a boolean if a field has been set.

### GetAlwaysSignArtifactResponse

`func (o *BaseBrowserSso) GetAlwaysSignArtifactResponse() bool`

GetAlwaysSignArtifactResponse returns the AlwaysSignArtifactResponse field if non-nil, zero value otherwise.

### GetAlwaysSignArtifactResponseOk

`func (o *BaseBrowserSso) GetAlwaysSignArtifactResponseOk() (*bool, bool)`

GetAlwaysSignArtifactResponseOk returns a tuple with the AlwaysSignArtifactResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysSignArtifactResponse

`func (o *BaseBrowserSso) SetAlwaysSignArtifactResponse(v bool)`

SetAlwaysSignArtifactResponse sets AlwaysSignArtifactResponse field to given value.

### HasAlwaysSignArtifactResponse

`func (o *BaseBrowserSso) HasAlwaysSignArtifactResponse() bool`

HasAlwaysSignArtifactResponse returns a boolean if a field has been set.

### GetSsoApplicationEndpoint

`func (o *BaseBrowserSso) GetSsoApplicationEndpoint() string`

GetSsoApplicationEndpoint returns the SsoApplicationEndpoint field if non-nil, zero value otherwise.

### GetSsoApplicationEndpointOk

`func (o *BaseBrowserSso) GetSsoApplicationEndpointOk() (*string, bool)`

GetSsoApplicationEndpointOk returns a tuple with the SsoApplicationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoApplicationEndpoint

`func (o *BaseBrowserSso) SetSsoApplicationEndpoint(v string)`

SetSsoApplicationEndpoint sets SsoApplicationEndpoint field to given value.

### HasSsoApplicationEndpoint

`func (o *BaseBrowserSso) HasSsoApplicationEndpoint() bool`

HasSsoApplicationEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


