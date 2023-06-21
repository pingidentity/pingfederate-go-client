# IdpAdapterAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthnCtxClassRef** | Pointer to **string** | The fixed value that indicates how the user was authenticated. | [optional] 
**AttributeMapping** | Pointer to [**IdpAdapterContractMapping**](IdpAdapterContractMapping.md) |  | [optional] 
**AttributeContract** | Pointer to [**IdpAdapterAttributeContract**](IdpAdapterAttributeContract.md) |  | [optional] 

## Methods

### NewIdpAdapterAllOf

`func NewIdpAdapterAllOf() *IdpAdapterAllOf`

NewIdpAdapterAllOf instantiates a new IdpAdapterAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpAdapterAllOfWithDefaults

`func NewIdpAdapterAllOfWithDefaults() *IdpAdapterAllOf`

NewIdpAdapterAllOfWithDefaults instantiates a new IdpAdapterAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthnCtxClassRef

`func (o *IdpAdapterAllOf) GetAuthnCtxClassRef() string`

GetAuthnCtxClassRef returns the AuthnCtxClassRef field if non-nil, zero value otherwise.

### GetAuthnCtxClassRefOk

`func (o *IdpAdapterAllOf) GetAuthnCtxClassRefOk() (*string, bool)`

GetAuthnCtxClassRefOk returns a tuple with the AuthnCtxClassRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnCtxClassRef

`func (o *IdpAdapterAllOf) SetAuthnCtxClassRef(v string)`

SetAuthnCtxClassRef sets AuthnCtxClassRef field to given value.

### HasAuthnCtxClassRef

`func (o *IdpAdapterAllOf) HasAuthnCtxClassRef() bool`

HasAuthnCtxClassRef returns a boolean if a field has been set.

### GetAttributeMapping

`func (o *IdpAdapterAllOf) GetAttributeMapping() IdpAdapterContractMapping`

GetAttributeMapping returns the AttributeMapping field if non-nil, zero value otherwise.

### GetAttributeMappingOk

`func (o *IdpAdapterAllOf) GetAttributeMappingOk() (*IdpAdapterContractMapping, bool)`

GetAttributeMappingOk returns a tuple with the AttributeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeMapping

`func (o *IdpAdapterAllOf) SetAttributeMapping(v IdpAdapterContractMapping)`

SetAttributeMapping sets AttributeMapping field to given value.

### HasAttributeMapping

`func (o *IdpAdapterAllOf) HasAttributeMapping() bool`

HasAttributeMapping returns a boolean if a field has been set.

### GetAttributeContract

`func (o *IdpAdapterAllOf) GetAttributeContract() IdpAdapterAttributeContract`

GetAttributeContract returns the AttributeContract field if non-nil, zero value otherwise.

### GetAttributeContractOk

`func (o *IdpAdapterAllOf) GetAttributeContractOk() (*IdpAdapterAttributeContract, bool)`

GetAttributeContractOk returns a tuple with the AttributeContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContract

`func (o *IdpAdapterAllOf) SetAttributeContract(v IdpAdapterAttributeContract)`

SetAttributeContract sets AttributeContract field to given value.

### HasAttributeContract

`func (o *IdpAdapterAllOf) HasAttributeContract() bool`

HasAttributeContract returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


