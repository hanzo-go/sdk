# O11yOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**O11yGlobal**](O11yGlobal.md) |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yOut

`func NewO11yOut() *O11yOut`

NewO11yOut instantiates a new O11yOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yOutWithDefaults

`func NewO11yOutWithDefaults() *O11yOut`

NewO11yOutWithDefaults instantiates a new O11yOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *O11yOut) GetData() O11yGlobal`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yOut) GetDataOk() (*O11yGlobal, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yOut) SetData(v O11yGlobal)`

SetData sets Data field to given value.

### HasData

`func (o *O11yOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *O11yOut) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *O11yOut) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *O11yOut) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *O11yOut) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *O11yOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


