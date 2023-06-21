# ArtifactResolverLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int64** | The priority of the endpoint. | 
**Url** | **string** | Remote party URLs that you will use to resolve/translate the artifact and get the actual protocol message | 

## Methods

### NewArtifactResolverLocation

`func NewArtifactResolverLocation(index int64, url string, ) *ArtifactResolverLocation`

NewArtifactResolverLocation instantiates a new ArtifactResolverLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactResolverLocationWithDefaults

`func NewArtifactResolverLocationWithDefaults() *ArtifactResolverLocation`

NewArtifactResolverLocationWithDefaults instantiates a new ArtifactResolverLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *ArtifactResolverLocation) GetIndex() int64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ArtifactResolverLocation) GetIndexOk() (*int64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ArtifactResolverLocation) SetIndex(v int64)`

SetIndex sets Index field to given value.


### GetUrl

`func (o *ArtifactResolverLocation) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ArtifactResolverLocation) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ArtifactResolverLocation) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


