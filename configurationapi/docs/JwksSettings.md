# JwksSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JwksUrl** | Pointer to **string** | JSON Web Key Set (JWKS) URL of the OAuth client. Either &#39;jwks&#39; or &#39;jwksUrl&#39; must be provided if private key JWT client authentication or signed requests is enabled.  If the client signs its JWTs using an RSASSA-PSS signing algorithm, PingFederate must either use Java 11 or be integrated with a hardware security module (HSM) to process the digital signatures. | [optional] 
**Jwks** | Pointer to **string** | JSON Web Key Set (JWKS) document of the OAuth client. Either &#39;jwks&#39; or &#39;jwksUrl&#39; must be provided if private key JWT client authentication or signed requests is enabled.  If the client signs its JWTs using an RSASSA-PSS signing algorithm, PingFederate must either use Java 11 or be integrated with a hardware security module (HSM) to process the digital signatures. | [optional] 

## Methods

### NewJwksSettings

`func NewJwksSettings() *JwksSettings`

NewJwksSettings instantiates a new JwksSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJwksSettingsWithDefaults

`func NewJwksSettingsWithDefaults() *JwksSettings`

NewJwksSettingsWithDefaults instantiates a new JwksSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJwksUrl

`func (o *JwksSettings) GetJwksUrl() string`

GetJwksUrl returns the JwksUrl field if non-nil, zero value otherwise.

### GetJwksUrlOk

`func (o *JwksSettings) GetJwksUrlOk() (*string, bool)`

GetJwksUrlOk returns a tuple with the JwksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUrl

`func (o *JwksSettings) SetJwksUrl(v string)`

SetJwksUrl sets JwksUrl field to given value.

### HasJwksUrl

`func (o *JwksSettings) HasJwksUrl() bool`

HasJwksUrl returns a boolean if a field has been set.

### GetJwks

`func (o *JwksSettings) GetJwks() string`

GetJwks returns the Jwks field if non-nil, zero value otherwise.

### GetJwksOk

`func (o *JwksSettings) GetJwksOk() (*string, bool)`

GetJwksOk returns a tuple with the Jwks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwks

`func (o *JwksSettings) SetJwks(v string)`

SetJwks sets Jwks field to given value.

### HasJwks

`func (o *JwksSettings) HasJwks() bool`

HasJwks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


