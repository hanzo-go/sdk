# KnowledgeFileSearchIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to **[]string** | Files restricts the search to these file ids. Empty searches every file the org holds. | [optional] 
**Limit** | Pointer to **int64** | Limit bounds the passages returned. Default 8, maximum 50. | [optional] 
**Project** | Pointer to **string** | Project restricts the search to files indexed under one project scope. | [optional] 
**Query** | Pointer to **string** | Query is the question or phrase. Required. | [optional] 

## Methods

### NewKnowledgeFileSearchIn

`func NewKnowledgeFileSearchIn() *KnowledgeFileSearchIn`

NewKnowledgeFileSearchIn instantiates a new KnowledgeFileSearchIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnowledgeFileSearchInWithDefaults

`func NewKnowledgeFileSearchInWithDefaults() *KnowledgeFileSearchIn`

NewKnowledgeFileSearchInWithDefaults instantiates a new KnowledgeFileSearchIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *KnowledgeFileSearchIn) GetFiles() []string`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *KnowledgeFileSearchIn) GetFilesOk() (*[]string, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *KnowledgeFileSearchIn) SetFiles(v []string)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *KnowledgeFileSearchIn) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetLimit

`func (o *KnowledgeFileSearchIn) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *KnowledgeFileSearchIn) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *KnowledgeFileSearchIn) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *KnowledgeFileSearchIn) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetProject

`func (o *KnowledgeFileSearchIn) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *KnowledgeFileSearchIn) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *KnowledgeFileSearchIn) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *KnowledgeFileSearchIn) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetQuery

`func (o *KnowledgeFileSearchIn) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *KnowledgeFileSearchIn) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *KnowledgeFileSearchIn) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *KnowledgeFileSearchIn) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


