# IdpAdapterAssertionMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdpAdapterRef** | [**ResourceLink**](ResourceLink.md) |  | 
**RestrictVirtualEntityIds** | Pointer to **bool** | Restricts this mapping to specific virtual entity IDs. | [optional] 
**RestrictedVirtualEntityIds** | Pointer to **[]string** | The list of virtual server IDs that this mapping is restricted to. | [optional] 
**AdapterOverrideSettings** | Pointer to [**IdpAdapter**](IdpAdapter.md) |  | [optional] 
**AbortSsoTransactionAsFailSafe** | Pointer to **bool** | If set to true, SSO transaction will be aborted as a fail-safe when the data-store&#39;s attribute mappings fail to complete the attribute contract. Otherwise, the attribute contract with default values is used. By default, this value is false. | [optional] 
**AttributeSources** | Pointer to [**[]AttributeSource**](AttributeSource.md) | A list of configured data stores to look up attributes from. | [optional] 
**AttributeContractFulfillment** | [**map[string]AttributeFulfillmentValue**](AttributeFulfillmentValue.md) | A list of mappings from attribute names to their fulfillment values. | 
**IssuanceCriteria** | Pointer to [**IssuanceCriteria**](IssuanceCriteria.md) |  | [optional] 

## Methods

### NewIdpAdapterAssertionMapping

`func NewIdpAdapterAssertionMapping(idpAdapterRef ResourceLink, attributeContractFulfillment map[string]AttributeFulfillmentValue, ) *IdpAdapterAssertionMapping`

NewIdpAdapterAssertionMapping instantiates a new IdpAdapterAssertionMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpAdapterAssertionMappingWithDefaults

`func NewIdpAdapterAssertionMappingWithDefaults() *IdpAdapterAssertionMapping`

NewIdpAdapterAssertionMappingWithDefaults instantiates a new IdpAdapterAssertionMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdpAdapterRef

`func (o *IdpAdapterAssertionMapping) GetIdpAdapterRef() ResourceLink`

GetIdpAdapterRef returns the IdpAdapterRef field if non-nil, zero value otherwise.

### GetIdpAdapterRefOk

`func (o *IdpAdapterAssertionMapping) GetIdpAdapterRefOk() (*ResourceLink, bool)`

GetIdpAdapterRefOk returns a tuple with the IdpAdapterRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpAdapterRef

`func (o *IdpAdapterAssertionMapping) SetIdpAdapterRef(v ResourceLink)`

SetIdpAdapterRef sets IdpAdapterRef field to given value.


### GetRestrictVirtualEntityIds

`func (o *IdpAdapterAssertionMapping) GetRestrictVirtualEntityIds() bool`

GetRestrictVirtualEntityIds returns the RestrictVirtualEntityIds field if non-nil, zero value otherwise.

### GetRestrictVirtualEntityIdsOk

`func (o *IdpAdapterAssertionMapping) GetRestrictVirtualEntityIdsOk() (*bool, bool)`

GetRestrictVirtualEntityIdsOk returns a tuple with the RestrictVirtualEntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictVirtualEntityIds

`func (o *IdpAdapterAssertionMapping) SetRestrictVirtualEntityIds(v bool)`

SetRestrictVirtualEntityIds sets RestrictVirtualEntityIds field to given value.

### HasRestrictVirtualEntityIds

`func (o *IdpAdapterAssertionMapping) HasRestrictVirtualEntityIds() bool`

HasRestrictVirtualEntityIds returns a boolean if a field has been set.

### GetRestrictedVirtualEntityIds

`func (o *IdpAdapterAssertionMapping) GetRestrictedVirtualEntityIds() []string`

GetRestrictedVirtualEntityIds returns the RestrictedVirtualEntityIds field if non-nil, zero value otherwise.

### GetRestrictedVirtualEntityIdsOk

`func (o *IdpAdapterAssertionMapping) GetRestrictedVirtualEntityIdsOk() (*[]string, bool)`

GetRestrictedVirtualEntityIdsOk returns a tuple with the RestrictedVirtualEntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedVirtualEntityIds

`func (o *IdpAdapterAssertionMapping) SetRestrictedVirtualEntityIds(v []string)`

SetRestrictedVirtualEntityIds sets RestrictedVirtualEntityIds field to given value.

### HasRestrictedVirtualEntityIds

`func (o *IdpAdapterAssertionMapping) HasRestrictedVirtualEntityIds() bool`

HasRestrictedVirtualEntityIds returns a boolean if a field has been set.

