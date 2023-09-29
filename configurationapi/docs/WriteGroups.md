# WriteGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeFulfillment** | [**map[string]AttributeFulfillmentValue**](AttributeFulfillmentValue.md) | A list of user repository mappings from attribute names to their fulfillment values. | 

## Methods

### NewWriteGroups

`func NewWriteGroups(attributeFulfillment map[string]AttributeFulfillmentValue, ) *WriteGroups`

NewWriteGroups instantiates a new WriteGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWriteGroupsWithDefaults

`func NewWriteGroupsWithDefaults() *WriteGroups`

NewWriteGroupsWithDefaults instantiates a new WriteGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeFulfillment

`func (o *WriteGroups) GetAttributeFulfillment() map[string]AttributeFulfillmentValue`

GetAttributeFulfillment returns the AttributeFulfillment field if non-nil, zero value otherwise.

### GetAttributeFulfillmentOk

`func (o *WriteGroups) GetAttributeFulfillmentOk() (*map[string]AttributeFulfillmentValue, bool)`

GetAttributeFulfillmentOk returns a tuple with the AttributeFulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeFulfillment

`func (o *WriteGroups) SetAttributeFulfillment(v map[string]AttributeFulfillmentValue)`

SetAttributeFulfillment sets AttributeFulfillment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


