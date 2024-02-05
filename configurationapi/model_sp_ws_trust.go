/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the SpWsTrust type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpWsTrust{}

// SpWsTrust Ws-Trust STS provides security-token validation and creation to extend SSO access to identity-enabled Web Services
type SpWsTrust struct {
	// The partner service identifiers.
	PartnerServiceIds []string `json:"partnerServiceIds" tfsdk:"partner_service_ids"`
	// When selected, four additional token-type requests become available.
	OAuthAssertionProfiles *bool `json:"oAuthAssertionProfiles,omitempty" tfsdk:"oauth_assertion_profiles"`
	// The default token type when a web service client (WSC) does not specify in the token request which token type the STS should issue. Defaults to SAML 2.0.
	DefaultTokenType *string `json:"defaultTokenType,omitempty" tfsdk:"default_token_type"`
	// When selected, the STS generates a symmetric key to be used in conjunction with the \"Holder of Key\" (HoK) designation for the assertion's Subject Confirmation Method.  This option does not apply to OAuth assertion profiles.
	GenerateKey *bool `json:"generateKey,omitempty" tfsdk:"generate_key"`
	// When selected, the STS encrypts the SAML 2.0 assertion. Applicable only to SAML 2.0 security token.  This option does not apply to OAuth assertion profiles.
	EncryptSaml2Assertion *bool `json:"encryptSaml2Assertion,omitempty" tfsdk:"encrypt_saml2_assertion"`
	// The amount of time before the SAML token was issued during which it is to be considered valid. The default value is 5.
	MinutesBefore *int64 `json:"minutesBefore,omitempty" tfsdk:"minutes_before"`
	// The amount of time after the SAML token was issued during which it is to be considered valid. The default value is 30.
	MinutesAfter      *int64                     `json:"minutesAfter,omitempty" tfsdk:"minutes_after"`
	AttributeContract SpWsTrustAttributeContract `json:"attributeContract" tfsdk:"attribute_contract"`
	// A list of token processors to validate incoming tokens.
	TokenProcessorMappings []IdpTokenProcessorMapping `json:"tokenProcessorMappings" tfsdk:"token_processor_mappings"`
	// If the attribute contract cannot be fulfilled using data from the Request, abort the transaction.
	AbortIfNotFulfilledFromRequest *bool         `json:"abortIfNotFulfilledFromRequest,omitempty" tfsdk:"abort_if_not_fulfilled_from_request"`
	RequestContractRef             *ResourceLink `json:"requestContractRef,omitempty" tfsdk:"request_contract_ref"`
	// The message customizations for WS-Trust. Depending on server settings, connection type, and protocol this may or may not be supported.
	MessageCustomizations []ProtocolMessageCustomization `json:"messageCustomizations,omitempty" tfsdk:"message_customizations"`
}

// NewSpWsTrust instantiates a new SpWsTrust object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpWsTrust(partnerServiceIds []string, attributeContract SpWsTrustAttributeContract, tokenProcessorMappings []IdpTokenProcessorMapping) *SpWsTrust {
	this := SpWsTrust{}
	this.PartnerServiceIds = partnerServiceIds
	this.AttributeContract = attributeContract
	this.TokenProcessorMappings = tokenProcessorMappings
	return &this
}

// NewSpWsTrustWithDefaults instantiates a new SpWsTrust object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpWsTrustWithDefaults() *SpWsTrust {
	this := SpWsTrust{}
	return &this
}

// GetPartnerServiceIds returns the PartnerServiceIds field value
func (o *SpWsTrust) GetPartnerServiceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PartnerServiceIds
}

// GetPartnerServiceIdsOk returns a tuple with the PartnerServiceIds field value
// and a boolean to check if the value has been set.
func (o *SpWsTrust) GetPartnerServiceIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PartnerServiceIds, true
}

// SetPartnerServiceIds sets field value
func (o *SpWsTrust) SetPartnerServiceIds(v []string) {
	o.PartnerServiceIds = v
}

// GetOAuthAssertionProfiles returns the OAuthAssertionProfiles field value if set, zero value otherwise.
func (o *SpWsTrust) GetOAuthAssertionProfiles() bool {
	if o == nil || IsNil(o.OAuthAssertionProfiles) {
		var ret bool
		return ret
	}
	return *o.OAuthAssertionProfiles
}

