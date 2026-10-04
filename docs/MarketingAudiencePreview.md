# MarketingAudiencePreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Available is false when the roster or the warehouse could not be read; the counts are then zero because nothing was measured, not because the cohort is empty, and Reason says which read failed. | [optional] 
**Count** | Pointer to **int64** | Count is the cohort size: distinct warehouse identifiers for an event audience, mailable customers for an event-less (whole-org) one. | [optional] 
**Deliverable** | Pointer to **int64** | Deliverable is how many de-duplicated mailboxes a send would reach. Two customers sharing an address count once, so it is &lt;&#x3D; Count. | [optional] 
**Reason** | Pointer to **string** | Reason is the error text of the read that failed: the org&#39;s roster could not be loaded (\&quot;identity store unavailable…\&quot;), or the cohort query had no warehouse to run against (\&quot;analytics warehouse not configured\&quot;). Absent when the evaluation succeeded, so its presence and Available&#x3D;false are one fact seen twice. | [optional] 
**Sample** | Pointer to **[]string** | Sample is up to 1000 cohort IDENTIFIERS — never addresses, which product analytics does not hold. Empty for an event-less (whole-org) audience. | [optional] 
**Source** | Pointer to **string** | Source names where the cohort was read: the events table for an event audience, \&quot;iam:&lt;org&gt;\&quot; for the whole-org one. | [optional] 
**Unmatched** | Pointer to **int64** | Unmatched is how many cohort identifiers named nobody on the org&#39;s roster and so have no address to mail. It is reported rather than hidden: it is the honest explanation for a cohort of 500 that mails 3. Always 0 for an event-less (whole-org) audience, which starts from the roster and has nothing to match. | [optional] 

## Methods

### NewMarketingAudiencePreview

`func NewMarketingAudiencePreview() *MarketingAudiencePreview`

NewMarketingAudiencePreview instantiates a new MarketingAudiencePreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingAudiencePreviewWithDefaults

`func NewMarketingAudiencePreviewWithDefaults() *MarketingAudiencePreview`

NewMarketingAudiencePreviewWithDefaults instantiates a new MarketingAudiencePreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *MarketingAudiencePreview) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *MarketingAudiencePreview) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *MarketingAudiencePreview) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *MarketingAudiencePreview) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCount

`func (o *MarketingAudiencePreview) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *MarketingAudiencePreview) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *MarketingAudiencePreview) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *MarketingAudiencePreview) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDeliverable

`func (o *MarketingAudiencePreview) GetDeliverable() int64`

GetDeliverable returns the Deliverable field if non-nil, zero value otherwise.

### GetDeliverableOk

`func (o *MarketingAudiencePreview) GetDeliverableOk() (*int64, bool)`

GetDeliverableOk returns a tuple with the Deliverable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverable

`func (o *MarketingAudiencePreview) SetDeliverable(v int64)`

SetDeliverable sets Deliverable field to given value.

### HasDeliverable

`func (o *MarketingAudiencePreview) HasDeliverable() bool`

HasDeliverable returns a boolean if a field has been set.

### GetReason

`func (o *MarketingAudiencePreview) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MarketingAudiencePreview) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MarketingAudiencePreview) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *MarketingAudiencePreview) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSample

`func (o *MarketingAudiencePreview) GetSample() []string`

GetSample returns the Sample field if non-nil, zero value otherwise.

### GetSampleOk

`func (o *MarketingAudiencePreview) GetSampleOk() (*[]string, bool)`

GetSampleOk returns a tuple with the Sample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSample

`func (o *MarketingAudiencePreview) SetSample(v []string)`

SetSample sets Sample field to given value.

### HasSample

`func (o *MarketingAudiencePreview) HasSample() bool`

HasSample returns a boolean if a field has been set.

### GetSource

`func (o *MarketingAudiencePreview) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MarketingAudiencePreview) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MarketingAudiencePreview) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *MarketingAudiencePreview) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUnmatched

`func (o *MarketingAudiencePreview) GetUnmatched() int64`

GetUnmatched returns the Unmatched field if non-nil, zero value otherwise.

### GetUnmatchedOk

`func (o *MarketingAudiencePreview) GetUnmatchedOk() (*int64, bool)`

GetUnmatchedOk returns a tuple with the Unmatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatched

`func (o *MarketingAudiencePreview) SetUnmatched(v int64)`

SetUnmatched sets Unmatched field to given value.

### HasUnmatched

`func (o *MarketingAudiencePreview) HasUnmatched() bool`

HasUnmatched returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


