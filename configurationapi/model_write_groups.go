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

// checks if the WriteGroups type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WriteGroups{}

// WriteGroups Group creation configuration.
type WriteGroups struct {
	// A list of user repository mappings from attribute names to their fulfillment values.
	AttributeFulfillment map[string]AttributeFulfillmentValue `json:"attributeFulfillment" tfsdk:"attribute_fulfillment"`
}

type _WriteGroups WriteGroups

// NewWriteGroups instantiates a new WriteGroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWriteGroups(attributeFulfillment map[string]AttributeFulfillmentValue) *WriteGroups {
	this := WriteGroups{}
	this.AttributeFulfillment = attributeFulfillment
	return &this
}

// NewWriteGroupsWithDefaults instantiates a new WriteGroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteGroupsWithDefaults() *WriteGroups {
	this := WriteGroups{}
	return &this
}

// GetAttributeFulfillment returns the AttributeFulfillment field value
func (o *WriteGroups) GetAttributeFulfillment() map[string]AttributeFulfillmentValue {
	if o == nil {
		var ret map[string]AttributeFulfillmentValue
		return ret
	}

	return o.AttributeFulfillment
}

// GetAttributeFulfillmentOk returns a tuple with the AttributeFulfillment field value
// and a boolean to check if the value has been set.
func (o *WriteGroups) GetAttributeFulfillmentOk() (*map[string]AttributeFulfillmentValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeFulfillment, true
}

// SetAttributeFulfillment sets field value
func (o *WriteGroups) SetAttributeFulfillment(v map[string]AttributeFulfillmentValue) {
	o.AttributeFulfillment = v
}

func (o WriteGroups) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WriteGroups) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attributeFulfillment"] = o.AttributeFulfillment
	return toSerialize, nil
}

func (o *WriteGroups) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attributeFulfillment",
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

	varWriteGroups := _WriteGroups{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varWriteGroups)

	if err != nil {
		return err
	}

	*o = WriteGroups(varWriteGroups)

	return err
}

type NullableWriteGroups struct {
	value *WriteGroups
	isSet bool
}

func (v NullableWriteGroups) Get() *WriteGroups {
	return v.value
}

func (v *NullableWriteGroups) Set(val *WriteGroups) {
	v.value = val
	v.isSet = true
}

func (v NullableWriteGroups) IsSet() bool {
	return v.isSet
}

func (v *NullableWriteGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWriteGroups(val *WriteGroups) *NullableWriteGroups {
	return &NullableWriteGroups{value: val, isSet: true}
}

func (v NullableWriteGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWriteGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
