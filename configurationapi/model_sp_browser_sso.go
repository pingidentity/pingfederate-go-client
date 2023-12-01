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

// checks if the SpBrowserSso type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpBrowserSso{}

// SpBrowserSso The SAML settings used to enable secure browser-based SSO to resources at your partner's site.
type SpBrowserSso struct {
	// The browser-based SSO protocol to use.
	Protocol string `json:"protocol" tfsdk:"protocol"`
	// The WS-Federation Token Type to use.
	WsFedTokenType *string `json:"wsFedTokenType,omitempty" tfsdk:"ws_fed_token_type"`
	// The WS-Trust version for a WS-Federation connection. The default version is WSTRUST12.
	WsTrustVersion *string `json:"wsTrustVersion,omitempty" tfsdk:"ws_trust_version"`
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
	// Default Target URL for SAML1.x connections. For SP connections, this default URL represents the destination on the SP where the user will be directed. For IdP connections, entering a URL in the Default Target URL field overrides the SP Default URL SSO setting.
	DefaultTargetUrl *string `json:"defaultTargetUrl,omitempty" tfsdk:"default_target_url"`
	// Specify to always sign the SAML ArtifactResponse.
	AlwaysSignArtifactResponse *bool `json:"alwaysSignArtifactResponse,omitempty" tfsdk:"always_sign_artifact_response"`
	// A list of possible endpoints to send assertions to.
	SsoServiceEndpoints []SpSsoServiceEndpoint `json:"ssoServiceEndpoints" tfsdk:"sso_service_endpoints"`
	// Process in which users authenticated by the IdP are associated with user accounts local to the SP.
	SpSamlIdentityMapping *string `json:"spSamlIdentityMapping,omitempty" tfsdk:"sp_saml_identity_mapping"`
	// Process in which users authenticated by the IdP are associated with user accounts local to the SP for WS-Federation connection types.
	SpWsFedIdentityMapping *string `json:"spWsFedIdentityMapping,omitempty" tfsdk:"sp_ws_fed_identity_mapping"`
	// Sign SAML Response as required by the associated binding and encryption policy. Applicable to SAML2.0 only and is defaulted to true. It can be set to false only on SAML2.0 connections when signAssertions is set to true.
	SignResponseAsRequired *bool `json:"signResponseAsRequired,omitempty" tfsdk:"sign_response_as_required"`
	// Always sign the SAML Assertion.
	SignAssertions *bool `json:"signAssertions,omitempty" tfsdk:"sign_assertions"`
	// Require AuthN requests to be signed when received via the POST or Redirect bindings.
	RequireSignedAuthnRequests *bool                         `json:"requireSignedAuthnRequests,omitempty" tfsdk:"require_signed_authn_requests"`
	EncryptionPolicy           *EncryptionPolicy             `json:"encryptionPolicy,omitempty" tfsdk:"encryption_policy"`
	AttributeContract          SpBrowserSsoAttributeContract `json:"attributeContract" tfsdk:"attribute_contract"`
	// A list of adapters that map to outgoing assertions.
	AdapterMappings []IdpAdapterAssertionMapping `json:"adapterMappings" tfsdk:"adapter_mappings"`
	// A list of authentication policy contracts that map to outgoing assertions.
	AuthenticationPolicyContractAssertionMappings []AuthenticationPolicyContractAssertionMapping `json:"authenticationPolicyContractAssertionMappings,omitempty" tfsdk:"authentication_policy_contract_assertion_mappings"`
	AssertionLifetime                             AssertionLifetime                              `json:"assertionLifetime" tfsdk:"assertion_lifetime"`
}

// NewSpBrowserSso instantiates a new SpBrowserSso object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpBrowserSso(protocol string, ssoServiceEndpoints []SpSsoServiceEndpoint, attributeContract SpBrowserSsoAttributeContract, adapterMappings []IdpAdapterAssertionMapping, assertionLifetime AssertionLifetime) *SpBrowserSso {
	this := SpBrowserSso{}
	this.Protocol = protocol
	this.SsoServiceEndpoints = ssoServiceEndpoints
	this.AttributeContract = attributeContract
	this.AdapterMappings = adapterMappings
	this.AssertionLifetime = assertionLifetime
	return &this
}

