# CloudWaitlistBoostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email identifies the entry to boost. Either this or RefCode is required. | [optional] 
**Points** | Pointer to **int32** | Points is how many points to award. Must be positive — this seam exists to move someone UP toward the cutoff. | [optional] 
**Reason** | Pointer to **string** | Reason is the operator&#39;s justification. Not sent to the engine; it is recorded on the audit row, which is the point of asking for it. | [optional] 
**RefCode** | Pointer to **string** | RefCode identifies the entry by its referral code, when the email is unknown. | [optional] 
**Waitlist** | Pointer to **string** | Waitlist is the waitlist slug the grant lands on. Required. | [optional] 

## Methods

### NewCloudWaitlistBoostRequest

`func NewCloudWaitlistBoostRequest() *CloudWaitlistBoostRequest`

NewCloudWaitlistBoostRequest instantiates a new CloudWaitlistBoostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudWaitlistBoostRequestWithDefaults

`func NewCloudWaitlistBoostRequestWithDefaults() *CloudWaitlistBoostRequest`

NewCloudWaitlistBoostRequestWithDefaults instantiates a new CloudWaitlistBoostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CloudWaitlistBoostRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CloudWaitlistBoostRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CloudWaitlistBoostRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CloudWaitlistBoostRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPoints

`func (o *CloudWaitlistBoostRequest) GetPoints() int32`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *CloudWaitlistBoostRequest) GetPointsOk() (*int32, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *CloudWaitlistBoostRequest) SetPoints(v int32)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *CloudWaitlistBoostRequest) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetReason

`func (o *CloudWaitlistBoostRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CloudWaitlistBoostRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CloudWaitlistBoostRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CloudWaitlistBoostRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRefCode

`func (o *CloudWaitlistBoostRequest) GetRefCode() string`

GetRefCode returns the RefCode field if non-nil, zero value otherwise.

### GetRefCodeOk

`func (o *CloudWaitlistBoostRequest) GetRefCodeOk() (*string, bool)`

GetRefCodeOk returns a tuple with the RefCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefCode

`func (o *CloudWaitlistBoostRequest) SetRefCode(v string)`

SetRefCode sets RefCode field to given value.

### HasRefCode

`func (o *CloudWaitlistBoostRequest) HasRefCode() bool`

HasRefCode returns a boolean if a field has been set.

### GetWaitlist

`func (o *CloudWaitlistBoostRequest) GetWaitlist() string`

GetWaitlist returns the Waitlist field if non-nil, zero value otherwise.

### GetWaitlistOk

`func (o *CloudWaitlistBoostRequest) GetWaitlistOk() (*string, bool)`

GetWaitlistOk returns a tuple with the Waitlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitlist

`func (o *CloudWaitlistBoostRequest) SetWaitlist(v string)`

SetWaitlist sets Waitlist field to given value.

### HasWaitlist

`func (o *CloudWaitlistBoostRequest) HasWaitlist() bool`

HasWaitlist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


