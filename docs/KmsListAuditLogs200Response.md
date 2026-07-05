# KmsListAuditLogs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditLogs** | Pointer to [**[]KmsAuditLog**](KmsAuditLog.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewKmsListAuditLogs200Response

`func NewKmsListAuditLogs200Response() *KmsListAuditLogs200Response`

NewKmsListAuditLogs200Response instantiates a new KmsListAuditLogs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsListAuditLogs200ResponseWithDefaults

`func NewKmsListAuditLogs200ResponseWithDefaults() *KmsListAuditLogs200Response`

NewKmsListAuditLogs200ResponseWithDefaults instantiates a new KmsListAuditLogs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditLogs

`func (o *KmsListAuditLogs200Response) GetAuditLogs() []KmsAuditLog`

GetAuditLogs returns the AuditLogs field if non-nil, zero value otherwise.

### GetAuditLogsOk

`func (o *KmsListAuditLogs200Response) GetAuditLogsOk() (*[]KmsAuditLog, bool)`

GetAuditLogsOk returns a tuple with the AuditLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogs

`func (o *KmsListAuditLogs200Response) SetAuditLogs(v []KmsAuditLog)`

SetAuditLogs sets AuditLogs field to given value.

### HasAuditLogs

`func (o *KmsListAuditLogs200Response) HasAuditLogs() bool`

HasAuditLogs returns a boolean if a field has been set.

### GetTotalCount

`func (o *KmsListAuditLogs200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *KmsListAuditLogs200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *KmsListAuditLogs200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *KmsListAuditLogs200Response) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