// NewSpBrowserSsoWithDefaults instantiates a new SpBrowserSso object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpBrowserSsoWithDefaults() *SpBrowserSso {
	this := SpBrowserSso{}
	return &this
}

// GetProtocol returns the Protocol field value
func (o *SpBrowserSso) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *SpBrowserSso) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *SpBrowserSso) SetProtocol(v string) {
	o.Protocol = v
}

// GetWsFedTokenType returns the WsFedTokenType field value if set, zero value otherwise.
func (o *SpBrowserSso) GetWsFedTokenType() string {
	if o == nil || IsNil(o.WsFedTokenType) {
		var ret string
		return ret
	}
	return *o.WsFedTokenType
}

// GetWsFedTokenTypeOk returns a tuple with the WsFedTokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpBrowserSso) GetWsFedTokenTypeOk() (*string, bool) {
	if o == nil || IsNil(o.WsFedTokenType) {
		return nil, false
	}
	return o.WsFedTokenType, true
}

// HasWsFedTokenType returns a boolean if a field has been set.
func (o *SpBrowserSso) HasWsFedTokenType() bool {
	if o != nil && !IsNil(o.WsFedTokenType) {
		return true
	}

	return false
}

// SetWsFedTokenType gets a reference to the given string and assigns it to the WsFedTokenType field.
func (o *SpBrowserSso) SetWsFedTokenType(v string) {
	o.WsFedTokenType = &v
}

// GetWsTrustVersion returns the WsTrustVersion field value if set, zero value otherwise.
func (o *SpBrowserSso) GetWsTrustVersion() string {
	if o == nil || IsNil(o.WsTrustVersion) {
		var ret string
		return ret
	}
	return *o.WsTrustVersion
}

// GetWsTrustVersionOk returns a tuple with the WsTrustVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpBrowserSso) GetWsTrustVersionOk() (*string, bool) {
	if o == nil || IsNil(o.WsTrustVersion) {
		return nil, false
	}
	return o.WsTrustVersion, true
}

// HasWsTrustVersion returns a boolean if a field has been set.
func (o *SpBrowserSso) HasWsTrustVersion() bool {
	if o != nil && !IsNil(o.WsTrustVersion) {
		return true
	}

	return false
}

// SetWsTrustVersion gets a reference to the given string and assigns it to the WsTrustVersion field.
func (o *SpBrowserSso) SetWsTrustVersion(v string) {
	o.WsTrustVersion = &v
}

// GetEnabledProfiles returns the EnabledProfiles field value if set, zero value otherwise.
func (o *SpBrowserSso) GetEnabledProfiles() []string {
	if o == nil || IsNil(o.EnabledProfiles) {
		var ret []string
		return ret
	}
	return o.EnabledProfiles
}

// GetEnabledProfilesOk returns a tuple with the EnabledProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpBrowserSso) GetEnabledProfilesOk() ([]string, bool) {
	if o == nil || IsNil(o.EnabledProfiles) {
		return nil, false
	}
	return o.EnabledProfiles, true
}

// HasEnabledProfiles returns a boolean if a field has been set.
func (o *SpBrowserSso) HasEnabledProfiles() bool {
	if o != nil && !IsNil(o.EnabledProfiles) {
		return true
	}

	return false
}

// SetEnabledProfiles gets a reference to the given []string and assigns it to the EnabledProfiles field.
func (o *SpBrowserSso) SetEnabledProfiles(v []string) {
	o.EnabledProfiles = v
}

// GetIncomingBindings returns the IncomingBindings field value if set, zero value otherwise.
func (o *SpBrowserSso) GetIncomingBindings() []string {
	if o == nil || IsNil(o.IncomingBindings) {
		var ret []string
		return ret
	}
	return o.IncomingBindings
}

// GetIncomingBindingsOk returns a tuple with the IncomingBindings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpBrowserSso) GetIncomingBindingsOk() ([]string, bool) {
	if o == nil || IsNil(o.IncomingBindings) {
		return nil, false
	}
	return o.IncomingBindings, true
}

// HasIncomingBindings returns a boolean if a field has been set.
func (o *SpBrowserSso) HasIncomingBindings() bool {
	if o != nil && !IsNil(o.IncomingBindings) {
		return true
	}

	return false
}

