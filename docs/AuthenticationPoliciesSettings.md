# AuthenticationPoliciesSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableIdpAuthnSelection** | Pointer to **bool** | Enable IdP authentication policies. | [optional] 
**EnableSpAuthnSelection** | Pointer to **bool** | Enable SP authentication policies. | [optional] 

## Methods

### NewAuthenticationPoliciesSettings

`func NewAuthenticationPoliciesSettings() *AuthenticationPoliciesSettings`

NewAuthenticationPoliciesSettings instantiates a new AuthenticationPoliciesSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationPoliciesSettingsWithDefaults

`func NewAuthenticationPoliciesSettingsWithDefaults() *AuthenticationPoliciesSettings`

NewAuthenticationPoliciesSettingsWithDefaults instantiates a new AuthenticationPoliciesSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableIdpAuthnSelection

`func (o *AuthenticationPoliciesSettings) GetEnableIdpAuthnSelection() bool`

GetEnableIdpAuthnSelection returns the EnableIdpAuthnSelection field if non-nil, zero value otherwise.

### GetEnableIdpAuthnSelectionOk

`func (o *AuthenticationPoliciesSettings) GetEnableIdpAuthnSelectionOk() (*bool, bool)`

GetEnableIdpAuthnSelectionOk returns a tuple with the EnableIdpAuthnSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIdpAuthnSelection

`func (o *AuthenticationPoliciesSettings) SetEnableIdpAuthnSelection(v bool)`

SetEnableIdpAuthnSelection sets EnableIdpAuthnSelection field to given value.

### HasEnableIdpAuthnSelection

`func (o *AuthenticationPoliciesSettings) HasEnableIdpAuthnSelection() bool`

HasEnableIdpAuthnSelection returns a boolean if a field has been set.

### GetEnableSpAuthnSelection

`func (o *AuthenticationPoliciesSettings) GetEnableSpAuthnSelection() bool`

GetEnableSpAuthnSelection returns the EnableSpAuthnSelection field if non-nil, zero value otherwise.

### GetEnableSpAuthnSelectionOk

`func (o *AuthenticationPoliciesSettings) GetEnableSpAuthnSelectionOk() (*bool, bool)`

GetEnableSpAuthnSelectionOk returns a tuple with the EnableSpAuthnSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSpAuthnSelection

`func (o *AuthenticationPoliciesSettings) SetEnableSpAuthnSelection(v bool)`

SetEnableSpAuthnSelection sets EnableSpAuthnSelection field to given value.

### HasEnableSpAuthnSelection

`func (o *AuthenticationPoliciesSettings) HasEnableSpAuthnSelection() bool`

HasEnableSpAuthnSelection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


