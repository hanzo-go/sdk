# ResearchGrantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | an experiment (run) stable id | [optional] 
**Sha256** | Pointer to **string** | OR an artifact content hash | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Trainable** | Pointer to **bool** |  | [optional] 
**Publishable** | Pointer to **bool** |  | [optional] 

## Methods

### NewResearchGrantRequest

`func NewResearchGrantRequest() *ResearchGrantRequest`

NewResearchGrantRequest instantiates a new ResearchGrantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResearchGrantRequestWithDefaults

`func NewResearchGrantRequestWithDefaults() *ResearchGrantRequest`

NewResearchGrantRequestWithDefaults instantiates a new ResearchGrantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *ResearchGrantRequest) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ResearchGrantRequest) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ResearchGrantRequest) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ResearchGrantRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetId

`func (o *ResearchGrantRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResearchGrantRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResearchGrantRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResearchGrantRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSha256

`func (o *ResearchGrantRequest) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *ResearchGrantRequest) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *ResearchGrantRequest) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *ResearchGrantRequest) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetVisibility

`func (o *ResearchGrantRequest) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ResearchGrantRequest) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ResearchGrantRequest) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ResearchGrantRequest) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTrainable

`func (o *ResearchGrantRequest) GetTrainable() bool`

GetTrainable returns the Trainable field if non-nil, zero value otherwise.

### GetTrainableOk

`func (o *ResearchGrantRequest) GetTrainableOk() (*bool, bool)`

GetTrainableOk returns a tuple with the Trainable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainable

`func (o *ResearchGrantRequest) SetTrainable(v bool)`

SetTrainable sets Trainable field to given value.

### HasTrainable

`func (o *ResearchGrantRequest) HasTrainable() bool`

HasTrainable returns a boolean if a field has been set.

### GetPublishable

`func (o *ResearchGrantRequest) GetPublishable() bool`

GetPublishable returns the Publishable field if non-nil, zero value otherwise.

### GetPublishableOk

`func (o *ResearchGrantRequest) GetPublishableOk() (*bool, bool)`

GetPublishableOk returns a tuple with the Publishable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishable

`func (o *ResearchGrantRequest) SetPublishable(v bool)`

SetPublishable sets Publishable field to given value.

### HasPublishable

`func (o *ResearchGrantRequest) HasPublishable() bool`

HasPublishable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


