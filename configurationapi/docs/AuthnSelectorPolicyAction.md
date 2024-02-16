# AuthnSelectorPolicyAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The authentication selection type. | 
**Context** | Pointer to **string** | The result context. | [optional] 
**AuthenticationSelectorRef** | [**ResourceLink**](ResourceLink.md) |  | 

## Methods

### NewAuthnSelectorPolicyAction

`func NewAuthnSelectorPolicyAction(type_ string, authenticationSelectorRef ResourceLink, ) *AuthnSelectorPolicyAction`

NewAuthnSelectorPolicyAction instantiates a new AuthnSelectorPolicyAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthnSelectorPolicyActionWithDefaults

`func NewAuthnSelectorPolicyActionWithDefaults() *AuthnSelectorPolicyAction`

NewAuthnSelectorPolicyActionWithDefaults instantiates a new AuthnSelectorPolicyAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AuthnSelectorPolicyAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthnSelectorPolicyAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthnSelectorPolicyAction) SetType(v string)`

SetType sets Type field to given value.


### GetContext

`func (o *AuthnSelectorPolicyAction) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *AuthnSelectorPolicyAction) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *AuthnSelectorPolicyAction) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *AuthnSelectorPolicyAction) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetAuthenticationSelectorRef

`func (o *AuthnSelectorPolicyAction) GetAuthenticationSelectorRef() ResourceLink`

GetAuthenticationSelectorRef returns the AuthenticationSelectorRef field if non-nil, zero value otherwise.

### GetAuthenticationSelectorRefOk

`func (o *AuthnSelectorPolicyAction) GetAuthenticationSelectorRefOk() (*ResourceLink, bool)`

GetAuthenticationSelectorRefOk returns a tuple with the AuthenticationSelectorRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationSelectorRef

`func (o *AuthnSelectorPolicyAction) SetAuthenticationSelectorRef(v ResourceLink)`

SetAuthenticationSelectorRef sets AuthenticationSelectorRef field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


