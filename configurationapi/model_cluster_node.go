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

// checks if the ClusterNode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterNode{}

// ClusterNode Describes a node in a clustered deployment of PingFederate.
type ClusterNode struct {
	// The IP address and port this node is running on.
	Address *string `json:"address,omitempty" tfsdk:"address"`
	// Index of the node within the cluster, or -1 if an index is not assigned.
	Index *int64 `json:"index,omitempty" tfsdk:"index"`
	// The deployment mode of this node, from a clustering standpoint. CLUSTERED_DUAL is not supported.
	Mode *string `json:"mode,omitempty" tfsdk:"mode"`
	// The node group for this node. This field is only populated if adaptive clustering is enabled.
	NodeGroup *string `json:"nodeGroup,omitempty" tfsdk:"node_group"`
	// The PingFederate version this node is running on.
	Version *string `json:"version,omitempty" tfsdk:"version"`
	// The node tags for this node. This field is only populated for engine nodes.
	NodeTags *string `json:"nodeTags,omitempty" tfsdk:"node_tags"`
	// The time stamp of the configuration data retrieved by this node.
	ConfigurationTimestamp *time.Time `json:"configurationTimestamp,omitempty" tfsdk:"configuration_timestamp"`
	// The replication status of the node.
	ReplicationStatus *string           `json:"replicationStatus,omitempty" tfsdk:"replication_status"`
	AdminConsoleInfo  *AdminConsoleInfo `json:"adminConsoleInfo,omitempty" tfsdk:"admin_console_info"`
}

// NewClusterNode instantiates a new ClusterNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterNode() *ClusterNode {
	this := ClusterNode{}
	return &this
}

// NewClusterNodeWithDefaults instantiates a new ClusterNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterNodeWithDefaults() *ClusterNode {
	this := ClusterNode{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ClusterNode) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNode) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ClusterNode) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ClusterNode) SetAddress(v string) {
	o.Address = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *ClusterNode) GetIndex() int64 {
	if o == nil || IsNil(o.Index) {
		var ret int64
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNode) GetIndexOk() (*int64, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *ClusterNode) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int64 and assigns it to the Index field.
func (o *ClusterNode) SetIndex(v int64) {
	o.Index = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *ClusterNode) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNode) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *ClusterNode) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *ClusterNode) SetMode(v string) {
	o.Mode = &v
}

// GetNodeGroup returns the NodeGroup field value if set, zero value otherwise.
func (o *ClusterNode) GetNodeGroup() string {
	if o == nil || IsNil(o.NodeGroup) {
		var ret string
		return ret
	}
	return *o.NodeGroup
}

// GetNodeGroupOk returns a tuple with the NodeGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNode) GetNodeGroupOk() (*string, bool) {
	if o == nil || IsNil(o.NodeGroup) {
		return nil, false
	}
	return o.NodeGroup, true
}

// HasNodeGroup returns a boolean if a field has been set.
func (o *ClusterNode) HasNodeGroup() bool {
	if o != nil && !IsNil(o.NodeGroup) {
		return true
	}

	return false
}

// SetNodeGroup gets a reference to the given string and assigns it to the NodeGroup field.
func (o *ClusterNode) SetNodeGroup(v string) {
	o.NodeGroup = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ClusterNode) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNode) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ClusterNode) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ClusterNode) SetVersion(v string) {
	o.Version = &v
}

// GetNodeTags returns the NodeTags field value if set, zero value otherwise.
func (o *ClusterNode) GetNodeTags() string {
	if o == nil || IsNil(o.NodeTags) {
		var ret string
		return ret
	}
	return *o.NodeTags
}

// GetNodeTagsOk returns a tuple with the NodeTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNode) GetNodeTagsOk() (*string, bool) {
	if o == nil || IsNil(o.NodeTags) {
		return nil, false
	}
	return o.NodeTags, true
}

// HasNodeTags returns a boolean if a field has been set.
func (o *ClusterNode) HasNodeTags() bool {
	if o != nil && !IsNil(o.NodeTags) {
		return true
	}

	return false
}

// SetNodeTags gets a reference to the given string and assigns it to the NodeTags field.
func (o *ClusterNode) SetNodeTags(v string) {
	o.NodeTags = &v
}

