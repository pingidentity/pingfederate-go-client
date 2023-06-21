# AccessTokenManagerAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeContract** | Pointer to [**AccessTokenAttributeContract**](AccessTokenAttributeContract.md) |  | [optional] 
**SelectionSettings** | Pointer to [**AtmSelectionSettings**](AtmSelectionSettings.md) |  | [optional] 
**AccessControlSettings** | Pointer to [**AtmAccessControlSettings**](AtmAccessControlSettings.md) |  | [optional] 
**SessionValidationSettings** | Pointer to [**SessionValidationSettings**](SessionValidationSettings.md) |  | [optional] 
**SequenceNumber** | Pointer to **int64** | Number added to an access token to identify which Access Token Manager issued the token. | [optional] 

## Methods

### NewAccessTokenManagerAllOf

`func NewAccessTokenManagerAllOf() *AccessTokenManagerAllOf`

NewAccessTokenManagerAllOf instantiates a new AccessTokenManagerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenManagerAllOfWithDefaults

`func NewAccessTokenManagerAllOfWithDefaults() *AccessTokenManagerAllOf`

NewAccessTokenManagerAllOfWithDefaults instantiates a new AccessTokenManagerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeContract

`func (o *AccessTokenManagerAllOf) GetAttributeContract() AccessTokenAttributeContract`

GetAttributeContract returns the AttributeContract field if non-nil, zero value otherwise.

### GetAttributeContractOk

`func (o *AccessTokenManagerAllOf) GetAttributeContractOk() (*AccessTokenAttributeContract, bool)`

GetAttributeContractOk returns a tuple with the AttributeContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContract

`func (o *AccessTokenManagerAllOf) SetAttributeContract(v AccessTokenAttributeContract)`

SetAttributeContract sets AttributeContract field to given value.

### HasAttributeContract

`func (o *AccessTokenManagerAllOf) HasAttributeContract() bool`

HasAttributeContract returns a boolean if a field has been set.

### GetSelectionSettings

`func (o *AccessTokenManagerAllOf) GetSelectionSettings() AtmSelectionSettings`

GetSelectionSettings returns the SelectionSettings field if non-nil, zero value otherwise.

### GetSelectionSettingsOk

`func (o *AccessTokenManagerAllOf) GetSelectionSettingsOk() (*AtmSelectionSettings, bool)`

GetSelectionSettingsOk returns a tuple with the SelectionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionSettings

`func (o *AccessTokenManagerAllOf) SetSelectionSettings(v AtmSelectionSettings)`

SetSelectionSettings sets SelectionSettings field to given value.

### HasSelectionSettings

`func (o *AccessTokenManagerAllOf) HasSelectionSettings() bool`

HasSelectionSettings returns a boolean if a field has been set.

### GetAccessControlSettings

`func (o *AccessTokenManagerAllOf) GetAccessControlSettings() AtmAccessControlSettings`

GetAccessControlSettings returns the AccessControlSettings field if non-nil, zero value otherwise.

### GetAccessControlSettingsOk

`func (o *AccessTokenManagerAllOf) GetAccessControlSettingsOk() (*AtmAccessControlSettings, bool)`

GetAccessControlSettingsOk returns a tuple with the AccessControlSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlSettings

`func (o *AccessTokenManagerAllOf) SetAccessControlSettings(v AtmAccessControlSettings)`

SetAccessControlSettings sets AccessControlSettings field to given value.

### HasAccessControlSettings

`func (o *AccessTokenManagerAllOf) HasAccessControlSettings() bool`

HasAccessControlSettings returns a boolean if a field has been set.

### GetSessionValidationSettings

`func (o *AccessTokenManagerAllOf) GetSessionValidationSettings() SessionValidationSettings`

GetSessionValidationSettings returns the SessionValidationSettings field if non-nil, zero value otherwise.

### GetSessionValidationSettingsOk

`func (o *AccessTokenManagerAllOf) GetSessionValidationSettingsOk() (*SessionValidationSettings, bool)`

GetSessionValidationSettingsOk returns a tuple with the SessionValidationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionValidationSettings

`func (o *AccessTokenManagerAllOf) SetSessionValidationSettings(v SessionValidationSettings)`

SetSessionValidationSettings sets SessionValidationSettings field to given value.

### HasSessionValidationSettings

`func (o *AccessTokenManagerAllOf) HasSessionValidationSettings() bool`

HasSessionValidationSettings returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *AccessTokenManagerAllOf) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *AccessTokenManagerAllOf) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *AccessTokenManagerAllOf) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *AccessTokenManagerAllOf) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


