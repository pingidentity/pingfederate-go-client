# IdpBrowserSso

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | **string** | The browser-based SSO protocol to use. | 
**OidcProviderSettings** | Pointer to [**OIDCProviderSettings**](OIDCProviderSettings.md) |  | [optional] 
**EnabledProfiles** | Pointer to **[]string** | The profiles that are enabled for browser-based SSO. SAML 2.0 supports all profiles whereas SAML 1.x IdP connections support both IdP and SP (non-standard) initiated SSO. This is required for SAMLx.x Connections.  | [optional] 
**IncomingBindings** | Pointer to **[]string** | The SAML bindings that are enabled for browser-based SSO. This is required for SAML 2.0 connections when the enabled profiles contain the SP-initiated SSO profile or either SLO profile. For SAML 1.x based connections, it is not used for SP Connections and it is optional for IdP Connections. | [optional] 
**MessageCustomizations** | Pointer to [**[]ProtocolMessageCustomization**](ProtocolMessageCustomization.md) | The message customizations for browser-based SSO. Depending on server settings, connection type, and protocol this may or may not be supported. | [optional] 
**UrlWhitelistEntries** | Pointer to [**[]UrlWhitelistEntry**](UrlWhitelistEntry.md) | For WS-Federation connections, a whitelist of additional allowed domains and paths used to validate wreply for SLO, if enabled. | [optional] 
**Artifact** | Pointer to [**ArtifactSettings**](ArtifactSettings.md) |  | [optional] 
**SloServiceEndpoints** | Pointer to [**[]SloServiceEndpoint**](SloServiceEndpoint.md) | A list of possible endpoints to send SLO requests and responses. | [optional] 
**AlwaysSignArtifactResponse** | Pointer to **bool** | Specify to always sign the SAML ArtifactResponse. | [optional] 
**SsoApplicationEndpoint** | Pointer to **string** | Application endpoint that can be used to invoke single sign-on (SSO) for the connection. This is a read-only parameter. | [optional] 
**SsoServiceEndpoints** | Pointer to [**[]IdpSsoServiceEndpoint**](IdpSsoServiceEndpoint.md) | The IdP SSO endpoints that define where to send your authentication requests. Only required for SP initiated SSO. This is required for SAML x.x and WS-FED Connections. | [optional] 
**DefaultTargetUrl** | Pointer to **string** | The default target URL for this connection. If defined, this overrides the default URL. | [optional] 
**AuthnContextMappings** | Pointer to [**[]AuthnContextMapping**](AuthnContextMapping.md) | A list of authentication context mappings between local and remote values. Applicable for SAML 2.0 and OIDC protocol connections. | [optional] 
**AssertionsSigned** | Pointer to **bool** | Specify whether the incoming SAML assertions are signed rather than the entire SAML response being signed. | [optional] 
**SignAuthnRequests** | Pointer to **bool** | Determines whether SAML authentication requests should be signed. | [optional] 
**DecryptionPolicy** | Pointer to [**DecryptionPolicy**](DecryptionPolicy.md) |  | [optional] 
**IdpIdentityMapping** | **string** | Defines the process in which users authenticated by the IdP are associated with user accounts local to the SP. | 
**AttributeContract** | Pointer to [**IdpBrowserSsoAttributeContract**](IdpBrowserSsoAttributeContract.md) |  | [optional] 
**AdapterMappings** | Pointer to [**[]SpAdapterMapping**](SpAdapterMapping.md) | A list of adapters that map to incoming assertions. | [optional] 
**AuthenticationPolicyContractMappings** | Pointer to [**[]AuthenticationPolicyContractMapping**](AuthenticationPolicyContractMapping.md) | A list of Authentication Policy Contracts that map to incoming assertions. | [optional] 
**SsoOAuthMapping** | Pointer to [**SsoOAuthMapping**](SsoOAuthMapping.md) |  | [optional] 
**OauthAuthenticationPolicyContractRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**JitProvisioning** | Pointer to [**JitProvisioning**](JitProvisioning.md) |  | [optional] 

## Methods

### NewIdpBrowserSso

`func NewIdpBrowserSso(protocol string, idpIdentityMapping string, ) *IdpBrowserSso`

NewIdpBrowserSso instantiates a new IdpBrowserSso object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpBrowserSsoWithDefaults

`func NewIdpBrowserSsoWithDefaults() *IdpBrowserSso`

