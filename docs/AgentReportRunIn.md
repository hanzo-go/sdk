# AgentReportRunIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to **string** | Branch, CommitSha and Diffstat describe what the run produced; Error is the failure when OK is false. Each is clamped, never rejected. | [optional] 
**Changed** | Pointer to **bool** | Changed says whether the run produced any commit. It is INDEPENDENT of OK: a run can succeed and change nothing (there was nothing to do), and a run can fail after committing some of its work. Two questions, two booleans. | [optional] 
**CommitSha** | Pointer to **string** | CommitSha is the tip the run pushed, clamped to 128 characters. Empty when it pushed nothing, which is the same case Changed reports false for. | [optional] 
**Diffstat** | Pointer to **string** | Diffstat is the run&#39;s own summary of what it changed, as text, clamped to 64 KiB. Free-form: it is shown, never parsed. | [optional] 
**Error** | Pointer to **string** | Error is why the run failed, clamped to 64 KiB. It is CLAMPED rather than refused — a truncated reason is worth more than a rejected report, because a rejected report leaves the durable workflow waiting forever. | [optional] 
**Id** | Pointer to **string** | ID is the machine reporting, from the path. | [optional] 
**Ok** | Pointer to **bool** | OK is whether the run succeeded; Changed whether it produced any commit. | [optional] 
**RunId** | Pointer to **string** | RunID is the routed run being completed, from the path. | [optional] 

## Methods

### NewAgentReportRunIn

`func NewAgentReportRunIn() *AgentReportRunIn`

NewAgentReportRunIn instantiates a new AgentReportRunIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentReportRunInWithDefaults

`func NewAgentReportRunInWithDefaults() *AgentReportRunIn`

NewAgentReportRunInWithDefaults instantiates a new AgentReportRunIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *AgentReportRunIn) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *AgentReportRunIn) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *AgentReportRunIn) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *AgentReportRunIn) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetChanged

`func (o *AgentReportRunIn) GetChanged() bool`

GetChanged returns the Changed field if non-nil, zero value otherwise.

### GetChangedOk

`func (o *AgentReportRunIn) GetChangedOk() (*bool, bool)`

GetChangedOk returns a tuple with the Changed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanged

`func (o *AgentReportRunIn) SetChanged(v bool)`

SetChanged sets Changed field to given value.

### HasChanged

`func (o *AgentReportRunIn) HasChanged() bool`

HasChanged returns a boolean if a field has been set.

### GetCommitSha

`func (o *AgentReportRunIn) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *AgentReportRunIn) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *AgentReportRunIn) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.

### HasCommitSha

`func (o *AgentReportRunIn) HasCommitSha() bool`

HasCommitSha returns a boolean if a field has been set.

### GetDiffstat

`func (o *AgentReportRunIn) GetDiffstat() string`

GetDiffstat returns the Diffstat field if non-nil, zero value otherwise.

### GetDiffstatOk

`func (o *AgentReportRunIn) GetDiffstatOk() (*string, bool)`

GetDiffstatOk returns a tuple with the Diffstat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffstat

`func (o *AgentReportRunIn) SetDiffstat(v string)`

SetDiffstat sets Diffstat field to given value.

### HasDiffstat

`func (o *AgentReportRunIn) HasDiffstat() bool`

HasDiffstat returns a boolean if a field has been set.

### GetError

`func (o *AgentReportRunIn) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AgentReportRunIn) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AgentReportRunIn) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AgentReportRunIn) HasError() bool`

HasError returns a boolean if a field has been set.

### GetId

`func (o *AgentReportRunIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentReportRunIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentReportRunIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentReportRunIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOk

`func (o *AgentReportRunIn) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *AgentReportRunIn) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *AgentReportRunIn) SetOk(v bool)`

SetOk sets Ok field to given value.

### HasOk

`func (o *AgentReportRunIn) HasOk() bool`

HasOk returns a boolean if a field has been set.

### GetRunId

`func (o *AgentReportRunIn) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *AgentReportRunIn) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *AgentReportRunIn) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *AgentReportRunIn) HasRunId() bool`

HasRunId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


