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

// checks if the IdpOAuthAttributeContract type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdpOAuthAttributeContract{}

// IdpOAuthAttributeContract A set of user attributes that the IdP sends in the OAuth Assertion Grant.
type IdpOAuthAttributeContract struct {
	// A list of read-only assertion attributes that are automatically populated by PingFederate.
	CoreAttributes []IdpBrowserSsoAttribute `json:"coreAttributes,omitempty" tfsdk:"core_attributes"`
	// A list of additional attributes that are present in the incoming assertion.
	ExtendedAttributes []IdpBrowserSsoAttribute `json:"extendedAttributes,omitempty" tfsdk:"extended_attributes"`
}

// NewIdpOAuthAttributeContract instantiates a new IdpOAuthAttributeContract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpOAuthAttributeContract() *IdpOAuthAttributeContract {
	this := IdpOAuthAttributeContract{}
	return &this
}

// NewIdpOAuthAttributeContractWithDefaults instantiates a new IdpOAuthAttributeContract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpOAuthAttributeContractWithDefaults() *IdpOAuthAttributeContract {
	this := IdpOAuthAttributeContract{}
	return &this
}

// GetCoreAttributes returns the CoreAttributes field value if set, zero value otherwise.
func (o *IdpOAuthAttributeContract) GetCoreAttributes() []IdpBrowserSsoAttribute {
	if o == nil || IsNil(o.CoreAttributes) {
		var ret []IdpBrowserSsoAttribute
		return ret
	}
	return o.CoreAttributes
}

// GetCoreAttributesOk returns a tuple with the CoreAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpOAuthAttributeContract) GetCoreAttributesOk() ([]IdpBrowserSsoAttribute, bool) {
	if o == nil || IsNil(o.CoreAttributes) {
		return nil, false
	}
	return o.CoreAttributes, true
}

// HasCoreAttributes returns a boolean if a field has been set.
func (o *IdpOAuthAttributeContract) HasCoreAttributes() bool {
	if o != nil && !IsNil(o.CoreAttributes) {
		return true
	}

	return false
}

// SetCoreAttributes gets a reference to the given []IdpBrowserSsoAttribute and assigns it to the CoreAttributes field.
func (o *IdpOAuthAttributeContract) SetCoreAttributes(v []IdpBrowserSsoAttribute) {
	o.CoreAttributes = v
}

// GetExtendedAttributes returns the ExtendedAttributes field value if set, zero value otherwise.
func (o *IdpOAuthAttributeContract) GetExtendedAttributes() []IdpBrowserSsoAttribute {
	if o == nil || IsNil(o.ExtendedAttributes) {
		var ret []IdpBrowserSsoAttribute
		return ret
	}
	return o.ExtendedAttributes
}

// GetExtendedAttributesOk returns a tuple with the ExtendedAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpOAuthAttributeContract) GetExtendedAttributesOk() ([]IdpBrowserSsoAttribute, bool) {
	if o == nil || IsNil(o.ExtendedAttributes) {
		return nil, false
	}
	return o.ExtendedAttributes, true
}

// HasExtendedAttributes returns a boolean if a field has been set.
func (o *IdpOAuthAttributeContract) HasExtendedAttributes() bool {
	if o != nil && !IsNil(o.ExtendedAttributes) {
		return true
	}

	return false
}

// SetExtendedAttributes gets a reference to the given []IdpBrowserSsoAttribute and assigns it to the ExtendedAttributes field.
func (o *IdpOAuthAttributeContract) SetExtendedAttributes(v []IdpBrowserSsoAttribute) {
	o.ExtendedAttributes = v
}

func (o IdpOAuthAttributeContract) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdpOAuthAttributeContract) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CoreAttributes) {
		toSerialize["coreAttributes"] = o.CoreAttributes
	}
	if !IsNil(o.ExtendedAttributes) {
		toSerialize["extendedAttributes"] = o.ExtendedAttributes
	}
	return toSerialize, nil
}

type NullableIdpOAuthAttributeContract struct {
	value *IdpOAuthAttributeContract
	isSet bool
}

func (v NullableIdpOAuthAttributeContract) Get() *IdpOAuthAttributeContract {
	return v.value
}

func (v *NullableIdpOAuthAttributeContract) Set(val *IdpOAuthAttributeContract) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpOAuthAttributeContract) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpOAuthAttributeContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpOAuthAttributeContract(val *IdpOAuthAttributeContract) *NullableIdpOAuthAttributeContract {
	return &NullableIdpOAuthAttributeContract{value: val, isSet: true}
}

func (v NullableIdpOAuthAttributeContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpOAuthAttributeContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}