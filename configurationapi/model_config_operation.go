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

// checks if the ConfigOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigOperation{}

// ConfigOperation Model describing a list of configuration operations for a given resource type.
type ConfigOperation struct {
	// The identifier for the resource type the operation applies to.
	ResourceType string `json:"resourceType" tfsdk:"resource_type"`
	// The subresource for the operation.
	SubResource *string `json:"subResource,omitempty" tfsdk:"sub_resource"`
	// The type of operation to be performed.
	OperationType string `json:"operationType" tfsdk:"operation_type"`
	// The configuration items for the operation. This field only applies to the SAVE operation type.
	Items []map[string]interface{} `json:"items,omitempty" tfsdk:"items"`
	// The item ID's for the operation. This field only applies to the DELETE operation type.
	ItemIds []string `json:"itemIds,omitempty" tfsdk:"item_ids"`
}

// NewConfigOperation instantiates a new ConfigOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigOperation(resourceType string, operationType string) *ConfigOperation {
	this := ConfigOperation{}
	this.ResourceType = resourceType
	this.OperationType = operationType
	return &this
}

// NewConfigOperationWithDefaults instantiates a new ConfigOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigOperationWithDefaults() *ConfigOperation {
	this := ConfigOperation{}
	return &this
}

// GetResourceType returns the ResourceType field value
func (o *ConfigOperation) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *ConfigOperation) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *ConfigOperation) SetResourceType(v string) {
	o.ResourceType = v
}

// GetSubResource returns the SubResource field value if set, zero value otherwise.
func (o *ConfigOperation) GetSubResource() string {
	if o == nil || IsNil(o.SubResource) {
		var ret string
		return ret
	}
	return *o.SubResource
}

// GetSubResourceOk returns a tuple with the SubResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigOperation) GetSubResourceOk() (*string, bool) {
	if o == nil || IsNil(o.SubResource) {
		return nil, false
	}
	return o.SubResource, true
}

// HasSubResource returns a boolean if a field has been set.
func (o *ConfigOperation) HasSubResource() bool {
	if o != nil && !IsNil(o.SubResource) {
		return true
	}

	return false
}

// SetSubResource gets a reference to the given string and assigns it to the SubResource field.
func (o *ConfigOperation) SetSubResource(v string) {
	o.SubResource = &v
}

// GetOperationType returns the OperationType field value
func (o *ConfigOperation) GetOperationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationType
}

// GetOperationTypeOk returns a tuple with the OperationType field value
// and a boolean to check if the value has been set.
func (o *ConfigOperation) GetOperationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationType, true
}

// SetOperationType sets field value
func (o *ConfigOperation) SetOperationType(v string) {
	o.OperationType = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ConfigOperation) GetItems() []map[string]interface{} {
	if o == nil || IsNil(o.Items) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigOperation) GetItemsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ConfigOperation) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []map[string]interface{} and assigns it to the Items field.
func (o *ConfigOperation) SetItems(v []map[string]interface{}) {
	o.Items = v
}

// GetItemIds returns the ItemIds field value if set, zero value otherwise.
func (o *ConfigOperation) GetItemIds() []string {
	if o == nil || IsNil(o.ItemIds) {
		var ret []string
		return ret
	}
	return o.ItemIds
}

// GetItemIdsOk returns a tuple with the ItemIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigOperation) GetItemIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ItemIds) {
		return nil, false
	}
	return o.ItemIds, true
}

// HasItemIds returns a boolean if a field has been set.
func (o *ConfigOperation) HasItemIds() bool {
	if o != nil && !IsNil(o.ItemIds) {
		return true
	}

	return false
}

// SetItemIds gets a reference to the given []string and assigns it to the ItemIds field.
func (o *ConfigOperation) SetItemIds(v []string) {
	o.ItemIds = v
}

func (o ConfigOperation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resourceType"] = o.ResourceType
	if !IsNil(o.SubResource) {
		toSerialize["subResource"] = o.SubResource
	}
	toSerialize["operationType"] = o.OperationType
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ItemIds) {
		toSerialize["itemIds"] = o.ItemIds
	}
	return toSerialize, nil
}

type NullableConfigOperation struct {
	value *ConfigOperation
	isSet bool
}

func (v NullableConfigOperation) Get() *ConfigOperation {
	return v.value
}

func (v *NullableConfigOperation) Set(val *ConfigOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigOperation(val *ConfigOperation) *NullableConfigOperation {
	return &NullableConfigOperation{value: val, isSet: true}
}

func (v NullableConfigOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
