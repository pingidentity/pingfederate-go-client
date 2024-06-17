# JitProvisioning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserAttributes** | [**JitProvisioningUserAttributes**](JitProvisioningUserAttributes.md) |  | 
**UserRepository** | [**DataStoreRepositoryAggregation**](DataStoreRepositoryAggregation.md) |  | 
**EventTrigger** | Pointer to **string** | Specify when provisioning occurs during assertion processing. The default is &#39;NEW_USER_ONLY&#39;. | [optional] 
**ErrorHandling** | Pointer to **string** | Specify behavior when provisioning request fails. The default is &#39;CONTINUE_SSO&#39;. | [optional] 

## Methods

### NewJitProvisioning

`func NewJitProvisioning(userAttributes JitProvisioningUserAttributes, userRepository DataStoreRepositoryAggregation, ) *JitProvisioning`

NewJitProvisioning instantiates a new JitProvisioning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJitProvisioningWithDefaults

`func NewJitProvisioningWithDefaults() *JitProvisioning`

NewJitProvisioningWithDefaults instantiates a new JitProvisioning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserAttributes

`func (o *JitProvisioning) GetUserAttributes() JitProvisioningUserAttributes`

GetUserAttributes returns the UserAttributes field if non-nil, zero value otherwise.

### GetUserAttributesOk

`func (o *JitProvisioning) GetUserAttributesOk() (*JitProvisioningUserAttributes, bool)`

GetUserAttributesOk returns a tuple with the UserAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttributes

`func (o *JitProvisioning) SetUserAttributes(v JitProvisioningUserAttributes)`

SetUserAttributes sets UserAttributes field to given value.


### GetUserRepository

`func (o *JitProvisioning) GetUserRepository() DataStoreRepositoryAggregation`

GetUserRepository returns the UserRepository field if non-nil, zero value otherwise.

### GetUserRepositoryOk

`func (o *JitProvisioning) GetUserRepositoryOk() (*DataStoreRepositoryAggregation, bool)`

GetUserRepositoryOk returns a tuple with the UserRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRepository

`func (o *JitProvisioning) SetUserRepository(v DataStoreRepositoryAggregation)`

SetUserRepository sets UserRepository field to given value.


### GetEventTrigger

`func (o *JitProvisioning) GetEventTrigger() string`

GetEventTrigger returns the EventTrigger field if non-nil, zero value otherwise.

### GetEventTriggerOk

`func (o *JitProvisioning) GetEventTriggerOk() (*string, bool)`

GetEventTriggerOk returns a tuple with the EventTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTrigger

`func (o *JitProvisioning) SetEventTrigger(v string)`

SetEventTrigger sets EventTrigger field to given value.

### HasEventTrigger

`func (o *JitProvisioning) HasEventTrigger() bool`

HasEventTrigger returns a boolean if a field has been set.

### GetErrorHandling

`func (o *JitProvisioning) GetErrorHandling() string`

GetErrorHandling returns the ErrorHandling field if non-nil, zero value otherwise.

### GetErrorHandlingOk

`func (o *JitProvisioning) GetErrorHandlingOk() (*string, bool)`

GetErrorHandlingOk returns a tuple with the ErrorHandling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorHandling

`func (o *JitProvisioning) SetErrorHandling(v string)`

SetErrorHandling sets ErrorHandling field to given value.

### HasErrorHandling

`func (o *JitProvisioning) HasErrorHandling() bool`

HasErrorHandling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


