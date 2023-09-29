# AccountManagementSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountStatusAttributeName** | **string** | The account status attribute name. | 
**AccountStatusAlgorithm** | **string** | The account status algorithm name.  ACCOUNT_STATUS_ALGORITHM_AD -  Algorithm name for Active Directory, which uses a bitmap for each user entry.  ACCOUNT_STATUS_ALGORITHM_FLAG - Algorithm name for Oracle Directory Server and other LDAP directories that use a separate attribute to store the user&#39;s status. When this option is selected, the Flag Comparison Value and Flag Comparison Status fields should be used. | 
**FlagComparisonValue** | Pointer to **string** | The flag that represents comparison value. | [optional] 
**FlagComparisonStatus** | Pointer to **bool** | The flag that represents comparison status. | [optional] 
**DefaultStatus** | Pointer to **bool** | The default status of the account. | [optional] 

## Methods

### NewAccountManagementSettings

`func NewAccountManagementSettings(accountStatusAttributeName string, accountStatusAlgorithm string, ) *AccountManagementSettings`

NewAccountManagementSettings instantiates a new AccountManagementSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountManagementSettingsWithDefaults

`func NewAccountManagementSettingsWithDefaults() *AccountManagementSettings`

NewAccountManagementSettingsWithDefaults instantiates a new AccountManagementSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountStatusAttributeName

`func (o *AccountManagementSettings) GetAccountStatusAttributeName() string`

GetAccountStatusAttributeName returns the AccountStatusAttributeName field if non-nil, zero value otherwise.

### GetAccountStatusAttributeNameOk

`func (o *AccountManagementSettings) GetAccountStatusAttributeNameOk() (*string, bool)`

GetAccountStatusAttributeNameOk returns a tuple with the AccountStatusAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStatusAttributeName

`func (o *AccountManagementSettings) SetAccountStatusAttributeName(v string)`

SetAccountStatusAttributeName sets AccountStatusAttributeName field to given value.


### GetAccountStatusAlgorithm

`func (o *AccountManagementSettings) GetAccountStatusAlgorithm() string`

GetAccountStatusAlgorithm returns the AccountStatusAlgorithm field if non-nil, zero value otherwise.

### GetAccountStatusAlgorithmOk

`func (o *AccountManagementSettings) GetAccountStatusAlgorithmOk() (*string, bool)`

GetAccountStatusAlgorithmOk returns a tuple with the AccountStatusAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStatusAlgorithm

`func (o *AccountManagementSettings) SetAccountStatusAlgorithm(v string)`

SetAccountStatusAlgorithm sets AccountStatusAlgorithm field to given value.


### GetFlagComparisonValue

`func (o *AccountManagementSettings) GetFlagComparisonValue() string`

GetFlagComparisonValue returns the FlagComparisonValue field if non-nil, zero value otherwise.

### GetFlagComparisonValueOk

`func (o *AccountManagementSettings) GetFlagComparisonValueOk() (*string, bool)`

GetFlagComparisonValueOk returns a tuple with the FlagComparisonValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagComparisonValue

`func (o *AccountManagementSettings) SetFlagComparisonValue(v string)`

SetFlagComparisonValue sets FlagComparisonValue field to given value.

### HasFlagComparisonValue

`func (o *AccountManagementSettings) HasFlagComparisonValue() bool`

HasFlagComparisonValue returns a boolean if a field has been set.

### GetFlagComparisonStatus

`func (o *AccountManagementSettings) GetFlagComparisonStatus() bool`

GetFlagComparisonStatus returns the FlagComparisonStatus field if non-nil, zero value otherwise.

### GetFlagComparisonStatusOk

`func (o *AccountManagementSettings) GetFlagComparisonStatusOk() (*bool, bool)`

GetFlagComparisonStatusOk returns a tuple with the FlagComparisonStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagComparisonStatus

`func (o *AccountManagementSettings) SetFlagComparisonStatus(v bool)`

SetFlagComparisonStatus sets FlagComparisonStatus field to given value.

### HasFlagComparisonStatus

`func (o *AccountManagementSettings) HasFlagComparisonStatus() bool`

HasFlagComparisonStatus returns a boolean if a field has been set.

### GetDefaultStatus

`func (o *AccountManagementSettings) GetDefaultStatus() bool`

GetDefaultStatus returns the DefaultStatus field if non-nil, zero value otherwise.

### GetDefaultStatusOk

`func (o *AccountManagementSettings) GetDefaultStatusOk() (*bool, bool)`

GetDefaultStatusOk returns a tuple with the DefaultStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStatus

`func (o *AccountManagementSettings) SetDefaultStatus(v bool)`

SetDefaultStatus sets DefaultStatus field to given value.

### HasDefaultStatus

`func (o *AccountManagementSettings) HasDefaultStatus() bool`

HasDefaultStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


