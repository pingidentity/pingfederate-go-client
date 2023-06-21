/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the IdpInboundProvisioningAttributeContract type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdpInboundProvisioningAttributeContract{}

// IdpInboundProvisioningAttributeContract A set of user attributes that the IdP sends in the SCIM response.
type IdpInboundProvisioningAttributeContract struct {
	// A list of read-only assertion attributes that are automatically populated by PingFederate.
	CoreAttributes []IdpInboundProvisioningAttribute `json:"coreAttributes,omitempty"`
	// A list of additional attributes that are added to the SCIM response.
	ExtendedAttributes []IdpInboundProvisioningAttribute `json:"extendedAttributes,omitempty"`
}

// NewIdpInboundProvisioningAttributeContract instantiates a new IdpInboundProvisioningAttributeContract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpInboundProvisioningAttributeContract() *IdpInboundProvisioningAttributeContract {
	this := IdpInboundProvisioningAttributeContract{}
	return &this
}

// NewIdpInboundProvisioningAttributeContractWithDefaults instantiates a new IdpInboundProvisioningAttributeContract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpInboundProvisioningAttributeContractWithDefaults() *IdpInboundProvisioningAttributeContract {
	this := IdpInboundProvisioningAttributeContract{}
	return &this
}

// GetCoreAttributes returns the CoreAttributes field value if set, zero value otherwise.
func (o *IdpInboundProvisioningAttributeContract) GetCoreAttributes() []IdpInboundProvisioningAttribute {
	if o == nil || IsNil(o.CoreAttributes) {
		var ret []IdpInboundProvisioningAttribute
		return ret
	}
	return o.CoreAttributes
}

// GetCoreAttributesOk returns a tuple with the CoreAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpInboundProvisioningAttributeContract) GetCoreAttributesOk() ([]IdpInboundProvisioningAttribute, bool) {
	if o == nil || IsNil(o.CoreAttributes) {
		return nil, false
	}
	return o.CoreAttributes, true
}

// HasCoreAttributes returns a boolean if a field has been set.
func (o *IdpInboundProvisioningAttributeContract) HasCoreAttributes() bool {
	if o != nil && !IsNil(o.CoreAttributes) {
		return true
	}

	return false
}

// SetCoreAttributes gets a reference to the given []IdpInboundProvisioningAttribute and assigns it to the CoreAttributes field.
func (o *IdpInboundProvisioningAttributeContract) SetCoreAttributes(v []IdpInboundProvisioningAttribute) {
	o.CoreAttributes = v
}

// GetExtendedAttributes returns the ExtendedAttributes field value if set, zero value otherwise.
func (o *IdpInboundProvisioningAttributeContract) GetExtendedAttributes() []IdpInboundProvisioningAttribute {
	if o == nil || IsNil(o.ExtendedAttributes) {
		var ret []IdpInboundProvisioningAttribute
		return ret
	}
	return o.ExtendedAttributes
}

// GetExtendedAttributesOk returns a tuple with the ExtendedAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpInboundProvisioningAttributeContract) GetExtendedAttributesOk() ([]IdpInboundProvisioningAttribute, bool) {
	if o == nil || IsNil(o.ExtendedAttributes) {
		return nil, false
	}
	return o.ExtendedAttributes, true
}

// HasExtendedAttributes returns a boolean if a field has been set.
func (o *IdpInboundProvisioningAttributeContract) HasExtendedAttributes() bool {
	if o != nil && !IsNil(o.ExtendedAttributes) {
		return true
	}

	return false
}

// SetExtendedAttributes gets a reference to the given []IdpInboundProvisioningAttribute and assigns it to the ExtendedAttributes field.
func (o *IdpInboundProvisioningAttributeContract) SetExtendedAttributes(v []IdpInboundProvisioningAttribute) {
	o.ExtendedAttributes = v
}

func (o IdpInboundProvisioningAttributeContract) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdpInboundProvisioningAttributeContract) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CoreAttributes) {
		toSerialize["coreAttributes"] = o.CoreAttributes
	}
	if !IsNil(o.ExtendedAttributes) {
		toSerialize["extendedAttributes"] = o.ExtendedAttributes
	}
	return toSerialize, nil
}

type NullableIdpInboundProvisioningAttributeContract struct {
	value *IdpInboundProvisioningAttributeContract
	isSet bool
}

func (v NullableIdpInboundProvisioningAttributeContract) Get() *IdpInboundProvisioningAttributeContract {
	return v.value
}

func (v *NullableIdpInboundProvisioningAttributeContract) Set(val *IdpInboundProvisioningAttributeContract) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpInboundProvisioningAttributeContract) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpInboundProvisioningAttributeContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpInboundProvisioningAttributeContract(val *IdpInboundProvisioningAttributeContract) *NullableIdpInboundProvisioningAttributeContract {
	return &NullableIdpInboundProvisioningAttributeContract{value: val, isSet: true}
}

func (v NullableIdpInboundProvisioningAttributeContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpInboundProvisioningAttributeContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
