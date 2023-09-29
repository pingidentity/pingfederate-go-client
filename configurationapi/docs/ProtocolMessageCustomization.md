# ProtocolMessageCustomization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextName** | Pointer to **string** | The context in which the customization will be applied. Depending on the connection type and protocol, this can either be &#39;assertion&#39;, &#39;authn-response&#39; or &#39;authn-request&#39;. | [optional] 
**MessageExpression** | Pointer to **string** | The OGNL expression that will be executed. Refer to the Admin Manual for a list of variables provided by PingFederate. | [optional] 

## Methods

### NewProtocolMessageCustomization

`func NewProtocolMessageCustomization() *ProtocolMessageCustomization`

NewProtocolMessageCustomization instantiates a new ProtocolMessageCustomization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtocolMessageCustomizationWithDefaults

`func NewProtocolMessageCustomizationWithDefaults() *ProtocolMessageCustomization`

NewProtocolMessageCustomizationWithDefaults instantiates a new ProtocolMessageCustomization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextName

`func (o *ProtocolMessageCustomization) GetContextName() string`

GetContextName returns the ContextName field if non-nil, zero value otherwise.

### GetContextNameOk

`func (o *ProtocolMessageCustomization) GetContextNameOk() (*string, bool)`

GetContextNameOk returns a tuple with the ContextName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextName

`func (o *ProtocolMessageCustomization) SetContextName(v string)`

SetContextName sets ContextName field to given value.

### HasContextName

`func (o *ProtocolMessageCustomization) HasContextName() bool`

HasContextName returns a boolean if a field has been set.

### GetMessageExpression

`func (o *ProtocolMessageCustomization) GetMessageExpression() string`

GetMessageExpression returns the MessageExpression field if non-nil, zero value otherwise.

### GetMessageExpressionOk

`func (o *ProtocolMessageCustomization) GetMessageExpressionOk() (*string, bool)`

GetMessageExpressionOk returns a tuple with the MessageExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageExpression

`func (o *ProtocolMessageCustomization) SetMessageExpression(v string)`

SetMessageExpression sets MessageExpression field to given value.

### HasMessageExpression

`func (o *ProtocolMessageCustomization) HasMessageExpression() bool`

HasMessageExpression returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


