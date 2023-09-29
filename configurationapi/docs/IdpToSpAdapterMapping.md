# IdpToSpAdapterMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeSources** | Pointer to [**[]AttributeSource**](AttributeSource.md) | A list of configured data stores to look up attributes from. | [optional] 
**AttributeContractFulfillment** | [**map[string]AttributeFulfillmentValue**](AttributeFulfillmentValue.md) | A list of mappings from attribute names to their fulfillment values. | 
**IssuanceCriteria** | Pointer to [**IssuanceCriteria**](IssuanceCriteria.md) |  | [optional] 
**SourceId** | **string** | The id of the IdP Adapter. | 
**TargetId** | **string** | The id of the SP Adapter. | 
**Id** | Pointer to **string** | The id of the IdP-to-SP Adapter mapping. This field is read-only and is ignored when passed in with the payload. | [optional] 
**DefaultTargetResource** | Pointer to **string** | Default target URL for this adapter-to-adapter mapping configuration. | [optional] 
**LicenseConnectionGroupAssignment** | Pointer to **string** | The license connection group. | [optional] 
**ApplicationName** | Pointer to **string** | The application name. | [optional] 
**ApplicationIconUrl** | Pointer to **string** | The application icon URL. | [optional] 

## Methods

### NewIdpToSpAdapterMapping

`func NewIdpToSpAdapterMapping(attributeContractFulfillment map[string]AttributeFulfillmentValue, sourceId string, targetId string, ) *IdpToSpAdapterMapping`

NewIdpToSpAdapterMapping instantiates a new IdpToSpAdapterMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpToSpAdapterMappingWithDefaults

`func NewIdpToSpAdapterMappingWithDefaults() *IdpToSpAdapterMapping`

NewIdpToSpAdapterMappingWithDefaults instantiates a new IdpToSpAdapterMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeSources

`func (o *IdpToSpAdapterMapping) GetAttributeSources() []AttributeSource`

GetAttributeSources returns the AttributeSources field if non-nil, zero value otherwise.

### GetAttributeSourcesOk

`func (o *IdpToSpAdapterMapping) GetAttributeSourcesOk() (*[]AttributeSource, bool)`

GetAttributeSourcesOk returns a tuple with the AttributeSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeSources

`func (o *IdpToSpAdapterMapping) SetAttributeSources(v []AttributeSource)`

SetAttributeSources sets AttributeSources field to given value.

### HasAttributeSources

`func (o *IdpToSpAdapterMapping) HasAttributeSources() bool`

HasAttributeSources returns a boolean if a field has been set.

### GetAttributeContractFulfillment

`func (o *IdpToSpAdapterMapping) GetAttributeContractFulfillment() map[string]AttributeFulfillmentValue`

GetAttributeContractFulfillment returns the AttributeContractFulfillment field if non-nil, zero value otherwise.

### GetAttributeContractFulfillmentOk

`func (o *IdpToSpAdapterMapping) GetAttributeContractFulfillmentOk() (*map[string]AttributeFulfillmentValue, bool)`

GetAttributeContractFulfillmentOk returns a tuple with the AttributeContractFulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContractFulfillment

`func (o *IdpToSpAdapterMapping) SetAttributeContractFulfillment(v map[string]AttributeFulfillmentValue)`

SetAttributeContractFulfillment sets AttributeContractFulfillment field to given value.


### GetIssuanceCriteria

`func (o *IdpToSpAdapterMapping) GetIssuanceCriteria() IssuanceCriteria`

GetIssuanceCriteria returns the IssuanceCriteria field if non-nil, zero value otherwise.

### GetIssuanceCriteriaOk

`func (o *IdpToSpAdapterMapping) GetIssuanceCriteriaOk() (*IssuanceCriteria, bool)`

GetIssuanceCriteriaOk returns a tuple with the IssuanceCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceCriteria

`func (o *IdpToSpAdapterMapping) SetIssuanceCriteria(v IssuanceCriteria)`

SetIssuanceCriteria sets IssuanceCriteria field to given value.

### HasIssuanceCriteria

`func (o *IdpToSpAdapterMapping) HasIssuanceCriteria() bool`

HasIssuanceCriteria returns a boolean if a field has been set.

### GetSourceId

`func (o *IdpToSpAdapterMapping) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *IdpToSpAdapterMapping) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *IdpToSpAdapterMapping) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetTargetId

`func (o *IdpToSpAdapterMapping) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *IdpToSpAdapterMapping) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *IdpToSpAdapterMapping) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.


### GetId

`func (o *IdpToSpAdapterMapping) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdpToSpAdapterMapping) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdpToSpAdapterMapping) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdpToSpAdapterMapping) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDefaultTargetResource

`func (o *IdpToSpAdapterMapping) GetDefaultTargetResource() string`

GetDefaultTargetResource returns the DefaultTargetResource field if non-nil, zero value otherwise.

### GetDefaultTargetResourceOk

`func (o *IdpToSpAdapterMapping) GetDefaultTargetResourceOk() (*string, bool)`

GetDefaultTargetResourceOk returns a tuple with the DefaultTargetResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTargetResource

`func (o *IdpToSpAdapterMapping) SetDefaultTargetResource(v string)`

SetDefaultTargetResource sets DefaultTargetResource field to given value.

### HasDefaultTargetResource

`func (o *IdpToSpAdapterMapping) HasDefaultTargetResource() bool`

HasDefaultTargetResource returns a boolean if a field has been set.

### GetLicenseConnectionGroupAssignment

`func (o *IdpToSpAdapterMapping) GetLicenseConnectionGroupAssignment() string`

GetLicenseConnectionGroupAssignment returns the LicenseConnectionGroupAssignment field if non-nil, zero value otherwise.

### GetLicenseConnectionGroupAssignmentOk

`func (o *IdpToSpAdapterMapping) GetLicenseConnectionGroupAssignmentOk() (*string, bool)`

GetLicenseConnectionGroupAssignmentOk returns a tuple with the LicenseConnectionGroupAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseConnectionGroupAssignment

`func (o *IdpToSpAdapterMapping) SetLicenseConnectionGroupAssignment(v string)`

SetLicenseConnectionGroupAssignment sets LicenseConnectionGroupAssignment field to given value.

### HasLicenseConnectionGroupAssignment

`func (o *IdpToSpAdapterMapping) HasLicenseConnectionGroupAssignment() bool`

HasLicenseConnectionGroupAssignment returns a boolean if a field has been set.

### GetApplicationName

`func (o *IdpToSpAdapterMapping) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *IdpToSpAdapterMapping) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *IdpToSpAdapterMapping) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *IdpToSpAdapterMapping) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetApplicationIconUrl

`func (o *IdpToSpAdapterMapping) GetApplicationIconUrl() string`

GetApplicationIconUrl returns the ApplicationIconUrl field if non-nil, zero value otherwise.

### GetApplicationIconUrlOk

`func (o *IdpToSpAdapterMapping) GetApplicationIconUrlOk() (*string, bool)`

GetApplicationIconUrlOk returns a tuple with the ApplicationIconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIconUrl

`func (o *IdpToSpAdapterMapping) SetApplicationIconUrl(v string)`

SetApplicationIconUrl sets ApplicationIconUrl field to given value.

### HasApplicationIconUrl

`func (o *IdpToSpAdapterMapping) HasApplicationIconUrl() bool`

HasApplicationIconUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


