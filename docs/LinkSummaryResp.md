# LinkSummaryResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**LinkSourceState**](LinkSourceState.md) | Account reports the linked-accounts ledger&#39;s own availability, so a partial answer never fabricates this half. It is scoped to the CALLER: the accounts they linked, metered from each provider&#39;s own login. | [optional] 
**From** | Pointer to **string** | From is when the window opens, RFC 3339 UTC. ONE resolver fixes it for both ledgers, so the account rows and the Hanzo rows always cover the same period — two resolvers could drift and turn the union into a lie. | [optional] 
**Hanzo** | Pointer to [**LinkSourceState**](LinkSourceState.md) | Hanzo reports the same for the Hanzo-routed ledger, which is scoped to the ORG rather than the caller — a different question over the same window. The two are independent: either can be unavailable while the other answers, and Rows then carries only the half that did. | [optional] 
**Range** | Pointer to **string** | Range is the resolved period label. | [optional] 
**Rows** | Pointer to [**[]LinkTotalView**](LinkTotalView.md) | Rows is the union of both ledgers, each row labelled by source and scope — concatenated, NEVER summed: a plan&#39;s percentage is not money. | [optional] 
**To** | Pointer to **string** | To is where the window closes, EXCLUSIVE, RFC 3339 UTC — the instant the read was served. Shared by both ledgers, for the reason From gives. | [optional] 

## Methods

### NewLinkSummaryResp

`func NewLinkSummaryResp() *LinkSummaryResp`

NewLinkSummaryResp instantiates a new LinkSummaryResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkSummaryRespWithDefaults

`func NewLinkSummaryRespWithDefaults() *LinkSummaryResp`

NewLinkSummaryRespWithDefaults instantiates a new LinkSummaryResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *LinkSummaryResp) GetAccount() LinkSourceState`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *LinkSummaryResp) GetAccountOk() (*LinkSourceState, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *LinkSummaryResp) SetAccount(v LinkSourceState)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *LinkSummaryResp) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetFrom

`func (o *LinkSummaryResp) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *LinkSummaryResp) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *LinkSummaryResp) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *LinkSummaryResp) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetHanzo

`func (o *LinkSummaryResp) GetHanzo() LinkSourceState`

GetHanzo returns the Hanzo field if non-nil, zero value otherwise.

### GetHanzoOk

`func (o *LinkSummaryResp) GetHanzoOk() (*LinkSourceState, bool)`

GetHanzoOk returns a tuple with the Hanzo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHanzo

`func (o *LinkSummaryResp) SetHanzo(v LinkSourceState)`

SetHanzo sets Hanzo field to given value.

### HasHanzo

`func (o *LinkSummaryResp) HasHanzo() bool`

HasHanzo returns a boolean if a field has been set.

### GetRange

`func (o *LinkSummaryResp) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *LinkSummaryResp) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *LinkSummaryResp) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *LinkSummaryResp) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetRows

`func (o *LinkSummaryResp) GetRows() []LinkTotalView`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *LinkSummaryResp) GetRowsOk() (*[]LinkTotalView, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *LinkSummaryResp) SetRows(v []LinkTotalView)`

SetRows sets Rows field to given value.

### HasRows

`func (o *LinkSummaryResp) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTo

`func (o *LinkSummaryResp) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *LinkSummaryResp) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *LinkSummaryResp) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *LinkSummaryResp) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


