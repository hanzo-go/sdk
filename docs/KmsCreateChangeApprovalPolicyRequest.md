# KmsCreateChangeApprovalPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceId** | **string** |  | 
**Name** | **string** |  | 
**Approvals** | **int32** |  | 
**Environment** | **string** |  | 
**SecretPath** | Pointer to **string** |  | [optional] 

## Methods

### NewKmsCreateChangeApprovalPolicyRequest

`func NewKmsCreateChangeApprovalPolicyRequest(workspaceId string, name string, approvals int32, environment string, ) *KmsCreateChangeApprovalPolicyRequest`

NewKmsCreateChangeApprovalPolicyRequest instantiates a new KmsCreateChangeApprovalPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsCreateChangeApprovalPolicyRequestWithDefaults

`func NewKmsCreateChangeApprovalPolicyRequestWithDefaults() *KmsCreateChangeApprovalPolicyRequest`

NewKmsCreateChangeApprovalPolicyRequestWithDefaults instantiates a new KmsCreateChangeApprovalPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceId

`func (o *KmsCreateChangeApprovalPolicyRequest) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *KmsCreateChangeApprovalPolicyRequest) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *KmsCreateChangeApprovalPolicyRequest) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetName

`func (o *KmsCreateChangeApprovalPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KmsCreateChangeApprovalPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KmsCreateChangeApprovalPolicyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetApprovals

`func (o *KmsCreateChangeApprovalPolicyRequest) GetApprovals() int32`

GetApprovals returns the Approvals field if non-nil, zero value otherwise.

### GetApprovalsOk

`func (o *KmsCreateChangeApprovalPolicyRequest) GetApprovalsOk() (*int32, bool)`

GetApprovalsOk returns a tuple with the Approvals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovals

`func (o *KmsCreateChangeApprovalPolicyRequest) SetApprovals(v int32)`

SetApprovals sets Approvals field to given value.


### GetEnvironment

`func (o *KmsCreateChangeApprovalPolicyRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *KmsCreateChangeApprovalPolicyRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *KmsCreateChangeApprovalPolicyRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetSecretPath

`func (o *KmsCreateChangeApprovalPolicyRequest) GetSecretPath() string`

GetSecretPath returns the SecretPath field if non-nil, zero value otherwise.

### GetSecretPathOk

`func (o *KmsCreateChangeApprovalPolicyRequest) GetSecretPathOk() (*string, bool)`

GetSecretPathOk returns a tuple with the SecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretPath

`func (o *KmsCreateChangeApprovalPolicyRequest) SetSecretPath(v string)`

SetSecretPath sets SecretPath field to given value.

### HasSecretPath

`func (o *KmsCreateChangeApprovalPolicyRequest) HasSecretPath() bool`

HasSecretPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


