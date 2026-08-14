# LeaderRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gap** | Pointer to **float32** | published − measured (the arena signal) | [optional] 
**Measured** | Pointer to **float32** | hanzo-measured accuracy % (nil if unrun) | [optional] 
**Model** | Pointer to **string** | the model this row scores | [optional] 
**N** | Pointer to **int32** | coverage — NEVER compare across different n | [optional] 
**Protocol** | Pointer to **string** | how the vendor scored their claim: single-attempt, pass@k or agentic | [optional] 
**Published** | Pointer to **float32** | provider-claimed % (nil if none) | [optional] 

## Methods

### NewLeaderRow

`func NewLeaderRow() *LeaderRow`

NewLeaderRow instantiates a new LeaderRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaderRowWithDefaults

`func NewLeaderRowWithDefaults() *LeaderRow`

NewLeaderRowWithDefaults instantiates a new LeaderRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGap

`func (o *LeaderRow) GetGap() float32`

GetGap returns the Gap field if non-nil, zero value otherwise.

### GetGapOk

`func (o *LeaderRow) GetGapOk() (*float32, bool)`

GetGapOk returns a tuple with the Gap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGap

`func (o *LeaderRow) SetGap(v float32)`

SetGap sets Gap field to given value.

### HasGap

`func (o *LeaderRow) HasGap() bool`

HasGap returns a boolean if a field has been set.

### GetMeasured

`func (o *LeaderRow) GetMeasured() float32`

GetMeasured returns the Measured field if non-nil, zero value otherwise.

### GetMeasuredOk

`func (o *LeaderRow) GetMeasuredOk() (*float32, bool)`

GetMeasuredOk returns a tuple with the Measured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasured

`func (o *LeaderRow) SetMeasured(v float32)`

SetMeasured sets Measured field to given value.

### HasMeasured

`func (o *LeaderRow) HasMeasured() bool`

HasMeasured returns a boolean if a field has been set.

### GetModel

`func (o *LeaderRow) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *LeaderRow) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *LeaderRow) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *LeaderRow) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetN

`func (o *LeaderRow) GetN() int32`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *LeaderRow) GetNOk() (*int32, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *LeaderRow) SetN(v int32)`

SetN sets N field to given value.

### HasN

`func (o *LeaderRow) HasN() bool`

HasN returns a boolean if a field has been set.

### GetProtocol

`func (o *LeaderRow) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *LeaderRow) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *LeaderRow) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *LeaderRow) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPublished

`func (o *LeaderRow) GetPublished() float32`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *LeaderRow) GetPublishedOk() (*float32, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *LeaderRow) SetPublished(v float32)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *LeaderRow) HasPublished() bool`

HasPublished returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


