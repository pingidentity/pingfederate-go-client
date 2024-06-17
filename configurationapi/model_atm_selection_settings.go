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

// checks if the AtmSelectionSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AtmSelectionSettings{}

// AtmSelectionSettings Selection settings for an access token management plugin instance.
type AtmSelectionSettings struct {
	// The list of base resource URI's which map to this token manager. A resource URI, specified via the 'aud' parameter, can be used to select a specific token manager for an OAuth request.
	ResourceUris []string `json:"resourceUris,omitempty" tfsdk:"resource_uris"`
}

// NewAtmSelectionSettings instantiates a new AtmSelectionSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAtmSelectionSettings() *AtmSelectionSettings {
	this := AtmSelectionSettings{}
	return &this
}

// NewAtmSelectionSettingsWithDefaults instantiates a new AtmSelectionSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAtmSelectionSettingsWithDefaults() *AtmSelectionSettings {
	this := AtmSelectionSettings{}
	return &this
}

// GetResourceUris returns the ResourceUris field value if set, zero value otherwise.
func (o *AtmSelectionSettings) GetResourceUris() []string {
	if o == nil || IsNil(o.ResourceUris) {
		var ret []string
		return ret
	}
	return o.ResourceUris
}

// GetResourceUrisOk returns a tuple with the ResourceUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtmSelectionSettings) GetResourceUrisOk() ([]string, bool) {
	if o == nil || IsNil(o.ResourceUris) {
		return nil, false
	}
	return o.ResourceUris, true
}

// HasResourceUris returns a boolean if a field has been set.
func (o *AtmSelectionSettings) HasResourceUris() bool {
	if o != nil && !IsNil(o.ResourceUris) {
		return true
	}

	return false
}

// SetResourceUris gets a reference to the given []string and assigns it to the ResourceUris field.
func (o *AtmSelectionSettings) SetResourceUris(v []string) {
	o.ResourceUris = v
}

func (o AtmSelectionSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AtmSelectionSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResourceUris) {
		toSerialize["resourceUris"] = o.ResourceUris
	}
	return toSerialize, nil
}

type NullableAtmSelectionSettings struct {
	value *AtmSelectionSettings
	isSet bool
}

func (v NullableAtmSelectionSettings) Get() *AtmSelectionSettings {
	return v.value
}

func (v *NullableAtmSelectionSettings) Set(val *AtmSelectionSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAtmSelectionSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAtmSelectionSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAtmSelectionSettings(val *AtmSelectionSettings) *NullableAtmSelectionSettings {
	return &NullableAtmSelectionSettings{value: val, isSet: true}
}

func (v NullableAtmSelectionSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAtmSelectionSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
