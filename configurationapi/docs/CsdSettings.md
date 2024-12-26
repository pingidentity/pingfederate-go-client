# CsdSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodesToCollect** | Pointer to **[]int64** | The list of nodes to collect support data archives from. | [optional] 
**NodeTypeToCollect** | Pointer to **string** | The type of PingFederate nodes to collect support data archives from. | [optional] 
**TruncateLogs** | Pointer to **bool** | If set, PingFederate logs will be truncated. | [optional] 
**FileHeadCollectionKBSize** | Pointer to **int64** | The amount of data in kilobytes to collect at the beginning of truncated files. Data will not be truncated from the beginning of files if left blank. | [optional] 
**FileTailCollectionKBSize** | Pointer to **int64** | The amount of data in kilobytes to collect at the end of truncated files. Data will not be truncated from the end of files if left blank. | [optional] 
**RolledLogCount** | Pointer to **int64** | The number of rolled server log files to collect. | [optional] 
**EncryptArchive** | Pointer to **bool** | Indicates that the resulting support data archive should be encrypted. | [optional] 
**EncryptionPassphrase** | Pointer to **string** | The passphrase to use to encrypt and decrypt the support data archive. Required if encryptArchive is true. | [optional] 
**IncludeBinaryFiles** | Pointer to **bool** | If set, binary files will be included in the collected archive. | [optional] 
**CollectExpensiveData** | Pointer to **bool** | Collect data from expensive or long running processes. These processes may make the PingFederate server unresponsive for a couple of minutes. | [optional] 
**NumHeartbeatSamples** | Pointer to **int64** | Number of heartbeat samples to take. | [optional] 
**IntervalBetweenHeartbeatSamples** | Pointer to **int64** | Interval between heartbeat calls in seconds. | [optional] 
**ReportCount** | Pointer to **int64** | Number of reports generated for commands that support sampling (for example, mpstat). A value of 0 (zero) indicates that no reports will be generated for these commands | [optional] 
**ReportInterval** | Pointer to **int64** | Number of seconds between reports for commands that support sampling (for example, mpstat). | [optional] 
**Comment** | Pointer to **string** | Specify additional information about the collected data set.  This comment will be added to the generated archive as a README file. | [optional] 

## Methods

### NewCsdSettings

`func NewCsdSettings() *CsdSettings`

NewCsdSettings instantiates a new CsdSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCsdSettingsWithDefaults

`func NewCsdSettingsWithDefaults() *CsdSettings`

NewCsdSettingsWithDefaults instantiates a new CsdSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodesToCollect

`func (o *CsdSettings) GetNodesToCollect() []int64`

GetNodesToCollect returns the NodesToCollect field if non-nil, zero value otherwise.

### GetNodesToCollectOk

`func (o *CsdSettings) GetNodesToCollectOk() (*[]int64, bool)`

GetNodesToCollectOk returns a tuple with the NodesToCollect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesToCollect

`func (o *CsdSettings) SetNodesToCollect(v []int64)`

SetNodesToCollect sets NodesToCollect field to given value.

### HasNodesToCollect

`func (o *CsdSettings) HasNodesToCollect() bool`

HasNodesToCollect returns a boolean if a field has been set.

### GetNodeTypeToCollect

`func (o *CsdSettings) GetNodeTypeToCollect() string`

GetNodeTypeToCollect returns the NodeTypeToCollect field if non-nil, zero value otherwise.

### GetNodeTypeToCollectOk

`func (o *CsdSettings) GetNodeTypeToCollectOk() (*string, bool)`

GetNodeTypeToCollectOk returns a tuple with the NodeTypeToCollect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeTypeToCollect

`func (o *CsdSettings) SetNodeTypeToCollect(v string)`

SetNodeTypeToCollect sets NodeTypeToCollect field to given value.

### HasNodeTypeToCollect

`func (o *CsdSettings) HasNodeTypeToCollect() bool`

HasNodeTypeToCollect returns a boolean if a field has been set.

### GetTruncateLogs

`func (o *CsdSettings) GetTruncateLogs() bool`

GetTruncateLogs returns the TruncateLogs field if non-nil, zero value otherwise.

### GetTruncateLogsOk

`func (o *CsdSettings) GetTruncateLogsOk() (*bool, bool)`

GetTruncateLogsOk returns a tuple with the TruncateLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncateLogs

`func (o *CsdSettings) SetTruncateLogs(v bool)`

