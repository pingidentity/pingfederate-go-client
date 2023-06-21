# SpAttributeQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | **[]string** | The list of attributes that may be returned to the SP in the response to an attribute request. | 
**AttributeContractFulfillment** | [**map[string]AttributeFulfillmentValue**](AttributeFulfillmentValue.md) | A list of mappings from attribute names to their fulfillment values. | 
**IssuanceCriteria** | Pointer to [**IssuanceCriteria**](IssuanceCriteria.md) |  | [optional] 
**Policy** | Pointer to [**SpAttributeQueryPolicy**](SpAttributeQueryPolicy.md) |  | [optional] 
**AttributeSources** | [**[]AttributeSource**](AttributeSource.md) | A list of configured data stores to look up attributes from. | 

## Methods

### NewSpAttributeQuery

`func NewSpAttributeQuery(attributes []string, attributeContractFulfillment map[string]AttributeFulfillmentValue, attributeSources []AttributeSource, ) *SpAttributeQuery`

NewSpAttributeQuery instantiates a new SpAttributeQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpAttributeQueryWithDefaults

`func NewSpAttributeQueryWithDefaults() *SpAttributeQuery`

NewSpAttributeQueryWithDefaults instantiates a new SpAttributeQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *SpAttributeQuery) GetAttributes() []string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SpAttributeQuery) GetAttributesOk() (*[]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SpAttributeQuery) SetAttributes(v []string)`

SetAttributes sets Attributes field to given value.


### GetAttributeContractFulfillment

`func (o *SpAttributeQuery) GetAttributeContractFulfillment() map[string]AttributeFulfillmentValue`

GetAttributeContractFulfillment returns the AttributeContractFulfillment field if non-nil, zero value otherwise.

### GetAttributeContractFulfillmentOk

`func (o *SpAttributeQuery) GetAttributeContractFulfillmentOk() (*map[string]AttributeFulfillmentValue, bool)`

GetAttributeContractFulfillmentOk returns a tuple with the AttributeContractFulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContractFulfillment

`func (o *SpAttributeQuery) SetAttributeContractFulfillment(v map[string]AttributeFulfillmentValue)`

SetAttributeContractFulfillment sets AttributeContractFulfillment field to given value.


### GetIssuanceCriteria

`func (o *SpAttributeQuery) GetIssuanceCriteria() IssuanceCriteria`

GetIssuanceCriteria returns the IssuanceCriteria field if non-nil, zero value otherwise.

### GetIssuanceCriteriaOk

`func (o *SpAttributeQuery) GetIssuanceCriteriaOk() (*IssuanceCriteria, bool)`

GetIssuanceCriteriaOk returns a tuple with the IssuanceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceCriteria

`func (o *SpAttributeQuery) SetIssuanceCriteria(v IssuanceCriteria)`

SetIssuanceCriteria sets IssuanceCriteria field to given value.

### HasIssuanceCriteria

`func (o *SpAttributeQuery) HasIssuanceCriteria() bool`

HasIssuanceCriteria returns a boolean if a field has been set.

### GetPolicy

`func (o *SpAttributeQuery) GetPolicy() SpAttributeQueryPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *SpAttributeQuery) GetPolicyOk() (*SpAttributeQueryPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *SpAttributeQuery) SetPolicy(v SpAttributeQueryPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *SpAttributeQuery) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetAttributeSources

`func (o *SpAttributeQuery) GetAttributeSources() []AttributeSource`

GetAttributeSources returns the AttributeSources field if non-nil, zero value otherwise.

### GetAttributeSourcesOk

`func (o *SpAttributeQuery) GetAttributeSourcesOk() (*[]AttributeSource, bool)`

GetAttributeSourcesOk returns a tuple with the AttributeSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeSources

`func (o *SpAttributeQuery) SetAttributeSources(v []AttributeSource)`

SetAttributeSources sets AttributeSources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


