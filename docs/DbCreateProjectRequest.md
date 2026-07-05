# DbCreateProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | [**DbProjectCreate**](DbProjectCreate.md) |  | 

## Methods

### NewDbCreateProjectRequest

`func NewDbCreateProjectRequest(project DbProjectCreate, ) *DbCreateProjectRequest`

NewDbCreateProjectRequest instantiates a new DbCreateProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbCreateProjectRequestWithDefaults

`func NewDbCreateProjectRequestWithDefaults() *DbCreateProjectRequest`

NewDbCreateProjectRequestWithDefaults instantiates a new DbCreateProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *DbCreateProjectRequest) GetProject() DbProjectCreate`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *DbCreateProjectRequest) GetProjectOk() (*DbProjectCreate, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *DbCreateProjectRequest) SetProject(v DbProjectCreate)`

SetProject sets Project field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


