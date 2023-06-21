# ChangeDetectionSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserObjectClass** | **string** | The user object class. | 
**GroupObjectClass** | **string** | The group object class. | 
**ChangedUsersAlgorithm** | **string** | The changed user algorithm.  ACTIVE_DIRECTORY_USN - For Active Directory only, this algorithm queries for update sequence numbers on user records that are larger than the last time records were checked.  TIMESTAMP - Queries for timestamps on user records that are not older than the last time records were checked. This check is more efficient from the point of view of the PingFederate provisioner but can be more time consuming on the LDAP side, particularly with the Oracle Directory Server.  TIMESTAMP_NO_NEGATION - Queries for timestamps on user records that are newer than the last time records were checked. This algorithm is recommended for the Oracle Directory Server. | 
**UsnAttributeName** | Pointer to **string** | The USN attribute name. | [optional] 
**TimeStampAttributeName** | **string** | The timestamp attribute name. | 

## Methods

### NewChangeDetectionSettings

`func NewChangeDetectionSettings(userObjectClass string, groupObjectClass string, changedUsersAlgorithm string, timeStampAttributeName string, ) *ChangeDetectionSettings`

NewChangeDetectionSettings instantiates a new ChangeDetectionSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeDetectionSettingsWithDefaults

`func NewChangeDetectionSettingsWithDefaults() *ChangeDetectionSettings`

NewChangeDetectionSettingsWithDefaults instantiates a new ChangeDetectionSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserObjectClass

`func (o *ChangeDetectionSettings) GetUserObjectClass() string`

GetUserObjectClass returns the UserObjectClass field if non-nil, zero value otherwise.

### GetUserObjectClassOk

`func (o *ChangeDetectionSettings) GetUserObjectClassOk() (*string, bool)`

GetUserObjectClassOk returns a tuple with the UserObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserObjectClass

`func (o *ChangeDetectionSettings) SetUserObjectClass(v string)`

SetUserObjectClass sets UserObjectClass field to given value.


### GetGroupObjectClass

`func (o *ChangeDetectionSettings) GetGroupObjectClass() string`

GetGroupObjectClass returns the GroupObjectClass field if non-nil, zero value otherwise.

### GetGroupObjectClassOk

`func (o *ChangeDetectionSettings) GetGroupObjectClassOk() (*string, bool)`

GetGroupObjectClassOk returns a tuple with the GroupObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupObjectClass

`func (o *ChangeDetectionSettings) SetGroupObjectClass(v string)`

SetGroupObjectClass sets GroupObjectClass field to given value.


### GetChangedUsersAlgorithm

`func (o *ChangeDetectionSettings) GetChangedUsersAlgorithm() string`

GetChangedUsersAlgorithm returns the ChangedUsersAlgorithm field if non-nil, zero value otherwise.

### GetChangedUsersAlgorithmOk

`func (o *ChangeDetectionSettings) GetChangedUsersAlgorithmOk() (*string, bool)`

GetChangedUsersAlgorithmOk returns a tuple with the ChangedUsersAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedUsersAlgorithm

`func (o *ChangeDetectionSettings) SetChangedUsersAlgorithm(v string)`

SetChangedUsersAlgorithm sets ChangedUsersAlgorithm field to given value.


### GetUsnAttributeName

`func (o *ChangeDetectionSettings) GetUsnAttributeName() string`

GetUsnAttributeName returns the UsnAttributeName field if non-nil, zero value otherwise.

### GetUsnAttributeNameOk

`func (o *ChangeDetectionSettings) GetUsnAttributeNameOk() (*string, bool)`

GetUsnAttributeNameOk returns a tuple with the UsnAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsnAttributeName

`func (o *ChangeDetectionSettings) SetUsnAttributeName(v string)`

SetUsnAttributeName sets UsnAttributeName field to given value.

### HasUsnAttributeName

`func (o *ChangeDetectionSettings) HasUsnAttributeName() bool`

HasUsnAttributeName returns a boolean if a field has been set.

### GetTimeStampAttributeName

`func (o *ChangeDetectionSettings) GetTimeStampAttributeName() string`

GetTimeStampAttributeName returns the TimeStampAttributeName field if non-nil, zero value otherwise.

### GetTimeStampAttributeNameOk

`func (o *ChangeDetectionSettings) GetTimeStampAttributeNameOk() (*string, bool)`

GetTimeStampAttributeNameOk returns a tuple with the TimeStampAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStampAttributeName

`func (o *ChangeDetectionSettings) SetTimeStampAttributeName(v string)`

SetTimeStampAttributeName sets TimeStampAttributeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


