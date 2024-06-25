# ClusterNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The IP address and port this node is running on. | [optional] 
**Index** | Pointer to **int64** | Index of the node within the cluster, or -1 if an index is not assigned. | [optional] 
**Mode** | Pointer to **string** | The deployment mode of this node, from a clustering standpoint. CLUSTERED_DUAL is not supported. | [optional] 
**NodeGroup** | Pointer to **string** | The node group for this node. This field is only populated if adaptive clustering is enabled. | [optional] 
**Version** | Pointer to **string** | The PingFederate version this node is running on. | [optional] 
**NodeTags** | Pointer to **string** | The node tags for this node. This field is only populated for engine nodes. | [optional] 
**ConfigurationTimestamp** | Pointer to **time.Time** | The time stamp of the configuration data retrieved by this node. | [optional] 
**ReplicationStatus** | Pointer to **string** | The replication status of the node. | [optional] 
**AdminConsoleInfo** | Pointer to [**AdminConsoleInfo**](AdminConsoleInfo.md) |  | [optional] 

## Methods

### NewClusterNode

`func NewClusterNode() *ClusterNode`

NewClusterNode instantiates a new ClusterNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNodeWithDefaults

`func NewClusterNodeWithDefaults() *ClusterNode`

NewClusterNodeWithDefaults instantiates a new ClusterNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ClusterNode) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ClusterNode) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ClusterNode) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ClusterNode) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetIndex

`func (o *ClusterNode) GetIndex() int64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ClusterNode) GetIndexOk() (*int64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ClusterNode) SetIndex(v int64)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ClusterNode) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetMode

`func (o *ClusterNode) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ClusterNode) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ClusterNode) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ClusterNode) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetNodeGroup

`func (o *ClusterNode) GetNodeGroup() string`

GetNodeGroup returns the NodeGroup field if non-nil, zero value otherwise.

### GetNodeGroupOk

`func (o *ClusterNode) GetNodeGroupOk() (*string, bool)`

GetNodeGroupOk returns a tuple with the NodeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeGroup

`func (o *ClusterNode) SetNodeGroup(v string)`

SetNodeGroup sets NodeGroup field to given value.

### HasNodeGroup

`func (o *ClusterNode) HasNodeGroup() bool`

HasNodeGroup returns a boolean if a field has been set.

### GetVersion

`func (o *ClusterNode) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ClusterNode) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ClusterNode) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ClusterNode) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetNodeTags

`func (o *ClusterNode) GetNodeTags() string`

GetNodeTags returns the NodeTags field if non-nil, zero value otherwise.

### GetNodeTagsOk

`func (o *ClusterNode) GetNodeTagsOk() (*string, bool)`

GetNodeTagsOk returns a tuple with the NodeTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeTags

`func (o *ClusterNode) SetNodeTags(v string)`

SetNodeTags sets NodeTags field to given value.

### HasNodeTags

`func (o *ClusterNode) HasNodeTags() bool`

HasNodeTags returns a boolean if a field has been set.

### GetConfigurationTimestamp

`func (o *ClusterNode) GetConfigurationTimestamp() time.Time`

GetConfigurationTimestamp returns the ConfigurationTimestamp field if non-nil, zero value otherwise.

### GetConfigurationTimestampOk

`func (o *ClusterNode) GetConfigurationTimestampOk() (*time.Time, bool)`

GetConfigurationTimestampOk returns a tuple with the ConfigurationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationTimestamp

`func (o *ClusterNode) SetConfigurationTimestamp(v time.Time)`

SetConfigurationTimestamp sets ConfigurationTimestamp field to given value.

### HasConfigurationTimestamp

`func (o *ClusterNode) HasConfigurationTimestamp() bool`

HasConfigurationTimestamp returns a boolean if a field has been set.

### GetReplicationStatus

`func (o *ClusterNode) GetReplicationStatus() string`

GetReplicationStatus returns the ReplicationStatus field if non-nil, zero value otherwise.

### GetReplicationStatusOk

`func (o *ClusterNode) GetReplicationStatusOk() (*string, bool)`

GetReplicationStatusOk returns a tuple with the ReplicationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStatus

`func (o *ClusterNode) SetReplicationStatus(v string)`

SetReplicationStatus sets ReplicationStatus field to given value.

### HasReplicationStatus

`func (o *ClusterNode) HasReplicationStatus() bool`

HasReplicationStatus returns a boolean if a field has been set.

### GetAdminConsoleInfo

`func (o *ClusterNode) GetAdminConsoleInfo() AdminConsoleInfo`

GetAdminConsoleInfo returns the AdminConsoleInfo field if non-nil, zero value otherwise.

### GetAdminConsoleInfoOk

`func (o *ClusterNode) GetAdminConsoleInfoOk() (*AdminConsoleInfo, bool)`

GetAdminConsoleInfoOk returns a tuple with the AdminConsoleInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminConsoleInfo

`func (o *ClusterNode) SetAdminConsoleInfo(v AdminConsoleInfo)`

SetAdminConsoleInfo sets AdminConsoleInfo field to given value.

### HasAdminConsoleInfo

`func (o *ClusterNode) HasAdminConsoleInfo() bool`

HasAdminConsoleInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


