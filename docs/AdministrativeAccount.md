# AdministrativeAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | Username for the Administrative Account. | 
**Password** | Pointer to **string** | Password for the Account. This field is only applicable during a POST operation. | [optional] 
**EncryptedPassword** | Pointer to **string** | For GET requests, this field contains the encrypted account password. For POST and PUT requests, if you wish to re-use the password from an API response to this endpoint, this field should be passed back unchanged. | [optional] 
**Active** | Pointer to **bool** | Indicates whether the account is active or not. | [optional] 
**Description** | Pointer to **string** | Description of the account. | [optional] 
**Auditor** | Pointer to **bool** | Indicates whether the account belongs to an Auditor. An Auditor has View-only permissions for all administrative functions. An Auditor cannot have any administrative roles. | [optional] 
**PhoneNumber** | Pointer to **string** | Phone number associated with the account. | [optional] 
**EmailAddress** | Pointer to **string** | Email address associated with the account. | [optional] 
**Department** | Pointer to **string** | The Department name of account user. | [optional] 
**Roles** | Pointer to **[]string** | Roles available for an administrator. &lt;br&gt;USER_ADMINISTRATOR - Can create, deactivate or delete accounts and reset passwords. Additionally, install replacement license keys. &lt;br&gt; CRYPTO_ADMINISTRATOR - Can manage local keys and certificates. &lt;br&gt; ADMINISTRATOR - Can configure partner connections and most system settings (except the management of native accounts and the handling of local keys and certificates. &lt;br&gt;EXPRESSION_ADMINISTRATOR - Can add and update OGNL expressions. &lt;br&gt; | [optional] 

## Methods

### NewAdministrativeAccount

`func NewAdministrativeAccount(username string, ) *AdministrativeAccount`

NewAdministrativeAccount instantiates a new AdministrativeAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdministrativeAccountWithDefaults

`func NewAdministrativeAccountWithDefaults() *AdministrativeAccount`

NewAdministrativeAccountWithDefaults instantiates a new AdministrativeAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *AdministrativeAccount) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AdministrativeAccount) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AdministrativeAccount) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *AdministrativeAccount) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AdministrativeAccount) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AdministrativeAccount) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AdministrativeAccount) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetEncryptedPassword

`func (o *AdministrativeAccount) GetEncryptedPassword() string`

GetEncryptedPassword returns the EncryptedPassword field if non-nil, zero value otherwise.

### GetEncryptedPasswordOk

`func (o *AdministrativeAccount) GetEncryptedPasswordOk() (*string, bool)`

GetEncryptedPasswordOk returns a tuple with the EncryptedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedPassword

`func (o *AdministrativeAccount) SetEncryptedPassword(v string)`

SetEncryptedPassword sets EncryptedPassword field to given value.

### HasEncryptedPassword

`func (o *AdministrativeAccount) HasEncryptedPassword() bool`

HasEncryptedPassword returns a boolean if a field has been set.

### GetActive

`func (o *AdministrativeAccount) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AdministrativeAccount) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AdministrativeAccount) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AdministrativeAccount) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDescription

`func (o *AdministrativeAccount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdministrativeAccount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdministrativeAccount) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdministrativeAccount) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAuditor

`func (o *AdministrativeAccount) GetAuditor() bool`

GetAuditor returns the Auditor field if non-nil, zero value otherwise.

### GetAuditorOk

`func (o *AdministrativeAccount) GetAuditorOk() (*bool, bool)`

GetAuditorOk returns a tuple with the Auditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditor

`func (o *AdministrativeAccount) SetAuditor(v bool)`

SetAuditor sets Auditor field to given value.

### HasAuditor

`func (o *AdministrativeAccount) HasAuditor() bool`

HasAuditor returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *AdministrativeAccount) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *AdministrativeAccount) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *AdministrativeAccount) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *AdministrativeAccount) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetEmailAddress

`func (o *AdministrativeAccount) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *AdministrativeAccount) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *AdministrativeAccount) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *AdministrativeAccount) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetDepartment

`func (o *AdministrativeAccount) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *AdministrativeAccount) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *AdministrativeAccount) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *AdministrativeAccount) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetRoles

`func (o *AdministrativeAccount) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AdministrativeAccount) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AdministrativeAccount) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *AdministrativeAccount) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


