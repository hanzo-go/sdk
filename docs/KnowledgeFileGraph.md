# KnowledgeFileGraph

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | Pointer to [**[]KnowledgeFileEntity**](KnowledgeFileEntity.md) | Entities are the entities it names, most mentioned first. | [optional] 
**File** | Pointer to [**KnowledgeFile**](KnowledgeFile.md) | File is the file the graph is of. | [optional] 
**Links** | Pointer to [**[]KnowledgeFileLink**](KnowledgeFileLink.md) | Links are the other files it is joined to, strongest first. | [optional] 

## Methods

### NewKnowledgeFileGraph

`func NewKnowledgeFileGraph() *KnowledgeFileGraph`

NewKnowledgeFileGraph instantiates a new KnowledgeFileGraph object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnowledgeFileGraphWithDefaults

`func NewKnowledgeFileGraphWithDefaults() *KnowledgeFileGraph`

NewKnowledgeFileGraphWithDefaults instantiates a new KnowledgeFileGraph object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *KnowledgeFileGraph) GetEntities() []KnowledgeFileEntity`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *KnowledgeFileGraph) GetEntitiesOk() (*[]KnowledgeFileEntity, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *KnowledgeFileGraph) SetEntities(v []KnowledgeFileEntity)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *KnowledgeFileGraph) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetFile

`func (o *KnowledgeFileGraph) GetFile() KnowledgeFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *KnowledgeFileGraph) GetFileOk() (*KnowledgeFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *KnowledgeFileGraph) SetFile(v KnowledgeFile)`

SetFile sets File field to given value.

### HasFile

`func (o *KnowledgeFileGraph) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetLinks

`func (o *KnowledgeFileGraph) GetLinks() []KnowledgeFileLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *KnowledgeFileGraph) GetLinksOk() (*[]KnowledgeFileLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *KnowledgeFileGraph) SetLinks(v []KnowledgeFileLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *KnowledgeFileGraph) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


