# RiskLabelCoverage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contested** | Pointer to **int32** | Contested is how many matured events have two visible assertions that disagree. It is the number that says whether the precedence rule is load-bearing or decorative, and it is the one to watch after wiring a new source. | [optional] 
**Events** | Pointer to **int32** | Events is how many DISTINCT judged events those assertions name, keyed on (kind, subject, at). It counts only events something was ASSERTED about: what share of the whole event stream carries a label is a question about the feature plane&#39;s denominator and is not answerable here. Matured + Unmatured is Events. | [optional] 
**Explore** | Pointer to **float32** | Explore is the share of judged events whose winning assertion came from the below-the-line sample. A blocked transaction never produces a chargeback, so a training set with no exploration in it is a description of the incumbent block list rather than of the world — and a champion measured on it is measured on whether it agrees with the incumbent. | [optional] 
**Facts** | Pointer to **int32** | Facts is how many assertions the window holds; Events is how many distinct judged events they cover. The two differ by exactly the corroboration and the conflict in the plane. | [optional] 
**From** | Pointer to **string** | From is the INCLUSIVE start of the EVENT window these counts were folded over, RFC 3339, echoed with the defaults filled in — the caller&#39;s, or 90 days before To. An assertion is in the window when its event time satisfies at &gt;&#x3D; From. | [optional] 
**Horizon** | Pointer to **int32** | Horizon is the maturity horizon these counts were measured under, IN DAYS — the caller&#39;s, or 120. It decides Matured (an event is matured when its &#x60;at&#x60; plus this many days is not after now), it sets each event&#39;s own as-of and so which assertions were visible to it, and when the caller bounds nothing it also places the default window&#39;s end. | [optional] 
**Judged** | Pointer to **int32** | Judged is how many MATURED events resolve, at their own as-of, to something other than unjudged. | [optional] 
**Matured** | Pointer to **int32** | Matured is how many of those events have aged past the horizon and may therefore be admitted to a supervised set at all. It counts every matured event, judged or not — it is the DENOMINATOR an operator divides Judged by, and a denominator that excluded the unjudged would read 1.0 on a plane with one label in it. | [optional] 
**Pending** | Pointer to **int32** | Pending is how many of this tenant&#39;s assertions the DERIVED columnar copy is not known to hold yet. Every count above is folded from the record, so they are right regardless — but a materialiser that joins in the warehouse while this is non-zero is joining against an incomplete answer key, and a missing fraud label is indistinguishable from an honest customer. It is reported at the training gate because that is where somebody is deciding whether the ground truth is good enough to fit on. Counted under a cap, so it saturates rather than costing a full scan on every read. | [optional] 
**Productive** | Pointer to **int32** | Productive is how many matured events resolve, at their own as-of, to a WINNING assertion of &#x60;productive&#x60; — the event led somewhere: escalated, reported, charged back. It is the positive class a supervised fit would train on, and a near-zero count is the number that says the fit is not worth running. | [optional] 
**Sources** | Pointer to [**[]RiskSourceCoverage**](RiskSourceCoverage.md) | Sources breaks the judged events down by the source that WON, so a plane that looks labelled because one noisy source dominates is visible as such. | [optional] 
**To** | Pointer to **string** | To is the EXCLUSIVE end of that window (at &lt; To). Unstated it is one horizon before now, never now: a window running to now under a maturity horizon can hold no matured event at all, so every count below would read zero however much ground truth the tenant held. | [optional] 
**Unlabelled** | Pointer to **int32** | Unlabelled is how many MATURED events had no assertion knowable by their own as-of — including every assertion that arrived after that instant. It is the field that says WHY judged is low: a tenant whose ground truth was filed long after the events it judges reads matured&#x3D;n, judged&#x3D;0, unlabelled&#x3D;n, which is diagnosable, rather than a bare zero, which is not. | [optional] 
**Unmatured** | Pointer to **int32** | Unmatured is how many events in the window have NOT aged past the horizon. They are not unlabelled — they are not yet askable, and a supervised set must exclude them rather than treat them as negatives. Matured + Unmatured is Events. | [optional] 
**Unproductive** | Pointer to **int32** | Unproductive is every OTHER judged event: the winner claimed &#x60;unproductive&#x60;, judged not suspicious. Productive + Unproductive is Judged exactly, because a winner of the explicit unjudged is counted in neither — it is a matured event somebody looked at and could not conclude about, and rolling it into the negatives would hand a model a claim nobody made. | [optional] 

