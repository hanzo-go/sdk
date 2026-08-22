# IssueHits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Count is how many rows Issues carries — the size of THIS answer after the cap, not how many issues matched. A count equal to the limit means there are probably more; there is no total and no cursor. | [optional] 
**Issues** | Pointer to [**[]IssueHit**](IssueHit.md) | Issues are the matching rows grouped by status and oldest-first within a group, capped by the search&#39;s limit (50 by default, 200 at most). The cap is applied to that order, so a broad search returns the head of it rather than a sample. | [optional] 

## Methods

### NewIssueHits

`func NewIssueHits() *IssueHits`

NewIssueHits instantiates a new IssueHits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueHitsWithDefaults

`func NewIssueHitsWithDefaults() *IssueHits`

NewIssueHitsWithDefaults instantiates a new IssueHits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *IssueHits) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *IssueHits) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *IssueHits) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *IssueHits) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetIssues

`func (o *IssueHits) GetIssues() []IssueHit`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *IssueHits) GetIssuesOk() (*[]IssueHit, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *IssueHits) SetIssues(v []IssueHit)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *IssueHits) HasIssues() bool`

HasIssues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