// GetOAuthAssertionProfilesOk returns a tuple with the OAuthAssertionProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpWsTrust) GetOAuthAssertionProfilesOk() (*bool, bool) {
	if o == nil || IsNil(o.OAuthAssertionProfiles) {
		return nil, false
	}
	return o.OAuthAssertionProfiles, true
}

// HasOAuthAssertionProfiles returns a boolean if a field has been set.
func (o *SpWsTrust) HasOAuthAssertionProfiles() bool {
	if o != nil && !IsNil(o.OAuthAssertionProfiles) {
		return true
	}

	return false
}

// SetOAuthAssertionProfiles gets a reference to the given bool and assigns it to the OAuthAssertionProfiles field.
func (o *SpWsTrust) SetOAuthAssertionProfiles(v bool) {
	o.OAuthAssertionProfiles = &v
}

// GetDefaultTokenType returns the DefaultTokenType field value if set, zero value otherwise.
func (o *SpWsTrust) GetDefaultTokenType() string {
	if o == nil || IsNil(o.DefaultTokenType) {
		var ret string
		return ret
	}
	return *o.DefaultTokenType
}

// GetDefaultTokenTypeOk returns a tuple with the DefaultTokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpWsTrust) GetDefaultTokenTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultTokenType) {
		return nil, false
	}
	return o.DefaultTokenType, true
}

// HasDefaultTokenType returns a boolean if a field has been set.
func (o *SpWsTrust) HasDefaultTokenType() bool {
	if o != nil && !IsNil(o.DefaultTokenType) {
		return true
	}

	return false
}

// SetDefaultTokenType gets a reference to the given string and assigns it to the DefaultTokenType field.
func (o *SpWsTrust) SetDefaultTokenType(v string) {
	o.DefaultTokenType = &v
}

// GetGenerateKey returns the GenerateKey field value if set, zero value otherwise.
func (o *SpWsTrust) GetGenerateKey() bool {
	if o == nil || IsNil(o.GenerateKey) {
		var ret bool
		return ret
	}
	return *o.GenerateKey
}

// GetGenerateKeyOk returns a tuple with the GenerateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpWsTrust) GetGenerateKeyOk() (*bool, bool) {
	if o == nil || IsNil(o.GenerateKey) {
		return nil, false
	}
	return o.GenerateKey, true
}

// HasGenerateKey returns a boolean if a field has been set.
func (o *SpWsTrust) HasGenerateKey() bool {
	if o != nil && !IsNil(o.GenerateKey) {
		return true
	}

	return false
}

// SetGenerateKey gets a reference to the given bool and assigns it to the GenerateKey field.
func (o *SpWsTrust) SetGenerateKey(v bool) {
	o.GenerateKey = &v
}

// GetEncryptSaml2Assertion returns the EncryptSaml2Assertion field value if set, zero value otherwise.
func (o *SpWsTrust) GetEncryptSaml2Assertion() bool {
	if o == nil || IsNil(o.EncryptSaml2Assertion) {
		var ret bool
		return ret
	}
	return *o.EncryptSaml2Assertion
}

// GetEncryptSaml2AssertionOk returns a tuple with the EncryptSaml2Assertion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpWsTrust) GetEncryptSaml2AssertionOk() (*bool, bool) {
	if o == nil || IsNil(o.EncryptSaml2Assertion) {
		return nil, false
	}
	return o.EncryptSaml2Assertion, true
}

// HasEncryptSaml2Assertion returns a boolean if a field has been set.
func (o *SpWsTrust) HasEncryptSaml2Assertion() bool {
	if o != nil && !IsNil(o.EncryptSaml2Assertion) {
		return true
	}

	return false
}

// SetEncryptSaml2Assertion gets a reference to the given bool and assigns it to the EncryptSaml2Assertion field.
func (o *SpWsTrust) SetEncryptSaml2Assertion(v bool) {
	o.EncryptSaml2Assertion = &v
}

// GetMinutesBefore returns the MinutesBefore field value if set, zero value otherwise.
func (o *SpWsTrust) GetMinutesBefore() int64 {
	if o == nil || IsNil(o.MinutesBefore) {
		var ret int64
		return ret
	}
	return *o.MinutesBefore
}

// GetMinutesBeforeOk returns a tuple with the MinutesBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpWsTrust) GetMinutesBeforeOk() (*int64, bool) {
	if o == nil || IsNil(o.MinutesBefore) {
		return nil, false
	}
	return o.MinutesBefore, true
}

