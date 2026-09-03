# AnalyzeQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alpha** | Pointer to **float64** | Alpha overrides the 0.05 two-tailed significance threshold when it lies strictly between 0 and 1; anything else leaves the default in place. | [optional] 
**Days** | Pointer to **int64** | Days is how far back to read when no start is given: 1 to 365, 30 by default. A value outside that range leaves the default in place. | [optional] 
**End** | Pointer to **string** | End is the window&#39;s exclusive end in RFC3339, defaulting to now. | [optional] 
**Id** | Pointer to **string** | ID is the experiment the URL names. | [optional] 
**Start** | Pointer to **string** | Start is the window&#39;s inclusive start in RFC3339. Given, it wins over days. | [optional] 

## Methods

### NewAnalyzeQuery

`func NewAnalyzeQuery() *AnalyzeQuery`

NewAnalyzeQuery instantiates a new AnalyzeQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyzeQueryWithDefaults

`func NewAnalyzeQueryWithDefaults() *AnalyzeQuery`

NewAnalyzeQueryWithDefaults instantiates a new AnalyzeQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlpha

`func (o *AnalyzeQuery) GetAlpha() float64`

GetAlpha returns the Alpha field if non-nil, zero value otherwise.

### GetAlphaOk

`func (o *AnalyzeQuery) GetAlphaOk() (*float64, bool)`

GetAlphaOk returns a tuple with the Alpha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpha

`func (o *AnalyzeQuery) SetAlpha(v float64)`

SetAlpha sets Alpha field to given value.

### HasAlpha

`func (o *AnalyzeQuery) HasAlpha() bool`

HasAlpha returns a boolean if a field has been set.

### GetDays

`func (o *AnalyzeQuery) GetDays() int64`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *AnalyzeQuery) GetDaysOk() (*int64, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *AnalyzeQuery) SetDays(v int64)`

SetDays sets Days field to given value.

### HasDays

`func (o *AnalyzeQuery) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetEnd

`func (o *AnalyzeQuery) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *AnalyzeQuery) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *AnalyzeQuery) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *AnalyzeQuery) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetId

`func (o *AnalyzeQuery) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AnalyzeQuery) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AnalyzeQuery) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AnalyzeQuery) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStart

`func (o *AnalyzeQuery) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *AnalyzeQuery) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *AnalyzeQuery) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *AnalyzeQuery) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


