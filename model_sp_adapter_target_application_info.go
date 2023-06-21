/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the SpAdapterTargetApplicationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpAdapterTargetApplicationInfo{}

// SpAdapterTargetApplicationInfo Target Application Information exposed by an SP adapter.
type SpAdapterTargetApplicationInfo struct {
	// The application name.
	ApplicationName *string `json:"applicationName,omitempty"`
	// The application icon URL.
	ApplicationIconUrl *string `json:"applicationIconUrl,omitempty"`
	// Specifies Whether target application information is inherited from its parent instance. If true, the rest of the properties in this model become read-only. The default value is false.
	Inherited *bool `json:"inherited,omitempty"`
}

// NewSpAdapterTargetApplicationInfo instantiates a new SpAdapterTargetApplicationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpAdapterTargetApplicationInfo() *SpAdapterTargetApplicationInfo {
	this := SpAdapterTargetApplicationInfo{}
	return &this
}

// NewSpAdapterTargetApplicationInfoWithDefaults instantiates a new SpAdapterTargetApplicationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpAdapterTargetApplicationInfoWithDefaults() *SpAdapterTargetApplicationInfo {
	this := SpAdapterTargetApplicationInfo{}
	return &this
}

// GetApplicationName returns the ApplicationName field value if set, zero value otherwise.
func (o *SpAdapterTargetApplicationInfo) GetApplicationName() string {
	if o == nil || IsNil(o.ApplicationName) {
		var ret string
		return ret
	}
	return *o.ApplicationName
}

// GetApplicationNameOk returns a tuple with the ApplicationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpAdapterTargetApplicationInfo) GetApplicationNameOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationName) {
		return nil, false
	}
	return o.ApplicationName, true
}

// HasApplicationName returns a boolean if a field has been set.
func (o *SpAdapterTargetApplicationInfo) HasApplicationName() bool {
	if o != nil && !IsNil(o.ApplicationName) {
		return true
	}

	return false
}

// SetApplicationName gets a reference to the given string and assigns it to the ApplicationName field.
func (o *SpAdapterTargetApplicationInfo) SetApplicationName(v string) {
	o.ApplicationName = &v
}

// GetApplicationIconUrl returns the ApplicationIconUrl field value if set, zero value otherwise.
func (o *SpAdapterTargetApplicationInfo) GetApplicationIconUrl() string {
	if o == nil || IsNil(o.ApplicationIconUrl) {
		var ret string
		return ret
	}
	return *o.ApplicationIconUrl
}

// GetApplicationIconUrlOk returns a tuple with the ApplicationIconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpAdapterTargetApplicationInfo) GetApplicationIconUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationIconUrl) {
		return nil, false
	}
	return o.ApplicationIconUrl, true
}

// HasApplicationIconUrl returns a boolean if a field has been set.
func (o *SpAdapterTargetApplicationInfo) HasApplicationIconUrl() bool {
	if o != nil && !IsNil(o.ApplicationIconUrl) {
		return true
	}

	return false
}

// SetApplicationIconUrl gets a reference to the given string and assigns it to the ApplicationIconUrl field.
func (o *SpAdapterTargetApplicationInfo) SetApplicationIconUrl(v string) {
	o.ApplicationIconUrl = &v
}

// GetInherited returns the Inherited field value if set, zero value otherwise.
func (o *SpAdapterTargetApplicationInfo) GetInherited() bool {
	if o == nil || IsNil(o.Inherited) {
		var ret bool
		return ret
	}
	return *o.Inherited
}

// GetInheritedOk returns a tuple with the Inherited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpAdapterTargetApplicationInfo) GetInheritedOk() (*bool, bool) {
	if o == nil || IsNil(o.Inherited) {
		return nil, false
	}
	return o.Inherited, true
}

// HasInherited returns a boolean if a field has been set.
func (o *SpAdapterTargetApplicationInfo) HasInherited() bool {
	if o != nil && !IsNil(o.Inherited) {
		return true
	}

	return false
}

// SetInherited gets a reference to the given bool and assigns it to the Inherited field.
func (o *SpAdapterTargetApplicationInfo) SetInherited(v bool) {
	o.Inherited = &v
}

func (o SpAdapterTargetApplicationInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpAdapterTargetApplicationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationName) {
		toSerialize["applicationName"] = o.ApplicationName
	}
	if !IsNil(o.ApplicationIconUrl) {
		toSerialize["applicationIconUrl"] = o.ApplicationIconUrl
	}
	if !IsNil(o.Inherited) {
		toSerialize["inherited"] = o.Inherited
	}
	return toSerialize, nil
}

type NullableSpAdapterTargetApplicationInfo struct {
	value *SpAdapterTargetApplicationInfo
	isSet bool
}

func (v NullableSpAdapterTargetApplicationInfo) Get() *SpAdapterTargetApplicationInfo {
	return v.value
}

func (v *NullableSpAdapterTargetApplicationInfo) Set(val *SpAdapterTargetApplicationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSpAdapterTargetApplicationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSpAdapterTargetApplicationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpAdapterTargetApplicationInfo(val *SpAdapterTargetApplicationInfo) *NullableSpAdapterTargetApplicationInfo {
	return &NullableSpAdapterTargetApplicationInfo{value: val, isSet: true}
}

func (v NullableSpAdapterTargetApplicationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpAdapterTargetApplicationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
