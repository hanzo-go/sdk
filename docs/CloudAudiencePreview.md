# CloudAudiencePreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Available is false when the roster or the warehouse could not be read; the counts are then zero because nothing was measured, not because the cohort is empty, and Reason says which read failed. | [optional] 
**Count** | Pointer to **int32** | Count is the cohort size: distinct warehouse identifiers for an event audience, mailable customers for an event-less (whole-org) one. | [optional] 
**Deliverable** | Pointer to **int32** | Deliverable is how many de-duplicated addresses a send would reach, and Unmatched how many cohort identifiers named no customer. Unmatched is reported rather than hidden: it is the honest explanation for a cohort of 500 that mails 3. | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Sample** | Pointer to **[]string** | Sample is up to 1000 cohort IDENTIFIERS — never addresses, which product analytics does not hold. Empty for an event-less (whole-org) audience. | [optional] 
**Source** | Pointer to **string** | Source names where the cohort was read: the events table for an event audience, \&quot;iam:&lt;org&gt;\&quot; for the whole-org one. | [optional] 
**Unmatched** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudAudiencePreview

`func NewCloudAudiencePreview() *CloudAudiencePreview`

NewCloudAudiencePreview instantiates a new CloudAudiencePreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAudiencePreviewWithDefaults

`func NewCloudAudiencePreviewWithDefaults() *CloudAudiencePreview`

NewCloudAudiencePreviewWithDefaults instantiates a new CloudAudiencePreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *CloudAudiencePreview) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *CloudAudiencePreview) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *CloudAudiencePreview) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *CloudAudiencePreview) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCount

`func (o *CloudAudiencePreview) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CloudAudiencePreview) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CloudAudiencePreview) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CloudAudiencePreview) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDeliverable

`func (o *CloudAudiencePreview) GetDeliverable() int32`

GetDeliverable returns the Deliverable field if non-nil, zero value otherwise.

### GetDeliverableOk

`func (o *CloudAudiencePreview) GetDeliverableOk() (*int32, bool)`

GetDeliverableOk returns a tuple with the Deliverable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverable

`func (o *CloudAudiencePreview) SetDeliverable(v int32)`

SetDeliverable sets Deliverable field to given value.

### HasDeliverable

`func (o *CloudAudiencePreview) HasDeliverable() bool`

HasDeliverable returns a boolean if a field has been set.

### GetReason

`func (o *CloudAudiencePreview) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CloudAudiencePreview) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CloudAudiencePreview) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CloudAudiencePreview) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSample

`func (o *CloudAudiencePreview) GetSample() []string`

GetSample returns the Sample field if non-nil, zero value otherwise.

### GetSampleOk

`func (o *CloudAudiencePreview) GetSampleOk() (*[]string, bool)`

GetSampleOk returns a tuple with the Sample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSample

`func (o *CloudAudiencePreview) SetSample(v []string)`

SetSample sets Sample field to given value.

### HasSample

`func (o *CloudAudiencePreview) HasSample() bool`

HasSample returns a boolean if a field has been set.

### GetSource

`func (o *CloudAudiencePreview) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CloudAudiencePreview) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CloudAudiencePreview) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CloudAudiencePreview) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUnmatched

`func (o *CloudAudiencePreview) GetUnmatched() int32`

GetUnmatched returns the Unmatched field if non-nil, zero value otherwise.

### GetUnmatchedOk

`func (o *CloudAudiencePreview) GetUnmatchedOk() (*int32, bool)`

GetUnmatchedOk returns a tuple with the Unmatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmatched

`func (o *CloudAudiencePreview) SetUnmatched(v int32)`

SetUnmatched sets Unmatched field to given value.

### HasUnmatched

`func (o *CloudAudiencePreview) HasUnmatched() bool`

HasUnmatched returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