SetTruncateLogs sets TruncateLogs field to given value.

### HasTruncateLogs

`func (o *CsdSettings) HasTruncateLogs() bool`

HasTruncateLogs returns a boolean if a field has been set.

### GetFileHeadCollectionKBSize

`func (o *CsdSettings) GetFileHeadCollectionKBSize() int64`

GetFileHeadCollectionKBSize returns the FileHeadCollectionKBSize field if non-nil, zero value otherwise.

### GetFileHeadCollectionKBSizeOk

`func (o *CsdSettings) GetFileHeadCollectionKBSizeOk() (*int64, bool)`

GetFileHeadCollectionKBSizeOk returns a tuple with the FileHeadCollectionKBSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileHeadCollectionKBSize

`func (o *CsdSettings) SetFileHeadCollectionKBSize(v int64)`

SetFileHeadCollectionKBSize sets FileHeadCollectionKBSize field to given value.

### HasFileHeadCollectionKBSize

`func (o *CsdSettings) HasFileHeadCollectionKBSize() bool`

HasFileHeadCollectionKBSize returns a boolean if a field has been set.

### GetFileTailCollectionKBSize

`func (o *CsdSettings) GetFileTailCollectionKBSize() int64`

GetFileTailCollectionKBSize returns the FileTailCollectionKBSize field if non-nil, zero value otherwise.

### GetFileTailCollectionKBSizeOk

`func (o *CsdSettings) GetFileTailCollectionKBSizeOk() (*int64, bool)`

GetFileTailCollectionKBSizeOk returns a tuple with the FileTailCollectionKBSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileTailCollectionKBSize

`func (o *CsdSettings) SetFileTailCollectionKBSize(v int64)`

SetFileTailCollectionKBSize sets FileTailCollectionKBSize field to given value.

### HasFileTailCollectionKBSize

`func (o *CsdSettings) HasFileTailCollectionKBSize() bool`

HasFileTailCollectionKBSize returns a boolean if a field has been set.

### GetRolledLogCount

`func (o *CsdSettings) GetRolledLogCount() int64`

GetRolledLogCount returns the RolledLogCount field if non-nil, zero value otherwise.

### GetRolledLogCountOk

`func (o *CsdSettings) GetRolledLogCountOk() (*int64, bool)`

GetRolledLogCountOk returns a tuple with the RolledLogCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolledLogCount

`func (o *CsdSettings) SetRolledLogCount(v int64)`

SetRolledLogCount sets RolledLogCount field to given value.

### HasRolledLogCount

`func (o *CsdSettings) HasRolledLogCount() bool`

HasRolledLogCount returns a boolean if a field has been set.

### GetEncryptArchive

`func (o *CsdSettings) GetEncryptArchive() bool`

GetEncryptArchive returns the EncryptArchive field if non-nil, zero value otherwise.

### GetEncryptArchiveOk

`func (o *CsdSettings) GetEncryptArchiveOk() (*bool, bool)`

GetEncryptArchiveOk returns a tuple with the EncryptArchive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptArchive

`func (o *CsdSettings) SetEncryptArchive(v bool)`

SetEncryptArchive sets EncryptArchive field to given value.

### HasEncryptArchive

`func (o *CsdSettings) HasEncryptArchive() bool`

HasEncryptArchive returns a boolean if a field has been set.

### GetEncryptionPassphrase

`func (o *CsdSettings) GetEncryptionPassphrase() string`

GetEncryptionPassphrase returns the EncryptionPassphrase field if non-nil, zero value otherwise.

### GetEncryptionPassphraseOk

`func (o *CsdSettings) GetEncryptionPassphraseOk() (*string, bool)`

GetEncryptionPassphraseOk returns a tuple with the EncryptionPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionPassphrase

`func (o *CsdSettings) SetEncryptionPassphrase(v string)`

SetEncryptionPassphrase sets EncryptionPassphrase field to given value.

### HasEncryptionPassphrase

`func (o *CsdSettings) HasEncryptionPassphrase() bool`

HasEncryptionPassphrase returns a boolean if a field has been set.

### GetIncludeBinaryFiles

`func (o *CsdSettings) GetIncludeBinaryFiles() bool`

GetIncludeBinaryFiles returns the IncludeBinaryFiles field if non-nil, zero value otherwise.

### GetIncludeBinaryFilesOk

