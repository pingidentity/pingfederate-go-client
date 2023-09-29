# ArtifactSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lifetime** | **int64** | The lifetime of the artifact in seconds. | 
**ResolverLocations** | [**[]ArtifactResolverLocation**](ArtifactResolverLocation.md) | Remote party URLs that you will use to resolve/translate the artifact and get the actual protocol message | 
**SourceId** | Pointer to **string** | Source ID for SAML1.x connections | [optional] 

## Methods

### NewArtifactSettings

`func NewArtifactSettings(lifetime int64, resolverLocations []ArtifactResolverLocation, ) *ArtifactSettings`

NewArtifactSettings instantiates a new ArtifactSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactSettingsWithDefaults

`func NewArtifactSettingsWithDefaults() *ArtifactSettings`

NewArtifactSettingsWithDefaults instantiates a new ArtifactSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLifetime

`func (o *ArtifactSettings) GetLifetime() int64`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *ArtifactSettings) GetLifetimeOk() (*int64, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *ArtifactSettings) SetLifetime(v int64)`

SetLifetime sets Lifetime field to given value.


### GetResolverLocations

`func (o *ArtifactSettings) GetResolverLocations() []ArtifactResolverLocation`

GetResolverLocations returns the ResolverLocations field if non-nil, zero value otherwise.

### GetResolverLocationsOk

`func (o *ArtifactSettings) GetResolverLocationsOk() (*[]ArtifactResolverLocation, bool)`

GetResolverLocationsOk returns a tuple with the ResolverLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolverLocations

`func (o *ArtifactSettings) SetResolverLocations(v []ArtifactResolverLocation)`

SetResolverLocations sets ResolverLocations field to given value.


### GetSourceId

`func (o *ArtifactSettings) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ArtifactSettings) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ArtifactSettings) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *ArtifactSettings) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


