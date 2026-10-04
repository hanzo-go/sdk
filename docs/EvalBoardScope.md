# EvalBoardScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllOrgs** | Pointer to **bool** | true when a platform admin is seeing every org at once | [optional] 
**Org** | Pointer to **string** | the org the board covers; \&quot;\&quot; when it covers all of them | [optional] 
**Project** | Pointer to **string** | the sub-scope within the org; \&quot;\&quot; is the whole org | [optional] 

## Methods

### NewEvalBoardScope

`func NewEvalBoardScope() *EvalBoardScope`

NewEvalBoardScope instantiates a new EvalBoardScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvalBoardScopeWithDefaults

`func NewEvalBoardScopeWithDefaults() *EvalBoardScope`

NewEvalBoardScopeWithDefaults instantiates a new EvalBoardScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllOrgs

`func (o *EvalBoardScope) GetAllOrgs() bool`

GetAllOrgs returns the AllOrgs field if non-nil, zero value otherwise.

### GetAllOrgsOk

`func (o *EvalBoardScope) GetAllOrgsOk() (*bool, bool)`

GetAllOrgsOk returns a tuple with the AllOrgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllOrgs

`func (o *EvalBoardScope) SetAllOrgs(v bool)`

SetAllOrgs sets AllOrgs field to given value.

### HasAllOrgs

`func (o *EvalBoardScope) HasAllOrgs() bool`

HasAllOrgs returns a boolean if a field has been set.

### GetOrg

`func (o *EvalBoardScope) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *EvalBoardScope) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *EvalBoardScope) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *EvalBoardScope) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetProject

`func (o *EvalBoardScope) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *EvalBoardScope) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *EvalBoardScope) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *EvalBoardScope) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


