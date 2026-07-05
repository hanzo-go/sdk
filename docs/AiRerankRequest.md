# AiRerankRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** |  | 
**Query** | **string** |  | 
**Documents** | **[]string** |  | 
**TopN** | Pointer to **int32** |  | [optional] 

## Methods

### NewAiRerankRequest

`func NewAiRerankRequest(model string, query string, documents []string, ) *AiRerankRequest`

NewAiRerankRequest instantiates a new AiRerankRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiRerankRequestWithDefaults

`func NewAiRerankRequestWithDefaults() *AiRerankRequest`

NewAiRerankRequestWithDefaults instantiates a new AiRerankRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *AiRerankRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiRerankRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiRerankRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetQuery

`func (o *AiRerankRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *AiRerankRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *AiRerankRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetDocuments

`func (o *AiRerankRequest) GetDocuments() []string`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *AiRerankRequest) GetDocumentsOk() (*[]string, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *AiRerankRequest) SetDocuments(v []string)`

SetDocuments sets Documents field to given value.


### GetTopN

`func (o *AiRerankRequest) GetTopN() int32`

GetTopN returns the TopN field if non-nil, zero value otherwise.

### GetTopNOk

`func (o *AiRerankRequest) GetTopNOk() (*int32, bool)`

GetTopNOk returns a tuple with the TopN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopN

`func (o *AiRerankRequest) SetTopN(v int32)`

SetTopN sets TopN field to given value.

### HasTopN

`func (o *AiRerankRequest) HasTopN() bool`

HasTopN returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


