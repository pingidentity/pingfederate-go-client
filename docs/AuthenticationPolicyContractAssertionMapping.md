# AuthenticationPolicyContractAssertionMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationPolicyContractRef** | [**ResourceLink**](ResourceLink.md) |  | 
**RestrictVirtualEntityIds** | Pointer to **bool** | Restricts this mapping to specific virtual entity IDs. | [optional] 
**RestrictedVirtualEntityIds** | Pointer to **[]string** | The list of virtual server IDs that this mapping is restricted to. | [optional] 
**AbortSsoTransactionAsFailSafe** | Pointer to **bool** | If set to true, SSO transaction will be aborted as a fail-safe when the data-store&#39;s attribute mappings fail to complete the attribute contract. Otherwise, the attribute contract with default values is used. By default, this value is false. | [optional] 
**AttributeSources** | Pointer to [**[]AttributeSource**](AttributeSource.md) | A list of configured data stores to look up attributes from. | [optional] 
**AttributeContractFulfillment** | [**map[string]AttributeFulfillmentValue**](AttributeFulfillmentValue.md) | A list of mappings from attribute names to their fulfillment values. | 
**IssuanceCriteria** | Pointer to [**IssuanceCriteria**](IssuanceCriteria.md) |  | [optional] 

## Methods

### NewAuthenticationPolicyContractAssertionMapping

`func NewAuthenticationPolicyContractAssertionMapping(authenticationPolicyContractRef ResourceLink, attributeContractFulfillment map[string]AttributeFulfillmentValue, ) *AuthenticationPolicyContractAssertionMapping`

NewAuthenticationPolicyContractAssertionMapping instantiates a new AuthenticationPolicyContractAssertionMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationPolicyContractAssertionMappingWithDefaults

`func NewAuthenticationPolicyContractAssertionMappingWithDefaults() *AuthenticationPolicyContractAssertionMapping`

NewAuthenticationPolicyContractAssertionMappingWithDefaults instantiates a new AuthenticationPolicyContractAssertionMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationPolicyContractRef

`func (o *AuthenticationPolicyContractAssertionMapping) GetAuthenticationPolicyContractRef() ResourceLink`

GetAuthenticationPolicyContractRef returns the AuthenticationPolicyContractRef field if non-nil, zero value otherwise.

### GetAuthenticationPolicyContractRefOk

`func (o *AuthenticationPolicyContractAssertionMapping) GetAuthenticationPolicyContractRefOk() (*ResourceLink, bool)`

GetAuthenticationPolicyContractRefOk returns a tuple with the AuthenticationPolicyContractRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationPolicyContractRef

`func (o *AuthenticationPolicyContractAssertionMapping) SetAuthenticationPolicyContractRef(v ResourceLink)`

SetAuthenticationPolicyContractRef sets AuthenticationPolicyContractRef field to given value.


### GetRestrictVirtualEntityIds

`func (o *AuthenticationPolicyContractAssertionMapping) GetRestrictVirtualEntityIds() bool`

GetRestrictVirtualEntityIds returns the RestrictVirtualEntityIds field if non-nil, zero value otherwise.

### GetRestrictVirtualEntityIdsOk

`func (o *AuthenticationPolicyContractAssertionMapping) GetRestrictVirtualEntityIdsOk() (*bool, bool)`

GetRestrictVirtualEntityIdsOk returns a tuple with the RestrictVirtualEntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictVirtualEntityIds

`func (o *AuthenticationPolicyContractAssertionMapping) SetRestrictVirtualEntityIds(v bool)`

SetRestrictVirtualEntityIds sets RestrictVirtualEntityIds field to given value.

### HasRestrictVirtualEntityIds

`func (o *AuthenticationPolicyContractAssertionMapping) HasRestrictVirtualEntityIds() bool`

HasRestrictVirtualEntityIds returns a boolean if a field has been set.

### GetRestrictedVirtualEntityIds

`func (o *AuthenticationPolicyContractAssertionMapping) GetRestrictedVirtualEntityIds() []string`

GetRestrictedVirtualEntityIds returns the RestrictedVirtualEntityIds field if non-nil, zero value otherwise.

### GetRestrictedVirtualEntityIdsOk

`func (o *AuthenticationPolicyContractAssertionMapping) GetRestrictedVirtualEntityIdsOk() (*[]string, bool)`

