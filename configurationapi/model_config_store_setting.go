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

// checks if the ConfigStoreSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigStoreSetting{}

// ConfigStoreSetting Single configuration setting.
type ConfigStoreSetting struct {
	// The id of the configuration setting.
	Id string `json:"id" tfsdk:"id"`
	// The value of the configuration setting. This is used when the setting has a single string value.
	StringValue *string `json:"stringValue,omitempty" tfsdk:"string_value"`
	// The list of values for the configuration setting. This is used when the setting has a list of string values.
	ListValue []string `json:"listValue,omitempty" tfsdk:"list_value"`
	// The map of key/value pairs for the configuration setting. This is used when the setting has a map of string keys and values.
	MapValue *map[string]string `json:"mapValue,omitempty" tfsdk:"map_value"`
	// The type of configuration setting. This could be a single string, list of strings, or map of string keys and values.
	Type string `json:"type" tfsdk:"type"`
}

// NewConfigStoreSetting instantiates a new ConfigStoreSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigStoreSetting(id string, type_ string) *ConfigStoreSetting {
	this := ConfigStoreSetting{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewConfigStoreSettingWithDefaults instantiates a new ConfigStoreSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigStoreSettingWithDefaults() *ConfigStoreSetting {
	this := ConfigStoreSetting{}
	return &this
}

// GetId returns the Id field value
func (o *ConfigStoreSetting) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConfigStoreSetting) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConfigStoreSetting) SetId(v string) {
	o.Id = v
}

// GetStringValue returns the StringValue field value if set, zero value otherwise.
func (o *ConfigStoreSetting) GetStringValue() string {
	if o == nil || IsNil(o.StringValue) {
		var ret string
		return ret
	}
	return *o.StringValue
}

// GetStringValueOk returns a tuple with the StringValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigStoreSetting) GetStringValueOk() (*string, bool) {
	if o == nil || IsNil(o.StringValue) {
		return nil, false
	}
	return o.StringValue, true
}

// HasStringValue returns a boolean if a field has been set.
func (o *ConfigStoreSetting) HasStringValue() bool {
	if o != nil && !IsNil(o.StringValue) {
		return true
	}

	return false
}

// SetStringValue gets a reference to the given string and assigns it to the StringValue field.
func (o *ConfigStoreSetting) SetStringValue(v string) {
	o.StringValue = &v
}

// GetListValue returns the ListValue field value if set, zero value otherwise.
func (o *ConfigStoreSetting) GetListValue() []string {
	if o == nil || IsNil(o.ListValue) {
		var ret []string
		return ret
	}
	return o.ListValue
}

// GetListValueOk returns a tuple with the ListValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigStoreSetting) GetListValueOk() ([]string, bool) {
	if o == nil || IsNil(o.ListValue) {
		return nil, false
	}
	return o.ListValue, true
}

// HasListValue returns a boolean if a field has been set.
func (o *ConfigStoreSetting) HasListValue() bool {
	if o != nil && !IsNil(o.ListValue) {
		return true
	}

	return false
}

// SetListValue gets a reference to the given []string and assigns it to the ListValue field.
func (o *ConfigStoreSetting) SetListValue(v []string) {
	o.ListValue = v
}

// GetMapValue returns the MapValue field value if set, zero value otherwise.
func (o *ConfigStoreSetting) GetMapValue() map[string]string {
	if o == nil || IsNil(o.MapValue) {
		var ret map[string]string
		return ret
	}
	return *o.MapValue
}

// GetMapValueOk returns a tuple with the MapValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigStoreSetting) GetMapValueOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.MapValue) {
		return nil, false
	}
	return o.MapValue, true
}

// HasMapValue returns a boolean if a field has been set.
func (o *ConfigStoreSetting) HasMapValue() bool {
	if o != nil && !IsNil(o.MapValue) {
		return true
	}

	return false
}

// SetMapValue gets a reference to the given map[string]string and assigns it to the MapValue field.
func (o *ConfigStoreSetting) SetMapValue(v map[string]string) {
	o.MapValue = &v
}

// GetType returns the Type field value
func (o *ConfigStoreSetting) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ConfigStoreSetting) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ConfigStoreSetting) SetType(v string) {
	o.Type = v
}

func (o ConfigStoreSetting) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigStoreSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.StringValue) {
		toSerialize["stringValue"] = o.StringValue
	}
	if !IsNil(o.ListValue) {
		toSerialize["listValue"] = o.ListValue
	}
	if !IsNil(o.MapValue) {
		toSerialize["mapValue"] = o.MapValue
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableConfigStoreSetting struct {
	value *ConfigStoreSetting
	isSet bool
}

func (v NullableConfigStoreSetting) Get() *ConfigStoreSetting {
	return v.value
}

func (v *NullableConfigStoreSetting) Set(val *ConfigStoreSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigStoreSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigStoreSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigStoreSetting(val *ConfigStoreSetting) *NullableConfigStoreSetting {
	return &NullableConfigStoreSetting{value: val, isSet: true}
}

func (v NullableConfigStoreSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigStoreSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
