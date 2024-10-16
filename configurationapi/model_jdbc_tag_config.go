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

// checks if the JdbcTagConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JdbcTagConfig{}

// JdbcTagConfig A JDBC data store's connection URLs and tags configuration. For regional deployments, provide a separate JdbcTagConfig for each region, containing the region-specific connection URL and the corresponding tags.
type JdbcTagConfig struct {
	// The location of the JDBC database.
	ConnectionUrl string `json:"connectionUrl" tfsdk:"connection_url"`
	// Tags associated with the connection URL. At runtime, nodes will use the first JdbcTagConfig that has a tag that matches with node.tags in run.properties.
	Tags *string `json:"tags,omitempty" tfsdk:"tags"`
	// Whether this is the default connection. Defaults to false if not specified.
	DefaultSource *bool `json:"defaultSource,omitempty" tfsdk:"default_source"`
}

// NewJdbcTagConfig instantiates a new JdbcTagConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJdbcTagConfig(connectionUrl string) *JdbcTagConfig {
	this := JdbcTagConfig{}
	this.ConnectionUrl = connectionUrl
	return &this
}

// NewJdbcTagConfigWithDefaults instantiates a new JdbcTagConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJdbcTagConfigWithDefaults() *JdbcTagConfig {
	this := JdbcTagConfig{}
	return &this
}

// GetConnectionUrl returns the ConnectionUrl field value
func (o *JdbcTagConfig) GetConnectionUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionUrl
}

// GetConnectionUrlOk returns a tuple with the ConnectionUrl field value
// and a boolean to check if the value has been set.
func (o *JdbcTagConfig) GetConnectionUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionUrl, true
}

// SetConnectionUrl sets field value
func (o *JdbcTagConfig) SetConnectionUrl(v string) {
	o.ConnectionUrl = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *JdbcTagConfig) GetTags() string {
	if o == nil || IsNil(o.Tags) {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcTagConfig) GetTagsOk() (*string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *JdbcTagConfig) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *JdbcTagConfig) SetTags(v string) {
	o.Tags = &v
}

// GetDefaultSource returns the DefaultSource field value if set, zero value otherwise.
func (o *JdbcTagConfig) GetDefaultSource() bool {
	if o == nil || IsNil(o.DefaultSource) {
		var ret bool
		return ret
	}
	return *o.DefaultSource
}

// GetDefaultSourceOk returns a tuple with the DefaultSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcTagConfig) GetDefaultSourceOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultSource) {
		return nil, false
	}
	return o.DefaultSource, true
}

// HasDefaultSource returns a boolean if a field has been set.
func (o *JdbcTagConfig) HasDefaultSource() bool {
	if o != nil && !IsNil(o.DefaultSource) {
		return true
	}

	return false
}

// SetDefaultSource gets a reference to the given bool and assigns it to the DefaultSource field.
func (o *JdbcTagConfig) SetDefaultSource(v bool) {
	o.DefaultSource = &v
}

func (o JdbcTagConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JdbcTagConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connectionUrl"] = o.ConnectionUrl
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.DefaultSource) {
		toSerialize["defaultSource"] = o.DefaultSource
	}
	return toSerialize, nil
}

type NullableJdbcTagConfig struct {
	value *JdbcTagConfig
	isSet bool
}

func (v NullableJdbcTagConfig) Get() *JdbcTagConfig {
	return v.value
}

func (v *NullableJdbcTagConfig) Set(val *JdbcTagConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableJdbcTagConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableJdbcTagConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJdbcTagConfig(val *JdbcTagConfig) *NullableJdbcTagConfig {
	return &NullableJdbcTagConfig{value: val, isSet: true}
}

func (v NullableJdbcTagConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJdbcTagConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