`func (o *CsdSettings) GetIncludeBinaryFilesOk() (*bool, bool)`

GetIncludeBinaryFilesOk returns a tuple with the IncludeBinaryFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBinaryFiles

`func (o *CsdSettings) SetIncludeBinaryFiles(v bool)`

SetIncludeBinaryFiles sets IncludeBinaryFiles field to given value.

### HasIncludeBinaryFiles

`func (o *CsdSettings) HasIncludeBinaryFiles() bool`

HasIncludeBinaryFiles returns a boolean if a field has been set.

### GetCollectExpensiveData

`func (o *CsdSettings) GetCollectExpensiveData() bool`

GetCollectExpensiveData returns the CollectExpensiveData field if non-nil, zero value otherwise.

### GetCollectExpensiveDataOk

`func (o *CsdSettings) GetCollectExpensiveDataOk() (*bool, bool)`

GetCollectExpensiveDataOk returns a tuple with the CollectExpensiveData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectExpensiveData

`func (o *CsdSettings) SetCollectExpensiveData(v bool)`

SetCollectExpensiveData sets CollectExpensiveData field to given value.

### HasCollectExpensiveData

`func (o *CsdSettings) HasCollectExpensiveData() bool`

HasCollectExpensiveData returns a boolean if a field has been set.

### GetNumHeartbeatSamples

`func (o *CsdSettings) GetNumHeartbeatSamples() int64`

GetNumHeartbeatSamples returns the NumHeartbeatSamples field if non-nil, zero value otherwise.

### GetNumHeartbeatSamplesOk

`func (o *CsdSettings) GetNumHeartbeatSamplesOk() (*int64, bool)`

GetNumHeartbeatSamplesOk returns a tuple with the NumHeartbeatSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumHeartbeatSamples

`func (o *CsdSettings) SetNumHeartbeatSamples(v int64)`

SetNumHeartbeatSamples sets NumHeartbeatSamples field to given value.

### HasNumHeartbeatSamples

`func (o *CsdSettings) HasNumHeartbeatSamples() bool`

HasNumHeartbeatSamples returns a boolean if a field has been set.

### GetIntervalBetweenHeartbeatSamples

`func (o *CsdSettings) GetIntervalBetweenHeartbeatSamples() int64`

GetIntervalBetweenHeartbeatSamples returns the IntervalBetweenHeartbeatSamples field if non-nil, zero value otherwise.

### GetIntervalBetweenHeartbeatSamplesOk

`func (o *CsdSettings) GetIntervalBetweenHeartbeatSamplesOk() (*int64, bool)`

GetIntervalBetweenHeartbeatSamplesOk returns a tuple with the IntervalBetweenHeartbeatSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalBetweenHeartbeatSamples

`func (o *CsdSettings) SetIntervalBetweenHeartbeatSamples(v int64)`

SetIntervalBetweenHeartbeatSamples sets IntervalBetweenHeartbeatSamples field to given value.

### HasIntervalBetweenHeartbeatSamples

`func (o *CsdSettings) HasIntervalBetweenHeartbeatSamples() bool`

HasIntervalBetweenHeartbeatSamples returns a boolean if a field has been set.

### GetReportCount

`func (o *CsdSettings) GetReportCount() int64`

GetReportCount returns the ReportCount field if non-nil, zero value otherwise.

### GetReportCountOk

`func (o *CsdSettings) GetReportCountOk() (*int64, bool)`

GetReportCountOk returns a tuple with the ReportCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportCount

`func (o *CsdSettings) SetReportCount(v int64)`

SetReportCount sets ReportCount field to given value.

### HasReportCount

`func (o *CsdSettings) HasReportCount() bool`

HasReportCount returns a boolean if a field has been set.

### GetReportInterval

`func (o *CsdSettings) GetReportInterval() int64`

GetReportInterval returns the ReportInterval field if non-nil, zero value otherwise.

### GetReportIntervalOk

`func (o *CsdSettings) GetReportIntervalOk() (*int64, bool)`

GetReportIntervalOk returns a tuple with the ReportInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportInterval

`func (o *CsdSettings) SetReportInterval(v int64)`

SetReportInterval sets ReportInterval field to given value.

### HasReportInterval

`func (o *CsdSettings) HasReportInterval() bool`

HasReportInterval returns a boolean if a field has been set.

### GetComment

`func (o *CsdSettings) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CsdSettings) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CsdSettings) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CsdSettings) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


