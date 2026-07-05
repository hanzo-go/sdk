# DbBranch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** | Parent branch ID (null for root) | [optional] 
**ParentLsn** | Pointer to **string** | LSN at which branch was forked | [optional] 
**ParentTimestamp** | Pointer to **time.Time** | Timestamp at which branch was forked | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**CurrentState** | Pointer to **string** |  | [optional] 
**LogicalSize** | Pointer to **int64** |  | [optional] 
**Primary** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDbBranch

`func NewDbBranch() *DbBranch`

NewDbBranch instantiates a new DbBranch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbBranchWithDefaults

`func NewDbBranchWithDefaults() *DbBranch`

NewDbBranchWithDefaults instantiates a new DbBranch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DbBranch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbBranch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbBranch) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DbBranch) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProjectId

`func (o *DbBranch) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DbBranch) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DbBranch) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *DbBranch) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetParentId

`func (o *DbBranch) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *DbBranch) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *DbBranch) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *DbBranch) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetParentLsn

`func (o *DbBranch) GetParentLsn() string`

GetParentLsn returns the ParentLsn field if non-nil, zero value otherwise.

### GetParentLsnOk

`func (o *DbBranch) GetParentLsnOk() (*string, bool)`

GetParentLsnOk returns a tuple with the ParentLsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentLsn

`func (o *DbBranch) SetParentLsn(v string)`

SetParentLsn sets ParentLsn field to given value.

### HasParentLsn

`func (o *DbBranch) HasParentLsn() bool`

HasParentLsn returns a boolean if a field has been set.

### GetParentTimestamp

`func (o *DbBranch) GetParentTimestamp() time.Time`

GetParentTimestamp returns the ParentTimestamp field if non-nil, zero value otherwise.

### GetParentTimestampOk

`func (o *DbBranch) GetParentTimestampOk() (*time.Time, bool)`

GetParentTimestampOk returns a tuple with the ParentTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTimestamp

`func (o *DbBranch) SetParentTimestamp(v time.Time)`

SetParentTimestamp sets ParentTimestamp field to given value.

### HasParentTimestamp

`func (o *DbBranch) HasParentTimestamp() bool`

HasParentTimestamp returns a boolean if a field has been set.

### GetName

`func (o *DbBranch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DbBranch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DbBranch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DbBranch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCurrentState

`func (o *DbBranch) GetCurrentState() string`

GetCurrentState returns the CurrentState field if non-nil, zero value otherwise.

### GetCurrentStateOk

`func (o *DbBranch) GetCurrentStateOk() (*string, bool)`

GetCurrentStateOk returns a tuple with the CurrentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentState

`func (o *DbBranch) SetCurrentState(v string)`

SetCurrentState sets CurrentState field to given value.

### HasCurrentState

`func (o *DbBranch) HasCurrentState() bool`

HasCurrentState returns a boolean if a field has been set.

### GetLogicalSize

`func (o *DbBranch) GetLogicalSize() int64`

GetLogicalSize returns the LogicalSize field if non-nil, zero value otherwise.

### GetLogicalSizeOk

`func (o *DbBranch) GetLogicalSizeOk() (*int64, bool)`

GetLogicalSizeOk returns a tuple with the LogicalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalSize

`func (o *DbBranch) SetLogicalSize(v int64)`

SetLogicalSize sets LogicalSize field to given value.

### HasLogicalSize

`func (o *DbBranch) HasLogicalSize() bool`

HasLogicalSize returns a boolean if a field has been set.

### GetPrimary

`func (o *DbBranch) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *DbBranch) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *DbBranch) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *DbBranch) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DbBranch) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DbBranch) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DbBranch) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DbBranch) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DbBranch) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DbBranch) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DbBranch) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DbBranch) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


