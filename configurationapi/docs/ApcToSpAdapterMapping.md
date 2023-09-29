# ApcToSpAdapterMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeSources** | Pointer to [**[]AttributeSource**](AttributeSource.md) | A list of configured data stores to look up attributes from. | [optional] 
**AttributeContractFulfillment** | [**map[string]AttributeFulfillmentValue**](AttributeFulfillmentValue.md) | A list of mappings from attribute names to their fulfillment values. | 
**IssuanceCriteria** | Pointer to [**IssuanceCriteria**](IssuanceCriteria.md) |  | [optional] 
**SourceId** | **string** | The id of the Authentication Policy Contract. | 
**TargetId** | **string** | The id of the SP Adapter. | 
**Id** | Pointer to **string** | The id of the APC-to-SP Adapter mapping. This field is read-only and is ignored when passed in with the payload. | [optional] 
**DefaultTargetResource** | Pointer to **string** | Default target URL for this APC-to-adapter mapping configuration. | [optional] 
**LicenseConnectionGroupAssignment** | Pointer to **string** | The license connection group. | [optional] 

## Methods

### NewApcToSpAdapterMapping

`func NewApcToSpAdapterMapping(attributeContractFulfillment map[string]AttributeFulfillmentValue, sourceId string, targetId string, ) *ApcToSpAdapterMapping`

NewApcToSpAdapterMapping instantiates a new ApcToSpAdapterMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApcToSpAdapterMappingWithDefaults

`func NewApcToSpAdapterMappingWithDefaults() *ApcToSpAdapterMapping`

NewApcToSpAdapterMappingWithDefaults instantiates a new ApcToSpAdapterMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeSources

`func (o *ApcToSpAdapterMapping) GetAttributeSources() []AttributeSource`

GetAttributeSources returns the AttributeSources field if non-nil, zero value otherwise.

### GetAttributeSourcesOk

`func (o *ApcToSpAdapterMapping) GetAttributeSourcesOk() (*[]AttributeSource, bool)`

GetAttributeSourcesOk returns a tuple with the AttributeSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeSources

`func (o *ApcToSpAdapterMapping) SetAttributeSources(v []AttributeSource)`

SetAttributeSources sets AttributeSources field to given value.

### HasAttributeSources

`func (o *ApcToSpAdapterMapping) HasAttributeSources() bool`

HasAttributeSources returns a boolean if a field has been set.

### GetAttributeContractFulfillment

`func (o *ApcToSpAdapterMapping) GetAttributeContractFulfillment() map[string]AttributeFulfillmentValue`

GetAttributeContractFulfillment returns the AttributeContractFulfillment field if non-nil, zero value otherwise.

### GetAttributeContractFulfillmentOk

`func (o *ApcToSpAdapterMapping) GetAttributeContractFulfillmentOk() (*map[string]AttributeFulfillmentValue, bool)`

GetAttributeContractFulfillmentOk returns a tuple with the AttributeContractFulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContractFulfillment

`func (o *ApcToSpAdapterMapping) SetAttributeContractFulfillment(v map[string]AttributeFulfillmentValue)`

SetAttributeContractFulfillment sets AttributeContractFulfillment field to given value.


### GetIssuanceCriteria

`func (o *ApcToSpAdapterMapping) GetIssuanceCriteria() IssuanceCriteria`

GetIssuanceCriteria returns the IssuanceCriteria field if non-nil, zero value otherwise.

### GetIssuanceCriteriaOk

`func (o *ApcToSpAdapterMapping) GetIssuanceCriteriaOk() (*IssuanceCriteria, bool)`

GetIssuanceCriteriaOk returns a tuple with the IssuanceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceCriteria

`func (o *ApcToSpAdapterMapping) SetIssuanceCriteria(v IssuanceCriteria)`

SetIssuanceCriteria sets IssuanceCriteria field to given value.

### HasIssuanceCriteria

`func (o *ApcToSpAdapterMapping) HasIssuanceCriteria() bool`

HasIssuanceCriteria returns a boolean if a field has been set.

### GetSourceId

`func (o *ApcToSpAdapterMapping) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ApcToSpAdapterMapping) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ApcToSpAdapterMapping) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetTargetId

`func (o *ApcToSpAdapterMapping) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *ApcToSpAdapterMapping) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *ApcToSpAdapterMapping) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.


### GetId

`func (o *ApcToSpAdapterMapping) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApcToSpAdapterMapping) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApcToSpAdapterMapping) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApcToSpAdapterMapping) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDefaultTargetResource

`func (o *ApcToSpAdapterMapping) GetDefaultTargetResource() string`

GetDefaultTargetResource returns the DefaultTargetResource field if non-nil, zero value otherwise.

### GetDefaultTargetResourceOk

`func (o *ApcToSpAdapterMapping) GetDefaultTargetResourceOk() (*string, bool)`

GetDefaultTargetResourceOk returns a tuple with the DefaultTargetResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTargetResource

`func (o *ApcToSpAdapterMapping) SetDefaultTargetResource(v string)`

SetDefaultTargetResource sets DefaultTargetResource field to given value.

### HasDefaultTargetResource

`func (o *ApcToSpAdapterMapping) HasDefaultTargetResource() bool`

HasDefaultTargetResource returns a boolean if a field has been set.

### GetLicenseConnectionGroupAssignment

`func (o *ApcToSpAdapterMapping) GetLicenseConnectionGroupAssignment() string`

GetLicenseConnectionGroupAssignment returns the LicenseConnectionGroupAssignment field if non-nil, zero value otherwise.

### GetLicenseConnectionGroupAssignmentOk

`func (o *ApcToSpAdapterMapping) GetLicenseConnectionGroupAssignmentOk() (*string, bool)`

GetLicenseConnectionGroupAssignmentOk returns a tuple with the LicenseConnectionGroupAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseConnectionGroupAssignment

`func (o *ApcToSpAdapterMapping) SetLicenseConnectionGroupAssignment(v string)`

SetLicenseConnectionGroupAssignment sets LicenseConnectionGroupAssignment field to given value.

### HasLicenseConnectionGroupAssignment

`func (o *ApcToSpAdapterMapping) HasLicenseConnectionGroupAssignment() bool`

HasLicenseConnectionGroupAssignment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


