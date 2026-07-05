# DbRestoreBranch200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to [**DbBranch**](DbBranch.md) |  | [optional] 
**Operations** | Pointer to [**[]DbOperation**](DbOperation.md) |  | [optional] 

## Methods

### NewDbRestoreBranch200Response

`func NewDbRestoreBranch200Response() *DbRestoreBranch200Response`

NewDbRestoreBranch200Response instantiates a new DbRestoreBranch200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbRestoreBranch200ResponseWithDefaults

`func NewDbRestoreBranch200ResponseWithDefaults() *DbRestoreBranch200Response`

NewDbRestoreBranch200ResponseWithDefaults instantiates a new DbRestoreBranch200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *DbRestoreBranch200Response) GetBranch() DbBranch`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *DbRestoreBranch200Response) GetBranchOk() (*DbBranch, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *DbRestoreBranch200Response) SetBranch(v DbBranch)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *DbRestoreBranch200Response) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetOperations

`func (o *DbRestoreBranch200Response) GetOperations() []DbOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *DbRestoreBranch200Response) GetOperationsOk() (*[]DbOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *DbRestoreBranch200Response) SetOperations(v []DbOperation)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *DbRestoreBranch200Response) HasOperations() bool`

HasOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


