/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
	"time"
)

// checks if the P14EKeyPairView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &P14EKeyPairView{}

// P14EKeyPairView PingOne for Enterprise connection key pair details.
type P14EKeyPairView struct {
	// Indicates whether this is the current key used to authenticate with PingOne.
	CurrentAuthenticationKey *bool `json:"currentAuthenticationKey,omitempty" tfsdk:"current_authentication_key"`
	// Indicates whether this is the previous key used to authenticate with PingOne.
	PreviousAuthenticationKey *bool     `json:"previousAuthenticationKey,omitempty" tfsdk:"previous_authentication_key"`
	KeyPairView               *CertView `json:"keyPairView,omitempty" tfsdk:"key_pair_view"`
	// The creation time of the key.
	CreationTime *time.Time `json:"creationTime,omitempty" tfsdk:"creation_time"`
}

// NewP14EKeyPairView instantiates a new P14EKeyPairView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewP14EKeyPairView() *P14EKeyPairView {
	this := P14EKeyPairView{}
	return &this
}

// NewP14EKeyPairViewWithDefaults instantiates a new P14EKeyPairView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewP14EKeyPairViewWithDefaults() *P14EKeyPairView {
	this := P14EKeyPairView{}
	return &this
}

// GetCurrentAuthenticationKey returns the CurrentAuthenticationKey field value if set, zero value otherwise.
func (o *P14EKeyPairView) GetCurrentAuthenticationKey() bool {
	if o == nil || IsNil(o.CurrentAuthenticationKey) {
		var ret bool
		return ret
	}
	return *o.CurrentAuthenticationKey
}

// GetCurrentAuthenticationKeyOk returns a tuple with the CurrentAuthenticationKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *P14EKeyPairView) GetCurrentAuthenticationKeyOk() (*bool, bool) {
	if o == nil || IsNil(o.CurrentAuthenticationKey) {
		return nil, false
	}
	return o.CurrentAuthenticationKey, true
}

// HasCurrentAuthenticationKey returns a boolean if a field has been set.
func (o *P14EKeyPairView) HasCurrentAuthenticationKey() bool {
	if o != nil && !IsNil(o.CurrentAuthenticationKey) {
		return true
	}

	return false
}

// SetCurrentAuthenticationKey gets a reference to the given bool and assigns it to the CurrentAuthenticationKey field.
func (o *P14EKeyPairView) SetCurrentAuthenticationKey(v bool) {
	o.CurrentAuthenticationKey = &v
}

// GetPreviousAuthenticationKey returns the PreviousAuthenticationKey field value if set, zero value otherwise.
func (o *P14EKeyPairView) GetPreviousAuthenticationKey() bool {
	if o == nil || IsNil(o.PreviousAuthenticationKey) {
		var ret bool
		return ret
	}
	return *o.PreviousAuthenticationKey
}

// GetPreviousAuthenticationKeyOk returns a tuple with the PreviousAuthenticationKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *P14EKeyPairView) GetPreviousAuthenticationKeyOk() (*bool, bool) {
	if o == nil || IsNil(o.PreviousAuthenticationKey) {
		return nil, false
	}
	return o.PreviousAuthenticationKey, true
}

// HasPreviousAuthenticationKey returns a boolean if a field has been set.
func (o *P14EKeyPairView) HasPreviousAuthenticationKey() bool {
	if o != nil && !IsNil(o.PreviousAuthenticationKey) {
		return true
	}

	return false
}

// SetPreviousAuthenticationKey gets a reference to the given bool and assigns it to the PreviousAuthenticationKey field.
func (o *P14EKeyPairView) SetPreviousAuthenticationKey(v bool) {
	o.PreviousAuthenticationKey = &v
}

// GetKeyPairView returns the KeyPairView field value if set, zero value otherwise.
func (o *P14EKeyPairView) GetKeyPairView() CertView {
	if o == nil || IsNil(o.KeyPairView) {
		var ret CertView
		return ret
	}
	return *o.KeyPairView
}

// GetKeyPairViewOk returns a tuple with the KeyPairView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *P14EKeyPairView) GetKeyPairViewOk() (*CertView, bool) {
	if o == nil || IsNil(o.KeyPairView) {
		return nil, false
	}
	return o.KeyPairView, true
}

// HasKeyPairView returns a boolean if a field has been set.
func (o *P14EKeyPairView) HasKeyPairView() bool {
	if o != nil && !IsNil(o.KeyPairView) {
		return true
	}

	return false
}

// SetKeyPairView gets a reference to the given CertView and assigns it to the KeyPairView field.
func (o *P14EKeyPairView) SetKeyPairView(v CertView) {
	o.KeyPairView = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *P14EKeyPairView) GetCreationTime() time.Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *P14EKeyPairView) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *P14EKeyPairView) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *P14EKeyPairView) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

func (o P14EKeyPairView) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o P14EKeyPairView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrentAuthenticationKey) {
		toSerialize["currentAuthenticationKey"] = o.CurrentAuthenticationKey
	}
	if !IsNil(o.PreviousAuthenticationKey) {
		toSerialize["previousAuthenticationKey"] = o.PreviousAuthenticationKey
	}
	if !IsNil(o.KeyPairView) {
		toSerialize["keyPairView"] = o.KeyPairView
	}
	if !IsNil(o.CreationTime) {
		toSerialize["creationTime"] = o.CreationTime
	}
	return toSerialize, nil
}

type NullableP14EKeyPairView struct {
	value *P14EKeyPairView
	isSet bool
}

func (v NullableP14EKeyPairView) Get() *P14EKeyPairView {
	return v.value
}

func (v *NullableP14EKeyPairView) Set(val *P14EKeyPairView) {
	v.value = val
	v.isSet = true
}

func (v NullableP14EKeyPairView) IsSet() bool {
	return v.isSet
}

func (v *NullableP14EKeyPairView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableP14EKeyPairView(val *P14EKeyPairView) *NullableP14EKeyPairView {
	return &NullableP14EKeyPairView{value: val, isSet: true}
}

func (v NullableP14EKeyPairView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableP14EKeyPairView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
