# SpBrowserSso

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | **string** | The browser-based SSO protocol to use. | 
**WsFedTokenType** | Pointer to **string** | The WS-Federation Token Type to use. | [optional] 
**WsTrustVersion** | Pointer to **string** | The WS-Trust version for a WS-Federation connection. The default version is WSTRUST12. | [optional] 
**EnabledProfiles** | Pointer to **[]string** | The profiles that are enabled for browser-based SSO. SAML 2.0 supports all profiles whereas SAML 1.x IdP connections support both IdP and SP (non-standard) initiated SSO. This is required for SAMLx.x Connections.  | [optional] 
**IncomingBindings** | Pointer to **[]string** | The SAML bindings that are enabled for browser-based SSO. This is required for SAML 2.0 connections when the enabled profiles contain the SP-initiated SSO profile or either SLO profile. For SAML 1.x based connections, it is not used for SP Connections and it is optional for IdP Connections. | [optional] 
**MessageCustomizations** | Pointer to [**[]ProtocolMessageCustomization**](ProtocolMessageCustomization.md) | The message customizations for browser-based SSO. Depending on server settings, connection type, and protocol this may or may not be supported. | [optional] 
**UrlWhitelistEntries** | Pointer to [**[]UrlWhitelistEntry**](UrlWhitelistEntry.md) | For WS-Federation connections, a whitelist of additional allowed domains and paths used to validate wreply for SLO, if enabled. | [optional] 
**Artifact** | Pointer to [**ArtifactSettings**](ArtifactSettings.md) |  | [optional] 
**SloServiceEndpoints** | Pointer to [**[]SloServiceEndpoint**](SloServiceEndpoint.md) | A list of possible endpoints to send SLO requests and responses. | [optional] 
**DefaultTargetUrl** | Pointer to **string** | Default Target URL for SAML1.x connections. For SP connections, this default URL represents the destination on the SP where the user will be directed. For IdP connections, entering a URL in the Default Target URL field overrides the SP Default URL SSO setting. | [optional] 
**AlwaysSignArtifactResponse** | Pointer to **bool** | Specify to always sign the SAML ArtifactResponse. | [optional] 
**SsoServiceEndpoints** | [**[]SpSsoServiceEndpoint**](SpSsoServiceEndpoint.md) | A list of possible endpoints to send assertions to. | 
**SpSamlIdentityMapping** | Pointer to **string** | Process in which users authenticated by the IdP are associated with user accounts local to the SP. | [optional] 
**SpWsFedIdentityMapping** | Pointer to **string** | Process in which users authenticated by the IdP are associated with user accounts local to the SP for WS-Federation connection types. | [optional] 
**SignResponseAsRequired** | Pointer to **bool** | Sign SAML Response as required by the associated binding and encryption policy. Applicable to SAML2.0 only and is defaulted to true. It can be set to false only on SAML2.0 connections when signAssertions is set to true. | [optional] 
**SignAssertions** | Pointer to **bool** | Always sign the SAML Assertion. | [optional] 
**RequireSignedAuthnRequests** | Pointer to **bool** | Require AuthN requests to be signed when received via the POST or Redirect bindings. | [optional] 
**EncryptionPolicy** | [**EncryptionPolicy**](EncryptionPolicy.md) |  | 
**AttributeContract** | [**SpBrowserSsoAttributeContract**](SpBrowserSsoAttributeContract.md) |  | 
**AdapterMappings** | [**[]IdpAdapterAssertionMapping**](IdpAdapterAssertionMapping.md) | A list of adapters that map to outgoing assertions. | 
**AuthenticationPolicyContractAssertionMappings** | Pointer to [**[]AuthenticationPolicyContractAssertionMapping**](AuthenticationPolicyContractAssertionMapping.md) | A list of authentication policy contracts that map to outgoing assertions. | [optional] 
**AssertionLifetime** | [**AssertionLifetime**](AssertionLifetime.md) |  | 

## Methods

### NewSpBrowserSso

`func NewSpBrowserSso(protocol string, ssoServiceEndpoints []SpSsoServiceEndpoint, encryptionPolicy EncryptionPolicy, attributeContract SpBrowserSsoAttributeContract, adapterMappings []IdpAdapterAssertionMapping, assertionLifetime AssertionLifetime, ) *SpBrowserSso`

