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

// checks if the SigningKeys type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SigningKeys{}

// SigningKeys Setting for a OAuth/OpenID Connect signing key set while using multiple virtual issuers.
type SigningKeys struct {
	P256ActiveCertRef   *ResourceLink `json:"p256ActiveCertRef,omitempty" tfsdk:"p256_active_cert_ref"`
	P256PreviousCertRef *ResourceLink `json:"p256PreviousCertRef,omitempty" tfsdk:"p256_previous_cert_ref"`
	// Enable publishing of the P-256 certificate chain associated with the active key.
	P256PublishX5cParameter *bool         `json:"p256PublishX5cParameter,omitempty" tfsdk:"p256_publish_x5c_parameter"`
	P384ActiveCertRef       *ResourceLink `json:"p384ActiveCertRef,omitempty" tfsdk:"p384_active_cert_ref"`
	P384PreviousCertRef     *ResourceLink `json:"p384PreviousCertRef,omitempty" tfsdk:"p384_previous_cert_ref"`
	// Enable publishing of the P-384 certificate chain associated with the active key.
	P384PublishX5cParameter *bool         `json:"p384PublishX5cParameter,omitempty" tfsdk:"p384_publish_x5c_parameter"`
	P521ActiveCertRef       *ResourceLink `json:"p521ActiveCertRef,omitempty" tfsdk:"p521_active_cert_ref"`
	P521PreviousCertRef     *ResourceLink `json:"p521PreviousCertRef,omitempty" tfsdk:"p521_previous_cert_ref"`
	// Enable publishing of the P-521 certificate chain associated with the active key.
	P521PublishX5cParameter *bool         `json:"p521PublishX5cParameter,omitempty" tfsdk:"p521_publish_x5c_parameter"`
	RsaActiveCertRef        *ResourceLink `json:"rsaActiveCertRef,omitempty" tfsdk:"rsa_active_cert_ref"`
	RsaPreviousCertRef      *ResourceLink `json:"rsaPreviousCertRef,omitempty" tfsdk:"rsa_previous_cert_ref"`
	// Enable publishing of the RSA certificate chain associated with the active key.
	RsaPublishX5cParameter *bool `json:"rsaPublishX5cParameter,omitempty" tfsdk:"rsa_publish_x5c_parameter"`
}

// NewSigningKeys instantiates a new SigningKeys object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSigningKeys() *SigningKeys {
	this := SigningKeys{}
	return &this
}

// NewSigningKeysWithDefaults instantiates a new SigningKeys object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSigningKeysWithDefaults() *SigningKeys {
	this := SigningKeys{}
	return &this
}

// GetP256ActiveCertRef returns the P256ActiveCertRef field value if set, zero value otherwise.
func (o *SigningKeys) GetP256ActiveCertRef() ResourceLink {
	if o == nil || IsNil(o.P256ActiveCertRef) {
		var ret ResourceLink
		return ret
	}
	return *o.P256ActiveCertRef
}

// GetP256ActiveCertRefOk returns a tuple with the P256ActiveCertRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigningKeys) GetP256ActiveCertRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.P256ActiveCertRef) {
		return nil, false
	}
	return o.P256ActiveCertRef, true
}

// HasP256ActiveCertRef returns a boolean if a field has been set.
func (o *SigningKeys) HasP256ActiveCertRef() bool {
	if o != nil && !IsNil(o.P256ActiveCertRef) {
		return true
	}

	return false
}

// SetP256ActiveCertRef gets a reference to the given ResourceLink and assigns it to the P256ActiveCertRef field.
func (o *SigningKeys) SetP256ActiveCertRef(v ResourceLink) {
	o.P256ActiveCertRef = &v
}

// GetP256PreviousCertRef returns the P256PreviousCertRef field value if set, zero value otherwise.
func (o *SigningKeys) GetP256PreviousCertRef() ResourceLink {
	if o == nil || IsNil(o.P256PreviousCertRef) {
		var ret ResourceLink
		return ret
	}
	return *o.P256PreviousCertRef
}

// GetP256PreviousCertRefOk returns a tuple with the P256PreviousCertRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigningKeys) GetP256PreviousCertRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.P256PreviousCertRef) {
		return nil, false
	}
	return o.P256PreviousCertRef, true
}

// HasP256PreviousCertRef returns a boolean if a field has been set.
func (o *SigningKeys) HasP256PreviousCertRef() bool {
	if o != nil && !IsNil(o.P256PreviousCertRef) {
		return true
	}

	return false
}

