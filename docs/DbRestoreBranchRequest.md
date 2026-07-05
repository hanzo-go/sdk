# DbRestoreBranchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceBranchId** | Pointer to **string** |  | [optional] 
**SourceLsn** | Pointer to **string** | Restore to this LSN | [optional] 
**SourceTimestamp** | Pointer to **time.Time** | Restore to this point in time | [optional] 
**PreserveUnderName** | Pointer to **string** | Preserve current state under this name before restoring | [optional] 

## Methods

### NewDbRestoreBranchRequest

`func NewDbRestoreBranchRequest() *DbRestoreBranchRequest`

NewDbRestoreBranchRequest instantiates a new DbRestoreBranchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbRestoreBranchRequestWithDefaults

`func NewDbRestoreBranchRequestWithDefaults() *DbRestoreBranchRequest`

NewDbRestoreBranchRequestWithDefaults instantiates a new DbRestoreBranchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceBranchId

`func (o *DbRestoreBranchRequest) GetSourceBranchId() string`

GetSourceBranchId returns the SourceBranchId field if non-nil, zero value otherwise.

### GetSourceBranchIdOk

`func (o *DbRestoreBranchRequest) GetSourceBranchIdOk() (*string, bool)`

GetSourceBranchIdOk returns a tuple with the SourceBranchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBranchId

`func (o *DbRestoreBranchRequest) SetSourceBranchId(v string)`

SetSourceBranchId sets SourceBranchId field to given value.

### HasSourceBranchId

`func (o *DbRestoreBranchRequest) HasSourceBranchId() bool`

HasSourceBranchId returns a boolean if a field has been set.

### GetSourceLsn

`func (o *DbRestoreBranchRequest) GetSourceLsn() string`

GetSourceLsn returns the SourceLsn field if non-nil, zero value otherwise.

### GetSourceLsnOk

`func (o *DbRestoreBranchRequest) GetSourceLsnOk() (*string, bool)`

GetSourceLsnOk returns a tuple with the SourceLsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLsn

`func (o *DbRestoreBranchRequest) SetSourceLsn(v string)`

SetSourceLsn sets SourceLsn field to given value.

### HasSourceLsn

`func (o *DbRestoreBranchRequest) HasSourceLsn() bool`

HasSourceLsn returns a boolean if a field has been set.

### GetSourceTimestamp

`func (o *DbRestoreBranchRequest) GetSourceTimestamp() time.Time`

GetSourceTimestamp returns the SourceTimestamp field if non-nil, zero value otherwise.

### GetSourceTimestampOk

`func (o *DbRestoreBranchRequest) GetSourceTimestampOk() (*time.Time, bool)`

GetSourceTimestampOk returns a tuple with the SourceTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTimestamp

`func (o *DbRestoreBranchRequest) SetSourceTimestamp(v time.Time)`

SetSourceTimestamp sets SourceTimestamp field to given value.

### HasSourceTimestamp

`func (o *DbRestoreBranchRequest) HasSourceTimestamp() bool`

HasSourceTimestamp returns a boolean if a field has been set.

### GetPreserveUnderName

`func (o *DbRestoreBranchRequest) GetPreserveUnderName() string`

GetPreserveUnderName returns the PreserveUnderName field if non-nil, zero value otherwise.

### GetPreserveUnderNameOk

`func (o *DbRestoreBranchRequest) GetPreserveUnderNameOk() (*string, bool)`

GetPreserveUnderNameOk returns a tuple with the PreserveUnderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveUnderName

`func (o *DbRestoreBranchRequest) SetPreserveUnderName(v string)`

SetPreserveUnderName sets PreserveUnderName field to given value.

### HasPreserveUnderName

`func (o *DbRestoreBranchRequest) HasPreserveUnderName() bool`

HasPreserveUnderName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


