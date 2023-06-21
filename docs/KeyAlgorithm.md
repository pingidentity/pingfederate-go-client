# KeyAlgorithm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the key algorithm. | [optional] 
**KeySizes** | Pointer to **[]int64** | Possible key sizes for this algorithm, in bits. | [optional] 
**DefaultKeySize** | Pointer to **int64** | Default key size for this algorithm. | [optional] 
**SignatureAlgorithms** | Pointer to **[]string** | Possible signature algorithms for this key algorithm. | [optional] 
**DefaultSignatureAlgorithm** | Pointer to **string** | Default signature algorithm for this key algorithm. | [optional] 

## Methods

### NewKeyAlgorithm

`func NewKeyAlgorithm() *KeyAlgorithm`

NewKeyAlgorithm instantiates a new KeyAlgorithm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyAlgorithmWithDefaults

`func NewKeyAlgorithmWithDefaults() *KeyAlgorithm`

NewKeyAlgorithmWithDefaults instantiates a new KeyAlgorithm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KeyAlgorithm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyAlgorithm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyAlgorithm) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KeyAlgorithm) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKeySizes

`func (o *KeyAlgorithm) GetKeySizes() []int64`

GetKeySizes returns the KeySizes field if non-nil, zero value otherwise.

### GetKeySizesOk

`func (o *KeyAlgorithm) GetKeySizesOk() (*[]int64, bool)`

GetKeySizesOk returns a tuple with the KeySizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySizes

`func (o *KeyAlgorithm) SetKeySizes(v []int64)`

SetKeySizes sets KeySizes field to given value.

### HasKeySizes

`func (o *KeyAlgorithm) HasKeySizes() bool`

HasKeySizes returns a boolean if a field has been set.

### GetDefaultKeySize

`func (o *KeyAlgorithm) GetDefaultKeySize() int64`

GetDefaultKeySize returns the DefaultKeySize field if non-nil, zero value otherwise.

### GetDefaultKeySizeOk

`func (o *KeyAlgorithm) GetDefaultKeySizeOk() (*int64, bool)`

GetDefaultKeySizeOk returns a tuple with the DefaultKeySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultKeySize

`func (o *KeyAlgorithm) SetDefaultKeySize(v int64)`

SetDefaultKeySize sets DefaultKeySize field to given value.

### HasDefaultKeySize

`func (o *KeyAlgorithm) HasDefaultKeySize() bool`

HasDefaultKeySize returns a boolean if a field has been set.

### GetSignatureAlgorithms

`func (o *KeyAlgorithm) GetSignatureAlgorithms() []string`

GetSignatureAlgorithms returns the SignatureAlgorithms field if non-nil, zero value otherwise.

### GetSignatureAlgorithmsOk

`func (o *KeyAlgorithm) GetSignatureAlgorithmsOk() (*[]string, bool)`

GetSignatureAlgorithmsOk returns a tuple with the SignatureAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithms

`func (o *KeyAlgorithm) SetSignatureAlgorithms(v []string)`

SetSignatureAlgorithms sets SignatureAlgorithms field to given value.

### HasSignatureAlgorithms

`func (o *KeyAlgorithm) HasSignatureAlgorithms() bool`

HasSignatureAlgorithms returns a boolean if a field has been set.

### GetDefaultSignatureAlgorithm

`func (o *KeyAlgorithm) GetDefaultSignatureAlgorithm() string`

GetDefaultSignatureAlgorithm returns the DefaultSignatureAlgorithm field if non-nil, zero value otherwise.

### GetDefaultSignatureAlgorithmOk

`func (o *KeyAlgorithm) GetDefaultSignatureAlgorithmOk() (*string, bool)`

GetDefaultSignatureAlgorithmOk returns a tuple with the DefaultSignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSignatureAlgorithm

`func (o *KeyAlgorithm) SetDefaultSignatureAlgorithm(v string)`

SetDefaultSignatureAlgorithm sets DefaultSignatureAlgorithm field to given value.

### HasDefaultSignatureAlgorithm

`func (o *KeyAlgorithm) HasDefaultSignatureAlgorithm() bool`

HasDefaultSignatureAlgorithm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


