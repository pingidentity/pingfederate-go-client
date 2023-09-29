# SpWsTrust

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartnerServiceIds** | **[]string** | The partner service identifiers. | 
**OAuthAssertionProfiles** | Pointer to **bool** | When selected, four additional token-type requests become available. | [optional] 
**DefaultTokenType** | Pointer to **string** | The default token type when a web service client (WSC) does not specify in the token request which token type the STS should issue. Defaults to SAML 2.0. | [optional] 
**GenerateKey** | Pointer to **bool** | When selected, the STS generates a symmetric key to be used in conjunction with the \&quot;Holder of Key\&quot; (HoK) designation for the assertion&#39;s Subject Confirmation Method.  This option does not apply to OAuth assertion profiles. | [optional] 
**EncryptSaml2Assertion** | Pointer to **bool** | When selected, the STS encrypts the SAML 2.0 assertion. Applicable only to SAML 2.0 security token.  This option does not apply to OAuth assertion profiles. | [optional] 
**MinutesBefore** | Pointer to **int64** | The amount of time before the SAML token was issued during which it is to be considered valid. The default value is 5. | [optional] 
**MinutesAfter** | Pointer to **int64** | The amount of time after the SAML token was issued during which it is to be considered valid. The default value is 30. | [optional] 
**AttributeContract** | [**SpWsTrustAttributeContract**](SpWsTrustAttributeContract.md) |  | 
**TokenProcessorMappings** | [**[]IdpTokenProcessorMapping**](IdpTokenProcessorMapping.md) | A list of token processors to validate incoming tokens. | 
**AbortIfNotFulfilledFromRequest** | Pointer to **bool** | If the attribute contract cannot be fulfilled using data from the Request, abort the transaction. | [optional] 
**RequestContractRef** | Pointer to [**ResourceLink**](ResourceLink.md) |  | [optional] 
**MessageCustomizations** | Pointer to [**[]ProtocolMessageCustomization**](ProtocolMessageCustomization.md) | The message customizations for WS-Trust. Depending on server settings, connection type, and protocol this may or may not be supported. | [optional] 

## Methods

### NewSpWsTrust

`func NewSpWsTrust(partnerServiceIds []string, attributeContract SpWsTrustAttributeContract, tokenProcessorMappings []IdpTokenProcessorMapping, ) *SpWsTrust`

NewSpWsTrust instantiates a new SpWsTrust object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpWsTrustWithDefaults

`func NewSpWsTrustWithDefaults() *SpWsTrust`

NewSpWsTrustWithDefaults instantiates a new SpWsTrust object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartnerServiceIds

`func (o *SpWsTrust) GetPartnerServiceIds() []string`

GetPartnerServiceIds returns the PartnerServiceIds field if non-nil, zero value otherwise.

### GetPartnerServiceIdsOk

`func (o *SpWsTrust) GetPartnerServiceIdsOk() (*[]string, bool)`

GetPartnerServiceIdsOk returns a tuple with the PartnerServiceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerServiceIds

`func (o *SpWsTrust) SetPartnerServiceIds(v []string)`

SetPartnerServiceIds sets PartnerServiceIds field to given value.


### GetOAuthAssertionProfiles

`func (o *SpWsTrust) GetOAuthAssertionProfiles() bool`

GetOAuthAssertionProfiles returns the OAuthAssertionProfiles field if non-nil, zero value otherwise.

### GetOAuthAssertionProfilesOk

`func (o *SpWsTrust) GetOAuthAssertionProfilesOk() (*bool, bool)`

GetOAuthAssertionProfilesOk returns a tuple with the OAuthAssertionProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthAssertionProfiles

`func (o *SpWsTrust) SetOAuthAssertionProfiles(v bool)`

SetOAuthAssertionProfiles sets OAuthAssertionProfiles field to given value.

### HasOAuthAssertionProfiles

`func (o *SpWsTrust) HasOAuthAssertionProfiles() bool`

HasOAuthAssertionProfiles returns a boolean if a field has been set.

### GetDefaultTokenType

`func (o *SpWsTrust) GetDefaultTokenType() string`

GetDefaultTokenType returns the DefaultTokenType field if non-nil, zero value otherwise.

### GetDefaultTokenTypeOk

`func (o *SpWsTrust) GetDefaultTokenTypeOk() (*string, bool)`

GetDefaultTokenTypeOk returns a tuple with the DefaultTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTokenType

`func (o *SpWsTrust) SetDefaultTokenType(v string)`

SetDefaultTokenType sets DefaultTokenType field to given value.

### HasDefaultTokenType

`func (o *SpWsTrust) HasDefaultTokenType() bool`

HasDefaultTokenType returns a boolean if a field has been set.

### GetGenerateKey

`func (o *SpWsTrust) GetGenerateKey() bool`

GetGenerateKey returns the GenerateKey field if non-nil, zero value otherwise.

### GetGenerateKeyOk

`func (o *SpWsTrust) GetGenerateKeyOk() (*bool, bool)`

GetGenerateKeyOk returns a tuple with the GenerateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateKey

`func (o *SpWsTrust) SetGenerateKey(v bool)`

SetGenerateKey sets GenerateKey field to given value.

### HasGenerateKey

`func (o *SpWsTrust) HasGenerateKey() bool`

HasGenerateKey returns a boolean if a field has been set.

### GetEncryptSaml2Assertion

`func (o *SpWsTrust) GetEncryptSaml2Assertion() bool`

GetEncryptSaml2Assertion returns the EncryptSaml2Assertion field if non-nil, zero value otherwise.

### GetEncryptSaml2AssertionOk

`func (o *SpWsTrust) GetEncryptSaml2AssertionOk() (*bool, bool)`

