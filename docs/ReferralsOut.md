# ReferralsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ReferralBoard**](ReferralBoard.md) | Data is the referral board: leaders, funnel, tally and per-level liability. | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewReferralsOut

`func NewReferralsOut() *ReferralsOut`

NewReferralsOut instantiates a new ReferralsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferralsOutWithDefaults

`func NewReferralsOutWithDefaults() *ReferralsOut`

NewReferralsOutWithDefaults instantiates a new ReferralsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ReferralsOut) GetData() ReferralBoard`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ReferralsOut) GetDataOk() (*ReferralBoard, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ReferralsOut) SetData(v ReferralBoard)`

SetData sets Data field to given value.

### HasData

`func (o *ReferralsOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *ReferralsOut) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ReferralsOut) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ReferralsOut) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *ReferralsOut) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *ReferralsOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReferralsOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReferralsOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReferralsOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