// SetIncomingBindings gets a reference to the given []string and assigns it to the IncomingBindings field.
func (o *SpBrowserSso) SetIncomingBindings(v []string) {
	o.IncomingBindings = v
}

// GetMessageCustomizations returns the MessageCustomizations field value if set, zero value otherwise.
func (o *SpBrowserSso) GetMessageCustomizations() []ProtocolMessageCustomization {
	if o == nil || IsNil(o.MessageCustomizations) {
		var ret []ProtocolMessageCustomization
		return ret
	}
	return o.MessageCustomizations
}

// GetMessageCustomizationsOk returns a tuple with the MessageCustomizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpBrowserSso) GetMessageCustomizationsOk() ([]ProtocolMessageCustomization, bool) {
	if o == nil || IsNil(o.MessageCustomizations) {
		return nil, false
	}
	return o.MessageCustomizations, true
}

// HasMessageCustomizations returns a boolean if a field has been set.
func (o *SpBrowserSso) HasMessageCustomizations() bool {
	if o != nil && !IsNil(o.MessageCustomizations) {
		return true
	}

	return false
}

// SetMessageCustomizations gets a reference to the given []ProtocolMessageCustomization and assigns it to the MessageCustomizations field.
func (o *SpBrowserSso) SetMessageCustomizations(v []ProtocolMessageCustomization) {
	o.MessageCustomizations = v
}

// GetUrlWhitelistEntries returns the UrlWhitelistEntries field value if set, zero value otherwise.
func (o *SpBrowserSso) GetUrlWhitelistEntries() []UrlWhitelistEntry {
	if o == nil || IsNil(o.UrlWhitelistEntries) {
		var ret []UrlWhitelistEntry
		return ret
	}
	return o.UrlWhitelistEntries
}

// GetUrlWhitelistEntriesOk returns a tuple with the UrlWhitelistEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpBrowserSso) GetUrlWhitelistEntriesOk() ([]UrlWhitelistEntry, bool) {
	if o == nil || IsNil(o.UrlWhitelistEntries) {
		return nil, false
	}
	return o.UrlWhitelistEntries, true
}

// HasUrlWhitelistEntries returns a boolean if a field has been set.
func (o *SpBrowserSso) HasUrlWhitelistEntries() bool {
	if o != nil && !IsNil(o.UrlWhitelistEntries) {
		return true
	}

	return false
}

// SetUrlWhitelistEntries gets a reference to the given []UrlWhitelistEntry and assigns it to the UrlWhitelistEntries field.
func (o *SpBrowserSso) SetUrlWhitelistEntries(v []UrlWhitelistEntry) {
	o.UrlWhitelistEntries = v
}

// GetArtifact returns the Artifact field value if set, zero value otherwise.
func (o *SpBrowserSso) GetArtifact() ArtifactSettings {
	if o == nil || IsNil(o.Artifact) {
		var ret ArtifactSettings
		return ret
	}
	return *o.Artifact
}

// GetArtifactOk returns a tuple with the Artifact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpBrowserSso) GetArtifactOk() (*ArtifactSettings, bool) {
	if o == nil || IsNil(o.Artifact) {
		return nil, false
	}
	return o.Artifact, true
}

// HasArtifact returns a boolean if a field has been set.
func (o *SpBrowserSso) HasArtifact() bool {
	if o != nil && !IsNil(o.Artifact) {
		return true
	}

	return false
}

// SetArtifact gets a reference to the given ArtifactSettings and assigns it to the Artifact field.
func (o *SpBrowserSso) SetArtifact(v ArtifactSettings) {
	o.Artifact = &v
}

// GetSloServiceEndpoints returns the SloServiceEndpoints field value if set, zero value otherwise.
func (o *SpBrowserSso) GetSloServiceEndpoints() []SloServiceEndpoint {
	if o == nil || IsNil(o.SloServiceEndpoints) {
		var ret []SloServiceEndpoint
		return ret
	}
	return o.SloServiceEndpoints
}

// GetSloServiceEndpointsOk returns a tuple with the SloServiceEndpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpBrowserSso) GetSloServiceEndpointsOk() ([]SloServiceEndpoint, bool) {
	if o == nil || IsNil(o.SloServiceEndpoints) {
		return nil, false
	}
	return o.SloServiceEndpoints, true
}

