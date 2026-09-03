# TrailPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Wire**](Wire.md) | Data is one page of the org&#39;s events, newest first. Empty, never null. | [optional] 
**Msg** | Pointer to **string** | Msg is the envelope&#39;s message slot, empty on success. | [optional] 
**Status** | Pointer to **string** | Status is the envelope&#39;s status slot, \&quot;ok\&quot; on success. | [optional] 
**Total** | Pointer to **int64** | Total is how many events match the filter, across all pages — what a pager needs to size itself. | [optional] 

## Methods

### NewTrailPage

`func NewTrailPage() *TrailPage`

NewTrailPage instantiates a new TrailPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrailPageWithDefaults

`func NewTrailPageWithDefaults() *TrailPage`

NewTrailPageWithDefaults instantiates a new TrailPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TrailPage) GetData() []Wire`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TrailPage) GetDataOk() (*[]Wire, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TrailPage) SetData(v []Wire)`

SetData sets Data field to given value.

### HasData

`func (o *TrailPage) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *TrailPage) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *TrailPage) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *TrailPage) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *TrailPage) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *TrailPage) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TrailPage) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TrailPage) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TrailPage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotal

`func (o *TrailPage) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TrailPage) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TrailPage) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *TrailPage) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


