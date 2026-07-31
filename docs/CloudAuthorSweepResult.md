# CloudAuthorSweepResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**CloudSweepCounts**](CloudSweepCounts.md) | Data is the sweep&#39;s counts. | [optional] 
**Msg** | Pointer to **string** | Msg is the envelope&#39;s message slot, empty on success. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot;. | [optional] 

## Methods

### NewCloudAuthorSweepResult

`func NewCloudAuthorSweepResult() *CloudAuthorSweepResult`

NewCloudAuthorSweepResult instantiates a new CloudAuthorSweepResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAuthorSweepResultWithDefaults

`func NewCloudAuthorSweepResultWithDefaults() *CloudAuthorSweepResult`

NewCloudAuthorSweepResultWithDefaults instantiates a new CloudAuthorSweepResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudAuthorSweepResult) GetData() CloudSweepCounts`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudAuthorSweepResult) GetDataOk() (*CloudSweepCounts, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudAuthorSweepResult) SetData(v CloudSweepCounts)`

SetData sets Data field to given value.

### HasData

`func (o *CloudAuthorSweepResult) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *CloudAuthorSweepResult) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CloudAuthorSweepResult) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CloudAuthorSweepResult) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CloudAuthorSweepResult) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *CloudAuthorSweepResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudAuthorSweepResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudAuthorSweepResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudAuthorSweepResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


