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

// checks if the CustomDataStore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomDataStore{}

// CustomDataStore struct for CustomDataStore
type CustomDataStore struct {
	DataStore
	// The data store type.
	Type string `json:"type" tfsdk:"type"`
	// The persistent, unique ID for the data store. It can be any combination of [a-zA-Z0-9._-]. This property is system-assigned if not specified.
	Id *string `json:"id,omitempty" tfsdk:"id"`
	// Whether attribute values should be masked in the log.
	MaskAttributeValues *bool `json:"maskAttributeValues,omitempty" tfsdk:"mask_attribute_values"`
	// The plugin instance name.
	Name                string              `json:"name" tfsdk:"name"`
	PluginDescriptorRef ResourceLink        `json:"pluginDescriptorRef" tfsdk:"plugin_descriptor_ref"`
	ParentRef           *ResourceLink       `json:"parentRef,omitempty" tfsdk:"parent_ref"`
	Configuration       PluginConfiguration `json:"configuration" tfsdk:"configuration"`
}

// NewCustomDataStore instantiates a new CustomDataStore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomDataStore(type_ string, name string, pluginDescriptorRef ResourceLink, configuration PluginConfiguration) *CustomDataStore {
	this := CustomDataStore{}
	this.Type = type_
	this.Name = name
	this.PluginDescriptorRef = pluginDescriptorRef
	this.Configuration = configuration
	return &this
}

// NewCustomDataStoreWithDefaults instantiates a new CustomDataStore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomDataStoreWithDefaults() *CustomDataStore {
	this := CustomDataStore{}
	return &this
}

// GetType returns the Type field value
func (o *CustomDataStore) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomDataStore) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomDataStore) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomDataStore) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDataStore) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomDataStore) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomDataStore) SetId(v string) {
	o.Id = &v
}

// GetMaskAttributeValues returns the MaskAttributeValues field value if set, zero value otherwise.
func (o *CustomDataStore) GetMaskAttributeValues() bool {
	if o == nil || IsNil(o.MaskAttributeValues) {
		var ret bool
		return ret
	}
	return *o.MaskAttributeValues
}

// GetMaskAttributeValuesOk returns a tuple with the MaskAttributeValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDataStore) GetMaskAttributeValuesOk() (*bool, bool) {
	if o == nil || IsNil(o.MaskAttributeValues) {
		return nil, false
	}
	return o.MaskAttributeValues, true
}

// HasMaskAttributeValues returns a boolean if a field has been set.
func (o *CustomDataStore) HasMaskAttributeValues() bool {
	if o != nil && !IsNil(o.MaskAttributeValues) {
		return true
	}

	return false
}

// SetMaskAttributeValues gets a reference to the given bool and assigns it to the MaskAttributeValues field.
func (o *CustomDataStore) SetMaskAttributeValues(v bool) {
	o.MaskAttributeValues = &v
}

// GetName returns the Name field value
func (o *CustomDataStore) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomDataStore) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CustomDataStore) SetName(v string) {
	o.Name = v
}

// GetPluginDescriptorRef returns the PluginDescriptorRef field value
func (o *CustomDataStore) GetPluginDescriptorRef() ResourceLink {
	if o == nil {
		var ret ResourceLink
		return ret
	}

	return o.PluginDescriptorRef
}

// GetPluginDescriptorRefOk returns a tuple with the PluginDescriptorRef field value
// and a boolean to check if the value has been set.
func (o *CustomDataStore) GetPluginDescriptorRefOk() (*ResourceLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PluginDescriptorRef, true
}

// SetPluginDescriptorRef sets field value
func (o *CustomDataStore) SetPluginDescriptorRef(v ResourceLink) {
	o.PluginDescriptorRef = v
}

// GetParentRef returns the ParentRef field value if set, zero value otherwise.
func (o *CustomDataStore) GetParentRef() ResourceLink {
	if o == nil || IsNil(o.ParentRef) {
		var ret ResourceLink
		return ret
	}
	return *o.ParentRef
}

// GetParentRefOk returns a tuple with the ParentRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDataStore) GetParentRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.ParentRef) {
		return nil, false
	}
	return o.ParentRef, true
}

// HasParentRef returns a boolean if a field has been set.
func (o *CustomDataStore) HasParentRef() bool {
	if o != nil && !IsNil(o.ParentRef) {
		return true
	}

	return false
}

// SetParentRef gets a reference to the given ResourceLink and assigns it to the ParentRef field.
func (o *CustomDataStore) SetParentRef(v ResourceLink) {
	o.ParentRef = &v
}

// GetConfiguration returns the Configuration field value
func (o *CustomDataStore) GetConfiguration() PluginConfiguration {
	if o == nil {
		var ret PluginConfiguration
		return ret
	}

	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value
// and a boolean to check if the value has been set.
func (o *CustomDataStore) GetConfigurationOk() (*PluginConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Configuration, true
}

// SetConfiguration sets field value
func (o *CustomDataStore) SetConfiguration(v PluginConfiguration) {
	o.Configuration = v
}

func (o CustomDataStore) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomDataStore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedDataStore, errDataStore := json.Marshal(o.DataStore)
	if errDataStore != nil {
		return map[string]interface{}{}, errDataStore
	}
	errDataStore = json.Unmarshal([]byte(serializedDataStore), &toSerialize)
	if errDataStore != nil {
		return map[string]interface{}{}, errDataStore
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MaskAttributeValues) {
		toSerialize["maskAttributeValues"] = o.MaskAttributeValues
	}
	toSerialize["name"] = o.Name
	toSerialize["pluginDescriptorRef"] = o.PluginDescriptorRef
	if !IsNil(o.ParentRef) {
		toSerialize["parentRef"] = o.ParentRef
	}
	toSerialize["configuration"] = o.Configuration
	return toSerialize, nil
}

type NullableCustomDataStore struct {
	value *CustomDataStore
	isSet bool
}

func (v NullableCustomDataStore) Get() *CustomDataStore {
	return v.value
}

func (v *NullableCustomDataStore) Set(val *CustomDataStore) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomDataStore) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomDataStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomDataStore(val *CustomDataStore) *NullableCustomDataStore {
	return &NullableCustomDataStore{value: val, isSet: true}
}

func (v NullableCustomDataStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomDataStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
