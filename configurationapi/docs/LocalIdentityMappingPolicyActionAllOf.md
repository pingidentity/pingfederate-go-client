# LocalIdentityMappingPolicyActionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalIdentityRef** | [**ResourceLink**](ResourceLink.md) |  | 
**InboundMapping** | Pointer to [**AttributeMapping**](AttributeMapping.md) |  | [optional] 
**OutboundAttributeMapping** | [**AttributeMapping**](AttributeMapping.md) |  | 

## Methods

### NewLocalIdentityMappingPolicyActionAllOf

`func NewLocalIdentityMappingPolicyActionAllOf(localIdentityRef ResourceLink, outboundAttributeMapping AttributeMapping, ) *LocalIdentityMappingPolicyActionAllOf`

NewLocalIdentityMappingPolicyActionAllOf instantiates a new LocalIdentityMappingPolicyActionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalIdentityMappingPolicyActionAllOfWithDefaults

`func NewLocalIdentityMappingPolicyActionAllOfWithDefaults() *LocalIdentityMappingPolicyActionAllOf`

NewLocalIdentityMappingPolicyActionAllOfWithDefaults instantiates a new LocalIdentityMappingPolicyActionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalIdentityRef

`func (o *LocalIdentityMappingPolicyActionAllOf) GetLocalIdentityRef() ResourceLink`

GetLocalIdentityRef returns the LocalIdentityRef field if non-nil, zero value otherwise.

### GetLocalIdentityRefOk

`func (o *LocalIdentityMappingPolicyActionAllOf) GetLocalIdentityRefOk() (*ResourceLink, bool)`

GetLocalIdentityRefOk returns a tuple with the LocalIdentityRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIdentityRef

`func (o *LocalIdentityMappingPolicyActionAllOf) SetLocalIdentityRef(v ResourceLink)`

SetLocalIdentityRef sets LocalIdentityRef field to given value.


### GetInboundMapping

`func (o *LocalIdentityMappingPolicyActionAllOf) GetInboundMapping() AttributeMapping`

GetInboundMapping returns the InboundMapping field if non-nil, zero value otherwise.

### GetInboundMappingOk

`func (o *LocalIdentityMappingPolicyActionAllOf) GetInboundMappingOk() (*AttributeMapping, bool)`

GetInboundMappingOk returns a tuple with the InboundMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundMapping

`func (o *LocalIdentityMappingPolicyActionAllOf) SetInboundMapping(v AttributeMapping)`

SetInboundMapping sets InboundMapping field to given value.

### HasInboundMapping

`func (o *LocalIdentityMappingPolicyActionAllOf) HasInboundMapping() bool`

HasInboundMapping returns a boolean if a field has been set.

### GetOutboundAttributeMapping

`func (o *LocalIdentityMappingPolicyActionAllOf) GetOutboundAttributeMapping() AttributeMapping`

GetOutboundAttributeMapping returns the OutboundAttributeMapping field if non-nil, zero value otherwise.

### GetOutboundAttributeMappingOk

`func (o *LocalIdentityMappingPolicyActionAllOf) GetOutboundAttributeMappingOk() (*AttributeMapping, bool)`

GetOutboundAttributeMappingOk returns a tuple with the OutboundAttributeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundAttributeMapping

`func (o *LocalIdentityMappingPolicyActionAllOf) SetOutboundAttributeMapping(v AttributeMapping)`

SetOutboundAttributeMapping sets OutboundAttributeMapping field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


