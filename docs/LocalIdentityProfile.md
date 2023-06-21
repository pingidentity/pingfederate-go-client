# LocalIdentityProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The persistent, unique ID for the local identity profile. It can be any combination of [a-zA-Z0-9._-]. This property is system-assigned if not specified. | [optional] 
**Name** | **string** | The local identity profile name. Name is unique. | 
**ApcId** | [**ResourceLink**](ResourceLink.md) |  | 
**AuthSources** | Pointer to [**[]LocalIdentityAuthSource**](LocalIdentityAuthSource.md) | The local identity authentication sources. Sources are unique. | [optional] 
**AuthSourceUpdatePolicy** | Pointer to [**LocalIdentityAuthSourceUpdatePolicy**](LocalIdentityAuthSourceUpdatePolicy.md) |  | [optional] 
**RegistrationEnabled** | Pointer to **bool** | Whether the registration configuration is enabled or not. | [optional] 
**RegistrationConfig** | Pointer to [**RegistrationConfig**](RegistrationConfig.md) |  | [optional] 
**ProfileConfig** | Pointer to [**ProfileConfig**](ProfileConfig.md) |  | [optional] 
**FieldConfig** | Pointer to [**FieldConfig**](FieldConfig.md) |  | [optional] 
**EmailVerificationConfig** | Pointer to [**EmailVerificationConfig**](EmailVerificationConfig.md) |  | [optional] 
**DataStoreConfig** | Pointer to [**DataStoreConfig**](DataStoreConfig.md) |  | [optional] 
**ProfileEnabled** | Pointer to **bool** | Whether the profile configuration is enabled or not. | [optional] 

## Methods

### NewLocalIdentityProfile

`func NewLocalIdentityProfile(name string, apcId ResourceLink, ) *LocalIdentityProfile`

NewLocalIdentityProfile instantiates a new LocalIdentityProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalIdentityProfileWithDefaults

`func NewLocalIdentityProfileWithDefaults() *LocalIdentityProfile`

