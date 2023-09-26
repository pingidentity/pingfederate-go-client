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

// checks if the IdpWsTrustAttributeContract type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdpWsTrustAttributeContract{}

// IdpWsTrustAttributeContract A set of user attributes that this server will receive in the token.
type IdpWsTrustAttributeContract struct {
	// A list of read-only assertion attributes that are automatically populated by PingFederate.
	CoreAttributes []IdpWsTrustAttribute `json:"coreAttributes,omitempty" tfsdk:"core_attributes"`
	// A list of additional attributes that are receive in the incoming assertion.
	ExtendedAttributes []IdpWsTrustAttribute `json:"extendedAttributes,omitempty" tfsdk:"extended_attributes"`
}

// NewIdpWsTrustAttributeContract instantiates a new IdpWsTrustAttributeContract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpWsTrustAttributeContract() *IdpWsTrustAttributeContract {
	this := IdpWsTrustAttributeContract{}
	return &this
}

// NewIdpWsTrustAttributeContractWithDefaults instantiates a new IdpWsTrustAttributeContract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpWsTrustAttributeContractWithDefaults() *IdpWsTrustAttributeContract {
	this := IdpWsTrustAttributeContract{}
	return &this
}

// GetCoreAttributes returns the CoreAttributes field value if set, zero value otherwise.
func (o *IdpWsTrustAttributeContract) GetCoreAttributes() []IdpWsTrustAttribute {
	if o == nil || IsNil(o.CoreAttributes) {
		var ret []IdpWsTrustAttribute
		return ret
	}
	return o.CoreAttributes
}

// GetCoreAttributesOk returns a tuple with the CoreAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpWsTrustAttributeContract) GetCoreAttributesOk() ([]IdpWsTrustAttribute, bool) {
	if o == nil || IsNil(o.CoreAttributes) {
		return nil, false
	}
	return o.CoreAttributes, true
}

// HasCoreAttributes returns a boolean if a field has been set.
func (o *IdpWsTrustAttributeContract) HasCoreAttributes() bool {
	if o != nil && !IsNil(o.CoreAttributes) {
		return true
	}

	return false
}

// SetCoreAttributes gets a reference to the given []IdpWsTrustAttribute and assigns it to the CoreAttributes field.
func (o *IdpWsTrustAttributeContract) SetCoreAttributes(v []IdpWsTrustAttribute) {
	o.CoreAttributes = v
}

// GetExtendedAttributes returns the ExtendedAttributes field value if set, zero value otherwise.
func (o *IdpWsTrustAttributeContract) GetExtendedAttributes() []IdpWsTrustAttribute {
	if o == nil || IsNil(o.ExtendedAttributes) {
		var ret []IdpWsTrustAttribute
		return ret
	}
	return o.ExtendedAttributes
}

// GetExtendedAttributesOk returns a tuple with the ExtendedAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpWsTrustAttributeContract) GetExtendedAttributesOk() ([]IdpWsTrustAttribute, bool) {
	if o == nil || IsNil(o.ExtendedAttributes) {
		return nil, false
	}
	return o.ExtendedAttributes, true
}

// HasExtendedAttributes returns a boolean if a field has been set.
func (o *IdpWsTrustAttributeContract) HasExtendedAttributes() bool {
	if o != nil && !IsNil(o.ExtendedAttributes) {
		return true
	}

	return false
}

// SetExtendedAttributes gets a reference to the given []IdpWsTrustAttribute and assigns it to the ExtendedAttributes field.
func (o *IdpWsTrustAttributeContract) SetExtendedAttributes(v []IdpWsTrustAttribute) {
	o.ExtendedAttributes = v
}

func (o IdpWsTrustAttributeContract) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdpWsTrustAttributeContract) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CoreAttributes) {
		toSerialize["coreAttributes"] = o.CoreAttributes
	}
	if !IsNil(o.ExtendedAttributes) {
		toSerialize["extendedAttributes"] = o.ExtendedAttributes
	}
	return toSerialize, nil
}

type NullableIdpWsTrustAttributeContract struct {
	value *IdpWsTrustAttributeContract
	isSet bool
}

func (v NullableIdpWsTrustAttributeContract) Get() *IdpWsTrustAttributeContract {
	return v.value
}

func (v *NullableIdpWsTrustAttributeContract) Set(val *IdpWsTrustAttributeContract) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpWsTrustAttributeContract) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpWsTrustAttributeContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpWsTrustAttributeContract(val *IdpWsTrustAttributeContract) *NullableIdpWsTrustAttributeContract {
	return &NullableIdpWsTrustAttributeContract{value: val, isSet: true}
}

func (v NullableIdpWsTrustAttributeContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpWsTrustAttributeContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}