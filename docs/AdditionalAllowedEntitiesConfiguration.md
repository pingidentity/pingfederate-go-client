# AdditionalAllowedEntitiesConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowAdditionalEntities** | Pointer to **bool** | Set to true to configure additional entities or issuers to be accepted during entity or issuer validation. | [optional] 
**AllowAllEntities** | Pointer to **bool** | Set to true to accept any entity or issuer during entity or issuer validation. (Not Recommended) | [optional] 
**AdditionalAllowedEntities** | Pointer to [**[]Entity**](Entity.md) | An array of additional allowed entities or issuers to be accepted during entity or issuer validation. | [optional] 

## Methods

### NewAdditionalAllowedEntitiesConfiguration

`func NewAdditionalAllowedEntitiesConfiguration() *AdditionalAllowedEntitiesConfiguration`

NewAdditionalAllowedEntitiesConfiguration instantiates a new AdditionalAllowedEntitiesConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalAllowedEntitiesConfigurationWithDefaults

`func NewAdditionalAllowedEntitiesConfigurationWithDefaults() *AdditionalAllowedEntitiesConfiguration`

NewAdditionalAllowedEntitiesConfigurationWithDefaults instantiates a new AdditionalAllowedEntitiesConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowAdditionalEntities

`func (o *AdditionalAllowedEntitiesConfiguration) GetAllowAdditionalEntities() bool`

GetAllowAdditionalEntities returns the AllowAdditionalEntities field if non-nil, zero value otherwise.

### GetAllowAdditionalEntitiesOk

`func (o *AdditionalAllowedEntitiesConfiguration) GetAllowAdditionalEntitiesOk() (*bool, bool)`

GetAllowAdditionalEntitiesOk returns a tuple with the AllowAdditionalEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAdditionalEntities

`func (o *AdditionalAllowedEntitiesConfiguration) SetAllowAdditionalEntities(v bool)`

SetAllowAdditionalEntities sets AllowAdditionalEntities field to given value.

### HasAllowAdditionalEntities

`func (o *AdditionalAllowedEntitiesConfiguration) HasAllowAdditionalEntities() bool`

HasAllowAdditionalEntities returns a boolean if a field has been set.

### GetAllowAllEntities

`func (o *AdditionalAllowedEntitiesConfiguration) GetAllowAllEntities() bool`

GetAllowAllEntities returns the AllowAllEntities field if non-nil, zero value otherwise.

### GetAllowAllEntitiesOk

`func (o *AdditionalAllowedEntitiesConfiguration) GetAllowAllEntitiesOk() (*bool, bool)`

GetAllowAllEntitiesOk returns a tuple with the AllowAllEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAllEntities

`func (o *AdditionalAllowedEntitiesConfiguration) SetAllowAllEntities(v bool)`

SetAllowAllEntities sets AllowAllEntities field to given value.

### HasAllowAllEntities

`func (o *AdditionalAllowedEntitiesConfiguration) HasAllowAllEntities() bool`

HasAllowAllEntities returns a boolean if a field has been set.

### GetAdditionalAllowedEntities

`func (o *AdditionalAllowedEntitiesConfiguration) GetAdditionalAllowedEntities() []Entity`

GetAdditionalAllowedEntities returns the AdditionalAllowedEntities field if non-nil, zero value otherwise.

### GetAdditionalAllowedEntitiesOk

`func (o *AdditionalAllowedEntitiesConfiguration) GetAdditionalAllowedEntitiesOk() (*[]Entity, bool)`

GetAdditionalAllowedEntitiesOk returns a tuple with the AdditionalAllowedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAllowedEntities

`func (o *AdditionalAllowedEntitiesConfiguration) SetAdditionalAllowedEntities(v []Entity)`

SetAdditionalAllowedEntities sets AdditionalAllowedEntities field to given value.

### HasAdditionalAllowedEntities

`func (o *AdditionalAllowedEntitiesConfiguration) HasAdditionalAllowedEntities() bool`

HasAdditionalAllowedEntities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


