# Scim2SchemaAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the attribute. | [optional] 
**Description** | Pointer to **string** | Description of the attribute. | [optional] 
**Mutability** | Pointer to **string** | Mutability of the attribute. | [optional] 
**Type** | Pointer to **string** | Type of the attribute. | [optional] 
**SubAttributes** | Pointer to [**[]Scim2SchemaSubAttribute**](Scim2SchemaSubAttribute.md) | List of sub-attributes for complex attributes. | [optional] 
**Returned** | Pointer to **string** | Whether the attribute is returned in the response. | [optional] 
**Uniqueness** | Pointer to **string** | The uniqueness of the attribute. | [optional] 
**Required** | Pointer to **bool** | Whether the attribute is required. | [optional] 
**CaseExact** | Pointer to **bool** | Whether the attribute is case exact. | [optional] 
**MultiValued** | Pointer to **bool** | Indicates whether the attribute is multi-valued. | [optional] 
**CanonicalValues** | Pointer to **[]string** | List of canonical values for multi-valued attributes. | [optional] 

## Methods

### NewScim2SchemaAttribute

`func NewScim2SchemaAttribute() *Scim2SchemaAttribute`

NewScim2SchemaAttribute instantiates a new Scim2SchemaAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScim2SchemaAttributeWithDefaults

`func NewScim2SchemaAttributeWithDefaults() *Scim2SchemaAttribute`

NewScim2SchemaAttributeWithDefaults instantiates a new Scim2SchemaAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Scim2SchemaAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Scim2SchemaAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Scim2SchemaAttribute) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Scim2SchemaAttribute) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Scim2SchemaAttribute) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Scim2SchemaAttribute) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Scim2SchemaAttribute) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Scim2SchemaAttribute) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMutability

`func (o *Scim2SchemaAttribute) GetMutability() string`

GetMutability returns the Mutability field if non-nil, zero value otherwise.

### GetMutabilityOk

`func (o *Scim2SchemaAttribute) GetMutabilityOk() (*string, bool)`

GetMutabilityOk returns a tuple with the Mutability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutability

`func (o *Scim2SchemaAttribute) SetMutability(v string)`

SetMutability sets Mutability field to given value.

### HasMutability

`func (o *Scim2SchemaAttribute) HasMutability() bool`

HasMutability returns a boolean if a field has been set.

### GetType

`func (o *Scim2SchemaAttribute) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Scim2SchemaAttribute) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Scim2SchemaAttribute) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Scim2SchemaAttribute) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubAttributes

`func (o *Scim2SchemaAttribute) GetSubAttributes() []Scim2SchemaSubAttribute`

GetSubAttributes returns the SubAttributes field if non-nil, zero value otherwise.

### GetSubAttributesOk

`func (o *Scim2SchemaAttribute) GetSubAttributesOk() (*[]Scim2SchemaSubAttribute, bool)`

GetSubAttributesOk returns a tuple with the SubAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAttributes

`func (o *Scim2SchemaAttribute) SetSubAttributes(v []Scim2SchemaSubAttribute)`

SetSubAttributes sets SubAttributes field to given value.

### HasSubAttributes

`func (o *Scim2SchemaAttribute) HasSubAttributes() bool`

HasSubAttributes returns a boolean if a field has been set.

### GetReturned

`func (o *Scim2SchemaAttribute) GetReturned() string`

GetReturned returns the Returned field if non-nil, zero value otherwise.

### GetReturnedOk

`func (o *Scim2SchemaAttribute) GetReturnedOk() (*string, bool)`

GetReturnedOk returns a tuple with the Returned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturned

`func (o *Scim2SchemaAttribute) SetReturned(v string)`

SetReturned sets Returned field to given value.

### HasReturned

`func (o *Scim2SchemaAttribute) HasReturned() bool`

HasReturned returns a boolean if a field has been set.

### GetUniqueness

`func (o *Scim2SchemaAttribute) GetUniqueness() string`

GetUniqueness returns the Uniqueness field if non-nil, zero value otherwise.

### GetUniquenessOk

`func (o *Scim2SchemaAttribute) GetUniquenessOk() (*string, bool)`

GetUniquenessOk returns a tuple with the Uniqueness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueness

`func (o *Scim2SchemaAttribute) SetUniqueness(v string)`

SetUniqueness sets Uniqueness field to given value.

### HasUniqueness

`func (o *Scim2SchemaAttribute) HasUniqueness() bool`

HasUniqueness returns a boolean if a field has been set.

### GetRequired

`func (o *Scim2SchemaAttribute) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *Scim2SchemaAttribute) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *Scim2SchemaAttribute) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *Scim2SchemaAttribute) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetCaseExact

`func (o *Scim2SchemaAttribute) GetCaseExact() bool`

GetCaseExact returns the CaseExact field if non-nil, zero value otherwise.

### GetCaseExactOk

`func (o *Scim2SchemaAttribute) GetCaseExactOk() (*bool, bool)`

GetCaseExactOk returns a tuple with the CaseExact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseExact

`func (o *Scim2SchemaAttribute) SetCaseExact(v bool)`

SetCaseExact sets CaseExact field to given value.

### HasCaseExact

`func (o *Scim2SchemaAttribute) HasCaseExact() bool`

HasCaseExact returns a boolean if a field has been set.

### GetMultiValued

`func (o *Scim2SchemaAttribute) GetMultiValued() bool`

GetMultiValued returns the MultiValued field if non-nil, zero value otherwise.

### GetMultiValuedOk

`func (o *Scim2SchemaAttribute) GetMultiValuedOk() (*bool, bool)`

GetMultiValuedOk returns a tuple with the MultiValued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValued

`func (o *Scim2SchemaAttribute) SetMultiValued(v bool)`

SetMultiValued sets MultiValued field to given value.

### HasMultiValued

`func (o *Scim2SchemaAttribute) HasMultiValued() bool`

HasMultiValued returns a boolean if a field has been set.

### GetCanonicalValues

`func (o *Scim2SchemaAttribute) GetCanonicalValues() []string`

GetCanonicalValues returns the CanonicalValues field if non-nil, zero value otherwise.

### GetCanonicalValuesOk

`func (o *Scim2SchemaAttribute) GetCanonicalValuesOk() (*[]string, bool)`

GetCanonicalValuesOk returns a tuple with the CanonicalValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalValues

`func (o *Scim2SchemaAttribute) SetCanonicalValues(v []string)`

SetCanonicalValues sets CanonicalValues field to given value.

### HasCanonicalValues

`func (o *Scim2SchemaAttribute) HasCanonicalValues() bool`

HasCanonicalValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


