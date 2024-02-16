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

// checks if the SpWsTrustAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpWsTrustAttribute{}

// SpWsTrustAttribute An attribute for the Ws-Trust attribute contract.
type SpWsTrustAttribute struct {
	// The name of this attribute.
	Name string `json:"name" tfsdk:"name"`
	// The attribute namespace.  This is required when the Default Token Type is SAML2.0 or SAML1.1 or SAML1.1 for Office 365.
	Namespace string `json:"namespace" tfsdk:"namespace"`
}

type _SpWsTrustAttribute SpWsTrustAttribute

// NewSpWsTrustAttribute instantiates a new SpWsTrustAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpWsTrustAttribute(name string, namespace string) *SpWsTrustAttribute {
	this := SpWsTrustAttribute{}
	this.Name = name
	this.Namespace = namespace
	return &this
}

// NewSpWsTrustAttributeWithDefaults instantiates a new SpWsTrustAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpWsTrustAttributeWithDefaults() *SpWsTrustAttribute {
	this := SpWsTrustAttribute{}
	return &this
}

// GetName returns the Name field value
func (o *SpWsTrustAttribute) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SpWsTrustAttribute) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SpWsTrustAttribute) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value
func (o *SpWsTrustAttribute) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *SpWsTrustAttribute) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *SpWsTrustAttribute) SetNamespace(v string) {
	o.Namespace = v
}

func (o SpWsTrustAttribute) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpWsTrustAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["namespace"] = o.Namespace
	return toSerialize, nil
}

func (o *SpWsTrustAttribute) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"namespace",
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

	varSpWsTrustAttribute := _SpWsTrustAttribute{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varSpWsTrustAttribute)

	if err != nil {
		return err
	}

	*o = SpWsTrustAttribute(varSpWsTrustAttribute)

	return err
}

type NullableSpWsTrustAttribute struct {
	value *SpWsTrustAttribute
	isSet bool
}

func (v NullableSpWsTrustAttribute) Get() *SpWsTrustAttribute {
	return v.value
}

func (v *NullableSpWsTrustAttribute) Set(val *SpWsTrustAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableSpWsTrustAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableSpWsTrustAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpWsTrustAttribute(val *SpWsTrustAttribute) *NullableSpWsTrustAttribute {
	return &NullableSpWsTrustAttribute{value: val, isSet: true}
}

func (v NullableSpWsTrustAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpWsTrustAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
