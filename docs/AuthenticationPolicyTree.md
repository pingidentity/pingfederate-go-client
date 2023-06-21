# AuthenticationPolicyTree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The authentication policy ID. ID is unique. | [optional] 
**Name** | Pointer to **string** | The authentication policy name. Name is unique. | [optional] 
**Description** | Pointer to **string** | A description for the authentication policy. | [optional] 
**AuthenticationApiApplicationRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Whether or not this authentication policy tree is enabled. Default is true. | [optional] 
**RootNode** | Pointer to [**AuthenticationPolicyTreeNode**](AuthenticationPolicyTreeNode.md) |  | [optional] 
**HandleFailuresLocally** | Pointer to **bool** | If a policy ends in failure keep the user local. | [optional] 

## Methods

### NewAuthenticationPolicyTree

`func NewAuthenticationPolicyTree() *AuthenticationPolicyTree`

NewAuthenticationPolicyTree instantiates a new AuthenticationPolicyTree object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationPolicyTreeWithDefaults

`func NewAuthenticationPolicyTreeWithDefaults() *AuthenticationPolicyTree`

NewAuthenticationPolicyTreeWithDefaults instantiates a new AuthenticationPolicyTree object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthenticationPolicyTree) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthenticationPolicyTree) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthenticationPolicyTree) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthenticationPolicyTree) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AuthenticationPolicyTree) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthenticationPolicyTree) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthenticationPolicyTree) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthenticationPolicyTree) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AuthenticationPolicyTree) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthenticationPolicyTree) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthenticationPolicyTree) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthenticationPolicyTree) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAuthenticationApiApplicationRef

`func (o *AuthenticationPolicyTree) GetAuthenticationApiApplicationRef() ResourceLink`

GetAuthenticationApiApplicationRef returns the AuthenticationApiApplicationRef field if non-nil, zero value otherwise.

### GetAuthenticationApiApplicationRefOk

`func (o *AuthenticationPolicyTree) GetAuthenticationApiApplicationRefOk() (*ResourceLink, bool)`

GetAuthenticationApiApplicationRefOk returns a tuple with the AuthenticationApiApplicationRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationApiApplicationRef

`func (o *AuthenticationPolicyTree) SetAuthenticationApiApplicationRef(v ResourceLink)`

SetAuthenticationApiApplicationRef sets AuthenticationApiApplicationRef field to given value.

### HasAuthenticationApiApplicationRef

`func (o *AuthenticationPolicyTree) HasAuthenticationApiApplicationRef() bool`

HasAuthenticationApiApplicationRef returns a boolean if a field has been set.

### GetEnabled

`func (o *AuthenticationPolicyTree) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AuthenticationPolicyTree) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AuthenticationPolicyTree) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AuthenticationPolicyTree) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRootNode

`func (o *AuthenticationPolicyTree) GetRootNode() AuthenticationPolicyTreeNode`

GetRootNode returns the RootNode field if non-nil, zero value otherwise.

### GetRootNodeOk

`func (o *AuthenticationPolicyTree) GetRootNodeOk() (*AuthenticationPolicyTreeNode, bool)`

GetRootNodeOk returns a tuple with the RootNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootNode

`func (o *AuthenticationPolicyTree) SetRootNode(v AuthenticationPolicyTreeNode)`

SetRootNode sets RootNode field to given value.

### HasRootNode

`func (o *AuthenticationPolicyTree) HasRootNode() bool`

HasRootNode returns a boolean if a field has been set.

### GetHandleFailuresLocally

`func (o *AuthenticationPolicyTree) GetHandleFailuresLocally() bool`

GetHandleFailuresLocally returns the HandleFailuresLocally field if non-nil, zero value otherwise.

### GetHandleFailuresLocallyOk

`func (o *AuthenticationPolicyTree) GetHandleFailuresLocallyOk() (*bool, bool)`

GetHandleFailuresLocallyOk returns a tuple with the HandleFailuresLocally field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandleFailuresLocally

`func (o *AuthenticationPolicyTree) SetHandleFailuresLocally(v bool)`

SetHandleFailuresLocally sets HandleFailuresLocally field to given value.

### HasHandleFailuresLocally

`func (o *AuthenticationPolicyTree) HasHandleFailuresLocally() bool`

HasHandleFailuresLocally returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


