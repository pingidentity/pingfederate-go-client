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

// checks if the IdentityHintContract type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityHintContract{}

// IdentityHintContract A set of attributes exposed by request policy contract.
type IdentityHintContract struct {
	// A list of required identity hint contract attributes.
	CoreAttributes []IdentityHintAttribute `json:"coreAttributes" tfsdk:"core_attributes"`
	// A list of additional identity hint contract attributes.
	ExtendedAttributes []IdentityHintAttribute `json:"extendedAttributes,omitempty" tfsdk:"extended_attributes"`
}

// NewIdentityHintContract instantiates a new IdentityHintContract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityHintContract(coreAttributes []IdentityHintAttribute) *IdentityHintContract {
	this := IdentityHintContract{}
	this.CoreAttributes = coreAttributes
	return &this
}

// NewIdentityHintContractWithDefaults instantiates a new IdentityHintContract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityHintContractWithDefaults() *IdentityHintContract {
	this := IdentityHintContract{}
	return &this
}

// GetCoreAttributes returns the CoreAttributes field value
func (o *IdentityHintContract) GetCoreAttributes() []IdentityHintAttribute {
	if o == nil {
		var ret []IdentityHintAttribute
		return ret
	}

	return o.CoreAttributes
}

// GetCoreAttributesOk returns a tuple with the CoreAttributes field value
// and a boolean to check if the value has been set.
func (o *IdentityHintContract) GetCoreAttributesOk() ([]IdentityHintAttribute, bool) {
	if o == nil {
		return nil, false
	}
	return o.CoreAttributes, true
}

// SetCoreAttributes sets field value
func (o *IdentityHintContract) SetCoreAttributes(v []IdentityHintAttribute) {
	o.CoreAttributes = v
}

// GetExtendedAttributes returns the ExtendedAttributes field value if set, zero value otherwise.
func (o *IdentityHintContract) GetExtendedAttributes() []IdentityHintAttribute {
	if o == nil || IsNil(o.ExtendedAttributes) {
		var ret []IdentityHintAttribute
		return ret
	}
	return o.ExtendedAttributes
}

// GetExtendedAttributesOk returns a tuple with the ExtendedAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityHintContract) GetExtendedAttributesOk() ([]IdentityHintAttribute, bool) {
	if o == nil || IsNil(o.ExtendedAttributes) {
		return nil, false
	}
	return o.ExtendedAttributes, true
}

// HasExtendedAttributes returns a boolean if a field has been set.
func (o *IdentityHintContract) HasExtendedAttributes() bool {
	if o != nil && !IsNil(o.ExtendedAttributes) {
		return true
	}

	return false
}

// SetExtendedAttributes gets a reference to the given []IdentityHintAttribute and assigns it to the ExtendedAttributes field.
func (o *IdentityHintContract) SetExtendedAttributes(v []IdentityHintAttribute) {
	o.ExtendedAttributes = v
}

func (o IdentityHintContract) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityHintContract) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["coreAttributes"] = o.CoreAttributes
	if !IsNil(o.ExtendedAttributes) {
		toSerialize["extendedAttributes"] = o.ExtendedAttributes
	}
	return toSerialize, nil
}

type NullableIdentityHintContract struct {
	value *IdentityHintContract
	isSet bool
}

func (v NullableIdentityHintContract) Get() *IdentityHintContract {
	return v.value
}

func (v *NullableIdentityHintContract) Set(val *IdentityHintContract) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityHintContract) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityHintContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityHintContract(val *IdentityHintContract) *NullableIdentityHintContract {
	return &NullableIdentityHintContract{value: val, isSet: true}
}

func (v NullableIdentityHintContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityHintContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
