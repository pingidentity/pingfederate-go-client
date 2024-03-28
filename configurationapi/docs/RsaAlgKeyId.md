# RsaAlgKeyId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RsaAlgType** | **string** | The RSA signing algorithm type. The supported RSA signing algorithm types are RS256, RS384, RS512, PS256, PS384 and PS512. | 
**KeyId** | **string** | Unique key identifier. | 

## Methods

### NewRsaAlgKeyId

`func NewRsaAlgKeyId(rsaAlgType string, keyId string, ) *RsaAlgKeyId`

NewRsaAlgKeyId instantiates a new RsaAlgKeyId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRsaAlgKeyIdWithDefaults

`func NewRsaAlgKeyIdWithDefaults() *RsaAlgKeyId`

NewRsaAlgKeyIdWithDefaults instantiates a new RsaAlgKeyId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRsaAlgType

`func (o *RsaAlgKeyId) GetRsaAlgType() string`

GetRsaAlgType returns the RsaAlgType field if non-nil, zero value otherwise.

### GetRsaAlgTypeOk

`func (o *RsaAlgKeyId) GetRsaAlgTypeOk() (*string, bool)`

GetRsaAlgTypeOk returns a tuple with the RsaAlgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsaAlgType

`func (o *RsaAlgKeyId) SetRsaAlgType(v string)`

SetRsaAlgType sets RsaAlgType field to given value.


### GetKeyId

`func (o *RsaAlgKeyId) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *RsaAlgKeyId) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *RsaAlgKeyId) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