// SetP256PreviousCertRef gets a reference to the given ResourceLink and assigns it to the P256PreviousCertRef field.
func (o *SigningKeys) SetP256PreviousCertRef(v ResourceLink) {
	o.P256PreviousCertRef = &v
}

// GetP256PublishX5cParameter returns the P256PublishX5cParameter field value if set, zero value otherwise.
func (o *SigningKeys) GetP256PublishX5cParameter() bool {
	if o == nil || IsNil(o.P256PublishX5cParameter) {
		var ret bool
		return ret
	}
	return *o.P256PublishX5cParameter
}

// GetP256PublishX5cParameterOk returns a tuple with the P256PublishX5cParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigningKeys) GetP256PublishX5cParameterOk() (*bool, bool) {
	if o == nil || IsNil(o.P256PublishX5cParameter) {
		return nil, false
	}
	return o.P256PublishX5cParameter, true
}

// HasP256PublishX5cParameter returns a boolean if a field has been set.
func (o *SigningKeys) HasP256PublishX5cParameter() bool {
	if o != nil && !IsNil(o.P256PublishX5cParameter) {
		return true
	}

	return false
}

// SetP256PublishX5cParameter gets a reference to the given bool and assigns it to the P256PublishX5cParameter field.
func (o *SigningKeys) SetP256PublishX5cParameter(v bool) {
	o.P256PublishX5cParameter = &v
}

// GetP384ActiveCertRef returns the P384ActiveCertRef field value if set, zero value otherwise.
func (o *SigningKeys) GetP384ActiveCertRef() ResourceLink {
	if o == nil || IsNil(o.P384ActiveCertRef) {
		var ret ResourceLink
		return ret
	}
	return *o.P384ActiveCertRef
}

// GetP384ActiveCertRefOk returns a tuple with the P384ActiveCertRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigningKeys) GetP384ActiveCertRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.P384ActiveCertRef) {
		return nil, false
	}
	return o.P384ActiveCertRef, true
}

// HasP384ActiveCertRef returns a boolean if a field has been set.
func (o *SigningKeys) HasP384ActiveCertRef() bool {
	if o != nil && !IsNil(o.P384ActiveCertRef) {
		return true
	}

	return false
}

// SetP384ActiveCertRef gets a reference to the given ResourceLink and assigns it to the P384ActiveCertRef field.
func (o *SigningKeys) SetP384ActiveCertRef(v ResourceLink) {
	o.P384ActiveCertRef = &v
}

// GetP384PreviousCertRef returns the P384PreviousCertRef field value if set, zero value otherwise.
func (o *SigningKeys) GetP384PreviousCertRef() ResourceLink {
	if o == nil || IsNil(o.P384PreviousCertRef) {
		var ret ResourceLink
		return ret
	}
	return *o.P384PreviousCertRef
}

// GetP384PreviousCertRefOk returns a tuple with the P384PreviousCertRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigningKeys) GetP384PreviousCertRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.P384PreviousCertRef) {
		return nil, false
	}
	return o.P384PreviousCertRef, true
}

// HasP384PreviousCertRef returns a boolean if a field has been set.
func (o *SigningKeys) HasP384PreviousCertRef() bool {
	if o != nil && !IsNil(o.P384PreviousCertRef) {
		return true
	}

	return false
}

// SetP384PreviousCertRef gets a reference to the given ResourceLink and assigns it to the P384PreviousCertRef field.
func (o *SigningKeys) SetP384PreviousCertRef(v ResourceLink) {
	o.P384PreviousCertRef = &v
}

// GetP384PublishX5cParameter returns the P384PublishX5cParameter field value if set, zero value otherwise.
func (o *SigningKeys) GetP384PublishX5cParameter() bool {
	if o == nil || IsNil(o.P384PublishX5cParameter) {
		var ret bool
		return ret
	}
	return *o.P384PublishX5cParameter
}

// GetP384PublishX5cParameterOk returns a tuple with the P384PublishX5cParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigningKeys) GetP384PublishX5cParameterOk() (*bool, bool) {
	if o == nil || IsNil(o.P384PublishX5cParameter) {
		return nil, false
	}
	return o.P384PublishX5cParameter, true
}

// HasP384PublishX5cParameter returns a boolean if a field has been set.
func (o *SigningKeys) HasP384PublishX5cParameter() bool {
	if o != nil && !IsNil(o.P384PublishX5cParameter) {
		return true
	}

	return false
}

// SetP384PublishX5cParameter gets a reference to the given bool and assigns it to the P384PublishX5cParameter field.
func (o *SigningKeys) SetP384PublishX5cParameter(v bool) {
	o.P384PublishX5cParameter = &v
}