NewSpBrowserSso instantiates a new SpBrowserSso object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpBrowserSsoWithDefaults

`func NewSpBrowserSsoWithDefaults() *SpBrowserSso`

NewSpBrowserSsoWithDefaults instantiates a new SpBrowserSso object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *SpBrowserSso) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SpBrowserSso) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SpBrowserSso) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetWsFedTokenType

`func (o *SpBrowserSso) GetWsFedTokenType() string`

GetWsFedTokenType returns the WsFedTokenType field if non-nil, zero value otherwise.

### GetWsFedTokenTypeOk

`func (o *SpBrowserSso) GetWsFedTokenTypeOk() (*string, bool)`

GetWsFedTokenTypeOk returns a tuple with the WsFedTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWsFedTokenType

`func (o *SpBrowserSso) SetWsFedTokenType(v string)`

SetWsFedTokenType sets WsFedTokenType field to given value.

### HasWsFedTokenType

`func (o *SpBrowserSso) HasWsFedTokenType() bool`

HasWsFedTokenType returns a boolean if a field has been set.

### GetWsTrustVersion

`func (o *SpBrowserSso) GetWsTrustVersion() string`

GetWsTrustVersion returns the WsTrustVersion field if non-nil, zero value otherwise.

### GetWsTrustVersionOk

`func (o *SpBrowserSso) GetWsTrustVersionOk() (*string, bool)`

GetWsTrustVersionOk returns a tuple with the WsTrustVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWsTrustVersion

`func (o *SpBrowserSso) SetWsTrustVersion(v string)`

SetWsTrustVersion sets WsTrustVersion field to given value.

### HasWsTrustVersion

`func (o *SpBrowserSso) HasWsTrustVersion() bool`

HasWsTrustVersion returns a boolean if a field has been set.

### GetEnabledProfiles

`func (o *SpBrowserSso) GetEnabledProfiles() []string`

GetEnabledProfiles returns the EnabledProfiles field if non-nil, zero value otherwise.

### GetEnabledProfilesOk

`func (o *SpBrowserSso) GetEnabledProfilesOk() (*[]string, bool)`

GetEnabledProfilesOk returns a tuple with the EnabledProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledProfiles

`func (o *SpBrowserSso) SetEnabledProfiles(v []string)`

SetEnabledProfiles sets EnabledProfiles field to given value.

### HasEnabledProfiles

`func (o *SpBrowserSso) HasEnabledProfiles() bool`

HasEnabledProfiles returns a boolean if a field has been set.

### GetIncomingBindings

`func (o *SpBrowserSso) GetIncomingBindings() []string`

GetIncomingBindings returns the IncomingBindings field if non-nil, zero value otherwise.

### GetIncomingBindingsOk

`func (o *SpBrowserSso) GetIncomingBindingsOk() (*[]string, bool)`

GetIncomingBindingsOk returns a tuple with the IncomingBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingBindings

`func (o *SpBrowserSso) SetIncomingBindings(v []string)`

SetIncomingBindings sets IncomingBindings field to given value.

### HasIncomingBindings

`func (o *SpBrowserSso) HasIncomingBindings() bool`

HasIncomingBindings returns a boolean if a field has been set.

### GetMessageCustomizations

`func (o *SpBrowserSso) GetMessageCustomizations() []ProtocolMessageCustomization`

GetMessageCustomizations returns the MessageCustomizations field if non-nil, zero value otherwise.

### GetMessageCustomizationsOk

`func (o *SpBrowserSso) GetMessageCustomizationsOk() (*[]ProtocolMessageCustomization, bool)`

GetMessageCustomizationsOk returns a tuple with the MessageCustomizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCustomizations

`func (o *SpBrowserSso) SetMessageCustomizations(v []ProtocolMessageCustomization)`

SetMessageCustomizations sets MessageCustomizations field to given value.

### HasMessageCustomizations

`func (o *SpBrowserSso) HasMessageCustomizations() bool`

