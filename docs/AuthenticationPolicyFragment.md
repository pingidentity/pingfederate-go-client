# AuthenticationPolicyFragment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The authentication policy fragment ID. ID is unique. | [optional] 
**Name** | Pointer to **string** | The authentication policy fragment name. Name is unique. | [optional] 
**Description** | Pointer to **string** | A description for the authentication policy fragment. | [optional] 
**RootNode** | Pointer to [**AuthenticationPolicyTreeNode**](AuthenticationPolicyTreeNode.md) |  | [optional] 
**Inputs** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**Outputs** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 

## Methods

### NewAuthenticationPolicyFragment

`func NewAuthenticationPolicyFragment() *AuthenticationPolicyFragment`

NewAuthenticationPolicyFragment instantiates a new AuthenticationPolicyFragment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationPolicyFragmentWithDefaults

`func NewAuthenticationPolicyFragmentWithDefaults() *AuthenticationPolicyFragment`

NewAuthenticationPolicyFragmentWithDefaults instantiates a new AuthenticationPolicyFragment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthenticationPolicyFragment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthenticationPolicyFragment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthenticationPolicyFragment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthenticationPolicyFragment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AuthenticationPolicyFragment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthenticationPolicyFragment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthenticationPolicyFragment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthenticationPolicyFragment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AuthenticationPolicyFragment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthenticationPolicyFragment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthenticationPolicyFragment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthenticationPolicyFragment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRootNode

`func (o *AuthenticationPolicyFragment) GetRootNode() AuthenticationPolicyTreeNode`

GetRootNode returns the RootNode field if non-nil, zero value otherwise.

### GetRootNodeOk

`func (o *AuthenticationPolicyFragment) GetRootNodeOk() (*AuthenticationPolicyTreeNode, bool)`

GetRootNodeOk returns a tuple with the RootNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootNode

`func (o *AuthenticationPolicyFragment) SetRootNode(v AuthenticationPolicyTreeNode)`

SetRootNode sets RootNode field to given value.

### HasRootNode

`func (o *AuthenticationPolicyFragment) HasRootNode() bool`

HasRootNode returns a boolean if a field has been set.

### GetInputs

`func (o *AuthenticationPolicyFragment) GetInputs() ResourceLink`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *AuthenticationPolicyFragment) GetInputsOk() (*ResourceLink, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *AuthenticationPolicyFragment) SetInputs(v ResourceLink)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *AuthenticationPolicyFragment) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetOutputs

`func (o *AuthenticationPolicyFragment) GetOutputs() ResourceLink`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *AuthenticationPolicyFragment) GetOutputsOk() (*ResourceLink, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *AuthenticationPolicyFragment) SetOutputs(v ResourceLink)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *AuthenticationPolicyFragment) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


