# AccessTokenManagerMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessTokenManagerRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**AttributeSources** | Pointer to [**[]AttributeSourceAggregation**](AttributeSourceAggregation.md) | A list of configured data stores to look up attributes from. | [optional] 
**AttributeContractFulfillment** | [**map[string]AttributeFulfillmentValue**](AttributeFulfillmentValue.md) | A list of mappings from attribute names to their fulfillment values. | 
**IssuanceCriteria** | Pointer to [**IssuanceCriteria**](IssuanceCriteria.md) |  | [optional] 

## Methods

### NewAccessTokenManagerMapping

`func NewAccessTokenManagerMapping(attributeContractFulfillment map[string]AttributeFulfillmentValue, ) *AccessTokenManagerMapping`

NewAccessTokenManagerMapping instantiates a new AccessTokenManagerMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenManagerMappingWithDefaults

`func NewAccessTokenManagerMappingWithDefaults() *AccessTokenManagerMapping`

NewAccessTokenManagerMappingWithDefaults instantiates a new AccessTokenManagerMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessTokenManagerRef

`func (o *AccessTokenManagerMapping) GetAccessTokenManagerRef() ResourceLink`

GetAccessTokenManagerRef returns the AccessTokenManagerRef field if non-nil, zero value otherwise.

### GetAccessTokenManagerRefOk

`func (o *AccessTokenManagerMapping) GetAccessTokenManagerRefOk() (*ResourceLink, bool)`

GetAccessTokenManagerRefOk returns a tuple with the AccessTokenManagerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenManagerRef

`func (o *AccessTokenManagerMapping) SetAccessTokenManagerRef(v ResourceLink)`

SetAccessTokenManagerRef sets AccessTokenManagerRef field to given value.

### HasAccessTokenManagerRef

`func (o *AccessTokenManagerMapping) HasAccessTokenManagerRef() bool`

HasAccessTokenManagerRef returns a boolean if a field has been set.

### GetAttributeSources

`func (o *AccessTokenManagerMapping) GetAttributeSources() []AttributeSourceAggregation`

GetAttributeSources returns the AttributeSources field if non-nil, zero value otherwise.

### GetAttributeSourcesOk

`func (o *AccessTokenManagerMapping) GetAttributeSourcesOk() (*[]AttributeSourceAggregation, bool)`

GetAttributeSourcesOk returns a tuple with the AttributeSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeSources

`func (o *AccessTokenManagerMapping) SetAttributeSources(v []AttributeSourceAggregation)`

SetAttributeSources sets AttributeSources field to given value.

### HasAttributeSources

`func (o *AccessTokenManagerMapping) HasAttributeSources() bool`

HasAttributeSources returns a boolean if a field has been set.

### GetAttributeContractFulfillment

`func (o *AccessTokenManagerMapping) GetAttributeContractFulfillment() map[string]AttributeFulfillmentValue`

GetAttributeContractFulfillment returns the AttributeContractFulfillment field if non-nil, zero value otherwise.

### GetAttributeContractFulfillmentOk

`func (o *AccessTokenManagerMapping) GetAttributeContractFulfillmentOk() (*map[string]AttributeFulfillmentValue, bool)`

GetAttributeContractFulfillmentOk returns a tuple with the AttributeContractFulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContractFulfillment

`func (o *AccessTokenManagerMapping) SetAttributeContractFulfillment(v map[string]AttributeFulfillmentValue)`

SetAttributeContractFulfillment sets AttributeContractFulfillment field to given value.


### GetIssuanceCriteria

`func (o *AccessTokenManagerMapping) GetIssuanceCriteria() IssuanceCriteria`

GetIssuanceCriteria returns the IssuanceCriteria field if non-nil, zero value otherwise.

### GetIssuanceCriteriaOk

`func (o *AccessTokenManagerMapping) GetIssuanceCriteriaOk() (*IssuanceCriteria, bool)`

GetIssuanceCriteriaOk returns a tuple with the IssuanceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceCriteria

`func (o *AccessTokenManagerMapping) SetIssuanceCriteria(v IssuanceCriteria)`

SetIssuanceCriteria sets IssuanceCriteria field to given value.

### HasIssuanceCriteria

`func (o *AccessTokenManagerMapping) HasIssuanceCriteria() bool`

HasIssuanceCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


