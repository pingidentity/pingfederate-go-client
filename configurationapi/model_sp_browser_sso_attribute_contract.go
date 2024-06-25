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

// checks if the SpBrowserSsoAttributeContract type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpBrowserSsoAttributeContract{}

// SpBrowserSsoAttributeContract A set of user attributes that the IdP sends in the SAML assertion.
type SpBrowserSsoAttributeContract struct {
	// A list of read-only assertion attributes (for example, SAML_SUBJECT) that are automatically populated by PingFederate.
	CoreAttributes []SpBrowserSsoAttribute `json:"coreAttributes,omitempty" tfsdk:"core_attributes"`
	// A list of additional attributes that are added to the outgoing assertion.
	ExtendedAttributes []SpBrowserSsoAttribute `json:"extendedAttributes,omitempty" tfsdk:"extended_attributes"`
}

// NewSpBrowserSsoAttributeContract instantiates a new SpBrowserSsoAttributeContract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpBrowserSsoAttributeContract() *SpBrowserSsoAttributeContract {
	this := SpBrowserSsoAttributeContract{}
	return &this
}

// NewSpBrowserSsoAttributeContractWithDefaults instantiates a new SpBrowserSsoAttributeContract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpBrowserSsoAttributeContractWithDefaults() *SpBrowserSsoAttributeContract {
	this := SpBrowserSsoAttributeContract{}
	return &this
}

// GetCoreAttributes returns the CoreAttributes field value if set, zero value otherwise.
func (o *SpBrowserSsoAttributeContract) GetCoreAttributes() []SpBrowserSsoAttribute {
	if o == nil || IsNil(o.CoreAttributes) {
		var ret []SpBrowserSsoAttribute
		return ret
	}
	return o.CoreAttributes
}

// GetCoreAttributesOk returns a tuple with the CoreAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpBrowserSsoAttributeContract) GetCoreAttributesOk() ([]SpBrowserSsoAttribute, bool) {
	if o == nil || IsNil(o.CoreAttributes) {
		return nil, false
	}
	return o.CoreAttributes, true
}

// HasCoreAttributes returns a boolean if a field has been set.
func (o *SpBrowserSsoAttributeContract) HasCoreAttributes() bool {
	if o != nil && !IsNil(o.CoreAttributes) {
		return true
	}

	return false
}

// SetCoreAttributes gets a reference to the given []SpBrowserSsoAttribute and assigns it to the CoreAttributes field.
func (o *SpBrowserSsoAttributeContract) SetCoreAttributes(v []SpBrowserSsoAttribute) {
	o.CoreAttributes = v
}

// GetExtendedAttributes returns the ExtendedAttributes field value if set, zero value otherwise.
func (o *SpBrowserSsoAttributeContract) GetExtendedAttributes() []SpBrowserSsoAttribute {
	if o == nil || IsNil(o.ExtendedAttributes) {
		var ret []SpBrowserSsoAttribute
		return ret
	}
	return o.ExtendedAttributes
}

// GetExtendedAttributesOk returns a tuple with the ExtendedAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpBrowserSsoAttributeContract) GetExtendedAttributesOk() ([]SpBrowserSsoAttribute, bool) {
	if o == nil || IsNil(o.ExtendedAttributes) {
		return nil, false
	}
	return o.ExtendedAttributes, true
}

// HasExtendedAttributes returns a boolean if a field has been set.
func (o *SpBrowserSsoAttributeContract) HasExtendedAttributes() bool {
	if o != nil && !IsNil(o.ExtendedAttributes) {
		return true
	}

	return false
}

// SetExtendedAttributes gets a reference to the given []SpBrowserSsoAttribute and assigns it to the ExtendedAttributes field.
func (o *SpBrowserSsoAttributeContract) SetExtendedAttributes(v []SpBrowserSsoAttribute) {
	o.ExtendedAttributes = v
}

func (o SpBrowserSsoAttributeContract) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpBrowserSsoAttributeContract) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CoreAttributes) {
		toSerialize["coreAttributes"] = o.CoreAttributes
	}
	if !IsNil(o.ExtendedAttributes) {
		toSerialize["extendedAttributes"] = o.ExtendedAttributes
	}
	return toSerialize, nil
}

type NullableSpBrowserSsoAttributeContract struct {
	value *SpBrowserSsoAttributeContract
	isSet bool
}

func (v NullableSpBrowserSsoAttributeContract) Get() *SpBrowserSsoAttributeContract {
	return v.value
}

func (v *NullableSpBrowserSsoAttributeContract) Set(val *SpBrowserSsoAttributeContract) {
	v.value = val
	v.isSet = true
}

func (v NullableSpBrowserSsoAttributeContract) IsSet() bool {
	return v.isSet
}

func (v *NullableSpBrowserSsoAttributeContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpBrowserSsoAttributeContract(val *SpBrowserSsoAttributeContract) *NullableSpBrowserSsoAttributeContract {
	return &NullableSpBrowserSsoAttributeContract{value: val, isSet: true}
}

func (v NullableSpBrowserSsoAttributeContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpBrowserSsoAttributeContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
