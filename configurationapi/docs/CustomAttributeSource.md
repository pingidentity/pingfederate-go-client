# CustomAttributeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The data store type of this attribute source. | 
**DataStoreRef** | [**ResourceLink**](ResourceLink.md) |  | 
**Id** | Pointer to **string** | The ID that defines this attribute source. Only alphanumeric characters allowed.&lt;br&gt;Note: Required for OpenID Connect policy attribute sources, OAuth IdP adapter mappings, OAuth access token mappings and APC-to-SP Adapter Mappings. IdP Connections will ignore this property since it only allows one attribute source to be defined per mapping. IdP-to-SP Adapter Mappings can contain multiple attribute sources. | [optional] 
**Description** | Pointer to **string** | The description of this attribute source. The description needs to be unique amongst the attribute sources for the mapping.&lt;br&gt;Note: Required for APC-to-SP Adapter Mappings | [optional] 
**AttributeContractFulfillment** | Pointer to [**map[string]AttributeFulfillmentValue**](AttributeFulfillmentValue.md) | A list of mappings from attribute names to their fulfillment values. This field is only valid for the SP Connection&#39;s Browser SSO mappings | [optional] 
**FilterFields** | Pointer to [**[]FieldEntry**](FieldEntry.md) | The list of fields that can be used to filter a request to the custom data store. | [optional] 

## Methods

### NewCustomAttributeSource

`func NewCustomAttributeSource(type_ string, dataStoreRef ResourceLink, ) *CustomAttributeSource`

NewCustomAttributeSource instantiates a new CustomAttributeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAttributeSourceWithDefaults

`func NewCustomAttributeSourceWithDefaults() *CustomAttributeSource`

NewCustomAttributeSourceWithDefaults instantiates a new CustomAttributeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomAttributeSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomAttributeSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomAttributeSource) SetType(v string)`

SetType sets Type field to given value.


### GetDataStoreRef

`func (o *CustomAttributeSource) GetDataStoreRef() ResourceLink`

GetDataStoreRef returns the DataStoreRef field if non-nil, zero value otherwise.

### GetDataStoreRefOk

`func (o *CustomAttributeSource) GetDataStoreRefOk() (*ResourceLink, bool)`

GetDataStoreRefOk returns a tuple with the DataStoreRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStoreRef

`func (o *CustomAttributeSource) SetDataStoreRef(v ResourceLink)`

SetDataStoreRef sets DataStoreRef field to given value.


### GetId

`func (o *CustomAttributeSource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomAttributeSource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomAttributeSource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomAttributeSource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *CustomAttributeSource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomAttributeSource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomAttributeSource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomAttributeSource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAttributeContractFulfillment

`func (o *CustomAttributeSource) GetAttributeContractFulfillment() map[string]AttributeFulfillmentValue`

GetAttributeContractFulfillment returns the AttributeContractFulfillment field if non-nil, zero value otherwise.

### GetAttributeContractFulfillmentOk

`func (o *CustomAttributeSource) GetAttributeContractFulfillmentOk() (*map[string]AttributeFulfillmentValue, bool)`

GetAttributeContractFulfillmentOk returns a tuple with the AttributeContractFulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeContractFulfillment

`func (o *CustomAttributeSource) SetAttributeContractFulfillment(v map[string]AttributeFulfillmentValue)`

SetAttributeContractFulfillment sets AttributeContractFulfillment field to given value.

### HasAttributeContractFulfillment

`func (o *CustomAttributeSource) HasAttributeContractFulfillment() bool`

HasAttributeContractFulfillment returns a boolean if a field has been set.

### GetFilterFields

`func (o *CustomAttributeSource) GetFilterFields() []FieldEntry`

GetFilterFields returns the FilterFields field if non-nil, zero value otherwise.

### GetFilterFieldsOk

`func (o *CustomAttributeSource) GetFilterFieldsOk() (*[]FieldEntry, bool)`

GetFilterFieldsOk returns a tuple with the FilterFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterFields

`func (o *CustomAttributeSource) SetFilterFields(v []FieldEntry)`

SetFilterFields sets FilterFields field to given value.

### HasFilterFields

`func (o *CustomAttributeSource) HasFilterFields() bool`

HasFilterFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


