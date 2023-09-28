# ScopeEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the scope. | 
**Description** | **string** | The description of the scope that appears when the user is prompted for authorization. | 
**Dynamic** | Pointer to **bool** | True if the scope is dynamic. (Defaults to false) | [optional] 

## Methods

### NewScopeEntry

`func NewScopeEntry(name string, description string, ) *ScopeEntry`

NewScopeEntry instantiates a new ScopeEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopeEntryWithDefaults

`func NewScopeEntryWithDefaults() *ScopeEntry`

NewScopeEntryWithDefaults instantiates a new ScopeEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ScopeEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScopeEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScopeEntry) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ScopeEntry) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScopeEntry) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScopeEntry) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDynamic

`func (o *ScopeEntry) GetDynamic() bool`

GetDynamic returns the Dynamic field if non-nil, zero value otherwise.

### GetDynamicOk

`func (o *ScopeEntry) GetDynamicOk() (*bool, bool)`

GetDynamicOk returns a tuple with the Dynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic

`func (o *ScopeEntry) SetDynamic(v bool)`

SetDynamic sets Dynamic field to given value.

### HasDynamic

`func (o *ScopeEntry) HasDynamic() bool`

HasDynamic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


