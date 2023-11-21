# SpAdapterMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpAdapterRef** | [**ResourceLink**](ResourceLink.md) |  | 
**RestrictVirtualEntityIds** | Pointer to **bool** | Restricts this mapping to specific virtual entity IDs. | [optional] 
**RestrictedVirtualEntityIds** | Pointer to **[]string** | The list of virtual server IDs that this mapping is restricted to. | [optional] 
**AdapterOverrideSettings** | Pointer to [**SpAdapter**](SpAdapter.md) |  | [optional] 
**AttributeSources** | Pointer to [**[]AttributeSourceAggregation**](AttributeSourceAggregation.md) | A list of configured data stores to look up attributes from. | [optional] 
**AttributeContractFulfillment** | [**map[string]AttributeFulfillmentValue**](AttributeFulfillmentValue.md) | A list of mappings from attribute names to their fulfillment values. | 
**IssuanceCriteria** | Pointer to [**IssuanceCriteria**](IssuanceCriteria.md) |  | [optional] 

## Methods

### NewSpAdapterMapping

`func NewSpAdapterMapping(spAdapterRef ResourceLink, attributeContractFulfillment map[string]AttributeFulfillmentValue, ) *SpAdapterMapping`

NewSpAdapterMapping instantiates a new SpAdapterMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpAdapterMappingWithDefaults

`func NewSpAdapterMappingWithDefaults() *SpAdapterMapping`

NewSpAdapterMappingWithDefaults instantiates a new SpAdapterMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpAdapterRef

`func (o *SpAdapterMapping) GetSpAdapterRef() ResourceLink`

GetSpAdapterRef returns the SpAdapterRef field if non-nil, zero value otherwise.

### GetSpAdapterRefOk

`func (o *SpAdapterMapping) GetSpAdapterRefOk() (*ResourceLink, bool)`

GetSpAdapterRefOk returns a tuple with the SpAdapterRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpAdapterRef

`func (o *SpAdapterMapping) SetSpAdapterRef(v ResourceLink)`

SetSpAdapterRef sets SpAdapterRef field to given value.


### GetRestrictVirtualEntityIds

`func (o *SpAdapterMapping) GetRestrictVirtualEntityIds() bool`

GetRestrictVirtualEntityIds returns the RestrictVirtualEntityIds field if non-nil, zero value otherwise.

### GetRestrictVirtualEntityIdsOk

`func (o *SpAdapterMapping) GetRestrictVirtualEntityIdsOk() (*bool, bool)`

GetRestrictVirtualEntityIdsOk returns a tuple with the RestrictVirtualEntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictVirtualEntityIds

`func (o *SpAdapterMapping) SetRestrictVirtualEntityIds(v bool)`

SetRestrictVirtualEntityIds sets RestrictVirtualEntityIds field to given value.

### HasRestrictVirtualEntityIds

`func (o *SpAdapterMapping) HasRestrictVirtualEntityIds() bool`

HasRestrictVirtualEntityIds returns a boolean if a field has been set.

### GetRestrictedVirtualEntityIds

`func (o *SpAdapterMapping) GetRestrictedVirtualEntityIds() []string`

GetRestrictedVirtualEntityIds returns the RestrictedVirtualEntityIds field if non-nil, zero value otherwise.

### GetRestrictedVirtualEntityIdsOk

`func (o *SpAdapterMapping) GetRestrictedVirtualEntityIdsOk() (*[]string, bool)`

GetRestrictedVirtualEntityIdsOk returns a tuple with the RestrictedVirtualEntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedVirtualEntityIds

`func (o *SpAdapterMapping) SetRestrictedVirtualEntityIds(v []string)`

SetRestrictedVirtualEntityIds sets RestrictedVirtualEntityIds field to given value.

### HasRestrictedVirtualEntityIds

`func (o *SpAdapterMapping) HasRestrictedVirtualEntityIds() bool`

HasRestrictedVirtualEntityIds returns a boolean if a field has been set.

### GetAdapterOverrideSettings

`func (o *SpAdapterMapping) GetAdapterOverrideSettings() SpAdapter`

GetAdapterOverrideSettings returns the AdapterOverrideSettings field if non-nil, zero value otherwise.

### GetAdapterOverrideSettingsOk

`func (o *SpAdapterMapping) GetAdapterOverrideSettingsOk() (*SpAdapter, bool)`

GetAdapterOverrideSettingsOk returns a tuple with the AdapterOverrideSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterOverrideSettings

`func (o *SpAdapterMapping) SetAdapterOverrideSettings(v SpAdapter)`

SetAdapterOverrideSettings sets AdapterOverrideSettings field to given value.

### HasAdapterOverrideSettings

`func (o *SpAdapterMapping) HasAdapterOverrideSettings() bool`

HasAdapterOverrideSettings returns a boolean if a field has been set.

### GetAttributeSources

`func (o *SpAdapterMapping) GetAttributeSources() []AttributeSourceAggregation`

GetAttributeSources returns the AttributeSources field if non-nil, zero value otherwise.

### GetAttributeSourcesOk

`func (o *SpAdapterMapping) GetAttributeSourcesOk() (*[]AttributeSourceAggregation, bool)`

GetAttributeSourcesOk returns a tuple with the AttributeSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeSources

`func (o *SpAdapterMapping) SetAttributeSources(v []AttributeSourceAggregation)`

SetAttributeSources sets AttributeSources field to given value.

### HasAttributeSources

`func (o *SpAdapterMapping) HasAttributeSources() bool`

HasAttributeSources returns a boolean if a field has been set.

### GetAttributeContractFulfillment

`func (o *SpAdapterMapping) GetAttributeContractFulfillment() map[string]AttributeFulfillmentValue`

GetAttributeContractFulfillment returns the AttributeContractFulfillment field if non-nil, zero value otherwise.

### GetAttributeContractFulfillmentOk

`func (o *SpAdapterMapping) GetAttributeContractFulfillmentOk() (*map[string]AttributeFulfillmentValue, bool)`

GetAttributeContractFulfillmentOk returns a tuple with the AttributeContractFulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContractFulfillment

`func (o *SpAdapterMapping) SetAttributeContractFulfillment(v map[string]AttributeFulfillmentValue)`

SetAttributeContractFulfillment sets AttributeContractFulfillment field to given value.


### GetIssuanceCriteria

`func (o *SpAdapterMapping) GetIssuanceCriteria() IssuanceCriteria`

GetIssuanceCriteria returns the IssuanceCriteria field if non-nil, zero value otherwise.

### GetIssuanceCriteriaOk

`func (o *SpAdapterMapping) GetIssuanceCriteriaOk() (*IssuanceCriteria, bool)`

GetIssuanceCriteriaOk returns a tuple with the IssuanceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceCriteria

`func (o *SpAdapterMapping) SetIssuanceCriteria(v IssuanceCriteria)`

SetIssuanceCriteria sets IssuanceCriteria field to given value.

### HasIssuanceCriteria

`func (o *SpAdapterMapping) HasIssuanceCriteria() bool`

HasIssuanceCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


