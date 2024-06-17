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

// checks if the IdpAdapterAttributeContract type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdpAdapterAttributeContract{}

// IdpAdapterAttributeContract A set of attributes exposed by an IdP adapter.
type IdpAdapterAttributeContract struct {
	// A list of IdP adapter attributes that correspond to the attributes exposed by the IdP adapter type.
	CoreAttributes []IdpAdapterAttribute `json:"coreAttributes" tfsdk:"core_attributes"`
	// A list of additional attributes that can be returned by the IdP adapter. The extended attributes are only used if the adapter supports them.
	ExtendedAttributes []IdpAdapterAttribute `json:"extendedAttributes,omitempty" tfsdk:"extended_attributes"`
	// The attribute to use for uniquely identify a user's authentication sessions.
	UniqueUserKeyAttribute *string `json:"uniqueUserKeyAttribute,omitempty" tfsdk:"unique_user_key_attribute"`
	// Whether or not all OGNL expressions used to fulfill an outgoing assertion contract should be masked in the logs. Defaults to false.
	MaskOgnlValues *bool `json:"maskOgnlValues,omitempty" tfsdk:"mask_ognl_values"`
}

// NewIdpAdapterAttributeContract instantiates a new IdpAdapterAttributeContract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpAdapterAttributeContract(coreAttributes []IdpAdapterAttribute) *IdpAdapterAttributeContract {
	this := IdpAdapterAttributeContract{}
	this.CoreAttributes = coreAttributes
	return &this
}

// NewIdpAdapterAttributeContractWithDefaults instantiates a new IdpAdapterAttributeContract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpAdapterAttributeContractWithDefaults() *IdpAdapterAttributeContract {
	this := IdpAdapterAttributeContract{}
	return &this
}

// GetCoreAttributes returns the CoreAttributes field value
func (o *IdpAdapterAttributeContract) GetCoreAttributes() []IdpAdapterAttribute {
	if o == nil {
		var ret []IdpAdapterAttribute
		return ret
	}

	return o.CoreAttributes
}

// GetCoreAttributesOk returns a tuple with the CoreAttributes field value
// and a boolean to check if the value has been set.
func (o *IdpAdapterAttributeContract) GetCoreAttributesOk() ([]IdpAdapterAttribute, bool) {
	if o == nil {
		return nil, false
	}
	return o.CoreAttributes, true
}

// SetCoreAttributes sets field value
func (o *IdpAdapterAttributeContract) SetCoreAttributes(v []IdpAdapterAttribute) {
	o.CoreAttributes = v
}

// GetExtendedAttributes returns the ExtendedAttributes field value if set, zero value otherwise.
func (o *IdpAdapterAttributeContract) GetExtendedAttributes() []IdpAdapterAttribute {
	if o == nil || IsNil(o.ExtendedAttributes) {
		var ret []IdpAdapterAttribute
		return ret
	}
	return o.ExtendedAttributes
}

// GetExtendedAttributesOk returns a tuple with the ExtendedAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpAdapterAttributeContract) GetExtendedAttributesOk() ([]IdpAdapterAttribute, bool) {
	if o == nil || IsNil(o.ExtendedAttributes) {
		return nil, false
	}
	return o.ExtendedAttributes, true
}

// HasExtendedAttributes returns a boolean if a field has been set.
func (o *IdpAdapterAttributeContract) HasExtendedAttributes() bool {
	if o != nil && !IsNil(o.ExtendedAttributes) {
		return true
	}

	return false
}

// SetExtendedAttributes gets a reference to the given []IdpAdapterAttribute and assigns it to the ExtendedAttributes field.
func (o *IdpAdapterAttributeContract) SetExtendedAttributes(v []IdpAdapterAttribute) {
	o.ExtendedAttributes = v
}

// GetUniqueUserKeyAttribute returns the UniqueUserKeyAttribute field value if set, zero value otherwise.
func (o *IdpAdapterAttributeContract) GetUniqueUserKeyAttribute() string {
	if o == nil || IsNil(o.UniqueUserKeyAttribute) {
		var ret string
		return ret
	}
	return *o.UniqueUserKeyAttribute
}

// GetUniqueUserKeyAttributeOk returns a tuple with the UniqueUserKeyAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpAdapterAttributeContract) GetUniqueUserKeyAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.UniqueUserKeyAttribute) {
		return nil, false
	}
	return o.UniqueUserKeyAttribute, true
}

// HasUniqueUserKeyAttribute returns a boolean if a field has been set.
func (o *IdpAdapterAttributeContract) HasUniqueUserKeyAttribute() bool {
	if o != nil && !IsNil(o.UniqueUserKeyAttribute) {
		return true
	}

	return false
}

// SetUniqueUserKeyAttribute gets a reference to the given string and assigns it to the UniqueUserKeyAttribute field.
func (o *IdpAdapterAttributeContract) SetUniqueUserKeyAttribute(v string) {
	o.UniqueUserKeyAttribute = &v
}

// GetMaskOgnlValues returns the MaskOgnlValues field value if set, zero value otherwise.
func (o *IdpAdapterAttributeContract) GetMaskOgnlValues() bool {
	if o == nil || IsNil(o.MaskOgnlValues) {
		var ret bool
		return ret
	}
	return *o.MaskOgnlValues
}

// GetMaskOgnlValuesOk returns a tuple with the MaskOgnlValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpAdapterAttributeContract) GetMaskOgnlValuesOk() (*bool, bool) {
	if o == nil || IsNil(o.MaskOgnlValues) {
		return nil, false
	}
	return o.MaskOgnlValues, true
}

// HasMaskOgnlValues returns a boolean if a field has been set.
func (o *IdpAdapterAttributeContract) HasMaskOgnlValues() bool {
	if o != nil && !IsNil(o.MaskOgnlValues) {
		return true
	}

	return false
}

// SetMaskOgnlValues gets a reference to the given bool and assigns it to the MaskOgnlValues field.
func (o *IdpAdapterAttributeContract) SetMaskOgnlValues(v bool) {
	o.MaskOgnlValues = &v
}

func (o IdpAdapterAttributeContract) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdpAdapterAttributeContract) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["coreAttributes"] = o.CoreAttributes
	if !IsNil(o.ExtendedAttributes) {
		toSerialize["extendedAttributes"] = o.ExtendedAttributes
	}
	if !IsNil(o.UniqueUserKeyAttribute) {
		toSerialize["uniqueUserKeyAttribute"] = o.UniqueUserKeyAttribute
	}
	if !IsNil(o.MaskOgnlValues) {
		toSerialize["maskOgnlValues"] = o.MaskOgnlValues
	}
	return toSerialize, nil
}

type NullableIdpAdapterAttributeContract struct {
	value *IdpAdapterAttributeContract
	isSet bool
}

func (v NullableIdpAdapterAttributeContract) Get() *IdpAdapterAttributeContract {
	return v.value
}

func (v *NullableIdpAdapterAttributeContract) Set(val *IdpAdapterAttributeContract) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpAdapterAttributeContract) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpAdapterAttributeContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpAdapterAttributeContract(val *IdpAdapterAttributeContract) *NullableIdpAdapterAttributeContract {
	return &NullableIdpAdapterAttributeContract{value: val, isSet: true}
}

func (v NullableIdpAdapterAttributeContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpAdapterAttributeContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
