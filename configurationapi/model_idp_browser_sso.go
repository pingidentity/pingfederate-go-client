/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the IdpBrowserSso type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdpBrowserSso{}

// IdpBrowserSso The settings used to enable secure browser-based SSO to resources at your site.
type IdpBrowserSso struct {
	// The browser-based SSO protocol to use.
	Protocol             string                `json:"protocol" tfsdk:"protocol"`
	OidcProviderSettings *OIDCProviderSettings `json:"oidcProviderSettings,omitempty" tfsdk:"oidc_provider_settings"`
	// The profiles that are enabled for browser-based SSO. SAML 2.0 supports all profiles whereas SAML 1.x IdP connections support both IdP and SP (non-standard) initiated SSO. This is required for SAMLx.x Connections.
	EnabledProfiles []string `json:"enabledProfiles,omitempty" tfsdk:"enabled_profiles"`
	// The SAML bindings that are enabled for browser-based SSO. This is required for SAML 2.0 connections when the enabled profiles contain the SP-initiated SSO profile or either SLO profile. For SAML 1.x based connections, it is not used for SP Connections and it is optional for IdP Connections.
	IncomingBindings []string `json:"incomingBindings,omitempty" tfsdk:"incoming_bindings"`
	// The message customizations for browser-based SSO. Depending on server settings, connection type, and protocol this may or may not be supported.
	MessageCustomizations []ProtocolMessageCustomization `json:"messageCustomizations,omitempty" tfsdk:"message_customizations"`
	// For WS-Federation connections, a whitelist of additional allowed domains and paths used to validate wreply for SLO, if enabled.
	UrlWhitelistEntries []UrlWhitelistEntry `json:"urlWhitelistEntries,omitempty" tfsdk:"url_whitelist_entries"`
	Artifact            *ArtifactSettings   `json:"artifact,omitempty" tfsdk:"artifact"`
	// A list of possible endpoints to send SLO requests and responses.
	SloServiceEndpoints []SloServiceEndpoint `json:"sloServiceEndpoints,omitempty" tfsdk:"slo_service_endpoints"`
	// Specify to always sign the SAML ArtifactResponse.
	AlwaysSignArtifactResponse *bool `json:"alwaysSignArtifactResponse,omitempty" tfsdk:"always_sign_artifact_response"`
	// Application endpoint that can be used to invoke single sign-on (SSO) for the connection. This is a read-only parameter.
	SsoApplicationEndpoint *string `json:"ssoApplicationEndpoint,omitempty" tfsdk:"sso_application_endpoint"`
	// The IdP SSO endpoints that define where to send your authentication requests. Only required for SP initiated SSO. This is required for SAML x.x and WS-FED Connections.
	SsoServiceEndpoints []IdpSsoServiceEndpoint `json:"ssoServiceEndpoints,omitempty" tfsdk:"sso_service_endpoints"`
	// The default target URL for this connection. If defined, this overrides the default URL.
	DefaultTargetUrl *string `json:"defaultTargetUrl,omitempty" tfsdk:"default_target_url"`
	// A list of authentication context mappings between local and remote values. Applicable for SAML 2.0 and OIDC protocol connections.
	AuthnContextMappings []AuthnContextMapping `json:"authnContextMappings,omitempty" tfsdk:"authn_context_mappings"`
	// Specify whether the incoming SAML assertions are signed rather than the entire SAML response being signed.
	AssertionsSigned *bool `json:"assertionsSigned,omitempty" tfsdk:"assertions_signed"`
	// Determines whether SAML authentication requests should be signed.
	SignAuthnRequests *bool             `json:"signAuthnRequests,omitempty" tfsdk:"sign_authn_requests"`
	DecryptionPolicy  *DecryptionPolicy `json:"decryptionPolicy,omitempty" tfsdk:"decryption_policy"`
	// Defines the process in which users authenticated by the IdP are associated with user accounts local to the SP.
	IdpIdentityMapping string                          `json:"idpIdentityMapping" tfsdk:"idp_identity_mapping"`
	AttributeContract  *IdpBrowserSsoAttributeContract `json:"attributeContract,omitempty" tfsdk:"attribute_contract"`
	// A list of adapters that map to incoming assertions.
	AdapterMappings []SpAdapterMapping `json:"adapterMappings,omitempty" tfsdk:"adapter_mappings"`
	// A list of Authentication Policy Contracts that map to incoming assertions.
	AuthenticationPolicyContractMappings []AuthenticationPolicyContractMapping `json:"authenticationPolicyContractMappings,omitempty" tfsdk:"authentication_policy_contract_mappings"`
	SsoOAuthMapping                      *SsoOAuthMapping                      `json:"ssoOAuthMapping,omitempty" tfsdk:"sso_oauth_mapping"`
	OauthAuthenticationPolicyContractRef *ResourceLink                         `json:"oauthAuthenticationPolicyContractRef,omitempty" tfsdk:"oauth_authentication_policy_contract_ref"`
	JitProvisioning                      *JitProvisioning                      `json:"jitProvisioning,omitempty" tfsdk:"jit_provisioning"`
}

