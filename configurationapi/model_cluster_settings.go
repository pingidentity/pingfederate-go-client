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

// checks if the ClusterSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterSettings{}

// ClusterSettings Cluster configuration settings.
type ClusterSettings struct {
	// Whether changes to connections will automatically be replicated to the cluster.
	ReplicateConnectionsOnSave *bool `json:"replicateConnectionsOnSave,omitempty" tfsdk:"replicate_connections_on_save"`
	// Whether changes to OAuth clients will automatically be replicated to the cluster. This setting only applies when using XML Client storage.
	ReplicateClientsOnSave *bool `json:"replicateClientsOnSave,omitempty" tfsdk:"replicate_clients_on_save"`
}

// NewClusterSettings instantiates a new ClusterSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterSettings() *ClusterSettings {
	this := ClusterSettings{}
	return &this
}

// NewClusterSettingsWithDefaults instantiates a new ClusterSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterSettingsWithDefaults() *ClusterSettings {
	this := ClusterSettings{}
	return &this
}

// GetReplicateConnectionsOnSave returns the ReplicateConnectionsOnSave field value if set, zero value otherwise.
func (o *ClusterSettings) GetReplicateConnectionsOnSave() bool {
	if o == nil || IsNil(o.ReplicateConnectionsOnSave) {
		var ret bool
		return ret
	}
	return *o.ReplicateConnectionsOnSave
}

// GetReplicateConnectionsOnSaveOk returns a tuple with the ReplicateConnectionsOnSave field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSettings) GetReplicateConnectionsOnSaveOk() (*bool, bool) {
	if o == nil || IsNil(o.ReplicateConnectionsOnSave) {
		return nil, false
	}
	return o.ReplicateConnectionsOnSave, true
}

// HasReplicateConnectionsOnSave returns a boolean if a field has been set.
func (o *ClusterSettings) HasReplicateConnectionsOnSave() bool {
	if o != nil && !IsNil(o.ReplicateConnectionsOnSave) {
		return true
	}

	return false
}

// SetReplicateConnectionsOnSave gets a reference to the given bool and assigns it to the ReplicateConnectionsOnSave field.
func (o *ClusterSettings) SetReplicateConnectionsOnSave(v bool) {
	o.ReplicateConnectionsOnSave = &v
}

// GetReplicateClientsOnSave returns the ReplicateClientsOnSave field value if set, zero value otherwise.
func (o *ClusterSettings) GetReplicateClientsOnSave() bool {
	if o == nil || IsNil(o.ReplicateClientsOnSave) {
		var ret bool
		return ret
	}
	return *o.ReplicateClientsOnSave
}

// GetReplicateClientsOnSaveOk returns a tuple with the ReplicateClientsOnSave field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSettings) GetReplicateClientsOnSaveOk() (*bool, bool) {
	if o == nil || IsNil(o.ReplicateClientsOnSave) {
		return nil, false
	}
	return o.ReplicateClientsOnSave, true
}

// HasReplicateClientsOnSave returns a boolean if a field has been set.
func (o *ClusterSettings) HasReplicateClientsOnSave() bool {
	if o != nil && !IsNil(o.ReplicateClientsOnSave) {
		return true
	}

	return false
}

// SetReplicateClientsOnSave gets a reference to the given bool and assigns it to the ReplicateClientsOnSave field.
func (o *ClusterSettings) SetReplicateClientsOnSave(v bool) {
	o.ReplicateClientsOnSave = &v
}

func (o ClusterSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReplicateConnectionsOnSave) {
		toSerialize["replicateConnectionsOnSave"] = o.ReplicateConnectionsOnSave
	}
	if !IsNil(o.ReplicateClientsOnSave) {
		toSerialize["replicateClientsOnSave"] = o.ReplicateClientsOnSave
	}
	return toSerialize, nil
}

type NullableClusterSettings struct {
	value *ClusterSettings
	isSet bool
}

func (v NullableClusterSettings) Get() *ClusterSettings {
	return v.value
}

func (v *NullableClusterSettings) Set(val *ClusterSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterSettings(val *ClusterSettings) *NullableClusterSettings {
	return &NullableClusterSettings{value: val, isSet: true}
}

func (v NullableClusterSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