NewIdpBrowserSsoWithDefaults instantiates a new IdpBrowserSso object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *IdpBrowserSso) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *IdpBrowserSso) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *IdpBrowserSso) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetOidcProviderSettings

`func (o *IdpBrowserSso) GetOidcProviderSettings() OIDCProviderSettings`

GetOidcProviderSettings returns the OidcProviderSettings field if non-nil, zero value otherwise.

### GetOidcProviderSettingsOk

`func (o *IdpBrowserSso) GetOidcProviderSettingsOk() (*OIDCProviderSettings, bool)`

GetOidcProviderSettingsOk returns a tuple with the OidcProviderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcProviderSettings

`func (o *IdpBrowserSso) SetOidcProviderSettings(v OIDCProviderSettings)`

SetOidcProviderSettings sets OidcProviderSettings field to given value.

### HasOidcProviderSettings

`func (o *IdpBrowserSso) HasOidcProviderSettings() bool`

HasOidcProviderSettings returns a boolean if a field has been set.

### GetEnabledProfiles

`func (o *IdpBrowserSso) GetEnabledProfiles() []string`

GetEnabledProfiles returns the EnabledProfiles field if non-nil, zero value otherwise.

### GetEnabledProfilesOk

`func (o *IdpBrowserSso) GetEnabledProfilesOk() (*[]string, bool)`

GetEnabledProfilesOk returns a tuple with the EnabledProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledProfiles

`func (o *IdpBrowserSso) SetEnabledProfiles(v []string)`

SetEnabledProfiles sets EnabledProfiles field to given value.

### HasEnabledProfiles

`func (o *IdpBrowserSso) HasEnabledProfiles() bool`

HasEnabledProfiles returns a boolean if a field has been set.

### GetIncomingBindings

`func (o *IdpBrowserSso) GetIncomingBindings() []string`

GetIncomingBindings returns the IncomingBindings field if non-nil, zero value otherwise.

### GetIncomingBindingsOk

`func (o *IdpBrowserSso) GetIncomingBindingsOk() (*[]string, bool)`

GetIncomingBindingsOk returns a tuple with the IncomingBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingBindings

`func (o *IdpBrowserSso) SetIncomingBindings(v []string)`

SetIncomingBindings sets IncomingBindings field to given value.

### HasIncomingBindings

`func (o *IdpBrowserSso) HasIncomingBindings() bool`

HasIncomingBindings returns a boolean if a field has been set.

### GetMessageCustomizations

`func (o *IdpBrowserSso) GetMessageCustomizations() []ProtocolMessageCustomization`

GetMessageCustomizations returns the MessageCustomizations field if non-nil, zero value otherwise.

### GetMessageCustomizationsOk

`func (o *IdpBrowserSso) GetMessageCustomizationsOk() (*[]ProtocolMessageCustomization, bool)`

GetMessageCustomizationsOk returns a tuple with the MessageCustomizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCustomizations

`func (o *IdpBrowserSso) SetMessageCustomizations(v []ProtocolMessageCustomization)`

SetMessageCustomizations sets MessageCustomizations field to given value.

### HasMessageCustomizations

`func (o *IdpBrowserSso) HasMessageCustomizations() bool`

HasMessageCustomizations returns a boolean if a field has been set.

### GetUrlWhitelistEntries

`func (o *IdpBrowserSso) GetUrlWhitelistEntries() []UrlWhitelistEntry`

GetUrlWhitelistEntries returns the UrlWhitelistEntries field if non-nil, zero value otherwise.

### GetUrlWhitelistEntriesOk

`func (o *IdpBrowserSso) GetUrlWhitelistEntriesOk() (*[]UrlWhitelistEntry, bool)`

GetUrlWhitelistEntriesOk returns a tuple with the UrlWhitelistEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlWhitelistEntries

`func (o *IdpBrowserSso) SetUrlWhitelistEntries(v []UrlWhitelistEntry)`

SetUrlWhitelistEntries sets UrlWhitelistEntries field to given value.

### HasUrlWhitelistEntries

`func (o *IdpBrowserSso) HasUrlWhitelistEntries() bool`

HasUrlWhitelistEntries returns a boolean if a field has been set.

### GetArtifact

`func (o *IdpBrowserSso) GetArtifact() ArtifactSettings`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *IdpBrowserSso) GetArtifactOk() (*ArtifactSettings, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *IdpBrowserSso) SetArtifact(v ArtifactSettings)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *IdpBrowserSso) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetSloServiceEndpoints

