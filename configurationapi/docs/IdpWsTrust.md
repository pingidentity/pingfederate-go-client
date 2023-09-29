# IdpWsTrust

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeContract** | [**IdpWsTrustAttributeContract**](IdpWsTrustAttributeContract.md) |  | 
**GenerateLocalToken** | **bool** | Indicates whether a local token needs to be generated. The default value is false. | 
**TokenGeneratorMappings** | Pointer to [**[]SpTokenGeneratorMapping**](SpTokenGeneratorMapping.md) | A list of token generators to generate local tokens. Required if a local token needs to be generated. | [optional] 

## Methods

### NewIdpWsTrust

`func NewIdpWsTrust(attributeContract IdpWsTrustAttributeContract, generateLocalToken bool, ) *IdpWsTrust`

NewIdpWsTrust instantiates a new IdpWsTrust object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpWsTrustWithDefaults

`func NewIdpWsTrustWithDefaults() *IdpWsTrust`

NewIdpWsTrustWithDefaults instantiates a new IdpWsTrust object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeContract

`func (o *IdpWsTrust) GetAttributeContract() IdpWsTrustAttributeContract`

GetAttributeContract returns the AttributeContract field if non-nil, zero value otherwise.

### GetAttributeContractOk

`func (o *IdpWsTrust) GetAttributeContractOk() (*IdpWsTrustAttributeContract, bool)`

GetAttributeContractOk returns a tuple with the AttributeContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContract

`func (o *IdpWsTrust) SetAttributeContract(v IdpWsTrustAttributeContract)`

SetAttributeContract sets AttributeContract field to given value.


### GetGenerateLocalToken

`func (o *IdpWsTrust) GetGenerateLocalToken() bool`

GetGenerateLocalToken returns the GenerateLocalToken field if non-nil, zero value otherwise.

### GetGenerateLocalTokenOk

`func (o *IdpWsTrust) GetGenerateLocalTokenOk() (*bool, bool)`

GetGenerateLocalTokenOk returns a tuple with the GenerateLocalToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateLocalToken

`func (o *IdpWsTrust) SetGenerateLocalToken(v bool)`

SetGenerateLocalToken sets GenerateLocalToken field to given value.


### GetTokenGeneratorMappings

`func (o *IdpWsTrust) GetTokenGeneratorMappings() []SpTokenGeneratorMapping`

GetTokenGeneratorMappings returns the TokenGeneratorMappings field if non-nil, zero value otherwise.

### GetTokenGeneratorMappingsOk

`func (o *IdpWsTrust) GetTokenGeneratorMappingsOk() (*[]SpTokenGeneratorMapping, bool)`

GetTokenGeneratorMappingsOk returns a tuple with the TokenGeneratorMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenGeneratorMappings

`func (o *IdpWsTrust) SetTokenGeneratorMappings(v []SpTokenGeneratorMapping)`

SetTokenGeneratorMappings sets TokenGeneratorMappings field to given value.

### HasTokenGeneratorMappings

`func (o *IdpWsTrust) HasTokenGeneratorMappings() bool`

HasTokenGeneratorMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