// GetConfigurationTimestamp returns the ConfigurationTimestamp field value if set, zero value otherwise.
func (o *ClusterNode) GetConfigurationTimestamp() time.Time {
	if o == nil || IsNil(o.ConfigurationTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.ConfigurationTimestamp
}

// GetConfigurationTimestampOk returns a tuple with the ConfigurationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNode) GetConfigurationTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ConfigurationTimestamp) {
		return nil, false
	}
	return o.ConfigurationTimestamp, true
}

// HasConfigurationTimestamp returns a boolean if a field has been set.
func (o *ClusterNode) HasConfigurationTimestamp() bool {
	if o != nil && !IsNil(o.ConfigurationTimestamp) {
		return true
	}

	return false
}

// SetConfigurationTimestamp gets a reference to the given time.Time and assigns it to the ConfigurationTimestamp field.
func (o *ClusterNode) SetConfigurationTimestamp(v time.Time) {
	o.ConfigurationTimestamp = &v
}

// GetReplicationStatus returns the ReplicationStatus field value if set, zero value otherwise.
func (o *ClusterNode) GetReplicationStatus() string {
	if o == nil || IsNil(o.ReplicationStatus) {
		var ret string
		return ret
	}
	return *o.ReplicationStatus
}

// GetReplicationStatusOk returns a tuple with the ReplicationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNode) GetReplicationStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicationStatus) {
		return nil, false
	}
	return o.ReplicationStatus, true
}

// HasReplicationStatus returns a boolean if a field has been set.
func (o *ClusterNode) HasReplicationStatus() bool {
	if o != nil && !IsNil(o.ReplicationStatus) {
		return true
	}

	return false
}

// SetReplicationStatus gets a reference to the given string and assigns it to the ReplicationStatus field.
func (o *ClusterNode) SetReplicationStatus(v string) {
	o.ReplicationStatus = &v
}

// GetAdminConsoleInfo returns the AdminConsoleInfo field value if set, zero value otherwise.
func (o *ClusterNode) GetAdminConsoleInfo() AdminConsoleInfo {
	if o == nil || IsNil(o.AdminConsoleInfo) {
		var ret AdminConsoleInfo
		return ret
	}
	return *o.AdminConsoleInfo
}

// GetAdminConsoleInfoOk returns a tuple with the AdminConsoleInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterNode) GetAdminConsoleInfoOk() (*AdminConsoleInfo, bool) {
	if o == nil || IsNil(o.AdminConsoleInfo) {
		return nil, false
	}
	return o.AdminConsoleInfo, true
}

// HasAdminConsoleInfo returns a boolean if a field has been set.
func (o *ClusterNode) HasAdminConsoleInfo() bool {
	if o != nil && !IsNil(o.AdminConsoleInfo) {
		return true
	}

	return false
}

// SetAdminConsoleInfo gets a reference to the given AdminConsoleInfo and assigns it to the AdminConsoleInfo field.
func (o *ClusterNode) SetAdminConsoleInfo(v AdminConsoleInfo) {
	o.AdminConsoleInfo = &v
}

func (o ClusterNode) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterNode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.NodeGroup) {
		toSerialize["nodeGroup"] = o.NodeGroup
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.NodeTags) {
		toSerialize["nodeTags"] = o.NodeTags
	}
	if !IsNil(o.ConfigurationTimestamp) {
		toSerialize["configurationTimestamp"] = o.ConfigurationTimestamp
	}
	if !IsNil(o.ReplicationStatus) {
		toSerialize["replicationStatus"] = o.ReplicationStatus
	}
	if !IsNil(o.AdminConsoleInfo) {
		toSerialize["adminConsoleInfo"] = o.AdminConsoleInfo
	}
	return toSerialize, nil
}

type NullableClusterNode struct {
	value *ClusterNode
	isSet bool
}

func (v NullableClusterNode) Get() *ClusterNode {
	return v.value
}

func (v *NullableClusterNode) Set(val *ClusterNode) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterNode) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterNode(val *ClusterNode) *NullableClusterNode {
	return &NullableClusterNode{value: val, isSet: true}
}

func (v NullableClusterNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
