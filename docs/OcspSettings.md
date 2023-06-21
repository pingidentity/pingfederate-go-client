# OcspSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequesterAddNonce** | Pointer to **bool** | Do not allow responder to use cached responses. This setting defaults to disabled. | [optional] 
**ResponderUrl** | Pointer to **string** | Default responder URL. This URL is used if the certificate being checked does not specify an OCSP responder URL. | [optional] 
**ResponderCertReference** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**CurrentUpdateGracePeriod** | Pointer to **int64** | Current update grace period in minutes. This value defaults to \&quot;5\&quot;. | [optional] 
**NextUpdateGracePeriod** | Pointer to **int64** | Next update grace period in minutes. This value defaults to \&quot;5\&quot;. | [optional] 
**ResponseCachePeriod** | Pointer to **int64** | Response cache period in hours. This value defaults to \&quot;48\&quot;. | [optional] 
**ResponderTimeout** | Pointer to **int64** | Responder connection timeout in seconds. This value defaults to \&quot;5\&quot;. | [optional] 
**ActionOnResponderUnavailable** | Pointer to **string** | Action on responder unavailable. This value defaults to  \&quot;CONTINUE\&quot;. | [optional] 
**ActionOnStatusUnknown** | Pointer to **string** | Action on status unknown. This value defaults to  \&quot;FAIL\&quot;. | [optional] 
**ActionOnUnsuccessfulResponse** | Pointer to **string** | Action on unsuccessful response. This value defaults to  \&quot;FAIL\&quot;. | [optional] 

## Methods

### NewOcspSettings

`func NewOcspSettings() *OcspSettings`

NewOcspSettings instantiates a new OcspSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOcspSettingsWithDefaults

`func NewOcspSettingsWithDefaults() *OcspSettings`

NewOcspSettingsWithDefaults instantiates a new OcspSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequesterAddNonce

`func (o *OcspSettings) GetRequesterAddNonce() bool`

GetRequesterAddNonce returns the RequesterAddNonce field if non-nil, zero value otherwise.

### GetRequesterAddNonceOk

`func (o *OcspSettings) GetRequesterAddNonceOk() (*bool, bool)`

GetRequesterAddNonceOk returns a tuple with the RequesterAddNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterAddNonce

`func (o *OcspSettings) SetRequesterAddNonce(v bool)`

SetRequesterAddNonce sets RequesterAddNonce field to given value.

### HasRequesterAddNonce

`func (o *OcspSettings) HasRequesterAddNonce() bool`

HasRequesterAddNonce returns a boolean if a field has been set.

### GetResponderUrl

`func (o *OcspSettings) GetResponderUrl() string`

GetResponderUrl returns the ResponderUrl field if non-nil, zero value otherwise.

### GetResponderUrlOk

`func (o *OcspSettings) GetResponderUrlOk() (*string, bool)`

GetResponderUrlOk returns a tuple with the ResponderUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponderUrl

`func (o *OcspSettings) SetResponderUrl(v string)`

SetResponderUrl sets ResponderUrl field to given value.

### HasResponderUrl

`func (o *OcspSettings) HasResponderUrl() bool`

HasResponderUrl returns a boolean if a field has been set.

### GetResponderCertReference

`func (o *OcspSettings) GetResponderCertReference() ResourceLink`

GetResponderCertReference returns the ResponderCertReference field if non-nil, zero value otherwise.

### GetResponderCertReferenceOk

`func (o *OcspSettings) GetResponderCertReferenceOk() (*ResourceLink, bool)`

GetResponderCertReferenceOk returns a tuple with the ResponderCertReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponderCertReference

`func (o *OcspSettings) SetResponderCertReference(v ResourceLink)`

SetResponderCertReference sets ResponderCertReference field to given value.

### HasResponderCertReference

`func (o *OcspSettings) HasResponderCertReference() bool`

HasResponderCertReference returns a boolean if a field has been set.

### GetCurrentUpdateGracePeriod

`func (o *OcspSettings) GetCurrentUpdateGracePeriod() int64`

GetCurrentUpdateGracePeriod returns the CurrentUpdateGracePeriod field if non-nil, zero value otherwise.

### GetCurrentUpdateGracePeriodOk

`func (o *OcspSettings) GetCurrentUpdateGracePeriodOk() (*int64, bool)`

GetCurrentUpdateGracePeriodOk returns a tuple with the CurrentUpdateGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUpdateGracePeriod

`func (o *OcspSettings) SetCurrentUpdateGracePeriod(v int64)`

SetCurrentUpdateGracePeriod sets CurrentUpdateGracePeriod field to given value.

### HasCurrentUpdateGracePeriod

`func (o *OcspSettings) HasCurrentUpdateGracePeriod() bool`

HasCurrentUpdateGracePeriod returns a boolean if a field has been set.

### GetNextUpdateGracePeriod

