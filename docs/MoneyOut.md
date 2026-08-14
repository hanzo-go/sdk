# MoneyOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**MoneyBoard**](MoneyBoard.md) |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewMoneyOut

`func NewMoneyOut() *MoneyOut`

NewMoneyOut instantiates a new MoneyOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoneyOutWithDefaults

`func NewMoneyOutWithDefaults() *MoneyOut`

NewMoneyOutWithDefaults instantiates a new MoneyOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MoneyOut) GetData() MoneyBoard`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MoneyOut) GetDataOk() (*MoneyBoard, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MoneyOut) SetData(v MoneyBoard)`

SetData sets Data field to given value.

### HasData

`func (o *MoneyOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *MoneyOut) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *MoneyOut) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *MoneyOut) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *MoneyOut) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *MoneyOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MoneyOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MoneyOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MoneyOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


