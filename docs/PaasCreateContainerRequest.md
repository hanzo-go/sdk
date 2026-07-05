# PaasCreateContainerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** |  | 
**Template** | Pointer to **string** | Template slug (nodejs, python, go, rust, etc.) | [optional] 
**Repo** | Pointer to [**PaasCreateContainerRequestRepo**](PaasCreateContainerRequestRepo.md) |  | [optional] 

## Methods

### NewPaasCreateContainerRequest

`func NewPaasCreateContainerRequest(name string, type_ string, ) *PaasCreateContainerRequest`

NewPaasCreateContainerRequest instantiates a new PaasCreateContainerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaasCreateContainerRequestWithDefaults

`func NewPaasCreateContainerRequestWithDefaults() *PaasCreateContainerRequest`

NewPaasCreateContainerRequestWithDefaults instantiates a new PaasCreateContainerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PaasCreateContainerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaasCreateContainerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaasCreateContainerRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *PaasCreateContainerRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaasCreateContainerRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaasCreateContainerRequest) SetType(v string)`

SetType sets Type field to given value.


### GetTemplate

`func (o *PaasCreateContainerRequest) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *PaasCreateContainerRequest) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *PaasCreateContainerRequest) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *PaasCreateContainerRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetRepo

`func (o *PaasCreateContainerRequest) GetRepo() PaasCreateContainerRequestRepo`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *PaasCreateContainerRequest) GetRepoOk() (*PaasCreateContainerRequestRepo, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *PaasCreateContainerRequest) SetRepo(v PaasCreateContainerRequestRepo)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *PaasCreateContainerRequest) HasRepo() bool`

HasRepo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


