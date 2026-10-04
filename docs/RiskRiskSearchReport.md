# RiskRiskSearchReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Done** | Pointer to **bool** | Done is false while the run is still going; the trials below are then the ones finished so far. | [optional] 
**Ended** | Pointer to **string** | Ended is when it finished, RFC 3339. Absent while it is still going. | [optional] 
**Events** | Pointer to **int64** | Events is how much of this organisation&#39;s history was replayed. | [optional] 
**Fitted** | Pointer to [**RiskRiskModelValue**](RiskRiskModelValue.md) | Fitted is the winning shape FITTED over your own history and published as one of your organisation&#39;s own model values. Name its address on PUT /v1/risk/state/model and the winning shape becomes the model you are running.  It is why this op answers something you can act on. A trial keeps counts and not the model that produced them, so a report without this named a shape nobody could install — and the adoption path refused a shape change besides. Fitting the winner once is a sixty-fifth pass over the same history; keeping all sixty-four fitted models resident instead would cost a measured 21 MiB per run for sixty-three shapes nobody adopts.  Two things about it are worth knowing before you adopt it. Its realised rate can differ from the winner&#39;s above, because the ranking measures every candidate under one fixed reference geometry so the comparison is a comparison, while this is fitted under YOUR geometry — the one an outsider cannot predict. And it has learned the window this search replayed and nothing older, so adopting it trades history for fit. | [optional] 
**Gap** | Pointer to **string** | Gap says why the winning shape could not be fitted into an adoptable value, when it could not. It is separate from Refusal because they are different facts: a refusal means the ranking below proves nothing, a gap means the ranking stands and only the value is missing. | [optional] 
**Id** | Pointer to **string** | ID is the run. | [optional] 
**Refusal** | Pointer to **string** | Refusal says why the run proves nothing, when it does. An empty history is REFUSED rather than reported as zero alerts: \&quot;no alerts\&quot; is exactly what a quiet model looks like, and choosing a shape on the strength of an empty replay is the failure a sandbox exists to prevent. | [optional] 
**Started** | Pointer to **string** | Started is when the run was accepted, RFC 3339. | [optional] 
**Trials** | Pointer to [**[]RiskRiskTrial**](RiskRiskTrial.md) | Trials is every shape tried, best first. | [optional] 
**Winner** | Pointer to [**RiskRiskTrial**](RiskRiskTrial.md) | Winner is the best-fitting shape, absent when nothing fit. | [optional] 

## Methods

### NewRiskRiskSearchReport

`func NewRiskRiskSearchReport() *RiskRiskSearchReport`

NewRiskRiskSearchReport instantiates a new RiskRiskSearchReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskRiskSearchReportWithDefaults

`func NewRiskRiskSearchReportWithDefaults() *RiskRiskSearchReport`

NewRiskRiskSearchReportWithDefaults instantiates a new RiskRiskSearchReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDone

`func (o *RiskRiskSearchReport) GetDone() bool`

GetDone returns the Done field if non-nil, zero value otherwise.

### GetDoneOk

`func (o *RiskRiskSearchReport) GetDoneOk() (*bool, bool)`

GetDoneOk returns a tuple with the Done field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDone

`func (o *RiskRiskSearchReport) SetDone(v bool)`

SetDone sets Done field to given value.

### HasDone

`func (o *RiskRiskSearchReport) HasDone() bool`

HasDone returns a boolean if a field has been set.

### GetEnded

`func (o *RiskRiskSearchReport) GetEnded() string`

GetEnded returns the Ended field if non-nil, zero value otherwise.

### GetEndedOk

`func (o *RiskRiskSearchReport) GetEndedOk() (*string, bool)`

GetEndedOk returns a tuple with the Ended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnded

`func (o *RiskRiskSearchReport) SetEnded(v string)`

SetEnded sets Ended field to given value.

### HasEnded

`func (o *RiskRiskSearchReport) HasEnded() bool`

HasEnded returns a boolean if a field has been set.

### GetEvents

`func (o *RiskRiskSearchReport) GetEvents() int64`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *RiskRiskSearchReport) GetEventsOk() (*int64, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *RiskRiskSearchReport) SetEvents(v int64)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *RiskRiskSearchReport) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetFitted

`func (o *RiskRiskSearchReport) GetFitted() RiskRiskModelValue`

GetFitted returns the Fitted field if non-nil, zero value otherwise.

### GetFittedOk

`func (o *RiskRiskSearchReport) GetFittedOk() (*RiskRiskModelValue, bool)`

GetFittedOk returns a tuple with the Fitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFitted

`func (o *RiskRiskSearchReport) SetFitted(v RiskRiskModelValue)`

SetFitted sets Fitted field to given value.

### HasFitted

`func (o *RiskRiskSearchReport) HasFitted() bool`

HasFitted returns a boolean if a field has been set.

### GetGap

`func (o *RiskRiskSearchReport) GetGap() string`

GetGap returns the Gap field if non-nil, zero value otherwise.

### GetGapOk

`func (o *RiskRiskSearchReport) GetGapOk() (*string, bool)`

GetGapOk returns a tuple with the Gap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGap

`func (o *RiskRiskSearchReport) SetGap(v string)`

SetGap sets Gap field to given value.

### HasGap

`func (o *RiskRiskSearchReport) HasGap() bool`

HasGap returns a boolean if a field has been set.

### GetId

`func (o *RiskRiskSearchReport) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskRiskSearchReport) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskRiskSearchReport) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RiskRiskSearchReport) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRefusal

`func (o *RiskRiskSearchReport) GetRefusal() string`

GetRefusal returns the Refusal field if non-nil, zero value otherwise.

### GetRefusalOk

`func (o *RiskRiskSearchReport) GetRefusalOk() (*string, bool)`

GetRefusalOk returns a tuple with the Refusal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusal

`func (o *RiskRiskSearchReport) SetRefusal(v string)`

SetRefusal sets Refusal field to given value.

### HasRefusal

`func (o *RiskRiskSearchReport) HasRefusal() bool`

HasRefusal returns a boolean if a field has been set.

### GetStarted

`func (o *RiskRiskSearchReport) GetStarted() string`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *RiskRiskSearchReport) GetStartedOk() (*string, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *RiskRiskSearchReport) SetStarted(v string)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *RiskRiskSearchReport) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetTrials

`func (o *RiskRiskSearchReport) GetTrials() []RiskRiskTrial`

GetTrials returns the Trials field if non-nil, zero value otherwise.

### GetTrialsOk

`func (o *RiskRiskSearchReport) GetTrialsOk() (*[]RiskRiskTrial, bool)`

GetTrialsOk returns a tuple with the Trials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrials

`func (o *RiskRiskSearchReport) SetTrials(v []RiskRiskTrial)`

SetTrials sets Trials field to given value.

### HasTrials

`func (o *RiskRiskSearchReport) HasTrials() bool`

HasTrials returns a boolean if a field has been set.

### GetWinner

`func (o *RiskRiskSearchReport) GetWinner() RiskRiskTrial`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *RiskRiskSearchReport) GetWinnerOk() (*RiskRiskTrial, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *RiskRiskSearchReport) SetWinner(v RiskRiskTrial)`

SetWinner sets Winner field to given value.

### HasWinner

`func (o *RiskRiskSearchReport) HasWinner() bool`

HasWinner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


