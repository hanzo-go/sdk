# RiskSearchReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Done** | Pointer to **bool** | Done is false while the run is still going; the trials below are then the ones finished so far. | [optional] 
**Ended** | Pointer to **string** | Ended is when it finished, RFC 3339. Absent while it is still going. | [optional] 
**Events** | Pointer to **int32** | Events is how much of this organisation&#39;s history was replayed. | [optional] 
**Fitted** | Pointer to [**RiskModelValue**](RiskModelValue.md) | Fitted is the winning shape FITTED over your own history and published as one of your organisation&#39;s own model values. Name its address on PUT /v1/risk/state/model and the winning shape becomes the model you are running.  It is why this op answers something you can act on. A trial keeps counts and not the model that produced them, so a report without this named a shape nobody could install — and the adoption path refused a shape change besides. Fitting the winner once is a sixty-fifth pass over the same history; keeping all sixty-four fitted models resident instead would cost a measured 21 MiB per run for sixty-three shapes nobody adopts.  Two things about it are worth knowing before you adopt it. Its realised rate can differ from the winner&#39;s above, because the ranking measures every candidate under one fixed reference geometry so the comparison is a comparison, while this is fitted under YOUR geometry — the one an outsider cannot predict. And it has learned the window this search replayed and nothing older, so adopting it trades history for fit. | [optional] 
**Gap** | Pointer to **string** | Gap says why the winning shape could not be fitted into an adoptable value, when it could not. It is separate from Refusal because they are different facts: a refusal means the ranking below proves nothing, a gap means the ranking stands and only the value is missing. | [optional] 
**Id** | Pointer to **string** | ID is the run. | [optional] 
**Refusal** | Pointer to **string** | Refusal says why the run proves nothing, when it does. An empty history is REFUSED rather than reported as zero alerts: \&quot;no alerts\&quot; is exactly what a quiet model looks like, and choosing a shape on the strength of an empty replay is the failure a sandbox exists to prevent. | [optional] 
**Started** | Pointer to **string** | Started is when the run was accepted, RFC 3339. | [optional] 
**Trials** | Pointer to [**[]RiskTrial**](RiskTrial.md) | Trials is every shape tried, best first. | [optional] 
**Winner** | Pointer to [**RiskTrial**](RiskTrial.md) | Winner is the best-fitting shape, absent when nothing fit. | [optional] 

## Methods

### NewRiskSearchReport

`func NewRiskSearchReport() *RiskSearchReport`

NewRiskSearchReport instantiates a new RiskSearchReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskSearchReportWithDefaults

`func NewRiskSearchReportWithDefaults() *RiskSearchReport`

NewRiskSearchReportWithDefaults instantiates a new RiskSearchReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDone

`func (o *RiskSearchReport) GetDone() bool`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *RiskSearchReport) GetDoneOk() (*bool, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *RiskSearchReport) SetDone(v bool)`

SetDone sets Done field to given value.

### HasDone

`func (o *RiskSearchReport) HasDone() bool`

HasDone returns a boolean if a field has been set.

### GetEnded

`func (o *RiskSearchReport) GetEnded() string`

GetEnded returns the Ended field if non-nil, zero value otherwise.

### GetEndedOk

`func (o *RiskSearchReport) GetEndedOk() (*string, bool)`

GetEndedOk returns a tuple with the Ended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnded

`func (o *RiskSearchReport) SetEnded(v string)`

SetEnded sets Ended field to given value.

### HasEnded

`func (o *RiskSearchReport) HasEnded() bool`

HasEnded returns a boolean if a field has been set.

### GetEvents

`func (o *RiskSearchReport) GetEvents() int32`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *RiskSearchReport) GetEventsOk() (*int32, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *RiskSearchReport) SetEvents(v int32)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *RiskSearchReport) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetFitted

`func (o *RiskSearchReport) GetFitted() RiskModelValue`

GetFitted returns the Fitted field if non-nil, zero value otherwise.

### GetFittedOk

`func (o *RiskSearchReport) GetFittedOk() (*RiskModelValue, bool)`

GetFittedOk returns a tuple with the Fitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFitted

`func (o *RiskSearchReport) SetFitted(v RiskModelValue)`

SetFitted sets Fitted field to given value.

### HasFitted

`func (o *RiskSearchReport) HasFitted() bool`

HasFitted returns a boolean if a field has been set.

### GetGap

`func (o *RiskSearchReport) GetGap() string`

GetGap returns the Gap field if non-nil, zero value otherwise.

### GetGapOk

`func (o *RiskSearchReport) GetGapOk() (*string, bool)`

GetGapOk returns a tuple with the Gap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGap

`func (o *RiskSearchReport) SetGap(v string)`

SetGap sets Gap field to given value.

### HasGap

`func (o *RiskSearchReport) HasGap() bool`

HasGap returns a boolean if a field has been set.

### GetId

`func (o *RiskSearchReport) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskSearchReport) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskSearchReport) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RiskSearchReport) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRefusal

`func (o *RiskSearchReport) GetRefusal() string`

GetRefusal returns the Refusal field if non-nil, zero value otherwise.

### GetRefusalOk

`func (o *RiskSearchReport) GetRefusalOk() (*string, bool)`

GetRefusalOk returns a tuple with the Refusal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusal

`func (o *RiskSearchReport) SetRefusal(v string)`

SetRefusal sets Refusal field to given value.

### HasRefusal

`func (o *RiskSearchReport) HasRefusal() bool`

HasRefusal returns a boolean if a field has been set.

### GetStarted

`func (o *RiskSearchReport) GetStarted() string`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *RiskSearchReport) GetStartedOk() (*string, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *RiskSearchReport) SetStarted(v string)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *RiskSearchReport) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetTrials

`func (o *RiskSearchReport) GetTrials() []RiskTrial`

GetTrials returns the Trials field if non-nil, zero value otherwise.

### GetTrialsOk

`func (o *RiskSearchReport) GetTrialsOk() (*[]RiskTrial, bool)`

GetTrialsOk returns a tuple with the Trials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrials

`func (o *RiskSearchReport) SetTrials(v []RiskTrial)`

SetTrials sets Trials field to given value.

### HasTrials

`func (o *RiskSearchReport) HasTrials() bool`

HasTrials returns a boolean if a field has been set.

### GetWinner

`func (o *RiskSearchReport) GetWinner() RiskTrial`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *RiskSearchReport) GetWinnerOk() (*RiskTrial, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *RiskSearchReport) SetWinner(v RiskTrial)`

SetWinner sets Winner field to given value.

### HasWinner

`func (o *RiskSearchReport) HasWinner() bool`

HasWinner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


