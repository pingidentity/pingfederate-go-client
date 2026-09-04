# CimdSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CimdEnabled** | **bool** | Whether Client ID Metadata Document feature is enabled. When enabled, PingFederate will apply CIMD logic on incoming OAuth requests. CIMD requests won&#39;t be processed unless an external OAuth client storage is configured. | 
**CacheLifetimeMinSec** | **int64** | Minimum CIMD metadata document cache lifetime in seconds. Acts as a floor on the client metadata server&#39;s Cache-Control: max-age. When Cache-Control: no-store is received, the effective TTL is set to this floor.(set to 0 to honour those directives up to the maximum cache life time value). Values must be between 0 and 86400. | 
**CacheLifetimeMaxSec** | **int64** | Maximum CIMD metadata document cache lifetime in seconds, regardless of Cache-Control: max-age. Must be greater than or equal to cacheLifetimeMinSec. Values must be between 1 and 86400. | 
**MaxRetry** | **int64** | Maximum number of retry attempts for transient failures when fetching a CIMD metadata document. Retried: IOException (network error, connection reset) and non-200 HTTP responses other than 3xx (redirects). Values must be between 0 and 5. | 
**RequestTimeoutSec** | **int64** | Maximum time in seconds to wait for a connection from the HTTP connection pool before failing the fetch. Values must be between 1 and 30. | 
**MaxFileSizeBytes** | **int64** | Maximum size in bytes of a retrieved CIMD metadata document. Responses exceeding this limit are rejected. Values must be between 1 and 65536. | 
**ReadTimeoutSec** | **int64** | Timeout in seconds applied to both TCP connection establishment and inter-packet idle read for CIMD metadata document fetches. Values must be between 1 and 120. | 

## Methods

### NewCimdSettings

`func NewCimdSettings(cimdEnabled bool, cacheLifetimeMinSec int64, cacheLifetimeMaxSec int64, maxRetry int64, requestTimeoutSec int64, maxFileSizeBytes int64, readTimeoutSec int64, ) *CimdSettings`

NewCimdSettings instantiates a new CimdSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCimdSettingsWithDefaults

`func NewCimdSettingsWithDefaults() *CimdSettings`

NewCimdSettingsWithDefaults instantiates a new CimdSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCimdEnabled

`func (o *CimdSettings) GetCimdEnabled() bool`

GetCimdEnabled returns the CimdEnabled field if non-nil, zero value otherwise.

### GetCimdEnabledOk

`func (o *CimdSettings) GetCimdEnabledOk() (*bool, bool)`

GetCimdEnabledOk returns a tuple with the CimdEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCimdEnabled

`func (o *CimdSettings) SetCimdEnabled(v bool)`

SetCimdEnabled sets CimdEnabled field to given value.


### GetCacheLifetimeMinSec

`func (o *CimdSettings) GetCacheLifetimeMinSec() int64`

GetCacheLifetimeMinSec returns the CacheLifetimeMinSec field if non-nil, zero value otherwise.

### GetCacheLifetimeMinSecOk

`func (o *CimdSettings) GetCacheLifetimeMinSecOk() (*int64, bool)`

GetCacheLifetimeMinSecOk returns a tuple with the CacheLifetimeMinSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheLifetimeMinSec

`func (o *CimdSettings) SetCacheLifetimeMinSec(v int64)`

SetCacheLifetimeMinSec sets CacheLifetimeMinSec field to given value.


### GetCacheLifetimeMaxSec

`func (o *CimdSettings) GetCacheLifetimeMaxSec() int64`

GetCacheLifetimeMaxSec returns the CacheLifetimeMaxSec field if non-nil, zero value otherwise.

### GetCacheLifetimeMaxSecOk

`func (o *CimdSettings) GetCacheLifetimeMaxSecOk() (*int64, bool)`

GetCacheLifetimeMaxSecOk returns a tuple with the CacheLifetimeMaxSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheLifetimeMaxSec

`func (o *CimdSettings) SetCacheLifetimeMaxSec(v int64)`

SetCacheLifetimeMaxSec sets CacheLifetimeMaxSec field to given value.


### GetMaxRetry

`func (o *CimdSettings) GetMaxRetry() int64`

GetMaxRetry returns the MaxRetry field if non-nil, zero value otherwise.

### GetMaxRetryOk

`func (o *CimdSettings) GetMaxRetryOk() (*int64, bool)`

GetMaxRetryOk returns a tuple with the MaxRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetry

`func (o *CimdSettings) SetMaxRetry(v int64)`

SetMaxRetry sets MaxRetry field to given value.


### GetRequestTimeoutSec

`func (o *CimdSettings) GetRequestTimeoutSec() int64`

GetRequestTimeoutSec returns the RequestTimeoutSec field if non-nil, zero value otherwise.

### GetRequestTimeoutSecOk

`func (o *CimdSettings) GetRequestTimeoutSecOk() (*int64, bool)`

GetRequestTimeoutSecOk returns a tuple with the RequestTimeoutSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTimeoutSec

`func (o *CimdSettings) SetRequestTimeoutSec(v int64)`

SetRequestTimeoutSec sets RequestTimeoutSec field to given value.


### GetMaxFileSizeBytes

`func (o *CimdSettings) GetMaxFileSizeBytes() int64`

GetMaxFileSizeBytes returns the MaxFileSizeBytes field if non-nil, zero value otherwise.

### GetMaxFileSizeBytesOk

`func (o *CimdSettings) GetMaxFileSizeBytesOk() (*int64, bool)`

GetMaxFileSizeBytesOk returns a tuple with the MaxFileSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFileSizeBytes

`func (o *CimdSettings) SetMaxFileSizeBytes(v int64)`

SetMaxFileSizeBytes sets MaxFileSizeBytes field to given value.


### GetReadTimeoutSec

`func (o *CimdSettings) GetReadTimeoutSec() int64`

GetReadTimeoutSec returns the ReadTimeoutSec field if non-nil, zero value otherwise.

### GetReadTimeoutSecOk

`func (o *CimdSettings) GetReadTimeoutSecOk() (*int64, bool)`

GetReadTimeoutSecOk returns a tuple with the ReadTimeoutSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadTimeoutSec

`func (o *CimdSettings) SetReadTimeoutSec(v int64)`

SetReadTimeoutSec sets ReadTimeoutSec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


