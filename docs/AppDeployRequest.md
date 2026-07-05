# AppDeployRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | Deploy source. &#x60;git&#x60; queues a CI build from a commit. | 
**Commit** | Pointer to **string** | Git commit SHA to build. | [optional] 

## Methods

### NewAppDeployRequest

`func NewAppDeployRequest(source string, ) *AppDeployRequest`

NewAppDeployRequest instantiates a new AppDeployRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDeployRequestWithDefaults

`func NewAppDeployRequestWithDefaults() *AppDeployRequest`

NewAppDeployRequestWithDefaults instantiates a new AppDeployRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *AppDeployRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AppDeployRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AppDeployRequest) SetSource(v string)`

SetSource sets Source field to given value.


### GetCommit

`func (o *AppDeployRequest) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *AppDeployRequest) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *AppDeployRequest) SetCommit(v string)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *AppDeployRequest) HasCommit() bool`

HasCommit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


