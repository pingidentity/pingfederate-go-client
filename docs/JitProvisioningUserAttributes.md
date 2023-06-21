# JitProvisioningUserAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeContract** | Pointer to [**[]IdpBrowserSsoAttribute**](IdpBrowserSsoAttribute.md) | A list of user attributes that the IdP sends in the SAML assertion. | [optional] 
**DoAttributeQuery** | Pointer to **bool** | Specify whether to use only attributes from the SAML Assertion or retrieve additional attributes from the IdP. The default is false. | [optional] 

## Methods

### NewJitProvisioningUserAttributes

`func NewJitProvisioningUserAttributes() *JitProvisioningUserAttributes`

NewJitProvisioningUserAttributes instantiates a new JitProvisioningUserAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJitProvisioningUserAttributesWithDefaults

`func NewJitProvisioningUserAttributesWithDefaults() *JitProvisioningUserAttributes`

NewJitProvisioningUserAttributesWithDefaults instantiates a new JitProvisioningUserAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeContract

`func (o *JitProvisioningUserAttributes) GetAttributeContract() []IdpBrowserSsoAttribute`

GetAttributeContract returns the AttributeContract field if non-nil, zero value otherwise.

### GetAttributeContractOk

`func (o *JitProvisioningUserAttributes) GetAttributeContractOk() (*[]IdpBrowserSsoAttribute, bool)`

GetAttributeContractOk returns a tuple with the AttributeContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContract

`func (o *JitProvisioningUserAttributes) SetAttributeContract(v []IdpBrowserSsoAttribute)`

SetAttributeContract sets AttributeContract field to given value.

### HasAttributeContract

`func (o *JitProvisioningUserAttributes) HasAttributeContract() bool`

HasAttributeContract returns a boolean if a field has been set.

### GetDoAttributeQuery

`func (o *JitProvisioningUserAttributes) GetDoAttributeQuery() bool`

GetDoAttributeQuery returns the DoAttributeQuery field if non-nil, zero value otherwise.

### GetDoAttributeQueryOk

`func (o *JitProvisioningUserAttributes) GetDoAttributeQueryOk() (*bool, bool)`

GetDoAttributeQueryOk returns a tuple with the DoAttributeQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoAttributeQuery

`func (o *JitProvisioningUserAttributes) SetDoAttributeQuery(v bool)`

SetDoAttributeQuery sets DoAttributeQuery field to given value.

### HasDoAttributeQuery

`func (o *JitProvisioningUserAttributes) HasDoAttributeQuery() bool`

HasDoAttributeQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