GetEncryptSaml2AssertionOk returns a tuple with the EncryptSaml2Assertion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptSaml2Assertion

`func (o *SpWsTrust) SetEncryptSaml2Assertion(v bool)`

SetEncryptSaml2Assertion sets EncryptSaml2Assertion field to given value.

### HasEncryptSaml2Assertion

`func (o *SpWsTrust) HasEncryptSaml2Assertion() bool`

HasEncryptSaml2Assertion returns a boolean if a field has been set.

### GetMinutesBefore

`func (o *SpWsTrust) GetMinutesBefore() int64`

GetMinutesBefore returns the MinutesBefore field if non-nil, zero value otherwise.

### GetMinutesBeforeOk

`func (o *SpWsTrust) GetMinutesBeforeOk() (*int64, bool)`

GetMinutesBeforeOk returns a tuple with the MinutesBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutesBefore

`func (o *SpWsTrust) SetMinutesBefore(v int64)`

SetMinutesBefore sets MinutesBefore field to given value.

### HasMinutesBefore

`func (o *SpWsTrust) HasMinutesBefore() bool`

HasMinutesBefore returns a boolean if a field has been set.

### GetMinutesAfter

`func (o *SpWsTrust) GetMinutesAfter() int64`

GetMinutesAfter returns the MinutesAfter field if non-nil, zero value otherwise.

### GetMinutesAfterOk

`func (o *SpWsTrust) GetMinutesAfterOk() (*int64, bool)`

GetMinutesAfterOk returns a tuple with the MinutesAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutesAfter

`func (o *SpWsTrust) SetMinutesAfter(v int64)`

SetMinutesAfter sets MinutesAfter field to given value.

### HasMinutesAfter

`func (o *SpWsTrust) HasMinutesAfter() bool`

HasMinutesAfter returns a boolean if a field has been set.

### GetAttributeContract

`func (o *SpWsTrust) GetAttributeContract() SpWsTrustAttributeContract`

GetAttributeContract returns the AttributeContract field if non-nil, zero value otherwise.

### GetAttributeContractOk

`func (o *SpWsTrust) GetAttributeContractOk() (*SpWsTrustAttributeContract, bool)`

GetAttributeContractOk returns a tuple with the AttributeContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContract

`func (o *SpWsTrust) SetAttributeContract(v SpWsTrustAttributeContract)`

SetAttributeContract sets AttributeContract field to given value.


### GetTokenProcessorMappings

`func (o *SpWsTrust) GetTokenProcessorMappings() []IdpTokenProcessorMapping`

GetTokenProcessorMappings returns the TokenProcessorMappings field if non-nil, zero value otherwise.

### GetTokenProcessorMappingsOk

`func (o *SpWsTrust) GetTokenProcessorMappingsOk() (*[]IdpTokenProcessorMapping, bool)`

GetTokenProcessorMappingsOk returns a tuple with the TokenProcessorMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenProcessorMappings

`func (o *SpWsTrust) SetTokenProcessorMappings(v []IdpTokenProcessorMapping)`

SetTokenProcessorMappings sets TokenProcessorMappings field to given value.


### GetAbortIfNotFulfilledFromRequest

`func (o *SpWsTrust) GetAbortIfNotFulfilledFromRequest() bool`

GetAbortIfNotFulfilledFromRequest returns the AbortIfNotFulfilledFromRequest field if non-nil, zero value otherwise.

### GetAbortIfNotFulfilledFromRequestOk

`func (o *SpWsTrust) GetAbortIfNotFulfilledFromRequestOk() (*bool, bool)`

GetAbortIfNotFulfilledFromRequestOk returns a tuple with the AbortIfNotFulfilledFromRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortIfNotFulfilledFromRequest

`func (o *SpWsTrust) SetAbortIfNotFulfilledFromRequest(v bool)`

SetAbortIfNotFulfilledFromRequest sets AbortIfNotFulfilledFromRequest field to given value.

### HasAbortIfNotFulfilledFromRequest

`func (o *SpWsTrust) HasAbortIfNotFulfilledFromRequest() bool`

HasAbortIfNotFulfilledFromRequest returns a boolean if a field has been set.

### GetRequestContractRef

`func (o *SpWsTrust) GetRequestContractRef() ResourceLink`

GetRequestContractRef returns the RequestContractRef field if non-nil, zero value otherwise.

### GetRequestContractRefOk

`func (o *SpWsTrust) GetRequestContractRefOk() (*ResourceLink, bool)`

GetRequestContractRefOk returns a tuple with the RequestContractRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestContractRef

`func (o *SpWsTrust) SetRequestContractRef(v ResourceLink)`

SetRequestContractRef sets RequestContractRef field to given value.

### HasRequestContractRef

`func (o *SpWsTrust) HasRequestContractRef() bool`

HasRequestContractRef returns a boolean if a field has been set.

### GetMessageCustomizations

`func (o *SpWsTrust) GetMessageCustomizations() []ProtocolMessageCustomization`

GetMessageCustomizations returns the MessageCustomizations field if non-nil, zero value otherwise.

### GetMessageCustomizationsOk

`func (o *SpWsTrust) GetMessageCustomizationsOk() (*[]ProtocolMessageCustomization, bool)`

GetMessageCustomizationsOk returns a tuple with the MessageCustomizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCustomizations

`func (o *SpWsTrust) SetMessageCustomizations(v []ProtocolMessageCustomization)`

SetMessageCustomizations sets MessageCustomizations field to given value.

### HasMessageCustomizations

`func (o *SpWsTrust) HasMessageCustomizations() bool`

HasMessageCustomizations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


