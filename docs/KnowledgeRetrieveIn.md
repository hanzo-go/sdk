# KnowledgeRetrieveIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to **[]string** | Files are the files the question is about — the files a chat message references. Empty asks the whole workspace. The graph may still reach files beyond them. | [optional] 
**Limit** | Pointer to **int64** | Limit bounds the passages drilled out of the chosen sections. Default 8, maximum 24. The graph adds up to four more. | [optional] 
**Project** | Pointer to **string** | Project restricts retrieval to files indexed under one project scope. | [optional] 
**Query** | Pointer to **string** | Query is the question to answer. Required. | [optional] 

## Methods

### NewKnowledgeRetrieveIn

`func NewKnowledgeRetrieveIn() *KnowledgeRetrieveIn`

NewKnowledgeRetrieveIn instantiates a new KnowledgeRetrieveIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnowledgeRetrieveInWithDefaults

`func NewKnowledgeRetrieveInWithDefaults() *KnowledgeRetrieveIn`

NewKnowledgeRetrieveInWithDefaults instantiates a new KnowledgeRetrieveIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *KnowledgeRetrieveIn) GetFiles() []string`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *KnowledgeRetrieveIn) GetFilesOk() (*[]string, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *KnowledgeRetrieveIn) SetFiles(v []string)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *KnowledgeRetrieveIn) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetLimit

`func (o *KnowledgeRetrieveIn) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *KnowledgeRetrieveIn) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *KnowledgeRetrieveIn) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *KnowledgeRetrieveIn) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetProject

`func (o *KnowledgeRetrieveIn) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *KnowledgeRetrieveIn) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *KnowledgeRetrieveIn) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *KnowledgeRetrieveIn) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetQuery

`func (o *KnowledgeRetrieveIn) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *KnowledgeRetrieveIn) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *KnowledgeRetrieveIn) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *KnowledgeRetrieveIn) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


