# ScopeGroupEntries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]ScopeGroupEntry**](ScopeGroupEntry.md) | The list of scope groups and their descriptions. | [optional] 
**TotalItems** | Pointer to **int64** | The total number of scope groups. | [optional] 

## Methods

### NewScopeGroupEntries

`func NewScopeGroupEntries() *ScopeGroupEntries`

NewScopeGroupEntries instantiates a new ScopeGroupEntries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopeGroupEntriesWithDefaults

`func NewScopeGroupEntriesWithDefaults() *ScopeGroupEntries`

NewScopeGroupEntriesWithDefaults instantiates a new ScopeGroupEntries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ScopeGroupEntries) GetItems() []ScopeGroupEntry`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ScopeGroupEntries) GetItemsOk() (*[]ScopeGroupEntry, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ScopeGroupEntries) SetItems(v []ScopeGroupEntry)`

SetItems sets Items field to given value.

### HasItems

`func (o *ScopeGroupEntries) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotalItems

`func (o *ScopeGroupEntries) GetTotalItems() int64`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *ScopeGroupEntries) GetTotalItemsOk() (*int64, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *ScopeGroupEntries) SetTotalItems(v int64)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *ScopeGroupEntries) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


