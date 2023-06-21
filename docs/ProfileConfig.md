# ProfileConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteIdentityEnabled** | Pointer to **bool** | Whether the end user is allowed to use delete functionality. | [optional] 
**TemplateName** | **string** | The template name for end-user profile management. | 

## Methods

### NewProfileConfig

`func NewProfileConfig(templateName string, ) *ProfileConfig`

NewProfileConfig instantiates a new ProfileConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileConfigWithDefaults

`func NewProfileConfigWithDefaults() *ProfileConfig`

NewProfileConfigWithDefaults instantiates a new ProfileConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteIdentityEnabled

`func (o *ProfileConfig) GetDeleteIdentityEnabled() bool`

GetDeleteIdentityEnabled returns the DeleteIdentityEnabled field if non-nil, zero value otherwise.

### GetDeleteIdentityEnabledOk

`func (o *ProfileConfig) GetDeleteIdentityEnabledOk() (*bool, bool)`

GetDeleteIdentityEnabledOk returns a tuple with the DeleteIdentityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteIdentityEnabled

`func (o *ProfileConfig) SetDeleteIdentityEnabled(v bool)`

SetDeleteIdentityEnabled sets DeleteIdentityEnabled field to given value.

### HasDeleteIdentityEnabled

`func (o *ProfileConfig) HasDeleteIdentityEnabled() bool`

HasDeleteIdentityEnabled returns a boolean if a field has been set.

### GetTemplateName

`func (o *ProfileConfig) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *ProfileConfig) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *ProfileConfig) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


