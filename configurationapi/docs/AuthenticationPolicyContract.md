# AuthenticationPolicyContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The persistent, unique ID for the authentication policy contract. It can be any combination of [a-zA-Z0-9._-]. This property is system-assigned if not specified. | [optional] 
**Name** | Pointer to **string** | The Authentication Policy Contract Name. Name is unique. | [optional] 
**CoreAttributes** | Pointer to [**[]AuthenticationPolicyContractAttribute**](AuthenticationPolicyContractAttribute.md) | A list of read-only assertion attributes (for example, subject) that are automatically populated by PingFederate. | [optional] 
**ExtendedAttributes** | Pointer to [**[]AuthenticationPolicyContractAttribute**](AuthenticationPolicyContractAttribute.md) | A list of additional attributes as needed. | [optional] 
**LastModified** | Pointer to **time.Time** | The time at which the authentication policy contract was last changed. This property is read only and is ignored on PUT and POST requests. | [optional] 

## Methods

### NewAuthenticationPolicyContract

`func NewAuthenticationPolicyContract() *AuthenticationPolicyContract`

NewAuthenticationPolicyContract instantiates a new AuthenticationPolicyContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationPolicyContractWithDefaults

`func NewAuthenticationPolicyContractWithDefaults() *AuthenticationPolicyContract`

NewAuthenticationPolicyContractWithDefaults instantiates a new AuthenticationPolicyContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthenticationPolicyContract) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthenticationPolicyContract) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthenticationPolicyContract) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthenticationPolicyContract) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AuthenticationPolicyContract) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthenticationPolicyContract) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthenticationPolicyContract) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthenticationPolicyContract) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCoreAttributes

`func (o *AuthenticationPolicyContract) GetCoreAttributes() []AuthenticationPolicyContractAttribute`

GetCoreAttributes returns the CoreAttributes field if non-nil, zero value otherwise.

### GetCoreAttributesOk

`func (o *AuthenticationPolicyContract) GetCoreAttributesOk() (*[]AuthenticationPolicyContractAttribute, bool)`

GetCoreAttributesOk returns a tuple with the CoreAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreAttributes

`func (o *AuthenticationPolicyContract) SetCoreAttributes(v []AuthenticationPolicyContractAttribute)`

SetCoreAttributes sets CoreAttributes field to given value.

### HasCoreAttributes

`func (o *AuthenticationPolicyContract) HasCoreAttributes() bool`

HasCoreAttributes returns a boolean if a field has been set.

### GetExtendedAttributes

`func (o *AuthenticationPolicyContract) GetExtendedAttributes() []AuthenticationPolicyContractAttribute`

GetExtendedAttributes returns the ExtendedAttributes field if non-nil, zero value otherwise.

### GetExtendedAttributesOk

`func (o *AuthenticationPolicyContract) GetExtendedAttributesOk() (*[]AuthenticationPolicyContractAttribute, bool)`

GetExtendedAttributesOk returns a tuple with the ExtendedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedAttributes

`func (o *AuthenticationPolicyContract) SetExtendedAttributes(v []AuthenticationPolicyContractAttribute)`

SetExtendedAttributes sets ExtendedAttributes field to given value.

### HasExtendedAttributes

`func (o *AuthenticationPolicyContract) HasExtendedAttributes() bool`

HasExtendedAttributes returns a boolean if a field has been set.

### GetLastModified

`func (o *AuthenticationPolicyContract) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *AuthenticationPolicyContract) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *AuthenticationPolicyContract) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *AuthenticationPolicyContract) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