// HasSloServiceEndpoints returns a boolean if a field has been set.
func (o *SpBrowserSso) HasSloServiceEndpoints() bool {
	if o != nil && !IsNil(o.SloServiceEndpoints) {
		return true
	}

	return false
}

// SetSloServiceEndpoints gets a reference to the given []SloServiceEndpoint and assigns it to the SloServiceEndpoints field.
func (o *SpBrowserSso) SetSloServiceEndpoints(v []SloServiceEndpoint) {
	o.SloServiceEndpoints = v
}

// GetDefaultTargetUrl returns the DefaultTargetUrl field value if set, zero value otherwise.
func (o *SpBrowserSso) GetDefaultTargetUrl() string {
	if o == nil || IsNil(o.DefaultTargetUrl) {
		var ret string
		return ret
	}
	return *o.DefaultTargetUrl
}

// GetDefaultTargetUrlOk returns a tuple with the DefaultTargetUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpBrowserSso) GetDefaultTargetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultTargetUrl) {
		return nil, false
	}
	return o.DefaultTargetUrl, true
}

// HasDefaultTargetUrl returns a boolean if a field has been set.
func (o *SpBrowserSso) HasDefaultTargetUrl() bool {
	if o != nil && !IsNil(o.DefaultTargetUrl) {
		return true
	}

	return false
}

// SetDefaultTargetUrl gets a reference to the given string and assigns it to the DefaultTargetUrl field.
func (o *SpBrowserSso) SetDefaultTargetUrl(v string) {
	o.DefaultTargetUrl = &v
}

// GetAlwaysSignArtifactResponse returns the AlwaysSignArtifactResponse field value if set, zero value otherwise.
func (o *SpBrowserSso) GetAlwaysSignArtifactResponse() bool {
	if o == nil || IsNil(o.AlwaysSignArtifactResponse) {
		var ret bool
		return ret
	}
	return *o.AlwaysSignArtifactResponse
}

// GetAlwaysSignArtifactResponseOk returns a tuple with the AlwaysSignArtifactResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpBrowserSso) GetAlwaysSignArtifactResponseOk() (*bool, bool) {
	if o == nil || IsNil(o.AlwaysSignArtifactResponse) {
		return nil, false
	}
	return o.AlwaysSignArtifactResponse, true
}

// HasAlwaysSignArtifactResponse returns a boolean if a field has been set.
func (o *SpBrowserSso) HasAlwaysSignArtifactResponse() bool {
	if o != nil && !IsNil(o.AlwaysSignArtifactResponse) {
		return true
	}

	return false
}

// SetAlwaysSignArtifactResponse gets a reference to the given bool and assigns it to the AlwaysSignArtifactResponse field.
func (o *SpBrowserSso) SetAlwaysSignArtifactResponse(v bool) {
	o.AlwaysSignArtifactResponse = &v
}

// GetSsoServiceEndpoints returns the SsoServiceEndpoints field value
func (o *SpBrowserSso) GetSsoServiceEndpoints() []SpSsoServiceEndpoint {
	if o == nil {
		var ret []SpSsoServiceEndpoint
		return ret
	}

	return o.SsoServiceEndpoints
}

// GetSsoServiceEndpointsOk returns a tuple with the SsoServiceEndpoints field value
// and a boolean to check if the value has been set.
func (o *SpBrowserSso) GetSsoServiceEndpointsOk() ([]SpSsoServiceEndpoint, bool) {
	if o == nil {
		return nil, false
	}
	return o.SsoServiceEndpoints, true
}

// SetSsoServiceEndpoints sets field value
func (o *SpBrowserSso) SetSsoServiceEndpoints(v []SpSsoServiceEndpoint) {
	o.SsoServiceEndpoints = v
}

// GetSpSamlIdentityMapping returns the SpSamlIdentityMapping field value if set, zero value otherwise.
func (o *SpBrowserSso) GetSpSamlIdentityMapping() string {
	if o == nil || IsNil(o.SpSamlIdentityMapping) {
		var ret string
		return ret
	}
	return *o.SpSamlIdentityMapping
}

// GetSpSamlIdentityMappingOk returns a tuple with the SpSamlIdentityMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpBrowserSso) GetSpSamlIdentityMappingOk() (*string, bool) {
	if o == nil || IsNil(o.SpSamlIdentityMapping) {
		return nil, false
	}
	return o.SpSamlIdentityMapping, true
}

