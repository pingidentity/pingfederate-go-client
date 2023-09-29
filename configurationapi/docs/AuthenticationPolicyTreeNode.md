# AuthenticationPolicyTreeNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**PolicyAction**](PolicyAction.md) |  | 
**Children** | Pointer to [**[]AuthenticationPolicyTreeNode**](AuthenticationPolicyTreeNode.md) | The nodes inside the authentication policy tree node of type AuthenticationPolicyTreeNode. | [optional] 

## Methods

### NewAuthenticationPolicyTreeNode

`func NewAuthenticationPolicyTreeNode(action PolicyAction, ) *AuthenticationPolicyTreeNode`

NewAuthenticationPolicyTreeNode instantiates a new AuthenticationPolicyTreeNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationPolicyTreeNodeWithDefaults

`func NewAuthenticationPolicyTreeNodeWithDefaults() *AuthenticationPolicyTreeNode`

NewAuthenticationPolicyTreeNodeWithDefaults instantiates a new AuthenticationPolicyTreeNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AuthenticationPolicyTreeNode) GetAction() PolicyAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AuthenticationPolicyTreeNode) GetActionOk() (*PolicyAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AuthenticationPolicyTreeNode) SetAction(v PolicyAction)`

SetAction sets Action field to given value.


### GetChildren

`func (o *AuthenticationPolicyTreeNode) GetChildren() []AuthenticationPolicyTreeNode`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *AuthenticationPolicyTreeNode) GetChildrenOk() (*[]AuthenticationPolicyTreeNode, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *AuthenticationPolicyTreeNode) SetChildren(v []AuthenticationPolicyTreeNode)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *AuthenticationPolicyTreeNode) HasChildren() bool`

HasChildren returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


