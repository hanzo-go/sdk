# AdminAdminListAudit200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]AdminAuditRow**](AdminAuditRow.md) |  | [optional] 
**Data2** | Pointer to **int32** |  | [optional] 
**Integrity** | Pointer to [**AdminAuditIntegrity**](AdminAuditIntegrity.md) |  | [optional] 

## Methods

### NewAdminAdminListAudit200Response

`func NewAdminAdminListAudit200Response() *AdminAdminListAudit200Response`

NewAdminAdminListAudit200Response instantiates a new AdminAdminListAudit200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminAdminListAudit200ResponseWithDefaults

`func NewAdminAdminListAudit200ResponseWithDefaults() *AdminAdminListAudit200Response`

NewAdminAdminListAudit200ResponseWithDefaults instantiates a new AdminAdminListAudit200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AdminAdminListAudit200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminAdminListAudit200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminAdminListAudit200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdminAdminListAudit200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMsg

`func (o *AdminAdminListAudit200Response) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *AdminAdminListAudit200Response) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *AdminAdminListAudit200Response) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *AdminAdminListAudit200Response) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetData

`func (o *AdminAdminListAudit200Response) GetData() []AdminAuditRow`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AdminAdminListAudit200Response) GetDataOk() (*[]AdminAuditRow, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AdminAdminListAudit200Response) SetData(v []AdminAuditRow)`

SetData sets Data field to given value.

### HasData

`func (o *AdminAdminListAudit200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetData2

`func (o *AdminAdminListAudit200Response) GetData2() int32`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *AdminAdminListAudit200Response) GetData2Ok() (*int32, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *AdminAdminListAudit200Response) SetData2(v int32)`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *AdminAdminListAudit200Response) HasData2() bool`

HasData2 returns a boolean if a field has been set.

### GetIntegrity

`func (o *AdminAdminListAudit200Response) GetIntegrity() AdminAuditIntegrity`

GetIntegrity returns the Integrity field if non-nil, zero value otherwise.

### GetIntegrityOk

`func (o *AdminAdminListAudit200Response) GetIntegrityOk() (*AdminAuditIntegrity, bool)`

GetIntegrityOk returns a tuple with the Integrity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrity

`func (o *AdminAdminListAudit200Response) SetIntegrity(v AdminAuditIntegrity)`

SetIntegrity sets Integrity field to given value.

### HasIntegrity

`func (o *AdminAdminListAudit200Response) HasIntegrity() bool`

HasIntegrity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


