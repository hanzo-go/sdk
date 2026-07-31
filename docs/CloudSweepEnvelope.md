# CloudSweepEnvelope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**CloudSweepResult**](CloudSweepResult.md) | Data is the sweep&#39;s counters. | [optional] 
**Msg** | Pointer to **string** | Msg is empty on success; the console surfaces it when status is not \&quot;ok\&quot;. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot; on success. | [optional] 

## Methods

### NewCloudSweepEnvelope

`func NewCloudSweepEnvelope() *CloudSweepEnvelope`

NewCloudSweepEnvelope instantiates a new CloudSweepEnvelope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSweepEnvelopeWithDefaults

`func NewCloudSweepEnvelopeWithDefaults() *CloudSweepEnvelope`

NewCloudSweepEnvelopeWithDefaults instantiates a new CloudSweepEnvelope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudSweepEnvelope) GetData() CloudSweepResult`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudSweepEnvelope) GetDataOk() (*CloudSweepResult, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudSweepEnvelope) SetData(v CloudSweepResult)`

SetData sets Data field to given value.

### HasData

`func (o *CloudSweepEnvelope) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *CloudSweepEnvelope) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CloudSweepEnvelope) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CloudSweepEnvelope) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CloudSweepEnvelope) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *CloudSweepEnvelope) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudSweepEnvelope) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudSweepEnvelope) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudSweepEnvelope) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


