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

// checks if the BulkConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkConfig{}

// BulkConfig Model describing a series of configuration operations.
type BulkConfig struct {
	Metadata BulkConfigMetadata `json:"metadata" tfsdk:"metadata"`
	// A list of configuration operations.
	Operations []ConfigOperation `json:"operations" tfsdk:"operations"`
}

// NewBulkConfig instantiates a new BulkConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkConfig(metadata BulkConfigMetadata, operations []ConfigOperation) *BulkConfig {
	this := BulkConfig{}
	this.Metadata = metadata
	this.Operations = operations
	return &this
}

// NewBulkConfigWithDefaults instantiates a new BulkConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkConfigWithDefaults() *BulkConfig {
	this := BulkConfig{}
	return &this
}

// GetMetadata returns the Metadata field value
func (o *BulkConfig) GetMetadata() BulkConfigMetadata {
	if o == nil {
		var ret BulkConfigMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *BulkConfig) GetMetadataOk() (*BulkConfigMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *BulkConfig) SetMetadata(v BulkConfigMetadata) {
	o.Metadata = v
}

// GetOperations returns the Operations field value
func (o *BulkConfig) GetOperations() []ConfigOperation {
	if o == nil {
		var ret []ConfigOperation
		return ret
	}

	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value
// and a boolean to check if the value has been set.
func (o *BulkConfig) GetOperationsOk() ([]ConfigOperation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operations, true
}

// SetOperations sets field value
func (o *BulkConfig) SetOperations(v []ConfigOperation) {
	o.Operations = v
}

func (o BulkConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metadata"] = o.Metadata
	toSerialize["operations"] = o.Operations
	return toSerialize, nil
}

type NullableBulkConfig struct {
	value *BulkConfig
	isSet bool
}

func (v NullableBulkConfig) Get() *BulkConfig {
	return v.value
}

func (v *NullableBulkConfig) Set(val *BulkConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkConfig(val *BulkConfig) *NullableBulkConfig {
	return &NullableBulkConfig{value: val, isSet: true}
}

func (v NullableBulkConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
