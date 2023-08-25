/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version:
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the RolesAndProtocolsSpRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RolesAndProtocolsSpRole{}

// RolesAndProtocolsSpRole struct for RolesAndProtocolsSpRole
type RolesAndProtocolsSpRole struct {
	// Enable Service Provider Role.
	Enable interface{} `json:"enable,omitempty" tfsdk:"enable"`
	// Enable SAML 1.1.
	EnableSaml11 interface{} `json:"enableSaml11,omitempty" tfsdk:"enable_saml11"`
	// Enable SAML 1.0.
	EnableSaml10 interface{} `json:"enableSaml10,omitempty" tfsdk:"enable_saml10"`
	// Enable WS Federation.
	EnableWsFed interface{} `json:"enableWsFed,omitempty" tfsdk:"enable_ws_fed"`
	// Enable WS Trust.
	EnableWsTrust interface{}      `json:"enableWsTrust,omitempty" tfsdk:"enable_ws_trust"`
	Saml20Profile *SpSAML20Profile `json:"saml20Profile,omitempty" tfsdk:"saml20_profile"`
	// Enable OpenID Connect.
	EnableOpenIDConnect interface{} `json:"enableOpenIDConnect,omitempty" tfsdk:"enable_open_idc_onnect"`
	// Enable Inbound Provisioning.
	EnableInboundProvisioning interface{} `json:"enableInboundProvisioning,omitempty" tfsdk:"enable_inbound_provisioning"`
}

// NewRolesAndProtocolsSpRole instantiates a new RolesAndProtocolsSpRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRolesAndProtocolsSpRole() *RolesAndProtocolsSpRole {
	this := RolesAndProtocolsSpRole{}
	return &this
}

// NewRolesAndProtocolsSpRoleWithDefaults instantiates a new RolesAndProtocolsSpRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRolesAndProtocolsSpRoleWithDefaults() *RolesAndProtocolsSpRole {
	this := RolesAndProtocolsSpRole{}
	return &this
}

// GetEnable returns the Enable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RolesAndProtocolsSpRole) GetEnable() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RolesAndProtocolsSpRole) GetEnableOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return &o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *RolesAndProtocolsSpRole) HasEnable() bool {
	if o != nil && IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given interface{} and assigns it to the Enable field.
func (o *RolesAndProtocolsSpRole) SetEnable(v interface{}) {
	o.Enable = v
}

// GetEnableSaml11 returns the EnableSaml11 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RolesAndProtocolsSpRole) GetEnableSaml11() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EnableSaml11
}

// GetEnableSaml11Ok returns a tuple with the EnableSaml11 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RolesAndProtocolsSpRole) GetEnableSaml11Ok() (*interface{}, bool) {
	if o == nil || IsNil(o.EnableSaml11) {
		return nil, false
	}
	return &o.EnableSaml11, true
}

// HasEnableSaml11 returns a boolean if a field has been set.
func (o *RolesAndProtocolsSpRole) HasEnableSaml11() bool {
	if o != nil && IsNil(o.EnableSaml11) {
		return true
	}

	return false
}

// SetEnableSaml11 gets a reference to the given interface{} and assigns it to the EnableSaml11 field.
func (o *RolesAndProtocolsSpRole) SetEnableSaml11(v interface{}) {
	o.EnableSaml11 = v
}

// GetEnableSaml10 returns the EnableSaml10 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RolesAndProtocolsSpRole) GetEnableSaml10() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EnableSaml10
}

// GetEnableSaml10Ok returns a tuple with the EnableSaml10 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RolesAndProtocolsSpRole) GetEnableSaml10Ok() (*interface{}, bool) {
	if o == nil || IsNil(o.EnableSaml10) {
		return nil, false
	}
	return &o.EnableSaml10, true
}

// HasEnableSaml10 returns a boolean if a field has been set.
func (o *RolesAndProtocolsSpRole) HasEnableSaml10() bool {
	if o != nil && IsNil(o.EnableSaml10) {
		return true
	}

	return false
}

// SetEnableSaml10 gets a reference to the given interface{} and assigns it to the EnableSaml10 field.
func (o *RolesAndProtocolsSpRole) SetEnableSaml10(v interface{}) {
	o.EnableSaml10 = v
}

// GetEnableWsFed returns the EnableWsFed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RolesAndProtocolsSpRole) GetEnableWsFed() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EnableWsFed
}

// GetEnableWsFedOk returns a tuple with the EnableWsFed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RolesAndProtocolsSpRole) GetEnableWsFedOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EnableWsFed) {
		return nil, false
	}
	return &o.EnableWsFed, true
}

// HasEnableWsFed returns a boolean if a field has been set.
func (o *RolesAndProtocolsSpRole) HasEnableWsFed() bool {
	if o != nil && IsNil(o.EnableWsFed) {
		return true
	}

	return false
}

// SetEnableWsFed gets a reference to the given interface{} and assigns it to the EnableWsFed field.
func (o *RolesAndProtocolsSpRole) SetEnableWsFed(v interface{}) {
	o.EnableWsFed = v
}

// GetEnableWsTrust returns the EnableWsTrust field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RolesAndProtocolsSpRole) GetEnableWsTrust() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EnableWsTrust
}

// GetEnableWsTrustOk returns a tuple with the EnableWsTrust field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RolesAndProtocolsSpRole) GetEnableWsTrustOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EnableWsTrust) {
		return nil, false
	}
	return &o.EnableWsTrust, true
}

