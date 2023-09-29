# ServiceAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentName** | Pointer to **string** | The display name for the component. | [optional] 
**ServiceNames** | Pointer to **[]string** | The list of PingOne services consumed by the plugin. The first service represents the primary service consumed by the plugin. | [optional] 
**Configured** | Pointer to **bool** | Indicates whether one or more instances of the plugin are configured for a given PingOne connection. | [optional] 

## Methods

### NewServiceAssociation

`func NewServiceAssociation() *ServiceAssociation`

NewServiceAssociation instantiates a new ServiceAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAssociationWithDefaults

`func NewServiceAssociationWithDefaults() *ServiceAssociation`

NewServiceAssociationWithDefaults instantiates a new ServiceAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentName

`func (o *ServiceAssociation) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *ServiceAssociation) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *ServiceAssociation) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.

### HasComponentName

`func (o *ServiceAssociation) HasComponentName() bool`

HasComponentName returns a boolean if a field has been set.

### GetServiceNames

`func (o *ServiceAssociation) GetServiceNames() []string`

GetServiceNames returns the ServiceNames field if non-nil, zero value otherwise.

### GetServiceNamesOk

`func (o *ServiceAssociation) GetServiceNamesOk() (*[]string, bool)`

GetServiceNamesOk returns a tuple with the ServiceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNames

`func (o *ServiceAssociation) SetServiceNames(v []string)`

SetServiceNames sets ServiceNames field to given value.

### HasServiceNames

`func (o *ServiceAssociation) HasServiceNames() bool`

HasServiceNames returns a boolean if a field has been set.

### GetConfigured

`func (o *ServiceAssociation) GetConfigured() bool`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *ServiceAssociation) GetConfiguredOk() (*bool, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *ServiceAssociation) SetConfigured(v bool)`

SetConfigured sets Configured field to given value.

### HasConfigured

`func (o *ServiceAssociation) HasConfigured() bool`

HasConfigured returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


