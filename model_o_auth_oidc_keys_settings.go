/*
Administrative API Documentation

The PingFederate Administrative API is a REST-based interface that provides a programmatic way to make configuration changes to PingFederate as an alternative to using the administrative console.<br/><br/>Expand the resources below to display implementation details on that resource such as the available endpoints, the parameter and response models for the operation, and the model structure of the resources themselves. Each resource operation comes with the ability to interact with the API. You are prompted for proper administration credentials when you try to perform an API operation.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the OAuthOidcKeysSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuthOidcKeysSettings{}

// OAuthOidcKeysSettings Setting for OAuth/OpenID Connect signing and decryption key settings.
type OAuthOidcKeysSettings struct {
	// Enable static keys.
	StaticJwksEnabled   bool          `json:"staticJwksEnabled" tfsdk:"static_jwks_enabled"`
	P256ActiveCertRef   *ResourceLink `json:"p256ActiveCertRef,omitempty" tfsdk:"p256_active_cert_ref"`
	P256PreviousCertRef *ResourceLink `json:"p256PreviousCertRef,omitempty" tfsdk:"p256_previous_cert_ref"`
	// Enable publishing of the P-256 certificate chain associated with the active key.
	P256PublishX5cParameter *bool         `json:"p256PublishX5cParameter,omitempty" tfsdk:"p256_publish_x5c_parameter"`
	P384ActiveCertRef       *ResourceLink `json:"p384ActiveCertRef,omitempty" tfsdk:"p384_active_cert_ref"`
	P384PreviousCertRef     *ResourceLink `json:"p384PreviousCertRef,omitempty" tfsdk:"p384_previous_cert_ref"`
	// Enable publishing of the P-384 certificate chain associated with the active key.
	P384PublishX5cParameter *bool         `json:"p384PublishX5cParameter,omitempty" tfsdk:"p384_publish_x5c_parameter"`
	P521ActiveCertRef       *ResourceLink `json:"p521ActiveCertRef,omitempty" tfsdk:"p521_active_cert_ref"`
	P521PreviousCertRef     *ResourceLink `json:"p521PreviousCertRef,omitempty" tfsdk:"p521_previous_cert_ref"`
	// Enable publishing of the P-521 certificate chain associated with the active key.
	P521PublishX5cParameter *bool         `json:"p521PublishX5cParameter,omitempty" tfsdk:"p521_publish_x5c_parameter"`
	RsaActiveCertRef        *ResourceLink `json:"rsaActiveCertRef,omitempty" tfsdk:"rsa_active_cert_ref"`
	RsaPreviousCertRef      *ResourceLink `json:"rsaPreviousCertRef,omitempty" tfsdk:"rsa_previous_cert_ref"`
	// Enable publishing of the RSA certificate chain associated with the active key.
	RsaPublishX5cParameter        *bool         `json:"rsaPublishX5cParameter,omitempty" tfsdk:"rsa_publish_x5c_parameter"`
	P256DecryptionActiveCertRef   *ResourceLink `json:"p256DecryptionActiveCertRef,omitempty" tfsdk:"p256_decryption_active_cert_ref"`
	P256DecryptionPreviousCertRef *ResourceLink `json:"p256DecryptionPreviousCertRef,omitempty" tfsdk:"p256_decryption_previous_cert_ref"`
	// Enable publishing of the P-256 certificate chain associated with the active key.
	P256DecryptionPublishX5cParameter *bool         `json:"p256DecryptionPublishX5cParameter,omitempty" tfsdk:"p256_decryption_publish_x5c_parameter"`
	P384DecryptionActiveCertRef       *ResourceLink `json:"p384DecryptionActiveCertRef,omitempty" tfsdk:"p384_decryption_active_cert_ref"`
	P384DecryptionPreviousCertRef     *ResourceLink `json:"p384DecryptionPreviousCertRef,omitempty" tfsdk:"p384_decryption_previous_cert_ref"`
	// Enable publishing of the P-384 certificate chain associated with the active key.
	P384DecryptionPublishX5cParameter *bool         `json:"p384DecryptionPublishX5cParameter,omitempty" tfsdk:"p384_decryption_publish_x5c_parameter"`
	P521DecryptionActiveCertRef       *ResourceLink `json:"p521DecryptionActiveCertRef,omitempty" tfsdk:"p521_decryption_active_cert_ref"`
	P521DecryptionPreviousCertRef     *ResourceLink `json:"p521DecryptionPreviousCertRef,omitempty" tfsdk:"p521_decryption_previous_cert_ref"`
	// Enable publishing of the P-521 certificate chain associated with the active key.
	P521DecryptionPublishX5cParameter *bool         `json:"p521DecryptionPublishX5cParameter,omitempty" tfsdk:"p521_decryption_publish_x5c_parameter"`
	RsaDecryptionActiveCertRef        *ResourceLink `json:"rsaDecryptionActiveCertRef,omitempty" tfsdk:"rsa_decryption_active_cert_ref"`
	RsaDecryptionPreviousCertRef      *ResourceLink `json:"rsaDecryptionPreviousCertRef,omitempty" tfsdk:"rsa_decryption_previous_cert_ref"`
	// Enable publishing of the RSA certificate chain associated with the active key.
	RsaDecryptionPublishX5cParameter *bool `json:"rsaDecryptionPublishX5cParameter,omitempty" tfsdk:"rsa_decryption_publish_x5c_parameter"`
}

// NewOAuthOidcKeysSettings instantiates a new OAuthOidcKeysSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthOidcKeysSettings(staticJwksEnabled bool) *OAuthOidcKeysSettings {
	this := OAuthOidcKeysSettings{}
	this.StaticJwksEnabled = staticJwksEnabled
	return &this
}

// NewOAuthOidcKeysSettingsWithDefaults instantiates a new OAuthOidcKeysSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthOidcKeysSettingsWithDefaults() *OAuthOidcKeysSettings {
	this := OAuthOidcKeysSettings{}
	return &this
}

// GetStaticJwksEnabled returns the StaticJwksEnabled field value
func (o *OAuthOidcKeysSettings) GetStaticJwksEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.StaticJwksEnabled
}

// GetStaticJwksEnabledOk returns a tuple with the StaticJwksEnabled field value
// and a boolean to check if the value has been set.
func (o *OAuthOidcKeysSettings) GetStaticJwksEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StaticJwksEnabled, true
}

// SetStaticJwksEnabled sets field value
func (o *OAuthOidcKeysSettings) SetStaticJwksEnabled(v bool) {
	o.StaticJwksEnabled = v
}

// GetP256ActiveCertRef returns the P256ActiveCertRef field value if set, zero value otherwise.
func (o *OAuthOidcKeysSettings) GetP256ActiveCertRef() ResourceLink {
	if o == nil || IsNil(o.P256ActiveCertRef) {
		var ret ResourceLink
		return ret
	}
	return *o.P256ActiveCertRef
}

// GetP256ActiveCertRefOk returns a tuple with the P256ActiveCertRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthOidcKeysSettings) GetP256ActiveCertRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.P256ActiveCertRef) {
		return nil, false
	}
	return o.P256ActiveCertRef, true
}

// HasP256ActiveCertRef returns a boolean if a field has been set.
func (o *OAuthOidcKeysSettings) HasP256ActiveCertRef() bool {
	if o != nil && !IsNil(o.P256ActiveCertRef) {
		return true
	}

	return false
}

// SetP256ActiveCertRef gets a reference to the given ResourceLink and assigns it to the P256ActiveCertRef field.
func (o *OAuthOidcKeysSettings) SetP256ActiveCertRef(v ResourceLink) {
	o.P256ActiveCertRef = &v
}

// GetP256PreviousCertRef returns the P256PreviousCertRef field value if set, zero value otherwise.
func (o *OAuthOidcKeysSettings) GetP256PreviousCertRef() ResourceLink {
	if o == nil || IsNil(o.P256PreviousCertRef) {
		var ret ResourceLink
		return ret
	}
	return *o.P256PreviousCertRef
}

// GetP256PreviousCertRefOk returns a tuple with the P256PreviousCertRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthOidcKeysSettings) GetP256PreviousCertRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.P256PreviousCertRef) {
		return nil, false
	}
	return o.P256PreviousCertRef, true
}

// HasP256PreviousCertRef returns a boolean if a field has been set.
func (o *OAuthOidcKeysSettings) HasP256PreviousCertRef() bool {
	if o != nil && !IsNil(o.P256PreviousCertRef) {
		return true
	}

	return false
}

// SetP256PreviousCertRef gets a reference to the given ResourceLink and assigns it to the P256PreviousCertRef field.
func (o *OAuthOidcKeysSettings) SetP256PreviousCertRef(v ResourceLink) {
	o.P256PreviousCertRef = &v
}

// GetP256PublishX5cParameter returns the P256PublishX5cParameter field value if set, zero value otherwise.
func (o *OAuthOidcKeysSettings) GetP256PublishX5cParameter() bool {
	if o == nil || IsNil(o.P256PublishX5cParameter) {
		var ret bool
		return ret
	}
	return *o.P256PublishX5cParameter
}

// GetP256PublishX5cParameterOk returns a tuple with the P256PublishX5cParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthOidcKeysSettings) GetP256PublishX5cParameterOk() (*bool, bool) {
	if o == nil || IsNil(o.P256PublishX5cParameter) {
		return nil, false
	}
	return o.P256PublishX5cParameter, true
}

// HasP256PublishX5cParameter returns a boolean if a field has been set.
func (o *OAuthOidcKeysSettings) HasP256PublishX5cParameter() bool {
	if o != nil && !IsNil(o.P256PublishX5cParameter) {
		return true
	}

	return false
}

// SetP256PublishX5cParameter gets a reference to the given bool and assigns it to the P256PublishX5cParameter field.
func (o *OAuthOidcKeysSettings) SetP256PublishX5cParameter(v bool) {
	o.P256PublishX5cParameter = &v
}

// GetP384ActiveCertRef returns the P384ActiveCertRef field value if set, zero value otherwise.
func (o *OAuthOidcKeysSettings) GetP384ActiveCertRef() ResourceLink {
	if o == nil || IsNil(o.P384ActiveCertRef) {
		var ret ResourceLink
		return ret
	}
	return *o.P384ActiveCertRef
}

// GetP384ActiveCertRefOk returns a tuple with the P384ActiveCertRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthOidcKeysSettings) GetP384ActiveCertRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.P384ActiveCertRef) {
		return nil, false
	}
	return o.P384ActiveCertRef, true
}

// HasP384ActiveCertRef returns a boolean if a field has been set.
func (o *OAuthOidcKeysSettings) HasP384ActiveCertRef() bool {
	if o != nil && !IsNil(o.P384ActiveCertRef) {
		return true
	}

	return false
}

// SetP384ActiveCertRef gets a reference to the given ResourceLink and assigns it to the P384ActiveCertRef field.
func (o *OAuthOidcKeysSettings) SetP384ActiveCertRef(v ResourceLink) {
	o.P384ActiveCertRef = &v
}

// GetP384PreviousCertRef returns the P384PreviousCertRef field value if set, zero value otherwise.
func (o *OAuthOidcKeysSettings) GetP384PreviousCertRef() ResourceLink {
	if o == nil || IsNil(o.P384PreviousCertRef) {
		var ret ResourceLink
		return ret
	}
	return *o.P384PreviousCertRef
}

// GetP384PreviousCertRefOk returns a tuple with the P384PreviousCertRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthOidcKeysSettings) GetP384PreviousCertRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.P384PreviousCertRef) {
		return nil, false
	}
	return o.P384PreviousCertRef, true
}

// HasP384PreviousCertRef returns a boolean if a field has been set.
func (o *OAuthOidcKeysSettings) HasP384PreviousCertRef() bool {
	if o != nil && !IsNil(o.P384PreviousCertRef) {
		return true
	}

	return false
}

// SetP384PreviousCertRef gets a reference to the given ResourceLink and assigns it to the P384PreviousCertRef field.
func (o *OAuthOidcKeysSettings) SetP384PreviousCertRef(v ResourceLink) {
	o.P384PreviousCertRef = &v
}

// GetP384PublishX5cParameter returns the P384PublishX5cParameter field value if set, zero value otherwise.
func (o *OAuthOidcKeysSettings) GetP384PublishX5cParameter() bool {
	if o == nil || IsNil(o.P384PublishX5cParameter) {
		var ret bool
		return ret
	}
	return *o.P384PublishX5cParameter
}

// GetP384PublishX5cParameterOk returns a tuple with the P384PublishX5cParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthOidcKeysSettings) GetP384PublishX5cParameterOk() (*bool, bool) {
	if o == nil || IsNil(o.P384PublishX5cParameter) {
		return nil, false
	}
	return o.P384PublishX5cParameter, true
}

// HasP384PublishX5cParameter returns a boolean if a field has been set.
func (o *OAuthOidcKeysSettings) HasP384PublishX5cParameter() bool {
	if o != nil && !IsNil(o.P384PublishX5cParameter) {
		return true
	}

	return false
}

// SetP384PublishX5cParameter gets a reference to the given bool and assigns it to the P384PublishX5cParameter field.
func (o *OAuthOidcKeysSettings) SetP384PublishX5cParameter(v bool) {
	o.P384PublishX5cParameter = &v
}

// GetP521ActiveCertRef returns the P521ActiveCertRef field value if set, zero value otherwise.
func (o *OAuthOidcKeysSettings) GetP521ActiveCertRef() ResourceLink {
	if o == nil || IsNil(o.P521ActiveCertRef) {
		var ret ResourceLink
		return ret
	}
	return *o.P521ActiveCertRef
}

// GetP521ActiveCertRefOk returns a tuple with the P521ActiveCertRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthOidcKeysSettings) GetP521ActiveCertRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.P521ActiveCertRef) {
		return nil, false
	}
	return o.P521ActiveCertRef, true
}

// HasP521ActiveCertRef returns a boolean if a field has been set.
func (o *OAuthOidcKeysSettings) HasP521ActiveCertRef() bool {
	if o != nil && !IsNil(o.P521ActiveCertRef) {
		return true
	}

	return false
}

// SetP521ActiveCertRef gets a reference to the given ResourceLink and assigns it to the P521ActiveCertRef field.
func (o *OAuthOidcKeysSettings) SetP521ActiveCertRef(v ResourceLink) {
	o.P521ActiveCertRef = &v
}

// GetP521PreviousCertRef returns the P521PreviousCertRef field value if set, zero value otherwise.
func (o *OAuthOidcKeysSettings) GetP521PreviousCertRef() ResourceLink {
	if o == nil || IsNil(o.P521PreviousCertRef) {
		var ret ResourceLink
		return ret
	}
	return *o.P521PreviousCertRef
}

// GetP521PreviousCertRefOk returns a tuple with the P521PreviousCertRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthOidcKeysSettings) GetP521PreviousCertRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.P521PreviousCertRef) {
		return nil, false
	}
	return o.P521PreviousCertRef, true
}

// HasP521PreviousCertRef returns a boolean if a field has been set.
func (o *OAuthOidcKeysSettings) HasP521PreviousCertRef() bool {
	if o != nil && !IsNil(o.P521PreviousCertRef) {
		return true
	}

	return false
}

// SetP521PreviousCertRef gets a reference to the given ResourceLink and assigns it to the P521PreviousCertRef field.
func (o *OAuthOidcKeysSettings) SetP521PreviousCertRef(v ResourceLink) {
	o.P521PreviousCertRef = &v
}

// GetP521PublishX5cParameter returns the P521PublishX5cParameter field value if set, zero value otherwise.
func (o *OAuthOidcKeysSettings) GetP521PublishX5cParameter() bool {
	if o == nil || IsNil(o.P521PublishX5cParameter) {
		var ret bool
		return ret
	}
	return *o.P521PublishX5cParameter
}

// GetP521PublishX5cParameterOk returns a tuple with the P521PublishX5cParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthOidcKeysSettings) GetP521PublishX5cParameterOk() (*bool, bool) {
	if o == nil || IsNil(o.P521PublishX5cParameter) {
		return nil, false
	}
	return o.P521PublishX5cParameter, true
}

// HasP521PublishX5cParameter returns a boolean if a field has been set.
func (o *OAuthOidcKeysSettings) HasP521PublishX5cParameter() bool {
	if o != nil && !IsNil(o.P521PublishX5cParameter) {
		return true
	}

	return false
}

// SetP521PublishX5cParameter gets a reference to the given bool and assigns it to the P521PublishX5cParameter field.
func (o *OAuthOidcKeysSettings) SetP521PublishX5cParameter(v bool) {
	o.P521PublishX5cParameter = &v
}

// GetRsaActiveCertRef returns the RsaActiveCertRef field value if set, zero value otherwise.
func (o *OAuthOidcKeysSettings) GetRsaActiveCertRef() ResourceLink {
	if o == nil || IsNil(o.RsaActiveCertRef) {
		var ret ResourceLink
		return ret
	}
	return *o.RsaActiveCertRef
}

// GetRsaActiveCertRefOk returns a tuple with the RsaActiveCertRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthOidcKeysSettings) GetRsaActiveCertRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.RsaActiveCertRef) {
		return nil, false
	}
	return o.RsaActiveCertRef, true
}

// HasRsaActiveCertRef returns a boolean if a field has been set.
func (o *OAuthOidcKeysSettings) HasRsaActiveCertRef() bool {
	if o != nil && !IsNil(o.RsaActiveCertRef) {
		return true
	}

	return false
}

// SetRsaActiveCertRef gets a reference to the given ResourceLink and assigns it to the RsaActiveCertRef field.
func (o *OAuthOidcKeysSettings) SetRsaActiveCertRef(v ResourceLink) {
	o.RsaActiveCertRef = &v
}

// GetRsaPreviousCertRef returns the RsaPreviousCertRef field value if set, zero value otherwise.
func (o *OAuthOidcKeysSettings) GetRsaPreviousCertRef() ResourceLink {
	if o == nil || IsNil(o.RsaPreviousCertRef) {
		var ret ResourceLink
		return ret
	}
	return *o.RsaPreviousCertRef
}

// GetRsaPreviousCertRefOk returns a tuple with the RsaPreviousCertRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthOidcKeysSettings) GetRsaPreviousCertRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.RsaPreviousCertRef) {
		return nil, false
	}
	return o.RsaPreviousCertRef, true
}

// HasRsaPreviousCertRef returns a boolean if a field has been set.
func (o *OAuthOidcKeysSettings) HasRsaPreviousCertRef() bool {
	if o != nil && !IsNil(o.RsaPreviousCertRef) {
		return true
	}

	return false
}

// SetRsaPreviousCertRef gets a reference to the given ResourceLink and assigns it to the RsaPreviousCertRef field.
func (o *OAuthOidcKeysSettings) SetRsaPreviousCertRef(v ResourceLink) {
	o.RsaPreviousCertRef = &v
}

// GetRsaPublishX5cParameter returns the RsaPublishX5cParameter field value if set, zero value otherwise.
func (o *OAuthOidcKeysSettings) GetRsaPublishX5cParameter() bool {
	if o == nil || IsNil(o.RsaPublishX5cParameter) {
		var ret bool
		return ret
	}
	return *o.RsaPublishX5cParameter
}

// GetRsaPublishX5cParameterOk returns a tuple with the RsaPublishX5cParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthOidcKeysSettings) GetRsaPublishX5cParameterOk() (*bool, bool) {
	if o == nil || IsNil(o.RsaPublishX5cParameter) {
		return nil, false
	}
	return o.RsaPublishX5cParameter, true
}

// HasRsaPublishX5cParameter returns a boolean if a field has been set.
func (o *OAuthOidcKeysSettings) HasRsaPublishX5cParameter() bool {
	if o != nil && !IsNil(o.RsaPublishX5cParameter) {
		return true
	}

	return false
}

// SetRsaPublishX5cParameter gets a reference to the given bool and assigns it to the RsaPublishX5cParameter field.
func (o *OAuthOidcKeysSettings) SetRsaPublishX5cParameter(v bool) {
	o.RsaPublishX5cParameter = &v
}

// GetP256DecryptionActiveCertRef returns the P256DecryptionActiveCertRef field value if set, zero value otherwise.
func (o *OAuthOidcKeysSettings) GetP256DecryptionActiveCertRef() ResourceLink {
	if o == nil || IsNil(o.P256DecryptionActiveCertRef) {
		var ret ResourceLink
		return ret
	}
	return *o.P256DecryptionActiveCertRef
}

// GetP256DecryptionActiveCertRefOk returns a tuple with the P256DecryptionActiveCertRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthOidcKeysSettings) GetP256DecryptionActiveCertRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.P256DecryptionActiveCertRef) {
		return nil, false
	}
	return o.P256DecryptionActiveCertRef, true
}

// HasP256DecryptionActiveCertRef returns a boolean if a field has been set.
func (o *OAuthOidcKeysSettings) HasP256DecryptionActiveCertRef() bool {
	if o != nil && !IsNil(o.P256DecryptionActiveCertRef) {
		return true
	}

	return false
}

// SetP256DecryptionActiveCertRef gets a reference to the given ResourceLink and assigns it to the P256DecryptionActiveCertRef field.
func (o *OAuthOidcKeysSettings) SetP256DecryptionActiveCertRef(v ResourceLink) {
	o.P256DecryptionActiveCertRef = &v
}

// GetP256DecryptionPreviousCertRef returns the P256DecryptionPreviousCertRef field value if set, zero value otherwise.
func (o *OAuthOidcKeysSettings) GetP256DecryptionPreviousCertRef() ResourceLink {
	if o == nil || IsNil(o.P256DecryptionPreviousCertRef) {
		var ret ResourceLink
		return ret
	}
	return *o.P256DecryptionPreviousCertRef
}

// GetP256DecryptionPreviousCertRefOk returns a tuple with the P256DecryptionPreviousCertRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthOidcKeysSettings) GetP256DecryptionPreviousCertRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.P256DecryptionPreviousCertRef) {
		return nil, false
	}
	return o.P256DecryptionPreviousCertRef, true
}

// HasP256DecryptionPreviousCertRef returns a boolean if a field has been set.
func (o *OAuthOidcKeysSettings) HasP256DecryptionPreviousCertRef() bool {
	if o != nil && !IsNil(o.P256DecryptionPreviousCertRef) {
		return true
	}

	return false
}

// SetP256DecryptionPreviousCertRef gets a reference to the given ResourceLink and assigns it to the P256DecryptionPreviousCertRef field.
func (o *OAuthOidcKeysSettings) SetP256DecryptionPreviousCertRef(v ResourceLink) {
	o.P256DecryptionPreviousCertRef = &v
}

// GetP256DecryptionPublishX5cParameter returns the P256DecryptionPublishX5cParameter field value if set, zero value otherwise.
func (o *OAuthOidcKeysSettings) GetP256DecryptionPublishX5cParameter() bool {
	if o == nil || IsNil(o.P256DecryptionPublishX5cParameter) {
		var ret bool
		return ret
	}
	return *o.P256DecryptionPublishX5cParameter
}

// GetP256DecryptionPublishX5cParameterOk returns a tuple with the P256DecryptionPublishX5cParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthOidcKeysSettings) GetP256DecryptionPublishX5cParameterOk() (*bool, bool) {
	if o == nil || IsNil(o.P256DecryptionPublishX5cParameter) {
		return nil, false
	}
	return o.P256DecryptionPublishX5cParameter, true
}

// HasP256DecryptionPublishX5cParameter returns a boolean if a field has been set.
func (o *OAuthOidcKeysSettings) HasP256DecryptionPublishX5cParameter() bool {
	if o != nil && !IsNil(o.P256DecryptionPublishX5cParameter) {
		return true
	}

	return false
}

// SetP256DecryptionPublishX5cParameter gets a reference to the given bool and assigns it to the P256DecryptionPublishX5cParameter field.
func (o *OAuthOidcKeysSettings) SetP256DecryptionPublishX5cParameter(v bool) {
	o.P256DecryptionPublishX5cParameter = &v
}

// GetP384DecryptionActiveCertRef returns the P384DecryptionActiveCertRef field value if set, zero value otherwise.
func (o *OAuthOidcKeysSettings) GetP384DecryptionActiveCertRef() ResourceLink {
	if o == nil || IsNil(o.P384DecryptionActiveCertRef) {
		var ret ResourceLink
		return ret
	}
	return *o.P384DecryptionActiveCertRef
}

// GetP384DecryptionActiveCertRefOk returns a tuple with the P384DecryptionActiveCertRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthOidcKeysSettings) GetP384DecryptionActiveCertRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.P384DecryptionActiveCertRef) {
		return nil, false
	}
	return o.P384DecryptionActiveCertRef, true
}

// HasP384DecryptionActiveCertRef returns a boolean if a field has been set.
func (o *OAuthOidcKeysSettings) HasP384DecryptionActiveCertRef() bool {
	if o != nil && !IsNil(o.P384DecryptionActiveCertRef) {
		return true
	}

	return false
}

// SetP384DecryptionActiveCertRef gets a reference to the given ResourceLink and assigns it to the P384DecryptionActiveCertRef field.
func (o *OAuthOidcKeysSettings) SetP384DecryptionActiveCertRef(v ResourceLink) {
	o.P384DecryptionActiveCertRef = &v
}

// GetP384DecryptionPreviousCertRef returns the P384DecryptionPreviousCertRef field value if set, zero value otherwise.
func (o *OAuthOidcKeysSettings) GetP384DecryptionPreviousCertRef() ResourceLink {
	if o == nil || IsNil(o.P384DecryptionPreviousCertRef) {
		var ret ResourceLink
		return ret
	}
	return *o.P384DecryptionPreviousCertRef
}

// GetP384DecryptionPreviousCertRefOk returns a tuple with the P384DecryptionPreviousCertRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthOidcKeysSettings) GetP384DecryptionPreviousCertRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.P384DecryptionPreviousCertRef) {
		return nil, false
	}
	return o.P384DecryptionPreviousCertRef, true
}

// HasP384DecryptionPreviousCertRef returns a boolean if a field has been set.
func (o *OAuthOidcKeysSettings) HasP384DecryptionPreviousCertRef() bool {
	if o != nil && !IsNil(o.P384DecryptionPreviousCertRef) {
		return true
	}

	return false
}

// SetP384DecryptionPreviousCertRef gets a reference to the given ResourceLink and assigns it to the P384DecryptionPreviousCertRef field.
func (o *OAuthOidcKeysSettings) SetP384DecryptionPreviousCertRef(v ResourceLink) {
	o.P384DecryptionPreviousCertRef = &v
}

// GetP384DecryptionPublishX5cParameter returns the P384DecryptionPublishX5cParameter field value if set, zero value otherwise.
func (o *OAuthOidcKeysSettings) GetP384DecryptionPublishX5cParameter() bool {
	if o == nil || IsNil(o.P384DecryptionPublishX5cParameter) {
		var ret bool
		return ret
	}
	return *o.P384DecryptionPublishX5cParameter
}

// GetP384DecryptionPublishX5cParameterOk returns a tuple with the P384DecryptionPublishX5cParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthOidcKeysSettings) GetP384DecryptionPublishX5cParameterOk() (*bool, bool) {
	if o == nil || IsNil(o.P384DecryptionPublishX5cParameter) {
		return nil, false
	}
	return o.P384DecryptionPublishX5cParameter, true
}

// HasP384DecryptionPublishX5cParameter returns a boolean if a field has been set.
func (o *OAuthOidcKeysSettings) HasP384DecryptionPublishX5cParameter() bool {
	if o != nil && !IsNil(o.P384DecryptionPublishX5cParameter) {
		return true
	}

	return false
}

// SetP384DecryptionPublishX5cParameter gets a reference to the given bool and assigns it to the P384DecryptionPublishX5cParameter field.
func (o *OAuthOidcKeysSettings) SetP384DecryptionPublishX5cParameter(v bool) {
	o.P384DecryptionPublishX5cParameter = &v
}

// GetP521DecryptionActiveCertRef returns the P521DecryptionActiveCertRef field value if set, zero value otherwise.
func (o *OAuthOidcKeysSettings) GetP521DecryptionActiveCertRef() ResourceLink {
	if o == nil || IsNil(o.P521DecryptionActiveCertRef) {
		var ret ResourceLink
		return ret
	}
	return *o.P521DecryptionActiveCertRef
}

// GetP521DecryptionActiveCertRefOk returns a tuple with the P521DecryptionActiveCertRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthOidcKeysSettings) GetP521DecryptionActiveCertRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.P521DecryptionActiveCertRef) {
		return nil, false
	}
	return o.P521DecryptionActiveCertRef, true
}

// HasP521DecryptionActiveCertRef returns a boolean if a field has been set.
func (o *OAuthOidcKeysSettings) HasP521DecryptionActiveCertRef() bool {
	if o != nil && !IsNil(o.P521DecryptionActiveCertRef) {
		return true
	}

	return false
}

// SetP521DecryptionActiveCertRef gets a reference to the given ResourceLink and assigns it to the P521DecryptionActiveCertRef field.
func (o *OAuthOidcKeysSettings) SetP521DecryptionActiveCertRef(v ResourceLink) {
	o.P521DecryptionActiveCertRef = &v
}

// GetP521DecryptionPreviousCertRef returns the P521DecryptionPreviousCertRef field value if set, zero value otherwise.
func (o *OAuthOidcKeysSettings) GetP521DecryptionPreviousCertRef() ResourceLink {
	if o == nil || IsNil(o.P521DecryptionPreviousCertRef) {
		var ret ResourceLink
		return ret
	}
	return *o.P521DecryptionPreviousCertRef
}

// GetP521DecryptionPreviousCertRefOk returns a tuple with the P521DecryptionPreviousCertRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthOidcKeysSettings) GetP521DecryptionPreviousCertRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.P521DecryptionPreviousCertRef) {
		return nil, false
	}
	return o.P521DecryptionPreviousCertRef, true
}

// HasP521DecryptionPreviousCertRef returns a boolean if a field has been set.
func (o *OAuthOidcKeysSettings) HasP521DecryptionPreviousCertRef() bool {
	if o != nil && !IsNil(o.P521DecryptionPreviousCertRef) {
		return true
	}

	return false
}

// SetP521DecryptionPreviousCertRef gets a reference to the given ResourceLink and assigns it to the P521DecryptionPreviousCertRef field.
func (o *OAuthOidcKeysSettings) SetP521DecryptionPreviousCertRef(v ResourceLink) {
	o.P521DecryptionPreviousCertRef = &v
}

// GetP521DecryptionPublishX5cParameter returns the P521DecryptionPublishX5cParameter field value if set, zero value otherwise.
func (o *OAuthOidcKeysSettings) GetP521DecryptionPublishX5cParameter() bool {
	if o == nil || IsNil(o.P521DecryptionPublishX5cParameter) {
		var ret bool
		return ret
	}
	return *o.P521DecryptionPublishX5cParameter
}

// GetP521DecryptionPublishX5cParameterOk returns a tuple with the P521DecryptionPublishX5cParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthOidcKeysSettings) GetP521DecryptionPublishX5cParameterOk() (*bool, bool) {
	if o == nil || IsNil(o.P521DecryptionPublishX5cParameter) {
		return nil, false
	}
	return o.P521DecryptionPublishX5cParameter, true
}

// HasP521DecryptionPublishX5cParameter returns a boolean if a field has been set.
func (o *OAuthOidcKeysSettings) HasP521DecryptionPublishX5cParameter() bool {
	if o != nil && !IsNil(o.P521DecryptionPublishX5cParameter) {
		return true
	}

	return false
}

// SetP521DecryptionPublishX5cParameter gets a reference to the given bool and assigns it to the P521DecryptionPublishX5cParameter field.
func (o *OAuthOidcKeysSettings) SetP521DecryptionPublishX5cParameter(v bool) {
	o.P521DecryptionPublishX5cParameter = &v
}

// GetRsaDecryptionActiveCertRef returns the RsaDecryptionActiveCertRef field value if set, zero value otherwise.
func (o *OAuthOidcKeysSettings) GetRsaDecryptionActiveCertRef() ResourceLink {
	if o == nil || IsNil(o.RsaDecryptionActiveCertRef) {
		var ret ResourceLink
		return ret
	}
	return *o.RsaDecryptionActiveCertRef
}

// GetRsaDecryptionActiveCertRefOk returns a tuple with the RsaDecryptionActiveCertRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthOidcKeysSettings) GetRsaDecryptionActiveCertRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.RsaDecryptionActiveCertRef) {
		return nil, false
	}
	return o.RsaDecryptionActiveCertRef, true
}

// HasRsaDecryptionActiveCertRef returns a boolean if a field has been set.
func (o *OAuthOidcKeysSettings) HasRsaDecryptionActiveCertRef() bool {
	if o != nil && !IsNil(o.RsaDecryptionActiveCertRef) {
		return true
	}

	return false
}

// SetRsaDecryptionActiveCertRef gets a reference to the given ResourceLink and assigns it to the RsaDecryptionActiveCertRef field.
func (o *OAuthOidcKeysSettings) SetRsaDecryptionActiveCertRef(v ResourceLink) {
	o.RsaDecryptionActiveCertRef = &v
}

// GetRsaDecryptionPreviousCertRef returns the RsaDecryptionPreviousCertRef field value if set, zero value otherwise.
func (o *OAuthOidcKeysSettings) GetRsaDecryptionPreviousCertRef() ResourceLink {
	if o == nil || IsNil(o.RsaDecryptionPreviousCertRef) {
		var ret ResourceLink
		return ret
	}
	return *o.RsaDecryptionPreviousCertRef
}

// GetRsaDecryptionPreviousCertRefOk returns a tuple with the RsaDecryptionPreviousCertRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthOidcKeysSettings) GetRsaDecryptionPreviousCertRefOk() (*ResourceLink, bool) {
	if o == nil || IsNil(o.RsaDecryptionPreviousCertRef) {
		return nil, false
	}
	return o.RsaDecryptionPreviousCertRef, true
}

// HasRsaDecryptionPreviousCertRef returns a boolean if a field has been set.
func (o *OAuthOidcKeysSettings) HasRsaDecryptionPreviousCertRef() bool {
	if o != nil && !IsNil(o.RsaDecryptionPreviousCertRef) {
		return true
	}

	return false
}

// SetRsaDecryptionPreviousCertRef gets a reference to the given ResourceLink and assigns it to the RsaDecryptionPreviousCertRef field.
func (o *OAuthOidcKeysSettings) SetRsaDecryptionPreviousCertRef(v ResourceLink) {
	o.RsaDecryptionPreviousCertRef = &v
}

// GetRsaDecryptionPublishX5cParameter returns the RsaDecryptionPublishX5cParameter field value if set, zero value otherwise.
func (o *OAuthOidcKeysSettings) GetRsaDecryptionPublishX5cParameter() bool {
	if o == nil || IsNil(o.RsaDecryptionPublishX5cParameter) {
		var ret bool
		return ret
	}
	return *o.RsaDecryptionPublishX5cParameter
}

// GetRsaDecryptionPublishX5cParameterOk returns a tuple with the RsaDecryptionPublishX5cParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthOidcKeysSettings) GetRsaDecryptionPublishX5cParameterOk() (*bool, bool) {
	if o == nil || IsNil(o.RsaDecryptionPublishX5cParameter) {
		return nil, false
	}
	return o.RsaDecryptionPublishX5cParameter, true
}

// HasRsaDecryptionPublishX5cParameter returns a boolean if a field has been set.
func (o *OAuthOidcKeysSettings) HasRsaDecryptionPublishX5cParameter() bool {
	if o != nil && !IsNil(o.RsaDecryptionPublishX5cParameter) {
		return true
	}

	return false
}

// SetRsaDecryptionPublishX5cParameter gets a reference to the given bool and assigns it to the RsaDecryptionPublishX5cParameter field.
func (o *OAuthOidcKeysSettings) SetRsaDecryptionPublishX5cParameter(v bool) {
	o.RsaDecryptionPublishX5cParameter = &v
}

func (o OAuthOidcKeysSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuthOidcKeysSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["staticJwksEnabled"] = o.StaticJwksEnabled
	if !IsNil(o.P256ActiveCertRef) {
		toSerialize["p256ActiveCertRef"] = o.P256ActiveCertRef
	}
	if !IsNil(o.P256PreviousCertRef) {
		toSerialize["p256PreviousCertRef"] = o.P256PreviousCertRef
	}
	if !IsNil(o.P256PublishX5cParameter) {
		toSerialize["p256PublishX5cParameter"] = o.P256PublishX5cParameter
	}
	if !IsNil(o.P384ActiveCertRef) {
		toSerialize["p384ActiveCertRef"] = o.P384ActiveCertRef
	}
	if !IsNil(o.P384PreviousCertRef) {
		toSerialize["p384PreviousCertRef"] = o.P384PreviousCertRef
	}
	if !IsNil(o.P384PublishX5cParameter) {
		toSerialize["p384PublishX5cParameter"] = o.P384PublishX5cParameter
	}
	if !IsNil(o.P521ActiveCertRef) {
		toSerialize["p521ActiveCertRef"] = o.P521ActiveCertRef
	}
	if !IsNil(o.P521PreviousCertRef) {
		toSerialize["p521PreviousCertRef"] = o.P521PreviousCertRef
	}
	if !IsNil(o.P521PublishX5cParameter) {
		toSerialize["p521PublishX5cParameter"] = o.P521PublishX5cParameter
	}
	if !IsNil(o.RsaActiveCertRef) {
		toSerialize["rsaActiveCertRef"] = o.RsaActiveCertRef
	}
	if !IsNil(o.RsaPreviousCertRef) {
		toSerialize["rsaPreviousCertRef"] = o.RsaPreviousCertRef
	}
	if !IsNil(o.RsaPublishX5cParameter) {
		toSerialize["rsaPublishX5cParameter"] = o.RsaPublishX5cParameter
	}
	if !IsNil(o.P256DecryptionActiveCertRef) {
		toSerialize["p256DecryptionActiveCertRef"] = o.P256DecryptionActiveCertRef
	}
	if !IsNil(o.P256DecryptionPreviousCertRef) {
		toSerialize["p256DecryptionPreviousCertRef"] = o.P256DecryptionPreviousCertRef
	}
	if !IsNil(o.P256DecryptionPublishX5cParameter) {
		toSerialize["p256DecryptionPublishX5cParameter"] = o.P256DecryptionPublishX5cParameter
	}
	if !IsNil(o.P384DecryptionActiveCertRef) {
		toSerialize["p384DecryptionActiveCertRef"] = o.P384DecryptionActiveCertRef
	}
	if !IsNil(o.P384DecryptionPreviousCertRef) {
		toSerialize["p384DecryptionPreviousCertRef"] = o.P384DecryptionPreviousCertRef
	}
	if !IsNil(o.P384DecryptionPublishX5cParameter) {
		toSerialize["p384DecryptionPublishX5cParameter"] = o.P384DecryptionPublishX5cParameter
	}
	if !IsNil(o.P521DecryptionActiveCertRef) {
		toSerialize["p521DecryptionActiveCertRef"] = o.P521DecryptionActiveCertRef
	}
	if !IsNil(o.P521DecryptionPreviousCertRef) {
		toSerialize["p521DecryptionPreviousCertRef"] = o.P521DecryptionPreviousCertRef
	}
	if !IsNil(o.P521DecryptionPublishX5cParameter) {
		toSerialize["p521DecryptionPublishX5cParameter"] = o.P521DecryptionPublishX5cParameter
	}
	if !IsNil(o.RsaDecryptionActiveCertRef) {
		toSerialize["rsaDecryptionActiveCertRef"] = o.RsaDecryptionActiveCertRef
	}
	if !IsNil(o.RsaDecryptionPreviousCertRef) {
		toSerialize["rsaDecryptionPreviousCertRef"] = o.RsaDecryptionPreviousCertRef
	}
	if !IsNil(o.RsaDecryptionPublishX5cParameter) {
		toSerialize["rsaDecryptionPublishX5cParameter"] = o.RsaDecryptionPublishX5cParameter
	}
	return toSerialize, nil
}

type NullableOAuthOidcKeysSettings struct {
	value *OAuthOidcKeysSettings
	isSet bool
}

func (v NullableOAuthOidcKeysSettings) Get() *OAuthOidcKeysSettings {
	return v.value
}

func (v *NullableOAuthOidcKeysSettings) Set(val *OAuthOidcKeysSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuthOidcKeysSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuthOidcKeysSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuthOidcKeysSettings(val *OAuthOidcKeysSettings) *NullableOAuthOidcKeysSettings {
	return &NullableOAuthOidcKeysSettings{value: val, isSet: true}
}

func (v NullableOAuthOidcKeysSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuthOidcKeysSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