GetRestrictedVirtualEntityIdsOk returns a tuple with the RestrictedVirtualEntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedVirtualEntityIds

`func (o *AuthenticationPolicyContractAssertionMapping) SetRestrictedVirtualEntityIds(v []string)`

SetRestrictedVirtualEntityIds sets RestrictedVirtualEntityIds field to given value.

### HasRestrictedVirtualEntityIds

`func (o *AuthenticationPolicyContractAssertionMapping) HasRestrictedVirtualEntityIds() bool`

HasRestrictedVirtualEntityIds returns a boolean if a field has been set.

### GetAbortSsoTransactionAsFailSafe

`func (o *AuthenticationPolicyContractAssertionMapping) GetAbortSsoTransactionAsFailSafe() bool`

GetAbortSsoTransactionAsFailSafe returns the AbortSsoTransactionAsFailSafe field if non-nil, zero value otherwise.

### GetAbortSsoTransactionAsFailSafeOk

`func (o *AuthenticationPolicyContractAssertionMapping) GetAbortSsoTransactionAsFailSafeOk() (*bool, bool)`

GetAbortSsoTransactionAsFailSafeOk returns a tuple with the AbortSsoTransactionAsFailSafe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortSsoTransactionAsFailSafe

`func (o *AuthenticationPolicyContractAssertionMapping) SetAbortSsoTransactionAsFailSafe(v bool)`

SetAbortSsoTransactionAsFailSafe sets AbortSsoTransactionAsFailSafe field to given value.

### HasAbortSsoTransactionAsFailSafe

`func (o *AuthenticationPolicyContractAssertionMapping) HasAbortSsoTransactionAsFailSafe() bool`

HasAbortSsoTransactionAsFailSafe returns a boolean if a field has been set.

### GetAttributeSources

`func (o *AuthenticationPolicyContractAssertionMapping) GetAttributeSources() []AttributeSource`

GetAttributeSources returns the AttributeSources field if non-nil, zero value otherwise.

### GetAttributeSourcesOk

`func (o *AuthenticationPolicyContractAssertionMapping) GetAttributeSourcesOk() (*[]AttributeSource, bool)`

GetAttributeSourcesOk returns a tuple with the AttributeSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeSources

`func (o *AuthenticationPolicyContractAssertionMapping) SetAttributeSources(v []AttributeSource)`

SetAttributeSources sets AttributeSources field to given value.

### HasAttributeSources

`func (o *AuthenticationPolicyContractAssertionMapping) HasAttributeSources() bool`

HasAttributeSources returns a boolean if a field has been set.

### GetAttributeContractFulfillment

`func (o *AuthenticationPolicyContractAssertionMapping) GetAttributeContractFulfillment() map[string]AttributeFulfillmentValue`

GetAttributeContractFulfillment returns the AttributeContractFulfillment field if non-nil, zero value otherwise.

### GetAttributeContractFulfillmentOk

`func (o *AuthenticationPolicyContractAssertionMapping) GetAttributeContractFulfillmentOk() (*map[string]AttributeFulfillmentValue, bool)`

GetAttributeContractFulfillmentOk returns a tuple with the AttributeContractFulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContractFulfillment

`func (o *AuthenticationPolicyContractAssertionMapping) SetAttributeContractFulfillment(v map[string]AttributeFulfillmentValue)`

SetAttributeContractFulfillment sets AttributeContractFulfillment field to given value.


### GetIssuanceCriteria

`func (o *AuthenticationPolicyContractAssertionMapping) GetIssuanceCriteria() IssuanceCriteria`

GetIssuanceCriteria returns the IssuanceCriteria field if non-nil, zero value otherwise.

### GetIssuanceCriteriaOk

`func (o *AuthenticationPolicyContractAssertionMapping) GetIssuanceCriteriaOk() (*IssuanceCriteria, bool)`

GetIssuanceCriteriaOk returns a tuple with the IssuanceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceCriteria

`func (o *AuthenticationPolicyContractAssertionMapping) SetIssuanceCriteria(v IssuanceCriteria)`

SetIssuanceCriteria sets IssuanceCriteria field to given value.

### HasIssuanceCriteria

`func (o *AuthenticationPolicyContractAssertionMapping) HasIssuanceCriteria() bool`

HasIssuanceCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


