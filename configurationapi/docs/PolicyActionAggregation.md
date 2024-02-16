# PolicyActionAggregation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The authentication selection type. | 
**Context** | Pointer to **string** | The result context. | [optional] 
**AuthenticationPolicyContractRef** | [**ResourceLink**](ResourceLink.md) |  | 
**AttributeMapping** | [**AttributeMapping**](AttributeMapping.md) |  | 
**AuthenticationSelectorRef** | [**ResourceLink**](ResourceLink.md) |  | 
**AttributeRules** | Pointer to [**AttributeRules**](AttributeRules.md) |  | [optional] 
**AuthenticationSource** | [**AuthenticationSource**](AuthenticationSource.md) |  | 
**InputUserIdMapping** | Pointer to [**AttributeFulfillmentValue**](AttributeFulfillmentValue.md) |  | [optional] 
**UserIdAuthenticated** | Pointer to **bool** | Indicates whether the user ID obtained by the user ID mapping is authenticated. | [optional] 
**Fragment** | [**ResourceLink**](ResourceLink.md) |  | 
**FragmentMapping** | Pointer to [**AttributeMapping**](AttributeMapping.md) |  | [optional] 
**LocalIdentityRef** | [**ResourceLink**](ResourceLink.md) |  | 
**InboundMapping** | Pointer to [**AttributeMapping**](AttributeMapping.md) |  | [optional] 
**OutboundAttributeMapping** | [**AttributeMapping**](AttributeMapping.md) |  | 

## Methods

### NewPolicyActionAggregation

`func NewPolicyActionAggregation(type_ string, authenticationPolicyContractRef ResourceLink, attributeMapping AttributeMapping, authenticationSelectorRef ResourceLink, authenticationSource AuthenticationSource, fragment ResourceLink, localIdentityRef ResourceLink, outboundAttributeMapping AttributeMapping, ) *PolicyActionAggregation`

NewPolicyActionAggregation instantiates a new PolicyActionAggregation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyActionAggregationWithDefaults

`func NewPolicyActionAggregationWithDefaults() *PolicyActionAggregation`

NewPolicyActionAggregationWithDefaults instantiates a new PolicyActionAggregation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PolicyActionAggregation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PolicyActionAggregation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PolicyActionAggregation) SetType(v string)`

SetType sets Type field to given value.


### GetContext

`func (o *PolicyActionAggregation) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PolicyActionAggregation) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PolicyActionAggregation) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *PolicyActionAggregation) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetAuthenticationPolicyContractRef

`func (o *PolicyActionAggregation) GetAuthenticationPolicyContractRef() ResourceLink`

GetAuthenticationPolicyContractRef returns the AuthenticationPolicyContractRef field if non-nil, zero value otherwise.

### GetAuthenticationPolicyContractRefOk

`func (o *PolicyActionAggregation) GetAuthenticationPolicyContractRefOk() (*ResourceLink, bool)`

GetAuthenticationPolicyContractRefOk returns a tuple with the AuthenticationPolicyContractRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationPolicyContractRef

`func (o *PolicyActionAggregation) SetAuthenticationPolicyContractRef(v ResourceLink)`

SetAuthenticationPolicyContractRef sets AuthenticationPolicyContractRef field to given value.


### GetAttributeMapping

`func (o *PolicyActionAggregation) GetAttributeMapping() AttributeMapping`

GetAttributeMapping returns the AttributeMapping field if non-nil, zero value otherwise.

### GetAttributeMappingOk

`func (o *PolicyActionAggregation) GetAttributeMappingOk() (*AttributeMapping, bool)`

GetAttributeMappingOk returns a tuple with the AttributeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeMapping

`func (o *PolicyActionAggregation) SetAttributeMapping(v AttributeMapping)`

SetAttributeMapping sets AttributeMapping field to given value.


### GetAuthenticationSelectorRef

`func (o *PolicyActionAggregation) GetAuthenticationSelectorRef() ResourceLink`

GetAuthenticationSelectorRef returns the AuthenticationSelectorRef field if non-nil, zero value otherwise.

### GetAuthenticationSelectorRefOk

`func (o *PolicyActionAggregation) GetAuthenticationSelectorRefOk() (*ResourceLink, bool)`

GetAuthenticationSelectorRefOk returns a tuple with the AuthenticationSelectorRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationSelectorRef

`func (o *PolicyActionAggregation) SetAuthenticationSelectorRef(v ResourceLink)`

SetAuthenticationSelectorRef sets AuthenticationSelectorRef field to given value.


### GetAttributeRules

`func (o *PolicyActionAggregation) GetAttributeRules() AttributeRules`

GetAttributeRules returns the AttributeRules field if non-nil, zero value otherwise.

### GetAttributeRulesOk

`func (o *PolicyActionAggregation) GetAttributeRulesOk() (*AttributeRules, bool)`

GetAttributeRulesOk returns a tuple with the AttributeRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeRules

`func (o *PolicyActionAggregation) SetAttributeRules(v AttributeRules)`

SetAttributeRules sets AttributeRules field to given value.

### HasAttributeRules

`func (o *PolicyActionAggregation) HasAttributeRules() bool`

HasAttributeRules returns a boolean if a field has been set.

### GetAuthenticationSource

`func (o *PolicyActionAggregation) GetAuthenticationSource() AuthenticationSource`

