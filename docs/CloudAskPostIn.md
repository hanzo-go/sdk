# CloudAskPostIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **string** | Query is the question, from the BODY. Takes precedence over &#x60;?q&#x3D;&#x60;. | [optional] 
**Repo** | Pointer to **string** | Repo is the repository narrowing, from the BODY. Takes precedence over &#x60;?repo&#x3D;&#x60;. | [optional] 

## Methods

### NewCloudAskPostIn

`func NewCloudAskPostIn() *CloudAskPostIn`

NewCloudAskPostIn instantiates a new CloudAskPostIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAskPostInWithDefaults

`func NewCloudAskPostInWithDefaults() *CloudAskPostIn`

NewCloudAskPostInWithDefaults instantiates a new CloudAskPostIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *CloudAskPostIn) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CloudAskPostIn) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CloudAskPostIn) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *CloudAskPostIn) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetRepo

`func (o *CloudAskPostIn) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CloudAskPostIn) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CloudAskPostIn) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CloudAskPostIn) HasRepo() bool`

HasRepo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


