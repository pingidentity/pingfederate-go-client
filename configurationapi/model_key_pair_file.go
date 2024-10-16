/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 12.1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the KeyPairFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyPairFile{}

// KeyPairFile Represents the contents of a PKCS12 or PEM file.
type KeyPairFile struct {
	// The persistent, unique ID for the certificate. It can be any combination of [a-z0-9._-]. This property is system-assigned if not specified.
	Id *string `json:"id,omitempty" tfsdk:"id"`
	// Base-64 encoded PKCS12 or PEM file data. In the case of PEM, the raw (non-base-64) data is also accepted. In BCFIPS mode, only PEM with PBES2 and AES or Triple DES encryption is accepted and 128-bit salt is required.
	FileData string `json:"fileData" tfsdk:"file_data"`
	// Key pair file format. If specified, this field will control what file format is expected, otherwise the format will be auto-detected. In BCFIPS mode, only PEM is supported.
	Format *string `json:"format,omitempty" tfsdk:"format"`
	// Password for the file. In BCFIPS mode, the password must be at least 14 characters.
	Password string `json:"password" tfsdk:"password"`
	// Encrypted password for the file. Only applicable for bulk export/import operations. For bulk import operation, either password or encrypted password must be set.
	EncryptedPassword *string `json:"encryptedPassword,omitempty" tfsdk:"encrypted_password"`
	// Cryptographic Provider. This is only applicable if Hybrid HSM mode is true.
	CryptoProvider *string `json:"cryptoProvider,omitempty" tfsdk:"crypto_provider"`
}

// NewKeyPairFile instantiates a new KeyPairFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyPairFile(fileData string, password string) *KeyPairFile {
	this := KeyPairFile{}
	this.FileData = fileData
	this.Password = password
	return &this
}

// NewKeyPairFileWithDefaults instantiates a new KeyPairFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyPairFileWithDefaults() *KeyPairFile {
	this := KeyPairFile{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KeyPairFile) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyPairFile) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KeyPairFile) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *KeyPairFile) SetId(v string) {
	o.Id = &v
}

// GetFileData returns the FileData field value
func (o *KeyPairFile) GetFileData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileData
}

// GetFileDataOk returns a tuple with the FileData field value
// and a boolean to check if the value has been set.
func (o *KeyPairFile) GetFileDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileData, true
}

// SetFileData sets field value
func (o *KeyPairFile) SetFileData(v string) {
	o.FileData = v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *KeyPairFile) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyPairFile) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *KeyPairFile) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *KeyPairFile) SetFormat(v string) {
	o.Format = &v
}

// GetPassword returns the Password field value
func (o *KeyPairFile) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *KeyPairFile) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *KeyPairFile) SetPassword(v string) {
	o.Password = v
}

// GetEncryptedPassword returns the EncryptedPassword field value if set, zero value otherwise.
func (o *KeyPairFile) GetEncryptedPassword() string {
	if o == nil || IsNil(o.EncryptedPassword) {
		var ret string
		return ret
	}
	return *o.EncryptedPassword
}

// GetEncryptedPasswordOk returns a tuple with the EncryptedPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyPairFile) GetEncryptedPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptedPassword) {
		return nil, false
	}
	return o.EncryptedPassword, true
}

// HasEncryptedPassword returns a boolean if a field has been set.
func (o *KeyPairFile) HasEncryptedPassword() bool {
	if o != nil && !IsNil(o.EncryptedPassword) {
		return true
	}

	return false
}

// SetEncryptedPassword gets a reference to the given string and assigns it to the EncryptedPassword field.
func (o *KeyPairFile) SetEncryptedPassword(v string) {
	o.EncryptedPassword = &v
}

// GetCryptoProvider returns the CryptoProvider field value if set, zero value otherwise.
func (o *KeyPairFile) GetCryptoProvider() string {
	if o == nil || IsNil(o.CryptoProvider) {
		var ret string
		return ret
	}
	return *o.CryptoProvider
}

// GetCryptoProviderOk returns a tuple with the CryptoProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyPairFile) GetCryptoProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CryptoProvider) {
		return nil, false
	}
	return o.CryptoProvider, true
}

// HasCryptoProvider returns a boolean if a field has been set.
func (o *KeyPairFile) HasCryptoProvider() bool {
	if o != nil && !IsNil(o.CryptoProvider) {
		return true
	}

	return false
}

// SetCryptoProvider gets a reference to the given string and assigns it to the CryptoProvider field.
func (o *KeyPairFile) SetCryptoProvider(v string) {
	o.CryptoProvider = &v
}

func (o KeyPairFile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyPairFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["fileData"] = o.FileData
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	toSerialize["password"] = o.Password
	if !IsNil(o.EncryptedPassword) {
		toSerialize["encryptedPassword"] = o.EncryptedPassword
	}
	if !IsNil(o.CryptoProvider) {
		toSerialize["cryptoProvider"] = o.CryptoProvider
	}
	return toSerialize, nil
}

type NullableKeyPairFile struct {
	value *KeyPairFile
	isSet bool
}

func (v NullableKeyPairFile) Get() *KeyPairFile {
	return v.value
}

func (v *NullableKeyPairFile) Set(val *KeyPairFile) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyPairFile) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyPairFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyPairFile(val *KeyPairFile) *NullableKeyPairFile {
	return &NullableKeyPairFile{value: val, isSet: true}
}

func (v NullableKeyPairFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyPairFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
