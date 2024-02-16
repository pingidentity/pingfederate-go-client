# FragmentPolicyAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The authentication selection type. | 
**Context** | Pointer to **string** | The result context. | [optional] 
**AttributeRules** | Pointer to [**AttributeRules**](AttributeRules.md) |  | [optional] 
**Fragment** | [**ResourceLink**](ResourceLink.md) |  | 
**FragmentMapping** | Pointer to [**AttributeMapping**](AttributeMapping.md) |  | [optional] 

## Methods

### NewFragmentPolicyAction

`func NewFragmentPolicyAction(type_ string, fragment ResourceLink, ) *FragmentPolicyAction`

NewFragmentPolicyAction instantiates a new FragmentPolicyAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFragmentPolicyActionWithDefaults

`func NewFragmentPolicyActionWithDefaults() *FragmentPolicyAction`

NewFragmentPolicyActionWithDefaults instantiates a new FragmentPolicyAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FragmentPolicyAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FragmentPolicyAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FragmentPolicyAction) SetType(v string)`

SetType sets Type field to given value.


### GetContext

`func (o *FragmentPolicyAction) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *FragmentPolicyAction) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *FragmentPolicyAction) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *FragmentPolicyAction) HasContext() bool`

HasContext returns a boolean if a field has been set.

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


