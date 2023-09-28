# AlternativeLoginHintTokenIssuer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issuer** | **string** | The issuer. Issuer is unique. | 
**JwksURL** | Pointer to **string** | The JWKS URL. | [optional] 
**Jwks** | Pointer to **string** | The JWKS. | [optional] 

## Methods

### NewAlternativeLoginHintTokenIssuer

`func NewAlternativeLoginHintTokenIssuer(issuer string, ) *AlternativeLoginHintTokenIssuer`

NewAlternativeLoginHintTokenIssuer instantiates a new AlternativeLoginHintTokenIssuer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlternativeLoginHintTokenIssuerWithDefaults

`func NewAlternativeLoginHintTokenIssuerWithDefaults() *AlternativeLoginHintTokenIssuer`

NewAlternativeLoginHintTokenIssuerWithDefaults instantiates a new AlternativeLoginHintTokenIssuer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuer

`func (o *AlternativeLoginHintTokenIssuer) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *AlternativeLoginHintTokenIssuer) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *AlternativeLoginHintTokenIssuer) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetJwksURL

`func (o *AlternativeLoginHintTokenIssuer) GetJwksURL() string`

GetJwksURL returns the JwksURL field if non-nil, zero value otherwise.

### GetJwksURLOk

`func (o *AlternativeLoginHintTokenIssuer) GetJwksURLOk() (*string, bool)`

GetJwksURLOk returns a tuple with the JwksURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksURL

`func (o *AlternativeLoginHintTokenIssuer) SetJwksURL(v string)`

SetJwksURL sets JwksURL field to given value.

### HasJwksURL

`func (o *AlternativeLoginHintTokenIssuer) HasJwksURL() bool`

HasJwksURL returns a boolean if a field has been set.

### GetJwks

`func (o *AlternativeLoginHintTokenIssuer) GetJwks() string`

GetJwks returns the Jwks field if non-nil, zero value otherwise.

### GetJwksOk

`func (o *AlternativeLoginHintTokenIssuer) GetJwksOk() (*string, bool)`

GetJwksOk returns a tuple with the Jwks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwks

`func (o *AlternativeLoginHintTokenIssuer) SetJwks(v string)`

SetJwks sets Jwks field to given value.

### HasJwks

`func (o *AlternativeLoginHintTokenIssuer) HasJwks() bool`

HasJwks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


