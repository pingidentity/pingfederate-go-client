# BackChannelAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The back channel authentication type. | 
**HttpBasicCredentials** | Pointer to [**UsernamePasswordCredentials**](UsernamePasswordCredentials.md) |  | [optional] 
**DigitalSignature** | Pointer to **bool** | If incoming or outgoing messages must be signed. | [optional] 

## Methods

### NewBackChannelAuth

`func NewBackChannelAuth(type_ string, ) *BackChannelAuth`

NewBackChannelAuth instantiates a new BackChannelAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackChannelAuthWithDefaults

`func NewBackChannelAuthWithDefaults() *BackChannelAuth`

NewBackChannelAuthWithDefaults instantiates a new BackChannelAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BackChannelAuth) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BackChannelAuth) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BackChannelAuth) SetType(v string)`

SetType sets Type field to given value.


### GetHttpBasicCredentials

`func (o *BackChannelAuth) GetHttpBasicCredentials() UsernamePasswordCredentials`

GetHttpBasicCredentials returns the HttpBasicCredentials field if non-nil, zero value otherwise.

### GetHttpBasicCredentialsOk

`func (o *BackChannelAuth) GetHttpBasicCredentialsOk() (*UsernamePasswordCredentials, bool)`

GetHttpBasicCredentialsOk returns a tuple with the HttpBasicCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpBasicCredentials

`func (o *BackChannelAuth) SetHttpBasicCredentials(v UsernamePasswordCredentials)`

SetHttpBasicCredentials sets HttpBasicCredentials field to given value.

### HasHttpBasicCredentials

`func (o *BackChannelAuth) HasHttpBasicCredentials() bool`

HasHttpBasicCredentials returns a boolean if a field has been set.

### GetDigitalSignature

`func (o *BackChannelAuth) GetDigitalSignature() bool`

GetDigitalSignature returns the DigitalSignature field if non-nil, zero value otherwise.

### GetDigitalSignatureOk

`func (o *BackChannelAuth) GetDigitalSignatureOk() (*bool, bool)`

GetDigitalSignatureOk returns a tuple with the DigitalSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalSignature

`func (o *BackChannelAuth) SetDigitalSignature(v bool)`

SetDigitalSignature sets DigitalSignature field to given value.

### HasDigitalSignature

`func (o *BackChannelAuth) HasDigitalSignature() bool`

HasDigitalSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