`func (o *IdpBrowserSso) GetSloServiceEndpoints() []SloServiceEndpoint`

GetSloServiceEndpoints returns the SloServiceEndpoints field if non-nil, zero value otherwise.

### GetSloServiceEndpointsOk

`func (o *IdpBrowserSso) GetSloServiceEndpointsOk() (*[]SloServiceEndpoint, bool)`

GetSloServiceEndpointsOk returns a tuple with the SloServiceEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloServiceEndpoints

`func (o *IdpBrowserSso) SetSloServiceEndpoints(v []SloServiceEndpoint)`

SetSloServiceEndpoints sets SloServiceEndpoints field to given value.

### HasSloServiceEndpoints

`func (o *IdpBrowserSso) HasSloServiceEndpoints() bool`

HasSloServiceEndpoints returns a boolean if a field has been set.

### GetAlwaysSignArtifactResponse

`func (o *IdpBrowserSso) GetAlwaysSignArtifactResponse() bool`

GetAlwaysSignArtifactResponse returns the AlwaysSignArtifactResponse field if non-nil, zero value otherwise.

### GetAlwaysSignArtifactResponseOk

`func (o *IdpBrowserSso) GetAlwaysSignArtifactResponseOk() (*bool, bool)`

GetAlwaysSignArtifactResponseOk returns a tuple with the AlwaysSignArtifactResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysSignArtifactResponse

`func (o *IdpBrowserSso) SetAlwaysSignArtifactResponse(v bool)`

SetAlwaysSignArtifactResponse sets AlwaysSignArtifactResponse field to given value.

### HasAlwaysSignArtifactResponse

`func (o *IdpBrowserSso) HasAlwaysSignArtifactResponse() bool`

HasAlwaysSignArtifactResponse returns a boolean if a field has been set.

### GetSsoApplicationEndpoint

`func (o *IdpBrowserSso) GetSsoApplicationEndpoint() string`

GetSsoApplicationEndpoint returns the SsoApplicationEndpoint field if non-nil, zero value otherwise.

### GetSsoApplicationEndpointOk

`func (o *IdpBrowserSso) GetSsoApplicationEndpointOk() (*string, bool)`

GetSsoApplicationEndpointOk returns a tuple with the SsoApplicationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoApplicationEndpoint

`func (o *IdpBrowserSso) SetSsoApplicationEndpoint(v string)`

SetSsoApplicationEndpoint sets SsoApplicationEndpoint field to given value.

### HasSsoApplicationEndpoint

`func (o *IdpBrowserSso) HasSsoApplicationEndpoint() bool`

HasSsoApplicationEndpoint returns a boolean if a field has been set.

### GetSsoServiceEndpoints

`func (o *IdpBrowserSso) GetSsoServiceEndpoints() []IdpSsoServiceEndpoint`

GetSsoServiceEndpoints returns the SsoServiceEndpoints field if non-nil, zero value otherwise.

### GetSsoServiceEndpointsOk

`func (o *IdpBrowserSso) GetSsoServiceEndpointsOk() (*[]IdpSsoServiceEndpoint, bool)`

GetSsoServiceEndpointsOk returns a tuple with the SsoServiceEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoServiceEndpoints

`func (o *IdpBrowserSso) SetSsoServiceEndpoints(v []IdpSsoServiceEndpoint)`

SetSsoServiceEndpoints sets SsoServiceEndpoints field to given value.

### HasSsoServiceEndpoints

`func (o *IdpBrowserSso) HasSsoServiceEndpoints() bool`

HasSsoServiceEndpoints returns a boolean if a field has been set.

### GetDefaultTargetUrl

`func (o *IdpBrowserSso) GetDefaultTargetUrl() string`

GetDefaultTargetUrl returns the DefaultTargetUrl field if non-nil, zero value otherwise.

### GetDefaultTargetUrlOk

`func (o *IdpBrowserSso) GetDefaultTargetUrlOk() (*string, bool)`

GetDefaultTargetUrlOk returns a tuple with the DefaultTargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTargetUrl

`func (o *IdpBrowserSso) SetDefaultTargetUrl(v string)`

SetDefaultTargetUrl sets DefaultTargetUrl field to given value.

### HasDefaultTargetUrl

`func (o *IdpBrowserSso) HasDefaultTargetUrl() bool`

HasDefaultTargetUrl returns a boolean if a field has been set.

### GetAuthnContextMappings

`func (o *IdpBrowserSso) GetAuthnContextMappings() []AuthnContextMapping`