HasMessageCustomizations returns a boolean if a field has been set.

### GetUrlWhitelistEntries

`func (o *SpBrowserSso) GetUrlWhitelistEntries() []UrlWhitelistEntry`

GetUrlWhitelistEntries returns the UrlWhitelistEntries field if non-nil, zero value otherwise.

### GetUrlWhitelistEntriesOk

`func (o *SpBrowserSso) GetUrlWhitelistEntriesOk() (*[]UrlWhitelistEntry, bool)`

GetUrlWhitelistEntriesOk returns a tuple with the UrlWhitelistEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlWhitelistEntries

`func (o *SpBrowserSso) SetUrlWhitelistEntries(v []UrlWhitelistEntry)`

SetUrlWhitelistEntries sets UrlWhitelistEntries field to given value.

### HasUrlWhitelistEntries

`func (o *SpBrowserSso) HasUrlWhitelistEntries() bool`

HasUrlWhitelistEntries returns a boolean if a field has been set.

### GetArtifact

`func (o *SpBrowserSso) GetArtifact() ArtifactSettings`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *SpBrowserSso) GetArtifactOk() (*ArtifactSettings, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *SpBrowserSso) SetArtifact(v ArtifactSettings)`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *SpBrowserSso) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### GetSloServiceEndpoints

`func (o *SpBrowserSso) GetSloServiceEndpoints() []SloServiceEndpoint`

GetSloServiceEndpoints returns the SloServiceEndpoints field if non-nil, zero value otherwise.

### GetSloServiceEndpointsOk

`func (o *SpBrowserSso) GetSloServiceEndpointsOk() (*[]SloServiceEndpoint, bool)`

GetSloServiceEndpointsOk returns a tuple with the SloServiceEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloServiceEndpoints

`func (o *SpBrowserSso) SetSloServiceEndpoints(v []SloServiceEndpoint)`

SetSloServiceEndpoints sets SloServiceEndpoints field to given value.

### HasSloServiceEndpoints

`func (o *SpBrowserSso) HasSloServiceEndpoints() bool`

HasSloServiceEndpoints returns a boolean if a field has been set.

### GetDefaultTargetUrl

`func (o *SpBrowserSso) GetDefaultTargetUrl() string`

GetDefaultTargetUrl returns the DefaultTargetUrl field if non-nil, zero value otherwise.

### GetDefaultTargetUrlOk

`func (o *SpBrowserSso) GetDefaultTargetUrlOk() (*string, bool)`

GetDefaultTargetUrlOk returns a tuple with the DefaultTargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTargetUrl

`func (o *SpBrowserSso) SetDefaultTargetUrl(v string)`

SetDefaultTargetUrl sets DefaultTargetUrl field to given value.

### HasDefaultTargetUrl

`func (o *SpBrowserSso) HasDefaultTargetUrl() bool`

HasDefaultTargetUrl returns a boolean if a field has been set.

### GetAlwaysSignArtifactResponse

`func (o *SpBrowserSso) GetAlwaysSignArtifactResponse() bool`

GetAlwaysSignArtifactResponse returns the AlwaysSignArtifactResponse field if non-nil, zero value otherwise.

### GetAlwaysSignArtifactResponseOk

`func (o *SpBrowserSso) GetAlwaysSignArtifactResponseOk() (*bool, bool)`

GetAlwaysSignArtifactResponseOk returns a tuple with the AlwaysSignArtifactResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysSignArtifactResponse

`func (o *SpBrowserSso) SetAlwaysSignArtifactResponse(v bool)`

SetAlwaysSignArtifactResponse sets AlwaysSignArtifactResponse field to given value.

### HasAlwaysSignArtifactResponse

`func (o *SpBrowserSso) HasAlwaysSignArtifactResponse() bool`

HasAlwaysSignArtifactResponse returns a boolean if a field has been set.

### GetSsoServiceEndpoints

`func (o *SpBrowserSso) GetSsoServiceEndpoints() []SpSsoServiceEndpoint`

GetSsoServiceEndpoints returns the SsoServiceEndpoints field if non-nil, zero value otherwise.

### GetSsoServiceEndpointsOk

`func (o *SpBrowserSso) GetSsoServiceEndpointsOk() (*[]SpSsoServiceEndpoint, bool)`

GetSsoServiceEndpointsOk returns a tuple with the SsoServiceEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoServiceEndpoints

`func (o *SpBrowserSso) SetSsoServiceEndpoints(v []SpSsoServiceEndpoint)`

SetSsoServiceEndpoints sets SsoServiceEndpoints field to given value.


### GetSpSamlIdentityMapping

`func (o *SpBrowserSso) GetSpSamlIdentityMapping() string`

GetSpSamlIdentityMapping returns the SpSamlIdentityMapping field if non-nil, zero value otherwise.

### GetSpSamlIdentityMappingOk

`func (o *SpBrowserSso) GetSpSamlIdentityMappingOk() (*string, bool)`

GetSpSamlIdentityMappingOk returns a tuple with the SpSamlIdentityMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpSamlIdentityMapping

`func (o *SpBrowserSso) SetSpSamlIdentityMapping(v string)`

SetSpSamlIdentityMapping sets SpSamlIdentityMapping field to given value.

### HasSpSamlIdentityMapping

`func (o *SpBrowserSso) HasSpSamlIdentityMapping() bool`

HasSpSamlIdentityMapping returns a boolean if a field has been set.

### GetSpWsFedIdentityMapping

`func (o *SpBrowserSso) GetSpWsFedIdentityMapping() string`

GetSpWsFedIdentityMapping returns the SpWsFedIdentityMapping field if non-nil, zero value otherwise.

### GetSpWsFedIdentityMappingOk

`func (o *SpBrowserSso) GetSpWsFedIdentityMappingOk() (*string, bool)`

GetSpWsFedIdentityMappingOk returns a tuple with the SpWsFedIdentityMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpWsFedIdentityMapping

`func (o *SpBrowserSso) SetSpWsFedIdentityMapping(v string)`

SetSpWsFedIdentityMapping sets SpWsFedIdentityMapping field to given value.

### HasSpWsFedIdentityMapping

`func (o *SpBrowserSso) HasSpWsFedIdentityMapping() bool`

HasSpWsFedIdentityMapping returns a boolean if a field has been set.

### GetSignResponseAsRequired

`func (o *SpBrowserSso) GetSignResponseAsRequired() bool`

GetSignResponseAsRequired returns the SignResponseAsRequired field if non-nil, zero value otherwise.

### GetSignResponseAsRequiredOk

`func (o *SpBrowserSso) GetSignResponseAsRequiredOk() (*bool, bool)`

GetSignResponseAsRequiredOk returns a tuple with the SignResponseAsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignResponseAsRequired

`func (o *SpBrowserSso) SetSignResponseAsRequired(v bool)`

SetSignResponseAsRequired sets SignResponseAsRequired field to given value.

### HasSignResponseAsRequired

`func (o *SpBrowserSso) HasSignResponseAsRequired() bool`

HasSignResponseAsRequired returns a boolean if a field has been set.

### GetSignAssertions

`func (o *SpBrowserSso) GetSignAssertions() bool`

GetSignAssertions returns the SignAssertions field if non-nil, zero value otherwise.

### GetSignAssertionsOk

`func (o *SpBrowserSso) GetSignAssertionsOk() (*bool, bool)`

GetSignAssertionsOk returns a tuple with the SignAssertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignAssertions

`func (o *SpBrowserSso) SetSignAssertions(v bool)`

SetSignAssertions sets SignAssertions field to given value.

### HasSignAssertions

`func (o *SpBrowserSso) HasSignAssertions() bool`

HasSignAssertions returns a boolean if a field has been set.

### GetRequireSignedAuthnRequests

`func (o *SpBrowserSso) GetRequireSignedAuthnRequests() bool`

GetRequireSignedAuthnRequests returns the RequireSignedAuthnRequests field if non-nil, zero value otherwise.

### GetRequireSignedAuthnRequestsOk

`func (o *SpBrowserSso) GetRequireSignedAuthnRequestsOk() (*bool, bool)`

GetRequireSignedAuthnRequestsOk returns a tuple with the RequireSignedAuthnRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSignedAuthnRequests

`func (o *SpBrowserSso) SetRequireSignedAuthnRequests(v bool)`

SetRequireSignedAuthnRequests sets RequireSignedAuthnRequests field to given value.

### HasRequireSignedAuthnRequests

`func (o *SpBrowserSso) HasRequireSignedAuthnRequests() bool`

HasRequireSignedAuthnRequests returns a boolean if a field has been set.

### GetEncryptionPolicy

`func (o *SpBrowserSso) GetEncryptionPolicy() EncryptionPolicy`

GetEncryptionPolicy returns the EncryptionPolicy field if non-nil, zero value otherwise.

### GetEncryptionPolicyOk

`func (o *SpBrowserSso) GetEncryptionPolicyOk() (*EncryptionPolicy, bool)`

GetEncryptionPolicyOk returns a tuple with the EncryptionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionPolicy

`func (o *SpBrowserSso) SetEncryptionPolicy(v EncryptionPolicy)`

SetEncryptionPolicy sets EncryptionPolicy field to given value.


### GetAttributeContract

`func (o *SpBrowserSso) GetAttributeContract() SpBrowserSsoAttributeContract`

GetAttributeContract returns the AttributeContract field if non-nil, zero value otherwise.

### GetAttributeContractOk

`func (o *SpBrowserSso) GetAttributeContractOk() (*SpBrowserSsoAttributeContract, bool)`

GetAttributeContractOk returns a tuple with the AttributeContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContract

`func (o *SpBrowserSso) SetAttributeContract(v SpBrowserSsoAttributeContract)`

SetAttributeContract sets AttributeContract field to given value.


### GetAdapterMappings

`func (o *SpBrowserSso) GetAdapterMappings() []IdpAdapterAssertionMapping`

GetAdapterMappings returns the AdapterMappings field if non-nil, zero value otherwise.

### GetAdapterMappingsOk

`func (o *SpBrowserSso) GetAdapterMappingsOk() (*[]IdpAdapterAssertionMapping, bool)`

GetAdapterMappingsOk returns a tuple with the AdapterMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterMappings

`func (o *SpBrowserSso) SetAdapterMappings(v []IdpAdapterAssertionMapping)`

SetAdapterMappings sets AdapterMappings field to given value.


### GetAuthenticationPolicyContractAssertionMappings

`func (o *SpBrowserSso) GetAuthenticationPolicyContractAssertionMappings() []AuthenticationPolicyContractAssertionMapping`

GetAuthenticationPolicyContractAssertionMappings returns the AuthenticationPolicyContractAssertionMappings field if non-nil, zero value otherwise.

### GetAuthenticationPolicyContractAssertionMappingsOk

`func (o *SpBrowserSso) GetAuthenticationPolicyContractAssertionMappingsOk() (*[]AuthenticationPolicyContractAssertionMapping, bool)`

GetAuthenticationPolicyContractAssertionMappingsOk returns a tuple with the AuthenticationPolicyContractAssertionMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationPolicyContractAssertionMappings

`func (o *SpBrowserSso) SetAuthenticationPolicyContractAssertionMappings(v []AuthenticationPolicyContractAssertionMapping)`

SetAuthenticationPolicyContractAssertionMappings sets AuthenticationPolicyContractAssertionMappings field to given value.

### HasAuthenticationPolicyContractAssertionMappings

`func (o *SpBrowserSso) HasAuthenticationPolicyContractAssertionMappings() bool`

HasAuthenticationPolicyContractAssertionMappings returns a boolean if a field has been set.

### GetAssertionLifetime

`func (o *SpBrowserSso) GetAssertionLifetime() AssertionLifetime`

GetAssertionLifetime returns the AssertionLifetime field if non-nil, zero value otherwise.

### GetAssertionLifetimeOk

`func (o *SpBrowserSso) GetAssertionLifetimeOk() (*AssertionLifetime, bool)`

GetAssertionLifetimeOk returns a tuple with the AssertionLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionLifetime

`func (o *SpBrowserSso) SetAssertionLifetime(v AssertionLifetime)`

SetAssertionLifetime sets AssertionLifetime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


