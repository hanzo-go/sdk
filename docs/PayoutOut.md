# PayoutOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**Settlement**](Settlement.md) | Data is the recorded payout and the balances it left behind. | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewPayoutOut

`func NewPayoutOut() *PayoutOut`

NewPayoutOut instantiates a new PayoutOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutOutWithDefaults

`func NewPayoutOutWithDefaults() *PayoutOut`

NewPayoutOutWithDefaults instantiates a new PayoutOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PayoutOut) GetData() Settlement`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PayoutOut) GetDataOk() (*Settlement, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PayoutOut) SetData(v Settlement)`

SetData sets Data field to given value.

### HasData

`func (o *PayoutOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *PayoutOut) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *PayoutOut) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *PayoutOut) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *PayoutOut) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *PayoutOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PayoutOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PayoutOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PayoutOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


