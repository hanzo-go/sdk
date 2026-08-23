# ProjectsStar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Starred** | Pointer to **bool** | Starred is whether THIS caller has starred the project after the toggle — their own bookmark, not a property the project carries, so two people see two answers for one project. | [optional] 

## Methods

### NewProjectsStar

`func NewProjectsStar() *ProjectsStar`

NewProjectsStar instantiates a new ProjectsStar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsStarWithDefaults

`func NewProjectsStarWithDefaults() *ProjectsStar`

NewProjectsStarWithDefaults instantiates a new ProjectsStar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStarred

`func (o *ProjectsStar) GetStarred() bool`

GetStarred returns the Starred field if non-nil, zero value otherwise.

### GetStarredOk

`func (o *ProjectsStar) GetStarredOk() (*bool, bool)`

GetStarredOk returns a tuple with the Starred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarred

`func (o *ProjectsStar) SetStarred(v bool)`

SetStarred sets Starred field to given value.

### HasStarred

`func (o *ProjectsStar) HasStarred() bool`

HasStarred returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


