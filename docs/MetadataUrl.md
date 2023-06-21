# MetadataUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The persistent, unique ID for the Metadata Url. It can be any combination of [a-z0-9._-]. This property is system-assigned if not specified. | [optional] 
**Name** | **string** | The name for the Metadata URL. | 
**Url** | **string** | The Metadata URL. | 
**CertView** | Pointer to [**CertView**](CertView.md) |  | [optional] 
**X509File** | Pointer to [**X509File**](X509File.md) |  | [optional] 
**ValidateSignature** | Pointer to **bool** | Perform Metadata Signature Validation. The default value is TRUE. | [optional] 

## Methods

### NewMetadataUrl

`func NewMetadataUrl(name string, url string, ) *MetadataUrl`

NewMetadataUrl instantiates a new MetadataUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataUrlWithDefaults

`func NewMetadataUrlWithDefaults() *MetadataUrl`

NewMetadataUrlWithDefaults instantiates a new MetadataUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetadataUrl) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetadataUrl) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetadataUrl) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MetadataUrl) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MetadataUrl) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetadataUrl) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetadataUrl) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *MetadataUrl) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MetadataUrl) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MetadataUrl) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCertView

`func (o *MetadataUrl) GetCertView() CertView`

GetCertView returns the CertView field if non-nil, zero value otherwise.

### GetCertViewOk

`func (o *MetadataUrl) GetCertViewOk() (*CertView, bool)`

GetCertViewOk returns a tuple with the CertView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertView

`func (o *MetadataUrl) SetCertView(v CertView)`

SetCertView sets CertView field to given value.

### HasCertView

`func (o *MetadataUrl) HasCertView() bool`

HasCertView returns a boolean if a field has been set.

### GetX509File

`func (o *MetadataUrl) GetX509File() X509File`

GetX509File returns the X509File field if non-nil, zero value otherwise.

### GetX509FileOk

`func (o *MetadataUrl) GetX509FileOk() (*X509File, bool)`

GetX509FileOk returns a tuple with the X509File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509File

`func (o *MetadataUrl) SetX509File(v X509File)`

SetX509File sets X509File field to given value.

### HasX509File

`func (o *MetadataUrl) HasX509File() bool`

HasX509File returns a boolean if a field has been set.

### GetValidateSignature

`func (o *MetadataUrl) GetValidateSignature() bool`

GetValidateSignature returns the ValidateSignature field if non-nil, zero value otherwise.

### GetValidateSignatureOk

`func (o *MetadataUrl) GetValidateSignatureOk() (*bool, bool)`

GetValidateSignatureOk returns a tuple with the ValidateSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateSignature

`func (o *MetadataUrl) SetValidateSignature(v bool)`

SetValidateSignature sets ValidateSignature field to given value.

### HasValidateSignature

`func (o *MetadataUrl) HasValidateSignature() bool`

HasValidateSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


