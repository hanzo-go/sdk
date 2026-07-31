# CloudTrailPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CloudWire**](CloudWire.md) | Data is one page of the org&#39;s events, newest first. Empty, never null. | [optional] 
**Msg** | Pointer to **string** | Msg is the envelope&#39;s message slot, empty on success. | [optional] 
**Status** | Pointer to **string** | Status is the envelope&#39;s status slot, \&quot;ok\&quot; on success. | [optional] 
**Total** | Pointer to **int32** | Total is how many events match the filter, across all pages — what a pager needs to size itself. | [optional] 

## Methods

### NewCloudTrailPage

`func NewCloudTrailPage() *CloudTrailPage`

NewCloudTrailPage instantiates a new CloudTrailPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTrailPageWithDefaults

`func NewCloudTrailPageWithDefaults() *CloudTrailPage`

NewCloudTrailPageWithDefaults instantiates a new CloudTrailPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudTrailPage) GetData() []CloudWire`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudTrailPage) GetDataOk() (*[]CloudWire, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudTrailPage) SetData(v []CloudWire)`

SetData sets Data field to given value.

### HasData

`func (o *CloudTrailPage) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *CloudTrailPage) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CloudTrailPage) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CloudTrailPage) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CloudTrailPage) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *CloudTrailPage) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudTrailPage) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudTrailPage) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudTrailPage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotal

`func (o *CloudTrailPage) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CloudTrailPage) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CloudTrailPage) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CloudTrailPage) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


