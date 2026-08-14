# SummaryResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**SourceState**](SourceState.md) | Account and Hanzo report each ledger&#39;s own availability, so a partial warehouse never fabricates the other half. | [optional] 
**From** | Pointer to **string** | From and To are the one [from, to) window BOTH halves resolved, RFC 3339 UTC. | [optional] 
**Hanzo** | Pointer to [**SourceState**](SourceState.md) |  | [optional] 
**Range** | Pointer to **string** | Range is the resolved period label. | [optional] 
**Rows** | Pointer to [**[]TotalView**](TotalView.md) | Rows is the union of both ledgers, each row labelled by source and scope — concatenated, NEVER summed: a plan&#39;s percentage is not money. | [optional] 
**To** | Pointer to **string** |  | [optional] 

## Methods

### NewSummaryResp

`func NewSummaryResp() *SummaryResp`

NewSummaryResp instantiates a new SummaryResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryRespWithDefaults

`func NewSummaryRespWithDefaults() *SummaryResp`

NewSummaryRespWithDefaults instantiates a new SummaryResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *SummaryResp) GetAccount() SourceState`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SummaryResp) GetAccountOk() (*SourceState, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SummaryResp) SetAccount(v SourceState)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *SummaryResp) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetFrom

`func (o *SummaryResp) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SummaryResp) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SummaryResp) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *SummaryResp) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetHanzo

`func (o *SummaryResp) GetHanzo() SourceState`

GetHanzo returns the Hanzo field if non-nil, zero value otherwise.

### GetHanzoOk

`func (o *SummaryResp) GetHanzoOk() (*SourceState, bool)`

GetHanzoOk returns a tuple with the Hanzo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHanzo

`func (o *SummaryResp) SetHanzo(v SourceState)`

SetHanzo sets Hanzo field to given value.

### HasHanzo

`func (o *SummaryResp) HasHanzo() bool`

HasHanzo returns a boolean if a field has been set.

### GetRange

`func (o *SummaryResp) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *SummaryResp) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *SummaryResp) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *SummaryResp) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetRows

`func (o *SummaryResp) GetRows() []TotalView`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *SummaryResp) GetRowsOk() (*[]TotalView, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *SummaryResp) SetRows(v []TotalView)`

SetRows sets Rows field to given value.

### HasRows

`func (o *SummaryResp) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTo

`func (o *SummaryResp) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SummaryResp) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SummaryResp) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *SummaryResp) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


