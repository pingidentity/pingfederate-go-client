# ProcessorPolicyToGeneratorMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeSources** | Pointer to [**[]AttributeSource**](AttributeSource.md) | A list of configured data stores to look up attributes from. | [optional] 
**AttributeContractFulfillment** | [**map[string]AttributeFulfillmentValue**](AttributeFulfillmentValue.md) | A list of mappings from attribute names to their fulfillment values. | 
**IssuanceCriteria** | Pointer to [**IssuanceCriteria**](IssuanceCriteria.md) |  | [optional] 
**Id** | Pointer to **string** | The id of the Token Exchange Processor policy to Token Generator mapping. This field is read-only and is ignored when passed in with the payload. | [optional] 
**SourceId** | **string** | The id of the Token Exchange Processor policy. | 
**TargetId** | **string** | The id of the Token Generator. | 
**LicenseConnectionGroupAssignment** | Pointer to **string** | The license connection group. | [optional] 

## Methods

### NewProcessorPolicyToGeneratorMapping

`func NewProcessorPolicyToGeneratorMapping(attributeContractFulfillment map[string]AttributeFulfillmentValue, sourceId string, targetId string, ) *ProcessorPolicyToGeneratorMapping`

NewProcessorPolicyToGeneratorMapping instantiates a new ProcessorPolicyToGeneratorMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorPolicyToGeneratorMappingWithDefaults

`func NewProcessorPolicyToGeneratorMappingWithDefaults() *ProcessorPolicyToGeneratorMapping`

NewProcessorPolicyToGeneratorMappingWithDefaults instantiates a new ProcessorPolicyToGeneratorMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeSources

`func (o *ProcessorPolicyToGeneratorMapping) GetAttributeSources() []AttributeSource`

GetAttributeSources returns the AttributeSources field if non-nil, zero value otherwise.

### GetAttributeSourcesOk

`func (o *ProcessorPolicyToGeneratorMapping) GetAttributeSourcesOk() (*[]AttributeSource, bool)`

GetAttributeSourcesOk returns a tuple with the AttributeSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeSources

`func (o *ProcessorPolicyToGeneratorMapping) SetAttributeSources(v []AttributeSource)`

SetAttributeSources sets AttributeSources field to given value.

### HasAttributeSources

`func (o *ProcessorPolicyToGeneratorMapping) HasAttributeSources() bool`

HasAttributeSources returns a boolean if a field has been set.

### GetAttributeContractFulfillment

`func (o *ProcessorPolicyToGeneratorMapping) GetAttributeContractFulfillment() map[string]AttributeFulfillmentValue`

GetAttributeContractFulfillment returns the AttributeContractFulfillment field if non-nil, zero value otherwise.

### GetAttributeContractFulfillmentOk

`func (o *ProcessorPolicyToGeneratorMapping) GetAttributeContractFulfillmentOk() (*map[string]AttributeFulfillmentValue, bool)`

GetAttributeContractFulfillmentOk returns a tuple with the AttributeContractFulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContractFulfillment

`func (o *ProcessorPolicyToGeneratorMapping) SetAttributeContractFulfillment(v map[string]AttributeFulfillmentValue)`

SetAttributeContractFulfillment sets AttributeContractFulfillment field to given value.


### GetIssuanceCriteria

`func (o *ProcessorPolicyToGeneratorMapping) GetIssuanceCriteria() IssuanceCriteria`

GetIssuanceCriteria returns the IssuanceCriteria field if non-nil, zero value otherwise.

### GetIssuanceCriteriaOk

`func (o *ProcessorPolicyToGeneratorMapping) GetIssuanceCriteriaOk() (*IssuanceCriteria, bool)`

GetIssuanceCriteriaOk returns a tuple with the IssuanceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceCriteria

`func (o *ProcessorPolicyToGeneratorMapping) SetIssuanceCriteria(v IssuanceCriteria)`

SetIssuanceCriteria sets IssuanceCriteria field to given value.

### HasIssuanceCriteria

`func (o *ProcessorPolicyToGeneratorMapping) HasIssuanceCriteria() bool`

HasIssuanceCriteria returns a boolean if a field has been set.

### GetId

`func (o *ProcessorPolicyToGeneratorMapping) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProcessorPolicyToGeneratorMapping) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProcessorPolicyToGeneratorMapping) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProcessorPolicyToGeneratorMapping) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSourceId

`func (o *ProcessorPolicyToGeneratorMapping) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ProcessorPolicyToGeneratorMapping) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ProcessorPolicyToGeneratorMapping) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetTargetId

`func (o *ProcessorPolicyToGeneratorMapping) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *ProcessorPolicyToGeneratorMapping) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *ProcessorPolicyToGeneratorMapping) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.


### GetLicenseConnectionGroupAssignment

`func (o *ProcessorPolicyToGeneratorMapping) GetLicenseConnectionGroupAssignment() string`

GetLicenseConnectionGroupAssignment returns the LicenseConnectionGroupAssignment field if non-nil, zero value otherwise.

### GetLicenseConnectionGroupAssignmentOk

`func (o *ProcessorPolicyToGeneratorMapping) GetLicenseConnectionGroupAssignmentOk() (*string, bool)`

GetLicenseConnectionGroupAssignmentOk returns a tuple with the LicenseConnectionGroupAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseConnectionGroupAssignment

`func (o *ProcessorPolicyToGeneratorMapping) SetLicenseConnectionGroupAssignment(v string)`

SetLicenseConnectionGroupAssignment sets LicenseConnectionGroupAssignment field to given value.

### HasLicenseConnectionGroupAssignment

`func (o *ProcessorPolicyToGeneratorMapping) HasLicenseConnectionGroupAssignment() bool`

HasLicenseConnectionGroupAssignment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


