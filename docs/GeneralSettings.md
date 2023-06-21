# GeneralSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisableAutomaticConnectionValidation** | Pointer to **bool** | Boolean that disables automatic connection validation when set to true. The default is false. | [optional] 
**IdpConnectionTransactionLoggingOverride** | Pointer to **string** | Determines the level of transaction logging for all identity provider connections. The default is DONT_OVERRIDE, in which case the logging level will be determined by each individual IdP connection | [optional] 
**SpConnectionTransactionLoggingOverride** | Pointer to **string** | Determines the level of transaction logging for all service provider connections. The default is DONT_OVERRIDE, in which case the logging level will be determined by each individual SP connection | [optional] 
**DatastoreValidationIntervalSecs** | Pointer to **int64** | Determines how long (in seconds) the result of testing a datastore connection is cached. The default is 300. | [optional] 
**RequestHeaderForCorrelationId** | Pointer to **string** | HTTP request header for retrieving correlation ID. | [optional] 

## Methods

### NewGeneralSettings

`func NewGeneralSettings() *GeneralSettings`

NewGeneralSettings instantiates a new GeneralSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneralSettingsWithDefaults

`func NewGeneralSettingsWithDefaults() *GeneralSettings`

NewGeneralSettingsWithDefaults instantiates a new GeneralSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisableAutomaticConnectionValidation

`func (o *GeneralSettings) GetDisableAutomaticConnectionValidation() bool`

GetDisableAutomaticConnectionValidation returns the DisableAutomaticConnectionValidation field if non-nil, zero value otherwise.

### GetDisableAutomaticConnectionValidationOk

`func (o *GeneralSettings) GetDisableAutomaticConnectionValidationOk() (*bool, bool)`

GetDisableAutomaticConnectionValidationOk returns a tuple with the DisableAutomaticConnectionValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAutomaticConnectionValidation

`func (o *GeneralSettings) SetDisableAutomaticConnectionValidation(v bool)`

SetDisableAutomaticConnectionValidation sets DisableAutomaticConnectionValidation field to given value.

### HasDisableAutomaticConnectionValidation

`func (o *GeneralSettings) HasDisableAutomaticConnectionValidation() bool`

HasDisableAutomaticConnectionValidation returns a boolean if a field has been set.

### GetIdpConnectionTransactionLoggingOverride

`func (o *GeneralSettings) GetIdpConnectionTransactionLoggingOverride() string`

GetIdpConnectionTransactionLoggingOverride returns the IdpConnectionTransactionLoggingOverride field if non-nil, zero value otherwise.

### GetIdpConnectionTransactionLoggingOverrideOk

`func (o *GeneralSettings) GetIdpConnectionTransactionLoggingOverrideOk() (*string, bool)`

GetIdpConnectionTransactionLoggingOverrideOk returns a tuple with the IdpConnectionTransactionLoggingOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpConnectionTransactionLoggingOverride

`func (o *GeneralSettings) SetIdpConnectionTransactionLoggingOverride(v string)`

SetIdpConnectionTransactionLoggingOverride sets IdpConnectionTransactionLoggingOverride field to given value.

### HasIdpConnectionTransactionLoggingOverride

`func (o *GeneralSettings) HasIdpConnectionTransactionLoggingOverride() bool`

HasIdpConnectionTransactionLoggingOverride returns a boolean if a field has been set.

### GetSpConnectionTransactionLoggingOverride

`func (o *GeneralSettings) GetSpConnectionTransactionLoggingOverride() string`

GetSpConnectionTransactionLoggingOverride returns the SpConnectionTransactionLoggingOverride field if non-nil, zero value otherwise.

### GetSpConnectionTransactionLoggingOverrideOk

`func (o *GeneralSettings) GetSpConnectionTransactionLoggingOverrideOk() (*string, bool)`

GetSpConnectionTransactionLoggingOverrideOk returns a tuple with the SpConnectionTransactionLoggingOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpConnectionTransactionLoggingOverride

`func (o *GeneralSettings) SetSpConnectionTransactionLoggingOverride(v string)`

SetSpConnectionTransactionLoggingOverride sets SpConnectionTransactionLoggingOverride field to given value.

### HasSpConnectionTransactionLoggingOverride

`func (o *GeneralSettings) HasSpConnectionTransactionLoggingOverride() bool`

HasSpConnectionTransactionLoggingOverride returns a boolean if a field has been set.

### GetDatastoreValidationIntervalSecs

`func (o *GeneralSettings) GetDatastoreValidationIntervalSecs() int64`

GetDatastoreValidationIntervalSecs returns the DatastoreValidationIntervalSecs field if non-nil, zero value otherwise.

### GetDatastoreValidationIntervalSecsOk

`func (o *GeneralSettings) GetDatastoreValidationIntervalSecsOk() (*int64, bool)`

GetDatastoreValidationIntervalSecsOk returns a tuple with the DatastoreValidationIntervalSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreValidationIntervalSecs

`func (o *GeneralSettings) SetDatastoreValidationIntervalSecs(v int64)`

SetDatastoreValidationIntervalSecs sets DatastoreValidationIntervalSecs field to given value.

### HasDatastoreValidationIntervalSecs

`func (o *GeneralSettings) HasDatastoreValidationIntervalSecs() bool`

HasDatastoreValidationIntervalSecs returns a boolean if a field has been set.

### GetRequestHeaderForCorrelationId

`func (o *GeneralSettings) GetRequestHeaderForCorrelationId() string`

GetRequestHeaderForCorrelationId returns the RequestHeaderForCorrelationId field if non-nil, zero value otherwise.

### GetRequestHeaderForCorrelationIdOk

`func (o *GeneralSettings) GetRequestHeaderForCorrelationIdOk() (*string, bool)`

GetRequestHeaderForCorrelationIdOk returns a tuple with the RequestHeaderForCorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaderForCorrelationId

`func (o *GeneralSettings) SetRequestHeaderForCorrelationId(v string)`

SetRequestHeaderForCorrelationId sets RequestHeaderForCorrelationId field to given value.

### HasRequestHeaderForCorrelationId

`func (o *GeneralSettings) HasRequestHeaderForCorrelationId() bool`

HasRequestHeaderForCorrelationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