// NewIdpBrowserSso instantiates a new IdpBrowserSso object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpBrowserSso(protocol string, idpIdentityMapping string) *IdpBrowserSso {
	this := IdpBrowserSso{}
	this.Protocol = protocol
	this.IdpIdentityMapping = idpIdentityMapping
	return &this
}

// NewIdpBrowserSsoWithDefaults instantiates a new IdpBrowserSso object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpBrowserSsoWithDefaults() *IdpBrowserSso {
	this := IdpBrowserSso{}
	return &this
}

// GetProtocol returns the Protocol field value
func (o *IdpBrowserSso) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *IdpBrowserSso) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *IdpBrowserSso) SetProtocol(v string) {
	o.Protocol = v
}

// GetOidcProviderSettings returns the OidcProviderSettings field value if set, zero value otherwise.
func (o *IdpBrowserSso) GetOidcProviderSettings() OIDCProviderSettings {
	if o == nil || IsNil(o.OidcProviderSettings) {
		var ret OIDCProviderSettings
		return ret
	}
	return *o.OidcProviderSettings
}

// GetOidcProviderSettingsOk returns a tuple with the OidcProviderSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpBrowserSso) GetOidcProviderSettingsOk() (*OIDCProviderSettings, bool) {
	if o == nil || IsNil(o.OidcProviderSettings) {
		return nil, false
	}
	return o.OidcProviderSettings, true
}

// HasOidcProviderSettings returns a boolean if a field has been set.
func (o *IdpBrowserSso) HasOidcProviderSettings() bool {
	if o != nil && !IsNil(o.OidcProviderSettings) {
		return true
	}

	return false
}

// SetOidcProviderSettings gets a reference to the given OIDCProviderSettings and assigns it to the OidcProviderSettings field.
func (o *IdpBrowserSso) SetOidcProviderSettings(v OIDCProviderSettings) {
	o.OidcProviderSettings = &v
}

// GetEnabledProfiles returns the EnabledProfiles field value if set, zero value otherwise.
func (o *IdpBrowserSso) GetEnabledProfiles() []string {
	if o == nil || IsNil(o.EnabledProfiles) {
		var ret []string
		return ret
	}
	return o.EnabledProfiles
}

// GetEnabledProfilesOk returns a tuple with the EnabledProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpBrowserSso) GetEnabledProfilesOk() ([]string, bool) {
	if o == nil || IsNil(o.EnabledProfiles) {
		return nil, false
	}
	return o.EnabledProfiles, true
}

// HasEnabledProfiles returns a boolean if a field has been set.
func (o *IdpBrowserSso) HasEnabledProfiles() bool {
	if o != nil && !IsNil(o.EnabledProfiles) {
		return true
	}

	return false
}

// SetEnabledProfiles gets a reference to the given []string and assigns it to the EnabledProfiles field.
func (o *IdpBrowserSso) SetEnabledProfiles(v []string) {
	o.EnabledProfiles = v
}

// GetIncomingBindings returns the IncomingBindings field value if set, zero value otherwise.
func (o *IdpBrowserSso) GetIncomingBindings() []string {
	if o == nil || IsNil(o.IncomingBindings) {
		var ret []string
		return ret
	}
	return o.IncomingBindings
}

// GetIncomingBindingsOk returns a tuple with the IncomingBindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpBrowserSso) GetIncomingBindingsOk() ([]string, bool) {
	if o == nil || IsNil(o.IncomingBindings) {
		return nil, false
	}
	return o.IncomingBindings, true
}

