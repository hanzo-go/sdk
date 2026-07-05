# KmsApprovalPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Approvals** | Pointer to **int32** |  | [optional] 
**SecretPath** | Pointer to **string** |  | [optional] 
**EnvId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewKmsApprovalPolicy

`func NewKmsApprovalPolicy() *KmsApprovalPolicy`

NewKmsApprovalPolicy instantiates a new KmsApprovalPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsApprovalPolicyWithDefaults

`func NewKmsApprovalPolicyWithDefaults() *KmsApprovalPolicy`

NewKmsApprovalPolicyWithDefaults instantiates a new KmsApprovalPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KmsApprovalPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KmsApprovalPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KmsApprovalPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KmsApprovalPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KmsApprovalPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KmsApprovalPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KmsApprovalPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KmsApprovalPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *KmsApprovalPolicy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KmsApprovalPolicy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KmsApprovalPolicy) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KmsApprovalPolicy) HasType() bool`

HasType returns a boolean if a field has been set.

### GetApprovals

`func (o *KmsApprovalPolicy) GetApprovals() int32`

GetApprovals returns the Approvals field if non-nil, zero value otherwise.

### GetApprovalsOk

`func (o *KmsApprovalPolicy) GetApprovalsOk() (*int32, bool)`

GetApprovalsOk returns a tuple with the Approvals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovals

`func (o *KmsApprovalPolicy) SetApprovals(v int32)`

SetApprovals sets Approvals field to given value.

### HasApprovals

`func (o *KmsApprovalPolicy) HasApprovals() bool`

HasApprovals returns a boolean if a field has been set.

### GetSecretPath

`func (o *KmsApprovalPolicy) GetSecretPath() string`

GetSecretPath returns the SecretPath field if non-nil, zero value otherwise.

### GetSecretPathOk

`func (o *KmsApprovalPolicy) GetSecretPathOk() (*string, bool)`

GetSecretPathOk returns a tuple with the SecretPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretPath

`func (o *KmsApprovalPolicy) SetSecretPath(v string)`

SetSecretPath sets SecretPath field to given value.

### HasSecretPath

`func (o *KmsApprovalPolicy) HasSecretPath() bool`

HasSecretPath returns a boolean if a field has been set.

### GetEnvId

`func (o *KmsApprovalPolicy) GetEnvId() string`

GetEnvId returns the EnvId field if non-nil, zero value otherwise.

### GetEnvIdOk

`func (o *KmsApprovalPolicy) GetEnvIdOk() (*string, bool)`

GetEnvIdOk returns a tuple with the EnvId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvId

`func (o *KmsApprovalPolicy) SetEnvId(v string)`

SetEnvId sets EnvId field to given value.

### HasEnvId

`func (o *KmsApprovalPolicy) HasEnvId() bool`

HasEnvId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KmsApprovalPolicy) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KmsApprovalPolicy) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KmsApprovalPolicy) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KmsApprovalPolicy) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KmsApprovalPolicy) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KmsApprovalPolicy) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KmsApprovalPolicy) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KmsApprovalPolicy) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


