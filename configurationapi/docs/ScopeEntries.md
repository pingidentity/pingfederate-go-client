# ScopeEntries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]ScopeEntry**](ScopeEntry.md) | The list of scopes and their descriptions. | [optional] 
**TotalItems** | Pointer to **int64** | The total number of scopes. | [optional] 

## Methods

### NewScopeEntries

`func NewScopeEntries() *ScopeEntries`

NewScopeEntries instantiates a new ScopeEntries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopeEntriesWithDefaults

`func NewScopeEntriesWithDefaults() *ScopeEntries`

NewScopeEntriesWithDefaults instantiates a new ScopeEntries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ScopeEntries) GetItems() []ScopeEntry`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ScopeEntries) GetItemsOk() (*[]ScopeEntry, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ScopeEntries) SetItems(v []ScopeEntry)`

SetItems sets Items field to given value.

### HasItems

`func (o *ScopeEntries) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotalItems

`func (o *ScopeEntries) GetTotalItems() int64`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *ScopeEntries) GetTotalItemsOk() (*int64, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *ScopeEntries) SetTotalItems(v int64)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *ScopeEntries) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


