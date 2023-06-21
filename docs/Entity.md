# Entity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to **string** | Unique entity identifier. | [optional] 
**EntityDescription** | Pointer to **string** | Entity description. | [optional] 

## Methods

### NewEntity

`func NewEntity() *Entity`

NewEntity instantiates a new Entity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityWithDefaults

`func NewEntityWithDefaults() *Entity`

NewEntityWithDefaults instantiates a new Entity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *Entity) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *Entity) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *Entity) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *Entity) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEntityDescription

`func (o *Entity) GetEntityDescription() string`

GetEntityDescription returns the EntityDescription field if non-nil, zero value otherwise.

### GetEntityDescriptionOk

`func (o *Entity) GetEntityDescriptionOk() (*string, bool)`

GetEntityDescriptionOk returns a tuple with the EntityDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityDescription

`func (o *Entity) SetEntityDescription(v string)`

SetEntityDescription sets EntityDescription field to given value.

### HasEntityDescription

`func (o *Entity) HasEntityDescription() bool`

HasEntityDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


