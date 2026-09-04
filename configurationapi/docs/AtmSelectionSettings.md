# AtmSelectionSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inherited** | Pointer to **bool** | If this token manager has a parent, this flag determines whether selection settings, such as resource URI&#39;s, are inherited from the parent. When set to true, the other fields in this model become read-only. The default value is false. | [optional] 
**ResourceUris** | Pointer to **[]string** | The list of base resource URI&#39;s which map to this token manager. A resource URI, specified via the &#39;aud&#39; parameter, can be used to select a specific token manager for an OAuth request. | [optional] 

## Methods

### NewAtmSelectionSettings

`func NewAtmSelectionSettings() *AtmSelectionSettings`

NewAtmSelectionSettings instantiates a new AtmSelectionSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtmSelectionSettingsWithDefaults

`func NewAtmSelectionSettingsWithDefaults() *AtmSelectionSettings`

NewAtmSelectionSettingsWithDefaults instantiates a new AtmSelectionSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInherited

`func (o *AtmSelectionSettings) GetInherited() bool`

GetInherited returns the Inherited field if non-nil, zero value otherwise.

### GetInheritedOk

`func (o *AtmSelectionSettings) GetInheritedOk() (*bool, bool)`

GetInheritedOk returns a tuple with the Inherited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInherited

`func (o *AtmSelectionSettings) SetInherited(v bool)`

SetInherited sets Inherited field to given value.

### HasInherited

`func (o *AtmSelectionSettings) HasInherited() bool`

HasInherited returns a boolean if a field has been set.

### GetResourceUris

`func (o *AtmSelectionSettings) GetResourceUris() []string`

GetResourceUris returns the ResourceUris field if non-nil, zero value otherwise.

### GetResourceUrisOk

`func (o *AtmSelectionSettings) GetResourceUrisOk() (*[]string, bool)`

GetResourceUrisOk returns a tuple with the ResourceUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUris

`func (o *AtmSelectionSettings) SetResourceUris(v []string)`

SetResourceUris sets ResourceUris field to given value.

### HasResourceUris

`func (o *AtmSelectionSettings) HasResourceUris() bool`

HasResourceUris returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


