# RiskResolveOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to [**[]RiskResolved**](RiskResolved.md) | Labels is one entry per named event that BOTH matured and had at least one assertion knowable by its own as-of, in the order the events were named. The three outcomes partition the ask: len(labels) + Unmatured + Unlabelled is the number of DISTINCT events named, an event named twice having been answered once. | [optional] 
**Unlabelled** | Pointer to **int32** | Unlabelled is how many matured events had no assertion knowable by their own as-of. That is the ordinary state of most traffic and it is reported rather than answered as unproductive: manufacturing negatives is how a fraud model comes to describe the incumbent block list. | [optional] 
**Unmatured** | Pointer to **int32** | Unmatured is how many named events had not aged past the horizon. They are not unlabelled — they are not yet ASKABLE, and a supervised training set must exclude them rather than treat them as negatives. | [optional] 

## Methods

### NewRiskResolveOut

`func NewRiskResolveOut() *RiskResolveOut`

NewRiskResolveOut instantiates a new RiskResolveOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskResolveOutWithDefaults

`func NewRiskResolveOutWithDefaults() *RiskResolveOut`

NewRiskResolveOutWithDefaults instantiates a new RiskResolveOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHorizon

`func (o *RiskResolveOut) GetHorizon() int32`

GetHorizon returns the Horizon field if non-nil, zero value otherwise.

### GetHorizonOk

`func (o *RiskResolveOut) GetHorizonOk() (*int32, bool)`

GetHorizonOk returns a tuple with the Horizon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizon

`func (o *RiskResolveOut) SetHorizon(v int32)`

SetHorizon sets Horizon field to given value.

### HasHorizon

`func (o *RiskResolveOut) HasHorizon() bool`

HasHorizon returns a boolean if a field has been set.

### GetLabels

`func (o *RiskResolveOut) GetLabels() []RiskResolved`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *RiskResolveOut) GetLabelsOk() (*[]RiskResolved, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *RiskResolveOut) SetLabels(v []RiskResolved)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *RiskResolveOut) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetNow

`func (o *RiskResolveOut) GetNow() string`

GetNow returns the Now field if non-nil, zero value otherwise.

### GetNowOk

`func (o *RiskResolveOut) GetNowOk() (*string, bool)`

GetNowOk returns a tuple with the Now field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNow

`func (o *RiskResolveOut) SetNow(v string)`

SetNow sets Now field to given value.

### HasNow

`func (o *RiskResolveOut) HasNow() bool`

HasNow returns a boolean if a field has been set.

### GetUnlabelled

`func (o *RiskResolveOut) GetUnlabelled() int32`

GetUnlabelled returns the Unlabelled field if non-nil, zero value otherwise.

### GetUnlabelledOk

`func (o *RiskResolveOut) GetUnlabelledOk() (*int32, bool)`

GetUnlabelledOk returns a tuple with the Unlabelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlabelled

`func (o *RiskResolveOut) SetUnlabelled(v int32)`

SetUnlabelled sets Unlabelled field to given value.

### HasUnlabelled

`func (o *RiskResolveOut) HasUnlabelled() bool`

HasUnlabelled returns a boolean if a field has been set.

### GetUnmatured

`func (o *RiskResolveOut) GetUnmatured() int32`

GetUnmatured returns the Unmatured field if non-nil, zero value otherwise.

### GetUnmaturedOk

`func (o *RiskResolveOut) GetUnmaturedOk() (*int32, bool)`

GetUnmaturedOk returns a tuple with the Unmatured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatured

`func (o *RiskResolveOut) SetUnmatured(v int32)`

SetUnmatured sets Unmatured field to given value.

### HasUnmatured

`func (o *RiskResolveOut) HasUnmatured() bool`

HasUnmatured returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