// HasSpSamlIdentityMapping returns a boolean if a field has been set.
func (o *SpBrowserSso) HasSpSamlIdentityMapping() bool {
	if o != nil && !IsNil(o.SpSamlIdentityMapping) {
		return true
	}

	return false
}

// SetSpSamlIdentityMapping gets a reference to the given string and assigns it to the SpSamlIdentityMapping field.
func (o *SpBrowserSso) SetSpSamlIdentityMapping(v string) {
	o.SpSamlIdentityMapping = &v
}

// GetSpWsFedIdentityMapping returns the SpWsFedIdentityMapping field value if set, zero value otherwise.
func (o *SpBrowserSso) GetSpWsFedIdentityMapping() string {
	if o == nil || IsNil(o.SpWsFedIdentityMapping) {
		var ret string
		return ret
	}
	return *o.SpWsFedIdentityMapping
}

// GetSpWsFedIdentityMappingOk returns a tuple with the SpWsFedIdentityMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpBrowserSso) GetSpWsFedIdentityMappingOk() (*string, bool) {
	if o == nil || IsNil(o.SpWsFedIdentityMapping) {
		return nil, false
	}
	return o.SpWsFedIdentityMapping, true
}

// HasSpWsFedIdentityMapping returns a boolean if a field has been set.
func (o *SpBrowserSso) HasSpWsFedIdentityMapping() bool {
	if o != nil && !IsNil(o.SpWsFedIdentityMapping) {
		return true
	}

	return false
}

// SetSpWsFedIdentityMapping gets a reference to the given string and assigns it to the SpWsFedIdentityMapping field.
func (o *SpBrowserSso) SetSpWsFedIdentityMapping(v string) {
	o.SpWsFedIdentityMapping = &v
}

// GetSignResponseAsRequired returns the SignResponseAsRequired field value if set, zero value otherwise.
func (o *SpBrowserSso) GetSignResponseAsRequired() bool {
	if o == nil || IsNil(o.SignResponseAsRequired) {
		var ret bool
		return ret
	}
	return *o.SignResponseAsRequired
}

// GetSignResponseAsRequiredOk returns a tuple with the SignResponseAsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpBrowserSso) GetSignResponseAsRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.SignResponseAsRequired) {
		return nil, false
	}
	return o.SignResponseAsRequired, true
}

// HasSignResponseAsRequired returns a boolean if a field has been set.
func (o *SpBrowserSso) HasSignResponseAsRequired() bool {
	if o != nil && !IsNil(o.SignResponseAsRequired) {
		return true
	}

	return false
}

// SetSignResponseAsRequired gets a reference to the given bool and assigns it to the SignResponseAsRequired field.
func (o *SpBrowserSso) SetSignResponseAsRequired(v bool) {
	o.SignResponseAsRequired = &v
}

// GetSignAssertions returns the SignAssertions field value if set, zero value otherwise.
func (o *SpBrowserSso) GetSignAssertions() bool {
	if o == nil || IsNil(o.SignAssertions) {
		var ret bool
		return ret
	}
	return *o.SignAssertions
}

// GetSignAssertionsOk returns a tuple with the SignAssertions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpBrowserSso) GetSignAssertionsOk() (*bool, bool) {
	if o == nil || IsNil(o.SignAssertions) {
		return nil, false
	}
	return o.SignAssertions, true
}

// HasSignAssertions returns a boolean if a field has been set.
func (o *SpBrowserSso) HasSignAssertions() bool {
	if o != nil && !IsNil(o.SignAssertions) {
		return true
	}

	return false
}

// SetSignAssertions gets a reference to the given bool and assigns it to the SignAssertions field.
func (o *SpBrowserSso) SetSignAssertions(v bool) {
	o.SignAssertions = &v
}

// GetRequireSignedAuthnRequests returns the RequireSignedAuthnRequests field value if set, zero value otherwise.
func (o *SpBrowserSso) GetRequireSignedAuthnRequests() bool {
	if o == nil || IsNil(o.RequireSignedAuthnRequests) {
		var ret bool
		return ret
	}
	return *o.RequireSignedAuthnRequests
}