GetAuthenticationSource returns the AuthenticationSource field if non-nil, zero value otherwise.

### GetAuthenticationSourceOk

`func (o *PolicyActionAggregation) GetAuthenticationSourceOk() (*AuthenticationSource, bool)`

GetAuthenticationSourceOk returns a tuple with the AuthenticationSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationSource

`func (o *PolicyActionAggregation) SetAuthenticationSource(v AuthenticationSource)`

SetAuthenticationSource sets AuthenticationSource field to given value.


### GetInputUserIdMapping

`func (o *PolicyActionAggregation) GetInputUserIdMapping() AttributeFulfillmentValue`

GetInputUserIdMapping returns the InputUserIdMapping field if non-nil, zero value otherwise.

### GetInputUserIdMappingOk

`func (o *PolicyActionAggregation) GetInputUserIdMappingOk() (*AttributeFulfillmentValue, bool)`

GetInputUserIdMappingOk returns a tuple with the InputUserIdMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputUserIdMapping

`func (o *PolicyActionAggregation) SetInputUserIdMapping(v AttributeFulfillmentValue)`

SetInputUserIdMapping sets InputUserIdMapping field to given value.

### HasInputUserIdMapping

`func (o *PolicyActionAggregation) HasInputUserIdMapping() bool`

HasInputUserIdMapping returns a boolean if a field has been set.

### GetUserIdAuthenticated

`func (o *PolicyActionAggregation) GetUserIdAuthenticated() bool`

GetUserIdAuthenticated returns the UserIdAuthenticated field if non-nil, zero value otherwise.

### GetUserIdAuthenticatedOk

`func (o *PolicyActionAggregation) GetUserIdAuthenticatedOk() (*bool, bool)`

GetUserIdAuthenticatedOk returns a tuple with the UserIdAuthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdAuthenticated

`func (o *PolicyActionAggregation) SetUserIdAuthenticated(v bool)`

SetUserIdAuthenticated sets UserIdAuthenticated field to given value.

### HasUserIdAuthenticated

`func (o *PolicyActionAggregation) HasUserIdAuthenticated() bool`

HasUserIdAuthenticated returns a boolean if a field has been set.

### GetFragment

`func (o *PolicyActionAggregation) GetFragment() ResourceLink`

GetFragment returns the Fragment field if non-nil, zero value otherwise.

### GetFragmentOk

`func (o *PolicyActionAggregation) GetFragmentOk() (*ResourceLink, bool)`

GetFragmentOk returns a tuple with the Fragment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragment

`func (o *PolicyActionAggregation) SetFragment(v ResourceLink)`

SetFragment sets Fragment field to given value.


### GetFragmentMapping

`func (o *PolicyActionAggregation) GetFragmentMapping() AttributeMapping`

GetFragmentMapping returns the FragmentMapping field if non-nil, zero value otherwise.

### GetFragmentMappingOk

`func (o *PolicyActionAggregation) GetFragmentMappingOk() (*AttributeMapping, bool)`

GetFragmentMappingOk returns a tuple with the FragmentMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragmentMapping

`func (o *PolicyActionAggregation) SetFragmentMapping(v AttributeMapping)`

SetFragmentMapping sets FragmentMapping field to given value.

### HasFragmentMapping

`func (o *PolicyActionAggregation) HasFragmentMapping() bool`

HasFragmentMapping returns a boolean if a field has been set.

### GetLocalIdentityRef

`func (o *PolicyActionAggregation) GetLocalIdentityRef() ResourceLink`

GetLocalIdentityRef returns the LocalIdentityRef field if non-nil, zero value otherwise.

### GetLocalIdentityRefOk

`func (o *PolicyActionAggregation) GetLocalIdentityRefOk() (*ResourceLink, bool)`

GetLocalIdentityRefOk returns a tuple with the LocalIdentityRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIdentityRef

`func (o *PolicyActionAggregation) SetLocalIdentityRef(v ResourceLink)`

SetLocalIdentityRef sets LocalIdentityRef field to given value.


### GetInboundMapping

`func (o *PolicyActionAggregation) GetInboundMapping() AttributeMapping`

GetInboundMapping returns the InboundMapping field if non-nil, zero value otherwise.

### GetInboundMappingOk

`func (o *PolicyActionAggregation) GetInboundMappingOk() (*AttributeMapping, bool)`

GetInboundMappingOk returns a tuple with the InboundMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundMapping

`func (o *PolicyActionAggregation) SetInboundMapping(v AttributeMapping)`

SetInboundMapping sets InboundMapping field to given value.

### HasInboundMapping

`func (o *PolicyActionAggregation) HasInboundMapping() bool`

HasInboundMapping returns a boolean if a field has been set.

### GetOutboundAttributeMapping

`func (o *PolicyActionAggregation) GetOutboundAttributeMapping() AttributeMapping`

GetOutboundAttributeMapping returns the OutboundAttributeMapping field if non-nil, zero value otherwise.

### GetOutboundAttributeMappingOk

`func (o *PolicyActionAggregation) GetOutboundAttributeMappingOk() (*AttributeMapping, bool)`

GetOutboundAttributeMappingOk returns a tuple with the OutboundAttributeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundAttributeMapping

`func (o *PolicyActionAggregation) SetOutboundAttributeMapping(v AttributeMapping)`

SetOutboundAttributeMapping sets OutboundAttributeMapping field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


