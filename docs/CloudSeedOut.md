# CloudSeedOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**CloudSeedData**](CloudSeedData.md) | Data is the injection result. | [optional] 
**Msg** | Pointer to **string** | Msg carries an operator-facing note; empty on success. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot; on success. | [optional] 

## Methods

### NewCloudSeedOut

`func NewCloudSeedOut() *CloudSeedOut`

NewCloudSeedOut instantiates a new CloudSeedOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSeedOutWithDefaults

`func NewCloudSeedOutWithDefaults() *CloudSeedOut`

NewCloudSeedOutWithDefaults instantiates a new CloudSeedOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudSeedOut) GetData() CloudSeedData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudSeedOut) GetDataOk() (*CloudSeedData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudSeedOut) SetData(v CloudSeedData)`

SetData sets Data field to given value.

### HasData

`func (o *CloudSeedOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *CloudSeedOut) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CloudSeedOut) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CloudSeedOut) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CloudSeedOut) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *CloudSeedOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudSeedOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudSeedOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudSeedOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