// GetP521ActiveCertRef returns the P521ActiveCertRef field value if set, zero value otherwise.
func (o *SigningKeys) GetP521ActiveCertRef() ResourceLink {
	if o == nil || IsNil(o.P521ActiveCertRef) {
		var ret ResourceLink
		return ret
	}
	return *o.P521ActiveCertRef
}

// GetP521ActiveCertRefOk returns a tuple with the P521ActiveCertRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigningKeys) GetP521ActiveCertRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.P521ActiveCertRef) {
		return nil, false
	}
	return o.P521ActiveCertRef, true
}

// HasP521ActiveCertRef returns a boolean if a field has been set.
func (o *SigningKeys) HasP521ActiveCertRef() bool {
	if o != nil && !IsNil(o.P521ActiveCertRef) {
		return true
	}

	return false
}

// SetP521ActiveCertRef gets a reference to the given ResourceLink and assigns it to the P521ActiveCertRef field.
func (o *SigningKeys) SetP521ActiveCertRef(v ResourceLink) {
	o.P521ActiveCertRef = &v
}

// GetP521PreviousCertRef returns the P521PreviousCertRef field value if set, zero value otherwise.
func (o *SigningKeys) GetP521PreviousCertRef() ResourceLink {
	if o == nil || IsNil(o.P521PreviousCertRef) {
		var ret ResourceLink
		return ret
	}
	return *o.P521PreviousCertRef
}

// GetP521PreviousCertRefOk returns a tuple with the P521PreviousCertRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigningKeys) GetP521PreviousCertRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.P521PreviousCertRef) {
		return nil, false
	}
	return o.P521PreviousCertRef, true
}

// HasP521PreviousCertRef returns a boolean if a field has been set.
func (o *SigningKeys) HasP521PreviousCertRef() bool {
	if o != nil && !IsNil(o.P521PreviousCertRef) {
		return true
	}

	return false
}

// SetP521PreviousCertRef gets a reference to the given ResourceLink and assigns it to the P521PreviousCertRef field.
func (o *SigningKeys) SetP521PreviousCertRef(v ResourceLink) {
	o.P521PreviousCertRef = &v
}

// GetP521PublishX5cParameter returns the P521PublishX5cParameter field value if set, zero value otherwise.
func (o *SigningKeys) GetP521PublishX5cParameter() bool {
	if o == nil || IsNil(o.P521PublishX5cParameter) {
		var ret bool
		return ret
	}
	return *o.P521PublishX5cParameter
}

// GetP521PublishX5cParameterOk returns a tuple with the P521PublishX5cParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigningKeys) GetP521PublishX5cParameterOk() (*bool, bool) {
	if o == nil || IsNil(o.P521PublishX5cParameter) {
		return nil, false
	}
	return o.P521PublishX5cParameter, true
}

// HasP521PublishX5cParameter returns a boolean if a field has been set.
func (o *SigningKeys) HasP521PublishX5cParameter() bool {
	if o != nil && !IsNil(o.P521PublishX5cParameter) {
		return true
	}

	return false
}

// SetP521PublishX5cParameter gets a reference to the given bool and assigns it to the P521PublishX5cParameter field.
func (o *SigningKeys) SetP521PublishX5cParameter(v bool) {
	o.P521PublishX5cParameter = &v
}

// GetRsaActiveCertRef returns the RsaActiveCertRef field value if set, zero value otherwise.
func (o *SigningKeys) GetRsaActiveCertRef() ResourceLink {
	if o == nil || IsNil(o.RsaActiveCertRef) {
		var ret ResourceLink
		return ret
	}
	return *o.RsaActiveCertRef
}

// GetRsaActiveCertRefOk returns a tuple with the RsaActiveCertRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigningKeys) GetRsaActiveCertRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.RsaActiveCertRef) {
		return nil, false
	}
	return o.RsaActiveCertRef, true
}

// HasRsaActiveCertRef returns a boolean if a field has been set.
func (o *SigningKeys) HasRsaActiveCertRef() bool {
	if o != nil && !IsNil(o.RsaActiveCertRef) {
		return true
	}

	return false
}

// SetRsaActiveCertRef gets a reference to the given ResourceLink and assigns it to the RsaActiveCertRef field.
func (o *SigningKeys) SetRsaActiveCertRef(v ResourceLink) {
	o.RsaActiveCertRef = &v
}

// GetRsaPreviousCertRef returns the RsaPreviousCertRef field value if set, zero value otherwise.
func (o *SigningKeys) GetRsaPreviousCertRef() ResourceLink {
	if o == nil || IsNil(o.RsaPreviousCertRef) {
		var ret ResourceLink
		return ret
	}
	return *o.RsaPreviousCertRef
}

