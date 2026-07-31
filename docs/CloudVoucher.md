# CloudVoucher

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description is the human line for the event, e.g. the vendor a bill came from. | [optional] 
**Legs** | Pointer to [**[]CloudLeg**](CloudLeg.md) | Legs are the sides of the posting. They must balance: Σdebit &#x3D;&#x3D; Σcredit, give or take the 2¢ round-off allowance. | [optional] 
**PostingAt** | Pointer to **string** | PostingAt is the RFC3339 instant the event posts at — the time every statement window filters on. | [optional] 
**SourceId** | Pointer to **string** | SourceID is the source event&#39;s own id within that namespace. Together with SourceKind it is the key that makes a repeat posting a no-op. | [optional] 
**SourceKind** | Pointer to **string** | SourceKind is the idempotency namespace naming what booked this, e.g. \&quot;scan\&quot;. | [optional] 

## Methods

### NewCloudVoucher

`func NewCloudVoucher() *CloudVoucher`

NewCloudVoucher instantiates a new CloudVoucher object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudVoucherWithDefaults

`func NewCloudVoucherWithDefaults() *CloudVoucher`

NewCloudVoucherWithDefaults instantiates a new CloudVoucher object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CloudVoucher) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudVoucher) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudVoucher) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudVoucher) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLegs

`func (o *CloudVoucher) GetLegs() []CloudLeg`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *CloudVoucher) GetLegsOk() (*[]CloudLeg, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *CloudVoucher) SetLegs(v []CloudLeg)`

SetLegs sets Legs field to given value.

### HasLegs

`func (o *CloudVoucher) HasLegs() bool`

HasLegs returns a boolean if a field has been set.

### GetPostingAt

`func (o *CloudVoucher) GetPostingAt() string`

GetPostingAt returns the PostingAt field if non-nil, zero value otherwise.

### GetPostingAtOk

`func (o *CloudVoucher) GetPostingAtOk() (*string, bool)`

GetPostingAtOk returns a tuple with the PostingAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostingAt

`func (o *CloudVoucher) SetPostingAt(v string)`

SetPostingAt sets PostingAt field to given value.

### HasPostingAt

`func (o *CloudVoucher) HasPostingAt() bool`

HasPostingAt returns a boolean if a field has been set.

### GetSourceId

`func (o *CloudVoucher) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *CloudVoucher) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *CloudVoucher) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *CloudVoucher) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceKind

`func (o *CloudVoucher) GetSourceKind() string`

GetSourceKind returns the SourceKind field if non-nil, zero value otherwise.

### GetSourceKindOk

`func (o *CloudVoucher) GetSourceKindOk() (*string, bool)`

GetSourceKindOk returns a tuple with the SourceKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceKind

`func (o *CloudVoucher) SetSourceKind(v string)`

SetSourceKind sets SourceKind field to given value.

### HasSourceKind

`func (o *CloudVoucher) HasSourceKind() bool`

HasSourceKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


