# AuthenticationPolicyContractMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationPolicyContractRef** | [**ResourceLink**](ResourceLink.md) |  | 
**RestrictVirtualServerIds** | Pointer to **bool** | Restricts this mapping to specific virtual entity IDs. | [optional] 
**RestrictedVirtualServerIds** | Pointer to **[]string** | The list of virtual server IDs that this mapping is restricted to. | [optional] 
**AttributeSources** | Pointer to [**[]AttributeSourceAggregation**](AttributeSourceAggregation.md) | A list of configured data stores to look up attributes from. | [optional] 
**AttributeContractFulfillment** | [**map[string]AttributeFulfillmentValue**](AttributeFulfillmentValue.md) | A list of mappings from attribute names to their fulfillment values. | 
**IssuanceCriteria** | Pointer to [**IssuanceCriteria**](IssuanceCriteria.md) |  | [optional] 

## Methods

### NewAuthenticationPolicyContractMapping

`func NewAuthenticationPolicyContractMapping(authenticationPolicyContractRef ResourceLink, attributeContractFulfillment map[string]AttributeFulfillmentValue, ) *AuthenticationPolicyContractMapping`

NewAuthenticationPolicyContractMapping instantiates a new AuthenticationPolicyContractMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationPolicyContractMappingWithDefaults

`func NewAuthenticationPolicyContractMappingWithDefaults() *AuthenticationPolicyContractMapping`

NewAuthenticationPolicyContractMappingWithDefaults instantiates a new AuthenticationPolicyContractMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationPolicyContractRef

`func (o *AuthenticationPolicyContractMapping) GetAuthenticationPolicyContractRef() ResourceLink`

GetAuthenticationPolicyContractRef returns the AuthenticationPolicyContractRef field if non-nil, zero value otherwise.

### GetAuthenticationPolicyContractRefOk

`func (o *AuthenticationPolicyContractMapping) GetAuthenticationPolicyContractRefOk() (*ResourceLink, bool)`

GetAuthenticationPolicyContractRefOk returns a tuple with the AuthenticationPolicyContractRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationPolicyContractRef

`func (o *AuthenticationPolicyContractMapping) SetAuthenticationPolicyContractRef(v ResourceLink)`

SetAuthenticationPolicyContractRef sets AuthenticationPolicyContractRef field to given value.


### GetRestrictVirtualServerIds

`func (o *AuthenticationPolicyContractMapping) GetRestrictVirtualServerIds() bool`

GetRestrictVirtualServerIds returns the RestrictVirtualServerIds field if non-nil, zero value otherwise.

### GetRestrictVirtualServerIdsOk

`func (o *AuthenticationPolicyContractMapping) GetRestrictVirtualServerIdsOk() (*bool, bool)`

GetRestrictVirtualServerIdsOk returns a tuple with the RestrictVirtualServerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictVirtualServerIds

`func (o *AuthenticationPolicyContractMapping) SetRestrictVirtualServerIds(v bool)`

SetRestrictVirtualServerIds sets RestrictVirtualServerIds field to given value.

### HasRestrictVirtualServerIds

`func (o *AuthenticationPolicyContractMapping) HasRestrictVirtualServerIds() bool`

HasRestrictVirtualServerIds returns a boolean if a field has been set.

### GetRestrictedVirtualServerIds

`func (o *AuthenticationPolicyContractMapping) GetRestrictedVirtualServerIds() []string`

GetRestrictedVirtualServerIds returns the RestrictedVirtualServerIds field if non-nil, zero value otherwise.

### GetRestrictedVirtualServerIdsOk

`func (o *AuthenticationPolicyContractMapping) GetRestrictedVirtualServerIdsOk() (*[]string, bool)`

GetRestrictedVirtualServerIdsOk returns a tuple with the RestrictedVirtualServerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedVirtualServerIds

`func (o *AuthenticationPolicyContractMapping) SetRestrictedVirtualServerIds(v []string)`

SetRestrictedVirtualServerIds sets RestrictedVirtualServerIds field to given value.

### HasRestrictedVirtualServerIds

`func (o *AuthenticationPolicyContractMapping) HasRestrictedVirtualServerIds() bool`

HasRestrictedVirtualServerIds returns a boolean if a field has been set.

### GetAttributeSources

`func (o *AuthenticationPolicyContractMapping) GetAttributeSources() []AttributeSourceAggregation`

GetAttributeSources returns the AttributeSources field if non-nil, zero value otherwise.

### GetAttributeSourcesOk

`func (o *AuthenticationPolicyContractMapping) GetAttributeSourcesOk() (*[]AttributeSourceAggregation, bool)`

GetAttributeSourcesOk returns a tuple with the AttributeSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeSources

`func (o *AuthenticationPolicyContractMapping) SetAttributeSources(v []AttributeSourceAggregation)`

SetAttributeSources sets AttributeSources field to given value.

### HasAttributeSources

`func (o *AuthenticationPolicyContractMapping) HasAttributeSources() bool`

HasAttributeSources returns a boolean if a field has been set.

### GetAttributeContractFulfillment

`func (o *AuthenticationPolicyContractMapping) GetAttributeContractFulfillment() map[string]AttributeFulfillmentValue`

GetAttributeContractFulfillment returns the AttributeContractFulfillment field if non-nil, zero value otherwise.

### GetAttributeContractFulfillmentOk

`func (o *AuthenticationPolicyContractMapping) GetAttributeContractFulfillmentOk() (*map[string]AttributeFulfillmentValue, bool)`

GetAttributeContractFulfillmentOk returns a tuple with the AttributeContractFulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContractFulfillment

`func (o *AuthenticationPolicyContractMapping) SetAttributeContractFulfillment(v map[string]AttributeFulfillmentValue)`

SetAttributeContractFulfillment sets AttributeContractFulfillment field to given value.


### GetIssuanceCriteria

`func (o *AuthenticationPolicyContractMapping) GetIssuanceCriteria() IssuanceCriteria`

GetIssuanceCriteria returns the IssuanceCriteria field if non-nil, zero value otherwise.

### GetIssuanceCriteriaOk

`func (o *AuthenticationPolicyContractMapping) GetIssuanceCriteriaOk() (*IssuanceCriteria, bool)`

GetIssuanceCriteriaOk returns a tuple with the IssuanceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceCriteria

`func (o *AuthenticationPolicyContractMapping) SetIssuanceCriteria(v IssuanceCriteria)`

SetIssuanceCriteria sets IssuanceCriteria field to given value.

### HasIssuanceCriteria

`func (o *AuthenticationPolicyContractMapping) HasIssuanceCriteria() bool`

HasIssuanceCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


