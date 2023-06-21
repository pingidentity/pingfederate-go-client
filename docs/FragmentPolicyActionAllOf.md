# FragmentPolicyActionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeRules** | Pointer to [**AttributeRules**](AttributeRules.md) |  | [optional] 
**Fragment** | [**ResourceLink**](ResourceLink.md) |  | 
**FragmentMapping** | Pointer to [**AttributeMapping**](AttributeMapping.md) |  | [optional] 

## Methods

### NewFragmentPolicyActionAllOf

`func NewFragmentPolicyActionAllOf(fragment ResourceLink, ) *FragmentPolicyActionAllOf`

NewFragmentPolicyActionAllOf instantiates a new FragmentPolicyActionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFragmentPolicyActionAllOfWithDefaults

`func NewFragmentPolicyActionAllOfWithDefaults() *FragmentPolicyActionAllOf`

NewFragmentPolicyActionAllOfWithDefaults instantiates a new FragmentPolicyActionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeRules

`func (o *FragmentPolicyActionAllOf) GetAttributeRules() AttributeRules`

GetAttributeRules returns the AttributeRules field if non-nil, zero value otherwise.

### GetAttributeRulesOk

`func (o *FragmentPolicyActionAllOf) GetAttributeRulesOk() (*AttributeRules, bool)`

GetAttributeRulesOk returns a tuple with the AttributeRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeRules

`func (o *FragmentPolicyActionAllOf) SetAttributeRules(v AttributeRules)`

SetAttributeRules sets AttributeRules field to given value.

### HasAttributeRules

`func (o *FragmentPolicyActionAllOf) HasAttributeRules() bool`

HasAttributeRules returns a boolean if a field has been set.

### GetFragment

`func (o *FragmentPolicyActionAllOf) GetFragment() ResourceLink`

GetFragment returns the Fragment field if non-nil, zero value otherwise.

### GetFragmentOk

`func (o *FragmentPolicyActionAllOf) GetFragmentOk() (*ResourceLink, bool)`

GetFragmentOk returns a tuple with the Fragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragment

`func (o *FragmentPolicyActionAllOf) SetFragment(v ResourceLink)`

SetFragment sets Fragment field to given value.


### GetFragmentMapping

`func (o *FragmentPolicyActionAllOf) GetFragmentMapping() AttributeMapping`

GetFragmentMapping returns the FragmentMapping field if non-nil, zero value otherwise.

### GetFragmentMappingOk

`func (o *FragmentPolicyActionAllOf) GetFragmentMappingOk() (*AttributeMapping, bool)`

GetFragmentMappingOk returns a tuple with the FragmentMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragmentMapping

`func (o *FragmentPolicyActionAllOf) SetFragmentMapping(v AttributeMapping)`

SetFragmentMapping sets FragmentMapping field to given value.

### HasFragmentMapping

`func (o *FragmentPolicyActionAllOf) HasFragmentMapping() bool`

HasFragmentMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


