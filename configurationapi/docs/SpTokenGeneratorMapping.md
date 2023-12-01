# SpTokenGeneratorMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpTokenGeneratorRef** | [**ResourceLink**](ResourceLink.md) |  | 
**RestrictedVirtualEntityIds** | Pointer to **[]string** | The list of virtual server IDs that this mapping is restricted to. | [optional] 
**DefaultMapping** | Pointer to **bool** | Indicates whether the token generator mapping is the default mapping. The default value is false. | [optional] 
**AttributeSources** | Pointer to [**[]AttributeSourceAggregation**](AttributeSourceAggregation.md) | A list of configured data stores to look up attributes from. | [optional] 
**AttributeContractFulfillment** | [**map[string]AttributeFulfillmentValue**](AttributeFulfillmentValue.md) | A list of mappings from attribute names to their fulfillment values. | 
**IssuanceCriteria** | Pointer to [**IssuanceCriteria**](IssuanceCriteria.md) |  | [optional] 

## Methods

### NewSpTokenGeneratorMapping

`func NewSpTokenGeneratorMapping(spTokenGeneratorRef ResourceLink, attributeContractFulfillment map[string]AttributeFulfillmentValue, ) *SpTokenGeneratorMapping`

NewSpTokenGeneratorMapping instantiates a new SpTokenGeneratorMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpTokenGeneratorMappingWithDefaults

`func NewSpTokenGeneratorMappingWithDefaults() *SpTokenGeneratorMapping`

NewSpTokenGeneratorMappingWithDefaults instantiates a new SpTokenGeneratorMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpTokenGeneratorRef

`func (o *SpTokenGeneratorMapping) GetSpTokenGeneratorRef() ResourceLink`

GetSpTokenGeneratorRef returns the SpTokenGeneratorRef field if non-nil, zero value otherwise.

### GetSpTokenGeneratorRefOk

`func (o *SpTokenGeneratorMapping) GetSpTokenGeneratorRefOk() (*ResourceLink, bool)`

GetSpTokenGeneratorRefOk returns a tuple with the SpTokenGeneratorRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpTokenGeneratorRef

`func (o *SpTokenGeneratorMapping) SetSpTokenGeneratorRef(v ResourceLink)`

SetSpTokenGeneratorRef sets SpTokenGeneratorRef field to given value.


### GetRestrictedVirtualEntityIds

`func (o *SpTokenGeneratorMapping) GetRestrictedVirtualEntityIds() []string`

GetRestrictedVirtualEntityIds returns the RestrictedVirtualEntityIds field if non-nil, zero value otherwise.

### GetRestrictedVirtualEntityIdsOk

`func (o *SpTokenGeneratorMapping) GetRestrictedVirtualEntityIdsOk() (*[]string, bool)`

GetRestrictedVirtualEntityIdsOk returns a tuple with the RestrictedVirtualEntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedVirtualEntityIds

`func (o *SpTokenGeneratorMapping) SetRestrictedVirtualEntityIds(v []string)`

SetRestrictedVirtualEntityIds sets RestrictedVirtualEntityIds field to given value.

### HasRestrictedVirtualEntityIds

`func (o *SpTokenGeneratorMapping) HasRestrictedVirtualEntityIds() bool`

HasRestrictedVirtualEntityIds returns a boolean if a field has been set.

### GetDefaultMapping

`func (o *SpTokenGeneratorMapping) GetDefaultMapping() bool`

GetDefaultMapping returns the DefaultMapping field if non-nil, zero value otherwise.

### GetDefaultMappingOk

`func (o *SpTokenGeneratorMapping) GetDefaultMappingOk() (*bool, bool)`

GetDefaultMappingOk returns a tuple with the DefaultMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMapping

`func (o *SpTokenGeneratorMapping) SetDefaultMapping(v bool)`

SetDefaultMapping sets DefaultMapping field to given value.

### HasDefaultMapping

`func (o *SpTokenGeneratorMapping) HasDefaultMapping() bool`

HasDefaultMapping returns a boolean if a field has been set.

### GetAttributeSources

`func (o *SpTokenGeneratorMapping) GetAttributeSources() []AttributeSourceAggregation`

GetAttributeSources returns the AttributeSources field if non-nil, zero value otherwise.

### GetAttributeSourcesOk

`func (o *SpTokenGeneratorMapping) GetAttributeSourcesOk() (*[]AttributeSourceAggregation, bool)`

GetAttributeSourcesOk returns a tuple with the AttributeSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeSources

`func (o *SpTokenGeneratorMapping) SetAttributeSources(v []AttributeSourceAggregation)`

SetAttributeSources sets AttributeSources field to given value.

### HasAttributeSources

`func (o *SpTokenGeneratorMapping) HasAttributeSources() bool`

HasAttributeSources returns a boolean if a field has been set.

### GetAttributeContractFulfillment

`func (o *SpTokenGeneratorMapping) GetAttributeContractFulfillment() map[string]AttributeFulfillmentValue`

GetAttributeContractFulfillment returns the AttributeContractFulfillment field if non-nil, zero value otherwise.

### GetAttributeContractFulfillmentOk

`func (o *SpTokenGeneratorMapping) GetAttributeContractFulfillmentOk() (*map[string]AttributeFulfillmentValue, bool)`

GetAttributeContractFulfillmentOk returns a tuple with the AttributeContractFulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContractFulfillment

`func (o *SpTokenGeneratorMapping) SetAttributeContractFulfillment(v map[string]AttributeFulfillmentValue)`

SetAttributeContractFulfillment sets AttributeContractFulfillment field to given value.


### GetIssuanceCriteria

`func (o *SpTokenGeneratorMapping) GetIssuanceCriteria() IssuanceCriteria`

GetIssuanceCriteria returns the IssuanceCriteria field if non-nil, zero value otherwise.

### GetIssuanceCriteriaOk

`func (o *SpTokenGeneratorMapping) GetIssuanceCriteriaOk() (*IssuanceCriteria, bool)`

GetIssuanceCriteriaOk returns a tuple with the IssuanceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceCriteria

`func (o *SpTokenGeneratorMapping) SetIssuanceCriteria(v IssuanceCriteria)`

SetIssuanceCriteria sets IssuanceCriteria field to given value.

### HasIssuanceCriteria

`func (o *SpTokenGeneratorMapping) HasIssuanceCriteria() bool`

HasIssuanceCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


