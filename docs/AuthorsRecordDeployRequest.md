# AuthorsRecordDeployRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepoUrl** | Pointer to **string** | The sourceRepo the project was built from. Empty → honest no-op. | [optional] 
**Project** | **string** | The published project id. | 

## Methods

### NewAuthorsRecordDeployRequest

`func NewAuthorsRecordDeployRequest(project string, ) *AuthorsRecordDeployRequest`

NewAuthorsRecordDeployRequest instantiates a new AuthorsRecordDeployRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorsRecordDeployRequestWithDefaults

`func NewAuthorsRecordDeployRequestWithDefaults() *AuthorsRecordDeployRequest`

NewAuthorsRecordDeployRequestWithDefaults instantiates a new AuthorsRecordDeployRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepoUrl

`func (o *AuthorsRecordDeployRequest) GetRepoUrl() string`

GetRepoUrl returns the RepoUrl field if non-nil, zero value otherwise.

### GetRepoUrlOk

`func (o *AuthorsRecordDeployRequest) GetRepoUrlOk() (*string, bool)`

GetRepoUrlOk returns a tuple with the RepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoUrl

`func (o *AuthorsRecordDeployRequest) SetRepoUrl(v string)`

SetRepoUrl sets RepoUrl field to given value.

### HasRepoUrl

`func (o *AuthorsRecordDeployRequest) HasRepoUrl() bool`

HasRepoUrl returns a boolean if a field has been set.

### GetProject

`func (o *AuthorsRecordDeployRequest) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AuthorsRecordDeployRequest) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AuthorsRecordDeployRequest) SetProject(v string)`

SetProject sets Project field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


