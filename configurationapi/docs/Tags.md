# Tags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Tag**](Tag.md) | The list of tags. | [optional] 
**TotalItems** | Pointer to **int64** | The total number of tags matching the filter (before pagination). | [optional] 

## Methods

### NewTags

`func NewTags() *Tags`

NewTags instantiates a new Tags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagsWithDefaults

`func NewTagsWithDefaults() *Tags`

NewTagsWithDefaults instantiates a new Tags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *Tags) GetItems() []Tag`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Tags) GetItemsOk() (*[]Tag, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Tags) SetItems(v []Tag)`

SetItems sets Items field to given value.

### HasItems

`func (o *Tags) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotalItems

`func (o *Tags) GetTotalItems() int64`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *Tags) GetTotalItemsOk() (*int64, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *Tags) SetTotalItems(v int64)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *Tags) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


