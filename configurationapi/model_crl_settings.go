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

// checks if the CrlSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CrlSettings{}

// CrlSettings CRL settings.
type CrlSettings struct {
	// Treat non retrievable CRL as revoked. This setting defaults to disabled.
	TreatNonRetrievableCrlAsRevoked *bool `json:"treatNonRetrievableCrlAsRevoked,omitempty" tfsdk:"treat_non_retrievable_crl_as_revoked"`
	// Verify CRL signature. This setting defaults to enabled.
	VerifyCrlSignature *bool `json:"verifyCrlSignature,omitempty" tfsdk:"verify_crl_signature"`
	// Next retry on resolution failure in minutes. This value defaults to \"1440\".
	NextRetryMinsWhenResolveFailed *int64 `json:"nextRetryMinsWhenResolveFailed,omitempty" tfsdk:"next_retry_mins_when_resolve_failed"`
	// Next retry on next update expiration in minutes. This value defaults to \"60\".
	NextRetryMinsWhenNextUpdateInPast *int64 `json:"nextRetryMinsWhenNextUpdateInPast,omitempty" tfsdk:"next_retry_mins_when_next_update_in_past"`
}

// NewCrlSettings instantiates a new CrlSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCrlSettings() *CrlSettings {
	this := CrlSettings{}
	return &this
}

// NewCrlSettingsWithDefaults instantiates a new CrlSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCrlSettingsWithDefaults() *CrlSettings {
	this := CrlSettings{}
	return &this
}

// GetTreatNonRetrievableCrlAsRevoked returns the TreatNonRetrievableCrlAsRevoked field value if set, zero value otherwise.
func (o *CrlSettings) GetTreatNonRetrievableCrlAsRevoked() bool {
	if o == nil || IsNil(o.TreatNonRetrievableCrlAsRevoked) {
		var ret bool
		return ret
	}
	return *o.TreatNonRetrievableCrlAsRevoked
}

// GetTreatNonRetrievableCrlAsRevokedOk returns a tuple with the TreatNonRetrievableCrlAsRevoked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrlSettings) GetTreatNonRetrievableCrlAsRevokedOk() (*bool, bool) {
	if o == nil || IsNil(o.TreatNonRetrievableCrlAsRevoked) {
		return nil, false
	}
	return o.TreatNonRetrievableCrlAsRevoked, true
}

// HasTreatNonRetrievableCrlAsRevoked returns a boolean if a field has been set.
func (o *CrlSettings) HasTreatNonRetrievableCrlAsRevoked() bool {
	if o != nil && !IsNil(o.TreatNonRetrievableCrlAsRevoked) {
		return true
	}

	return false
}

// SetTreatNonRetrievableCrlAsRevoked gets a reference to the given bool and assigns it to the TreatNonRetrievableCrlAsRevoked field.
func (o *CrlSettings) SetTreatNonRetrievableCrlAsRevoked(v bool) {
	o.TreatNonRetrievableCrlAsRevoked = &v
}

// GetVerifyCrlSignature returns the VerifyCrlSignature field value if set, zero value otherwise.
func (o *CrlSettings) GetVerifyCrlSignature() bool {
	if o == nil || IsNil(o.VerifyCrlSignature) {
		var ret bool
		return ret
	}
	return *o.VerifyCrlSignature
}

// GetVerifyCrlSignatureOk returns a tuple with the VerifyCrlSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrlSettings) GetVerifyCrlSignatureOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifyCrlSignature) {
		return nil, false
	}
	return o.VerifyCrlSignature, true
}

// HasVerifyCrlSignature returns a boolean if a field has been set.
func (o *CrlSettings) HasVerifyCrlSignature() bool {
	if o != nil && !IsNil(o.VerifyCrlSignature) {
		return true
	}

	return false
}

// SetVerifyCrlSignature gets a reference to the given bool and assigns it to the VerifyCrlSignature field.
func (o *CrlSettings) SetVerifyCrlSignature(v bool) {
	o.VerifyCrlSignature = &v
}

