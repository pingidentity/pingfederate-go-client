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

// checks if the PasswordCredentialValidator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordCredentialValidator{}

// PasswordCredentialValidator struct for PasswordCredentialValidator
type PasswordCredentialValidator struct {
	// The ID of the plugin instance. The ID cannot be modified once the instance is created.<br>Note: Ignored when specifying a connection's adapter override.
	Id string `json:"id" tfsdk:"id"`
	// The plugin instance name. The name can be modified once the instance is created.<br>Note: Ignored when specifying a connection's adapter override.
	Name                string              `json:"name" tfsdk:"name"`
	PluginDescriptorRef ResourceLink        `json:"pluginDescriptorRef" tfsdk:"plugin_descriptor_ref"`
	ParentRef           *ResourceLink       `json:"parentRef,omitempty" tfsdk:"parent_ref"`
	Configuration       PluginConfiguration `json:"configuration" tfsdk:"configuration"`
	// The time at which the plugin instance was last changed. This property is read only and is ignored on PUT and POST requests.
	LastModified      *time.Time                                    `json:"lastModified,omitempty" tfsdk:"last_modified"`
	AttributeContract *PasswordCredentialValidatorAttributeContract `json:"attributeContract,omitempty" tfsdk:"attribute_contract"`
}

// NewPasswordCredentialValidator instantiates a new PasswordCredentialValidator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordCredentialValidator(id string, name string, pluginDescriptorRef ResourceLink, configuration PluginConfiguration) *PasswordCredentialValidator {
	this := PasswordCredentialValidator{}
	this.Id = id
	this.Name = name
	this.PluginDescriptorRef = pluginDescriptorRef
	this.Configuration = configuration
	return &this
}

// NewPasswordCredentialValidatorWithDefaults instantiates a new PasswordCredentialValidator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordCredentialValidatorWithDefaults() *PasswordCredentialValidator {
	this := PasswordCredentialValidator{}
	return &this
}

// GetId returns the Id field value
func (o *PasswordCredentialValidator) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PasswordCredentialValidator) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PasswordCredentialValidator) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *PasswordCredentialValidator) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PasswordCredentialValidator) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PasswordCredentialValidator) SetName(v string) {
	o.Name = v
}

// GetPluginDescriptorRef returns the PluginDescriptorRef field value
func (o *PasswordCredentialValidator) GetPluginDescriptorRef() ResourceLink {
	if o == nil {
		var ret ResourceLink
		return ret
	}

	return o.PluginDescriptorRef
}

// GetPluginDescriptorRefOk returns a tuple with the PluginDescriptorRef field value
// and a boolean to check if the value has been set.
func (o *PasswordCredentialValidator) GetPluginDescriptorRefOk() (*ResourceLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PluginDescriptorRef, true
}

// SetPluginDescriptorRef sets field value
func (o *PasswordCredentialValidator) SetPluginDescriptorRef(v ResourceLink) {
	o.PluginDescriptorRef = v
}

// GetParentRef returns the ParentRef field value if set, zero value otherwise.
func (o *PasswordCredentialValidator) GetParentRef() ResourceLink {
	if o == nil || IsNil(o.ParentRef) {
		var ret ResourceLink
		return ret
	}
	return *o.ParentRef
}

// GetParentRefOk returns a tuple with the ParentRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordCredentialValidator) GetParentRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.ParentRef) {
		return nil, false
	}
	return o.ParentRef, true
}

// HasParentRef returns a boolean if a field has been set.
func (o *PasswordCredentialValidator) HasParentRef() bool {
	if o != nil && !IsNil(o.ParentRef) {
		return true
	}

	return false
}

// SetParentRef gets a reference to the given ResourceLink and assigns it to the ParentRef field.
func (o *PasswordCredentialValidator) SetParentRef(v ResourceLink) {
	o.ParentRef = &v
}

// GetConfiguration returns the Configuration field value
func (o *PasswordCredentialValidator) GetConfiguration() PluginConfiguration {
	if o == nil {
		var ret PluginConfiguration
		return ret
	}

	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value
// and a boolean to check if the value has been set.
func (o *PasswordCredentialValidator) GetConfigurationOk() (*PluginConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Configuration, true
}

// SetConfiguration sets field value
func (o *PasswordCredentialValidator) SetConfiguration(v PluginConfiguration) {
	o.Configuration = v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *PasswordCredentialValidator) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordCredentialValidator) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *PasswordCredentialValidator) HasLastModified() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *PasswordCredentialValidator) SetLastModified(v time.Time) {
	o.LastModified = &v
}

// GetAttributeContract returns the AttributeContract field value if set, zero value otherwise.
func (o *PasswordCredentialValidator) GetAttributeContract() PasswordCredentialValidatorAttributeContract {
	if o == nil || IsNil(o.AttributeContract) {
		var ret PasswordCredentialValidatorAttributeContract
		return ret
	}
	return *o.AttributeContract
}

// GetAttributeContractOk returns a tuple with the AttributeContract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordCredentialValidator) GetAttributeContractOk() (*PasswordCredentialValidatorAttributeContract, bool) {
	if o == nil || IsNil(o.AttributeContract) {
		return nil, false
	}
	return o.AttributeContract, true
}

// HasAttributeContract returns a boolean if a field has been set.
func (o *PasswordCredentialValidator) HasAttributeContract() bool {
	if o != nil && !IsNil(o.AttributeContract) {
		return true
	}

	return false
}

// SetAttributeContract gets a reference to the given PasswordCredentialValidatorAttributeContract and assigns it to the AttributeContract field.
func (o *PasswordCredentialValidator) SetAttributeContract(v PasswordCredentialValidatorAttributeContract) {
	o.AttributeContract = &v
}

func (o PasswordCredentialValidator) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordCredentialValidator) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AttributeContract) {
		toSerialize["attributeContract"] = o.AttributeContract
	}
	return toSerialize, nil
}

type NullablePasswordCredentialValidator struct {
	value *PasswordCredentialValidator
	isSet bool
}

func (v NullablePasswordCredentialValidator) Get() *PasswordCredentialValidator {
	return v.value
}

func (v *NullablePasswordCredentialValidator) Set(val *PasswordCredentialValidator) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordCredentialValidator) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordCredentialValidator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordCredentialValidator(val *PasswordCredentialValidator) *NullablePasswordCredentialValidator {
	return &NullablePasswordCredentialValidator{value: val, isSet: true}
}

func (v NullablePasswordCredentialValidator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordCredentialValidator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
