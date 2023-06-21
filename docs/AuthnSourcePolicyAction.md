# AuthnSourcePolicyAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeRules** | Pointer to [**AttributeRules**](AttributeRules.md) |  | [optional] 
**AuthenticationSource** | [**AuthenticationSource**](AuthenticationSource.md) |  | 
**InputUserIdMapping** | Pointer to [**AttributeFulfillmentValue**](AttributeFulfillmentValue.md) |  | [optional] 
**UserIdAuthenticated** | Pointer to **bool** | Indicates whether the user ID obtained by the user ID mapping is authenticated. | [optional] 

## Methods

### NewAuthnSourcePolicyAction

`func NewAuthnSourcePolicyAction(authenticationSource AuthenticationSource, ) *AuthnSourcePolicyAction`

NewAuthnSourcePolicyAction instantiates a new AuthnSourcePolicyAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthnSourcePolicyActionWithDefaults

`func NewAuthnSourcePolicyActionWithDefaults() *AuthnSourcePolicyAction`

NewAuthnSourcePolicyActionWithDefaults instantiates a new AuthnSourcePolicyAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeRules

`func (o *AuthnSourcePolicyAction) GetAttributeRules() AttributeRules`

GetAttributeRules returns the AttributeRules field if non-nil, zero value otherwise.

### GetAttributeRulesOk

`func (o *AuthnSourcePolicyAction) GetAttributeRulesOk() (*AttributeRules, bool)`

GetAttributeRulesOk returns a tuple with the AttributeRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeRules

`func (o *AuthnSourcePolicyAction) SetAttributeRules(v AttributeRules)`

SetAttributeRules sets AttributeRules field to given value.

### HasAttributeRules

`func (o *AuthnSourcePolicyAction) HasAttributeRules() bool`

HasAttributeRules returns a boolean if a field has been set.

### GetAuthenticationSource

`func (o *AuthnSourcePolicyAction) GetAuthenticationSource() AuthenticationSource`

GetAuthenticationSource returns the AuthenticationSource field if non-nil, zero value otherwise.

### GetAuthenticationSourceOk

`func (o *AuthnSourcePolicyAction) GetAuthenticationSourceOk() (*AuthenticationSource, bool)`

GetAuthenticationSourceOk returns a tuple with the AuthenticationSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationSource

`func (o *AuthnSourcePolicyAction) SetAuthenticationSource(v AuthenticationSource)`

SetAuthenticationSource sets AuthenticationSource field to given value.


### GetInputUserIdMapping

`func (o *AuthnSourcePolicyAction) GetInputUserIdMapping() AttributeFulfillmentValue`

GetInputUserIdMapping returns the InputUserIdMapping field if non-nil, zero value otherwise.

### GetInputUserIdMappingOk

`func (o *AuthnSourcePolicyAction) GetInputUserIdMappingOk() (*AttributeFulfillmentValue, bool)`

GetInputUserIdMappingOk returns a tuple with the InputUserIdMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputUserIdMapping

`func (o *AuthnSourcePolicyAction) SetInputUserIdMapping(v AttributeFulfillmentValue)`

SetInputUserIdMapping sets InputUserIdMapping field to given value.

### HasInputUserIdMapping

`func (o *AuthnSourcePolicyAction) HasInputUserIdMapping() bool`

HasInputUserIdMapping returns a boolean if a field has been set.

### GetUserIdAuthenticated

`func (o *AuthnSourcePolicyAction) GetUserIdAuthenticated() bool`

GetUserIdAuthenticated returns the UserIdAuthenticated field if non-nil, zero value otherwise.

### GetUserIdAuthenticatedOk

`func (o *AuthnSourcePolicyAction) GetUserIdAuthenticatedOk() (*bool, bool)`

GetUserIdAuthenticatedOk returns a tuple with the UserIdAuthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdAuthenticated

`func (o *AuthnSourcePolicyAction) SetUserIdAuthenticated(v bool)`

SetUserIdAuthenticated sets UserIdAuthenticated field to given value.

### HasUserIdAuthenticated

`func (o *AuthnSourcePolicyAction) HasUserIdAuthenticated() bool`

HasUserIdAuthenticated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


