# ConfigOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | The identifier for the resource type the operation applies to. | 
**SubResource** | Pointer to **string** | The subresource for the operation. | [optional] 
**OperationType** | **string** | The type of operation to be performed. | 
**Items** | Pointer to **[]map[string]interface{}** | The configuration items for the operation. This field only applies to the SAVE operation type. | [optional] 
**ItemIds** | Pointer to **[]string** | The item ID&#39;s for the operation. This field only applies to the DELETE operation type. | [optional] 

## Methods

### NewConfigOperation

`func NewConfigOperation(resourceType string, operationType string, ) *ConfigOperation`

NewConfigOperation instantiates a new ConfigOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigOperationWithDefaults

`func NewConfigOperationWithDefaults() *ConfigOperation`

NewConfigOperationWithDefaults instantiates a new ConfigOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *ConfigOperation) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ConfigOperation) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ConfigOperation) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetSubResource

`func (o *ConfigOperation) GetSubResource() string`

GetSubResource returns the SubResource field if non-nil, zero value otherwise.

### GetSubResourceOk

`func (o *ConfigOperation) GetSubResourceOk() (*string, bool)`

GetSubResourceOk returns a tuple with the SubResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubResource

`func (o *ConfigOperation) SetSubResource(v string)`

SetSubResource sets SubResource field to given value.

### HasSubResource

`func (o *ConfigOperation) HasSubResource() bool`

HasSubResource returns a boolean if a field has been set.

### GetOperationType

`func (o *ConfigOperation) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *ConfigOperation) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *ConfigOperation) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetItems

`func (o *ConfigOperation) GetItems() []map[string]interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ConfigOperation) GetItemsOk() (*[]map[string]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ConfigOperation) SetItems(v []map[string]interface{})`

SetItems sets Items field to given value.

### HasItems

`func (o *ConfigOperation) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetItemIds

`func (o *ConfigOperation) GetItemIds() []string`

GetItemIds returns the ItemIds field if non-nil, zero value otherwise.

### GetItemIdsOk

`func (o *ConfigOperation) GetItemIdsOk() (*[]string, bool)`

GetItemIdsOk returns a tuple with the ItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemIds

`func (o *ConfigOperation) SetItemIds(v []string)`

SetItemIds sets ItemIds field to given value.

### HasItemIds

`func (o *ConfigOperation) HasItemIds() bool`

HasItemIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


