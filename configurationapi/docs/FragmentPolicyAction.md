# FragmentPolicyAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeRules** | Pointer to [**AttributeRules**](AttributeRules.md) |  | [optional] 
**Fragment** | [**ResourceLink**](ResourceLink.md) |  | 
**FragmentMapping** | Pointer to [**AttributeMapping**](AttributeMapping.md) |  | [optional] 

## Methods

### NewFragmentPolicyAction

`func NewFragmentPolicyAction(fragment ResourceLink, ) *FragmentPolicyAction`

NewFragmentPolicyAction instantiates a new FragmentPolicyAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFragmentPolicyActionWithDefaults

`func NewFragmentPolicyActionWithDefaults() *FragmentPolicyAction`

NewFragmentPolicyActionWithDefaults instantiates a new FragmentPolicyAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeRules

`func (o *FragmentPolicyAction) GetAttributeRules() AttributeRules`

GetAttributeRules returns the AttributeRules field if non-nil, zero value otherwise.

### GetAttributeRulesOk

`func (o *FragmentPolicyAction) GetAttributeRulesOk() (*AttributeRules, bool)`

GetAttributeRulesOk returns a tuple with the AttributeRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeRules

`func (o *FragmentPolicyAction) SetAttributeRules(v AttributeRules)`

SetAttributeRules sets AttributeRules field to given value.

### HasAttributeRules

`func (o *FragmentPolicyAction) HasAttributeRules() bool`

HasAttributeRules returns a boolean if a field has been set.

### GetFragment

`func (o *FragmentPolicyAction) GetFragment() ResourceLink`

GetFragment returns the Fragment field if non-nil, zero value otherwise.

### GetFragmentOk

`func (o *FragmentPolicyAction) GetFragmentOk() (*ResourceLink, bool)`

GetFragmentOk returns a tuple with the Fragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragment

`func (o *FragmentPolicyAction) SetFragment(v ResourceLink)`

SetFragment sets Fragment field to given value.


### GetFragmentMapping

`func (o *FragmentPolicyAction) GetFragmentMapping() AttributeMapping`

GetFragmentMapping returns the FragmentMapping field if non-nil, zero value otherwise.

### GetFragmentMappingOk

`func (o *FragmentPolicyAction) GetFragmentMappingOk() (*AttributeMapping, bool)`

GetFragmentMappingOk returns a tuple with the FragmentMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragmentMapping

`func (o *FragmentPolicyAction) SetFragmentMapping(v AttributeMapping)`

SetFragmentMapping sets FragmentMapping field to given value.

### HasFragmentMapping

`func (o *FragmentPolicyAction) HasFragmentMapping() bool`

HasFragmentMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


