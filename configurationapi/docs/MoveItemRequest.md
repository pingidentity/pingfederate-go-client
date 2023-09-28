# MoveItemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | **string** | Enumeration for where to move the item. | 
**MoveToId** | Pointer to **string** | When moving an item relative to another, this value indicates the target move-to ID. | [optional] 

## Methods

### NewMoveItemRequest

`func NewMoveItemRequest(location string, ) *MoveItemRequest`

NewMoveItemRequest instantiates a new MoveItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveItemRequestWithDefaults

`func NewMoveItemRequestWithDefaults() *MoveItemRequest`

NewMoveItemRequestWithDefaults instantiates a new MoveItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *MoveItemRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MoveItemRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *MoveItemRequest) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetMoveToId

`func (o *MoveItemRequest) GetMoveToId() string`

GetMoveToId returns the MoveToId field if non-nil, zero value otherwise.

### GetMoveToIdOk

`func (o *MoveItemRequest) GetMoveToIdOk() (*string, bool)`

GetMoveToIdOk returns a tuple with the MoveToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveToId

`func (o *MoveItemRequest) SetMoveToId(v string)`

SetMoveToId sets MoveToId field to given value.

### HasMoveToId

`func (o *MoveItemRequest) HasMoveToId() bool`

HasMoveToId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


