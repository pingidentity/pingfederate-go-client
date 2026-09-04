# SpOAuthProcessorPolicyMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeSources** | Pointer to [**[]AttributeSourceAggregation**](AttributeSourceAggregation.md) | A list of configured data stores to look up attributes from. | [optional] 
**AttributeContractFulfillment** | [**map[string]AttributeFulfillmentValue**](AttributeFulfillmentValue.md) | A list of mappings from attribute names to their fulfillment values. | 
**IssuanceCriteria** | Pointer to [**IssuanceCriteria**](IssuanceCriteria.md) |  | [optional] 
**ProcessorPolicyRef** | [**ResourceLink**](ResourceLink.md) |  | 
**AttributeSource** | Pointer to [**[]AttributeSourceAggregation**](AttributeSourceAggregation.md) |  | [optional] 

## Methods

### NewSpOAuthProcessorPolicyMapping

`func NewSpOAuthProcessorPolicyMapping(attributeContractFulfillment map[string]AttributeFulfillmentValue, processorPolicyRef ResourceLink, ) *SpOAuthProcessorPolicyMapping`

NewSpOAuthProcessorPolicyMapping instantiates a new SpOAuthProcessorPolicyMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpOAuthProcessorPolicyMappingWithDefaults

`func NewSpOAuthProcessorPolicyMappingWithDefaults() *SpOAuthProcessorPolicyMapping`

NewSpOAuthProcessorPolicyMappingWithDefaults instantiates a new SpOAuthProcessorPolicyMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeSources

`func (o *SpOAuthProcessorPolicyMapping) GetAttributeSources() []AttributeSourceAggregation`

GetAttributeSources returns the AttributeSources field if non-nil, zero value otherwise.

### GetAttributeSourcesOk

`func (o *SpOAuthProcessorPolicyMapping) GetAttributeSourcesOk() (*[]AttributeSourceAggregation, bool)`

GetAttributeSourcesOk returns a tuple with the AttributeSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeSources

`func (o *SpOAuthProcessorPolicyMapping) SetAttributeSources(v []AttributeSourceAggregation)`

SetAttributeSources sets AttributeSources field to given value.

### HasAttributeSources

`func (o *SpOAuthProcessorPolicyMapping) HasAttributeSources() bool`

HasAttributeSources returns a boolean if a field has been set.

### GetAttributeContractFulfillment

`func (o *SpOAuthProcessorPolicyMapping) GetAttributeContractFulfillment() map[string]AttributeFulfillmentValue`

GetAttributeContractFulfillment returns the AttributeContractFulfillment field if non-nil, zero value otherwise.

### GetAttributeContractFulfillmentOk

`func (o *SpOAuthProcessorPolicyMapping) GetAttributeContractFulfillmentOk() (*map[string]AttributeFulfillmentValue, bool)`

GetAttributeContractFulfillmentOk returns a tuple with the AttributeContractFulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContractFulfillment

`func (o *SpOAuthProcessorPolicyMapping) SetAttributeContractFulfillment(v map[string]AttributeFulfillmentValue)`

SetAttributeContractFulfillment sets AttributeContractFulfillment field to given value.


### GetIssuanceCriteria

`func (o *SpOAuthProcessorPolicyMapping) GetIssuanceCriteria() IssuanceCriteria`

GetIssuanceCriteria returns the IssuanceCriteria field if non-nil, zero value otherwise.

### GetIssuanceCriteriaOk

`func (o *SpOAuthProcessorPolicyMapping) GetIssuanceCriteriaOk() (*IssuanceCriteria, bool)`

GetIssuanceCriteriaOk returns a tuple with the IssuanceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceCriteria

`func (o *SpOAuthProcessorPolicyMapping) SetIssuanceCriteria(v IssuanceCriteria)`

SetIssuanceCriteria sets IssuanceCriteria field to given value.

### HasIssuanceCriteria

`func (o *SpOAuthProcessorPolicyMapping) HasIssuanceCriteria() bool`

HasIssuanceCriteria returns a boolean if a field has been set.

### GetProcessorPolicyRef

`func (o *SpOAuthProcessorPolicyMapping) GetProcessorPolicyRef() ResourceLink`

GetProcessorPolicyRef returns the ProcessorPolicyRef field if non-nil, zero value otherwise.

### GetProcessorPolicyRefOk

`func (o *SpOAuthProcessorPolicyMapping) GetProcessorPolicyRefOk() (*ResourceLink, bool)`

GetProcessorPolicyRefOk returns a tuple with the ProcessorPolicyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorPolicyRef

`func (o *SpOAuthProcessorPolicyMapping) SetProcessorPolicyRef(v ResourceLink)`

SetProcessorPolicyRef sets ProcessorPolicyRef field to given value.


### GetAttributeSource

`func (o *SpOAuthProcessorPolicyMapping) GetAttributeSource() []AttributeSourceAggregation`

GetAttributeSource returns the AttributeSource field if non-nil, zero value otherwise.

### GetAttributeSourceOk

`func (o *SpOAuthProcessorPolicyMapping) GetAttributeSourceOk() (*[]AttributeSourceAggregation, bool)`

GetAttributeSourceOk returns a tuple with the AttributeSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeSource

`func (o *SpOAuthProcessorPolicyMapping) SetAttributeSource(v []AttributeSourceAggregation)`

SetAttributeSource sets AttributeSource field to given value.

### HasAttributeSource

`func (o *SpOAuthProcessorPolicyMapping) HasAttributeSource() bool`

HasAttributeSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


