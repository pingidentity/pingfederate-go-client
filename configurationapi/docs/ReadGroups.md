# ReadGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeContract** | [**IdpInboundProvisioningAttributeContract**](IdpInboundProvisioningAttributeContract.md) |  | 
**Attributes** | [**[]Attribute**](Attribute.md) | A list of LDAP data store attributes to populate a response to a user-provisioning request. | 
**AttributeFulfillment** | [**map[string]AttributeFulfillmentValue**](AttributeFulfillmentValue.md) | A list of user repository mappings from attribute names to their fulfillment values. | 

## Methods

### NewReadGroups

`func NewReadGroups(attributeContract IdpInboundProvisioningAttributeContract, attributes []Attribute, attributeFulfillment map[string]AttributeFulfillmentValue, ) *ReadGroups`

NewReadGroups instantiates a new ReadGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadGroupsWithDefaults

`func NewReadGroupsWithDefaults() *ReadGroups`

NewReadGroupsWithDefaults instantiates a new ReadGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeContract

`func (o *ReadGroups) GetAttributeContract() IdpInboundProvisioningAttributeContract`

GetAttributeContract returns the AttributeContract field if non-nil, zero value otherwise.

### GetAttributeContractOk

`func (o *ReadGroups) GetAttributeContractOk() (*IdpInboundProvisioningAttributeContract, bool)`

GetAttributeContractOk returns a tuple with the AttributeContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContract

`func (o *ReadGroups) SetAttributeContract(v IdpInboundProvisioningAttributeContract)`

SetAttributeContract sets AttributeContract field to given value.


### GetAttributes

`func (o *ReadGroups) GetAttributes() []Attribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ReadGroups) GetAttributesOk() (*[]Attribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ReadGroups) SetAttributes(v []Attribute)`

SetAttributes sets Attributes field to given value.


### GetAttributeFulfillment

`func (o *ReadGroups) GetAttributeFulfillment() map[string]AttributeFulfillmentValue`

GetAttributeFulfillment returns the AttributeFulfillment field if non-nil, zero value otherwise.

### GetAttributeFulfillmentOk

`func (o *ReadGroups) GetAttributeFulfillmentOk() (*map[string]AttributeFulfillmentValue, bool)`

GetAttributeFulfillmentOk returns a tuple with the AttributeFulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeFulfillment

`func (o *ReadGroups) SetAttributeFulfillment(v map[string]AttributeFulfillmentValue)`

SetAttributeFulfillment sets AttributeFulfillment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