// HasMinutesBefore returns a boolean if a field has been set.
func (o *SpWsTrust) HasMinutesBefore() bool {
	if o != nil && !IsNil(o.MinutesBefore) {
		return true
	}

	return false
}

// SetMinutesBefore gets a reference to the given int64 and assigns it to the MinutesBefore field.
func (o *SpWsTrust) SetMinutesBefore(v int64) {
	o.MinutesBefore = &v
}

// GetMinutesAfter returns the MinutesAfter field value if set, zero value otherwise.
func (o *SpWsTrust) GetMinutesAfter() int64 {
	if o == nil || IsNil(o.MinutesAfter) {
		var ret int64
		return ret
	}
	return *o.MinutesAfter
}

// GetMinutesAfterOk returns a tuple with the MinutesAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpWsTrust) GetMinutesAfterOk() (*int64, bool) {
	if o == nil || IsNil(o.MinutesAfter) {
		return nil, false
	}
	return o.MinutesAfter, true
}

// HasMinutesAfter returns a boolean if a field has been set.
func (o *SpWsTrust) HasMinutesAfter() bool {
	if o != nil && !IsNil(o.MinutesAfter) {
		return true
	}

	return false
}

// SetMinutesAfter gets a reference to the given int64 and assigns it to the MinutesAfter field.
func (o *SpWsTrust) SetMinutesAfter(v int64) {
	o.MinutesAfter = &v
}

// GetAttributeContract returns the AttributeContract field value
func (o *SpWsTrust) GetAttributeContract() SpWsTrustAttributeContract {
	if o == nil {
		var ret SpWsTrustAttributeContract
		return ret
	}

	return o.AttributeContract
}

// GetAttributeContractOk returns a tuple with the AttributeContract field value
// and a boolean to check if the value has been set.
func (o *SpWsTrust) GetAttributeContractOk() (*SpWsTrustAttributeContract, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeContract, true
}

// SetAttributeContract sets field value
func (o *SpWsTrust) SetAttributeContract(v SpWsTrustAttributeContract) {
	o.AttributeContract = v
}

// GetTokenProcessorMappings returns the TokenProcessorMappings field value
func (o *SpWsTrust) GetTokenProcessorMappings() []IdpTokenProcessorMapping {
	if o == nil {
		var ret []IdpTokenProcessorMapping
		return ret
	}

	return o.TokenProcessorMappings
}

// GetTokenProcessorMappingsOk returns a tuple with the TokenProcessorMappings field value
// and a boolean to check if the value has been set.
func (o *SpWsTrust) GetTokenProcessorMappingsOk() ([]IdpTokenProcessorMapping, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenProcessorMappings, true
}

// SetTokenProcessorMappings sets field value
func (o *SpWsTrust) SetTokenProcessorMappings(v []IdpTokenProcessorMapping) {
	o.TokenProcessorMappings = v
}

// GetAbortIfNotFulfilledFromRequest returns the AbortIfNotFulfilledFromRequest field value if set, zero value otherwise.
func (o *SpWsTrust) GetAbortIfNotFulfilledFromRequest() bool {
	if o == nil || IsNil(o.AbortIfNotFulfilledFromRequest) {
		var ret bool
		return ret
	}
	return *o.AbortIfNotFulfilledFromRequest
}

// GetAbortIfNotFulfilledFromRequestOk returns a tuple with the AbortIfNotFulfilledFromRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpWsTrust) GetAbortIfNotFulfilledFromRequestOk() (*bool, bool) {
	if o == nil || IsNil(o.AbortIfNotFulfilledFromRequest) {
		return nil, false
	}
	return o.AbortIfNotFulfilledFromRequest, true
}

// HasAbortIfNotFulfilledFromRequest returns a boolean if a field has been set.
func (o *SpWsTrust) HasAbortIfNotFulfilledFromRequest() bool {
	if o != nil && !IsNil(o.AbortIfNotFulfilledFromRequest) {
		return true
	}

	return false
}

// SetAbortIfNotFulfilledFromRequest gets a reference to the given bool and assigns it to the AbortIfNotFulfilledFromRequest field.
func (o *SpWsTrust) SetAbortIfNotFulfilledFromRequest(v bool) {
	o.AbortIfNotFulfilledFromRequest = &v
}

