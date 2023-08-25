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

// checks if the RolesAndProtocolsIdpRoleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RolesAndProtocolsIdpRoleAllOf{}

// RolesAndProtocolsIdpRoleAllOf This property has been deprecated and is no longer used. All Roles and protocols are always enabled.
type RolesAndProtocolsIdpRoleAllOf struct {
	// Enable Identity Provider Role.
	Enable        interface{}    `json:"enable,omitempty" tfsdk:"enable"`
	Saml20Profile *SAML20Profile `json:"saml20Profile,omitempty" tfsdk:"saml20_profile"`
	// Enable Outbound Provisioning.
	EnableOutboundProvisioning interface{} `json:"enableOutboundProvisioning,omitempty" tfsdk:"enable_outbound_provisioning"`
}

// NewRolesAndProtocolsIdpRoleAllOf instantiates a new RolesAndProtocolsIdpRoleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRolesAndProtocolsIdpRoleAllOf() *RolesAndProtocolsIdpRoleAllOf {
	this := RolesAndProtocolsIdpRoleAllOf{}
	return &this
}

// NewRolesAndProtocolsIdpRoleAllOfWithDefaults instantiates a new RolesAndProtocolsIdpRoleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRolesAndProtocolsIdpRoleAllOfWithDefaults() *RolesAndProtocolsIdpRoleAllOf {
	this := RolesAndProtocolsIdpRoleAllOf{}
	return &this
}

// GetEnable returns the Enable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RolesAndProtocolsIdpRoleAllOf) GetEnable() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RolesAndProtocolsIdpRoleAllOf) GetEnableOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return &o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *RolesAndProtocolsIdpRoleAllOf) HasEnable() bool {
	if o != nil && IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given interface{} and assigns it to the Enable field.
func (o *RolesAndProtocolsIdpRoleAllOf) SetEnable(v interface{}) {
	o.Enable = v
}

// GetSaml20Profile returns the Saml20Profile field value if set, zero value otherwise.
func (o *RolesAndProtocolsIdpRoleAllOf) GetSaml20Profile() SAML20Profile {
	if o == nil || IsNil(o.Saml20Profile) {
		var ret SAML20Profile
		return ret
	}
	return *o.Saml20Profile
}

// GetSaml20ProfileOk returns a tuple with the Saml20Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolesAndProtocolsIdpRoleAllOf) GetSaml20ProfileOk() (*SAML20Profile, bool) {
	if o == nil || IsNil(o.Saml20Profile) {
		return nil, false
	}
	return o.Saml20Profile, true
}

// HasSaml20Profile returns a boolean if a field has been set.
func (o *RolesAndProtocolsIdpRoleAllOf) HasSaml20Profile() bool {
	if o != nil && !IsNil(o.Saml20Profile) {
		return true
	}

	return false
}

// SetSaml20Profile gets a reference to the given SAML20Profile and assigns it to the Saml20Profile field.
func (o *RolesAndProtocolsIdpRoleAllOf) SetSaml20Profile(v SAML20Profile) {
	o.Saml20Profile = &v
}

// GetEnableOutboundProvisioning returns the EnableOutboundProvisioning field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RolesAndProtocolsIdpRoleAllOf) GetEnableOutboundProvisioning() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EnableOutboundProvisioning
}

// GetEnableOutboundProvisioningOk returns a tuple with the EnableOutboundProvisioning field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RolesAndProtocolsIdpRoleAllOf) GetEnableOutboundProvisioningOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EnableOutboundProvisioning) {
		return nil, false
	}
	return &o.EnableOutboundProvisioning, true
}

// HasEnableOutboundProvisioning returns a boolean if a field has been set.
func (o *RolesAndProtocolsIdpRoleAllOf) HasEnableOutboundProvisioning() bool {
	if o != nil && IsNil(o.EnableOutboundProvisioning) {
		return true
	}

	return false
}

// SetEnableOutboundProvisioning gets a reference to the given interface{} and assigns it to the EnableOutboundProvisioning field.
func (o *RolesAndProtocolsIdpRoleAllOf) SetEnableOutboundProvisioning(v interface{}) {
	o.EnableOutboundProvisioning = v
}

func (o RolesAndProtocolsIdpRoleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RolesAndProtocolsIdpRoleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Enable != nil {
		toSerialize["enable"] = o.Enable
	}
	if !IsNil(o.Saml20Profile) {
		toSerialize["saml20Profile"] = o.Saml20Profile
	}
	if o.EnableOutboundProvisioning != nil {
		toSerialize["enableOutboundProvisioning"] = o.EnableOutboundProvisioning
	}
	return toSerialize, nil
}

type NullableRolesAndProtocolsIdpRoleAllOf struct {
	value *RolesAndProtocolsIdpRoleAllOf
	isSet bool
}

func (v NullableRolesAndProtocolsIdpRoleAllOf) Get() *RolesAndProtocolsIdpRoleAllOf {
	return v.value
}

func (v *NullableRolesAndProtocolsIdpRoleAllOf) Set(val *RolesAndProtocolsIdpRoleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRolesAndProtocolsIdpRoleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRolesAndProtocolsIdpRoleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRolesAndProtocolsIdpRoleAllOf(val *RolesAndProtocolsIdpRoleAllOf) *NullableRolesAndProtocolsIdpRoleAllOf {
	return &NullableRolesAndProtocolsIdpRoleAllOf{value: val, isSet: true}
}

func (v NullableRolesAndProtocolsIdpRoleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRolesAndProtocolsIdpRoleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
