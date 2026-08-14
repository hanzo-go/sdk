# BoardScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllOrgs** | Pointer to **bool** | true when a platform admin is seeing every org at once | [optional] 
**Org** | Pointer to **string** | the org the board covers; \&quot;\&quot; when it covers all of them | [optional] 
**Project** | Pointer to **string** | the sub-scope within the org; \&quot;\&quot; is the whole org | [optional] 

## Methods

### NewBoardScope

`func NewBoardScope() *BoardScope`

NewBoardScope instantiates a new BoardScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardScopeWithDefaults

`func NewBoardScopeWithDefaults() *BoardScope`

NewBoardScopeWithDefaults instantiates a new BoardScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllOrgs

`func (o *BoardScope) GetAllOrgs() bool`

GetAllOrgs returns the AllOrgs field if non-nil, zero value otherwise.

### GetAllOrgsOk

`func (o *BoardScope) GetAllOrgsOk() (*bool, bool)`

GetAllOrgsOk returns a tuple with the AllOrgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllOrgs

`func (o *BoardScope) SetAllOrgs(v bool)`

SetAllOrgs sets AllOrgs field to given value.

### HasAllOrgs

`func (o *BoardScope) HasAllOrgs() bool`

HasAllOrgs returns a boolean if a field has been set.

### GetOrg

`func (o *BoardScope) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *BoardScope) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *BoardScope) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *BoardScope) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetProject

`func (o *BoardScope) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *BoardScope) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *BoardScope) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *BoardScope) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