// HasIncomingBindings returns a boolean if a field has been set.
func (o *IdpBrowserSso) HasIncomingBindings() bool {
	if o != nil && !IsNil(o.IncomingBindings) {
		return true
	}

	return false
}

// SetIncomingBindings gets a reference to the given []string and assigns it to the IncomingBindings field.
func (o *IdpBrowserSso) SetIncomingBindings(v []string) {
	o.IncomingBindings = v
}

// GetMessageCustomizations returns the MessageCustomizations field value if set, zero value otherwise.
func (o *IdpBrowserSso) GetMessageCustomizations() []ProtocolMessageCustomization {
	if o == nil || IsNil(o.MessageCustomizations) {
		var ret []ProtocolMessageCustomization
		return ret
	}
	return o.MessageCustomizations
}

// GetMessageCustomizationsOk returns a tuple with the MessageCustomizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpBrowserSso) GetMessageCustomizationsOk() ([]ProtocolMessageCustomization, bool) {
	if o == nil || IsNil(o.MessageCustomizations) {
		return nil, false
	}
	return o.MessageCustomizations, true
}

// HasMessageCustomizations returns a boolean if a field has been set.
func (o *IdpBrowserSso) HasMessageCustomizations() bool {
	if o != nil && !IsNil(o.MessageCustomizations) {
		return true
	}

	return false
}

// SetMessageCustomizations gets a reference to the given []ProtocolMessageCustomization and assigns it to the MessageCustomizations field.
func (o *IdpBrowserSso) SetMessageCustomizations(v []ProtocolMessageCustomization) {
	o.MessageCustomizations = v
}

// GetUrlWhitelistEntries returns the UrlWhitelistEntries field value if set, zero value otherwise.
func (o *IdpBrowserSso) GetUrlWhitelistEntries() []UrlWhitelistEntry {
	if o == nil || IsNil(o.UrlWhitelistEntries) {
		var ret []UrlWhitelistEntry
		return ret
	}
	return o.UrlWhitelistEntries
}

// GetUrlWhitelistEntriesOk returns a tuple with the UrlWhitelistEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpBrowserSso) GetUrlWhitelistEntriesOk() ([]UrlWhitelistEntry, bool) {
	if o == nil || IsNil(o.UrlWhitelistEntries) {
		return nil, false
	}
	return o.UrlWhitelistEntries, true
}

// HasUrlWhitelistEntries returns a boolean if a field has been set.
func (o *IdpBrowserSso) HasUrlWhitelistEntries() bool {
	if o != nil && !IsNil(o.UrlWhitelistEntries) {
		return true
	}

	return false
}

// SetUrlWhitelistEntries gets a reference to the given []UrlWhitelistEntry and assigns it to the UrlWhitelistEntries field.
func (o *IdpBrowserSso) SetUrlWhitelistEntries(v []UrlWhitelistEntry) {
	o.UrlWhitelistEntries = v
}

// GetArtifact returns the Artifact field value if set, zero value otherwise.
func (o *IdpBrowserSso) GetArtifact() ArtifactSettings {
	if o == nil || IsNil(o.Artifact) {
		var ret ArtifactSettings
		return ret
	}
	return *o.Artifact
}

// GetArtifactOk returns a tuple with the Artifact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpBrowserSso) GetArtifactOk() (*ArtifactSettings, bool) {
	if o == nil || IsNil(o.Artifact) {
		return nil, false
	}
	return o.Artifact, true
}

// HasArtifact returns a boolean if a field has been set.
func (o *IdpBrowserSso) HasArtifact() bool {
	if o != nil && !IsNil(o.Artifact) {
		return true
	}

	return false
}

// SetArtifact gets a reference to the given ArtifactSettings and assigns it to the Artifact field.
func (o *IdpBrowserSso) SetArtifact(v ArtifactSettings) {
	o.Artifact = &v
}

// GetSloServiceEndpoints returns the SloServiceEndpoints field value if set, zero value otherwise.
func (o *IdpBrowserSso) GetSloServiceEndpoints() []SloServiceEndpoint {
	if o == nil || IsNil(o.SloServiceEndpoints) {
		var ret []SloServiceEndpoint
		return ret
	}
	return o.SloServiceEndpoints
}

// GetSloServiceEndpointsOk returns a tuple with the SloServiceEndpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpBrowserSso) GetSloServiceEndpointsOk() ([]SloServiceEndpoint, bool) {
	if o == nil || IsNil(o.SloServiceEndpoints) {
		return nil, false
	}
	return o.SloServiceEndpoints, true
}

