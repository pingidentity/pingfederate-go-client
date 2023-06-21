# MetadataSigningSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SigningKeyRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**SignatureAlgorithm** | Pointer to **string** | Signature algorithm. If this property is unset, the default signature algorithm for the key algorithm will be used. Supported signature algorithms are available through the /keyPairs/keyAlgorithms endpoint. | [optional] 

## Methods

### NewMetadataSigningSettings

`func NewMetadataSigningSettings() *MetadataSigningSettings`

NewMetadataSigningSettings instantiates a new MetadataSigningSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataSigningSettingsWithDefaults

`func NewMetadataSigningSettingsWithDefaults() *MetadataSigningSettings`

NewMetadataSigningSettingsWithDefaults instantiates a new MetadataSigningSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSigningKeyRef

`func (o *MetadataSigningSettings) GetSigningKeyRef() ResourceLink`

GetSigningKeyRef returns the SigningKeyRef field if non-nil, zero value otherwise.

### GetSigningKeyRefOk

`func (o *MetadataSigningSettings) GetSigningKeyRefOk() (*ResourceLink, bool)`

GetSigningKeyRefOk returns a tuple with the SigningKeyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKeyRef

`func (o *MetadataSigningSettings) SetSigningKeyRef(v ResourceLink)`

SetSigningKeyRef sets SigningKeyRef field to given value.

### HasSigningKeyRef

`func (o *MetadataSigningSettings) HasSigningKeyRef() bool`

HasSigningKeyRef returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *MetadataSigningSettings) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *MetadataSigningSettings) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *MetadataSigningSettings) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.

### HasSignatureAlgorithm

`func (o *MetadataSigningSettings) HasSignatureAlgorithm() bool`

HasSignatureAlgorithm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


