# AskPostIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **string** | Query is the question, from the BODY. Takes precedence over &#x60;?q&#x3D;&#x60;. | [optional] 
**Repo** | Pointer to **string** | Repo is the repository narrowing, from the BODY. Takes precedence over &#x60;?repo&#x3D;&#x60;. | [optional] 

## Methods

### NewAskPostIn

`func NewAskPostIn() *AskPostIn`

NewAskPostIn instantiates a new AskPostIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAskPostInWithDefaults

`func NewAskPostInWithDefaults() *AskPostIn`

NewAskPostInWithDefaults instantiates a new AskPostIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *AskPostIn) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *AskPostIn) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *AskPostIn) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *AskPostIn) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetRepo

`func (o *AskPostIn) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *AskPostIn) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *AskPostIn) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *AskPostIn) HasRepo() bool`

HasRepo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