// GetRequireSignedAuthnRequestsOk returns a tuple with the RequireSignedAuthnRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpBrowserSso) GetRequireSignedAuthnRequestsOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireSignedAuthnRequests) {
		return nil, false
	}
	return o.RequireSignedAuthnRequests, true
}

// HasRequireSignedAuthnRequests returns a boolean if a field has been set.
func (o *SpBrowserSso) HasRequireSignedAuthnRequests() bool {
	if o != nil && !IsNil(o.RequireSignedAuthnRequests) {
		return true
	}

	return false
}

// SetRequireSignedAuthnRequests gets a reference to the given bool and assigns it to the RequireSignedAuthnRequests field.
func (o *SpBrowserSso) SetRequireSignedAuthnRequests(v bool) {
	o.RequireSignedAuthnRequests = &v
}

// GetEncryptionPolicy returns the EncryptionPolicy field value if set, zero value otherwise.
func (o *SpBrowserSso) GetEncryptionPolicy() EncryptionPolicy {
	if o == nil || IsNil(o.EncryptionPolicy) {
		var ret EncryptionPolicy
		return ret
	}
	return *o.EncryptionPolicy
}

// GetEncryptionPolicyOk returns a tuple with the EncryptionPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpBrowserSso) GetEncryptionPolicyOk() (*EncryptionPolicy, bool) {
	if o == nil || IsNil(o.EncryptionPolicy) {
		return nil, false
	}
	return o.EncryptionPolicy, true
}

// HasEncryptionPolicy returns a boolean if a field has been set.
func (o *SpBrowserSso) HasEncryptionPolicy() bool {
	if o != nil && !IsNil(o.EncryptionPolicy) {
		return true
	}

	return false
}

// SetEncryptionPolicy gets a reference to the given EncryptionPolicy and assigns it to the EncryptionPolicy field.
func (o *SpBrowserSso) SetEncryptionPolicy(v EncryptionPolicy) {
	o.EncryptionPolicy = &v
}

// GetAttributeContract returns the AttributeContract field value
func (o *SpBrowserSso) GetAttributeContract() SpBrowserSsoAttributeContract {
	if o == nil {
		var ret SpBrowserSsoAttributeContract
		return ret
	}

	return o.AttributeContract
}

// GetAttributeContractOk returns a tuple with the AttributeContract field value
// and a boolean to check if the value has been set.
func (o *SpBrowserSso) GetAttributeContractOk() (*SpBrowserSsoAttributeContract, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeContract, true
}

// SetAttributeContract sets field value
func (o *SpBrowserSso) SetAttributeContract(v SpBrowserSsoAttributeContract) {
	o.AttributeContract = v
}

// GetAdapterMappings returns the AdapterMappings field value
func (o *SpBrowserSso) GetAdapterMappings() []IdpAdapterAssertionMapping {
	if o == nil {
		var ret []IdpAdapterAssertionMapping
		return ret
	}

	return o.AdapterMappings
}

// GetAdapterMappingsOk returns a tuple with the AdapterMappings field value
// and a boolean to check if the value has been set.
func (o *SpBrowserSso) GetAdapterMappingsOk() ([]IdpAdapterAssertionMapping, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdapterMappings, true
}

// SetAdapterMappings sets field value
func (o *SpBrowserSso) SetAdapterMappings(v []IdpAdapterAssertionMapping) {
	o.AdapterMappings = v
}

// GetAuthenticationPolicyContractAssertionMappings returns the AuthenticationPolicyContractAssertionMappings field value if set, zero value otherwise.
func (o *SpBrowserSso) GetAuthenticationPolicyContractAssertionMappings() []AuthenticationPolicyContractAssertionMapping {
	if o == nil || IsNil(o.AuthenticationPolicyContractAssertionMappings) {
		var ret []AuthenticationPolicyContractAssertionMapping
		return ret
	}
	return o.AuthenticationPolicyContractAssertionMappings
}

// GetAuthenticationPolicyContractAssertionMappingsOk returns a tuple with the AuthenticationPolicyContractAssertionMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpBrowserSso) GetAuthenticationPolicyContractAssertionMappingsOk() ([]AuthenticationPolicyContractAssertionMapping, bool) {
	if o == nil || IsNil(o.AuthenticationPolicyContractAssertionMappings) {
		return nil, false
	}
	return o.AuthenticationPolicyContractAssertionMappings, true
}

