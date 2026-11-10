# PatentApprovalBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approvals** | Pointer to [**[]PatentPatentApproval**](PatentPatentApproval.md) | Approvals are the decisions to append, at most 2000. | [optional] 

## Methods

### NewPatentApprovalBatch

`func NewPatentApprovalBatch() *PatentApprovalBatch`

NewPatentApprovalBatch instantiates a new PatentApprovalBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatentApprovalBatchWithDefaults

`func NewPatentApprovalBatchWithDefaults() *PatentApprovalBatch`

NewPatentApprovalBatchWithDefaults instantiates a new PatentApprovalBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovals

`func (o *PatentApprovalBatch) GetApprovals() []PatentPatentApproval`

GetApprovals returns the Approvals field if non-nil, zero value otherwise.

### GetApprovalsOk

`func (o *PatentApprovalBatch) GetApprovalsOk() (*[]PatentPatentApproval, bool)`

GetApprovalsOk returns a tuple with the Approvals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovals

`func (o *PatentApprovalBatch) SetApprovals(v []PatentPatentApproval)`

SetApprovals sets Approvals field to given value.

### HasApprovals

`func (o *PatentApprovalBatch) HasApprovals() bool`

HasApprovals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


