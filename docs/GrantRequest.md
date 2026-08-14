# GrantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | an experiment (run) stable id | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Publishable** | Pointer to **bool** |  | [optional] 
**Sha256** | Pointer to **string** | OR an artifact content hash | [optional] 
**Trainable** | Pointer to **bool** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 

## Methods

### NewGrantRequest

`func NewGrantRequest() *GrantRequest`

NewGrantRequest instantiates a new GrantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantRequestWithDefaults

`func NewGrantRequestWithDefaults() *GrantRequest`

NewGrantRequestWithDefaults instantiates a new GrantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GrantRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GrantRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GrantRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GrantRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProject

`func (o *GrantRequest) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *GrantRequest) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *GrantRequest) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *GrantRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetPublishable

`func (o *GrantRequest) GetPublishable() bool`

GetPublishable returns the Publishable field if non-nil, zero value otherwise.

### GetPublishableOk

`func (o *GrantRequest) GetPublishableOk() (*bool, bool)`

GetPublishableOk returns a tuple with the Publishable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishable

`func (o *GrantRequest) SetPublishable(v bool)`

SetPublishable sets Publishable field to given value.

### HasPublishable

`func (o *GrantRequest) HasPublishable() bool`

HasPublishable returns a boolean if a field has been set.

### GetSha256

`func (o *GrantRequest) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *GrantRequest) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *GrantRequest) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *GrantRequest) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetTrainable

`func (o *GrantRequest) GetTrainable() bool`

GetTrainable returns the Trainable field if non-nil, zero value otherwise.

### GetTrainableOk

`func (o *GrantRequest) GetTrainableOk() (*bool, bool)`

GetTrainableOk returns a tuple with the Trainable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainable

`func (o *GrantRequest) SetTrainable(v bool)`

SetTrainable sets Trainable field to given value.

### HasTrainable

`func (o *GrantRequest) HasTrainable() bool`

HasTrainable returns a boolean if a field has been set.

### GetVisibility

`func (o *GrantRequest) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GrantRequest) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GrantRequest) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GrantRequest) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


