# KnowledgeFileEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to **int64** | Files is how many other files of the org name it too. | [optional] 
**Kind** | Pointer to **string** | Kind is person, org, product, place or term. | [optional] 
**Name** | Pointer to **string** | Name is the entity as the text names it. | [optional] 
**Sections** | Pointer to **int64** | Sections is how many of the file&#39;s sections name it. | [optional] 

## Methods

### NewKnowledgeFileEntity

`func NewKnowledgeFileEntity() *KnowledgeFileEntity`

NewKnowledgeFileEntity instantiates a new KnowledgeFileEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnowledgeFileEntityWithDefaults

`func NewKnowledgeFileEntityWithDefaults() *KnowledgeFileEntity`

NewKnowledgeFileEntityWithDefaults instantiates a new KnowledgeFileEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *KnowledgeFileEntity) GetFiles() int64`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *KnowledgeFileEntity) GetFilesOk() (*int64, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *KnowledgeFileEntity) SetFiles(v int64)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *KnowledgeFileEntity) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetKind

`func (o *KnowledgeFileEntity) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KnowledgeFileEntity) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KnowledgeFileEntity) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *KnowledgeFileEntity) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *KnowledgeFileEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KnowledgeFileEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KnowledgeFileEntity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KnowledgeFileEntity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSections

`func (o *KnowledgeFileEntity) GetSections() int64`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *KnowledgeFileEntity) GetSectionsOk() (*int64, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *KnowledgeFileEntity) SetSections(v int64)`

SetSections sets Sections field to given value.

### HasSections

`func (o *KnowledgeFileEntity) HasSections() bool`

HasSections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


