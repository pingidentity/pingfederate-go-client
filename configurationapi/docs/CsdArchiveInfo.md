# CsdArchiveInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeIndex** | Pointer to **string** | The node index of the PingFederate node that the archive is being collected from. | [optional] 
**ArchiveId** | Pointer to **string** | The ID of the CSD archive. | [optional] 
**Address** | Pointer to **string** | The address of the PingFederate node that the archive is being collected from. | [optional] 
**Timestamp** | Pointer to **time.Time** | The timestamp of when the collection of the archive started. | [optional] 
**Status** | Pointer to **string** | The status of the archive. | [optional] 
**StatusLink** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**ExportLink** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 

## Methods

### NewCsdArchiveInfo

`func NewCsdArchiveInfo() *CsdArchiveInfo`

NewCsdArchiveInfo instantiates a new CsdArchiveInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCsdArchiveInfoWithDefaults

`func NewCsdArchiveInfoWithDefaults() *CsdArchiveInfo`

NewCsdArchiveInfoWithDefaults instantiates a new CsdArchiveInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeIndex

`func (o *CsdArchiveInfo) GetNodeIndex() string`

GetNodeIndex returns the NodeIndex field if non-nil, zero value otherwise.

### GetNodeIndexOk

`func (o *CsdArchiveInfo) GetNodeIndexOk() (*string, bool)`

GetNodeIndexOk returns a tuple with the NodeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIndex

`func (o *CsdArchiveInfo) SetNodeIndex(v string)`

SetNodeIndex sets NodeIndex field to given value.

### HasNodeIndex

`func (o *CsdArchiveInfo) HasNodeIndex() bool`

HasNodeIndex returns a boolean if a field has been set.

### GetArchiveId

`func (o *CsdArchiveInfo) GetArchiveId() string`

GetArchiveId returns the ArchiveId field if non-nil, zero value otherwise.

### GetArchiveIdOk

`func (o *CsdArchiveInfo) GetArchiveIdOk() (*string, bool)`

GetArchiveIdOk returns a tuple with the ArchiveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveId

`func (o *CsdArchiveInfo) SetArchiveId(v string)`

SetArchiveId sets ArchiveId field to given value.

### HasArchiveId

`func (o *CsdArchiveInfo) HasArchiveId() bool`

HasArchiveId returns a boolean if a field has been set.

### GetAddress

`func (o *CsdArchiveInfo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CsdArchiveInfo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CsdArchiveInfo) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CsdArchiveInfo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetTimestamp

`func (o *CsdArchiveInfo) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CsdArchiveInfo) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CsdArchiveInfo) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CsdArchiveInfo) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetStatus

`func (o *CsdArchiveInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CsdArchiveInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CsdArchiveInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CsdArchiveInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusLink

`func (o *CsdArchiveInfo) GetStatusLink() ResourceLink`

GetStatusLink returns the StatusLink field if non-nil, zero value otherwise.

### GetStatusLinkOk

`func (o *CsdArchiveInfo) GetStatusLinkOk() (*ResourceLink, bool)`

GetStatusLinkOk returns a tuple with the StatusLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusLink

`func (o *CsdArchiveInfo) SetStatusLink(v ResourceLink)`

SetStatusLink sets StatusLink field to given value.

### HasStatusLink

`func (o *CsdArchiveInfo) HasStatusLink() bool`

HasStatusLink returns a boolean if a field has been set.

### GetExportLink

`func (o *CsdArchiveInfo) GetExportLink() ResourceLink`

GetExportLink returns the ExportLink field if non-nil, zero value otherwise.

### GetExportLinkOk

`func (o *CsdArchiveInfo) GetExportLinkOk() (*ResourceLink, bool)`

GetExportLinkOk returns a tuple with the ExportLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportLink

`func (o *CsdArchiveInfo) SetExportLink(v ResourceLink)`

SetExportLink sets ExportLink field to given value.

### HasExportLink

`func (o *CsdArchiveInfo) HasExportLink() bool`

HasExportLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


