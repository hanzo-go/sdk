# CodeView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clicks** | Pointer to **int64** | Clicks is how many pings this code has taken. The one STORED counter here and pure vanity: no accrual or payout reads it, pings are coalesced in memory and flushed in batches, and a dropped tally is accepted rather than contending with the money write path. Do not reconcile it against anything. | [optional] 
**Code** | Pointer to **string** | Code is the link&#39;s slug — 3–32 chars of a–z, 0–9 and hyphen — unique across the WHOLE directory, so any affiliate&#39;s code resolves an attribution. | [optional] 
**Conversions** | Pointer to **int64** | Conversions is how many of those signups have actually produced positive commission for the caller. Also derived, from the accrual rows, so it is ≤ signups and lags a referral until the first sweep after it spends. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is when the link was minted, Unix seconds UTC. | [optional] 
**Label** | Pointer to **string** | Label is the caller&#39;s own note for the link (\&quot;twitter\&quot;, \&quot;newsletter\&quot;). Cosmetic: trimmed, stripped of control characters, capped at 48 bytes, and never part of the code. \&quot;primary\&quot; on the link mirrored at approval. | [optional] 
**Signups** | Pointer to **int64** | Signups is how many orgs were attributed with this code — DERIVED by counting attribution edges, never stored, so it cannot drift from the ledger. | [optional] 
**Url** | Pointer to **string** | URL is the full shareable link, the brand host plus ?aff&#x3D;&lt;code&gt;. The host is the deployment&#39;s own brand, so a Lux or Zoo install never mints a hanzo.ai link. | [optional] 

## Methods

### NewCodeView

`func NewCodeView() *CodeView`

NewCodeView instantiates a new CodeView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeViewWithDefaults

`func NewCodeViewWithDefaults() *CodeView`

NewCodeViewWithDefaults instantiates a new CodeView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClicks

`func (o *CodeView) GetClicks() int64`

GetClicks returns the Clicks field if non-nil, zero value otherwise.

### GetClicksOk

`func (o *CodeView) GetClicksOk() (*int64, bool)`

GetClicksOk returns a tuple with the Clicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicks

`func (o *CodeView) SetClicks(v int64)`

SetClicks sets Clicks field to given value.

### HasClicks

`func (o *CodeView) HasClicks() bool`

HasClicks returns a boolean if a field has been set.

### GetCode

`func (o *CodeView) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CodeView) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CodeView) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CodeView) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetConversions

`func (o *CodeView) GetConversions() int64`

GetConversions returns the Conversions field if non-nil, zero value otherwise.

### GetConversionsOk

`func (o *CodeView) GetConversionsOk() (*int64, bool)`

GetConversionsOk returns a tuple with the Conversions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversions

`func (o *CodeView) SetConversions(v int64)`

SetConversions sets Conversions field to given value.

### HasConversions

`func (o *CodeView) HasConversions() bool`

HasConversions returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CodeView) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CodeView) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CodeView) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CodeView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLabel

`func (o *CodeView) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CodeView) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CodeView) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CodeView) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetSignups

`func (o *CodeView) GetSignups() int64`

GetSignups returns the Signups field if non-nil, zero value otherwise.

### GetSignupsOk

`func (o *CodeView) GetSignupsOk() (*int64, bool)`

GetSignupsOk returns a tuple with the Signups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignups

`func (o *CodeView) SetSignups(v int64)`

SetSignups sets Signups field to given value.

### HasSignups

`func (o *CodeView) HasSignups() bool`

HasSignups returns a boolean if a field has been set.

### GetUrl

`func (o *CodeView) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CodeView) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CodeView) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CodeView) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