// HasSloServiceEndpoints returns a boolean if a field has been set.
func (o *IdpBrowserSso) HasSloServiceEndpoints() bool {
	if o != nil && !IsNil(o.SloServiceEndpoints) {
		return true
	}

	return false
}

// SetSloServiceEndpoints gets a reference to the given []SloServiceEndpoint and assigns it to the SloServiceEndpoints field.
func (o *IdpBrowserSso) SetSloServiceEndpoints(v []SloServiceEndpoint) {
	o.SloServiceEndpoints = v
}

// GetAlwaysSignArtifactResponse returns the AlwaysSignArtifactResponse field value if set, zero value otherwise.
func (o *IdpBrowserSso) GetAlwaysSignArtifactResponse() bool {
	if o == nil || IsNil(o.AlwaysSignArtifactResponse) {
		var ret bool
		return ret
	}
	return *o.AlwaysSignArtifactResponse
}

// GetAlwaysSignArtifactResponseOk returns a tuple with the AlwaysSignArtifactResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpBrowserSso) GetAlwaysSignArtifactResponseOk() (*bool, bool) {
	if o == nil || IsNil(o.AlwaysSignArtifactResponse) {
		return nil, false
	}
	return o.AlwaysSignArtifactResponse, true
}

// HasAlwaysSignArtifactResponse returns a boolean if a field has been set.
func (o *IdpBrowserSso) HasAlwaysSignArtifactResponse() bool {
	if o != nil && !IsNil(o.AlwaysSignArtifactResponse) {
		return true
	}

	return false
}

// SetAlwaysSignArtifactResponse gets a reference to the given bool and assigns it to the AlwaysSignArtifactResponse field.
func (o *IdpBrowserSso) SetAlwaysSignArtifactResponse(v bool) {
	o.AlwaysSignArtifactResponse = &v
}

// GetSsoApplicationEndpoint returns the SsoApplicationEndpoint field value if set, zero value otherwise.
func (o *IdpBrowserSso) GetSsoApplicationEndpoint() string {
	if o == nil || IsNil(o.SsoApplicationEndpoint) {
		var ret string
		return ret
	}
	return *o.SsoApplicationEndpoint
}

// GetSsoApplicationEndpointOk returns a tuple with the SsoApplicationEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpBrowserSso) GetSsoApplicationEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.SsoApplicationEndpoint) {
		return nil, false
	}
	return o.SsoApplicationEndpoint, true
}

// HasSsoApplicationEndpoint returns a boolean if a field has been set.
func (o *IdpBrowserSso) HasSsoApplicationEndpoint() bool {
	if o != nil && !IsNil(o.SsoApplicationEndpoint) {
		return true
	}

	return false
}

// SetSsoApplicationEndpoint gets a reference to the given string and assigns it to the SsoApplicationEndpoint field.
func (o *IdpBrowserSso) SetSsoApplicationEndpoint(v string) {
	o.SsoApplicationEndpoint = &v
}

// GetSsoServiceEndpoints returns the SsoServiceEndpoints field value if set, zero value otherwise.
func (o *IdpBrowserSso) GetSsoServiceEndpoints() []IdpSsoServiceEndpoint {
	if o == nil || IsNil(o.SsoServiceEndpoints) {
		var ret []IdpSsoServiceEndpoint
		return ret
	}
	return o.SsoServiceEndpoints
}

// GetSsoServiceEndpointsOk returns a tuple with the SsoServiceEndpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpBrowserSso) GetSsoServiceEndpointsOk() ([]IdpSsoServiceEndpoint, bool) {
	if o == nil || IsNil(o.SsoServiceEndpoints) {
		return nil, false
	}
	return o.SsoServiceEndpoints, true
}

// HasSsoServiceEndpoints returns a boolean if a field has been set.
func (o *IdpBrowserSso) HasSsoServiceEndpoints() bool {
	if o != nil && !IsNil(o.SsoServiceEndpoints) {
		return true
	}

	return false
}

// SetSsoServiceEndpoints gets a reference to the given []IdpSsoServiceEndpoint and assigns it to the SsoServiceEndpoints field.
func (o *IdpBrowserSso) SetSsoServiceEndpoints(v []IdpSsoServiceEndpoint) {
	o.SsoServiceEndpoints = v
}

