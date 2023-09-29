# OIDCRequestParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Request parameter name. | 
**AttributeValue** | [**AttributeFulfillmentValue**](AttributeFulfillmentValue.md) |  | 
**Value** | Pointer to **string** | A request parameter value. A parameter can have either a value or a attribute value but not both. Value set here will be converted to an attribute value of source type TEXT. An empty value will be converted to attribute value of source type NO_MAPPING. | [optional] 
**ApplicationEndpointOverride** | **bool** | Indicates whether the parameter value can be overridden by an Application Endpoint parameter | 

## Methods

### NewOIDCRequestParameter

`func NewOIDCRequestParameter(name string, attributeValue AttributeFulfillmentValue, applicationEndpointOverride bool, ) *OIDCRequestParameter`

NewOIDCRequestParameter instantiates a new OIDCRequestParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOIDCRequestParameterWithDefaults

`func NewOIDCRequestParameterWithDefaults() *OIDCRequestParameter`

NewOIDCRequestParameterWithDefaults instantiates a new OIDCRequestParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OIDCRequestParameter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OIDCRequestParameter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OIDCRequestParameter) SetName(v string)`

SetName sets Name field to given value.


### GetAttributeValue

`func (o *OIDCRequestParameter) GetAttributeValue() AttributeFulfillmentValue`

GetAttributeValue returns the AttributeValue field if non-nil, zero value otherwise.

### GetAttributeValueOk

`func (o *OIDCRequestParameter) GetAttributeValueOk() (*AttributeFulfillmentValue, bool)`

GetAttributeValueOk returns a tuple with the AttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeValue

`func (o *OIDCRequestParameter) SetAttributeValue(v AttributeFulfillmentValue)`

SetAttributeValue sets AttributeValue field to given value.


### GetValue

`func (o *OIDCRequestParameter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OIDCRequestParameter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OIDCRequestParameter) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *OIDCRequestParameter) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetApplicationEndpointOverride

`func (o *OIDCRequestParameter) GetApplicationEndpointOverride() bool`

GetApplicationEndpointOverride returns the ApplicationEndpointOverride field if non-nil, zero value otherwise.

### GetApplicationEndpointOverrideOk

`func (o *OIDCRequestParameter) GetApplicationEndpointOverrideOk() (*bool, bool)`

GetApplicationEndpointOverrideOk returns a tuple with the ApplicationEndpointOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationEndpointOverride

`func (o *OIDCRequestParameter) SetApplicationEndpointOverride(v bool)`

SetApplicationEndpointOverride sets ApplicationEndpointOverride field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


