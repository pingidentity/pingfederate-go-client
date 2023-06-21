# ScopeGroupEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the scope group. | 
**Description** | **string** | The description of the scope group. | 
**Scopes** | **[]string** | The set of scopes for this scope group. | 

## Methods

### NewScopeGroupEntry

`func NewScopeGroupEntry(name string, description string, scopes []string, ) *ScopeGroupEntry`

NewScopeGroupEntry instantiates a new ScopeGroupEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopeGroupEntryWithDefaults

`func NewScopeGroupEntryWithDefaults() *ScopeGroupEntry`

NewScopeGroupEntryWithDefaults instantiates a new ScopeGroupEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ScopeGroupEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScopeGroupEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScopeGroupEntry) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ScopeGroupEntry) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScopeGroupEntry) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScopeGroupEntry) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetScopes

`func (o *ScopeGroupEntry) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ScopeGroupEntry) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ScopeGroupEntry) SetScopes(v []string)`

SetScopes sets Scopes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


