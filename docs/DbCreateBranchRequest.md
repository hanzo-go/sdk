# DbCreateBranchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to [**DbBranchCreate**](DbBranchCreate.md) |  | [optional] 
**Endpoints** | Pointer to [**[]DbEndpointCreate**](DbEndpointCreate.md) |  | [optional] 

## Methods

### NewDbCreateBranchRequest

`func NewDbCreateBranchRequest() *DbCreateBranchRequest`

NewDbCreateBranchRequest instantiates a new DbCreateBranchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbCreateBranchRequestWithDefaults

`func NewDbCreateBranchRequestWithDefaults() *DbCreateBranchRequest`

NewDbCreateBranchRequestWithDefaults instantiates a new DbCreateBranchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *DbCreateBranchRequest) GetBranch() DbBranchCreate`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *DbCreateBranchRequest) GetBranchOk() (*DbBranchCreate, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *DbCreateBranchRequest) SetBranch(v DbBranchCreate)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *DbCreateBranchRequest) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetEndpoints

`func (o *DbCreateBranchRequest) GetEndpoints() []DbEndpointCreate`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *DbCreateBranchRequest) GetEndpointsOk() (*[]DbEndpointCreate, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *DbCreateBranchRequest) SetEndpoints(v []DbEndpointCreate)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *DbCreateBranchRequest) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


