# SpAttributeQueryPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignResponse** | Pointer to **bool** | Sign the response. | [optional] 
**SignAssertion** | Pointer to **bool** | Sign the assertion. | [optional] 
**EncryptAssertion** | Pointer to **bool** | Encrypt the assertion. | [optional] 
**RequireSignedAttributeQuery** | Pointer to **bool** | Require signed attribute query. | [optional] 
**RequireEncryptedNameId** | Pointer to **bool** | Require an encrypted name identifier. | [optional] 

## Methods

### NewSpAttributeQueryPolicy

`func NewSpAttributeQueryPolicy() *SpAttributeQueryPolicy`

NewSpAttributeQueryPolicy instantiates a new SpAttributeQueryPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpAttributeQueryPolicyWithDefaults

`func NewSpAttributeQueryPolicyWithDefaults() *SpAttributeQueryPolicy`

NewSpAttributeQueryPolicyWithDefaults instantiates a new SpAttributeQueryPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignResponse

`func (o *SpAttributeQueryPolicy) GetSignResponse() bool`

GetSignResponse returns the SignResponse field if non-nil, zero value otherwise.

### GetSignResponseOk

`func (o *SpAttributeQueryPolicy) GetSignResponseOk() (*bool, bool)`

GetSignResponseOk returns a tuple with the SignResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignResponse

`func (o *SpAttributeQueryPolicy) SetSignResponse(v bool)`

SetSignResponse sets SignResponse field to given value.

### HasSignResponse

`func (o *SpAttributeQueryPolicy) HasSignResponse() bool`

HasSignResponse returns a boolean if a field has been set.

### GetSignAssertion

`func (o *SpAttributeQueryPolicy) GetSignAssertion() bool`

GetSignAssertion returns the SignAssertion field if non-nil, zero value otherwise.

### GetSignAssertionOk

`func (o *SpAttributeQueryPolicy) GetSignAssertionOk() (*bool, bool)`

GetSignAssertionOk returns a tuple with the SignAssertion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignAssertion

`func (o *SpAttributeQueryPolicy) SetSignAssertion(v bool)`

SetSignAssertion sets SignAssertion field to given value.

### HasSignAssertion

`func (o *SpAttributeQueryPolicy) HasSignAssertion() bool`

HasSignAssertion returns a boolean if a field has been set.

### GetEncryptAssertion

`func (o *SpAttributeQueryPolicy) GetEncryptAssertion() bool`

GetEncryptAssertion returns the EncryptAssertion field if non-nil, zero value otherwise.

### GetEncryptAssertionOk

`func (o *SpAttributeQueryPolicy) GetEncryptAssertionOk() (*bool, bool)`

GetEncryptAssertionOk returns a tuple with the EncryptAssertion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAssertion

`func (o *SpAttributeQueryPolicy) SetEncryptAssertion(v bool)`

SetEncryptAssertion sets EncryptAssertion field to given value.

### HasEncryptAssertion

`func (o *SpAttributeQueryPolicy) HasEncryptAssertion() bool`

HasEncryptAssertion returns a boolean if a field has been set.

### GetRequireSignedAttributeQuery

`func (o *SpAttributeQueryPolicy) GetRequireSignedAttributeQuery() bool`

GetRequireSignedAttributeQuery returns the RequireSignedAttributeQuery field if non-nil, zero value otherwise.

### GetRequireSignedAttributeQueryOk

`func (o *SpAttributeQueryPolicy) GetRequireSignedAttributeQueryOk() (*bool, bool)`

GetRequireSignedAttributeQueryOk returns a tuple with the RequireSignedAttributeQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSignedAttributeQuery

`func (o *SpAttributeQueryPolicy) SetRequireSignedAttributeQuery(v bool)`

SetRequireSignedAttributeQuery sets RequireSignedAttributeQuery field to given value.

### HasRequireSignedAttributeQuery

`func (o *SpAttributeQueryPolicy) HasRequireSignedAttributeQuery() bool`

HasRequireSignedAttributeQuery returns a boolean if a field has been set.

### GetRequireEncryptedNameId

`func (o *SpAttributeQueryPolicy) GetRequireEncryptedNameId() bool`

GetRequireEncryptedNameId returns the RequireEncryptedNameId field if non-nil, zero value otherwise.

### GetRequireEncryptedNameIdOk

`func (o *SpAttributeQueryPolicy) GetRequireEncryptedNameIdOk() (*bool, bool)`

GetRequireEncryptedNameIdOk returns a tuple with the RequireEncryptedNameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireEncryptedNameId

`func (o *SpAttributeQueryPolicy) SetRequireEncryptedNameId(v bool)`

SetRequireEncryptedNameId sets RequireEncryptedNameId field to given value.

### HasRequireEncryptedNameId

`func (o *SpAttributeQueryPolicy) HasRequireEncryptedNameId() bool`

HasRequireEncryptedNameId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


