# CloudPayoutResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**CloudPayoutData**](CloudPayoutData.md) | Data carries the payout and the author. | [optional] 
**Msg** | Pointer to **string** | Msg is the envelope&#39;s message slot, empty on success. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot;. | [optional] 

## Methods

### NewCloudPayoutResult

`func NewCloudPayoutResult() *CloudPayoutResult`

NewCloudPayoutResult instantiates a new CloudPayoutResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPayoutResultWithDefaults

`func NewCloudPayoutResultWithDefaults() *CloudPayoutResult`

NewCloudPayoutResultWithDefaults instantiates a new CloudPayoutResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudPayoutResult) GetData() CloudPayoutData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudPayoutResult) GetDataOk() (*CloudPayoutData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudPayoutResult) SetData(v CloudPayoutData)`

SetData sets Data field to given value.

### HasData

`func (o *CloudPayoutResult) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *CloudPayoutResult) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CloudPayoutResult) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CloudPayoutResult) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CloudPayoutResult) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *CloudPayoutResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudPayoutResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudPayoutResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudPayoutResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