// GetRsaPreviousCertRefOk returns a tuple with the RsaPreviousCertRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigningKeys) GetRsaPreviousCertRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.RsaPreviousCertRef) {
		return nil, false
	}
	return o.RsaPreviousCertRef, true
}

// HasRsaPreviousCertRef returns a boolean if a field has been set.
func (o *SigningKeys) HasRsaPreviousCertRef() bool {
	if o != nil && !IsNil(o.RsaPreviousCertRef) {
		return true
	}

	return false
}

// SetRsaPreviousCertRef gets a reference to the given ResourceLink and assigns it to the RsaPreviousCertRef field.
func (o *SigningKeys) SetRsaPreviousCertRef(v ResourceLink) {
	o.RsaPreviousCertRef = &v
}

// GetRsaPublishX5cParameter returns the RsaPublishX5cParameter field value if set, zero value otherwise.
func (o *SigningKeys) GetRsaPublishX5cParameter() bool {
	if o == nil || IsNil(o.RsaPublishX5cParameter) {
		var ret bool
		return ret
	}
	return *o.RsaPublishX5cParameter
}

// GetRsaPublishX5cParameterOk returns a tuple with the RsaPublishX5cParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigningKeys) GetRsaPublishX5cParameterOk() (*bool, bool) {
	if o == nil || IsNil(o.RsaPublishX5cParameter) {
		return nil, false
	}
	return o.RsaPublishX5cParameter, true
}

// HasRsaPublishX5cParameter returns a boolean if a field has been set.
func (o *SigningKeys) HasRsaPublishX5cParameter() bool {
	if o != nil && !IsNil(o.RsaPublishX5cParameter) {
		return true
	}

	return false
}

// SetRsaPublishX5cParameter gets a reference to the given bool and assigns it to the RsaPublishX5cParameter field.
func (o *SigningKeys) SetRsaPublishX5cParameter(v bool) {
	o.RsaPublishX5cParameter = &v
}

func (o SigningKeys) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SigningKeys) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.P256ActiveCertRef) {
		toSerialize["p256ActiveCertRef"] = o.P256ActiveCertRef
	}
	if !IsNil(o.P256PreviousCertRef) {
		toSerialize["p256PreviousCertRef"] = o.P256PreviousCertRef
	}
	if !IsNil(o.P256PublishX5cParameter) {
		toSerialize["p256PublishX5cParameter"] = o.P256PublishX5cParameter
	}
	if !IsNil(o.P384ActiveCertRef) {
		toSerialize["p384ActiveCertRef"] = o.P384ActiveCertRef
	}
	if !IsNil(o.P384PreviousCertRef) {
		toSerialize["p384PreviousCertRef"] = o.P384PreviousCertRef
	}
	if !IsNil(o.P384PublishX5cParameter) {
		toSerialize["p384PublishX5cParameter"] = o.P384PublishX5cParameter
	}
	if !IsNil(o.P521ActiveCertRef) {
		toSerialize["p521ActiveCertRef"] = o.P521ActiveCertRef
	}
	if !IsNil(o.P521PreviousCertRef) {
		toSerialize["p521PreviousCertRef"] = o.P521PreviousCertRef
	}
	if !IsNil(o.P521PublishX5cParameter) {
		toSerialize["p521PublishX5cParameter"] = o.P521PublishX5cParameter
	}
	if !IsNil(o.RsaActiveCertRef) {
		toSerialize["rsaActiveCertRef"] = o.RsaActiveCertRef
	}
	if !IsNil(o.RsaPreviousCertRef) {
		toSerialize["rsaPreviousCertRef"] = o.RsaPreviousCertRef
	}
	if !IsNil(o.RsaPublishX5cParameter) {
		toSerialize["rsaPublishX5cParameter"] = o.RsaPublishX5cParameter
	}
	return toSerialize, nil
}

type NullableSigningKeys struct {
	value *SigningKeys
	isSet bool
}

func (v NullableSigningKeys) Get() *SigningKeys {
	return v.value
}

func (v *NullableSigningKeys) Set(val *SigningKeys) {
	v.value = val
	v.isSet = true
}

func (v NullableSigningKeys) IsSet() bool {
	return v.isSet
}

func (v *NullableSigningKeys) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSigningKeys(val *SigningKeys) *NullableSigningKeys {
	return &NullableSigningKeys{value: val, isSet: true}
}

func (v NullableSigningKeys) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSigningKeys) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
