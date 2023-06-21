# ChannelSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSource** | [**ResourceLink**](ResourceLink.md) |  | 
**GuidAttributeName** | **string** | the GUID attribute name. | 
**GuidBinary** | **bool** | Indicates whether the GUID is stored in binary format. | 
**ChangeDetectionSettings** | [**ChangeDetectionSettings**](ChangeDetectionSettings.md) |  | 
**GroupMembershipDetection** | [**GroupMembershipDetection**](GroupMembershipDetection.md) |  | 
**AccountManagementSettings** | [**AccountManagementSettings**](AccountManagementSettings.md) |  | 
**BaseDn** | **string** | The base DN where the user records are located. | 
**UserSourceLocation** | [**ChannelSourceLocation**](ChannelSourceLocation.md) |  | 
**GroupSourceLocation** | Pointer to [**ChannelSourceLocation**](ChannelSourceLocation.md) |  | [optional] 

## Methods

### NewChannelSource

`func NewChannelSource(dataSource ResourceLink, guidAttributeName string, guidBinary bool, changeDetectionSettings ChangeDetectionSettings, groupMembershipDetection GroupMembershipDetection, accountManagementSettings AccountManagementSettings, baseDn string, userSourceLocation ChannelSourceLocation, ) *ChannelSource`

NewChannelSource instantiates a new ChannelSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelSourceWithDefaults

`func NewChannelSourceWithDefaults() *ChannelSource`

NewChannelSourceWithDefaults instantiates a new ChannelSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSource

`func (o *ChannelSource) GetDataSource() ResourceLink`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *ChannelSource) GetDataSourceOk() (*ResourceLink, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *ChannelSource) SetDataSource(v ResourceLink)`

SetDataSource sets DataSource field to given value.


### GetGuidAttributeName

`func (o *ChannelSource) GetGuidAttributeName() string`

GetGuidAttributeName returns the GuidAttributeName field if non-nil, zero value otherwise.

### GetGuidAttributeNameOk

`func (o *ChannelSource) GetGuidAttributeNameOk() (*string, bool)`

GetGuidAttributeNameOk returns a tuple with the GuidAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuidAttributeName

`func (o *ChannelSource) SetGuidAttributeName(v string)`

SetGuidAttributeName sets GuidAttributeName field to given value.


### GetGuidBinary

`func (o *ChannelSource) GetGuidBinary() bool`

GetGuidBinary returns the GuidBinary field if non-nil, zero value otherwise.

### GetGuidBinaryOk

`func (o *ChannelSource) GetGuidBinaryOk() (*bool, bool)`

GetGuidBinaryOk returns a tuple with the GuidBinary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuidBinary

`func (o *ChannelSource) SetGuidBinary(v bool)`

SetGuidBinary sets GuidBinary field to given value.


### GetChangeDetectionSettings

`func (o *ChannelSource) GetChangeDetectionSettings() ChangeDetectionSettings`

GetChangeDetectionSettings returns the ChangeDetectionSettings field if non-nil, zero value otherwise.

### GetChangeDetectionSettingsOk

`func (o *ChannelSource) GetChangeDetectionSettingsOk() (*ChangeDetectionSettings, bool)`

GetChangeDetectionSettingsOk returns a tuple with the ChangeDetectionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeDetectionSettings

`func (o *ChannelSource) SetChangeDetectionSettings(v ChangeDetectionSettings)`

SetChangeDetectionSettings sets ChangeDetectionSettings field to given value.


### GetGroupMembershipDetection

`func (o *ChannelSource) GetGroupMembershipDetection() GroupMembershipDetection`

GetGroupMembershipDetection returns the GroupMembershipDetection field if non-nil, zero value otherwise.

### GetGroupMembershipDetectionOk

`func (o *ChannelSource) GetGroupMembershipDetectionOk() (*GroupMembershipDetection, bool)`

GetGroupMembershipDetectionOk returns a tuple with the GroupMembershipDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMembershipDetection

`func (o *ChannelSource) SetGroupMembershipDetection(v GroupMembershipDetection)`

SetGroupMembershipDetection sets GroupMembershipDetection field to given value.


### GetAccountManagementSettings

`func (o *ChannelSource) GetAccountManagementSettings() AccountManagementSettings`

GetAccountManagementSettings returns the AccountManagementSettings field if non-nil, zero value otherwise.

### GetAccountManagementSettingsOk

`func (o *ChannelSource) GetAccountManagementSettingsOk() (*AccountManagementSettings, bool)`

GetAccountManagementSettingsOk returns a tuple with the AccountManagementSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountManagementSettings

`func (o *ChannelSource) SetAccountManagementSettings(v AccountManagementSettings)`

SetAccountManagementSettings sets AccountManagementSettings field to given value.


### GetBaseDn

`func (o *ChannelSource) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *ChannelSource) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *ChannelSource) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.


### GetUserSourceLocation

`func (o *ChannelSource) GetUserSourceLocation() ChannelSourceLocation`

GetUserSourceLocation returns the UserSourceLocation field if non-nil, zero value otherwise.

### GetUserSourceLocationOk

`func (o *ChannelSource) GetUserSourceLocationOk() (*ChannelSourceLocation, bool)`

GetUserSourceLocationOk returns a tuple with the UserSourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSourceLocation

`func (o *ChannelSource) SetUserSourceLocation(v ChannelSourceLocation)`

SetUserSourceLocation sets UserSourceLocation field to given value.


### GetGroupSourceLocation

`func (o *ChannelSource) GetGroupSourceLocation() ChannelSourceLocation`

GetGroupSourceLocation returns the GroupSourceLocation field if non-nil, zero value otherwise.

### GetGroupSourceLocationOk

`func (o *ChannelSource) GetGroupSourceLocationOk() (*ChannelSourceLocation, bool)`

GetGroupSourceLocationOk returns a tuple with the GroupSourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSourceLocation

`func (o *ChannelSource) SetGroupSourceLocation(v ChannelSourceLocation)`

SetGroupSourceLocation sets GroupSourceLocation field to given value.

### HasGroupSourceLocation

`func (o *ChannelSource) HasGroupSourceLocation() bool`

HasGroupSourceLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


