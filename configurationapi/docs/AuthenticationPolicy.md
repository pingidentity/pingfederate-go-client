# AuthenticationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailIfNoSelection** | Pointer to **bool** | Fail if policy finds no authentication source. | [optional] 
**AuthnSelectionTrees** | Pointer to [**[]AuthenticationPolicyTree**](AuthenticationPolicyTree.md) | The list of authentication policy trees. | [optional] 
**DefaultAuthenticationSources** | Pointer to [**[]AuthenticationSource**](AuthenticationSource.md) | The default authentication sources. | [optional] 
**TrackedHttpParameters** | Pointer to **[]string** | The HTTP request parameters to track and make available to authentication sources, selectors, and contract mappings throughout the authentication policy. | [optional] 

## Methods

### NewAuthenticationPolicy

`func NewAuthenticationPolicy() *AuthenticationPolicy`

NewAuthenticationPolicy instantiates a new AuthenticationPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationPolicyWithDefaults

`func NewAuthenticationPolicyWithDefaults() *AuthenticationPolicy`

NewAuthenticationPolicyWithDefaults instantiates a new AuthenticationPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailIfNoSelection

`func (o *AuthenticationPolicy) GetFailIfNoSelection() bool`

GetFailIfNoSelection returns the FailIfNoSelection field if non-nil, zero value otherwise.

### GetFailIfNoSelectionOk

`func (o *AuthenticationPolicy) GetFailIfNoSelectionOk() (*bool, bool)`

GetFailIfNoSelectionOk returns a tuple with the FailIfNoSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailIfNoSelection

`func (o *AuthenticationPolicy) SetFailIfNoSelection(v bool)`

SetFailIfNoSelection sets FailIfNoSelection field to given value.

### HasFailIfNoSelection

`func (o *AuthenticationPolicy) HasFailIfNoSelection() bool`

HasFailIfNoSelection returns a boolean if a field has been set.

### GetAuthnSelectionTrees

`func (o *AuthenticationPolicy) GetAuthnSelectionTrees() []AuthenticationPolicyTree`

GetAuthnSelectionTrees returns the AuthnSelectionTrees field if non-nil, zero value otherwise.

### GetAuthnSelectionTreesOk

`func (o *AuthenticationPolicy) GetAuthnSelectionTreesOk() (*[]AuthenticationPolicyTree, bool)`

GetAuthnSelectionTreesOk returns a tuple with the AuthnSelectionTrees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnSelectionTrees

`func (o *AuthenticationPolicy) SetAuthnSelectionTrees(v []AuthenticationPolicyTree)`

SetAuthnSelectionTrees sets AuthnSelectionTrees field to given value.

### HasAuthnSelectionTrees

`func (o *AuthenticationPolicy) HasAuthnSelectionTrees() bool`

HasAuthnSelectionTrees returns a boolean if a field has been set.

### GetDefaultAuthenticationSources

`func (o *AuthenticationPolicy) GetDefaultAuthenticationSources() []AuthenticationSource`

GetDefaultAuthenticationSources returns the DefaultAuthenticationSources field if non-nil, zero value otherwise.

### GetDefaultAuthenticationSourcesOk

`func (o *AuthenticationPolicy) GetDefaultAuthenticationSourcesOk() (*[]AuthenticationSource, bool)`

GetDefaultAuthenticationSourcesOk returns a tuple with the DefaultAuthenticationSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAuthenticationSources

`func (o *AuthenticationPolicy) SetDefaultAuthenticationSources(v []AuthenticationSource)`

SetDefaultAuthenticationSources sets DefaultAuthenticationSources field to given value.

### HasDefaultAuthenticationSources

`func (o *AuthenticationPolicy) HasDefaultAuthenticationSources() bool`

HasDefaultAuthenticationSources returns a boolean if a field has been set.

### GetTrackedHttpParameters

`func (o *AuthenticationPolicy) GetTrackedHttpParameters() []string`

GetTrackedHttpParameters returns the TrackedHttpParameters field if non-nil, zero value otherwise.

### GetTrackedHttpParametersOk

`func (o *AuthenticationPolicy) GetTrackedHttpParametersOk() (*[]string, bool)`

GetTrackedHttpParametersOk returns a tuple with the TrackedHttpParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackedHttpParameters

`func (o *AuthenticationPolicy) SetTrackedHttpParameters(v []string)`

SetTrackedHttpParameters sets TrackedHttpParameters field to given value.

### HasTrackedHttpParameters

`func (o *AuthenticationPolicy) HasTrackedHttpParameters() bool`

HasTrackedHttpParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


