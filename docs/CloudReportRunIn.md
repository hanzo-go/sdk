# CloudReportRunIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to **string** | Branch, CommitSha and Diffstat describe what the run produced; Error is the failure when OK is false. Each is clamped, never rejected. | [optional] 
**Changed** | Pointer to **bool** |  | [optional] 
**CommitSha** | Pointer to **string** |  | [optional] 
**Diffstat** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | ID is the machine reporting, from the path. | [optional] 
**Ok** | Pointer to **bool** | OK is whether the run succeeded; Changed whether it produced any commit. | [optional] 
**RunId** | Pointer to **string** | RunID is the routed run being completed, from the path. | [optional] 

## Methods

### NewCloudReportRunIn

`func NewCloudReportRunIn() *CloudReportRunIn`

NewCloudReportRunIn instantiates a new CloudReportRunIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudReportRunInWithDefaults

`func NewCloudReportRunInWithDefaults() *CloudReportRunIn`

NewCloudReportRunInWithDefaults instantiates a new CloudReportRunIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *CloudReportRunIn) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *CloudReportRunIn) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *CloudReportRunIn) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *CloudReportRunIn) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetChanged

`func (o *CloudReportRunIn) GetChanged() bool`

GetChanged returns the Changed field if non-nil, zero value otherwise.

### GetChangedOk

`func (o *CloudReportRunIn) GetChangedOk() (*bool, bool)`

GetChangedOk returns a tuple with the Changed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanged

`func (o *CloudReportRunIn) SetChanged(v bool)`

SetChanged sets Changed field to given value.

### HasChanged

`func (o *CloudReportRunIn) HasChanged() bool`

HasChanged returns a boolean if a field has been set.

### GetCommitSha

`func (o *CloudReportRunIn) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *CloudReportRunIn) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *CloudReportRunIn) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.

### HasCommitSha

`func (o *CloudReportRunIn) HasCommitSha() bool`

HasCommitSha returns a boolean if a field has been set.

### GetDiffstat

`func (o *CloudReportRunIn) GetDiffstat() string`

GetDiffstat returns the Diffstat field if non-nil, zero value otherwise.

### GetDiffstatOk

`func (o *CloudReportRunIn) GetDiffstatOk() (*string, bool)`

GetDiffstatOk returns a tuple with the Diffstat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffstat

`func (o *CloudReportRunIn) SetDiffstat(v string)`

SetDiffstat sets Diffstat field to given value.

### HasDiffstat

`func (o *CloudReportRunIn) HasDiffstat() bool`

HasDiffstat returns a boolean if a field has been set.

### GetError

`func (o *CloudReportRunIn) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CloudReportRunIn) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CloudReportRunIn) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CloudReportRunIn) HasError() bool`

HasError returns a boolean if a field has been set.

### GetId

`func (o *CloudReportRunIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudReportRunIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudReportRunIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudReportRunIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOk

`func (o *CloudReportRunIn) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *CloudReportRunIn) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *CloudReportRunIn) SetOk(v bool)`

SetOk sets Ok field to given value.

### HasOk

`func (o *CloudReportRunIn) HasOk() bool`

HasOk returns a boolean if a field has been set.

### GetRunId

`func (o *CloudReportRunIn) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *CloudReportRunIn) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *CloudReportRunIn) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *CloudReportRunIn) HasRunId() bool`

HasRunId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


