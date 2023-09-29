# ResourceUsages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to [**[]ResourceCategoryInfo**](ResourceCategoryInfo.md) | The static list of available resource categories. | [optional] 
**Items** | Pointer to [**[]ResourceUsage**](ResourceUsage.md) | The actual list of resource usages. | [optional] 

## Methods

### NewResourceUsages

`func NewResourceUsages() *ResourceUsages`

NewResourceUsages instantiates a new ResourceUsages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceUsagesWithDefaults

`func NewResourceUsagesWithDefaults() *ResourceUsages`

NewResourceUsagesWithDefaults instantiates a new ResourceUsages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *ResourceUsages) GetCategories() []ResourceCategoryInfo`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ResourceUsages) GetCategoriesOk() (*[]ResourceCategoryInfo, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ResourceUsages) SetCategories(v []ResourceCategoryInfo)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ResourceUsages) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetItems

`func (o *ResourceUsages) GetItems() []ResourceUsage`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ResourceUsages) GetItemsOk() (*[]ResourceUsage, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ResourceUsages) SetItems(v []ResourceUsage)`

SetItems sets Items field to given value.

### HasItems

`func (o *ResourceUsages) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


