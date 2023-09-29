# CrlSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TreatNonRetrievableCrlAsRevoked** | Pointer to **bool** | Treat non retrievable CRL as revoked. This setting defaults to disabled. | [optional] 
**VerifyCrlSignature** | Pointer to **bool** | Verify CRL signature. This setting defaults to enabled. | [optional] 
**NextRetryMinsWhenResolveFailed** | Pointer to **int64** | Next retry on resolution failure in minutes. This value defaults to \&quot;1440\&quot;. | [optional] 
**NextRetryMinsWhenNextUpdateInPast** | Pointer to **int64** | Next retry on next update expiration in minutes. This value defaults to \&quot;60\&quot;. | [optional] 

## Methods

### NewCrlSettings

`func NewCrlSettings() *CrlSettings`

NewCrlSettings instantiates a new CrlSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrlSettingsWithDefaults

`func NewCrlSettingsWithDefaults() *CrlSettings`

NewCrlSettingsWithDefaults instantiates a new CrlSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTreatNonRetrievableCrlAsRevoked

`func (o *CrlSettings) GetTreatNonRetrievableCrlAsRevoked() bool`

GetTreatNonRetrievableCrlAsRevoked returns the TreatNonRetrievableCrlAsRevoked field if non-nil, zero value otherwise.

### GetTreatNonRetrievableCrlAsRevokedOk

`func (o *CrlSettings) GetTreatNonRetrievableCrlAsRevokedOk() (*bool, bool)`

GetTreatNonRetrievableCrlAsRevokedOk returns a tuple with the TreatNonRetrievableCrlAsRevoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreatNonRetrievableCrlAsRevoked

`func (o *CrlSettings) SetTreatNonRetrievableCrlAsRevoked(v bool)`

SetTreatNonRetrievableCrlAsRevoked sets TreatNonRetrievableCrlAsRevoked field to given value.

### HasTreatNonRetrievableCrlAsRevoked

`func (o *CrlSettings) HasTreatNonRetrievableCrlAsRevoked() bool`

HasTreatNonRetrievableCrlAsRevoked returns a boolean if a field has been set.

### GetVerifyCrlSignature

`func (o *CrlSettings) GetVerifyCrlSignature() bool`

GetVerifyCrlSignature returns the VerifyCrlSignature field if non-nil, zero value otherwise.

### GetVerifyCrlSignatureOk

`func (o *CrlSettings) GetVerifyCrlSignatureOk() (*bool, bool)`

GetVerifyCrlSignatureOk returns a tuple with the VerifyCrlSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCrlSignature

`func (o *CrlSettings) SetVerifyCrlSignature(v bool)`

SetVerifyCrlSignature sets VerifyCrlSignature field to given value.

### HasVerifyCrlSignature

`func (o *CrlSettings) HasVerifyCrlSignature() bool`

HasVerifyCrlSignature returns a boolean if a field has been set.

### GetNextRetryMinsWhenResolveFailed

`func (o *CrlSettings) GetNextRetryMinsWhenResolveFailed() int64`

GetNextRetryMinsWhenResolveFailed returns the NextRetryMinsWhenResolveFailed field if non-nil, zero value otherwise.

### GetNextRetryMinsWhenResolveFailedOk

`func (o *CrlSettings) GetNextRetryMinsWhenResolveFailedOk() (*int64, bool)`

GetNextRetryMinsWhenResolveFailedOk returns a tuple with the NextRetryMinsWhenResolveFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRetryMinsWhenResolveFailed

`func (o *CrlSettings) SetNextRetryMinsWhenResolveFailed(v int64)`

SetNextRetryMinsWhenResolveFailed sets NextRetryMinsWhenResolveFailed field to given value.

### HasNextRetryMinsWhenResolveFailed

`func (o *CrlSettings) HasNextRetryMinsWhenResolveFailed() bool`

HasNextRetryMinsWhenResolveFailed returns a boolean if a field has been set.

### GetNextRetryMinsWhenNextUpdateInPast

`func (o *CrlSettings) GetNextRetryMinsWhenNextUpdateInPast() int64`

GetNextRetryMinsWhenNextUpdateInPast returns the NextRetryMinsWhenNextUpdateInPast field if non-nil, zero value otherwise.

### GetNextRetryMinsWhenNextUpdateInPastOk

`func (o *CrlSettings) GetNextRetryMinsWhenNextUpdateInPastOk() (*int64, bool)`

GetNextRetryMinsWhenNextUpdateInPastOk returns a tuple with the NextRetryMinsWhenNextUpdateInPast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRetryMinsWhenNextUpdateInPast

`func (o *CrlSettings) SetNextRetryMinsWhenNextUpdateInPast(v int64)`

SetNextRetryMinsWhenNextUpdateInPast sets NextRetryMinsWhenNextUpdateInPast field to given value.

### HasNextRetryMinsWhenNextUpdateInPast

`func (o *CrlSettings) HasNextRetryMinsWhenNextUpdateInPast() bool`

HasNextRetryMinsWhenNextUpdateInPast returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


