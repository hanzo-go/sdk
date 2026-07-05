# DbBranchCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** | Branch to fork from (defaults to primary branch) | [optional] 
**ParentLsn** | Pointer to **string** | Fork at specific LSN | [optional] 
**ParentTimestamp** | Pointer to **time.Time** | Fork at specific point in time | [optional] 

## Methods

### NewDbBranchCreate

`func NewDbBranchCreate() *DbBranchCreate`

NewDbBranchCreate instantiates a new DbBranchCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbBranchCreateWithDefaults

`func NewDbBranchCreateWithDefaults() *DbBranchCreate`

NewDbBranchCreateWithDefaults instantiates a new DbBranchCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DbBranchCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DbBranchCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DbBranchCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DbBranchCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentId

`func (o *DbBranchCreate) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *DbBranchCreate) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *DbBranchCreate) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *DbBranchCreate) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetParentLsn

`func (o *DbBranchCreate) GetParentLsn() string`

GetParentLsn returns the ParentLsn field if non-nil, zero value otherwise.

### GetParentLsnOk

`func (o *DbBranchCreate) GetParentLsnOk() (*string, bool)`

GetParentLsnOk returns a tuple with the ParentLsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentLsn

`func (o *DbBranchCreate) SetParentLsn(v string)`

SetParentLsn sets ParentLsn field to given value.

### HasParentLsn

`func (o *DbBranchCreate) HasParentLsn() bool`

HasParentLsn returns a boolean if a field has been set.

### GetParentTimestamp

`func (o *DbBranchCreate) GetParentTimestamp() time.Time`

GetParentTimestamp returns the ParentTimestamp field if non-nil, zero value otherwise.

### GetParentTimestampOk

`func (o *DbBranchCreate) GetParentTimestampOk() (*time.Time, bool)`

GetParentTimestampOk returns a tuple with the ParentTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTimestamp

`func (o *DbBranchCreate) SetParentTimestamp(v time.Time)`

SetParentTimestamp sets ParentTimestamp field to given value.

### HasParentTimestamp

`func (o *DbBranchCreate) HasParentTimestamp() bool`

HasParentTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