// GetRequestContractRef returns the RequestContractRef field value if set, zero value otherwise.
func (o *SpWsTrust) GetRequestContractRef() ResourceLink {
	if o == nil || IsNil(o.RequestContractRef) {
		var ret ResourceLink
		return ret
	}
	return *o.RequestContractRef
}

// GetRequestContractRefOk returns a tuple with the RequestContractRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpWsTrust) GetRequestContractRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.RequestContractRef) {
		return nil, false
	}
	return o.RequestContractRef, true
}

// HasRequestContractRef returns a boolean if a field has been set.
func (o *SpWsTrust) HasRequestContractRef() bool {
	if o != nil && !IsNil(o.RequestContractRef) {
		return true
	}

	return false
}

// SetRequestContractRef gets a reference to the given ResourceLink and assigns it to the RequestContractRef field.
func (o *SpWsTrust) SetRequestContractRef(v ResourceLink) {
	o.RequestContractRef = &v
}

// GetMessageCustomizations returns the MessageCustomizations field value if set, zero value otherwise.
func (o *SpWsTrust) GetMessageCustomizations() []ProtocolMessageCustomization {
	if o == nil || IsNil(o.MessageCustomizations) {
		var ret []ProtocolMessageCustomization
		return ret
	}
	return o.MessageCustomizations
}

// GetMessageCustomizationsOk returns a tuple with the MessageCustomizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpWsTrust) GetMessageCustomizationsOk() ([]ProtocolMessageCustomization, bool) {
	if o == nil || IsNil(o.MessageCustomizations) {
		return nil, false
	}
	return o.MessageCustomizations, true
}

// HasMessageCustomizations returns a boolean if a field has been set.
func (o *SpWsTrust) HasMessageCustomizations() bool {
	if o != nil && !IsNil(o.MessageCustomizations) {
		return true
	}

	return false
}

// SetMessageCustomizations gets a reference to the given []ProtocolMessageCustomization and assigns it to the MessageCustomizations field.
func (o *SpWsTrust) SetMessageCustomizations(v []ProtocolMessageCustomization) {
	o.MessageCustomizations = v
}

func (o SpWsTrust) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpWsTrust) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["partnerServiceIds"] = o.PartnerServiceIds
	if !IsNil(o.OAuthAssertionProfiles) {
		toSerialize["oAuthAssertionProfiles"] = o.OAuthAssertionProfiles
	}
	if !IsNil(o.DefaultTokenType) {
		toSerialize["defaultTokenType"] = o.DefaultTokenType
	}
	if !IsNil(o.GenerateKey) {
		toSerialize["generateKey"] = o.GenerateKey
	}
	if !IsNil(o.EncryptSaml2Assertion) {
		toSerialize["encryptSaml2Assertion"] = o.EncryptSaml2Assertion
	}
	if !IsNil(o.MinutesBefore) {
		toSerialize["minutesBefore"] = o.MinutesBefore
	}
	if !IsNil(o.MinutesAfter) {
		toSerialize["minutesAfter"] = o.MinutesAfter
	}
	toSerialize["attributeContract"] = o.AttributeContract
	toSerialize["tokenProcessorMappings"] = o.TokenProcessorMappings
	if !IsNil(o.AbortIfNotFulfilledFromRequest) {
		toSerialize["abortIfNotFulfilledFromRequest"] = o.AbortIfNotFulfilledFromRequest
	}
	if !IsNil(o.RequestContractRef) {
		toSerialize["requestContractRef"] = o.RequestContractRef
	}
	if !IsNil(o.MessageCustomizations) {
		toSerialize["messageCustomizations"] = o.MessageCustomizations
	}
	return toSerialize, nil
}

type NullableSpWsTrust struct {
	value *SpWsTrust
	isSet bool
}

func (v NullableSpWsTrust) Get() *SpWsTrust {
	return v.value
}

func (v *NullableSpWsTrust) Set(val *SpWsTrust) {
	v.value = val
	v.isSet = true
}

func (v NullableSpWsTrust) IsSet() bool {
	return v.isSet
}

func (v *NullableSpWsTrust) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpWsTrust(val *SpWsTrust) *NullableSpWsTrust {
	return &NullableSpWsTrust{value: val, isSet: true}
}

func (v NullableSpWsTrust) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpWsTrust) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