## Methods

### NewRiskLabelCoverage

`func NewRiskLabelCoverage() *RiskLabelCoverage`

NewRiskLabelCoverage instantiates a new RiskLabelCoverage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskLabelCoverageWithDefaults

`func NewRiskLabelCoverageWithDefaults() *RiskLabelCoverage`

NewRiskLabelCoverageWithDefaults instantiates a new RiskLabelCoverage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContested

`func (o *RiskLabelCoverage) GetContested() int32`

GetContested returns the Contested field if non-nil, zero value otherwise.

### GetContestedOk

`func (o *RiskLabelCoverage) GetContestedOk() (*int32, bool)`

GetContestedOk returns a tuple with the Contested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContested

`func (o *RiskLabelCoverage) SetContested(v int32)`

SetContested sets Contested field to given value.

### HasContested

`func (o *RiskLabelCoverage) HasContested() bool`

HasContested returns a boolean if a field has been set.

### GetEvents

`func (o *RiskLabelCoverage) GetEvents() int32`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *RiskLabelCoverage) GetEventsOk() (*int32, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *RiskLabelCoverage) SetEvents(v int32)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *RiskLabelCoverage) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetExplore

`func (o *RiskLabelCoverage) GetExplore() float32`

GetExplore returns the Explore field if non-nil, zero value otherwise.

### GetExploreOk

`func (o *RiskLabelCoverage) GetExploreOk() (*float32, bool)`

GetExploreOk returns a tuple with the Explore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplore

`func (o *RiskLabelCoverage) SetExplore(v float32)`

SetExplore sets Explore field to given value.

### HasExplore

`func (o *RiskLabelCoverage) HasExplore() bool`

HasExplore returns a boolean if a field has been set.

### GetFacts

`func (o *RiskLabelCoverage) GetFacts() int32`

GetFacts returns the Facts field if non-nil, zero value otherwise.

### GetFactsOk

`func (o *RiskLabelCoverage) GetFactsOk() (*int32, bool)`

GetFactsOk returns a tuple with the Facts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacts

`func (o *RiskLabelCoverage) SetFacts(v int32)`

SetFacts sets Facts field to given value.

### HasFacts

`func (o *RiskLabelCoverage) HasFacts() bool`

HasFacts returns a boolean if a field has been set.

### GetFrom

`func (o *RiskLabelCoverage) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *RiskLabelCoverage) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *RiskLabelCoverage) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *RiskLabelCoverage) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetHorizon

`func (o *RiskLabelCoverage) GetHorizon() int32`

GetHorizon returns the Horizon field if non-nil, zero value otherwise.

### GetHorizonOk

`func (o *RiskLabelCoverage) GetHorizonOk() (*int32, bool)`

GetHorizonOk returns a tuple with the Horizon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizon

`func (o *RiskLabelCoverage) SetHorizon(v int32)`

SetHorizon sets Horizon field to given value.

### HasHorizon

`func (o *RiskLabelCoverage) HasHorizon() bool`

HasHorizon returns a boolean if a field has been set.

### GetJudged

`func (o *RiskLabelCoverage) GetJudged() int32`

GetJudged returns the Judged field if non-nil, zero value otherwise.

### GetJudgedOk

`func (o *RiskLabelCoverage) GetJudgedOk() (*int32, bool)`

GetJudgedOk returns a tuple with the Judged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudged

`func (o *RiskLabelCoverage) SetJudged(v int32)`

SetJudged sets Judged field to given value.

### HasJudged

`func (o *RiskLabelCoverage) HasJudged() bool`

HasJudged returns a boolean if a field has been set.

### GetMatured

`func (o *RiskLabelCoverage) GetMatured() int32`

GetMatured returns the Matured field if non-nil, zero value otherwise.

### GetMaturedOk

`func (o *RiskLabelCoverage) GetMaturedOk() (*int32, bool)`

GetMaturedOk returns a tuple with the Matured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatured

`func (o *RiskLabelCoverage) SetMatured(v int32)`

SetMatured sets Matured field to given value.

### HasMatured

`func (o *RiskLabelCoverage) HasMatured() bool`

HasMatured returns a boolean if a field has been set.

### GetPending

`func (o *RiskLabelCoverage) GetPending() int32`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *RiskLabelCoverage) GetPendingOk() (*int32, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *RiskLabelCoverage) SetPending(v int32)`

