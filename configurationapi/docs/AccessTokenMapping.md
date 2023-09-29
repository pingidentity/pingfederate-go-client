# AccessTokenMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the Access Token Mapping. | 
**Context** | [**AccessTokenMappingContext**](AccessTokenMappingContext.md) |  | 
**AccessTokenManagerRef** | [**ResourceLink**](ResourceLink.md) |  | 
**AttributeSources** | Pointer to [**[]AttributeSource**](AttributeSource.md) | A list of configured data stores to look up attributes from. | [optional] 
**AttributeContractFulfillment** | [**map[string]AttributeFulfillmentValue**](AttributeFulfillmentValue.md) | A list of mappings from attribute names to their fulfillment values. | 
**IssuanceCriteria** | Pointer to [**IssuanceCriteria**](IssuanceCriteria.md) |  | [optional] 

## Methods

### NewAccessTokenMapping

`func NewAccessTokenMapping(id string, context AccessTokenMappingContext, accessTokenManagerRef ResourceLink, attributeContractFulfillment map[string]AttributeFulfillmentValue, ) *AccessTokenMapping`

NewAccessTokenMapping instantiates a new AccessTokenMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenMappingWithDefaults

`func NewAccessTokenMappingWithDefaults() *AccessTokenMapping`

NewAccessTokenMappingWithDefaults instantiates a new AccessTokenMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccessTokenMapping) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessTokenMapping) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessTokenMapping) SetId(v string)`

SetId sets Id field to given value.


### GetContext

`func (o *AccessTokenMapping) GetContext() AccessTokenMappingContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *AccessTokenMapping) GetContextOk() (*AccessTokenMappingContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *AccessTokenMapping) SetContext(v AccessTokenMappingContext)`

SetContext sets Context field to given value.


### GetAccessTokenManagerRef

`func (o *AccessTokenMapping) GetAccessTokenManagerRef() ResourceLink`

GetAccessTokenManagerRef returns the AccessTokenManagerRef field if non-nil, zero value otherwise.

### GetAccessTokenManagerRefOk

`func (o *AccessTokenMapping) GetAccessTokenManagerRefOk() (*ResourceLink, bool)`

GetAccessTokenManagerRefOk returns a tuple with the AccessTokenManagerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenManagerRef

`func (o *AccessTokenMapping) SetAccessTokenManagerRef(v ResourceLink)`

SetAccessTokenManagerRef sets AccessTokenManagerRef field to given value.


### GetAttributeSources

`func (o *AccessTokenMapping) GetAttributeSources() []AttributeSource`

GetAttributeSources returns the AttributeSources field if non-nil, zero value otherwise.

### GetAttributeSourcesOk

`func (o *AccessTokenMapping) GetAttributeSourcesOk() (*[]AttributeSource, bool)`

GetAttributeSourcesOk returns a tuple with the AttributeSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeSources

`func (o *AccessTokenMapping) SetAttributeSources(v []AttributeSource)`

SetAttributeSources sets AttributeSources field to given value.

### HasAttributeSources

`func (o *AccessTokenMapping) HasAttributeSources() bool`

HasAttributeSources returns a boolean if a field has been set.

### GetAttributeContractFulfillment

`func (o *AccessTokenMapping) GetAttributeContractFulfillment() map[string]AttributeFulfillmentValue`

GetAttributeContractFulfillment returns the AttributeContractFulfillment field if non-nil, zero value otherwise.

### GetAttributeContractFulfillmentOk

`func (o *AccessTokenMapping) GetAttributeContractFulfillmentOk() (*map[string]AttributeFulfillmentValue, bool)`

GetAttributeContractFulfillmentOk returns a tuple with the AttributeContractFulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContractFulfillment

`func (o *AccessTokenMapping) SetAttributeContractFulfillment(v map[string]AttributeFulfillmentValue)`

SetAttributeContractFulfillment sets AttributeContractFulfillment field to given value.


### GetIssuanceCriteria

`func (o *AccessTokenMapping) GetIssuanceCriteria() IssuanceCriteria`

GetIssuanceCriteria returns the IssuanceCriteria field if non-nil, zero value otherwise.

### GetIssuanceCriteriaOk

`func (o *AccessTokenMapping) GetIssuanceCriteriaOk() (*IssuanceCriteria, bool)`

GetIssuanceCriteriaOk returns a tuple with the IssuanceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceCriteria

`func (o *AccessTokenMapping) SetIssuanceCriteria(v IssuanceCriteria)`

SetIssuanceCriteria sets IssuanceCriteria field to given value.

### HasIssuanceCriteria

`func (o *AccessTokenMapping) HasIssuanceCriteria() bool`

HasIssuanceCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


