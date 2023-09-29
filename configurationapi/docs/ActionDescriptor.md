# ActionDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of this action | [optional] 
**Description** | Pointer to **string** | The description of this action | [optional] 
**Download** | Pointer to **bool** | Whether this action will trigger a download or invoke an internal action that will return a string result. | [optional] 
**DownloadContentType** | Pointer to **string** | If this is a download, this is the Content-Type of the downloaded file. Otherwise, this is null. | [optional] 
**DownloadFileName** | Pointer to **string** | If this is a download, this is the suggested file name of the downloaded file. Otherwise, this is null. | [optional] 
**Parameters** | Pointer to [**[]FieldDescriptor**](FieldDescriptor.md) | List of parameters for this action. | [optional] 

## Methods

### NewActionDescriptor

`func NewActionDescriptor() *ActionDescriptor`

NewActionDescriptor instantiates a new ActionDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionDescriptorWithDefaults

`func NewActionDescriptorWithDefaults() *ActionDescriptor`

NewActionDescriptorWithDefaults instantiates a new ActionDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ActionDescriptor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActionDescriptor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActionDescriptor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ActionDescriptor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ActionDescriptor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ActionDescriptor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ActionDescriptor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ActionDescriptor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDownload

`func (o *ActionDescriptor) GetDownload() bool`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *ActionDescriptor) GetDownloadOk() (*bool, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *ActionDescriptor) SetDownload(v bool)`

SetDownload sets Download field to given value.

### HasDownload

`func (o *ActionDescriptor) HasDownload() bool`

HasDownload returns a boolean if a field has been set.

### GetDownloadContentType

`func (o *ActionDescriptor) GetDownloadContentType() string`

GetDownloadContentType returns the DownloadContentType field if non-nil, zero value otherwise.

### GetDownloadContentTypeOk

`func (o *ActionDescriptor) GetDownloadContentTypeOk() (*string, bool)`

GetDownloadContentTypeOk returns a tuple with the DownloadContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadContentType

`func (o *ActionDescriptor) SetDownloadContentType(v string)`

SetDownloadContentType sets DownloadContentType field to given value.

### HasDownloadContentType

`func (o *ActionDescriptor) HasDownloadContentType() bool`

HasDownloadContentType returns a boolean if a field has been set.

### GetDownloadFileName

`func (o *ActionDescriptor) GetDownloadFileName() string`

GetDownloadFileName returns the DownloadFileName field if non-nil, zero value otherwise.

### GetDownloadFileNameOk

`func (o *ActionDescriptor) GetDownloadFileNameOk() (*string, bool)`

GetDownloadFileNameOk returns a tuple with the DownloadFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadFileName

`func (o *ActionDescriptor) SetDownloadFileName(v string)`

SetDownloadFileName sets DownloadFileName field to given value.

### HasDownloadFileName

`func (o *ActionDescriptor) HasDownloadFileName() bool`

HasDownloadFileName returns a boolean if a field has been set.

### GetParameters

`func (o *ActionDescriptor) GetParameters() []FieldDescriptor`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ActionDescriptor) GetParametersOk() (*[]FieldDescriptor, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ActionDescriptor) SetParameters(v []FieldDescriptor)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ActionDescriptor) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


