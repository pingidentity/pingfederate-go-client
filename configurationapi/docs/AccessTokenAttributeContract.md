# AccessTokenAttributeContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoreAttributes** | Pointer to [**[]AccessTokenAttribute**](AccessTokenAttribute.md) | A list of core token attributes that are associated with the access token management plugin type. This field is read-only and is ignored on POST/PUT. | [optional] 
**ExtendedAttributes** | Pointer to [**[]AccessTokenAttribute**](AccessTokenAttribute.md) | A list of additional token attributes that are associated with this access token management plugin instance. | [optional] 
**DefaultSubjectAttribute** | Pointer to **string** | Default subject attribute to use for audit logging when validating the access token. Blank value means to use USER_KEY attribute value after grant lookup. | [optional] 

## Methods

### NewAccessTokenAttributeContract

`func NewAccessTokenAttributeContract() *AccessTokenAttributeContract`

NewAccessTokenAttributeContract instantiates a new AccessTokenAttributeContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenAttributeContractWithDefaults

`func NewAccessTokenAttributeContractWithDefaults() *AccessTokenAttributeContract`

NewAccessTokenAttributeContractWithDefaults instantiates a new AccessTokenAttributeContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoreAttributes

`func (o *AccessTokenAttributeContract) GetCoreAttributes() []AccessTokenAttribute`

GetCoreAttributes returns the CoreAttributes field if non-nil, zero value otherwise.

### GetCoreAttributesOk

`func (o *AccessTokenAttributeContract) GetCoreAttributesOk() (*[]AccessTokenAttribute, bool)`

GetCoreAttributesOk returns a tuple with the CoreAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreAttributes

`func (o *AccessTokenAttributeContract) SetCoreAttributes(v []AccessTokenAttribute)`

SetCoreAttributes sets CoreAttributes field to given value.

### HasCoreAttributes

`func (o *AccessTokenAttributeContract) HasCoreAttributes() bool`

HasCoreAttributes returns a boolean if a field has been set.

### GetExtendedAttributes

`func (o *AccessTokenAttributeContract) GetExtendedAttributes() []AccessTokenAttribute`

GetExtendedAttributes returns the ExtendedAttributes field if non-nil, zero value otherwise.

### GetExtendedAttributesOk

`func (o *AccessTokenAttributeContract) GetExtendedAttributesOk() (*[]AccessTokenAttribute, bool)`

GetExtendedAttributesOk returns a tuple with the ExtendedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedAttributes

`func (o *AccessTokenAttributeContract) SetExtendedAttributes(v []AccessTokenAttribute)`

SetExtendedAttributes sets ExtendedAttributes field to given value.

### HasExtendedAttributes

`func (o *AccessTokenAttributeContract) HasExtendedAttributes() bool`

HasExtendedAttributes returns a boolean if a field has been set.

### GetDefaultSubjectAttribute

`func (o *AccessTokenAttributeContract) GetDefaultSubjectAttribute() string`

GetDefaultSubjectAttribute returns the DefaultSubjectAttribute field if non-nil, zero value otherwise.

### GetDefaultSubjectAttributeOk

`func (o *AccessTokenAttributeContract) GetDefaultSubjectAttributeOk() (*string, bool)`

GetDefaultSubjectAttributeOk returns a tuple with the DefaultSubjectAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSubjectAttribute

`func (o *AccessTokenAttributeContract) SetDefaultSubjectAttribute(v string)`

SetDefaultSubjectAttribute sets DefaultSubjectAttribute field to given value.

### HasDefaultSubjectAttribute

`func (o *AccessTokenAttributeContract) HasDefaultSubjectAttribute() bool`

HasDefaultSubjectAttribute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


