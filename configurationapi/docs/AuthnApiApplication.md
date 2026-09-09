# AuthnApiApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The persistent, unique ID for the Authentication API application. It can be any combination of [a-zA-Z0-9._-]. This property is system-assigned if not specified. | 
**Name** | **string** | The Authentication API Application Name. Name must be unique. | 
**Url** | **string** | The Authentication API Application redirect URL. | 
**Description** | Pointer to **string** | The Authentication API Application description. | [optional] 
**AdditionalAllowedOrigins** | Pointer to **[]string** | The domain in the redirect URL is always whitelisted. This field contains a list of additional allowed origin URL&#39;s for cross-origin resource sharing. | [optional] 
**ClientForRedirectlessModeRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 

## Methods

### NewAuthnApiApplication

`func NewAuthnApiApplication(id string, name string, url string, ) *AuthnApiApplication`

NewAuthnApiApplication instantiates a new AuthnApiApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthnApiApplicationWithDefaults

`func NewAuthnApiApplicationWithDefaults() *AuthnApiApplication`

NewAuthnApiApplicationWithDefaults instantiates a new AuthnApiApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthnApiApplication) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthnApiApplication) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthnApiApplication) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AuthnApiApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthnApiApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthnApiApplication) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *AuthnApiApplication) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AuthnApiApplication) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AuthnApiApplication) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDescription

`func (o *AuthnApiApplication) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthnApiApplication) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthnApiApplication) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthnApiApplication) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAdditionalAllowedOrigins

`func (o *AuthnApiApplication) GetAdditionalAllowedOrigins() []string`

GetAdditionalAllowedOrigins returns the AdditionalAllowedOrigins field if non-nil, zero value otherwise.

### GetAdditionalAllowedOriginsOk

`func (o *AuthnApiApplication) GetAdditionalAllowedOriginsOk() (*[]string, bool)`

GetAdditionalAllowedOriginsOk returns a tuple with the AdditionalAllowedOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAllowedOrigins

`func (o *AuthnApiApplication) SetAdditionalAllowedOrigins(v []string)`

SetAdditionalAllowedOrigins sets AdditionalAllowedOrigins field to given value.

### HasAdditionalAllowedOrigins

`func (o *AuthnApiApplication) HasAdditionalAllowedOrigins() bool`

HasAdditionalAllowedOrigins returns a boolean if a field has been set.

### GetClientForRedirectlessModeRef

`func (o *AuthnApiApplication) GetClientForRedirectlessModeRef() ResourceLink`

GetClientForRedirectlessModeRef returns the ClientForRedirectlessModeRef field if non-nil, zero value otherwise.

### GetClientForRedirectlessModeRefOk

`func (o *AuthnApiApplication) GetClientForRedirectlessModeRefOk() (*ResourceLink, bool)`

GetClientForRedirectlessModeRefOk returns a tuple with the ClientForRedirectlessModeRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientForRedirectlessModeRef

`func (o *AuthnApiApplication) SetClientForRedirectlessModeRef(v ResourceLink)`

SetClientForRedirectlessModeRef sets ClientForRedirectlessModeRef field to given value.

### HasClientForRedirectlessModeRef

`func (o *AuthnApiApplication) HasClientForRedirectlessModeRef() bool`

HasClientForRedirectlessModeRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


