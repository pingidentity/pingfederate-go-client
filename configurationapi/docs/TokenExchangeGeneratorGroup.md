# TokenExchangeGeneratorGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Token Exchange Generator group ID. ID is unique. | 
**Name** | **string** | The Token Exchange Generator group name. Name is unique. | 
**ResourceUris** | Pointer to **[]string** | The list of  resource URI&#39;s which map to this Token Exchange Generator group. | [optional] 
**GeneratorMappings** | [**[]TokenExchangeGeneratorMapping**](TokenExchangeGeneratorMapping.md) | A list of Token Generator mapping into an OAuth 2.0 Token Exchange requested token type. | 

## Methods

### NewTokenExchangeGeneratorGroup

`func NewTokenExchangeGeneratorGroup(id string, name string, generatorMappings []TokenExchangeGeneratorMapping, ) *TokenExchangeGeneratorGroup`

NewTokenExchangeGeneratorGroup instantiates a new TokenExchangeGeneratorGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenExchangeGeneratorGroupWithDefaults

`func NewTokenExchangeGeneratorGroupWithDefaults() *TokenExchangeGeneratorGroup`

NewTokenExchangeGeneratorGroupWithDefaults instantiates a new TokenExchangeGeneratorGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TokenExchangeGeneratorGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenExchangeGeneratorGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenExchangeGeneratorGroup) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TokenExchangeGeneratorGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenExchangeGeneratorGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenExchangeGeneratorGroup) SetName(v string)`

SetName sets Name field to given value.


### GetResourceUris

`func (o *TokenExchangeGeneratorGroup) GetResourceUris() []string`

GetResourceUris returns the ResourceUris field if non-nil, zero value otherwise.

### GetResourceUrisOk

`func (o *TokenExchangeGeneratorGroup) GetResourceUrisOk() (*[]string, bool)`

GetResourceUrisOk returns a tuple with the ResourceUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUris

`func (o *TokenExchangeGeneratorGroup) SetResourceUris(v []string)`

SetResourceUris sets ResourceUris field to given value.

### HasResourceUris

`func (o *TokenExchangeGeneratorGroup) HasResourceUris() bool`

HasResourceUris returns a boolean if a field has been set.

### GetGeneratorMappings

`func (o *TokenExchangeGeneratorGroup) GetGeneratorMappings() []TokenExchangeGeneratorMapping`

GetGeneratorMappings returns the GeneratorMappings field if non-nil, zero value otherwise.

### GetGeneratorMappingsOk

`func (o *TokenExchangeGeneratorGroup) GetGeneratorMappingsOk() (*[]TokenExchangeGeneratorMapping, bool)`

GetGeneratorMappingsOk returns a tuple with the GeneratorMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratorMappings

`func (o *TokenExchangeGeneratorGroup) SetGeneratorMappings(v []TokenExchangeGeneratorMapping)`

SetGeneratorMappings sets GeneratorMappings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


