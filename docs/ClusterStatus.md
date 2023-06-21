# ClusterStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | Pointer to [**[]ClusterNode**](ClusterNode.md) | List of nodes in the cluster. | [optional] 
**LastConfigUpdateTime** | Pointer to **time.Time** | Time when the configuration of this node was last updated. | [optional] 
**LastReplicationTime** | Pointer to **time.Time** | Time when configuration changes were last replicated. | [optional] 
**ReplicationRequired** | Pointer to **bool** | Indicates whether a replication is required to propagate config updates. | [optional] 
**MixedMode** | Pointer to **bool** | Indicates whether there is more than one version of PingFederate in the cluster. | [optional] 

## Methods

### NewClusterStatus

`func NewClusterStatus() *ClusterStatus`

NewClusterStatus instantiates a new ClusterStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterStatusWithDefaults

`func NewClusterStatusWithDefaults() *ClusterStatus`

NewClusterStatusWithDefaults instantiates a new ClusterStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *ClusterStatus) GetNodes() []ClusterNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ClusterStatus) GetNodesOk() (*[]ClusterNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ClusterStatus) SetNodes(v []ClusterNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *ClusterStatus) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetLastConfigUpdateTime

`func (o *ClusterStatus) GetLastConfigUpdateTime() time.Time`

GetLastConfigUpdateTime returns the LastConfigUpdateTime field if non-nil, zero value otherwise.

### GetLastConfigUpdateTimeOk

`func (o *ClusterStatus) GetLastConfigUpdateTimeOk() (*time.Time, bool)`

GetLastConfigUpdateTimeOk returns a tuple with the LastConfigUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConfigUpdateTime

`func (o *ClusterStatus) SetLastConfigUpdateTime(v time.Time)`

SetLastConfigUpdateTime sets LastConfigUpdateTime field to given value.

### HasLastConfigUpdateTime

`func (o *ClusterStatus) HasLastConfigUpdateTime() bool`

HasLastConfigUpdateTime returns a boolean if a field has been set.

### GetLastReplicationTime

`func (o *ClusterStatus) GetLastReplicationTime() time.Time`

GetLastReplicationTime returns the LastReplicationTime field if non-nil, zero value otherwise.

### GetLastReplicationTimeOk

`func (o *ClusterStatus) GetLastReplicationTimeOk() (*time.Time, bool)`

GetLastReplicationTimeOk returns a tuple with the LastReplicationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReplicationTime

`func (o *ClusterStatus) SetLastReplicationTime(v time.Time)`

SetLastReplicationTime sets LastReplicationTime field to given value.

### HasLastReplicationTime

`func (o *ClusterStatus) HasLastReplicationTime() bool`

HasLastReplicationTime returns a boolean if a field has been set.

### GetReplicationRequired

`func (o *ClusterStatus) GetReplicationRequired() bool`

GetReplicationRequired returns the ReplicationRequired field if non-nil, zero value otherwise.

### GetReplicationRequiredOk

`func (o *ClusterStatus) GetReplicationRequiredOk() (*bool, bool)`

GetReplicationRequiredOk returns a tuple with the ReplicationRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationRequired

`func (o *ClusterStatus) SetReplicationRequired(v bool)`

SetReplicationRequired sets ReplicationRequired field to given value.

### HasReplicationRequired

`func (o *ClusterStatus) HasReplicationRequired() bool`

HasReplicationRequired returns a boolean if a field has been set.

### GetMixedMode

`func (o *ClusterStatus) GetMixedMode() bool`

GetMixedMode returns the MixedMode field if non-nil, zero value otherwise.

### GetMixedModeOk

`func (o *ClusterStatus) GetMixedModeOk() (*bool, bool)`

GetMixedModeOk returns a tuple with the MixedMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMixedMode

`func (o *ClusterStatus) SetMixedMode(v bool)`

SetMixedMode sets MixedMode field to given value.

### HasMixedMode

`func (o *ClusterStatus) HasMixedMode() bool`

HasMixedMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


