# AccrualsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**Accruals**](Accruals.md) | Data is what the run did: sources visited, new accruals, royalties alongside. | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewAccrualsOut

`func NewAccrualsOut() *AccrualsOut`

NewAccrualsOut instantiates a new AccrualsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccrualsOutWithDefaults

`func NewAccrualsOutWithDefaults() *AccrualsOut`

NewAccrualsOutWithDefaults instantiates a new AccrualsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AccrualsOut) GetData() Accruals`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AccrualsOut) GetDataOk() (*Accruals, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AccrualsOut) SetData(v Accruals)`

SetData sets Data field to given value.

### HasData

`func (o *AccrualsOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *AccrualsOut) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *AccrualsOut) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *AccrualsOut) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *AccrualsOut) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *AccrualsOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccrualsOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccrualsOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccrualsOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


