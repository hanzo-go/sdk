# CloudCrawlResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**CloudCrawlDocument**](CloudCrawlDocument.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewCloudCrawlResult

`func NewCloudCrawlResult() *CloudCrawlResult`

NewCloudCrawlResult instantiates a new CloudCrawlResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCrawlResultWithDefaults

`func NewCloudCrawlResultWithDefaults() *CloudCrawlResult`

NewCloudCrawlResultWithDefaults instantiates a new CloudCrawlResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudCrawlResult) GetData() CloudCrawlDocument`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudCrawlResult) GetDataOk() (*CloudCrawlDocument, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudCrawlResult) SetData(v CloudCrawlDocument)`

SetData sets Data field to given value.

### HasData

`func (o *CloudCrawlResult) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *CloudCrawlResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CloudCrawlResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CloudCrawlResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CloudCrawlResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetSuccess

`func (o *CloudCrawlResult) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CloudCrawlResult) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CloudCrawlResult) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CloudCrawlResult) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


