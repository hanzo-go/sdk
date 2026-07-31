# CloudRecordsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Integrity** | Pointer to [**CloudIntegrity**](CloudIntegrity.md) |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudRecordsOut

`func NewCloudRecordsOut() *CloudRecordsOut`

NewCloudRecordsOut instantiates a new CloudRecordsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRecordsOutWithDefaults

`func NewCloudRecordsOutWithDefaults() *CloudRecordsOut`

NewCloudRecordsOutWithDefaults instantiates a new CloudRecordsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudRecordsOut) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudRecordsOut) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudRecordsOut) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *CloudRecordsOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetIntegrity

`func (o *CloudRecordsOut) GetIntegrity() CloudIntegrity`

GetIntegrity returns the Integrity field if non-nil, zero value otherwise.

### GetIntegrityOk

`func (o *CloudRecordsOut) GetIntegrityOk() (*CloudIntegrity, bool)`

GetIntegrityOk returns a tuple with the Integrity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrity

`func (o *CloudRecordsOut) SetIntegrity(v CloudIntegrity)`

SetIntegrity sets Integrity field to given value.

### HasIntegrity

`func (o *CloudRecordsOut) HasIntegrity() bool`

HasIntegrity returns a boolean if a field has been set.

### GetMsg

`func (o *CloudRecordsOut) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CloudRecordsOut) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CloudRecordsOut) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CloudRecordsOut) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *CloudRecordsOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudRecordsOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudRecordsOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudRecordsOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotal

`func (o *CloudRecordsOut) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CloudRecordsOut) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CloudRecordsOut) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CloudRecordsOut) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


