/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
	"time"
)

// checks if the CaptchaProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CaptchaProvider{}

// CaptchaProvider A CAPTCHA provider plugin instance.
type CaptchaProvider struct {
	// The ID of the plugin instance. The ID cannot be modified once the instance is created.<br>Note: Ignored when specifying a connection's adapter override.
	Id string `json:"id" tfsdk:"id"`
	// The plugin instance name. The name can be modified once the instance is created.<br>Note: Ignored when specifying a connection's adapter override.
	Name                string              `json:"name" tfsdk:"name"`
	PluginDescriptorRef ResourceLink        `json:"pluginDescriptorRef" tfsdk:"plugin_descriptor_ref"`
	ParentRef           *ResourceLink       `json:"parentRef,omitempty" tfsdk:"parent_ref"`
	Configuration       PluginConfiguration `json:"configuration" tfsdk:"configuration"`
	// The time at which the plugin instance was last changed. This property is read only and is ignored on PUT and POST requests.
	LastModified *time.Time `json:"lastModified,omitempty" tfsdk:"last_modified"`
}

// NewCaptchaProvider instantiates a new CaptchaProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaptchaProvider(id string, name string, pluginDescriptorRef ResourceLink, configuration PluginConfiguration) *CaptchaProvider {
	this := CaptchaProvider{}
	this.Id = id
	this.Name = name
	this.PluginDescriptorRef = pluginDescriptorRef
	this.Configuration = configuration
	return &this
}

// NewCaptchaProviderWithDefaults instantiates a new CaptchaProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaptchaProviderWithDefaults() *CaptchaProvider {
	this := CaptchaProvider{}
	return &this
}

// GetId returns the Id field value
func (o *CaptchaProvider) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CaptchaProvider) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CaptchaProvider) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *CaptchaProvider) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CaptchaProvider) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CaptchaProvider) SetName(v string) {
	o.Name = v
}

// GetPluginDescriptorRef returns the PluginDescriptorRef field value
func (o *CaptchaProvider) GetPluginDescriptorRef() ResourceLink {
	if o == nil {
		var ret ResourceLink
		return ret
	}

	return o.PluginDescriptorRef
}

// GetPluginDescriptorRefOk returns a tuple with the PluginDescriptorRef field value
// and a boolean to check if the value has been set.
func (o *CaptchaProvider) GetPluginDescriptorRefOk() (*ResourceLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PluginDescriptorRef, true
}

// SetPluginDescriptorRef sets field value
func (o *CaptchaProvider) SetPluginDescriptorRef(v ResourceLink) {
	o.PluginDescriptorRef = v
}

// GetParentRef returns the ParentRef field value if set, zero value otherwise.
func (o *CaptchaProvider) GetParentRef() ResourceLink {
	if o == nil || IsNil(o.ParentRef) {
		var ret ResourceLink
		return ret
	}
	return *o.ParentRef
}

// GetParentRefOk returns a tuple with the ParentRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptchaProvider) GetParentRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.ParentRef) {
		return nil, false
	}
	return o.ParentRef, true
}

// HasParentRef returns a boolean if a field has been set.
func (o *CaptchaProvider) HasParentRef() bool {
	if o != nil && !IsNil(o.ParentRef) {
		return true
	}

	return false
}

// SetParentRef gets a reference to the given ResourceLink and assigns it to the ParentRef field.
func (o *CaptchaProvider) SetParentRef(v ResourceLink) {
	o.ParentRef = &v
}

// GetConfiguration returns the Configuration field value
func (o *CaptchaProvider) GetConfiguration() PluginConfiguration {
	if o == nil {
		var ret PluginConfiguration
		return ret
	}

	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value
// and a boolean to check if the value has been set.
func (o *CaptchaProvider) GetConfigurationOk() (*PluginConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Configuration, true
}

// SetConfiguration sets field value
func (o *CaptchaProvider) SetConfiguration(v PluginConfiguration) {
	o.Configuration = v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *CaptchaProvider) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptchaProvider) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *CaptchaProvider) HasLastModified() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *CaptchaProvider) SetLastModified(v time.Time) {
	o.LastModified = &v
}

func (o CaptchaProvider) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CaptchaProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["pluginDescriptorRef"] = o.PluginDescriptorRef
	if !IsNil(o.ParentRef) {
		toSerialize["parentRef"] = o.ParentRef
	}
	toSerialize["configuration"] = o.Configuration
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	return toSerialize, nil
}

type NullableCaptchaProvider struct {
	value *CaptchaProvider
	isSet bool
}

func (v NullableCaptchaProvider) Get() *CaptchaProvider {
	return v.value
}

func (v *NullableCaptchaProvider) Set(val *CaptchaProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptchaProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptchaProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptchaProvider(val *CaptchaProvider) *NullableCaptchaProvider {
	return &NullableCaptchaProvider{value: val, isSet: true}
}

func (v NullableCaptchaProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptchaProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
