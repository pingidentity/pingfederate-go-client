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

// checks if the IdpAttributeQueryPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdpAttributeQueryPolicy{}

// IdpAttributeQueryPolicy The attribute query profile's security policy.
type IdpAttributeQueryPolicy struct {
	// Require signed response.
	RequireSignedResponse *bool `json:"requireSignedResponse,omitempty" tfsdk:"require_signed_response"`
	// Require signed assertion.
	RequireSignedAssertion *bool `json:"requireSignedAssertion,omitempty" tfsdk:"require_signed_assertion"`
	// Require encrypted assertion.
	RequireEncryptedAssertion *bool `json:"requireEncryptedAssertion,omitempty" tfsdk:"require_encrypted_assertion"`
	// Sign the attribute query.
	SignAttributeQuery *bool `json:"signAttributeQuery,omitempty" tfsdk:"sign_attribute_query"`
	// Encrypt the name identifier.
	EncryptNameId *bool `json:"encryptNameId,omitempty" tfsdk:"encrypt_name_id"`
	// Mask attributes in log files.
	MaskAttributeValues *bool `json:"maskAttributeValues,omitempty" tfsdk:"mask_attribute_values"`
}

// NewIdpAttributeQueryPolicy instantiates a new IdpAttributeQueryPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdpAttributeQueryPolicy() *IdpAttributeQueryPolicy {
	this := IdpAttributeQueryPolicy{}
	return &this
}

// NewIdpAttributeQueryPolicyWithDefaults instantiates a new IdpAttributeQueryPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdpAttributeQueryPolicyWithDefaults() *IdpAttributeQueryPolicy {
	this := IdpAttributeQueryPolicy{}
	return &this
}

// GetRequireSignedResponse returns the RequireSignedResponse field value if set, zero value otherwise.
func (o *IdpAttributeQueryPolicy) GetRequireSignedResponse() bool {
	if o == nil || IsNil(o.RequireSignedResponse) {
		var ret bool
		return ret
	}
	return *o.RequireSignedResponse
}

// GetRequireSignedResponseOk returns a tuple with the RequireSignedResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpAttributeQueryPolicy) GetRequireSignedResponseOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireSignedResponse) {
		return nil, false
	}
	return o.RequireSignedResponse, true
}

// HasRequireSignedResponse returns a boolean if a field has been set.
func (o *IdpAttributeQueryPolicy) HasRequireSignedResponse() bool {
	if o != nil && !IsNil(o.RequireSignedResponse) {
		return true
	}

	return false
}

// SetRequireSignedResponse gets a reference to the given bool and assigns it to the RequireSignedResponse field.
func (o *IdpAttributeQueryPolicy) SetRequireSignedResponse(v bool) {
	o.RequireSignedResponse = &v
}

// GetRequireSignedAssertion returns the RequireSignedAssertion field value if set, zero value otherwise.
func (o *IdpAttributeQueryPolicy) GetRequireSignedAssertion() bool {
	if o == nil || IsNil(o.RequireSignedAssertion) {
		var ret bool
		return ret
	}
	return *o.RequireSignedAssertion
}

// GetRequireSignedAssertionOk returns a tuple with the RequireSignedAssertion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpAttributeQueryPolicy) GetRequireSignedAssertionOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireSignedAssertion) {
		return nil, false
	}
	return o.RequireSignedAssertion, true
}

// HasRequireSignedAssertion returns a boolean if a field has been set.
func (o *IdpAttributeQueryPolicy) HasRequireSignedAssertion() bool {
	if o != nil && !IsNil(o.RequireSignedAssertion) {
		return true
	}

	return false
}

// SetRequireSignedAssertion gets a reference to the given bool and assigns it to the RequireSignedAssertion field.
func (o *IdpAttributeQueryPolicy) SetRequireSignedAssertion(v bool) {
	o.RequireSignedAssertion = &v
}

// GetRequireEncryptedAssertion returns the RequireEncryptedAssertion field value if set, zero value otherwise.
func (o *IdpAttributeQueryPolicy) GetRequireEncryptedAssertion() bool {
	if o == nil || IsNil(o.RequireEncryptedAssertion) {
		var ret bool
		return ret
	}
	return *o.RequireEncryptedAssertion
}

// GetRequireEncryptedAssertionOk returns a tuple with the RequireEncryptedAssertion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpAttributeQueryPolicy) GetRequireEncryptedAssertionOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireEncryptedAssertion) {
		return nil, false
	}
	return o.RequireEncryptedAssertion, true
}

// HasRequireEncryptedAssertion returns a boolean if a field has been set.
func (o *IdpAttributeQueryPolicy) HasRequireEncryptedAssertion() bool {
	if o != nil && !IsNil(o.RequireEncryptedAssertion) {
		return true
	}

	return false
}

// SetRequireEncryptedAssertion gets a reference to the given bool and assigns it to the RequireEncryptedAssertion field.
func (o *IdpAttributeQueryPolicy) SetRequireEncryptedAssertion(v bool) {
	o.RequireEncryptedAssertion = &v
}

