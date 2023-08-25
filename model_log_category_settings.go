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

// checks if the LogCategorySettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogCategorySettings{}

// LogCategorySettings The settings for a log category. A log category represents a group of related loggers and is associated with a system property (the ID field of the category with 'pf.log.level.' prepended). The system property can be referenced in log4j2.xml to set the level for the associated loggers. Log category IDs, names, and descriptions are defined in log4j-categories.xml. Only the enabled state of the category can be modified through the administrative console or API.
type LogCategorySettings struct {
	// The ID of the log category. This field must match one of the category IDs defined in log4j-categories.xml.
	Id string `json:"id" tfsdk:"id"`
	// The name of the log category. This field is read-only and is ignored for PUT requests.
	Name *string `json:"name,omitempty" tfsdk:"name"`
	// The description of the log category. This field is read-only and is ignored for PUT requests.
	Description *string `json:"description,omitempty" tfsdk:"description"`
	// Determines whether or not the log category is enabled. The default is false.
	Enabled *bool `json:"enabled,omitempty" tfsdk:"enabled"`
}

// NewLogCategorySettings instantiates a new LogCategorySettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogCategorySettings(id string) *LogCategorySettings {
	this := LogCategorySettings{}
	this.Id = id
	return &this
}

// NewLogCategorySettingsWithDefaults instantiates a new LogCategorySettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogCategorySettingsWithDefaults() *LogCategorySettings {
	this := LogCategorySettings{}
	return &this
}

// GetId returns the Id field value
func (o *LogCategorySettings) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LogCategorySettings) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LogCategorySettings) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LogCategorySettings) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogCategorySettings) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LogCategorySettings) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LogCategorySettings) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LogCategorySettings) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogCategorySettings) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LogCategorySettings) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LogCategorySettings) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *LogCategorySettings) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogCategorySettings) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *LogCategorySettings) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *LogCategorySettings) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o LogCategorySettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogCategorySettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableLogCategorySettings struct {
	value *LogCategorySettings
	isSet bool
}

func (v NullableLogCategorySettings) Get() *LogCategorySettings {
	return v.value
}

func (v *NullableLogCategorySettings) Set(val *LogCategorySettings) {
	v.value = val
	v.isSet = true
}

func (v NullableLogCategorySettings) IsSet() bool {
	return v.isSet
}

func (v *NullableLogCategorySettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogCategorySettings(val *LogCategorySettings) *NullableLogCategorySettings {
	return &NullableLogCategorySettings{value: val, isSet: true}
}

func (v NullableLogCategorySettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogCategorySettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
