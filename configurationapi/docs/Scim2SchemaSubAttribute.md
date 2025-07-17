# Scim2SchemaSubAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the sub-attribute. | [optional] 
**Description** | Pointer to **string** | Description of the sub-attribute. | [optional] 
**Mutability** | Pointer to **string** | Mutability of the sub-attribute. | [optional] 
**Type** | Pointer to **string** | Type of the sub-attribute. | [optional] 
**Returned** | Pointer to **string** | Whether the sub-attribute is returned in the response. | [optional] 
**Uniqueness** | Pointer to **string** | The uniqueness of the sub-attribute. | [optional] 
**Required** | Pointer to **bool** | Whether the sub-attribute is required. | [optional] 
**CaseExact** | Pointer to **bool** | Whether the sub-attribute is case exact. | [optional] 

## Methods

### NewScim2SchemaSubAttribute

`func NewScim2SchemaSubAttribute() *Scim2SchemaSubAttribute`

NewScim2SchemaSubAttribute instantiates a new Scim2SchemaSubAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScim2SchemaSubAttributeWithDefaults

`func NewScim2SchemaSubAttributeWithDefaults() *Scim2SchemaSubAttribute`

NewScim2SchemaSubAttributeWithDefaults instantiates a new Scim2SchemaSubAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Scim2SchemaSubAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Scim2SchemaSubAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Scim2SchemaSubAttribute) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Scim2SchemaSubAttribute) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Scim2SchemaSubAttribute) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Scim2SchemaSubAttribute) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Scim2SchemaSubAttribute) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Scim2SchemaSubAttribute) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMutability

`func (o *Scim2SchemaSubAttribute) GetMutability() string`

GetMutability returns the Mutability field if non-nil, zero value otherwise.

### GetMutabilityOk

`func (o *Scim2SchemaSubAttribute) GetMutabilityOk() (*string, bool)`

GetMutabilityOk returns a tuple with the Mutability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutability

`func (o *Scim2SchemaSubAttribute) SetMutability(v string)`

SetMutability sets Mutability field to given value.

### HasMutability

`func (o *Scim2SchemaSubAttribute) HasMutability() bool`

HasMutability returns a boolean if a field has been set.

### GetType

`func (o *Scim2SchemaSubAttribute) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Scim2SchemaSubAttribute) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Scim2SchemaSubAttribute) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Scim2SchemaSubAttribute) HasType() bool`

HasType returns a boolean if a field has been set.

### GetReturned

`func (o *Scim2SchemaSubAttribute) GetReturned() string`

GetReturned returns the Returned field if non-nil, zero value otherwise.

### GetReturnedOk

`func (o *Scim2SchemaSubAttribute) GetReturnedOk() (*string, bool)`

GetReturnedOk returns a tuple with the Returned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturned

`func (o *Scim2SchemaSubAttribute) SetReturned(v string)`

SetReturned sets Returned field to given value.

### HasReturned

`func (o *Scim2SchemaSubAttribute) HasReturned() bool`

HasReturned returns a boolean if a field has been set.

### GetUniqueness

`func (o *Scim2SchemaSubAttribute) GetUniqueness() string`

GetUniqueness returns the Uniqueness field if non-nil, zero value otherwise.

### GetUniquenessOk

`func (o *Scim2SchemaSubAttribute) GetUniquenessOk() (*string, bool)`

GetUniquenessOk returns a tuple with the Uniqueness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueness

`func (o *Scim2SchemaSubAttribute) SetUniqueness(v string)`

SetUniqueness sets Uniqueness field to given value.

### HasUniqueness

`func (o *Scim2SchemaSubAttribute) HasUniqueness() bool`

HasUniqueness returns a boolean if a field has been set.

### GetRequired

`func (o *Scim2SchemaSubAttribute) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *Scim2SchemaSubAttribute) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *Scim2SchemaSubAttribute) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *Scim2SchemaSubAttribute) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetCaseExact

`func (o *Scim2SchemaSubAttribute) GetCaseExact() bool`

GetCaseExact returns the CaseExact field if non-nil, zero value otherwise.

### GetCaseExactOk

`func (o *Scim2SchemaSubAttribute) GetCaseExactOk() (*bool, bool)`

GetCaseExactOk returns a tuple with the CaseExact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseExact

`func (o *Scim2SchemaSubAttribute) SetCaseExact(v bool)`

SetCaseExact sets CaseExact field to given value.

### HasCaseExact

`func (o *Scim2SchemaSubAttribute) HasCaseExact() bool`

HasCaseExact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