// GetSignAttributeQuery returns the SignAttributeQuery field value if set, zero value otherwise.
func (o *IdpAttributeQueryPolicy) GetSignAttributeQuery() bool {
	if o == nil || IsNil(o.SignAttributeQuery) {
		var ret bool
		return ret
	}
	return *o.SignAttributeQuery
}

// GetSignAttributeQueryOk returns a tuple with the SignAttributeQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpAttributeQueryPolicy) GetSignAttributeQueryOk() (*bool, bool) {
	if o == nil || IsNil(o.SignAttributeQuery) {
		return nil, false
	}
	return o.SignAttributeQuery, true
}

// HasSignAttributeQuery returns a boolean if a field has been set.
func (o *IdpAttributeQueryPolicy) HasSignAttributeQuery() bool {
	if o != nil && !IsNil(o.SignAttributeQuery) {
		return true
	}

	return false
}

// SetSignAttributeQuery gets a reference to the given bool and assigns it to the SignAttributeQuery field.
func (o *IdpAttributeQueryPolicy) SetSignAttributeQuery(v bool) {
	o.SignAttributeQuery = &v
}

// GetEncryptNameId returns the EncryptNameId field value if set, zero value otherwise.
func (o *IdpAttributeQueryPolicy) GetEncryptNameId() bool {
	if o == nil || IsNil(o.EncryptNameId) {
		var ret bool
		return ret
	}
	return *o.EncryptNameId
}

// GetEncryptNameIdOk returns a tuple with the EncryptNameId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpAttributeQueryPolicy) GetEncryptNameIdOk() (*bool, bool) {
	if o == nil || IsNil(o.EncryptNameId) {
		return nil, false
	}
	return o.EncryptNameId, true
}

// HasEncryptNameId returns a boolean if a field has been set.
func (o *IdpAttributeQueryPolicy) HasEncryptNameId() bool {
	if o != nil && !IsNil(o.EncryptNameId) {
		return true
	}

	return false
}

// SetEncryptNameId gets a reference to the given bool and assigns it to the EncryptNameId field.
func (o *IdpAttributeQueryPolicy) SetEncryptNameId(v bool) {
	o.EncryptNameId = &v
}

// GetMaskAttributeValues returns the MaskAttributeValues field value if set, zero value otherwise.
func (o *IdpAttributeQueryPolicy) GetMaskAttributeValues() bool {
	if o == nil || IsNil(o.MaskAttributeValues) {
		var ret bool
		return ret
	}
	return *o.MaskAttributeValues
}

// GetMaskAttributeValuesOk returns a tuple with the MaskAttributeValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdpAttributeQueryPolicy) GetMaskAttributeValuesOk() (*bool, bool) {
	if o == nil || IsNil(o.MaskAttributeValues) {
		return nil, false
	}
	return o.MaskAttributeValues, true
}

// HasMaskAttributeValues returns a boolean if a field has been set.
func (o *IdpAttributeQueryPolicy) HasMaskAttributeValues() bool {
	if o != nil && !IsNil(o.MaskAttributeValues) {
		return true
	}

	return false
}

// SetMaskAttributeValues gets a reference to the given bool and assigns it to the MaskAttributeValues field.
func (o *IdpAttributeQueryPolicy) SetMaskAttributeValues(v bool) {
	o.MaskAttributeValues = &v
}

func (o IdpAttributeQueryPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdpAttributeQueryPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequireSignedResponse) {
		toSerialize["requireSignedResponse"] = o.RequireSignedResponse
	}
	if !IsNil(o.RequireSignedAssertion) {
		toSerialize["requireSignedAssertion"] = o.RequireSignedAssertion
	}
	if !IsNil(o.RequireEncryptedAssertion) {
		toSerialize["requireEncryptedAssertion"] = o.RequireEncryptedAssertion
	}
	if !IsNil(o.SignAttributeQuery) {
		toSerialize["signAttributeQuery"] = o.SignAttributeQuery
	}
	if !IsNil(o.EncryptNameId) {
		toSerialize["encryptNameId"] = o.EncryptNameId
	}
	if !IsNil(o.MaskAttributeValues) {
		toSerialize["maskAttributeValues"] = o.MaskAttributeValues
	}
	return toSerialize, nil
}

type NullableIdpAttributeQueryPolicy struct {
	value *IdpAttributeQueryPolicy
	isSet bool
}

func (v NullableIdpAttributeQueryPolicy) Get() *IdpAttributeQueryPolicy {
	return v.value
}

func (v *NullableIdpAttributeQueryPolicy) Set(val *IdpAttributeQueryPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableIdpAttributeQueryPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableIdpAttributeQueryPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdpAttributeQueryPolicy(val *IdpAttributeQueryPolicy) *NullableIdpAttributeQueryPolicy {
	return &NullableIdpAttributeQueryPolicy{value: val, isSet: true}
}

func (v NullableIdpAttributeQueryPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdpAttributeQueryPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}