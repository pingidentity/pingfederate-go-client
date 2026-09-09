# KerberosRealmsSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceTcp** | Pointer to **bool** | Reference to the default security. | [optional] 
**KdcRetries** | **string** | Reference to the default Key Distribution Center Retries. | 
**DebugLogOutput** | Pointer to **bool** | Reference to the default logging. | [optional] 
**KdcTimeout** | **string** | Reference to the default Key Distribution Center Timeout (in seconds). | 
**KeySetRetentionPeriodMins** | Pointer to **int64** | The key set retention period in minutes. When &#39;retainPreviousKeysOnPasswordChange&#39; is set to true for a realm, this setting determines how long keys will be retained after a password change occurs. If this field is omitted in a PUT request, the default of 610 minutes is applied. | [optional] 

## Methods

### NewKerberosRealmsSettings

`func NewKerberosRealmsSettings(kdcRetries string, kdcTimeout string, ) *KerberosRealmsSettings`

NewKerberosRealmsSettings instantiates a new KerberosRealmsSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKerberosRealmsSettingsWithDefaults

`func NewKerberosRealmsSettingsWithDefaults() *KerberosRealmsSettings`

NewKerberosRealmsSettingsWithDefaults instantiates a new KerberosRealmsSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForceTcp

`func (o *KerberosRealmsSettings) GetForceTcp() bool`

GetForceTcp returns the ForceTcp field if non-nil, zero value otherwise.

### GetForceTcpOk

`func (o *KerberosRealmsSettings) GetForceTcpOk() (*bool, bool)`

GetForceTcpOk returns a tuple with the ForceTcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceTcp

`func (o *KerberosRealmsSettings) SetForceTcp(v bool)`

SetForceTcp sets ForceTcp field to given value.

### HasForceTcp

`func (o *KerberosRealmsSettings) HasForceTcp() bool`

HasForceTcp returns a boolean if a field has been set.

### GetKdcRetries

`func (o *KerberosRealmsSettings) GetKdcRetries() string`

GetKdcRetries returns the KdcRetries field if non-nil, zero value otherwise.

### GetKdcRetriesOk

`func (o *KerberosRealmsSettings) GetKdcRetriesOk() (*string, bool)`

GetKdcRetriesOk returns a tuple with the KdcRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKdcRetries

`func (o *KerberosRealmsSettings) SetKdcRetries(v string)`

SetKdcRetries sets KdcRetries field to given value.


### GetDebugLogOutput

`func (o *KerberosRealmsSettings) GetDebugLogOutput() bool`

GetDebugLogOutput returns the DebugLogOutput field if non-nil, zero value otherwise.

### GetDebugLogOutputOk

`func (o *KerberosRealmsSettings) GetDebugLogOutputOk() (*bool, bool)`

GetDebugLogOutputOk returns a tuple with the DebugLogOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugLogOutput

`func (o *KerberosRealmsSettings) SetDebugLogOutput(v bool)`

SetDebugLogOutput sets DebugLogOutput field to given value.

### HasDebugLogOutput

`func (o *KerberosRealmsSettings) HasDebugLogOutput() bool`

HasDebugLogOutput returns a boolean if a field has been set.

### GetKdcTimeout

`func (o *KerberosRealmsSettings) GetKdcTimeout() string`

GetKdcTimeout returns the KdcTimeout field if non-nil, zero value otherwise.

### GetKdcTimeoutOk

`func (o *KerberosRealmsSettings) GetKdcTimeoutOk() (*string, bool)`

GetKdcTimeoutOk returns a tuple with the KdcTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKdcTimeout

`func (o *KerberosRealmsSettings) SetKdcTimeout(v string)`

SetKdcTimeout sets KdcTimeout field to given value.


### GetKeySetRetentionPeriodMins

`func (o *KerberosRealmsSettings) GetKeySetRetentionPeriodMins() int64`

GetKeySetRetentionPeriodMins returns the KeySetRetentionPeriodMins field if non-nil, zero value otherwise.

### GetKeySetRetentionPeriodMinsOk

`func (o *KerberosRealmsSettings) GetKeySetRetentionPeriodMinsOk() (*int64, bool)`

GetKeySetRetentionPeriodMinsOk returns a tuple with the KeySetRetentionPeriodMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySetRetentionPeriodMins

`func (o *KerberosRealmsSettings) SetKeySetRetentionPeriodMins(v int64)`

SetKeySetRetentionPeriodMins sets KeySetRetentionPeriodMins field to given value.

### HasKeySetRetentionPeriodMins

`func (o *KerberosRealmsSettings) HasKeySetRetentionPeriodMins() bool`

HasKeySetRetentionPeriodMins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


