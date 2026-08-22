# IndexHits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EstimatedTotalHits** | Pointer to **int32** | EstimatedTotalHits is the dialect&#39;s name for the match count. Every hit is materialised here, so for this page it is exact rather than estimated. | [optional] 
**Hits** | Pointer to **[]interface{}** | Hits are the matching documents, most relevant first, exactly as stored. | [optional] 
**Limit** | Pointer to **int32** | Limit is how many hits this page could hold. | [optional] 
**Offset** | Pointer to **int32** | Offset is where this page starts. | [optional] 
**ProcessingTimeMs** | Pointer to **int32** | ProcessingTimeMs is how long the query took, in milliseconds. | [optional] 
**Query** | Pointer to **string** | Query echoes the search terms, which is what a client renders above the results. | [optional] 

## Methods

### NewIndexHits

`func NewIndexHits() *IndexHits`

NewIndexHits instantiates a new IndexHits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexHitsWithDefaults

`func NewIndexHitsWithDefaults() *IndexHits`

NewIndexHitsWithDefaults instantiates a new IndexHits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEstimatedTotalHits

`func (o *IndexHits) GetEstimatedTotalHits() int32`

GetEstimatedTotalHits returns the EstimatedTotalHits field if non-nil, zero value otherwise.

### GetEstimatedTotalHitsOk

`func (o *IndexHits) GetEstimatedTotalHitsOk() (*int32, bool)`

GetEstimatedTotalHitsOk returns a tuple with the EstimatedTotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTotalHits

`func (o *IndexHits) SetEstimatedTotalHits(v int32)`

SetEstimatedTotalHits sets EstimatedTotalHits field to given value.

### HasEstimatedTotalHits

`func (o *IndexHits) HasEstimatedTotalHits() bool`

HasEstimatedTotalHits returns a boolean if a field has been set.

### GetHits

`func (o *IndexHits) GetHits() []interface{}`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *IndexHits) GetHitsOk() (*[]interface{}, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *IndexHits) SetHits(v []interface{})`

SetHits sets Hits field to given value.

### HasHits

`func (o *IndexHits) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetLimit

`func (o *IndexHits) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *IndexHits) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *IndexHits) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *IndexHits) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *IndexHits) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *IndexHits) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *IndexHits) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *IndexHits) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetProcessingTimeMs

`func (o *IndexHits) GetProcessingTimeMs() int32`

GetProcessingTimeMs returns the ProcessingTimeMs field if non-nil, zero value otherwise.

### GetProcessingTimeMsOk

`func (o *IndexHits) GetProcessingTimeMsOk() (*int32, bool)`

GetProcessingTimeMsOk returns a tuple with the ProcessingTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeMs

`func (o *IndexHits) SetProcessingTimeMs(v int32)`

SetProcessingTimeMs sets ProcessingTimeMs field to given value.

### HasProcessingTimeMs

`func (o *IndexHits) HasProcessingTimeMs() bool`

HasProcessingTimeMs returns a boolean if a field has been set.

### GetQuery

`func (o *IndexHits) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *IndexHits) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *IndexHits) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *IndexHits) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