// GetDefaultTargetUrl returns the DefaultTargetUrl field value if set, zero value otherwise.
func (o *IdpBrowserSso) GetDefaultTargetUrl() string {
	if o == nil || IsNil(o.DefaultTargetUrl) {
		var ret string
		return ret
	}
	return *o.DefaultTargetUrl
}

// GetDefaultTargetUrlOk returns a tuple with the DefaultTargetUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpBrowserSso) GetDefaultTargetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultTargetUrl) {
		return nil, false
	}
	return o.DefaultTargetUrl, true
}

// HasDefaultTargetUrl returns a boolean if a field has been set.
func (o *IdpBrowserSso) HasDefaultTargetUrl() bool {
	if o != nil && !IsNil(o.DefaultTargetUrl) {
		return true
	}

	return false
}

// SetDefaultTargetUrl gets a reference to the given string and assigns it to the DefaultTargetUrl field.
func (o *IdpBrowserSso) SetDefaultTargetUrl(v string) {
	o.DefaultTargetUrl = &v
}

// GetAuthnContextMappings returns the AuthnContextMappings field value if set, zero value otherwise.
func (o *IdpBrowserSso) GetAuthnContextMappings() []AuthnContextMapping {
	if o == nil || IsNil(o.AuthnContextMappings) {
		var ret []AuthnContextMapping
		return ret
	}
	return o.AuthnContextMappings
}

// GetAuthnContextMappingsOk returns a tuple with the AuthnContextMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpBrowserSso) GetAuthnContextMappingsOk() ([]AuthnContextMapping, bool) {
	if o == nil || IsNil(o.AuthnContextMappings) {
		return nil, false
	}
	return o.AuthnContextMappings, true
}

// HasAuthnContextMappings returns a boolean if a field has been set.
func (o *IdpBrowserSso) HasAuthnContextMappings() bool {
	if o != nil && !IsNil(o.AuthnContextMappings) {
		return true
	}

	return false
}

// SetAuthnContextMappings gets a reference to the given []AuthnContextMapping and assigns it to the AuthnContextMappings field.
func (o *IdpBrowserSso) SetAuthnContextMappings(v []AuthnContextMapping) {
	o.AuthnContextMappings = v
}

// GetAssertionsSigned returns the AssertionsSigned field value if set, zero value otherwise.
func (o *IdpBrowserSso) GetAssertionsSigned() bool {
	if o == nil || IsNil(o.AssertionsSigned) {
		var ret bool
		return ret
	}
	return *o.AssertionsSigned
}

// GetAssertionsSignedOk returns a tuple with the AssertionsSigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpBrowserSso) GetAssertionsSignedOk() (*bool, bool) {
	if o == nil || IsNil(o.AssertionsSigned) {
		return nil, false
	}
	return o.AssertionsSigned, true
}

// HasAssertionsSigned returns a boolean if a field has been set.
func (o *IdpBrowserSso) HasAssertionsSigned() bool {
	if o != nil && !IsNil(o.AssertionsSigned) {
		return true
	}

	return false
}

// SetAssertionsSigned gets a reference to the given bool and assigns it to the AssertionsSigned field.
func (o *IdpBrowserSso) SetAssertionsSigned(v bool) {
	o.AssertionsSigned = &v
}

// GetSignAuthnRequests returns the SignAuthnRequests field value if set, zero value otherwise.
func (o *IdpBrowserSso) GetSignAuthnRequests() bool {
	if o == nil || IsNil(o.SignAuthnRequests) {
		var ret bool
		return ret
	}
	return *o.SignAuthnRequests
}

// GetSignAuthnRequestsOk returns a tuple with the SignAuthnRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpBrowserSso) GetSignAuthnRequestsOk() (*bool, bool) {
	if o == nil || IsNil(o.SignAuthnRequests) {
		return nil, false
	}
	return o.SignAuthnRequests, true
}

// HasSignAuthnRequests returns a boolean if a field has been set.
func (o *IdpBrowserSso) HasSignAuthnRequests() bool {
	if o != nil && !IsNil(o.SignAuthnRequests) {
		return true
	}

	return false
}

// SetSignAuthnRequests gets a reference to the given bool and assigns it to the SignAuthnRequests field.
func (o *IdpBrowserSso) SetSignAuthnRequests(v bool) {
	o.SignAuthnRequests = &v
}

