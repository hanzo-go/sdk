# CrawlResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**CrawlDocument**](CrawlDocument.md) | Data is the page, present exactly when Success. | [optional] 
**Error** | Pointer to **string** | Error says what stopped the fetch: the host was refused, unreachable, or served something that is not a document. | [optional] 
**Success** | Pointer to **bool** | Success is whether the page was fetched and read. FALSE with an Error is a complete answer, not a fault — check this before reading Data. | [optional] 

## Methods

### NewCrawlResult

`func NewCrawlResult() *CrawlResult`

NewCrawlResult instantiates a new CrawlResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrawlResultWithDefaults

`func NewCrawlResultWithDefaults() *CrawlResult`

NewCrawlResultWithDefaults instantiates a new CrawlResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CrawlResult) GetData() CrawlDocument`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CrawlResult) GetDataOk() (*CrawlDocument, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CrawlResult) SetData(v CrawlDocument)`

SetData sets Data field to given value.

### HasData

`func (o *CrawlResult) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *CrawlResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CrawlResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CrawlResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CrawlResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetSuccess

`func (o *CrawlResult) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CrawlResult) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CrawlResult) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CrawlResult) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


