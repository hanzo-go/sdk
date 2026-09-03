# RiskSurface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Folded** | Pointer to **int64** | Folded is how many buckets of the tenant&#39;s own feature surface were folded into the model when it became resident. | [optional] 
**Gap** | Pointer to **string** | Gap says why the fold did not happen or did not complete, when that is the case. An empty surface and an unreachable warehouse are different facts and a model must not report them as the same one. | [optional] 
**Refused** | Pointer to **int64** | Refused is how many buckets of this organisation&#39;s own surface the fold could not fold, because a subject on them is longer than this plane&#39;s own field bound. It is history the model does not have, said out loud. | [optional] 
**Replayed** | Pointer to **int64** | Replayed is how many of this organisation&#39;s own recorded observations rebuilt its sliding aggregates when the model became resident. It is what says a rollout was a rebuild rather than a blindness: the aggregates are a projection of a durable record, so a restart costs a replay and not a control. | [optional] 
**Rolled** | Pointer to **int64** | Rolled is how many windows of this organisation&#39;s own source planes — product events, captured failures, metered inference — were rolled up into its feature surface before that fold. Zero with no gap means the surface was already current, which is a different fact from the rollup never running. | [optional] 
**Window** | Pointer to **string** | Window is the lookback the fold covered. | [optional] 

## Methods

### NewRiskSurface

`func NewRiskSurface() *RiskSurface`

NewRiskSurface instantiates a new RiskSurface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskSurfaceWithDefaults

`func NewRiskSurfaceWithDefaults() *RiskSurface`

NewRiskSurfaceWithDefaults instantiates a new RiskSurface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolded

`func (o *RiskSurface) GetFolded() int64`

GetFolded returns the Folded field if non-nil, zero value otherwise.

### GetFoldedOk

`func (o *RiskSurface) GetFoldedOk() (*int64, bool)`

GetFoldedOk returns a tuple with the Folded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolded

`func (o *RiskSurface) SetFolded(v int64)`

SetFolded sets Folded field to given value.

### HasFolded

`func (o *RiskSurface) HasFolded() bool`

HasFolded returns a boolean if a field has been set.

### GetGap

`func (o *RiskSurface) GetGap() string`

GetGap returns the Gap field if non-nil, zero value otherwise.

### GetGapOk

`func (o *RiskSurface) GetGapOk() (*string, bool)`

GetGapOk returns a tuple with the Gap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGap

`func (o *RiskSurface) SetGap(v string)`

SetGap sets Gap field to given value.

### HasGap

`func (o *RiskSurface) HasGap() bool`

HasGap returns a boolean if a field has been set.

### GetRefused

`func (o *RiskSurface) GetRefused() int64`

GetRefused returns the Refused field if non-nil, zero value otherwise.

### GetRefusedOk

`func (o *RiskSurface) GetRefusedOk() (*int64, bool)`

GetRefusedOk returns a tuple with the Refused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefused

`func (o *RiskSurface) SetRefused(v int64)`

SetRefused sets Refused field to given value.

### HasRefused

`func (o *RiskSurface) HasRefused() bool`

HasRefused returns a boolean if a field has been set.

### GetReplayed

`func (o *RiskSurface) GetReplayed() int64`

GetReplayed returns the Replayed field if non-nil, zero value otherwise.

### GetReplayedOk

`func (o *RiskSurface) GetReplayedOk() (*int64, bool)`

GetReplayedOk returns a tuple with the Replayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayed

`func (o *RiskSurface) SetReplayed(v int64)`

SetReplayed sets Replayed field to given value.

### HasReplayed

`func (o *RiskSurface) HasReplayed() bool`

HasReplayed returns a boolean if a field has been set.

### GetRolled

`func (o *RiskSurface) GetRolled() int64`

GetRolled returns the Rolled field if non-nil, zero value otherwise.

### GetRolledOk

`func (o *RiskSurface) GetRolledOk() (*int64, bool)`

GetRolledOk returns a tuple with the Rolled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolled

`func (o *RiskSurface) SetRolled(v int64)`

SetRolled sets Rolled field to given value.

### HasRolled

`func (o *RiskSurface) HasRolled() bool`

HasRolled returns a boolean if a field has been set.

### GetWindow

`func (o *RiskSurface) GetWindow() string`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *RiskSurface) GetWindowOk() (*string, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *RiskSurface) SetWindow(v string)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *RiskSurface) HasWindow() bool`

HasWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


