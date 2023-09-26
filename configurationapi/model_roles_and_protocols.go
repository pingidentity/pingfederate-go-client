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

// checks if the RolesAndProtocols type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RolesAndProtocols{}

// RolesAndProtocols This property has been deprecated and is no longer used. All Roles and protocols are always enabled.
type RolesAndProtocols struct {
	OauthRole *OAuthRole `json:"oauthRole,omitempty" tfsdk:"oauth_role"`
	IdpRole   *IdpRole   `json:"idpRole,omitempty" tfsdk:"idp_role"`
	SpRole    *SpRole    `json:"spRole,omitempty" tfsdk:"sp_role"`
	// Enable IdP Discovery.
	EnableIdpDiscovery *bool `json:"enableIdpDiscovery,omitempty" tfsdk:"enable_idp_discovery"`
}

// NewRolesAndProtocols instantiates a new RolesAndProtocols object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRolesAndProtocols() *RolesAndProtocols {
	this := RolesAndProtocols{}
	return &this
}

// NewRolesAndProtocolsWithDefaults instantiates a new RolesAndProtocols object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRolesAndProtocolsWithDefaults() *RolesAndProtocols {
	this := RolesAndProtocols{}
	return &this
}

// GetOauthRole returns the OauthRole field value if set, zero value otherwise.
func (o *RolesAndProtocols) GetOauthRole() OAuthRole {
	if o == nil || IsNil(o.OauthRole) {
		var ret OAuthRole
		return ret
	}
	return *o.OauthRole
}

// GetOauthRoleOk returns a tuple with the OauthRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolesAndProtocols) GetOauthRoleOk() (*OAuthRole, bool) {
	if o == nil || IsNil(o.OauthRole) {
		return nil, false
	}
	return o.OauthRole, true
}

// HasOauthRole returns a boolean if a field has been set.
func (o *RolesAndProtocols) HasOauthRole() bool {
	if o != nil && !IsNil(o.OauthRole) {
		return true
	}

	return false
}

// SetOauthRole gets a reference to the given OAuthRole and assigns it to the OauthRole field.
func (o *RolesAndProtocols) SetOauthRole(v OAuthRole) {
	o.OauthRole = &v
}

// GetIdpRole returns the IdpRole field value if set, zero value otherwise.
func (o *RolesAndProtocols) GetIdpRole() IdpRole {
	if o == nil || IsNil(o.IdpRole) {
		var ret IdpRole
		return ret
	}
	return *o.IdpRole
}

// GetIdpRoleOk returns a tuple with the IdpRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolesAndProtocols) GetIdpRoleOk() (*IdpRole, bool) {
	if o == nil || IsNil(o.IdpRole) {
		return nil, false
	}
	return o.IdpRole, true
}

// HasIdpRole returns a boolean if a field has been set.
func (o *RolesAndProtocols) HasIdpRole() bool {
	if o != nil && !IsNil(o.IdpRole) {
		return true
	}

	return false
}

// SetIdpRole gets a reference to the given IdpRole and assigns it to the IdpRole field.
func (o *RolesAndProtocols) SetIdpRole(v IdpRole) {
	o.IdpRole = &v
}

// GetSpRole returns the SpRole field value if set, zero value otherwise.
func (o *RolesAndProtocols) GetSpRole() SpRole {
	if o == nil || IsNil(o.SpRole) {
		var ret SpRole
		return ret
	}
	return *o.SpRole
}

// GetSpRoleOk returns a tuple with the SpRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolesAndProtocols) GetSpRoleOk() (*SpRole, bool) {
	if o == nil || IsNil(o.SpRole) {
		return nil, false
	}
	return o.SpRole, true
}

// HasSpRole returns a boolean if a field has been set.
func (o *RolesAndProtocols) HasSpRole() bool {
	if o != nil && !IsNil(o.SpRole) {
		return true
	}

	return false
}

// SetSpRole gets a reference to the given SpRole and assigns it to the SpRole field.
func (o *RolesAndProtocols) SetSpRole(v SpRole) {
	o.SpRole = &v
}

// GetEnableIdpDiscovery returns the EnableIdpDiscovery field value if set, zero value otherwise.
func (o *RolesAndProtocols) GetEnableIdpDiscovery() bool {
	if o == nil || IsNil(o.EnableIdpDiscovery) {
		var ret bool
		return ret
	}
	return *o.EnableIdpDiscovery
}

// GetEnableIdpDiscoveryOk returns a tuple with the EnableIdpDiscovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RolesAndProtocols) GetEnableIdpDiscoveryOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableIdpDiscovery) {
		return nil, false
	}
	return o.EnableIdpDiscovery, true
}

// HasEnableIdpDiscovery returns a boolean if a field has been set.
func (o *RolesAndProtocols) HasEnableIdpDiscovery() bool {
	if o != nil && !IsNil(o.EnableIdpDiscovery) {
		return true
	}

	return false
}

// SetEnableIdpDiscovery gets a reference to the given bool and assigns it to the EnableIdpDiscovery field.
func (o *RolesAndProtocols) SetEnableIdpDiscovery(v bool) {
	o.EnableIdpDiscovery = &v
}

func (o RolesAndProtocols) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RolesAndProtocols) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OauthRole) {
		toSerialize["oauthRole"] = o.OauthRole
	}
	if !IsNil(o.IdpRole) {
		toSerialize["idpRole"] = o.IdpRole
	}
	if !IsNil(o.SpRole) {
		toSerialize["spRole"] = o.SpRole
	}
	if !IsNil(o.EnableIdpDiscovery) {
		toSerialize["enableIdpDiscovery"] = o.EnableIdpDiscovery
	}
	return toSerialize, nil
}

type NullableRolesAndProtocols struct {
	value *RolesAndProtocols
	isSet bool
}

func (v NullableRolesAndProtocols) Get() *RolesAndProtocols {
	return v.value
}

func (v *NullableRolesAndProtocols) Set(val *RolesAndProtocols) {
	v.value = val
	v.isSet = true
}

func (v NullableRolesAndProtocols) IsSet() bool {
	return v.isSet
}

func (v *NullableRolesAndProtocols) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRolesAndProtocols(val *RolesAndProtocols) *NullableRolesAndProtocols {
	return &NullableRolesAndProtocols{value: val, isSet: true}
}

func (v NullableRolesAndProtocols) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRolesAndProtocols) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}