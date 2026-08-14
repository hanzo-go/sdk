# SubscriptionsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]SubscriptionRow**](SubscriptionRow.md) |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewSubscriptionsOut

`func NewSubscriptionsOut() *SubscriptionsOut`

NewSubscriptionsOut instantiates a new SubscriptionsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionsOutWithDefaults

`func NewSubscriptionsOutWithDefaults() *SubscriptionsOut`

NewSubscriptionsOutWithDefaults instantiates a new SubscriptionsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SubscriptionsOut) GetData() []SubscriptionRow`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SubscriptionsOut) GetDataOk() (*[]SubscriptionRow, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SubscriptionsOut) SetData(v []SubscriptionRow)`

SetData sets Data field to given value.

### HasData

`func (o *SubscriptionsOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *SubscriptionsOut) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *SubscriptionsOut) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *SubscriptionsOut) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *SubscriptionsOut) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *SubscriptionsOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionsOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionsOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubscriptionsOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotal

`func (o *SubscriptionsOut) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SubscriptionsOut) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SubscriptionsOut) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SubscriptionsOut) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


