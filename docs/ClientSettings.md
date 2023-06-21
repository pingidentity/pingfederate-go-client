# ClientSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientMetadata** | Pointer to [**[]ClientMetadata**](ClientMetadata.md) | The client metadata. | [optional] 
**DynamicClientRegistration** | Pointer to [**DynamicClientRegistration**](DynamicClientRegistration.md) |  | [optional] 

## Methods

### NewClientSettings

`func NewClientSettings() *ClientSettings`

NewClientSettings instantiates a new ClientSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientSettingsWithDefaults

`func NewClientSettingsWithDefaults() *ClientSettings`

NewClientSettingsWithDefaults instantiates a new ClientSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientMetadata

`func (o *ClientSettings) GetClientMetadata() []ClientMetadata`

GetClientMetadata returns the ClientMetadata field if non-nil, zero value otherwise.

### GetClientMetadataOk

`func (o *ClientSettings) GetClientMetadataOk() (*[]ClientMetadata, bool)`

GetClientMetadataOk returns a tuple with the ClientMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMetadata

`func (o *ClientSettings) SetClientMetadata(v []ClientMetadata)`

SetClientMetadata sets ClientMetadata field to given value.

### HasClientMetadata

`func (o *ClientSettings) HasClientMetadata() bool`

HasClientMetadata returns a boolean if a field has been set.

### GetDynamicClientRegistration

`func (o *ClientSettings) GetDynamicClientRegistration() DynamicClientRegistration`

GetDynamicClientRegistration returns the DynamicClientRegistration field if non-nil, zero value otherwise.

### GetDynamicClientRegistrationOk

`func (o *ClientSettings) GetDynamicClientRegistrationOk() (*DynamicClientRegistration, bool)`

GetDynamicClientRegistrationOk returns a tuple with the DynamicClientRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicClientRegistration

`func (o *ClientSettings) SetDynamicClientRegistration(v DynamicClientRegistration)`

SetDynamicClientRegistration sets DynamicClientRegistration field to given value.

### HasDynamicClientRegistration

`func (o *ClientSettings) HasDynamicClientRegistration() bool`

HasDynamicClientRegistration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


