# TokenToTokenMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeSources** | Pointer to [**[]AttributeSourceAggregation**](AttributeSourceAggregation.md) | A list of configured data stores to look up attributes from. | [optional] 
**AttributeContractFulfillment** | [**map[string]AttributeFulfillmentValue**](AttributeFulfillmentValue.md) | A list of mappings from attribute names to their fulfillment values. | 
**IssuanceCriteria** | Pointer to [**IssuanceCriteria**](IssuanceCriteria.md) |  | [optional] 
**SourceId** | **string** | The id of the Token Processor. | 
**TargetId** | **string** | The id of the Token Generator. | 
**Id** | Pointer to **string** | The id of the Token Processor to Token Generator mapping. This field is read-only and is ignored when passed in with the payload. | [optional] 
**DefaultTargetResource** | Pointer to **string** | Default target URL for this Token Processor to Token Generator mapping configuration. | [optional] 
**LicenseConnectionGroupAssignment** | Pointer to **string** | The license connection group. | [optional] 

## Methods

### NewTokenToTokenMapping

`func NewTokenToTokenMapping(attributeContractFulfillment map[string]AttributeFulfillmentValue, sourceId string, targetId string, ) *TokenToTokenMapping`

NewTokenToTokenMapping instantiates a new TokenToTokenMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenToTokenMappingWithDefaults

`func NewTokenToTokenMappingWithDefaults() *TokenToTokenMapping`

NewTokenToTokenMappingWithDefaults instantiates a new TokenToTokenMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeSources

`func (o *TokenToTokenMapping) GetAttributeSources() []AttributeSourceAggregation`

GetAttributeSources returns the AttributeSources field if non-nil, zero value otherwise.

### GetAttributeSourcesOk

`func (o *TokenToTokenMapping) GetAttributeSourcesOk() (*[]AttributeSourceAggregation, bool)`

GetAttributeSourcesOk returns a tuple with the AttributeSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeSources

`func (o *TokenToTokenMapping) SetAttributeSources(v []AttributeSourceAggregation)`

SetAttributeSources sets AttributeSources field to given value.

### HasAttributeSources

`func (o *TokenToTokenMapping) HasAttributeSources() bool`

HasAttributeSources returns a boolean if a field has been set.

### GetAttributeContractFulfillment

`func (o *TokenToTokenMapping) GetAttributeContractFulfillment() map[string]AttributeFulfillmentValue`

GetAttributeContractFulfillment returns the AttributeContractFulfillment field if non-nil, zero value otherwise.

### GetAttributeContractFulfillmentOk

`func (o *TokenToTokenMapping) GetAttributeContractFulfillmentOk() (*map[string]AttributeFulfillmentValue, bool)`

GetAttributeContractFulfillmentOk returns a tuple with the AttributeContractFulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContractFulfillment

`func (o *TokenToTokenMapping) SetAttributeContractFulfillment(v map[string]AttributeFulfillmentValue)`

SetAttributeContractFulfillment sets AttributeContractFulfillment field to given value.


### GetIssuanceCriteria

`func (o *TokenToTokenMapping) GetIssuanceCriteria() IssuanceCriteria`

GetIssuanceCriteria returns the IssuanceCriteria field if non-nil, zero value otherwise.

### GetIssuanceCriteriaOk

`func (o *TokenToTokenMapping) GetIssuanceCriteriaOk() (*IssuanceCriteria, bool)`

GetIssuanceCriteriaOk returns a tuple with the IssuanceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceCriteria

`func (o *TokenToTokenMapping) SetIssuanceCriteria(v IssuanceCriteria)`

SetIssuanceCriteria sets IssuanceCriteria field to given value.

### HasIssuanceCriteria

`func (o *TokenToTokenMapping) HasIssuanceCriteria() bool`

HasIssuanceCriteria returns a boolean if a field has been set.

### GetSourceId

`func (o *TokenToTokenMapping) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *TokenToTokenMapping) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *TokenToTokenMapping) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetTargetId

`func (o *TokenToTokenMapping) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *TokenToTokenMapping) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *TokenToTokenMapping) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.


### GetId

`func (o *TokenToTokenMapping) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenToTokenMapping) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenToTokenMapping) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TokenToTokenMapping) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDefaultTargetResource

`func (o *TokenToTokenMapping) GetDefaultTargetResource() string`

GetDefaultTargetResource returns the DefaultTargetResource field if non-nil, zero value otherwise.

### GetDefaultTargetResourceOk

`func (o *TokenToTokenMapping) GetDefaultTargetResourceOk() (*string, bool)`

GetDefaultTargetResourceOk returns a tuple with the DefaultTargetResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTargetResource

`func (o *TokenToTokenMapping) SetDefaultTargetResource(v string)`

SetDefaultTargetResource sets DefaultTargetResource field to given value.

### HasDefaultTargetResource

`func (o *TokenToTokenMapping) HasDefaultTargetResource() bool`

HasDefaultTargetResource returns a boolean if a field has been set.

### GetLicenseConnectionGroupAssignment

`func (o *TokenToTokenMapping) GetLicenseConnectionGroupAssignment() string`

GetLicenseConnectionGroupAssignment returns the LicenseConnectionGroupAssignment field if non-nil, zero value otherwise.

### GetLicenseConnectionGroupAssignmentOk

`func (o *TokenToTokenMapping) GetLicenseConnectionGroupAssignmentOk() (*string, bool)`

GetLicenseConnectionGroupAssignmentOk returns a tuple with the LicenseConnectionGroupAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseConnectionGroupAssignment

`func (o *TokenToTokenMapping) SetLicenseConnectionGroupAssignment(v string)`

SetLicenseConnectionGroupAssignment sets LicenseConnectionGroupAssignment field to given value.

### HasLicenseConnectionGroupAssignment

`func (o *TokenToTokenMapping) HasLicenseConnectionGroupAssignment() bool`

HasLicenseConnectionGroupAssignment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