// HasAuthenticationPolicyContractAssertionMappings returns a boolean if a field has been set.
func (o *SpBrowserSso) HasAuthenticationPolicyContractAssertionMappings() bool {
	if o != nil && !IsNil(o.AuthenticationPolicyContractAssertionMappings) {
		return true
	}

	return false
}

// SetAuthenticationPolicyContractAssertionMappings gets a reference to the given []AuthenticationPolicyContractAssertionMapping and assigns it to the AuthenticationPolicyContractAssertionMappings field.
func (o *SpBrowserSso) SetAuthenticationPolicyContractAssertionMappings(v []AuthenticationPolicyContractAssertionMapping) {
	o.AuthenticationPolicyContractAssertionMappings = v
}

// GetAssertionLifetime returns the AssertionLifetime field value
func (o *SpBrowserSso) GetAssertionLifetime() AssertionLifetime {
	if o == nil {
		var ret AssertionLifetime
		return ret
	}

	return o.AssertionLifetime
}

// GetAssertionLifetimeOk returns a tuple with the AssertionLifetime field value
// and a boolean to check if the value has been set.
func (o *SpBrowserSso) GetAssertionLifetimeOk() (*AssertionLifetime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssertionLifetime, true
}

// SetAssertionLifetime sets field value
func (o *SpBrowserSso) SetAssertionLifetime(v AssertionLifetime) {
	o.AssertionLifetime = v
}

func (o SpBrowserSso) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpBrowserSso) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["protocol"] = o.Protocol
	if !IsNil(o.WsFedTokenType) {
		toSerialize["wsFedTokenType"] = o.WsFedTokenType
	}
	if !IsNil(o.WsTrustVersion) {
		toSerialize["wsTrustVersion"] = o.WsTrustVersion
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
	if !IsNil(o.DefaultTargetUrl) {
		toSerialize["defaultTargetUrl"] = o.DefaultTargetUrl
	}
	if !IsNil(o.AlwaysSignArtifactResponse) {
		toSerialize["alwaysSignArtifactResponse"] = o.AlwaysSignArtifactResponse
	}
	toSerialize["ssoServiceEndpoints"] = o.SsoServiceEndpoints
	if !IsNil(o.SpSamlIdentityMapping) {
		toSerialize["spSamlIdentityMapping"] = o.SpSamlIdentityMapping
	}
	if !IsNil(o.SpWsFedIdentityMapping) {
		toSerialize["spWsFedIdentityMapping"] = o.SpWsFedIdentityMapping
	}
	if !IsNil(o.SignResponseAsRequired) {
		toSerialize["signResponseAsRequired"] = o.SignResponseAsRequired
	}
	if !IsNil(o.SignAssertions) {
		toSerialize["signAssertions"] = o.SignAssertions
	}
	if !IsNil(o.RequireSignedAuthnRequests) {
		toSerialize["requireSignedAuthnRequests"] = o.RequireSignedAuthnRequests
	}
	if !IsNil(o.EncryptionPolicy) {
		toSerialize["encryptionPolicy"] = o.EncryptionPolicy
	}
	toSerialize["attributeContract"] = o.AttributeContract
	toSerialize["adapterMappings"] = o.AdapterMappings
	if !IsNil(o.AuthenticationPolicyContractAssertionMappings) {
		toSerialize["authenticationPolicyContractAssertionMappings"] = o.AuthenticationPolicyContractAssertionMappings
	}
	toSerialize["assertionLifetime"] = o.AssertionLifetime
	return toSerialize, nil
}

type NullableSpBrowserSso struct {
	value *SpBrowserSso
	isSet bool
}

func (v NullableSpBrowserSso) Get() *SpBrowserSso {
	return v.value
}

func (v *NullableSpBrowserSso) Set(val *SpBrowserSso) {
	v.value = val
	v.isSet = true
}

func (v NullableSpBrowserSso) IsSet() bool {
	return v.isSet
}

func (v *NullableSpBrowserSso) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpBrowserSso(val *SpBrowserSso) *NullableSpBrowserSso {
	return &NullableSpBrowserSso{value: val, isSet: true}
}

func (v NullableSpBrowserSso) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpBrowserSso) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
