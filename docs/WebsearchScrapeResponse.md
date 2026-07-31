# WebsearchScrapeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Data** | Pointer to [**WebsearchScrapeResponseData**](WebsearchScrapeResponseData.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewWebsearchScrapeResponse

`func NewWebsearchScrapeResponse() *WebsearchScrapeResponse`

NewWebsearchScrapeResponse instantiates a new WebsearchScrapeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsearchScrapeResponseWithDefaults

`func NewWebsearchScrapeResponseWithDefaults() *WebsearchScrapeResponse`

NewWebsearchScrapeResponseWithDefaults instantiates a new WebsearchScrapeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *WebsearchScrapeResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *WebsearchScrapeResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *WebsearchScrapeResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *WebsearchScrapeResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetData

`func (o *WebsearchScrapeResponse) GetData() WebsearchScrapeResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebsearchScrapeResponse) GetDataOk() (*WebsearchScrapeResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WebsearchScrapeResponse) SetData(v WebsearchScrapeResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *WebsearchScrapeResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *WebsearchScrapeResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *WebsearchScrapeResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *WebsearchScrapeResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *WebsearchScrapeResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


