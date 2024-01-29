/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
	"fmt"
)

// checks if the AdministrativeAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdministrativeAccount{}

// AdministrativeAccount A PingFederate administrator account.
type AdministrativeAccount struct {
	// Username for the Administrative Account.
	Username string `json:"username" tfsdk:"username"`
	// Password for the Account. This field is only applicable during a POST operation.
	Password *string `json:"password,omitempty" tfsdk:"password"`
	// For GET requests, this field contains the encrypted account password. For POST and PUT requests, if you wish to re-use the password from an API response to this endpoint, this field should be passed back unchanged.
	EncryptedPassword *string `json:"encryptedPassword,omitempty" tfsdk:"encrypted_password"`
	// Indicates whether the account is active or not.
	Active *bool `json:"active,omitempty" tfsdk:"active"`
	// Description of the account.
	Description *string `json:"description,omitempty" tfsdk:"description"`
	// Indicates whether the account belongs to an Auditor. An Auditor has View-only permissions for all administrative functions. An Auditor cannot have any administrative roles.
	Auditor *bool `json:"auditor,omitempty" tfsdk:"auditor"`
	// Phone number associated with the account.
	PhoneNumber *string `json:"phoneNumber,omitempty" tfsdk:"phone_number"`
	// Email address associated with the account.
	EmailAddress *string `json:"emailAddress,omitempty" tfsdk:"email_address"`
	// The Department name of account user.
	Department *string `json:"department,omitempty" tfsdk:"department"`
	// Roles available for an administrator. <br>USER_ADMINISTRATOR - Can create, deactivate or delete accounts and reset passwords. Additionally, install replacement license keys. <br> CRYPTO_ADMINISTRATOR - Can manage local keys and certificates. <br> ADMINISTRATOR - Can configure partner connections and most system settings (except the management of native accounts and the handling of local keys and certificates. <br>EXPRESSION_ADMINISTRATOR - Can add and update OGNL expressions. <br>
	Roles []string `json:"roles,omitempty" tfsdk:"roles"`
}

type _AdministrativeAccount AdministrativeAccount

// NewAdministrativeAccount instantiates a new AdministrativeAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdministrativeAccount(username string) *AdministrativeAccount {
	this := AdministrativeAccount{}
	this.Username = username
	return &this
}

// NewAdministrativeAccountWithDefaults instantiates a new AdministrativeAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdministrativeAccountWithDefaults() *AdministrativeAccount {
	this := AdministrativeAccount{}
	return &this
}

// GetUsername returns the Username field value
func (o *AdministrativeAccount) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *AdministrativeAccount) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *AdministrativeAccount) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *AdministrativeAccount) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministrativeAccount) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *AdministrativeAccount) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *AdministrativeAccount) SetPassword(v string) {
	o.Password = &v
}

// GetEncryptedPassword returns the EncryptedPassword field value if set, zero value otherwise.
func (o *AdministrativeAccount) GetEncryptedPassword() string {
	if o == nil || IsNil(o.EncryptedPassword) {
		var ret string
		return ret
	}
	return *o.EncryptedPassword
}

// GetEncryptedPasswordOk returns a tuple with the EncryptedPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministrativeAccount) GetEncryptedPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptedPassword) {
		return nil, false
	}
	return o.EncryptedPassword, true
}

// HasEncryptedPassword returns a boolean if a field has been set.
func (o *AdministrativeAccount) HasEncryptedPassword() bool {
	if o != nil && !IsNil(o.EncryptedPassword) {
		return true
	}

	return false
}

// SetEncryptedPassword gets a reference to the given string and assigns it to the EncryptedPassword field.
func (o *AdministrativeAccount) SetEncryptedPassword(v string) {
	o.EncryptedPassword = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *AdministrativeAccount) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministrativeAccount) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *AdministrativeAccount) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *AdministrativeAccount) SetActive(v bool) {
	o.Active = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AdministrativeAccount) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministrativeAccount) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AdministrativeAccount) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AdministrativeAccount) SetDescription(v string) {
	o.Description = &v
}

// GetAuditor returns the Auditor field value if set, zero value otherwise.
func (o *AdministrativeAccount) GetAuditor() bool {
	if o == nil || IsNil(o.Auditor) {
		var ret bool
		return ret
	}
	return *o.Auditor
}

// GetAuditorOk returns a tuple with the Auditor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministrativeAccount) GetAuditorOk() (*bool, bool) {
	if o == nil || IsNil(o.Auditor) {
		return nil, false
	}
	return o.Auditor, true
}

// HasAuditor returns a boolean if a field has been set.
func (o *AdministrativeAccount) HasAuditor() bool {
	if o != nil && !IsNil(o.Auditor) {
		return true
	}

	return false
}

// SetAuditor gets a reference to the given bool and assigns it to the Auditor field.
func (o *AdministrativeAccount) SetAuditor(v bool) {
	o.Auditor = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *AdministrativeAccount) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministrativeAccount) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *AdministrativeAccount) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *AdministrativeAccount) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *AdministrativeAccount) GetEmailAddress() string {
	if o == nil || IsNil(o.EmailAddress) {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministrativeAccount) GetEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAddress) {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *AdministrativeAccount) HasEmailAddress() bool {
	if o != nil && !IsNil(o.EmailAddress) {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *AdministrativeAccount) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetDepartment returns the Department field value if set, zero value otherwise.
func (o *AdministrativeAccount) GetDepartment() string {
	if o == nil || IsNil(o.Department) {
		var ret string
		return ret
	}
	return *o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministrativeAccount) GetDepartmentOk() (*string, bool) {
	if o == nil || IsNil(o.Department) {
		return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *AdministrativeAccount) HasDepartment() bool {
	if o != nil && !IsNil(o.Department) {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given string and assigns it to the Department field.
func (o *AdministrativeAccount) SetDepartment(v string) {
	o.Department = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *AdministrativeAccount) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdministrativeAccount) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *AdministrativeAccount) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *AdministrativeAccount) SetRoles(v []string) {
	o.Roles = v
}

func (o AdministrativeAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdministrativeAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["username"] = o.Username
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.EncryptedPassword) {
		toSerialize["encryptedPassword"] = o.EncryptedPassword
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Auditor) {
		toSerialize["auditor"] = o.Auditor
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if !IsNil(o.EmailAddress) {
		toSerialize["emailAddress"] = o.EmailAddress
	}
	if !IsNil(o.Department) {
		toSerialize["department"] = o.Department
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

func (o *AdministrativeAccount) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"username",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAdministrativeAccount := _AdministrativeAccount{}

	err = json.Unmarshal(bytes, &varAdministrativeAccount)

	if err != nil {
		return err
	}

	*o = AdministrativeAccount(varAdministrativeAccount)

	return err
}

type NullableAdministrativeAccount struct {
	value *AdministrativeAccount
	isSet bool
}

func (v NullableAdministrativeAccount) Get() *AdministrativeAccount {
	return v.value
}

func (v *NullableAdministrativeAccount) Set(val *AdministrativeAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableAdministrativeAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAdministrativeAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdministrativeAccount(val *AdministrativeAccount) *NullableAdministrativeAccount {
	return &NullableAdministrativeAccount{value: val, isSet: true}
}

func (v NullableAdministrativeAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdministrativeAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
