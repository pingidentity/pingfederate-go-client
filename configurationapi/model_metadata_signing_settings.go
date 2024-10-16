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

// checks if the MetadataSigningSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataSigningSettings{}

// MetadataSigningSettings Metadata signing settings. If metadata is not signed, this model will be empty.
type MetadataSigningSettings struct {
	SigningKeyRef *ResourceLink `json:"signingKeyRef,omitempty" tfsdk:"signing_key_ref"`
	// Signature algorithm. If this property is unset, the default signature algorithm for the key algorithm will be used. Supported signature algorithms are available through the /keyPairs/keyAlgorithms endpoint.
	SignatureAlgorithm *string `json:"signatureAlgorithm,omitempty" tfsdk:"signature_algorithm"`
}

// NewMetadataSigningSettings instantiates a new MetadataSigningSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataSigningSettings() *MetadataSigningSettings {
	this := MetadataSigningSettings{}
	return &this
}

// NewMetadataSigningSettingsWithDefaults instantiates a new MetadataSigningSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataSigningSettingsWithDefaults() *MetadataSigningSettings {
	this := MetadataSigningSettings{}
	return &this
}

// GetSigningKeyRef returns the SigningKeyRef field value if set, zero value otherwise.
func (o *MetadataSigningSettings) GetSigningKeyRef() ResourceLink {
	if o == nil || IsNil(o.SigningKeyRef) {
		var ret ResourceLink
		return ret
	}
	return *o.SigningKeyRef
}

// GetSigningKeyRefOk returns a tuple with the SigningKeyRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataSigningSettings) GetSigningKeyRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.SigningKeyRef) {
		return nil, false
	}
	return o.SigningKeyRef, true
}

// HasSigningKeyRef returns a boolean if a field has been set.
func (o *MetadataSigningSettings) HasSigningKeyRef() bool {
	if o != nil && !IsNil(o.SigningKeyRef) {
		return true
	}

	return false
}

// SetSigningKeyRef gets a reference to the given ResourceLink and assigns it to the SigningKeyRef field.
func (o *MetadataSigningSettings) SetSigningKeyRef(v ResourceLink) {
	o.SigningKeyRef = &v
}

// GetSignatureAlgorithm returns the SignatureAlgorithm field value if set, zero value otherwise.
func (o *MetadataSigningSettings) GetSignatureAlgorithm() string {
	if o == nil || IsNil(o.SignatureAlgorithm) {
		var ret string
		return ret
	}
	return *o.SignatureAlgorithm
}

// GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataSigningSettings) GetSignatureAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureAlgorithm) {
		return nil, false
	}
	return o.SignatureAlgorithm, true
}

// HasSignatureAlgorithm returns a boolean if a field has been set.
func (o *MetadataSigningSettings) HasSignatureAlgorithm() bool {
	if o != nil && !IsNil(o.SignatureAlgorithm) {
		return true
	}

	return false
}

// SetSignatureAlgorithm gets a reference to the given string and assigns it to the SignatureAlgorithm field.
func (o *MetadataSigningSettings) SetSignatureAlgorithm(v string) {
	o.SignatureAlgorithm = &v
}

func (o MetadataSigningSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataSigningSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SigningKeyRef) {
		toSerialize["signingKeyRef"] = o.SigningKeyRef
	}
	if !IsNil(o.SignatureAlgorithm) {
		toSerialize["signatureAlgorithm"] = o.SignatureAlgorithm
	}
	return toSerialize, nil
}

type NullableMetadataSigningSettings struct {
	value *MetadataSigningSettings
	isSet bool
}

func (v NullableMetadataSigningSettings) Get() *MetadataSigningSettings {
	return v.value
}

func (v *NullableMetadataSigningSettings) Set(val *MetadataSigningSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataSigningSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataSigningSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataSigningSettings(val *MetadataSigningSettings) *NullableMetadataSigningSettings {
	return &NullableMetadataSigningSettings{value: val, isSet: true}
}

func (v NullableMetadataSigningSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataSigningSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