// GetDecryptionPolicy returns the DecryptionPolicy field value if set, zero value otherwise.
func (o *IdpBrowserSso) GetDecryptionPolicy() DecryptionPolicy {
	if o == nil || IsNil(o.DecryptionPolicy) {
		var ret DecryptionPolicy
		return ret
	}
	return *o.DecryptionPolicy
}

// GetDecryptionPolicyOk returns a tuple with the DecryptionPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpBrowserSso) GetDecryptionPolicyOk() (*DecryptionPolicy, bool) {
	if o == nil || IsNil(o.DecryptionPolicy) {
		return nil, false
	}
	return o.DecryptionPolicy, true
}

// HasDecryptionPolicy returns a boolean if a field has been set.
func (o *IdpBrowserSso) HasDecryptionPolicy() bool {
	if o != nil && !IsNil(o.DecryptionPolicy) {
		return true
	}

	return false
}

// SetDecryptionPolicy gets a reference to the given DecryptionPolicy and assigns it to the DecryptionPolicy field.
func (o *IdpBrowserSso) SetDecryptionPolicy(v DecryptionPolicy) {
	o.DecryptionPolicy = &v
}

// GetIdpIdentityMapping returns the IdpIdentityMapping field value
func (o *IdpBrowserSso) GetIdpIdentityMapping() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdpIdentityMapping
}

// GetIdpIdentityMappingOk returns a tuple with the IdpIdentityMapping field value
// and a boolean to check if the value has been set.
func (o *IdpBrowserSso) GetIdpIdentityMappingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdpIdentityMapping, true
}

// SetIdpIdentityMapping sets field value
func (o *IdpBrowserSso) SetIdpIdentityMapping(v string) {
	o.IdpIdentityMapping = v
}

// GetAttributeContract returns the AttributeContract field value if set, zero value otherwise.
func (o *IdpBrowserSso) GetAttributeContract() IdpBrowserSsoAttributeContract {
	if o == nil || IsNil(o.AttributeContract) {
		var ret IdpBrowserSsoAttributeContract
		return ret
	}
	return *o.AttributeContract
}

// GetAttributeContractOk returns a tuple with the AttributeContract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpBrowserSso) GetAttributeContractOk() (*IdpBrowserSsoAttributeContract, bool) {
	if o == nil || IsNil(o.AttributeContract) {
		return nil, false
	}
	return o.AttributeContract, true
}

// HasAttributeContract returns a boolean if a field has been set.
func (o *IdpBrowserSso) HasAttributeContract() bool {
	if o != nil && !IsNil(o.AttributeContract) {
		return true
	}

	return false
}

// SetAttributeContract gets a reference to the given IdpBrowserSsoAttributeContract and assigns it to the AttributeContract field.
func (o *IdpBrowserSso) SetAttributeContract(v IdpBrowserSsoAttributeContract) {
	o.AttributeContract = &v
}

// GetAdapterMappings returns the AdapterMappings field value if set, zero value otherwise.
func (o *IdpBrowserSso) GetAdapterMappings() []SpAdapterMapping {
	if o == nil || IsNil(o.AdapterMappings) {
		var ret []SpAdapterMapping
		return ret
	}
	return o.AdapterMappings
}

// GetAdapterMappingsOk returns a tuple with the AdapterMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpBrowserSso) GetAdapterMappingsOk() ([]SpAdapterMapping, bool) {
	if o == nil || IsNil(o.AdapterMappings) {
		return nil, false
	}
	return o.AdapterMappings, true
}

// HasAdapterMappings returns a boolean if a field has been set.
func (o *IdpBrowserSso) HasAdapterMappings() bool {
	if o != nil && !IsNil(o.AdapterMappings) {
		return true
	}

	return false
}

// SetAdapterMappings gets a reference to the given []SpAdapterMapping and assigns it to the AdapterMappings field.
func (o *IdpBrowserSso) SetAdapterMappings(v []SpAdapterMapping) {
	o.AdapterMappings = v
}

// GetAuthenticationPolicyContractMappings returns the AuthenticationPolicyContractMappings field value if set, zero value otherwise.
func (o *IdpBrowserSso) GetAuthenticationPolicyContractMappings() []AuthenticationPolicyContractMapping {
	if o == nil || IsNil(o.AuthenticationPolicyContractMappings) {
		var ret []AuthenticationPolicyContractMapping
		return ret
	}
	return o.AuthenticationPolicyContractMappings
}