GetAuthnContextMappings returns the AuthnContextMappings field if non-nil, zero value otherwise.

### GetAuthnContextMappingsOk

`func (o *IdpBrowserSso) GetAuthnContextMappingsOk() (*[]AuthnContextMapping, bool)`

GetAuthnContextMappingsOk returns a tuple with the AuthnContextMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnContextMappings

`func (o *IdpBrowserSso) SetAuthnContextMappings(v []AuthnContextMapping)`

SetAuthnContextMappings sets AuthnContextMappings field to given value.

### HasAuthnContextMappings

`func (o *IdpBrowserSso) HasAuthnContextMappings() bool`

HasAuthnContextMappings returns a boolean if a field has been set.

### GetAssertionsSigned

`func (o *IdpBrowserSso) GetAssertionsSigned() bool`

GetAssertionsSigned returns the AssertionsSigned field if non-nil, zero value otherwise.

### GetAssertionsSignedOk

`func (o *IdpBrowserSso) GetAssertionsSignedOk() (*bool, bool)`

GetAssertionsSignedOk returns a tuple with the AssertionsSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionsSigned

`func (o *IdpBrowserSso) SetAssertionsSigned(v bool)`

SetAssertionsSigned sets AssertionsSigned field to given value.

### HasAssertionsSigned

`func (o *IdpBrowserSso) HasAssertionsSigned() bool`

HasAssertionsSigned returns a boolean if a field has been set.

### GetSignAuthnRequests

`func (o *IdpBrowserSso) GetSignAuthnRequests() bool`

GetSignAuthnRequests returns the SignAuthnRequests field if non-nil, zero value otherwise.

### GetSignAuthnRequestsOk

`func (o *IdpBrowserSso) GetSignAuthnRequestsOk() (*bool, bool)`

GetSignAuthnRequestsOk returns a tuple with the SignAuthnRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignAuthnRequests

`func (o *IdpBrowserSso) SetSignAuthnRequests(v bool)`

SetSignAuthnRequests sets SignAuthnRequests field to given value.

### HasSignAuthnRequests

`func (o *IdpBrowserSso) HasSignAuthnRequests() bool`

HasSignAuthnRequests returns a boolean if a field has been set.

### GetDecryptionPolicy

`func (o *IdpBrowserSso) GetDecryptionPolicy() DecryptionPolicy`

GetDecryptionPolicy returns the DecryptionPolicy field if non-nil, zero value otherwise.

### GetDecryptionPolicyOk

`func (o *IdpBrowserSso) GetDecryptionPolicyOk() (*DecryptionPolicy, bool)`

GetDecryptionPolicyOk returns a tuple with the DecryptionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecryptionPolicy

`func (o *IdpBrowserSso) SetDecryptionPolicy(v DecryptionPolicy)`

SetDecryptionPolicy sets DecryptionPolicy field to given value.

### HasDecryptionPolicy

`func (o *IdpBrowserSso) HasDecryptionPolicy() bool`

HasDecryptionPolicy returns a boolean if a field has been set.

### GetIdpIdentityMapping

`func (o *IdpBrowserSso) GetIdpIdentityMapping() string`

GetIdpIdentityMapping returns the IdpIdentityMapping field if non-nil, zero value otherwise.

### GetIdpIdentityMappingOk

`func (o *IdpBrowserSso) GetIdpIdentityMappingOk() (*string, bool)`

GetIdpIdentityMappingOk returns a tuple with the IdpIdentityMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpIdentityMapping

`func (o *IdpBrowserSso) SetIdpIdentityMapping(v string)`

SetIdpIdentityMapping sets IdpIdentityMapping field to given value.


### GetAttributeContract

`func (o *IdpBrowserSso) GetAttributeContract() IdpBrowserSsoAttributeContract`

GetAttributeContract returns the AttributeContract field if non-nil, zero value otherwise.

### GetAttributeContractOk

`func (o *IdpBrowserSso) GetAttributeContractOk() (*IdpBrowserSsoAttributeContract, bool)`

GetAttributeContractOk returns a tuple with the AttributeContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContract

`func (o *IdpBrowserSso) SetAttributeContract(v IdpBrowserSsoAttributeContract)`

SetAttributeContract sets AttributeContract field to given value.

### HasAttributeContract

`func (o *IdpBrowserSso) HasAttributeContract() bool`

HasAttributeContract returns a boolean if a field has been set.

### GetAdapterMappings

