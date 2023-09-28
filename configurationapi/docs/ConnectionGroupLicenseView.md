# ConnectionGroupLicenseView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Group name from the license file. | [optional] 
**ConnectionCount** | Pointer to **int64** | Maximum number of connections permitted under the group. | [optional] 
**StartDate** | Pointer to **time.Time** | Start date for the group. | [optional] 
**EndDate** | Pointer to **time.Time** | End date for the group. | [optional] 

## Methods

### NewConnectionGroupLicenseView

`func NewConnectionGroupLicenseView() *ConnectionGroupLicenseView`

NewConnectionGroupLicenseView instantiates a new ConnectionGroupLicenseView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionGroupLicenseViewWithDefaults

`func NewConnectionGroupLicenseViewWithDefaults() *ConnectionGroupLicenseView`

NewConnectionGroupLicenseViewWithDefaults instantiates a new ConnectionGroupLicenseView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConnectionGroupLicenseView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectionGroupLicenseView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectionGroupLicenseView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectionGroupLicenseView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConnectionCount

`func (o *ConnectionGroupLicenseView) GetConnectionCount() int64`

GetConnectionCount returns the ConnectionCount field if non-nil, zero value otherwise.

### GetConnectionCountOk

`func (o *ConnectionGroupLicenseView) GetConnectionCountOk() (*int64, bool)`

GetConnectionCountOk returns a tuple with the ConnectionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCount

`func (o *ConnectionGroupLicenseView) SetConnectionCount(v int64)`

SetConnectionCount sets ConnectionCount field to given value.

### HasConnectionCount

`func (o *ConnectionGroupLicenseView) HasConnectionCount() bool`

HasConnectionCount returns a boolean if a field has been set.

### GetStartDate

`func (o *ConnectionGroupLicenseView) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ConnectionGroupLicenseView) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ConnectionGroupLicenseView) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ConnectionGroupLicenseView) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ConnectionGroupLicenseView) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ConnectionGroupLicenseView) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ConnectionGroupLicenseView) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ConnectionGroupLicenseView) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


