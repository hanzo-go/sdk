# DbOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**BranchId** | Pointer to **string** |  | [optional] 
**EndpointId** | Pointer to **string** |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**FailuresCount** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDbOperation

`func NewDbOperation() *DbOperation`

NewDbOperation instantiates a new DbOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbOperationWithDefaults

`func NewDbOperationWithDefaults() *DbOperation`

NewDbOperationWithDefaults instantiates a new DbOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DbOperation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DbOperation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DbOperation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DbOperation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProjectId

`func (o *DbOperation) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DbOperation) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DbOperation) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *DbOperation) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetBranchId

`func (o *DbOperation) GetBranchId() string`

GetBranchId returns the BranchId field if non-nil, zero value otherwise.

### GetBranchIdOk

`func (o *DbOperation) GetBranchIdOk() (*string, bool)`

GetBranchIdOk returns a tuple with the BranchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchId

`func (o *DbOperation) SetBranchId(v string)`

SetBranchId sets BranchId field to given value.

### HasBranchId

`func (o *DbOperation) HasBranchId() bool`

HasBranchId returns a boolean if a field has been set.

### GetEndpointId

`func (o *DbOperation) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *DbOperation) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *DbOperation) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *DbOperation) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetAction

`func (o *DbOperation) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DbOperation) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DbOperation) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *DbOperation) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetStatus

`func (o *DbOperation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DbOperation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DbOperation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DbOperation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFailuresCount

`func (o *DbOperation) GetFailuresCount() int32`

GetFailuresCount returns the FailuresCount field if non-nil, zero value otherwise.

### GetFailuresCountOk

`func (o *DbOperation) GetFailuresCountOk() (*int32, bool)`

GetFailuresCountOk returns a tuple with the FailuresCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailuresCount

`func (o *DbOperation) SetFailuresCount(v int32)`

SetFailuresCount sets FailuresCount field to given value.

### HasFailuresCount

`func (o *DbOperation) HasFailuresCount() bool`

HasFailuresCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DbOperation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DbOperation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DbOperation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DbOperation) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DbOperation) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DbOperation) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DbOperation) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DbOperation) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


