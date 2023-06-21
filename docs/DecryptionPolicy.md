# DecryptionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssertionEncrypted** | Pointer to **bool** | Specify whether the incoming SAML assertion is encrypted for an IdP connection. | [optional] 
**AttributesEncrypted** | Pointer to **bool** | Specify whether one or more incoming SAML attributes are encrypted for an IdP connection. | [optional] 
**SubjectNameIdEncrypted** | Pointer to **bool** | Specify whether the incoming Subject Name ID is encrypted for an IdP connection. | [optional] 
**SloEncryptSubjectNameID** | Pointer to **bool** | Encrypt the Subject Name ID in SLO messages to the IdP. | [optional] 
**SloSubjectNameIDEncrypted** | Pointer to **bool** | Allow encrypted Subject Name ID in SLO messages from the IdP. | [optional] 

## Methods

### NewDecryptionPolicy

`func NewDecryptionPolicy() *DecryptionPolicy`

NewDecryptionPolicy instantiates a new DecryptionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecryptionPolicyWithDefaults

`func NewDecryptionPolicyWithDefaults() *DecryptionPolicy`

NewDecryptionPolicyWithDefaults instantiates a new DecryptionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssertionEncrypted

`func (o *DecryptionPolicy) GetAssertionEncrypted() bool`

GetAssertionEncrypted returns the AssertionEncrypted field if non-nil, zero value otherwise.

### GetAssertionEncryptedOk

`func (o *DecryptionPolicy) GetAssertionEncryptedOk() (*bool, bool)`

GetAssertionEncryptedOk returns a tuple with the AssertionEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionEncrypted

`func (o *DecryptionPolicy) SetAssertionEncrypted(v bool)`

SetAssertionEncrypted sets AssertionEncrypted field to given value.

### HasAssertionEncrypted

`func (o *DecryptionPolicy) HasAssertionEncrypted() bool`

HasAssertionEncrypted returns a boolean if a field has been set.

### GetAttributesEncrypted

`func (o *DecryptionPolicy) GetAttributesEncrypted() bool`

GetAttributesEncrypted returns the AttributesEncrypted field if non-nil, zero value otherwise.

### GetAttributesEncryptedOk

`func (o *DecryptionPolicy) GetAttributesEncryptedOk() (*bool, bool)`

GetAttributesEncryptedOk returns a tuple with the AttributesEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributesEncrypted

`func (o *DecryptionPolicy) SetAttributesEncrypted(v bool)`

SetAttributesEncrypted sets AttributesEncrypted field to given value.

### HasAttributesEncrypted

`func (o *DecryptionPolicy) HasAttributesEncrypted() bool`

HasAttributesEncrypted returns a boolean if a field has been set.

### GetSubjectNameIdEncrypted

`func (o *DecryptionPolicy) GetSubjectNameIdEncrypted() bool`

GetSubjectNameIdEncrypted returns the SubjectNameIdEncrypted field if non-nil, zero value otherwise.

### GetSubjectNameIdEncryptedOk

`func (o *DecryptionPolicy) GetSubjectNameIdEncryptedOk() (*bool, bool)`

GetSubjectNameIdEncryptedOk returns a tuple with the SubjectNameIdEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectNameIdEncrypted

`func (o *DecryptionPolicy) SetSubjectNameIdEncrypted(v bool)`

SetSubjectNameIdEncrypted sets SubjectNameIdEncrypted field to given value.

### HasSubjectNameIdEncrypted

`func (o *DecryptionPolicy) HasSubjectNameIdEncrypted() bool`

HasSubjectNameIdEncrypted returns a boolean if a field has been set.

### GetSloEncryptSubjectNameID

`func (o *DecryptionPolicy) GetSloEncryptSubjectNameID() bool`

GetSloEncryptSubjectNameID returns the SloEncryptSubjectNameID field if non-nil, zero value otherwise.

### GetSloEncryptSubjectNameIDOk

`func (o *DecryptionPolicy) GetSloEncryptSubjectNameIDOk() (*bool, bool)`

GetSloEncryptSubjectNameIDOk returns a tuple with the SloEncryptSubjectNameID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloEncryptSubjectNameID

`func (o *DecryptionPolicy) SetSloEncryptSubjectNameID(v bool)`

SetSloEncryptSubjectNameID sets SloEncryptSubjectNameID field to given value.

### HasSloEncryptSubjectNameID

`func (o *DecryptionPolicy) HasSloEncryptSubjectNameID() bool`

HasSloEncryptSubjectNameID returns a boolean if a field has been set.

### GetSloSubjectNameIDEncrypted

`func (o *DecryptionPolicy) GetSloSubjectNameIDEncrypted() bool`

GetSloSubjectNameIDEncrypted returns the SloSubjectNameIDEncrypted field if non-nil, zero value otherwise.

### GetSloSubjectNameIDEncryptedOk

`func (o *DecryptionPolicy) GetSloSubjectNameIDEncryptedOk() (*bool, bool)`

GetSloSubjectNameIDEncryptedOk returns a tuple with the SloSubjectNameIDEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloSubjectNameIDEncrypted

`func (o *DecryptionPolicy) SetSloSubjectNameIDEncrypted(v bool)`

SetSloSubjectNameIDEncrypted sets SloSubjectNameIDEncrypted field to given value.

### HasSloSubjectNameIDEncrypted

`func (o *DecryptionPolicy) HasSloSubjectNameIDEncrypted() bool`

HasSloSubjectNameIDEncrypted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


