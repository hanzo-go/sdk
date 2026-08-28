# Memory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**Dimension** | Pointer to **int32** |  | [optional] 
**Embedding** | Pointer to **[]float32** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Score** | Pointer to **float32** |  | [optional] 
**UpdatedTime** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewMemory

`func NewMemory() *Memory`

NewMemory instantiates a new Memory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryWithDefaults

`func NewMemoryWithDefaults() *Memory`

NewMemoryWithDefaults instantiates a new Memory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *Memory) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Memory) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Memory) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *Memory) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Memory) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Memory) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Memory) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Memory) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDimension

`func (o *Memory) GetDimension() int32`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *Memory) GetDimensionOk() (*int32, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *Memory) SetDimension(v int32)`

SetDimension sets Dimension field to given value.

### HasDimension

`func (o *Memory) HasDimension() bool`

HasDimension returns a boolean if a field has been set.

### GetEmbedding

`func (o *Memory) GetEmbedding() []float32`

GetEmbedding returns the Embedding field if non-nil, zero value otherwise.

### GetEmbeddingOk

`func (o *Memory) GetEmbeddingOk() (*[]float32, bool)`

GetEmbeddingOk returns a tuple with the Embedding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedding

`func (o *Memory) SetEmbedding(v []float32)`

SetEmbedding sets Embedding field to given value.

### HasEmbedding

`func (o *Memory) HasEmbedding() bool`

HasEmbedding returns a boolean if a field has been set.

### GetKind

`func (o *Memory) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Memory) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Memory) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Memory) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *Memory) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Memory) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Memory) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Memory) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *Memory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Memory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Memory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Memory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *Memory) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Memory) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Memory) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Memory) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetScore

`func (o *Memory) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *Memory) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *Memory) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *Memory) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *Memory) GetUpdatedTime() string`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *Memory) GetUpdatedTimeOk() (*string, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *Memory) SetUpdatedTime(v string)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *Memory) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### GetUserId

`func (o *Memory) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Memory) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Memory) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Memory) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


