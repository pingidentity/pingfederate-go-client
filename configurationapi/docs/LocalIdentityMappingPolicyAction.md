# LocalIdentityMappingPolicyAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The authentication selection type. | 
**Context** | Pointer to **string** | The result context. | [optional] 
**LocalIdentityRef** | [**ResourceLink**](ResourceLink.md) |  | 
**InboundMapping** | Pointer to [**AttributeMapping**](AttributeMapping.md) |  | [optional] 
**OutboundAttributeMapping** | [**AttributeMapping**](AttributeMapping.md) |  | 

## Methods

### NewLocalIdentityMappingPolicyAction

`func NewLocalIdentityMappingPolicyAction(type_ string, localIdentityRef ResourceLink, outboundAttributeMapping AttributeMapping, ) *LocalIdentityMappingPolicyAction`

NewLocalIdentityMappingPolicyAction instantiates a new LocalIdentityMappingPolicyAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalIdentityMappingPolicyActionWithDefaults

`func NewLocalIdentityMappingPolicyActionWithDefaults() *LocalIdentityMappingPolicyAction`

NewLocalIdentityMappingPolicyActionWithDefaults instantiates a new LocalIdentityMappingPolicyAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LocalIdentityMappingPolicyAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LocalIdentityMappingPolicyAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LocalIdentityMappingPolicyAction) SetType(v string)`

SetType sets Type field to given value.


### GetContext

`func (o *LocalIdentityMappingPolicyAction) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *LocalIdentityMappingPolicyAction) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *LocalIdentityMappingPolicyAction) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *LocalIdentityMappingPolicyAction) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetLocalIdentityRef

`func (o *LocalIdentityMappingPolicyAction) GetLocalIdentityRef() ResourceLink`

GetLocalIdentityRef returns the LocalIdentityRef field if non-nil, zero value otherwise.

### GetLocalIdentityRefOk

`func (o *LocalIdentityMappingPolicyAction) GetLocalIdentityRefOk() (*ResourceLink, bool)`

GetLocalIdentityRefOk returns a tuple with the LocalIdentityRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIdentityRef

`func (o *LocalIdentityMappingPolicyAction) SetLocalIdentityRef(v ResourceLink)`

SetLocalIdentityRef sets LocalIdentityRef field to given value.


### GetInboundMapping

`func (o *LocalIdentityMappingPolicyAction) GetInboundMapping() AttributeMapping`

GetInboundMapping returns the InboundMapping field if non-nil, zero value otherwise.

### GetInboundMappingOk

`func (o *LocalIdentityMappingPolicyAction) GetInboundMappingOk() (*AttributeMapping, bool)`

GetInboundMappingOk returns a tuple with the InboundMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundMapping

`func (o *LocalIdentityMappingPolicyAction) SetInboundMapping(v AttributeMapping)`

SetInboundMapping sets InboundMapping field to given value.

### HasInboundMapping

`func (o *LocalIdentityMappingPolicyAction) HasInboundMapping() bool`

HasInboundMapping returns a boolean if a field has been set.

### GetOutboundAttributeMapping

`func (o *LocalIdentityMappingPolicyAction) GetOutboundAttributeMapping() AttributeMapping`

GetOutboundAttributeMapping returns the OutboundAttributeMapping field if non-nil, zero value otherwise.

### GetOutboundAttributeMappingOk

`func (o *LocalIdentityMappingPolicyAction) GetOutboundAttributeMappingOk() (*AttributeMapping, bool)`

GetOutboundAttributeMappingOk returns a tuple with the OutboundAttributeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundAttributeMapping

`func (o *LocalIdentityMappingPolicyAction) SetOutboundAttributeMapping(v AttributeMapping)`

SetOutboundAttributeMapping sets OutboundAttributeMapping field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


