# CloudSyncOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**CloudSyncStarted**](CloudSyncStarted.md) |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudSyncOut

`func NewCloudSyncOut() *CloudSyncOut`

NewCloudSyncOut instantiates a new CloudSyncOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSyncOutWithDefaults

`func NewCloudSyncOutWithDefaults() *CloudSyncOut`

NewCloudSyncOutWithDefaults instantiates a new CloudSyncOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudSyncOut) GetData() CloudSyncStarted`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudSyncOut) GetDataOk() (*CloudSyncStarted, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudSyncOut) SetData(v CloudSyncStarted)`

SetData sets Data field to given value.

### HasData

`func (o *CloudSyncOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *CloudSyncOut) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CloudSyncOut) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CloudSyncOut) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CloudSyncOut) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *CloudSyncOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudSyncOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudSyncOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudSyncOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


