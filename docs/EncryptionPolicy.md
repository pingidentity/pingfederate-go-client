# EncryptionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptAssertion** | Pointer to **bool** | Whether the outgoing SAML assertion will be encrypted. | [optional] 
**EncryptedAttributes** | Pointer to **[]string** | The list of outgoing SAML assertion attributes that will be encrypted. The &#39;encryptAssertion&#39; property takes precedence over this. | [optional] 
**EncryptSloSubjectNameId** | Pointer to **bool** | Encrypt the name-identifier attribute in outbound SLO messages.  This can be set if the name id is encrypted. | [optional] 
**SloSubjectNameIDEncrypted** | Pointer to **bool** | Allow the encryption of the name-identifier attribute for inbound SLO messages. This can be set if SP initiated SLO is enabled. | [optional] 

## Methods

### NewEncryptionPolicy

`func NewEncryptionPolicy() *EncryptionPolicy`

NewEncryptionPolicy instantiates a new EncryptionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptionPolicyWithDefaults

`func NewEncryptionPolicyWithDefaults() *EncryptionPolicy`

NewEncryptionPolicyWithDefaults instantiates a new EncryptionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptAssertion

`func (o *EncryptionPolicy) GetEncryptAssertion() bool`

GetEncryptAssertion returns the EncryptAssertion field if non-nil, zero value otherwise.

### GetEncryptAssertionOk

`func (o *EncryptionPolicy) GetEncryptAssertionOk() (*bool, bool)`

GetEncryptAssertionOk returns a tuple with the EncryptAssertion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAssertion

`func (o *EncryptionPolicy) SetEncryptAssertion(v bool)`

SetEncryptAssertion sets EncryptAssertion field to given value.

### HasEncryptAssertion

`func (o *EncryptionPolicy) HasEncryptAssertion() bool`

HasEncryptAssertion returns a boolean if a field has been set.

### GetEncryptedAttributes

`func (o *EncryptionPolicy) GetEncryptedAttributes() []string`

GetEncryptedAttributes returns the EncryptedAttributes field if non-nil, zero value otherwise.

### GetEncryptedAttributesOk

`func (o *EncryptionPolicy) GetEncryptedAttributesOk() (*[]string, bool)`

GetEncryptedAttributesOk returns a tuple with the EncryptedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedAttributes

`func (o *EncryptionPolicy) SetEncryptedAttributes(v []string)`

SetEncryptedAttributes sets EncryptedAttributes field to given value.

### HasEncryptedAttributes

`func (o *EncryptionPolicy) HasEncryptedAttributes() bool`

HasEncryptedAttributes returns a boolean if a field has been set.

### GetEncryptSloSubjectNameId

`func (o *EncryptionPolicy) GetEncryptSloSubjectNameId() bool`

GetEncryptSloSubjectNameId returns the EncryptSloSubjectNameId field if non-nil, zero value otherwise.

### GetEncryptSloSubjectNameIdOk

`func (o *EncryptionPolicy) GetEncryptSloSubjectNameIdOk() (*bool, bool)`

GetEncryptSloSubjectNameIdOk returns a tuple with the EncryptSloSubjectNameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptSloSubjectNameId

`func (o *EncryptionPolicy) SetEncryptSloSubjectNameId(v bool)`

SetEncryptSloSubjectNameId sets EncryptSloSubjectNameId field to given value.

### HasEncryptSloSubjectNameId

`func (o *EncryptionPolicy) HasEncryptSloSubjectNameId() bool`

HasEncryptSloSubjectNameId returns a boolean if a field has been set.

### GetSloSubjectNameIDEncrypted

`func (o *EncryptionPolicy) GetSloSubjectNameIDEncrypted() bool`

GetSloSubjectNameIDEncrypted returns the SloSubjectNameIDEncrypted field if non-nil, zero value otherwise.

### GetSloSubjectNameIDEncryptedOk

`func (o *EncryptionPolicy) GetSloSubjectNameIDEncryptedOk() (*bool, bool)`

GetSloSubjectNameIDEncryptedOk returns a tuple with the SloSubjectNameIDEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloSubjectNameIDEncrypted

`func (o *EncryptionPolicy) SetSloSubjectNameIDEncrypted(v bool)`

SetSloSubjectNameIDEncrypted sets SloSubjectNameIDEncrypted field to given value.

### HasSloSubjectNameIDEncrypted

`func (o *EncryptionPolicy) HasSloSubjectNameIDEncrypted() bool`

HasSloSubjectNameIDEncrypted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


