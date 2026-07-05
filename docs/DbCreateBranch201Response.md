# DbCreateBranch201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to [**DbBranch**](DbBranch.md) |  | [optional] 
**Endpoints** | Pointer to [**[]DbEndpoint**](DbEndpoint.md) |  | [optional] 
**Operations** | Pointer to [**[]DbOperation**](DbOperation.md) |  | [optional] 

## Methods

### NewDbCreateBranch201Response

`func NewDbCreateBranch201Response() *DbCreateBranch201Response`

NewDbCreateBranch201Response instantiates a new DbCreateBranch201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbCreateBranch201ResponseWithDefaults

`func NewDbCreateBranch201ResponseWithDefaults() *DbCreateBranch201Response`

NewDbCreateBranch201ResponseWithDefaults instantiates a new DbCreateBranch201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *DbCreateBranch201Response) GetBranch() DbBranch`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *DbCreateBranch201Response) GetBranchOk() (*DbBranch, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *DbCreateBranch201Response) SetBranch(v DbBranch)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *DbCreateBranch201Response) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetEndpoints

`func (o *DbCreateBranch201Response) GetEndpoints() []DbEndpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *DbCreateBranch201Response) GetEndpointsOk() (*[]DbEndpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *DbCreateBranch201Response) SetEndpoints(v []DbEndpoint)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *DbCreateBranch201Response) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetOperations

`func (o *DbCreateBranch201Response) GetOperations() []DbOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *DbCreateBranch201Response) GetOperationsOk() (*[]DbOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *DbCreateBranch201Response) SetOperations(v []DbOperation)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *DbCreateBranch201Response) HasOperations() bool`

HasOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


