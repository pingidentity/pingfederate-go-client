# CaptchaSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteKey** | Pointer to **string** | Site key for reCAPTCHA. | [optional] 
**SecretKey** | Pointer to **string** | Secret key for reCAPTCHA. GETs will not return this attribute. To update this field, specify the new value in this attribute. | [optional] 
**EncryptedSecretKey** | Pointer to **string** | The encrypted secret key for reCAPTCHA. If you do not want to update the stored value, this attribute should be passed back unchanged. | [optional] 

## Methods

### NewCaptchaSettings

`func NewCaptchaSettings() *CaptchaSettings`

NewCaptchaSettings instantiates a new CaptchaSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptchaSettingsWithDefaults

`func NewCaptchaSettingsWithDefaults() *CaptchaSettings`

NewCaptchaSettingsWithDefaults instantiates a new CaptchaSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteKey

`func (o *CaptchaSettings) GetSiteKey() string`

GetSiteKey returns the SiteKey field if non-nil, zero value otherwise.

### GetSiteKeyOk

`func (o *CaptchaSettings) GetSiteKeyOk() (*string, bool)`

GetSiteKeyOk returns a tuple with the SiteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteKey

`func (o *CaptchaSettings) SetSiteKey(v string)`

SetSiteKey sets SiteKey field to given value.

### HasSiteKey

`func (o *CaptchaSettings) HasSiteKey() bool`

HasSiteKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *CaptchaSettings) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *CaptchaSettings) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *CaptchaSettings) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *CaptchaSettings) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetEncryptedSecretKey

`func (o *CaptchaSettings) GetEncryptedSecretKey() string`

GetEncryptedSecretKey returns the EncryptedSecretKey field if non-nil, zero value otherwise.

### GetEncryptedSecretKeyOk

`func (o *CaptchaSettings) GetEncryptedSecretKeyOk() (*string, bool)`

GetEncryptedSecretKeyOk returns a tuple with the EncryptedSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedSecretKey

`func (o *CaptchaSettings) SetEncryptedSecretKey(v string)`

SetEncryptedSecretKey sets EncryptedSecretKey field to given value.

### HasEncryptedSecretKey

`func (o *CaptchaSettings) HasEncryptedSecretKey() bool`

HasEncryptedSecretKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