// GetAuthenticationPolicyContractMappingsOk returns a tuple with the AuthenticationPolicyContractMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpBrowserSso) GetAuthenticationPolicyContractMappingsOk() ([]AuthenticationPolicyContractMapping, bool) {
	if o == nil || IsNil(o.AuthenticationPolicyContractMappings) {
		return nil, false
	}
	return o.AuthenticationPolicyContractMappings, true
}

// HasAuthenticationPolicyContractMappings returns a boolean if a field has been set.
func (o *IdpBrowserSso) HasAuthenticationPolicyContractMappings() bool {
	if o != nil && !IsNil(o.AuthenticationPolicyContractMappings) {
		return true
	}

	return false
}

// SetAuthenticationPolicyContractMappings gets a reference to the given []AuthenticationPolicyContractMapping and assigns it to the AuthenticationPolicyContractMappings field.
func (o *IdpBrowserSso) SetAuthenticationPolicyContractMappings(v []AuthenticationPolicyContractMapping) {
	o.AuthenticationPolicyContractMappings = v
}

// GetSsoOAuthMapping returns the SsoOAuthMapping field value if set, zero value otherwise.
func (o *IdpBrowserSso) GetSsoOAuthMapping() SsoOAuthMapping {
	if o == nil || IsNil(o.SsoOAuthMapping) {
		var ret SsoOAuthMapping
		return ret
	}
	return *o.SsoOAuthMapping
}

// GetSsoOAuthMappingOk returns a tuple with the SsoOAuthMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpBrowserSso) GetSsoOAuthMappingOk() (*SsoOAuthMapping, bool) {
	if o == nil || IsNil(o.SsoOAuthMapping) {
		return nil, false
	}
	return o.SsoOAuthMapping, true
}

// HasSsoOAuthMapping returns a boolean if a field has been set.
func (o *IdpBrowserSso) HasSsoOAuthMapping() bool {
	if o != nil && !IsNil(o.SsoOAuthMapping) {
		return true
	}

	return false
}

// SetSsoOAuthMapping gets a reference to the given SsoOAuthMapping and assigns it to the SsoOAuthMapping field.
func (o *IdpBrowserSso) SetSsoOAuthMapping(v SsoOAuthMapping) {
	o.SsoOAuthMapping = &v
}

// GetOauthAuthenticationPolicyContractRef returns the OauthAuthenticationPolicyContractRef field value if set, zero value otherwise.
func (o *IdpBrowserSso) GetOauthAuthenticationPolicyContractRef() ResourceLink {
	if o == nil || IsNil(o.OauthAuthenticationPolicyContractRef) {
		var ret ResourceLink
		return ret
	}
	return *o.OauthAuthenticationPolicyContractRef
}

// GetOauthAuthenticationPolicyContractRefOk returns a tuple with the OauthAuthenticationPolicyContractRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpBrowserSso) GetOauthAuthenticationPolicyContractRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.OauthAuthenticationPolicyContractRef) {
		return nil, false
	}
	return o.OauthAuthenticationPolicyContractRef, true
}

// HasOauthAuthenticationPolicyContractRef returns a boolean if a field has been set.
func (o *IdpBrowserSso) HasOauthAuthenticationPolicyContractRef() bool {
	if o != nil && !IsNil(o.OauthAuthenticationPolicyContractRef) {
		return true
	}

	return false
}

// SetOauthAuthenticationPolicyContractRef gets a reference to the given ResourceLink and assigns it to the OauthAuthenticationPolicyContractRef field.
func (o *IdpBrowserSso) SetOauthAuthenticationPolicyContractRef(v ResourceLink) {
	o.OauthAuthenticationPolicyContractRef = &v
}

// GetJitProvisioning returns the JitProvisioning field value if set, zero value otherwise.
func (o *IdpBrowserSso) GetJitProvisioning() JitProvisioning {
	if o == nil || IsNil(o.JitProvisioning) {
		var ret JitProvisioning
		return ret
	}
	return *o.JitProvisioning
}

// GetJitProvisioningOk returns a tuple with the JitProvisioning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpBrowserSso) GetJitProvisioningOk() (*JitProvisioning, bool) {
	if o == nil || IsNil(o.JitProvisioning) {
		return nil, false
	}
	return o.JitProvisioning, true
}