SetPending sets Pending field to given value.

### HasPending

`func (o *RiskLabelCoverage) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetProductive

`func (o *RiskLabelCoverage) GetProductive() int32`

GetProductive returns the Productive field if non-nil, zero value otherwise.

### GetProductiveOk

`func (o *RiskLabelCoverage) GetProductiveOk() (*int32, bool)`

GetProductiveOk returns a tuple with the Productive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductive

`func (o *RiskLabelCoverage) SetProductive(v int32)`

SetProductive sets Productive field to given value.

### HasProductive

`func (o *RiskLabelCoverage) HasProductive() bool`

HasProductive returns a boolean if a field has been set.

### GetSources

`func (o *RiskLabelCoverage) GetSources() []RiskSourceCoverage`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *RiskLabelCoverage) GetSourcesOk() (*[]RiskSourceCoverage, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *RiskLabelCoverage) SetSources(v []RiskSourceCoverage)`

SetSources sets Sources field to given value.

### HasSources

`func (o *RiskLabelCoverage) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetTo

`func (o *RiskLabelCoverage) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *RiskLabelCoverage) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *RiskLabelCoverage) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *RiskLabelCoverage) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetUnlabelled

`func (o *RiskLabelCoverage) GetUnlabelled() int32`

GetUnlabelled returns the Unlabelled field if non-nil, zero value otherwise.

### GetUnlabelledOk

`func (o *RiskLabelCoverage) GetUnlabelledOk() (*int32, bool)`

GetUnlabelledOk returns a tuple with the Unlabelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlabelled

`func (o *RiskLabelCoverage) SetUnlabelled(v int32)`

SetUnlabelled sets Unlabelled field to given value.

### HasUnlabelled

`func (o *RiskLabelCoverage) HasUnlabelled() bool`

HasUnlabelled returns a boolean if a field has been set.

### GetUnmatured

`func (o *RiskLabelCoverage) GetUnmatured() int32`

GetUnmatured returns the Unmatured field if non-nil, zero value otherwise.

### GetUnmaturedOk

`func (o *RiskLabelCoverage) GetUnmaturedOk() (*int32, bool)`

GetUnmaturedOk returns a tuple with the Unmatured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatured

`func (o *RiskLabelCoverage) SetUnmatured(v int32)`

SetUnmatured sets Unmatured field to given value.

### HasUnmatured

`func (o *RiskLabelCoverage) HasUnmatured() bool`

HasUnmatured returns a boolean if a field has been set.

### GetUnproductive

`func (o *RiskLabelCoverage) GetUnproductive() int32`

GetUnproductive returns the Unproductive field if non-nil, zero value otherwise.

### GetUnproductiveOk

`func (o *RiskLabelCoverage) GetUnproductiveOk() (*int32, bool)`

GetUnproductiveOk returns a tuple with the Unproductive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnproductive

`func (o *RiskLabelCoverage) SetUnproductive(v int32)`

SetUnproductive sets Unproductive field to given value.

### HasUnproductive

`func (o *RiskLabelCoverage) HasUnproductive() bool`

HasUnproductive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


