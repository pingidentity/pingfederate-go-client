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

// checks if the PluginDescriptor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PluginDescriptor{}

// PluginDescriptor Defines a plugin type, including available configuration parameters.
type PluginDescriptor struct {
	// Unique ID of the plugin.
	Id *string `json:"id,omitempty" tfsdk:"id"`
	// Friendly name for the plugin.
	Name *string `json:"name,omitempty" tfsdk:"name"`
	// Full class name of the class that implements this plugin.
	ClassName *string `json:"className,omitempty" tfsdk:"class_name"`
	// The attribute contract for this plugin.
	AttributeContract []string `json:"attributeContract,omitempty" tfsdk:"attribute_contract"`
	// Determines whether this plugin supports extending the attribute contract.
	SupportsExtendedContract *bool                   `json:"supportsExtendedContract,omitempty" tfsdk:"supports_extended_contract"`
	ConfigDescriptor         *PluginConfigDescriptor `json:"configDescriptor,omitempty" tfsdk:"config_descriptor"`
}

// NewPluginDescriptor instantiates a new PluginDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginDescriptor() *PluginDescriptor {
	this := PluginDescriptor{}
	return &this
}

// NewPluginDescriptorWithDefaults instantiates a new PluginDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginDescriptorWithDefaults() *PluginDescriptor {
	this := PluginDescriptor{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PluginDescriptor) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginDescriptor) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PluginDescriptor) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PluginDescriptor) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PluginDescriptor) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginDescriptor) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PluginDescriptor) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PluginDescriptor) SetName(v string) {
	o.Name = &v
}

// GetClassName returns the ClassName field value if set, zero value otherwise.
func (o *PluginDescriptor) GetClassName() string {
	if o == nil || IsNil(o.ClassName) {
		var ret string
		return ret
	}
	return *o.ClassName
}

// GetClassNameOk returns a tuple with the ClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginDescriptor) GetClassNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClassName) {
		return nil, false
	}
	return o.ClassName, true
}

// HasClassName returns a boolean if a field has been set.
func (o *PluginDescriptor) HasClassName() bool {
	if o != nil && !IsNil(o.ClassName) {
		return true
	}

	return false
}

// SetClassName gets a reference to the given string and assigns it to the ClassName field.
func (o *PluginDescriptor) SetClassName(v string) {
	o.ClassName = &v
}

// GetAttributeContract returns the AttributeContract field value if set, zero value otherwise.
func (o *PluginDescriptor) GetAttributeContract() []string {
	if o == nil || IsNil(o.AttributeContract) {
		var ret []string
		return ret
	}
	return o.AttributeContract
}

// GetAttributeContractOk returns a tuple with the AttributeContract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginDescriptor) GetAttributeContractOk() ([]string, bool) {
	if o == nil || IsNil(o.AttributeContract) {
		return nil, false
	}
	return o.AttributeContract, true
}

// HasAttributeContract returns a boolean if a field has been set.
func (o *PluginDescriptor) HasAttributeContract() bool {
	if o != nil && !IsNil(o.AttributeContract) {
		return true
	}

	return false
}

// SetAttributeContract gets a reference to the given []string and assigns it to the AttributeContract field.
func (o *PluginDescriptor) SetAttributeContract(v []string) {
	o.AttributeContract = v
}

// GetSupportsExtendedContract returns the SupportsExtendedContract field value if set, zero value otherwise.
func (o *PluginDescriptor) GetSupportsExtendedContract() bool {
	if o == nil || IsNil(o.SupportsExtendedContract) {
		var ret bool
		return ret
	}
	return *o.SupportsExtendedContract
}

// GetSupportsExtendedContractOk returns a tuple with the SupportsExtendedContract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginDescriptor) GetSupportsExtendedContractOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsExtendedContract) {
		return nil, false
	}
	return o.SupportsExtendedContract, true
}

// HasSupportsExtendedContract returns a boolean if a field has been set.
func (o *PluginDescriptor) HasSupportsExtendedContract() bool {
	if o != nil && !IsNil(o.SupportsExtendedContract) {
		return true
	}

	return false
}

// SetSupportsExtendedContract gets a reference to the given bool and assigns it to the SupportsExtendedContract field.
func (o *PluginDescriptor) SetSupportsExtendedContract(v bool) {
	o.SupportsExtendedContract = &v
}

// GetConfigDescriptor returns the ConfigDescriptor field value if set, zero value otherwise.
func (o *PluginDescriptor) GetConfigDescriptor() PluginConfigDescriptor {
	if o == nil || IsNil(o.ConfigDescriptor) {
		var ret PluginConfigDescriptor
		return ret
	}
	return *o.ConfigDescriptor
}

// GetConfigDescriptorOk returns a tuple with the ConfigDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginDescriptor) GetConfigDescriptorOk() (*PluginConfigDescriptor, bool) {
	if o == nil || IsNil(o.ConfigDescriptor) {
		return nil, false
	}
	return o.ConfigDescriptor, true
}

// HasConfigDescriptor returns a boolean if a field has been set.
func (o *PluginDescriptor) HasConfigDescriptor() bool {
	if o != nil && !IsNil(o.ConfigDescriptor) {
		return true
	}

	return false
}

// SetConfigDescriptor gets a reference to the given PluginConfigDescriptor and assigns it to the ConfigDescriptor field.
func (o *PluginDescriptor) SetConfigDescriptor(v PluginConfigDescriptor) {
	o.ConfigDescriptor = &v
}

func (o PluginDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PluginDescriptor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ClassName) {
		toSerialize["className"] = o.ClassName
	}
	if !IsNil(o.AttributeContract) {
		toSerialize["attributeContract"] = o.AttributeContract
	}
	if !IsNil(o.SupportsExtendedContract) {
		toSerialize["supportsExtendedContract"] = o.SupportsExtendedContract
	}
	if !IsNil(o.ConfigDescriptor) {
		toSerialize["configDescriptor"] = o.ConfigDescriptor
	}
	return toSerialize, nil
}

type NullablePluginDescriptor struct {
	value *PluginDescriptor
	isSet bool
}

func (v NullablePluginDescriptor) Get() *PluginDescriptor {
	return v.value
}

func (v *NullablePluginDescriptor) Set(val *PluginDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginDescriptor(val *PluginDescriptor) *NullablePluginDescriptor {
	return &NullablePluginDescriptor{value: val, isSet: true}
}

func (v NullablePluginDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
