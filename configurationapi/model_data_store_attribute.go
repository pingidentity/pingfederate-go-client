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

// checks if the DataStoreAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataStoreAttribute{}

// DataStoreAttribute The data store attribute.
type DataStoreAttribute struct {
	// The data store attribute type.
	Type string `json:"type" tfsdk:"type"`
	// The data store attribute name.
	Name string `json:"name" tfsdk:"name"`
	// The data store attribute metadata.
	Metadata *map[string]string `json:"metadata,omitempty" tfsdk:"metadata"`
}

type _DataStoreAttribute DataStoreAttribute

// NewDataStoreAttribute instantiates a new DataStoreAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataStoreAttribute(type_ string, name string) *DataStoreAttribute {
	this := DataStoreAttribute{}
	this.Type = type_
	this.Name = name
	return &this
}

// NewDataStoreAttributeWithDefaults instantiates a new DataStoreAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataStoreAttributeWithDefaults() *DataStoreAttribute {
	this := DataStoreAttribute{}
	return &this
}

// GetType returns the Type field value
func (o *DataStoreAttribute) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DataStoreAttribute) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DataStoreAttribute) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *DataStoreAttribute) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DataStoreAttribute) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DataStoreAttribute) SetName(v string) {
	o.Name = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *DataStoreAttribute) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataStoreAttribute) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *DataStoreAttribute) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *DataStoreAttribute) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

func (o DataStoreAttribute) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataStoreAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *DataStoreAttribute) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"name",
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

	varDataStoreAttribute := _DataStoreAttribute{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varDataStoreAttribute)

	if err != nil {
		return err
	}

	*o = DataStoreAttribute(varDataStoreAttribute)

	return err
}

type NullableDataStoreAttribute struct {
	value *DataStoreAttribute
	isSet bool
}

func (v NullableDataStoreAttribute) Get() *DataStoreAttribute {
	return v.value
}

func (v *NullableDataStoreAttribute) Set(val *DataStoreAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableDataStoreAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableDataStoreAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataStoreAttribute(val *DataStoreAttribute) *NullableDataStoreAttribute {
	return &NullableDataStoreAttribute{value: val, isSet: true}
}

func (v NullableDataStoreAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataStoreAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
