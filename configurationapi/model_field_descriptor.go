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

// checks if the FieldDescriptor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FieldDescriptor{}

// FieldDescriptor Describes a plugin configuration field.
type FieldDescriptor struct {
	// The type of field descriptor.
	Type *string `json:"type,omitempty" tfsdk:"type"`
	// Name of the field.
	Name *string `json:"name,omitempty" tfsdk:"name"`
	// Description of the field.
	Description *string `json:"description,omitempty" tfsdk:"description"`
	// Default value of the field. This is the value pre-populated in the UI on new plugin instance configuration. This is also the value used to populate the field if it is missing in a POST or PUT request and no 'defaultForLegacyConfig' is defined.
	DefaultValue *string `json:"defaultValue,omitempty" tfsdk:"default_value"`
	// Default value of the field when it is missing from the configuration (e.g. in upgrade scenarios). This is the value pre-populated in the UI for existing plugin configurations without values for the field. This is also the value used to populate the field if it is missing in a POST or PUT request. If 'defaultForLegacyConfig' is not defined, PingFederate will fall back to applying the 'defaultValue' to the field.
	DefaultForLegacyConfig *string `json:"defaultForLegacyConfig,omitempty" tfsdk:"default_for_legacy_config"`
	// Whether this is an advanced field or not.
	Advanced *bool `json:"advanced,omitempty" tfsdk:"advanced"`
	// Whether a value is required for this field or not.
	Required *bool `json:"required,omitempty" tfsdk:"required"`
	// Label of the field to be displayed in the administrative console.
	Label *string `json:"label,omitempty" tfsdk:"label"`
}

// NewFieldDescriptor instantiates a new FieldDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldDescriptor() *FieldDescriptor {
	this := FieldDescriptor{}
	return &this
}

// NewFieldDescriptorWithDefaults instantiates a new FieldDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldDescriptorWithDefaults() *FieldDescriptor {
	this := FieldDescriptor{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FieldDescriptor) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldDescriptor) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FieldDescriptor) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FieldDescriptor) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FieldDescriptor) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldDescriptor) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FieldDescriptor) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FieldDescriptor) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FieldDescriptor) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldDescriptor) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FieldDescriptor) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FieldDescriptor) SetDescription(v string) {
	o.Description = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *FieldDescriptor) GetDefaultValue() string {
	if o == nil || IsNil(o.DefaultValue) {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldDescriptor) GetDefaultValueOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultValue) {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *FieldDescriptor) HasDefaultValue() bool {
	if o != nil && !IsNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *FieldDescriptor) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// GetDefaultForLegacyConfig returns the DefaultForLegacyConfig field value if set, zero value otherwise.
func (o *FieldDescriptor) GetDefaultForLegacyConfig() string {
	if o == nil || IsNil(o.DefaultForLegacyConfig) {
		var ret string
		return ret
	}
	return *o.DefaultForLegacyConfig
}

// GetDefaultForLegacyConfigOk returns a tuple with the DefaultForLegacyConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldDescriptor) GetDefaultForLegacyConfigOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultForLegacyConfig) {
		return nil, false
	}
	return o.DefaultForLegacyConfig, true
}

// HasDefaultForLegacyConfig returns a boolean if a field has been set.
func (o *FieldDescriptor) HasDefaultForLegacyConfig() bool {
	if o != nil && !IsNil(o.DefaultForLegacyConfig) {
		return true
	}

	return false
}

// SetDefaultForLegacyConfig gets a reference to the given string and assigns it to the DefaultForLegacyConfig field.
func (o *FieldDescriptor) SetDefaultForLegacyConfig(v string) {
	o.DefaultForLegacyConfig = &v
}

// GetAdvanced returns the Advanced field value if set, zero value otherwise.
func (o *FieldDescriptor) GetAdvanced() bool {
	if o == nil || IsNil(o.Advanced) {
		var ret bool
		return ret
	}
	return *o.Advanced
}

// GetAdvancedOk returns a tuple with the Advanced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldDescriptor) GetAdvancedOk() (*bool, bool) {
	if o == nil || IsNil(o.Advanced) {
		return nil, false
	}
	return o.Advanced, true
}

// HasAdvanced returns a boolean if a field has been set.
func (o *FieldDescriptor) HasAdvanced() bool {
	if o != nil && !IsNil(o.Advanced) {
		return true
	}

	return false
}

// SetAdvanced gets a reference to the given bool and assigns it to the Advanced field.
func (o *FieldDescriptor) SetAdvanced(v bool) {
	o.Advanced = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *FieldDescriptor) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldDescriptor) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *FieldDescriptor) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *FieldDescriptor) SetRequired(v bool) {
	o.Required = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *FieldDescriptor) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldDescriptor) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *FieldDescriptor) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *FieldDescriptor) SetLabel(v string) {
	o.Label = &v
}

func (o FieldDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FieldDescriptor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DefaultValue) {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if !IsNil(o.DefaultForLegacyConfig) {
		toSerialize["defaultForLegacyConfig"] = o.DefaultForLegacyConfig
	}
	if !IsNil(o.Advanced) {
		toSerialize["advanced"] = o.Advanced
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	return toSerialize, nil
}

type NullableFieldDescriptor struct {
	value *FieldDescriptor
	isSet bool
}

func (v NullableFieldDescriptor) Get() *FieldDescriptor {
	return v.value
}

func (v *NullableFieldDescriptor) Set(val *FieldDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldDescriptor(val *FieldDescriptor) *NullableFieldDescriptor {
	return &NullableFieldDescriptor{value: val, isSet: true}
}

func (v NullableFieldDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
