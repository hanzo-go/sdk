# CloudGrantRequest

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

### NewCloudGrantRequest

`func NewCloudGrantRequest() *CloudGrantRequest`

NewCloudGrantRequest instantiates a new CloudGrantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudGrantRequestWithDefaults

`func NewCloudGrantRequestWithDefaults() *CloudGrantRequest`

NewCloudGrantRequestWithDefaults instantiates a new CloudGrantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudGrantRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudGrantRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudGrantRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudGrantRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProject

`func (o *CloudGrantRequest) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CloudGrantRequest) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CloudGrantRequest) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *CloudGrantRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetPublishable

`func (o *CloudGrantRequest) GetPublishable() bool`

GetPublishable returns the Publishable field if non-nil, zero value otherwise.

### GetPublishableOk

`func (o *CloudGrantRequest) GetPublishableOk() (*bool, bool)`

GetPublishableOk returns a tuple with the Publishable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishable

`func (o *CloudGrantRequest) SetPublishable(v bool)`

SetPublishable sets Publishable field to given value.

### HasPublishable

`func (o *CloudGrantRequest) HasPublishable() bool`

HasPublishable returns a boolean if a field has been set.

### GetSha256

`func (o *CloudGrantRequest) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *CloudGrantRequest) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *CloudGrantRequest) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *CloudGrantRequest) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetTrainable

`func (o *CloudGrantRequest) GetTrainable() bool`

GetTrainable returns the Trainable field if non-nil, zero value otherwise.

### GetTrainableOk

`func (o *CloudGrantRequest) GetTrainableOk() (*bool, bool)`

GetTrainableOk returns a tuple with the Trainable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainable

`func (o *CloudGrantRequest) SetTrainable(v bool)`

SetTrainable sets Trainable field to given value.

### HasTrainable

`func (o *CloudGrantRequest) HasTrainable() bool`

HasTrainable returns a boolean if a field has been set.

### GetVisibility

`func (o *CloudGrantRequest) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *CloudGrantRequest) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *CloudGrantRequest) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *CloudGrantRequest) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


