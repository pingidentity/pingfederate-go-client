# Action

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of this action. | [optional] 
**Name** | Pointer to **string** | The name of this action. | [optional] 
**Description** | Pointer to **string** | The description of this action. | [optional] 
**Download** | Pointer to **bool** | Whether this action will trigger a download or invoke an internal action that will return a string result. | [optional] 
**InvocationRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**Parameters** | Pointer to [**[]FieldDescriptor**](FieldDescriptor.md) | List of parameters for this action. | [optional] 

## Methods

### NewAction

`func NewAction() *Action`

NewAction instantiates a new Action object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionWithDefaults

`func NewActionWithDefaults() *Action`

NewActionWithDefaults instantiates a new Action object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Action) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Action) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Action) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Action) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Action) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Action) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Action) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Action) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Action) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Action) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Action) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Action) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDownload

`func (o *Action) GetDownload() bool`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *Action) GetDownloadOk() (*bool, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *Action) SetDownload(v bool)`

SetDownload sets Download field to given value.

### HasDownload

`func (o *Action) HasDownload() bool`

HasDownload returns a boolean if a field has been set.

### GetInvocationRef

`func (o *Action) GetInvocationRef() ResourceLink`

GetInvocationRef returns the InvocationRef field if non-nil, zero value otherwise.

### GetInvocationRefOk

`func (o *Action) GetInvocationRefOk() (*ResourceLink, bool)`

GetInvocationRefOk returns a tuple with the InvocationRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvocationRef

`func (o *Action) SetInvocationRef(v ResourceLink)`

SetInvocationRef sets InvocationRef field to given value.

### HasInvocationRef

`func (o *Action) HasInvocationRef() bool`

HasInvocationRef returns a boolean if a field has been set.

### GetParameters

`func (o *Action) GetParameters() []FieldDescriptor`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *Action) GetParametersOk() (*[]FieldDescriptor, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *Action) SetParameters(v []FieldDescriptor)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *Action) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


