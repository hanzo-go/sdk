# SearchListBatches200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]SearchBatchView**](SearchBatchView.md) |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**From** | Pointer to **int32** |  | [optional] 
**Next** | Pointer to **int32** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewSearchListBatches200Response

`func NewSearchListBatches200Response() *SearchListBatches200Response`

NewSearchListBatches200Response instantiates a new SearchListBatches200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchListBatches200ResponseWithDefaults

`func NewSearchListBatches200ResponseWithDefaults() *SearchListBatches200Response`

NewSearchListBatches200ResponseWithDefaults instantiates a new SearchListBatches200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *SearchListBatches200Response) GetResults() []SearchBatchView`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SearchListBatches200Response) GetResultsOk() (*[]SearchBatchView, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SearchListBatches200Response) SetResults(v []SearchBatchView)`

SetResults sets Results field to given value.

### HasResults

`func (o *SearchListBatches200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetLimit

`func (o *SearchListBatches200Response) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SearchListBatches200Response) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SearchListBatches200Response) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SearchListBatches200Response) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetFrom

`func (o *SearchListBatches200Response) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SearchListBatches200Response) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SearchListBatches200Response) SetFrom(v int32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *SearchListBatches200Response) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetNext

`func (o *SearchListBatches200Response) GetNext() int32`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *SearchListBatches200Response) GetNextOk() (*int32, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *SearchListBatches200Response) SetNext(v int32)`

SetNext sets Next field to given value.

### HasNext

`func (o *SearchListBatches200Response) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetTotal

`func (o *SearchListBatches200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SearchListBatches200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SearchListBatches200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SearchListBatches200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