### GetAdapterOverrideSettings

`func (o *IdpAdapterAssertionMapping) GetAdapterOverrideSettings() IdpAdapter`

GetAdapterOverrideSettings returns the AdapterOverrideSettings field if non-nil, zero value otherwise.

### GetAdapterOverrideSettingsOk

`func (o *IdpAdapterAssertionMapping) GetAdapterOverrideSettingsOk() (*IdpAdapter, bool)`

GetAdapterOverrideSettingsOk returns a tuple with the AdapterOverrideSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterOverrideSettings

`func (o *IdpAdapterAssertionMapping) SetAdapterOverrideSettings(v IdpAdapter)`

SetAdapterOverrideSettings sets AdapterOverrideSettings field to given value.

### HasAdapterOverrideSettings

`func (o *IdpAdapterAssertionMapping) HasAdapterOverrideSettings() bool`

HasAdapterOverrideSettings returns a boolean if a field has been set.

### GetAbortSsoTransactionAsFailSafe

`func (o *IdpAdapterAssertionMapping) GetAbortSsoTransactionAsFailSafe() bool`

GetAbortSsoTransactionAsFailSafe returns the AbortSsoTransactionAsFailSafe field if non-nil, zero value otherwise.

### GetAbortSsoTransactionAsFailSafeOk

`func (o *IdpAdapterAssertionMapping) GetAbortSsoTransactionAsFailSafeOk() (*bool, bool)`

GetAbortSsoTransactionAsFailSafeOk returns a tuple with the AbortSsoTransactionAsFailSafe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortSsoTransactionAsFailSafe

`func (o *IdpAdapterAssertionMapping) SetAbortSsoTransactionAsFailSafe(v bool)`

SetAbortSsoTransactionAsFailSafe sets AbortSsoTransactionAsFailSafe field to given value.

### HasAbortSsoTransactionAsFailSafe

`func (o *IdpAdapterAssertionMapping) HasAbortSsoTransactionAsFailSafe() bool`

HasAbortSsoTransactionAsFailSafe returns a boolean if a field has been set.

### GetAttributeSources

`func (o *IdpAdapterAssertionMapping) GetAttributeSources() []AttributeSource`

GetAttributeSources returns the AttributeSources field if non-nil, zero value otherwise.

### GetAttributeSourcesOk

`func (o *IdpAdapterAssertionMapping) GetAttributeSourcesOk() (*[]AttributeSource, bool)`

GetAttributeSourcesOk returns a tuple with the AttributeSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeSources

`func (o *IdpAdapterAssertionMapping) SetAttributeSources(v []AttributeSource)`

SetAttributeSources sets AttributeSources field to given value.

### HasAttributeSources

`func (o *IdpAdapterAssertionMapping) HasAttributeSources() bool`

HasAttributeSources returns a boolean if a field has been set.

### GetAttributeContractFulfillment

`func (o *IdpAdapterAssertionMapping) GetAttributeContractFulfillment() map[string]AttributeFulfillmentValue`

GetAttributeContractFulfillment returns the AttributeContractFulfillment field if non-nil, zero value otherwise.

### GetAttributeContractFulfillmentOk

`func (o *IdpAdapterAssertionMapping) GetAttributeContractFulfillmentOk() (*map[string]AttributeFulfillmentValue, bool)`

GetAttributeContractFulfillmentOk returns a tuple with the AttributeContractFulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContractFulfillment

`func (o *IdpAdapterAssertionMapping) SetAttributeContractFulfillment(v map[string]AttributeFulfillmentValue)`

SetAttributeContractFulfillment sets AttributeContractFulfillment field to given value.


### GetIssuanceCriteria

`func (o *IdpAdapterAssertionMapping) GetIssuanceCriteria() IssuanceCriteria`

GetIssuanceCriteria returns the IssuanceCriteria field if non-nil, zero value otherwise.

### GetIssuanceCriteriaOk

`func (o *IdpAdapterAssertionMapping) GetIssuanceCriteriaOk() (*IssuanceCriteria, bool)`

GetIssuanceCriteriaOk returns a tuple with the IssuanceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceCriteria

`func (o *IdpAdapterAssertionMapping) SetIssuanceCriteria(v IssuanceCriteria)`

SetIssuanceCriteria sets IssuanceCriteria field to given value.

### HasIssuanceCriteria

`func (o *IdpAdapterAssertionMapping) HasIssuanceCriteria() bool`

HasIssuanceCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