// GetNextRetryMinsWhenResolveFailed returns the NextRetryMinsWhenResolveFailed field value if set, zero value otherwise.
func (o *CrlSettings) GetNextRetryMinsWhenResolveFailed() int64 {
	if o == nil || IsNil(o.NextRetryMinsWhenResolveFailed) {
		var ret int64
		return ret
	}
	return *o.NextRetryMinsWhenResolveFailed
}

// GetNextRetryMinsWhenResolveFailedOk returns a tuple with the NextRetryMinsWhenResolveFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrlSettings) GetNextRetryMinsWhenResolveFailedOk() (*int64, bool) {
	if o == nil || IsNil(o.NextRetryMinsWhenResolveFailed) {
		return nil, false
	}
	return o.NextRetryMinsWhenResolveFailed, true
}

// HasNextRetryMinsWhenResolveFailed returns a boolean if a field has been set.
func (o *CrlSettings) HasNextRetryMinsWhenResolveFailed() bool {
	if o != nil && !IsNil(o.NextRetryMinsWhenResolveFailed) {
		return true
	}

	return false
}

// SetNextRetryMinsWhenResolveFailed gets a reference to the given int64 and assigns it to the NextRetryMinsWhenResolveFailed field.
func (o *CrlSettings) SetNextRetryMinsWhenResolveFailed(v int64) {
	o.NextRetryMinsWhenResolveFailed = &v
}

// GetNextRetryMinsWhenNextUpdateInPast returns the NextRetryMinsWhenNextUpdateInPast field value if set, zero value otherwise.
func (o *CrlSettings) GetNextRetryMinsWhenNextUpdateInPast() int64 {
	if o == nil || IsNil(o.NextRetryMinsWhenNextUpdateInPast) {
		var ret int64
		return ret
	}
	return *o.NextRetryMinsWhenNextUpdateInPast
}

// GetNextRetryMinsWhenNextUpdateInPastOk returns a tuple with the NextRetryMinsWhenNextUpdateInPast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrlSettings) GetNextRetryMinsWhenNextUpdateInPastOk() (*int64, bool) {
	if o == nil || IsNil(o.NextRetryMinsWhenNextUpdateInPast) {
		return nil, false
	}
	return o.NextRetryMinsWhenNextUpdateInPast, true
}

// HasNextRetryMinsWhenNextUpdateInPast returns a boolean if a field has been set.
func (o *CrlSettings) HasNextRetryMinsWhenNextUpdateInPast() bool {
	if o != nil && !IsNil(o.NextRetryMinsWhenNextUpdateInPast) {
		return true
	}

	return false
}

// SetNextRetryMinsWhenNextUpdateInPast gets a reference to the given int64 and assigns it to the NextRetryMinsWhenNextUpdateInPast field.
func (o *CrlSettings) SetNextRetryMinsWhenNextUpdateInPast(v int64) {
	o.NextRetryMinsWhenNextUpdateInPast = &v
}

func (o CrlSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CrlSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TreatNonRetrievableCrlAsRevoked) {
		toSerialize["treatNonRetrievableCrlAsRevoked"] = o.TreatNonRetrievableCrlAsRevoked
	}
	if !IsNil(o.VerifyCrlSignature) {
		toSerialize["verifyCrlSignature"] = o.VerifyCrlSignature
	}
	if !IsNil(o.NextRetryMinsWhenResolveFailed) {
		toSerialize["nextRetryMinsWhenResolveFailed"] = o.NextRetryMinsWhenResolveFailed
	}
	if !IsNil(o.NextRetryMinsWhenNextUpdateInPast) {
		toSerialize["nextRetryMinsWhenNextUpdateInPast"] = o.NextRetryMinsWhenNextUpdateInPast
	}
	return toSerialize, nil
}

type NullableCrlSettings struct {
	value *CrlSettings
	isSet bool
}

func (v NullableCrlSettings) Get() *CrlSettings {
	return v.value
}

func (v *NullableCrlSettings) Set(val *CrlSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableCrlSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableCrlSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCrlSettings(val *CrlSettings) *NullableCrlSettings {
	return &NullableCrlSettings{value: val, isSet: true}
}

func (v NullableCrlSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCrlSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}