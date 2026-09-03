# IamListOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditLogs** | Pointer to [**[]IamAuditLog**](IamAuditLog.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewIamListOutput

`func NewIamListOutput() *IamListOutput`

NewIamListOutput instantiates a new IamListOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamListOutputWithDefaults

`func NewIamListOutputWithDefaults() *IamListOutput`

NewIamListOutputWithDefaults instantiates a new IamListOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditLogs

`func (o *IamListOutput) GetAuditLogs() []IamAuditLog`

GetAuditLogs returns the AuditLogs field if non-nil, zero value otherwise.

### GetAuditLogsOk

`func (o *IamListOutput) GetAuditLogsOk() (*[]IamAuditLog, bool)`

GetAuditLogsOk returns a tuple with the AuditLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogs

`func (o *IamListOutput) SetAuditLogs(v []IamAuditLog)`

SetAuditLogs sets AuditLogs field to given value.

### HasAuditLogs

`func (o *IamListOutput) HasAuditLogs() bool`

HasAuditLogs returns a boolean if a field has been set.

### GetTotal

`func (o *IamListOutput) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *IamListOutput) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *IamListOutput) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *IamListOutput) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


