/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the LdapTagConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LdapTagConfig{}

// LdapTagConfig An LDAP data store's hostnames and tags configuration. For regional deployments, provide a separate LdapTagConfig for each region, containing region-specific hostnames and the corresponding tags.
type LdapTagConfig struct {
	// The LDAP host names. Failover can be configured by providing multiple host names.
	Hostnames []string `json:"hostnames" tfsdk:"hostnames"`
	// Tags associated with the host names. At runtime, nodes will use the first LdapTagConfig that has a tag that matches with node.tags in run.properties.
	Tags *string `json:"tags,omitempty" tfsdk:"tags"`
	// Whether this is the default connection. Defaults to false if not specified.
	DefaultSource *bool `json:"defaultSource,omitempty" tfsdk:"default_source"`
}

type _LdapTagConfig LdapTagConfig

// NewLdapTagConfig instantiates a new LdapTagConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapTagConfig(hostnames []string) *LdapTagConfig {
	this := LdapTagConfig{}
	this.Hostnames = hostnames
	return &this
}

// NewLdapTagConfigWithDefaults instantiates a new LdapTagConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapTagConfigWithDefaults() *LdapTagConfig {
	this := LdapTagConfig{}
	return &this
}

// GetHostnames returns the Hostnames field value
func (o *LdapTagConfig) GetHostnames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Hostnames
}

// GetHostnamesOk returns a tuple with the Hostnames field value
// and a boolean to check if the value has been set.
func (o *LdapTagConfig) GetHostnamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hostnames, true
}

// SetHostnames sets field value
func (o *LdapTagConfig) SetHostnames(v []string) {
	o.Hostnames = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *LdapTagConfig) GetTags() string {
	if o == nil || IsNil(o.Tags) {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapTagConfig) GetTagsOk() (*string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *LdapTagConfig) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *LdapTagConfig) SetTags(v string) {
	o.Tags = &v
}

// GetDefaultSource returns the DefaultSource field value if set, zero value otherwise.
func (o *LdapTagConfig) GetDefaultSource() bool {
	if o == nil || IsNil(o.DefaultSource) {
		var ret bool
		return ret
	}
	return *o.DefaultSource
}

// GetDefaultSourceOk returns a tuple with the DefaultSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapTagConfig) GetDefaultSourceOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultSource) {
		return nil, false
	}
	return o.DefaultSource, true
}

// HasDefaultSource returns a boolean if a field has been set.
func (o *LdapTagConfig) HasDefaultSource() bool {
	if o != nil && !IsNil(o.DefaultSource) {
		return true
	}

	return false
}

// SetDefaultSource gets a reference to the given bool and assigns it to the DefaultSource field.
func (o *LdapTagConfig) SetDefaultSource(v bool) {
	o.DefaultSource = &v
}

func (o LdapTagConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LdapTagConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hostnames"] = o.Hostnames
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.DefaultSource) {
		toSerialize["defaultSource"] = o.DefaultSource
	}
	return toSerialize, nil
}

func (o *LdapTagConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"hostnames",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varLdapTagConfig := _LdapTagConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varLdapTagConfig)

	if err != nil {
		return err
	}

	*o = LdapTagConfig(varLdapTagConfig)

	return err
}

type NullableLdapTagConfig struct {
	value *LdapTagConfig
	isSet bool
}

func (v NullableLdapTagConfig) Get() *LdapTagConfig {
	return v.value
}

func (v *NullableLdapTagConfig) Set(val *LdapTagConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapTagConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapTagConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapTagConfig(val *LdapTagConfig) *NullableLdapTagConfig {
	return &NullableLdapTagConfig{value: val, isSet: true}
}

func (v NullableLdapTagConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapTagConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
