# ArtifactsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ResearchArtifact**](ResearchArtifact.md) | Data are the artifacts, newest first. Content bytes are never returned here. | [optional] 
**Total** | Pointer to **int32** | Total is len(data). | [optional] 

## Methods

### NewArtifactsOut

`func NewArtifactsOut() *ArtifactsOut`

NewArtifactsOut instantiates a new ArtifactsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactsOutWithDefaults

`func NewArtifactsOutWithDefaults() *ArtifactsOut`

NewArtifactsOutWithDefaults instantiates a new ArtifactsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ArtifactsOut) GetData() []ResearchArtifact`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ArtifactsOut) GetDataOk() (*[]ResearchArtifact, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ArtifactsOut) SetData(v []ResearchArtifact)`

SetData sets Data field to given value.

### HasData

`func (o *ArtifactsOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotal

`func (o *ArtifactsOut) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ArtifactsOut) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ArtifactsOut) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ArtifactsOut) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


