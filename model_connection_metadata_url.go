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

// checks if the ConnectionMetadataUrl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionMetadataUrl{}

// ConnectionMetadataUrl Configuration settings to enable automatic reload of partner's metadata.
type ConnectionMetadataUrl struct {
	MetadataUrlRef ResourceLink `json:"metadataUrlRef" tfsdk:"metadata_url_ref"`
	// Specifies whether the metadata of the connection will be automatically reloaded. The default value is true.
	EnableAutoMetadataUpdate *bool `json:"enableAutoMetadataUpdate,omitempty" tfsdk:"enable_auto_metadata_update"`
}

// NewConnectionMetadataUrl instantiates a new ConnectionMetadataUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionMetadataUrl(metadataUrlRef ResourceLink) *ConnectionMetadataUrl {
	this := ConnectionMetadataUrl{}
	this.MetadataUrlRef = metadataUrlRef
	return &this
}

// NewConnectionMetadataUrlWithDefaults instantiates a new ConnectionMetadataUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionMetadataUrlWithDefaults() *ConnectionMetadataUrl {
	this := ConnectionMetadataUrl{}
	return &this
}

// GetMetadataUrlRef returns the MetadataUrlRef field value
func (o *ConnectionMetadataUrl) GetMetadataUrlRef() ResourceLink {
	if o == nil {
		var ret ResourceLink
		return ret
	}

	return o.MetadataUrlRef
}

// GetMetadataUrlRefOk returns a tuple with the MetadataUrlRef field value
// and a boolean to check if the value has been set.
func (o *ConnectionMetadataUrl) GetMetadataUrlRefOk() (*ResourceLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetadataUrlRef, true
}

// SetMetadataUrlRef sets field value
func (o *ConnectionMetadataUrl) SetMetadataUrlRef(v ResourceLink) {
	o.MetadataUrlRef = v
}

// GetEnableAutoMetadataUpdate returns the EnableAutoMetadataUpdate field value if set, zero value otherwise.
func (o *ConnectionMetadataUrl) GetEnableAutoMetadataUpdate() bool {
	if o == nil || IsNil(o.EnableAutoMetadataUpdate) {
		var ret bool
		return ret
	}
	return *o.EnableAutoMetadataUpdate
}

// GetEnableAutoMetadataUpdateOk returns a tuple with the EnableAutoMetadataUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionMetadataUrl) GetEnableAutoMetadataUpdateOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableAutoMetadataUpdate) {
		return nil, false
	}
	return o.EnableAutoMetadataUpdate, true
}

// HasEnableAutoMetadataUpdate returns a boolean if a field has been set.
func (o *ConnectionMetadataUrl) HasEnableAutoMetadataUpdate() bool {
	if o != nil && !IsNil(o.EnableAutoMetadataUpdate) {
		return true
	}

	return false
}

// SetEnableAutoMetadataUpdate gets a reference to the given bool and assigns it to the EnableAutoMetadataUpdate field.
func (o *ConnectionMetadataUrl) SetEnableAutoMetadataUpdate(v bool) {
	o.EnableAutoMetadataUpdate = &v
}

func (o ConnectionMetadataUrl) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionMetadataUrl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metadataUrlRef"] = o.MetadataUrlRef
	if !IsNil(o.EnableAutoMetadataUpdate) {
		toSerialize["enableAutoMetadataUpdate"] = o.EnableAutoMetadataUpdate
	}
	return toSerialize, nil
}

type NullableConnectionMetadataUrl struct {
	value *ConnectionMetadataUrl
	isSet bool
}

func (v NullableConnectionMetadataUrl) Get() *ConnectionMetadataUrl {
	return v.value
}

func (v *NullableConnectionMetadataUrl) Set(val *ConnectionMetadataUrl) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionMetadataUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionMetadataUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionMetadataUrl(val *ConnectionMetadataUrl) *NullableConnectionMetadataUrl {
	return &NullableConnectionMetadataUrl{value: val, isSet: true}
}

func (v NullableConnectionMetadataUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionMetadataUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
