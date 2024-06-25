/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the IdpInboundProvisioningAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdpInboundProvisioningAttribute{}

// IdpInboundProvisioningAttribute An attribute for the IdP Inbound Provisioning attribute contract.
type IdpInboundProvisioningAttribute struct {
	// The name of this attribute.
	Name string `json:"name" tfsdk:"name"`
	// Specifies whether this attribute is masked in PingFederate logs. Defaults to false.
	Masked *bool `json:"masked,omitempty" tfsdk:"masked"`
}

// NewIdpInboundProvisioningAttribute instantiates a new IdpInboundProvisioningAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpInboundProvisioningAttribute(name string) *IdpInboundProvisioningAttribute {
	this := IdpInboundProvisioningAttribute{}
	this.Name = name
	return &this
}

// NewIdpInboundProvisioningAttributeWithDefaults instantiates a new IdpInboundProvisioningAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpInboundProvisioningAttributeWithDefaults() *IdpInboundProvisioningAttribute {
	this := IdpInboundProvisioningAttribute{}
	return &this
}

// GetName returns the Name field value
func (o *IdpInboundProvisioningAttribute) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IdpInboundProvisioningAttribute) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IdpInboundProvisioningAttribute) SetName(v string) {
	o.Name = v
}

// GetMasked returns the Masked field value if set, zero value otherwise.
func (o *IdpInboundProvisioningAttribute) GetMasked() bool {
	if o == nil || IsNil(o.Masked) {
		var ret bool
		return ret
	}
	return *o.Masked
}

// GetMaskedOk returns a tuple with the Masked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpInboundProvisioningAttribute) GetMaskedOk() (*bool, bool) {
	if o == nil || IsNil(o.Masked) {
		return nil, false
	}
	return o.Masked, true
}

// HasMasked returns a boolean if a field has been set.
func (o *IdpInboundProvisioningAttribute) HasMasked() bool {
	if o != nil && !IsNil(o.Masked) {
		return true
	}

	return false
}

// SetMasked gets a reference to the given bool and assigns it to the Masked field.
func (o *IdpInboundProvisioningAttribute) SetMasked(v bool) {
	o.Masked = &v
}

func (o IdpInboundProvisioningAttribute) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdpInboundProvisioningAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Masked) {
		toSerialize["masked"] = o.Masked
	}
	return toSerialize, nil
}

type NullableIdpInboundProvisioningAttribute struct {
	value *IdpInboundProvisioningAttribute
	isSet bool
}

func (v NullableIdpInboundProvisioningAttribute) Get() *IdpInboundProvisioningAttribute {
	return v.value
}

func (v *NullableIdpInboundProvisioningAttribute) Set(val *IdpInboundProvisioningAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpInboundProvisioningAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpInboundProvisioningAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpInboundProvisioningAttribute(val *IdpInboundProvisioningAttribute) *NullableIdpInboundProvisioningAttribute {
	return &NullableIdpInboundProvisioningAttribute{value: val, isSet: true}
}

func (v NullableIdpInboundProvisioningAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpInboundProvisioningAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
