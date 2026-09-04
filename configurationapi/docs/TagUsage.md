# TagUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientCount** | Pointer to **int64** | The number of OAuth clients that reference this tag. | [optional] 
**ClientIds** | Pointer to **[]string** | IDs of OAuth clients that reference this tag. | [optional] 
**SelectorInstances** | Pointer to [**[]ReferenceItem**](ReferenceItem.md) | Authentication selector instances that reference this tag. Deletion is blocked while non-empty. | [optional] 
**CimdPolicies** | Pointer to [**[]ReferenceItem**](ReferenceItem.md) | CIMD policies that reference this tag. Deletion is blocked while non-empty. | [optional] 
**ClientSettingsUsed** | Pointer to **bool** | Whether OAuth client settings reference this tag. Deletion is blocked when true. | [optional] 

## Methods

### NewTagUsage

`func NewTagUsage() *TagUsage`

NewTagUsage instantiates a new TagUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagUsageWithDefaults

`func NewTagUsageWithDefaults() *TagUsage`

NewTagUsageWithDefaults instantiates a new TagUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientCount

`func (o *TagUsage) GetClientCount() int64`

GetClientCount returns the ClientCount field if non-nil, zero value otherwise.

### GetClientCountOk

`func (o *TagUsage) GetClientCountOk() (*int64, bool)`

GetClientCountOk returns a tuple with the ClientCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCount

`func (o *TagUsage) SetClientCount(v int64)`

SetClientCount sets ClientCount field to given value.

### HasClientCount

`func (o *TagUsage) HasClientCount() bool`

HasClientCount returns a boolean if a field has been set.

### GetClientIds

`func (o *TagUsage) GetClientIds() []string`

GetClientIds returns the ClientIds field if non-nil, zero value otherwise.

### GetClientIdsOk

`func (o *TagUsage) GetClientIdsOk() (*[]string, bool)`

GetClientIdsOk returns a tuple with the ClientIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIds

`func (o *TagUsage) SetClientIds(v []string)`

SetClientIds sets ClientIds field to given value.

### HasClientIds

`func (o *TagUsage) HasClientIds() bool`

HasClientIds returns a boolean if a field has been set.

### GetSelectorInstances

`func (o *TagUsage) GetSelectorInstances() []ReferenceItem`

GetSelectorInstances returns the SelectorInstances field if non-nil, zero value otherwise.

### GetSelectorInstancesOk

`func (o *TagUsage) GetSelectorInstancesOk() (*[]ReferenceItem, bool)`

GetSelectorInstancesOk returns a tuple with the SelectorInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectorInstances

`func (o *TagUsage) SetSelectorInstances(v []ReferenceItem)`

SetSelectorInstances sets SelectorInstances field to given value.

### HasSelectorInstances

`func (o *TagUsage) HasSelectorInstances() bool`

HasSelectorInstances returns a boolean if a field has been set.

### GetCimdPolicies

`func (o *TagUsage) GetCimdPolicies() []ReferenceItem`

GetCimdPolicies returns the CimdPolicies field if non-nil, zero value otherwise.

### GetCimdPoliciesOk

`func (o *TagUsage) GetCimdPoliciesOk() (*[]ReferenceItem, bool)`

GetCimdPoliciesOk returns a tuple with the CimdPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCimdPolicies

`func (o *TagUsage) SetCimdPolicies(v []ReferenceItem)`

SetCimdPolicies sets CimdPolicies field to given value.

### HasCimdPolicies

`func (o *TagUsage) HasCimdPolicies() bool`

HasCimdPolicies returns a boolean if a field has been set.

### GetClientSettingsUsed

`func (o *TagUsage) GetClientSettingsUsed() bool`

GetClientSettingsUsed returns the ClientSettingsUsed field if non-nil, zero value otherwise.

### GetClientSettingsUsedOk

`func (o *TagUsage) GetClientSettingsUsedOk() (*bool, bool)`

GetClientSettingsUsedOk returns a tuple with the ClientSettingsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSettingsUsed

`func (o *TagUsage) SetClientSettingsUsed(v bool)`

SetClientSettingsUsed sets ClientSettingsUsed field to given value.

### HasClientSettingsUsed

`func (o *TagUsage) HasClientSettingsUsed() bool`

HasClientSettingsUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


