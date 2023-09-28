# RegistrationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaptchaEnabled** | Pointer to **bool** | Whether CAPTCHA is enabled or not in the registration configuration. | [optional] 
**CaptchaProviderRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**TemplateName** | **string** | The template name for the registration configuration. | 
**CreateAuthnSessionAfterRegistration** | Pointer to **bool** | Whether to create an Authentication Session when registering a local account. Default is true. | [optional] 
**UsernameField** | Pointer to **string** | When creating an Authentication Session after registering a local account, PingFederate will pass the Unique ID field&#39;s value as the username. If the Unique ID value is not the username, then override which field&#39;s value will be used as the username. | [optional] 
**ThisIsMyDeviceEnabled** | Pointer to **bool** | Allows users to indicate whether their device is shared or private. In this mode, PingFederate Authentication Sessions will not be stored unless the user indicates the device is private. | [optional] 
**RegistrationWorkflow** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**ExecuteWorkflow** | Pointer to **string** | This setting indicates whether PingFederate should execute the workflow before or after account creation. The default is to run the registration workflow after account creation. | [optional] 

## Methods

### NewRegistrationConfig

`func NewRegistrationConfig(templateName string, ) *RegistrationConfig`

NewRegistrationConfig instantiates a new RegistrationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrationConfigWithDefaults

`func NewRegistrationConfigWithDefaults() *RegistrationConfig`

NewRegistrationConfigWithDefaults instantiates a new RegistrationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaptchaEnabled

`func (o *RegistrationConfig) GetCaptchaEnabled() bool`

GetCaptchaEnabled returns the CaptchaEnabled field if non-nil, zero value otherwise.

### GetCaptchaEnabledOk

`func (o *RegistrationConfig) GetCaptchaEnabledOk() (*bool, bool)`

GetCaptchaEnabledOk returns a tuple with the CaptchaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptchaEnabled

`func (o *RegistrationConfig) SetCaptchaEnabled(v bool)`

SetCaptchaEnabled sets CaptchaEnabled field to given value.

### HasCaptchaEnabled

`func (o *RegistrationConfig) HasCaptchaEnabled() bool`

HasCaptchaEnabled returns a boolean if a field has been set.

### GetCaptchaProviderRef

`func (o *RegistrationConfig) GetCaptchaProviderRef() ResourceLink`

GetCaptchaProviderRef returns the CaptchaProviderRef field if non-nil, zero value otherwise.

### GetCaptchaProviderRefOk

`func (o *RegistrationConfig) GetCaptchaProviderRefOk() (*ResourceLink, bool)`

GetCaptchaProviderRefOk returns a tuple with the CaptchaProviderRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptchaProviderRef

`func (o *RegistrationConfig) SetCaptchaProviderRef(v ResourceLink)`

SetCaptchaProviderRef sets CaptchaProviderRef field to given value.

### HasCaptchaProviderRef

`func (o *RegistrationConfig) HasCaptchaProviderRef() bool`

HasCaptchaProviderRef returns a boolean if a field has been set.

### GetTemplateName

`func (o *RegistrationConfig) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *RegistrationConfig) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *RegistrationConfig) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.


### GetCreateAuthnSessionAfterRegistration

`func (o *RegistrationConfig) GetCreateAuthnSessionAfterRegistration() bool`

GetCreateAuthnSessionAfterRegistration returns the CreateAuthnSessionAfterRegistration field if non-nil, zero value otherwise.

### GetCreateAuthnSessionAfterRegistrationOk

`func (o *RegistrationConfig) GetCreateAuthnSessionAfterRegistrationOk() (*bool, bool)`

GetCreateAuthnSessionAfterRegistrationOk returns a tuple with the CreateAuthnSessionAfterRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAuthnSessionAfterRegistration

`func (o *RegistrationConfig) SetCreateAuthnSessionAfterRegistration(v bool)`

SetCreateAuthnSessionAfterRegistration sets CreateAuthnSessionAfterRegistration field to given value.

### HasCreateAuthnSessionAfterRegistration

`func (o *RegistrationConfig) HasCreateAuthnSessionAfterRegistration() bool`

HasCreateAuthnSessionAfterRegistration returns a boolean if a field has been set.

### GetUsernameField

`func (o *RegistrationConfig) GetUsernameField() string`

GetUsernameField returns the UsernameField field if non-nil, zero value otherwise.

### GetUsernameFieldOk

`func (o *RegistrationConfig) GetUsernameFieldOk() (*string, bool)`

GetUsernameFieldOk returns a tuple with the UsernameField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameField

`func (o *RegistrationConfig) SetUsernameField(v string)`

SetUsernameField sets UsernameField field to given value.

### HasUsernameField

`func (o *RegistrationConfig) HasUsernameField() bool`

HasUsernameField returns a boolean if a field has been set.

### GetThisIsMyDeviceEnabled

`func (o *RegistrationConfig) GetThisIsMyDeviceEnabled() bool`

GetThisIsMyDeviceEnabled returns the ThisIsMyDeviceEnabled field if non-nil, zero value otherwise.

### GetThisIsMyDeviceEnabledOk

`func (o *RegistrationConfig) GetThisIsMyDeviceEnabledOk() (*bool, bool)`

GetThisIsMyDeviceEnabledOk returns a tuple with the ThisIsMyDeviceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThisIsMyDeviceEnabled

`func (o *RegistrationConfig) SetThisIsMyDeviceEnabled(v bool)`

SetThisIsMyDeviceEnabled sets ThisIsMyDeviceEnabled field to given value.

### HasThisIsMyDeviceEnabled

`func (o *RegistrationConfig) HasThisIsMyDeviceEnabled() bool`

HasThisIsMyDeviceEnabled returns a boolean if a field has been set.

### GetRegistrationWorkflow

`func (o *RegistrationConfig) GetRegistrationWorkflow() ResourceLink`

GetRegistrationWorkflow returns the RegistrationWorkflow field if non-nil, zero value otherwise.

### GetRegistrationWorkflowOk

`func (o *RegistrationConfig) GetRegistrationWorkflowOk() (*ResourceLink, bool)`

GetRegistrationWorkflowOk returns a tuple with the RegistrationWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationWorkflow

`func (o *RegistrationConfig) SetRegistrationWorkflow(v ResourceLink)`

SetRegistrationWorkflow sets RegistrationWorkflow field to given value.

### HasRegistrationWorkflow

`func (o *RegistrationConfig) HasRegistrationWorkflow() bool`

HasRegistrationWorkflow returns a boolean if a field has been set.

### GetExecuteWorkflow

`func (o *RegistrationConfig) GetExecuteWorkflow() string`

GetExecuteWorkflow returns the ExecuteWorkflow field if non-nil, zero value otherwise.

### GetExecuteWorkflowOk

`func (o *RegistrationConfig) GetExecuteWorkflowOk() (*string, bool)`

GetExecuteWorkflowOk returns a tuple with the ExecuteWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteWorkflow

`func (o *RegistrationConfig) SetExecuteWorkflow(v string)`

SetExecuteWorkflow sets ExecuteWorkflow field to given value.

### HasExecuteWorkflow

`func (o *RegistrationConfig) HasExecuteWorkflow() bool`

HasExecuteWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


