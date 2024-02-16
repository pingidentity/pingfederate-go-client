# ApcMappingPolicyAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The authentication selection type. | 
**Context** | Pointer to **string** | The result context. | [optional] 
**AuthenticationPolicyContractRef** | [**ResourceLink**](ResourceLink.md) |  | 
**AttributeMapping** | [**AttributeMapping**](AttributeMapping.md) |  | 

## Methods

### NewApcMappingPolicyAction

`func NewApcMappingPolicyAction(type_ string, authenticationPolicyContractRef ResourceLink, attributeMapping AttributeMapping, ) *ApcMappingPolicyAction`

NewApcMappingPolicyAction instantiates a new ApcMappingPolicyAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApcMappingPolicyActionWithDefaults

`func NewApcMappingPolicyActionWithDefaults() *ApcMappingPolicyAction`

NewApcMappingPolicyActionWithDefaults instantiates a new ApcMappingPolicyAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApcMappingPolicyAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApcMappingPolicyAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApcMappingPolicyAction) SetType(v string)`

SetType sets Type field to given value.


### GetContext

`func (o *ApcMappingPolicyAction) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ApcMappingPolicyAction) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ApcMappingPolicyAction) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *ApcMappingPolicyAction) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetAuthenticationPolicyContractRef

`func (o *ApcMappingPolicyAction) GetAuthenticationPolicyContractRef() ResourceLink`

GetAuthenticationPolicyContractRef returns the AuthenticationPolicyContractRef field if non-nil, zero value otherwise.

### GetAuthenticationPolicyContractRefOk

`func (o *ApcMappingPolicyAction) GetAuthenticationPolicyContractRefOk() (*ResourceLink, bool)`

GetAuthenticationPolicyContractRefOk returns a tuple with the AuthenticationPolicyContractRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationPolicyContractRef

`func (o *ApcMappingPolicyAction) SetAuthenticationPolicyContractRef(v ResourceLink)`

SetAuthenticationPolicyContractRef sets AuthenticationPolicyContractRef field to given value.


### GetAttributeMapping

`func (o *ApcMappingPolicyAction) GetAttributeMapping() AttributeMapping`

GetAttributeMapping returns the AttributeMapping field if non-nil, zero value otherwise.

### GetAttributeMappingOk

`func (o *ApcMappingPolicyAction) GetAttributeMappingOk() (*AttributeMapping, bool)`

GetAttributeMappingOk returns a tuple with the AttributeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeMapping

`func (o *ApcMappingPolicyAction) SetAttributeMapping(v AttributeMapping)`

SetAttributeMapping sets AttributeMapping field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


