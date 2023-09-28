# IdpAttributeQueryPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequireSignedResponse** | Pointer to **bool** | Require signed response. | [optional] 
**RequireSignedAssertion** | Pointer to **bool** | Require signed assertion. | [optional] 
**RequireEncryptedAssertion** | Pointer to **bool** | Require encrypted assertion. | [optional] 
**SignAttributeQuery** | Pointer to **bool** | Sign the attribute query. | [optional] 
**EncryptNameId** | Pointer to **bool** | Encrypt the name identifier. | [optional] 
**MaskAttributeValues** | Pointer to **bool** | Mask attributes in log files. | [optional] 

## Methods

### NewIdpAttributeQueryPolicy

`func NewIdpAttributeQueryPolicy() *IdpAttributeQueryPolicy`

NewIdpAttributeQueryPolicy instantiates a new IdpAttributeQueryPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdpAttributeQueryPolicyWithDefaults

`func NewIdpAttributeQueryPolicyWithDefaults() *IdpAttributeQueryPolicy`

NewIdpAttributeQueryPolicyWithDefaults instantiates a new IdpAttributeQueryPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequireSignedResponse

`func (o *IdpAttributeQueryPolicy) GetRequireSignedResponse() bool`

GetRequireSignedResponse returns the RequireSignedResponse field if non-nil, zero value otherwise.

### GetRequireSignedResponseOk

`func (o *IdpAttributeQueryPolicy) GetRequireSignedResponseOk() (*bool, bool)`

GetRequireSignedResponseOk returns a tuple with the RequireSignedResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSignedResponse

`func (o *IdpAttributeQueryPolicy) SetRequireSignedResponse(v bool)`

SetRequireSignedResponse sets RequireSignedResponse field to given value.

### HasRequireSignedResponse

`func (o *IdpAttributeQueryPolicy) HasRequireSignedResponse() bool`

HasRequireSignedResponse returns a boolean if a field has been set.

### GetRequireSignedAssertion

`func (o *IdpAttributeQueryPolicy) GetRequireSignedAssertion() bool`

GetRequireSignedAssertion returns the RequireSignedAssertion field if non-nil, zero value otherwise.

### GetRequireSignedAssertionOk

`func (o *IdpAttributeQueryPolicy) GetRequireSignedAssertionOk() (*bool, bool)`

GetRequireSignedAssertionOk returns a tuple with the RequireSignedAssertion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSignedAssertion

`func (o *IdpAttributeQueryPolicy) SetRequireSignedAssertion(v bool)`

SetRequireSignedAssertion sets RequireSignedAssertion field to given value.

### HasRequireSignedAssertion

`func (o *IdpAttributeQueryPolicy) HasRequireSignedAssertion() bool`

HasRequireSignedAssertion returns a boolean if a field has been set.

### GetRequireEncryptedAssertion

`func (o *IdpAttributeQueryPolicy) GetRequireEncryptedAssertion() bool`

GetRequireEncryptedAssertion returns the RequireEncryptedAssertion field if non-nil, zero value otherwise.

### GetRequireEncryptedAssertionOk

`func (o *IdpAttributeQueryPolicy) GetRequireEncryptedAssertionOk() (*bool, bool)`

GetRequireEncryptedAssertionOk returns a tuple with the RequireEncryptedAssertion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireEncryptedAssertion

`func (o *IdpAttributeQueryPolicy) SetRequireEncryptedAssertion(v bool)`

SetRequireEncryptedAssertion sets RequireEncryptedAssertion field to given value.

### HasRequireEncryptedAssertion

`func (o *IdpAttributeQueryPolicy) HasRequireEncryptedAssertion() bool`

HasRequireEncryptedAssertion returns a boolean if a field has been set.

### GetSignAttributeQuery

`func (o *IdpAttributeQueryPolicy) GetSignAttributeQuery() bool`

GetSignAttributeQuery returns the SignAttributeQuery field if non-nil, zero value otherwise.

### GetSignAttributeQueryOk

`func (o *IdpAttributeQueryPolicy) GetSignAttributeQueryOk() (*bool, bool)`

GetSignAttributeQueryOk returns a tuple with the SignAttributeQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignAttributeQuery

`func (o *IdpAttributeQueryPolicy) SetSignAttributeQuery(v bool)`

SetSignAttributeQuery sets SignAttributeQuery field to given value.

### HasSignAttributeQuery

`func (o *IdpAttributeQueryPolicy) HasSignAttributeQuery() bool`

HasSignAttributeQuery returns a boolean if a field has been set.

### GetEncryptNameId

`func (o *IdpAttributeQueryPolicy) GetEncryptNameId() bool`

GetEncryptNameId returns the EncryptNameId field if non-nil, zero value otherwise.

### GetEncryptNameIdOk

`func (o *IdpAttributeQueryPolicy) GetEncryptNameIdOk() (*bool, bool)`

GetEncryptNameIdOk returns a tuple with the EncryptNameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptNameId

`func (o *IdpAttributeQueryPolicy) SetEncryptNameId(v bool)`

SetEncryptNameId sets EncryptNameId field to given value.

### HasEncryptNameId

`func (o *IdpAttributeQueryPolicy) HasEncryptNameId() bool`

HasEncryptNameId returns a boolean if a field has been set.

### GetMaskAttributeValues

`func (o *IdpAttributeQueryPolicy) GetMaskAttributeValues() bool`

GetMaskAttributeValues returns the MaskAttributeValues field if non-nil, zero value otherwise.

### GetMaskAttributeValuesOk

`func (o *IdpAttributeQueryPolicy) GetMaskAttributeValuesOk() (*bool, bool)`

GetMaskAttributeValuesOk returns a tuple with the MaskAttributeValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskAttributeValues

`func (o *IdpAttributeQueryPolicy) SetMaskAttributeValues(v bool)`

SetMaskAttributeValues sets MaskAttributeValues field to given value.

### HasMaskAttributeValues

`func (o *IdpAttributeQueryPolicy) HasMaskAttributeValues() bool`

HasMaskAttributeValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


