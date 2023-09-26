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

// checks if the KeyAlgorithm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyAlgorithm{}

// KeyAlgorithm Details for a key algorithm.
type KeyAlgorithm struct {
	// Name of the key algorithm.
	Name *string `json:"name,omitempty" tfsdk:"name"`
	// Possible key sizes for this algorithm, in bits.
	KeySizes []int64 `json:"keySizes,omitempty" tfsdk:"key_sizes"`
	// Default key size for this algorithm.
	DefaultKeySize *int64 `json:"defaultKeySize,omitempty" tfsdk:"default_key_size"`
	// Possible signature algorithms for this key algorithm.
	SignatureAlgorithms []string `json:"signatureAlgorithms,omitempty" tfsdk:"signature_algorithms"`
	// Default signature algorithm for this key algorithm.
	DefaultSignatureAlgorithm *string `json:"defaultSignatureAlgorithm,omitempty" tfsdk:"default_signature_algorithm"`
}

// NewKeyAlgorithm instantiates a new KeyAlgorithm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyAlgorithm() *KeyAlgorithm {
	this := KeyAlgorithm{}
	return &this
}

// NewKeyAlgorithmWithDefaults instantiates a new KeyAlgorithm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyAlgorithmWithDefaults() *KeyAlgorithm {
	this := KeyAlgorithm{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KeyAlgorithm) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyAlgorithm) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KeyAlgorithm) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KeyAlgorithm) SetName(v string) {
	o.Name = &v
}

// GetKeySizes returns the KeySizes field value if set, zero value otherwise.
func (o *KeyAlgorithm) GetKeySizes() []int64 {
	if o == nil || IsNil(o.KeySizes) {
		var ret []int64
		return ret
	}
	return o.KeySizes
}

// GetKeySizesOk returns a tuple with the KeySizes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyAlgorithm) GetKeySizesOk() ([]int64, bool) {
	if o == nil || IsNil(o.KeySizes) {
		return nil, false
	}
	return o.KeySizes, true
}

// HasKeySizes returns a boolean if a field has been set.
func (o *KeyAlgorithm) HasKeySizes() bool {
	if o != nil && !IsNil(o.KeySizes) {
		return true
	}

	return false
}

// SetKeySizes gets a reference to the given []int64 and assigns it to the KeySizes field.
func (o *KeyAlgorithm) SetKeySizes(v []int64) {
	o.KeySizes = v
}

// GetDefaultKeySize returns the DefaultKeySize field value if set, zero value otherwise.
func (o *KeyAlgorithm) GetDefaultKeySize() int64 {
	if o == nil || IsNil(o.DefaultKeySize) {
		var ret int64
		return ret
	}
	return *o.DefaultKeySize
}

// GetDefaultKeySizeOk returns a tuple with the DefaultKeySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyAlgorithm) GetDefaultKeySizeOk() (*int64, bool) {
	if o == nil || IsNil(o.DefaultKeySize) {
		return nil, false
	}
	return o.DefaultKeySize, true
}

// HasDefaultKeySize returns a boolean if a field has been set.
func (o *KeyAlgorithm) HasDefaultKeySize() bool {
	if o != nil && !IsNil(o.DefaultKeySize) {
		return true
	}

	return false
}

// SetDefaultKeySize gets a reference to the given int64 and assigns it to the DefaultKeySize field.
func (o *KeyAlgorithm) SetDefaultKeySize(v int64) {
	o.DefaultKeySize = &v
}

// GetSignatureAlgorithms returns the SignatureAlgorithms field value if set, zero value otherwise.
func (o *KeyAlgorithm) GetSignatureAlgorithms() []string {
	if o == nil || IsNil(o.SignatureAlgorithms) {
		var ret []string
		return ret
	}
	return o.SignatureAlgorithms
}

// GetSignatureAlgorithmsOk returns a tuple with the SignatureAlgorithms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyAlgorithm) GetSignatureAlgorithmsOk() ([]string, bool) {
	if o == nil || IsNil(o.SignatureAlgorithms) {
		return nil, false
	}
	return o.SignatureAlgorithms, true
}

// HasSignatureAlgorithms returns a boolean if a field has been set.
func (o *KeyAlgorithm) HasSignatureAlgorithms() bool {
	if o != nil && !IsNil(o.SignatureAlgorithms) {
		return true
	}

	return false
}

// SetSignatureAlgorithms gets a reference to the given []string and assigns it to the SignatureAlgorithms field.
func (o *KeyAlgorithm) SetSignatureAlgorithms(v []string) {
	o.SignatureAlgorithms = v
}

// GetDefaultSignatureAlgorithm returns the DefaultSignatureAlgorithm field value if set, zero value otherwise.
func (o *KeyAlgorithm) GetDefaultSignatureAlgorithm() string {
	if o == nil || IsNil(o.DefaultSignatureAlgorithm) {
		var ret string
		return ret
	}
	return *o.DefaultSignatureAlgorithm
}

// GetDefaultSignatureAlgorithmOk returns a tuple with the DefaultSignatureAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyAlgorithm) GetDefaultSignatureAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultSignatureAlgorithm) {
		return nil, false
	}
	return o.DefaultSignatureAlgorithm, true
}

// HasDefaultSignatureAlgorithm returns a boolean if a field has been set.
func (o *KeyAlgorithm) HasDefaultSignatureAlgorithm() bool {
	if o != nil && !IsNil(o.DefaultSignatureAlgorithm) {
		return true
	}

	return false
}

// SetDefaultSignatureAlgorithm gets a reference to the given string and assigns it to the DefaultSignatureAlgorithm field.
func (o *KeyAlgorithm) SetDefaultSignatureAlgorithm(v string) {
	o.DefaultSignatureAlgorithm = &v
}

func (o KeyAlgorithm) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyAlgorithm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.KeySizes) {
		toSerialize["keySizes"] = o.KeySizes
	}
	if !IsNil(o.DefaultKeySize) {
		toSerialize["defaultKeySize"] = o.DefaultKeySize
	}
	if !IsNil(o.SignatureAlgorithms) {
		toSerialize["signatureAlgorithms"] = o.SignatureAlgorithms
	}
	if !IsNil(o.DefaultSignatureAlgorithm) {
		toSerialize["defaultSignatureAlgorithm"] = o.DefaultSignatureAlgorithm
	}
	return toSerialize, nil
}

type NullableKeyAlgorithm struct {
	value *KeyAlgorithm
	isSet bool
}

func (v NullableKeyAlgorithm) Get() *KeyAlgorithm {
	return v.value
}

func (v *NullableKeyAlgorithm) Set(val *KeyAlgorithm) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyAlgorithm) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyAlgorithm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyAlgorithm(val *KeyAlgorithm) *NullableKeyAlgorithm {
	return &NullableKeyAlgorithm{value: val, isSet: true}
}

func (v NullableKeyAlgorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyAlgorithm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}