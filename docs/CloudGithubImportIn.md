# CloudGithubImportIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **bool** | All imports every repository the installation grants, instead of naming them. Archived and disabled repositories are skipped either way — they cannot be fetched. | [optional] 
**Repos** | Pointer to **[]string** | Repos names the repositories to import, either owner-qualified (\&quot;hanzo-apps/ai\&quot;) or as a bare name (\&quot;ai\&quot;); a trailing \&quot;.git\&quot; is stripped. A bare name that matches more than one granted repository is an error rather than a guess, because one Hanzo org may hold several GitHub installations and a name is only unique within an owner. Ignored when all is true. | [optional] 

## Methods

### NewCloudGithubImportIn

`func NewCloudGithubImportIn() *CloudGithubImportIn`

NewCloudGithubImportIn instantiates a new CloudGithubImportIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudGithubImportInWithDefaults

`func NewCloudGithubImportInWithDefaults() *CloudGithubImportIn`

NewCloudGithubImportInWithDefaults instantiates a new CloudGithubImportIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *CloudGithubImportIn) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *CloudGithubImportIn) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *CloudGithubImportIn) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *CloudGithubImportIn) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetRepos

`func (o *CloudGithubImportIn) GetRepos() []string`

GetRepos returns the Repos field if non-nil, zero value otherwise.

### GetReposOk

`func (o *CloudGithubImportIn) GetReposOk() (*[]string, bool)`

GetReposOk returns a tuple with the Repos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepos

`func (o *CloudGithubImportIn) SetRepos(v []string)`

SetRepos sets Repos field to given value.

### HasRepos

`func (o *CloudGithubImportIn) HasRepos() bool`

HasRepos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