// HasJitProvisioning returns a boolean if a field has been set.
func (o *IdpBrowserSso) HasJitProvisioning() bool {
	if o != nil && !IsNil(o.JitProvisioning) {
		return true
	}

	return false
}

// SetJitProvisioning gets a reference to the given JitProvisioning and assigns it to the JitProvisioning field.
func (o *IdpBrowserSso) SetJitProvisioning(v JitProvisioning) {
	o.JitProvisioning = &v
}

func (o IdpBrowserSso) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdpBrowserSso) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["protocol"] = o.Protocol
	if !IsNil(o.OidcProviderSettings) {
		toSerialize["oidcProviderSettings"] = o.OidcProviderSettings
	}
	if !IsNil(o.EnabledProfiles) {
		toSerialize["enabledProfiles"] = o.EnabledProfiles
	}
	if !IsNil(o.IncomingBindings) {
		toSerialize["incomingBindings"] = o.IncomingBindings
	}
	if !IsNil(o.MessageCustomizations) {
		toSerialize["messageCustomizations"] = o.MessageCustomizations
	}
	if !IsNil(o.UrlWhitelistEntries) {
		toSerialize["urlWhitelistEntries"] = o.UrlWhitelistEntries
	}
	if !IsNil(o.Artifact) {
		toSerialize["artifact"] = o.Artifact
	}
	if !IsNil(o.SloServiceEndpoints) {
		toSerialize["sloServiceEndpoints"] = o.SloServiceEndpoints
	}
	if !IsNil(o.AlwaysSignArtifactResponse) {
		toSerialize["alwaysSignArtifactResponse"] = o.AlwaysSignArtifactResponse
	}
	if !IsNil(o.SsoApplicationEndpoint) {
		toSerialize["ssoApplicationEndpoint"] = o.SsoApplicationEndpoint
	}
	if !IsNil(o.SsoServiceEndpoints) {
		toSerialize["ssoServiceEndpoints"] = o.SsoServiceEndpoints
	}
	if !IsNil(o.DefaultTargetUrl) {
		toSerialize["defaultTargetUrl"] = o.DefaultTargetUrl
	}
	if !IsNil(o.AuthnContextMappings) {
		toSerialize["authnContextMappings"] = o.AuthnContextMappings
	}
	if !IsNil(o.AssertionsSigned) {
		toSerialize["assertionsSigned"] = o.AssertionsSigned
	}
	if !IsNil(o.SignAuthnRequests) {
		toSerialize["signAuthnRequests"] = o.SignAuthnRequests
	}
	if !IsNil(o.DecryptionPolicy) {
		toSerialize["decryptionPolicy"] = o.DecryptionPolicy
	}
	toSerialize["idpIdentityMapping"] = o.IdpIdentityMapping
	if !IsNil(o.AttributeContract) {
		toSerialize["attributeContract"] = o.AttributeContract
	}
	if !IsNil(o.AdapterMappings) {
		toSerialize["adapterMappings"] = o.AdapterMappings
	}
	if !IsNil(o.AuthenticationPolicyContractMappings) {
		toSerialize["authenticationPolicyContractMappings"] = o.AuthenticationPolicyContractMappings
	}
	if !IsNil(o.SsoOAuthMapping) {
		toSerialize["ssoOAuthMapping"] = o.SsoOAuthMapping
	}
	if !IsNil(o.OauthAuthenticationPolicyContractRef) {
		toSerialize["oauthAuthenticationPolicyContractRef"] = o.OauthAuthenticationPolicyContractRef
	}
	if !IsNil(o.JitProvisioning) {
		toSerialize["jitProvisioning"] = o.JitProvisioning
	}
	return toSerialize, nil
}

type NullableIdpBrowserSso struct {
	value *IdpBrowserSso
	isSet bool
}

func (v NullableIdpBrowserSso) Get() *IdpBrowserSso {
	return v.value
}

func (v *NullableIdpBrowserSso) Set(val *IdpBrowserSso) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpBrowserSso) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpBrowserSso) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpBrowserSso(val *IdpBrowserSso) *NullableIdpBrowserSso {
	return &NullableIdpBrowserSso{value: val, isSet: true}
}

func (v NullableIdpBrowserSso) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpBrowserSso) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
