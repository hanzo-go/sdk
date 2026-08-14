# AdminReportOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**AdminReportData**](AdminReportData.md) | Data is the treasury board. | [optional] 
**Msg** | Pointer to **string** | Msg carries an operator-facing note; empty on success. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot; on success; the transport maps a non-ok envelope to an error. | [optional] 

## Methods

### NewAdminReportOut

`func NewAdminReportOut() *AdminReportOut`

NewAdminReportOut instantiates a new AdminReportOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminReportOutWithDefaults

`func NewAdminReportOutWithDefaults() *AdminReportOut`

NewAdminReportOutWithDefaults instantiates a new AdminReportOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AdminReportOut) GetData() AdminReportData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AdminReportOut) GetDataOk() (*AdminReportData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AdminReportOut) SetData(v AdminReportData)`

SetData sets Data field to given value.

### HasData

`func (o *AdminReportOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *AdminReportOut) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *AdminReportOut) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *AdminReportOut) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *AdminReportOut) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *AdminReportOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminReportOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminReportOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdminReportOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


