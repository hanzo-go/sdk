# DbCreateProject201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | Pointer to [**DbProject**](DbProject.md) |  | [optional] 
**ConnectionUris** | Pointer to [**[]DbConnectionUri**](DbConnectionUri.md) |  | [optional] 
**Branch** | Pointer to [**DbBranch**](DbBranch.md) |  | [optional] 
**Endpoints** | Pointer to [**[]DbEndpoint**](DbEndpoint.md) |  | [optional] 
**Databases** | Pointer to [**[]DbDatabase**](DbDatabase.md) |  | [optional] 
**Roles** | Pointer to [**[]DbRole**](DbRole.md) |  | [optional] 

## Methods

### NewDbCreateProject201Response

`func NewDbCreateProject201Response() *DbCreateProject201Response`

NewDbCreateProject201Response instantiates a new DbCreateProject201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbCreateProject201ResponseWithDefaults

`func NewDbCreateProject201ResponseWithDefaults() *DbCreateProject201Response`

NewDbCreateProject201ResponseWithDefaults instantiates a new DbCreateProject201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *DbCreateProject201Response) GetProject() DbProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *DbCreateProject201Response) GetProjectOk() (*DbProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *DbCreateProject201Response) SetProject(v DbProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *DbCreateProject201Response) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetConnectionUris

`func (o *DbCreateProject201Response) GetConnectionUris() []DbConnectionUri`

GetConnectionUris returns the ConnectionUris field if non-nil, zero value otherwise.

### GetConnectionUrisOk

`func (o *DbCreateProject201Response) GetConnectionUrisOk() (*[]DbConnectionUri, bool)`

GetConnectionUrisOk returns a tuple with the ConnectionUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionUris

`func (o *DbCreateProject201Response) SetConnectionUris(v []DbConnectionUri)`

SetConnectionUris sets ConnectionUris field to given value.

### HasConnectionUris

`func (o *DbCreateProject201Response) HasConnectionUris() bool`

HasConnectionUris returns a boolean if a field has been set.

### GetBranch

`func (o *DbCreateProject201Response) GetBranch() DbBranch`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *DbCreateProject201Response) GetBranchOk() (*DbBranch, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *DbCreateProject201Response) SetBranch(v DbBranch)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *DbCreateProject201Response) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetEndpoints

`func (o *DbCreateProject201Response) GetEndpoints() []DbEndpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *DbCreateProject201Response) GetEndpointsOk() (*[]DbEndpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *DbCreateProject201Response) SetEndpoints(v []DbEndpoint)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *DbCreateProject201Response) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetDatabases

`func (o *DbCreateProject201Response) GetDatabases() []DbDatabase`

GetDatabases returns the Databases field if non-nil, zero value otherwise.

### GetDatabasesOk

`func (o *DbCreateProject201Response) GetDatabasesOk() (*[]DbDatabase, bool)`

GetDatabasesOk returns a tuple with the Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabases

`func (o *DbCreateProject201Response) SetDatabases(v []DbDatabase)`

SetDatabases sets Databases field to given value.

### HasDatabases

`func (o *DbCreateProject201Response) HasDatabases() bool`

HasDatabases returns a boolean if a field has been set.

### GetRoles

`func (o *DbCreateProject201Response) GetRoles() []DbRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *DbCreateProject201Response) GetRolesOk() (*[]DbRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *DbCreateProject201Response) SetRoles(v []DbRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *DbCreateProject201Response) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


