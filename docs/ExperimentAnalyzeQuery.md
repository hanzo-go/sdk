# ExperimentAnalyzeQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alpha** | Pointer to **float64** | Alpha overrides the 0.05 two-tailed significance threshold when it lies strictly between 0 and 1; anything else leaves the default in place. | [optional] 
**Days** | Pointer to **int64** | Days is how far back to read when no start is given: 1 to 365, 30 by default. A value outside that range leaves the default in place. | [optional] 
**End** | Pointer to **string** | End is the window&#39;s exclusive end in RFC3339, defaulting to now. | [optional] 
**Id** | Pointer to **string** | ID is the experiment the URL names. | [optional] 
**Start** | Pointer to **string** | Start is the window&#39;s inclusive start in RFC3339. Given, it wins over days. | [optional] 

## Methods

### NewExperimentAnalyzeQuery

`func NewExperimentAnalyzeQuery() *ExperimentAnalyzeQuery`

NewExperimentAnalyzeQuery instantiates a new ExperimentAnalyzeQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentAnalyzeQueryWithDefaults

`func NewExperimentAnalyzeQueryWithDefaults() *ExperimentAnalyzeQuery`

NewExperimentAnalyzeQueryWithDefaults instantiates a new ExperimentAnalyzeQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlpha

`func (o *ExperimentAnalyzeQuery) GetAlpha() float64`

GetAlpha returns the Alpha field if non-nil, zero value otherwise.

### GetAlphaOk

`func (o *ExperimentAnalyzeQuery) GetAlphaOk() (*float64, bool)`

GetAlphaOk returns a tuple with the Alpha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpha

`func (o *ExperimentAnalyzeQuery) SetAlpha(v float64)`

SetAlpha sets Alpha field to given value.

### HasAlpha

`func (o *ExperimentAnalyzeQuery) HasAlpha() bool`

HasAlpha returns a boolean if a field has been set.

### GetDays

`func (o *ExperimentAnalyzeQuery) GetDays() int64`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *ExperimentAnalyzeQuery) GetDaysOk() (*int64, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *ExperimentAnalyzeQuery) SetDays(v int64)`

SetDays sets Days field to given value.

### HasDays

`func (o *ExperimentAnalyzeQuery) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetEnd

`func (o *ExperimentAnalyzeQuery) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ExperimentAnalyzeQuery) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ExperimentAnalyzeQuery) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ExperimentAnalyzeQuery) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetId

`func (o *ExperimentAnalyzeQuery) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExperimentAnalyzeQuery) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExperimentAnalyzeQuery) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExperimentAnalyzeQuery) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStart

`func (o *ExperimentAnalyzeQuery) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ExperimentAnalyzeQuery) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ExperimentAnalyzeQuery) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *ExperimentAnalyzeQuery) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


