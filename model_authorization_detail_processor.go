/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the AuthorizationDetailProcessor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizationDetailProcessor{}

// AuthorizationDetailProcessor struct for AuthorizationDetailProcessor
type AuthorizationDetailProcessor struct {
	// The ID of the plugin instance. The ID cannot be modified once the instance is created.<br>Note: Ignored when specifying a connection's adapter override.
	Id string `json:"id"`
	// The plugin instance name. The name can be modified once the instance is created.<br>Note: Ignored when specifying a connection's adapter override.
	Name                string              `json:"name"`
	PluginDescriptorRef ResourceLink        `json:"pluginDescriptorRef"`
	ParentRef           *ResourceLink       `json:"parentRef,omitempty"`
	Configuration       PluginConfiguration `json:"configuration"`
}

// NewAuthorizationDetailProcessor instantiates a new AuthorizationDetailProcessor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationDetailProcessor(id string, name string, pluginDescriptorRef ResourceLink, configuration PluginConfiguration) *AuthorizationDetailProcessor {
	this := AuthorizationDetailProcessor{}
	this.Id = id
	this.Name = name
	this.PluginDescriptorRef = pluginDescriptorRef
	this.Configuration = configuration
	return &this
}

// NewAuthorizationDetailProcessorWithDefaults instantiates a new AuthorizationDetailProcessor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationDetailProcessorWithDefaults() *AuthorizationDetailProcessor {
	this := AuthorizationDetailProcessor{}
	return &this
}

// GetId returns the Id field value
func (o *AuthorizationDetailProcessor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AuthorizationDetailProcessor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AuthorizationDetailProcessor) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AuthorizationDetailProcessor) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthorizationDetailProcessor) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthorizationDetailProcessor) SetName(v string) {
	o.Name = v
}

// GetPluginDescriptorRef returns the PluginDescriptorRef field value
func (o *AuthorizationDetailProcessor) GetPluginDescriptorRef() ResourceLink {
	if o == nil {
		var ret ResourceLink
		return ret
	}

	return o.PluginDescriptorRef
}

// GetPluginDescriptorRefOk returns a tuple with the PluginDescriptorRef field value
// and a boolean to check if the value has been set.
func (o *AuthorizationDetailProcessor) GetPluginDescriptorRefOk() (*ResourceLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PluginDescriptorRef, true
}

// SetPluginDescriptorRef sets field value
func (o *AuthorizationDetailProcessor) SetPluginDescriptorRef(v ResourceLink) {
	o.PluginDescriptorRef = v
}

// GetParentRef returns the ParentRef field value if set, zero value otherwise.
func (o *AuthorizationDetailProcessor) GetParentRef() ResourceLink {
	if o == nil || IsNil(o.ParentRef) {
		var ret ResourceLink
		return ret
	}
	return *o.ParentRef
}

// GetParentRefOk returns a tuple with the ParentRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationDetailProcessor) GetParentRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.ParentRef) {
		return nil, false
	}
	return o.ParentRef, true
}

// HasParentRef returns a boolean if a field has been set.
func (o *AuthorizationDetailProcessor) HasParentRef() bool {
	if o != nil && !IsNil(o.ParentRef) {
		return true
	}

	return false
}

// SetParentRef gets a reference to the given ResourceLink and assigns it to the ParentRef field.
func (o *AuthorizationDetailProcessor) SetParentRef(v ResourceLink) {
	o.ParentRef = &v
}

// GetConfiguration returns the Configuration field value
func (o *AuthorizationDetailProcessor) GetConfiguration() PluginConfiguration {
	if o == nil {
		var ret PluginConfiguration
		return ret
	}

	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value
// and a boolean to check if the value has been set.
func (o *AuthorizationDetailProcessor) GetConfigurationOk() (*PluginConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Configuration, true
}

// SetConfiguration sets field value
func (o *AuthorizationDetailProcessor) SetConfiguration(v PluginConfiguration) {
	o.Configuration = v
}

func (o AuthorizationDetailProcessor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizationDetailProcessor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["pluginDescriptorRef"] = o.PluginDescriptorRef
	if !IsNil(o.ParentRef) {
		toSerialize["parentRef"] = o.ParentRef
	}
	toSerialize["configuration"] = o.Configuration
	return toSerialize, nil
}

type NullableAuthorizationDetailProcessor struct {
	value *AuthorizationDetailProcessor
	isSet bool
}

func (v NullableAuthorizationDetailProcessor) Get() *AuthorizationDetailProcessor {
	return v.value
}

func (v *NullableAuthorizationDetailProcessor) Set(val *AuthorizationDetailProcessor) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationDetailProcessor) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationDetailProcessor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationDetailProcessor(val *AuthorizationDetailProcessor) *NullableAuthorizationDetailProcessor {
	return &NullableAuthorizationDetailProcessor{value: val, isSet: true}
}

func (v NullableAuthorizationDetailProcessor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationDetailProcessor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
