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

// checks if the LicenseFeatureView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LicenseFeatureView{}

// LicenseFeatureView PingFederate license feature details.
type LicenseFeatureView struct {
	// The name of the license feature.
	Name *string `json:"name,omitempty" tfsdk:"name"`
	// The value of the license feature.
	Value *string `json:"value,omitempty" tfsdk:"value"`
}

// NewLicenseFeatureView instantiates a new LicenseFeatureView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseFeatureView() *LicenseFeatureView {
	this := LicenseFeatureView{}
	return &this
}

// NewLicenseFeatureViewWithDefaults instantiates a new LicenseFeatureView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseFeatureViewWithDefaults() *LicenseFeatureView {
	this := LicenseFeatureView{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LicenseFeatureView) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseFeatureView) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LicenseFeatureView) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LicenseFeatureView) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *LicenseFeatureView) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseFeatureView) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *LicenseFeatureView) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *LicenseFeatureView) SetValue(v string) {
	o.Value = &v
}

func (o LicenseFeatureView) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LicenseFeatureView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableLicenseFeatureView struct {
	value *LicenseFeatureView
	isSet bool
}

func (v NullableLicenseFeatureView) Get() *LicenseFeatureView {
	return v.value
}

func (v *NullableLicenseFeatureView) Set(val *LicenseFeatureView) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseFeatureView) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseFeatureView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseFeatureView(val *LicenseFeatureView) *NullableLicenseFeatureView {
	return &NullableLicenseFeatureView{value: val, isSet: true}
}

func (v NullableLicenseFeatureView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseFeatureView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