`func (o *IdpBrowserSso) GetAdapterMappings() []SpAdapterMapping`

GetAdapterMappings returns the AdapterMappings field if non-nil, zero value otherwise.

### GetAdapterMappingsOk

`func (o *IdpBrowserSso) GetAdapterMappingsOk() (*[]SpAdapterMapping, bool)`

GetAdapterMappingsOk returns a tuple with the AdapterMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterMappings

`func (o *IdpBrowserSso) SetAdapterMappings(v []SpAdapterMapping)`

SetAdapterMappings sets AdapterMappings field to given value.

### HasAdapterMappings

`func (o *IdpBrowserSso) HasAdapterMappings() bool`

HasAdapterMappings returns a boolean if a field has been set.

### GetAuthenticationPolicyContractMappings

`func (o *IdpBrowserSso) GetAuthenticationPolicyContractMappings() []AuthenticationPolicyContractMapping`

GetAuthenticationPolicyContractMappings returns the AuthenticationPolicyContractMappings field if non-nil, zero value otherwise.

### GetAuthenticationPolicyContractMappingsOk

`func (o *IdpBrowserSso) GetAuthenticationPolicyContractMappingsOk() (*[]AuthenticationPolicyContractMapping, bool)`

GetAuthenticationPolicyContractMappingsOk returns a tuple with the AuthenticationPolicyContractMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationPolicyContractMappings

`func (o *IdpBrowserSso) SetAuthenticationPolicyContractMappings(v []AuthenticationPolicyContractMapping)`

SetAuthenticationPolicyContractMappings sets AuthenticationPolicyContractMappings field to given value.

### HasAuthenticationPolicyContractMappings

`func (o *IdpBrowserSso) HasAuthenticationPolicyContractMappings() bool`

HasAuthenticationPolicyContractMappings returns a boolean if a field has been set.

### GetSsoOAuthMapping

`func (o *IdpBrowserSso) GetSsoOAuthMapping() SsoOAuthMapping`

GetSsoOAuthMapping returns the SsoOAuthMapping field if non-nil, zero value otherwise.

### GetSsoOAuthMappingOk

`func (o *IdpBrowserSso) GetSsoOAuthMappingOk() (*SsoOAuthMapping, bool)`

GetSsoOAuthMappingOk returns a tuple with the SsoOAuthMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoOAuthMapping

`func (o *IdpBrowserSso) SetSsoOAuthMapping(v SsoOAuthMapping)`

SetSsoOAuthMapping sets SsoOAuthMapping field to given value.

### HasSsoOAuthMapping

`func (o *IdpBrowserSso) HasSsoOAuthMapping() bool`

HasSsoOAuthMapping returns a boolean if a field has been set.

### GetOauthAuthenticationPolicyContractRef

`func (o *IdpBrowserSso) GetOauthAuthenticationPolicyContractRef() ResourceLink`

GetOauthAuthenticationPolicyContractRef returns the OauthAuthenticationPolicyContractRef field if non-nil, zero value otherwise.

### GetOauthAuthenticationPolicyContractRefOk

`func (o *IdpBrowserSso) GetOauthAuthenticationPolicyContractRefOk() (*ResourceLink, bool)`

GetOauthAuthenticationPolicyContractRefOk returns a tuple with the OauthAuthenticationPolicyContractRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthAuthenticationPolicyContractRef

`func (o *IdpBrowserSso) SetOauthAuthenticationPolicyContractRef(v ResourceLink)`

SetOauthAuthenticationPolicyContractRef sets OauthAuthenticationPolicyContractRef field to given value.

### HasOauthAuthenticationPolicyContractRef

`func (o *IdpBrowserSso) HasOauthAuthenticationPolicyContractRef() bool`

HasOauthAuthenticationPolicyContractRef returns a boolean if a field has been set.

### GetJitProvisioning

`func (o *IdpBrowserSso) GetJitProvisioning() JitProvisioning`

GetJitProvisioning returns the JitProvisioning field if non-nil, zero value otherwise.

### GetJitProvisioningOk

`func (o *IdpBrowserSso) GetJitProvisioningOk() (*JitProvisioning, bool)`

GetJitProvisioningOk returns a tuple with the JitProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitProvisioning

`func (o *IdpBrowserSso) SetJitProvisioning(v JitProvisioning)`

SetJitProvisioning sets JitProvisioning field to given value.

### HasJitProvisioning

`func (o *IdpBrowserSso) HasJitProvisioning() bool`

HasJitProvisioning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


