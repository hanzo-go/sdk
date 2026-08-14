# ProfileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyMetrics** | Pointer to [**ProfileMetrics**](ProfileMetrics.md) |  | [optional] 
**Signals** | Pointer to **map[string]bool** |  | [optional] 
**Stage** | Pointer to **string** |  | [optional] 

## Methods

### NewProfileResponse

`func NewProfileResponse() *ProfileResponse`

NewProfileResponse instantiates a new ProfileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileResponseWithDefaults

`func NewProfileResponseWithDefaults() *ProfileResponse`

NewProfileResponseWithDefaults instantiates a new ProfileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyMetrics

`func (o *ProfileResponse) GetKeyMetrics() ProfileMetrics`

GetKeyMetrics returns the KeyMetrics field if non-nil, zero value otherwise.

### GetKeyMetricsOk

`func (o *ProfileResponse) GetKeyMetricsOk() (*ProfileMetrics, bool)`

GetKeyMetricsOk returns a tuple with the KeyMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyMetrics

`func (o *ProfileResponse) SetKeyMetrics(v ProfileMetrics)`

SetKeyMetrics sets KeyMetrics field to given value.

### HasKeyMetrics

`func (o *ProfileResponse) HasKeyMetrics() bool`

HasKeyMetrics returns a boolean if a field has been set.

### GetSignals

`func (o *ProfileResponse) GetSignals() map[string]bool`

GetSignals returns the Signals field if non-nil, zero value otherwise.

### GetSignalsOk

`func (o *ProfileResponse) GetSignalsOk() (*map[string]bool, bool)`

GetSignalsOk returns a tuple with the Signals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignals

`func (o *ProfileResponse) SetSignals(v map[string]bool)`

SetSignals sets Signals field to given value.

### HasSignals

`func (o *ProfileResponse) HasSignals() bool`

HasSignals returns a boolean if a field has been set.

### GetStage

`func (o *ProfileResponse) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *ProfileResponse) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *ProfileResponse) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *ProfileResponse) HasStage() bool`

HasStage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


