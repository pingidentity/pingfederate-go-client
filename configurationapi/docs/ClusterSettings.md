# ClusterSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReplicateConnectionsOnSave** | Pointer to **bool** | Whether changes to connections will automatically be replicated to the cluster. | [optional] 
**ReplicateClientsOnSave** | Pointer to **bool** | Whether changes to OAuth clients will automatically be replicated to the cluster. This setting only applies when using XML Client storage. | [optional] 

## Methods

### NewClusterSettings

`func NewClusterSettings() *ClusterSettings`

NewClusterSettings instantiates a new ClusterSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterSettingsWithDefaults

`func NewClusterSettingsWithDefaults() *ClusterSettings`

NewClusterSettingsWithDefaults instantiates a new ClusterSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplicateConnectionsOnSave

`func (o *ClusterSettings) GetReplicateConnectionsOnSave() bool`

GetReplicateConnectionsOnSave returns the ReplicateConnectionsOnSave field if non-nil, zero value otherwise.

### GetReplicateConnectionsOnSaveOk

`func (o *ClusterSettings) GetReplicateConnectionsOnSaveOk() (*bool, bool)`

GetReplicateConnectionsOnSaveOk returns a tuple with the ReplicateConnectionsOnSave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicateConnectionsOnSave

`func (o *ClusterSettings) SetReplicateConnectionsOnSave(v bool)`

SetReplicateConnectionsOnSave sets ReplicateConnectionsOnSave field to given value.

### HasReplicateConnectionsOnSave

`func (o *ClusterSettings) HasReplicateConnectionsOnSave() bool`

HasReplicateConnectionsOnSave returns a boolean if a field has been set.

### GetReplicateClientsOnSave

`func (o *ClusterSettings) GetReplicateClientsOnSave() bool`

GetReplicateClientsOnSave returns the ReplicateClientsOnSave field if non-nil, zero value otherwise.

### GetReplicateClientsOnSaveOk

`func (o *ClusterSettings) GetReplicateClientsOnSaveOk() (*bool, bool)`

GetReplicateClientsOnSaveOk returns a tuple with the ReplicateClientsOnSave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicateClientsOnSave

`func (o *ClusterSettings) SetReplicateClientsOnSave(v bool)`

SetReplicateClientsOnSave sets ReplicateClientsOnSave field to given value.

### HasReplicateClientsOnSave

`func (o *ClusterSettings) HasReplicateClientsOnSave() bool`

HasReplicateClientsOnSave returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