`func (o *OcspSettings) GetNextUpdateGracePeriod() int64`

GetNextUpdateGracePeriod returns the NextUpdateGracePeriod field if non-nil, zero value otherwise.

### GetNextUpdateGracePeriodOk

`func (o *OcspSettings) GetNextUpdateGracePeriodOk() (*int64, bool)`

GetNextUpdateGracePeriodOk returns a tuple with the NextUpdateGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextUpdateGracePeriod

`func (o *OcspSettings) SetNextUpdateGracePeriod(v int64)`

SetNextUpdateGracePeriod sets NextUpdateGracePeriod field to given value.

### HasNextUpdateGracePeriod

`func (o *OcspSettings) HasNextUpdateGracePeriod() bool`

HasNextUpdateGracePeriod returns a boolean if a field has been set.

### GetResponseCachePeriod

`func (o *OcspSettings) GetResponseCachePeriod() int64`

GetResponseCachePeriod returns the ResponseCachePeriod field if non-nil, zero value otherwise.

### GetResponseCachePeriodOk

`func (o *OcspSettings) GetResponseCachePeriodOk() (*int64, bool)`

GetResponseCachePeriodOk returns a tuple with the ResponseCachePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCachePeriod

`func (o *OcspSettings) SetResponseCachePeriod(v int64)`

SetResponseCachePeriod sets ResponseCachePeriod field to given value.

### HasResponseCachePeriod

`func (o *OcspSettings) HasResponseCachePeriod() bool`

HasResponseCachePeriod returns a boolean if a field has been set.

### GetResponderTimeout

`func (o *OcspSettings) GetResponderTimeout() int64`

GetResponderTimeout returns the ResponderTimeout field if non-nil, zero value otherwise.

### GetResponderTimeoutOk

`func (o *OcspSettings) GetResponderTimeoutOk() (*int64, bool)`

GetResponderTimeoutOk returns a tuple with the ResponderTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponderTimeout

`func (o *OcspSettings) SetResponderTimeout(v int64)`

SetResponderTimeout sets ResponderTimeout field to given value.

### HasResponderTimeout

`func (o *OcspSettings) HasResponderTimeout() bool`

HasResponderTimeout returns a boolean if a field has been set.

### GetActionOnResponderUnavailable

`func (o *OcspSettings) GetActionOnResponderUnavailable() string`

GetActionOnResponderUnavailable returns the ActionOnResponderUnavailable field if non-nil, zero value otherwise.

### GetActionOnResponderUnavailableOk

`func (o *OcspSettings) GetActionOnResponderUnavailableOk() (*string, bool)`

GetActionOnResponderUnavailableOk returns a tuple with the ActionOnResponderUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionOnResponderUnavailable

`func (o *OcspSettings) SetActionOnResponderUnavailable(v string)`

SetActionOnResponderUnavailable sets ActionOnResponderUnavailable field to given value.

### HasActionOnResponderUnavailable

`func (o *OcspSettings) HasActionOnResponderUnavailable() bool`

HasActionOnResponderUnavailable returns a boolean if a field has been set.

### GetActionOnStatusUnknown

`func (o *OcspSettings) GetActionOnStatusUnknown() string`

GetActionOnStatusUnknown returns the ActionOnStatusUnknown field if non-nil, zero value otherwise.

### GetActionOnStatusUnknownOk

`func (o *OcspSettings) GetActionOnStatusUnknownOk() (*string, bool)`

GetActionOnStatusUnknownOk returns a tuple with the ActionOnStatusUnknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionOnStatusUnknown

`func (o *OcspSettings) SetActionOnStatusUnknown(v string)`

SetActionOnStatusUnknown sets ActionOnStatusUnknown field to given value.

### HasActionOnStatusUnknown

`func (o *OcspSettings) HasActionOnStatusUnknown() bool`

HasActionOnStatusUnknown returns a boolean if a field has been set.

### GetActionOnUnsuccessfulResponse

`func (o *OcspSettings) GetActionOnUnsuccessfulResponse() string`

GetActionOnUnsuccessfulResponse returns the ActionOnUnsuccessfulResponse field if non-nil, zero value otherwise.

### GetActionOnUnsuccessfulResponseOk

`func (o *OcspSettings) GetActionOnUnsuccessfulResponseOk() (*string, bool)`

GetActionOnUnsuccessfulResponseOk returns a tuple with the ActionOnUnsuccessfulResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionOnUnsuccessfulResponse

`func (o *OcspSettings) SetActionOnUnsuccessfulResponse(v string)`

SetActionOnUnsuccessfulResponse sets ActionOnUnsuccessfulResponse field to given value.

### HasActionOnUnsuccessfulResponse

`func (o *OcspSettings) HasActionOnUnsuccessfulResponse() bool`

HasActionOnUnsuccessfulResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


