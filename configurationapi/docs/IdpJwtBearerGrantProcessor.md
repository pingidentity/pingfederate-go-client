# IdpJwtBearerGrantProcessor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JwtBearerGrantProcessorId** | **string** | The identifier of the JWT Bearer Grant Processor to use for this connection. | 
**AttributeContract** | Pointer to [**IdpOAuthAttributeContract**](IdpOAuthAttributeContract.md) |  | [optional] 
**AccessTokenManagerMappings** | Pointer to [**[]AccessTokenManagerMapping**](AccessTokenManagerMapping.md) | A list of Access Token Manager mappings for the JWT Bearer Grant Processor. | [optional] 

## Methods

### NewIdpJwtBearerGrantProcessor

`func NewIdpJwtBearerGrantProcessor(jwtBearerGrantProcessorId string, ) *IdpJwtBearerGrantProcessor`

NewIdpJwtBearerGrantProcessor instantiates a new IdpJwtBearerGrantProcessor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpJwtBearerGrantProcessorWithDefaults

`func NewIdpJwtBearerGrantProcessorWithDefaults() *IdpJwtBearerGrantProcessor`

NewIdpJwtBearerGrantProcessorWithDefaults instantiates a new IdpJwtBearerGrantProcessor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJwtBearerGrantProcessorId

`func (o *IdpJwtBearerGrantProcessor) GetJwtBearerGrantProcessorId() string`

GetJwtBearerGrantProcessorId returns the JwtBearerGrantProcessorId field if non-nil, zero value otherwise.

### GetJwtBearerGrantProcessorIdOk

`func (o *IdpJwtBearerGrantProcessor) GetJwtBearerGrantProcessorIdOk() (*string, bool)`

GetJwtBearerGrantProcessorIdOk returns a tuple with the JwtBearerGrantProcessorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtBearerGrantProcessorId

`func (o *IdpJwtBearerGrantProcessor) SetJwtBearerGrantProcessorId(v string)`

SetJwtBearerGrantProcessorId sets JwtBearerGrantProcessorId field to given value.


### GetAttributeContract

`func (o *IdpJwtBearerGrantProcessor) GetAttributeContract() IdpOAuthAttributeContract`

GetAttributeContract returns the AttributeContract field if non-nil, zero value otherwise.

### GetAttributeContractOk

`func (o *IdpJwtBearerGrantProcessor) GetAttributeContractOk() (*IdpOAuthAttributeContract, bool)`

GetAttributeContractOk returns a tuple with the AttributeContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContract

`func (o *IdpJwtBearerGrantProcessor) SetAttributeContract(v IdpOAuthAttributeContract)`

SetAttributeContract sets AttributeContract field to given value.

### HasAttributeContract

`func (o *IdpJwtBearerGrantProcessor) HasAttributeContract() bool`

HasAttributeContract returns a boolean if a field has been set.

### GetAccessTokenManagerMappings

`func (o *IdpJwtBearerGrantProcessor) GetAccessTokenManagerMappings() []AccessTokenManagerMapping`

GetAccessTokenManagerMappings returns the AccessTokenManagerMappings field if non-nil, zero value otherwise.

### GetAccessTokenManagerMappingsOk

`func (o *IdpJwtBearerGrantProcessor) GetAccessTokenManagerMappingsOk() (*[]AccessTokenManagerMapping, bool)`

GetAccessTokenManagerMappingsOk returns a tuple with the AccessTokenManagerMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenManagerMappings

`func (o *IdpJwtBearerGrantProcessor) SetAccessTokenManagerMappings(v []AccessTokenManagerMapping)`

SetAccessTokenManagerMappings sets AccessTokenManagerMappings field to given value.

### HasAccessTokenManagerMappings

`func (o *IdpJwtBearerGrantProcessor) HasAccessTokenManagerMappings() bool`

HasAccessTokenManagerMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