NewLocalIdentityProfileWithDefaults instantiates a new LocalIdentityProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LocalIdentityProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LocalIdentityProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LocalIdentityProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LocalIdentityProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *LocalIdentityProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LocalIdentityProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LocalIdentityProfile) SetName(v string)`

SetName sets Name field to given value.


### GetApcId

`func (o *LocalIdentityProfile) GetApcId() ResourceLink`

GetApcId returns the ApcId field if non-nil, zero value otherwise.

### GetApcIdOk

`func (o *LocalIdentityProfile) GetApcIdOk() (*ResourceLink, bool)`

GetApcIdOk returns a tuple with the ApcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApcId

`func (o *LocalIdentityProfile) SetApcId(v ResourceLink)`

SetApcId sets ApcId field to given value.


### GetAuthSources

`func (o *LocalIdentityProfile) GetAuthSources() []LocalIdentityAuthSource`

GetAuthSources returns the AuthSources field if non-nil, zero value otherwise.

### GetAuthSourcesOk

`func (o *LocalIdentityProfile) GetAuthSourcesOk() (*[]LocalIdentityAuthSource, bool)`

GetAuthSourcesOk returns a tuple with the AuthSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSources

`func (o *LocalIdentityProfile) SetAuthSources(v []LocalIdentityAuthSource)`

SetAuthSources sets AuthSources field to given value.

### HasAuthSources

`func (o *LocalIdentityProfile) HasAuthSources() bool`

HasAuthSources returns a boolean if a field has been set.

### GetAuthSourceUpdatePolicy

`func (o *LocalIdentityProfile) GetAuthSourceUpdatePolicy() LocalIdentityAuthSourceUpdatePolicy`

GetAuthSourceUpdatePolicy returns the AuthSourceUpdatePolicy field if non-nil, zero value otherwise.

### GetAuthSourceUpdatePolicyOk

`func (o *LocalIdentityProfile) GetAuthSourceUpdatePolicyOk() (*LocalIdentityAuthSourceUpdatePolicy, bool)`

GetAuthSourceUpdatePolicyOk returns a tuple with the AuthSourceUpdatePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSourceUpdatePolicy

`func (o *LocalIdentityProfile) SetAuthSourceUpdatePolicy(v LocalIdentityAuthSourceUpdatePolicy)`

SetAuthSourceUpdatePolicy sets AuthSourceUpdatePolicy field to given value.

### HasAuthSourceUpdatePolicy

`func (o *LocalIdentityProfile) HasAuthSourceUpdatePolicy() bool`

HasAuthSourceUpdatePolicy returns a boolean if a field has been set.

### GetRegistrationEnabled

`func (o *LocalIdentityProfile) GetRegistrationEnabled() bool`

GetRegistrationEnabled returns the RegistrationEnabled field if non-nil, zero value otherwise.

### GetRegistrationEnabledOk

`func (o *LocalIdentityProfile) GetRegistrationEnabledOk() (*bool, bool)`

GetRegistrationEnabledOk returns a tuple with the RegistrationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationEnabled

`func (o *LocalIdentityProfile) SetRegistrationEnabled(v bool)`

SetRegistrationEnabled sets RegistrationEnabled field to given value.

### HasRegistrationEnabled

`func (o *LocalIdentityProfile) HasRegistrationEnabled() bool`

HasRegistrationEnabled returns a boolean if a field has been set.

### GetRegistrationConfig

`func (o *LocalIdentityProfile) GetRegistrationConfig() RegistrationConfig`

GetRegistrationConfig returns the RegistrationConfig field if non-nil, zero value otherwise.

### GetRegistrationConfigOk

`func (o *LocalIdentityProfile) GetRegistrationConfigOk() (*RegistrationConfig, bool)`

GetRegistrationConfigOk returns a tuple with the RegistrationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationConfig

`func (o *LocalIdentityProfile) SetRegistrationConfig(v RegistrationConfig)`

SetRegistrationConfig sets RegistrationConfig field to given value.

### HasRegistrationConfig

`func (o *LocalIdentityProfile) HasRegistrationConfig() bool`

HasRegistrationConfig returns a boolean if a field has been set.

### GetProfileConfig

`func (o *LocalIdentityProfile) GetProfileConfig() ProfileConfig`

GetProfileConfig returns the ProfileConfig field if non-nil, zero value otherwise.

### GetProfileConfigOk

`func (o *LocalIdentityProfile) GetProfileConfigOk() (*ProfileConfig, bool)`

GetProfileConfigOk returns a tuple with the ProfileConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileConfig

`func (o *LocalIdentityProfile) SetProfileConfig(v ProfileConfig)`

SetProfileConfig sets ProfileConfig field to given value.

### HasProfileConfig

`func (o *LocalIdentityProfile) HasProfileConfig() bool`

HasProfileConfig returns a boolean if a field has been set.

### GetFieldConfig

`func (o *LocalIdentityProfile) GetFieldConfig() FieldConfig`

GetFieldConfig returns the FieldConfig field if non-nil, zero value otherwise.

### GetFieldConfigOk

`func (o *LocalIdentityProfile) GetFieldConfigOk() (*FieldConfig, bool)`

GetFieldConfigOk returns a tuple with the FieldConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldConfig

`func (o *LocalIdentityProfile) SetFieldConfig(v FieldConfig)`

SetFieldConfig sets FieldConfig field to given value.

### HasFieldConfig

`func (o *LocalIdentityProfile) HasFieldConfig() bool`

HasFieldConfig returns a boolean if a field has been set.

### GetEmailVerificationConfig

`func (o *LocalIdentityProfile) GetEmailVerificationConfig() EmailVerificationConfig`

GetEmailVerificationConfig returns the EmailVerificationConfig field if non-nil, zero value otherwise.

### GetEmailVerificationConfigOk

`func (o *LocalIdentityProfile) GetEmailVerificationConfigOk() (*EmailVerificationConfig, bool)`

GetEmailVerificationConfigOk returns a tuple with the EmailVerificationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerificationConfig

`func (o *LocalIdentityProfile) SetEmailVerificationConfig(v EmailVerificationConfig)`

SetEmailVerificationConfig sets EmailVerificationConfig field to given value.

### HasEmailVerificationConfig

`func (o *LocalIdentityProfile) HasEmailVerificationConfig() bool`

HasEmailVerificationConfig returns a boolean if a field has been set.

### GetDataStoreConfig

`func (o *LocalIdentityProfile) GetDataStoreConfig() DataStoreConfig`

GetDataStoreConfig returns the DataStoreConfig field if non-nil, zero value otherwise.

### GetDataStoreConfigOk

`func (o *LocalIdentityProfile) GetDataStoreConfigOk() (*DataStoreConfig, bool)`

GetDataStoreConfigOk returns a tuple with the DataStoreConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStoreConfig

`func (o *LocalIdentityProfile) SetDataStoreConfig(v DataStoreConfig)`

SetDataStoreConfig sets DataStoreConfig field to given value.

### HasDataStoreConfig

`func (o *LocalIdentityProfile) HasDataStoreConfig() bool`

HasDataStoreConfig returns a boolean if a field has been set.

### GetProfileEnabled

`func (o *LocalIdentityProfile) GetProfileEnabled() bool`

GetProfileEnabled returns the ProfileEnabled field if non-nil, zero value otherwise.

### GetProfileEnabledOk

`func (o *LocalIdentityProfile) GetProfileEnabledOk() (*bool, bool)`

GetProfileEnabledOk returns a tuple with the ProfileEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileEnabled

`func (o *LocalIdentityProfile) SetProfileEnabled(v bool)`

SetProfileEnabled sets ProfileEnabled field to given value.

### HasProfileEnabled

`func (o *LocalIdentityProfile) HasProfileEnabled() bool`

HasProfileEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


