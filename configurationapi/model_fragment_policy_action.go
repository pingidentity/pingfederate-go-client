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

// checks if the FragmentPolicyAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FragmentPolicyAction{}

// FragmentPolicyAction struct for FragmentPolicyAction
type FragmentPolicyAction struct {
	// The authentication selection type.
	Type string `json:"type" tfsdk:"type"`
	// The result context.
	Context         *string           `json:"context,omitempty" tfsdk:"context"`
	AttributeRules  *AttributeRules   `json:"attributeRules,omitempty" tfsdk:"attribute_rules"`
	Fragment        ResourceLink      `json:"fragment" tfsdk:"fragment"`
	FragmentMapping *AttributeMapping `json:"fragmentMapping,omitempty" tfsdk:"fragment_mapping"`
}

type _FragmentPolicyAction FragmentPolicyAction

// NewFragmentPolicyAction instantiates a new FragmentPolicyAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFragmentPolicyAction(type_ string, fragment ResourceLink) *FragmentPolicyAction {
	this := FragmentPolicyAction{}
	this.Type = type_
	this.Fragment = fragment
	return &this
}

// NewFragmentPolicyActionWithDefaults instantiates a new FragmentPolicyAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFragmentPolicyActionWithDefaults() *FragmentPolicyAction {
	this := FragmentPolicyAction{}
	return &this
}

// GetType returns the Type field value
func (o *FragmentPolicyAction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FragmentPolicyAction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FragmentPolicyAction) SetType(v string) {
	o.Type = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *FragmentPolicyAction) GetContext() string {
	if o == nil || IsNil(o.Context) {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FragmentPolicyAction) GetContextOk() (*string, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *FragmentPolicyAction) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *FragmentPolicyAction) SetContext(v string) {
	o.Context = &v
}

// GetAttributeRules returns the AttributeRules field value if set, zero value otherwise.
func (o *FragmentPolicyAction) GetAttributeRules() AttributeRules {
	if o == nil || IsNil(o.AttributeRules) {
		var ret AttributeRules
		return ret
	}
	return *o.AttributeRules
}

// GetAttributeRulesOk returns a tuple with the AttributeRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FragmentPolicyAction) GetAttributeRulesOk() (*AttributeRules, bool) {
	if o == nil || IsNil(o.AttributeRules) {
		return nil, false
	}
	return o.AttributeRules, true
}

// HasAttributeRules returns a boolean if a field has been set.
func (o *FragmentPolicyAction) HasAttributeRules() bool {
	if o != nil && !IsNil(o.AttributeRules) {
		return true
	}

	return false
}

// SetAttributeRules gets a reference to the given AttributeRules and assigns it to the AttributeRules field.
func (o *FragmentPolicyAction) SetAttributeRules(v AttributeRules) {
	o.AttributeRules = &v
}

// GetFragment returns the Fragment field value
func (o *FragmentPolicyAction) GetFragment() ResourceLink {
	if o == nil {
		var ret ResourceLink
		return ret
	}

	return o.Fragment
}

// GetFragmentOk returns a tuple with the Fragment field value
// and a boolean to check if the value has been set.
func (o *FragmentPolicyAction) GetFragmentOk() (*ResourceLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fragment, true
}

// SetFragment sets field value
func (o *FragmentPolicyAction) SetFragment(v ResourceLink) {
	o.Fragment = v
}

// GetFragmentMapping returns the FragmentMapping field value if set, zero value otherwise.
func (o *FragmentPolicyAction) GetFragmentMapping() AttributeMapping {
	if o == nil || IsNil(o.FragmentMapping) {
		var ret AttributeMapping
		return ret
	}
	return *o.FragmentMapping
}

// GetFragmentMappingOk returns a tuple with the FragmentMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FragmentPolicyAction) GetFragmentMappingOk() (*AttributeMapping, bool) {
	if o == nil || IsNil(o.FragmentMapping) {
		return nil, false
	}
	return o.FragmentMapping, true
}

// HasFragmentMapping returns a boolean if a field has been set.
func (o *FragmentPolicyAction) HasFragmentMapping() bool {
	if o != nil && !IsNil(o.FragmentMapping) {
		return true
	}

	return false
}

// SetFragmentMapping gets a reference to the given AttributeMapping and assigns it to the FragmentMapping field.
func (o *FragmentPolicyAction) SetFragmentMapping(v AttributeMapping) {
	o.FragmentMapping = &v
}

func (o FragmentPolicyAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FragmentPolicyAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.AttributeRules) {
		toSerialize["attributeRules"] = o.AttributeRules
	}
	toSerialize["fragment"] = o.Fragment
	if !IsNil(o.FragmentMapping) {
		toSerialize["fragmentMapping"] = o.FragmentMapping
	}
	return toSerialize, nil
}

func (o *FragmentPolicyAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"fragment",
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

	varFragmentPolicyAction := _FragmentPolicyAction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varFragmentPolicyAction)

	if err != nil {
		return err
	}

	*o = FragmentPolicyAction(varFragmentPolicyAction)

	return err
}

type NullableFragmentPolicyAction struct {
	value *FragmentPolicyAction
	isSet bool
}

func (v NullableFragmentPolicyAction) Get() *FragmentPolicyAction {
	return v.value
}

func (v *NullableFragmentPolicyAction) Set(val *FragmentPolicyAction) {
	v.value = val
	v.isSet = true
}

func (v NullableFragmentPolicyAction) IsSet() bool {
	return v.isSet
}

func (v *NullableFragmentPolicyAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFragmentPolicyAction(val *FragmentPolicyAction) *NullableFragmentPolicyAction {
	return &NullableFragmentPolicyAction{value: val, isSet: true}
}

func (v NullableFragmentPolicyAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFragmentPolicyAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