// HasEnableWsTrust returns a boolean if a field has been set.
func (o *RolesAndProtocolsSpRole) HasEnableWsTrust() bool {
	if o != nil && IsNil(o.EnableWsTrust) {
		return true
	}

	return false
}

// SetEnableWsTrust gets a reference to the given interface{} and assigns it to the EnableWsTrust field.
func (o *RolesAndProtocolsSpRole) SetEnableWsTrust(v interface{}) {
	o.EnableWsTrust = v
}

// GetSaml20Profile returns the Saml20Profile field value if set, zero value otherwise.
func (o *RolesAndProtocolsSpRole) GetSaml20Profile() SpSAML20Profile {
	if o == nil || IsNil(o.Saml20Profile) {
		var ret SpSAML20Profile
		return ret
	}
	return *o.Saml20Profile
}

// GetSaml20ProfileOk returns a tuple with the Saml20Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolesAndProtocolsSpRole) GetSaml20ProfileOk() (*SpSAML20Profile, bool) {
	if o == nil || IsNil(o.Saml20Profile) {
		return nil, false
	}
	return o.Saml20Profile, true
}

// HasSaml20Profile returns a boolean if a field has been set.
func (o *RolesAndProtocolsSpRole) HasSaml20Profile() bool {
	if o != nil && !IsNil(o.Saml20Profile) {
		return true
	}

	return false
}

// SetSaml20Profile gets a reference to the given SpSAML20Profile and assigns it to the Saml20Profile field.
func (o *RolesAndProtocolsSpRole) SetSaml20Profile(v SpSAML20Profile) {
	o.Saml20Profile = &v
}

// GetEnableOpenIDConnect returns the EnableOpenIDConnect field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RolesAndProtocolsSpRole) GetEnableOpenIDConnect() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EnableOpenIDConnect
}

// GetEnableOpenIDConnectOk returns a tuple with the EnableOpenIDConnect field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RolesAndProtocolsSpRole) GetEnableOpenIDConnectOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EnableOpenIDConnect) {
		return nil, false
	}
	return &o.EnableOpenIDConnect, true
}

// HasEnableOpenIDConnect returns a boolean if a field has been set.
func (o *RolesAndProtocolsSpRole) HasEnableOpenIDConnect() bool {
	if o != nil && IsNil(o.EnableOpenIDConnect) {
		return true
	}

	return false
}

// SetEnableOpenIDConnect gets a reference to the given interface{} and assigns it to the EnableOpenIDConnect field.
func (o *RolesAndProtocolsSpRole) SetEnableOpenIDConnect(v interface{}) {
	o.EnableOpenIDConnect = v
}

// GetEnableInboundProvisioning returns the EnableInboundProvisioning field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RolesAndProtocolsSpRole) GetEnableInboundProvisioning() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EnableInboundProvisioning
}

// GetEnableInboundProvisioningOk returns a tuple with the EnableInboundProvisioning field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RolesAndProtocolsSpRole) GetEnableInboundProvisioningOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EnableInboundProvisioning) {
		return nil, false
	}
	return &o.EnableInboundProvisioning, true
}

// HasEnableInboundProvisioning returns a boolean if a field has been set.
func (o *RolesAndProtocolsSpRole) HasEnableInboundProvisioning() bool {
	if o != nil && IsNil(o.EnableInboundProvisioning) {
		return true
	}

	return false
}

// SetEnableInboundProvisioning gets a reference to the given interface{} and assigns it to the EnableInboundProvisioning field.
func (o *RolesAndProtocolsSpRole) SetEnableInboundProvisioning(v interface{}) {
	o.EnableInboundProvisioning = v
}

func (o RolesAndProtocolsSpRole) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RolesAndProtocolsSpRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Enable != nil {
		toSerialize["enable"] = o.Enable
	}
	if o.EnableSaml11 != nil {
		toSerialize["enableSaml11"] = o.EnableSaml11
	}
	if o.EnableSaml10 != nil {
		toSerialize["enableSaml10"] = o.EnableSaml10
	}
	if o.EnableWsFed != nil {
		toSerialize["enableWsFed"] = o.EnableWsFed
	}
	if o.EnableWsTrust != nil {
		toSerialize["enableWsTrust"] = o.EnableWsTrust
	}
	if !IsNil(o.Saml20Profile) {
		toSerialize["saml20Profile"] = o.Saml20Profile
	}
	if o.EnableOpenIDConnect != nil {
		toSerialize["enableOpenIDConnect"] = o.EnableOpenIDConnect
	}
	if o.EnableInboundProvisioning != nil {
		toSerialize["enableInboundProvisioning"] = o.EnableInboundProvisioning
	}
	return toSerialize, nil
}

type NullableRolesAndProtocolsSpRole struct {
	value *RolesAndProtocolsSpRole
	isSet bool
}

func (v NullableRolesAndProtocolsSpRole) Get() *RolesAndProtocolsSpRole {
	return v.value
}

func (v *NullableRolesAndProtocolsSpRole) Set(val *RolesAndProtocolsSpRole) {
	v.value = val
	v.isSet = true
}

func (v NullableRolesAndProtocolsSpRole) IsSet() bool {
	return v.isSet
}

func (v *NullableRolesAndProtocolsSpRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRolesAndProtocolsSpRole(val *RolesAndProtocolsSpRole) *NullableRolesAndProtocolsSpRole {
	return &NullableRolesAndProtocolsSpRole{value: val, isSet: true}
}

func (v NullableRolesAndProtocolsSpRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRolesAndProtocolsSpRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
