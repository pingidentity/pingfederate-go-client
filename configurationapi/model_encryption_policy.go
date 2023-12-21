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

// checks if the EncryptionPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EncryptionPolicy{}

// EncryptionPolicy Defines what to encrypt in the browser-based SSO profile.
type EncryptionPolicy struct {
	// Whether the outgoing SAML assertion will be encrypted.
	EncryptAssertion *bool `json:"encryptAssertion,omitempty" tfsdk:"encrypt_assertion"`
	// The list of outgoing SAML assertion attributes that will be encrypted. The 'encryptAssertion' property takes precedence over this.
	EncryptedAttributes []string `json:"encryptedAttributes,omitempty" tfsdk:"encrypted_attributes"`
	// Encrypt the name-identifier attribute in outbound SLO messages.  This can be set if the name id is encrypted.
	EncryptSloSubjectNameId *bool `json:"encryptSloSubjectNameId,omitempty" tfsdk:"encrypt_slo_subject_name_id"`
	// Allow the encryption of the name-identifier attribute for inbound SLO messages. This can be set if SP initiated SLO is enabled.
	SloSubjectNameIDEncrypted *bool `json:"sloSubjectNameIDEncrypted,omitempty" tfsdk:"slo_subject_name_ide_ncrypted"`
}

// NewEncryptionPolicy instantiates a new EncryptionPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncryptionPolicy() *EncryptionPolicy {
	this := EncryptionPolicy{}
	return &this
}

// NewEncryptionPolicyWithDefaults instantiates a new EncryptionPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncryptionPolicyWithDefaults() *EncryptionPolicy {
	this := EncryptionPolicy{}
	return &this
}

// GetEncryptAssertion returns the EncryptAssertion field value if set, zero value otherwise.
func (o *EncryptionPolicy) GetEncryptAssertion() bool {
	if o == nil || IsNil(o.EncryptAssertion) {
		var ret bool
		return ret
	}
	return *o.EncryptAssertion
}

// GetEncryptAssertionOk returns a tuple with the EncryptAssertion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionPolicy) GetEncryptAssertionOk() (*bool, bool) {
	if o == nil || IsNil(o.EncryptAssertion) {
		return nil, false
	}
	return o.EncryptAssertion, true
}

// HasEncryptAssertion returns a boolean if a field has been set.
func (o *EncryptionPolicy) HasEncryptAssertion() bool {
	if o != nil && !IsNil(o.EncryptAssertion) {
		return true
	}

	return false
}

// SetEncryptAssertion gets a reference to the given bool and assigns it to the EncryptAssertion field.
func (o *EncryptionPolicy) SetEncryptAssertion(v bool) {
	o.EncryptAssertion = &v
}

// GetEncryptedAttributes returns the EncryptedAttributes field value if set, zero value otherwise.
func (o *EncryptionPolicy) GetEncryptedAttributes() []string {
	if o == nil || IsNil(o.EncryptedAttributes) {
		var ret []string
		return ret
	}
	return o.EncryptedAttributes
}

// GetEncryptedAttributesOk returns a tuple with the EncryptedAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionPolicy) GetEncryptedAttributesOk() ([]string, bool) {
	if o == nil || IsNil(o.EncryptedAttributes) {
		return nil, false
	}
	return o.EncryptedAttributes, true
}

// HasEncryptedAttributes returns a boolean if a field has been set.
func (o *EncryptionPolicy) HasEncryptedAttributes() bool {
	if o != nil && !IsNil(o.EncryptedAttributes) {
		return true
	}

	return false
}

// SetEncryptedAttributes gets a reference to the given []string and assigns it to the EncryptedAttributes field.
func (o *EncryptionPolicy) SetEncryptedAttributes(v []string) {
	o.EncryptedAttributes = v
}

// GetEncryptSloSubjectNameId returns the EncryptSloSubjectNameId field value if set, zero value otherwise.
func (o *EncryptionPolicy) GetEncryptSloSubjectNameId() bool {
	if o == nil || IsNil(o.EncryptSloSubjectNameId) {
		var ret bool
		return ret
	}
	return *o.EncryptSloSubjectNameId
}

// GetEncryptSloSubjectNameIdOk returns a tuple with the EncryptSloSubjectNameId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionPolicy) GetEncryptSloSubjectNameIdOk() (*bool, bool) {
	if o == nil || IsNil(o.EncryptSloSubjectNameId) {
		return nil, false
	}
	return o.EncryptSloSubjectNameId, true
}

// HasEncryptSloSubjectNameId returns a boolean if a field has been set.
func (o *EncryptionPolicy) HasEncryptSloSubjectNameId() bool {
	if o != nil && !IsNil(o.EncryptSloSubjectNameId) {
		return true
	}

	return false
}

// SetEncryptSloSubjectNameId gets a reference to the given bool and assigns it to the EncryptSloSubjectNameId field.
func (o *EncryptionPolicy) SetEncryptSloSubjectNameId(v bool) {
	o.EncryptSloSubjectNameId = &v
}

// GetSloSubjectNameIDEncrypted returns the SloSubjectNameIDEncrypted field value if set, zero value otherwise.
func (o *EncryptionPolicy) GetSloSubjectNameIDEncrypted() bool {
	if o == nil || IsNil(o.SloSubjectNameIDEncrypted) {
		var ret bool
		return ret
	}
	return *o.SloSubjectNameIDEncrypted
}

// GetSloSubjectNameIDEncryptedOk returns a tuple with the SloSubjectNameIDEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncryptionPolicy) GetSloSubjectNameIDEncryptedOk() (*bool, bool) {
	if o == nil || IsNil(o.SloSubjectNameIDEncrypted) {
		return nil, false
	}
	return o.SloSubjectNameIDEncrypted, true
}

// HasSloSubjectNameIDEncrypted returns a boolean if a field has been set.
func (o *EncryptionPolicy) HasSloSubjectNameIDEncrypted() bool {
	if o != nil && !IsNil(o.SloSubjectNameIDEncrypted) {
		return true
	}

	return false
}

// SetSloSubjectNameIDEncrypted gets a reference to the given bool and assigns it to the SloSubjectNameIDEncrypted field.
func (o *EncryptionPolicy) SetSloSubjectNameIDEncrypted(v bool) {
	o.SloSubjectNameIDEncrypted = &v
}

func (o EncryptionPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EncryptionPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EncryptAssertion) {
		toSerialize["encryptAssertion"] = o.EncryptAssertion
	}
	if !IsNil(o.EncryptedAttributes) {
		toSerialize["encryptedAttributes"] = o.EncryptedAttributes
	}
	if !IsNil(o.EncryptSloSubjectNameId) {
		toSerialize["encryptSloSubjectNameId"] = o.EncryptSloSubjectNameId
	}
	if !IsNil(o.SloSubjectNameIDEncrypted) {
		toSerialize["sloSubjectNameIDEncrypted"] = o.SloSubjectNameIDEncrypted
	}
	return toSerialize, nil
}

type NullableEncryptionPolicy struct {
	value *EncryptionPolicy
	isSet bool
}

func (v NullableEncryptionPolicy) Get() *EncryptionPolicy {
	return v.value
}

func (v *NullableEncryptionPolicy) Set(val *EncryptionPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableEncryptionPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableEncryptionPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncryptionPolicy(val *EncryptionPolicy) *NullableEncryptionPolicy {
	return &NullableEncryptionPolicy{value: val, isSet: true}
}

func (v NullableEncryptionPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncryptionPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
