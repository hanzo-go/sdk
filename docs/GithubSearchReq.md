# GithubSearchReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | Limit caps the answer; 0 takes the default and anything above the ceiling is clamped rather than refused. | [optional] 
**Q** | Pointer to **string** | Q is GitHub&#39;s own search syntax, passed through: \&quot;tetris language:go\&quot;, \&quot;org:hanzoai stars:&gt;10\&quot;. Passing it through rather than inventing a vocabulary means one thing to learn, and it is theirs. | [optional] 

## Methods

### NewGithubSearchReq

`func NewGithubSearchReq() *GithubSearchReq`

NewGithubSearchReq instantiates a new GithubSearchReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubSearchReqWithDefaults

`func NewGithubSearchReqWithDefaults() *GithubSearchReq`

NewGithubSearchReqWithDefaults instantiates a new GithubSearchReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *GithubSearchReq) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GithubSearchReq) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GithubSearchReq) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GithubSearchReq) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetQ

`func (o *GithubSearchReq) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *GithubSearchReq) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *GithubSearchReq) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *GithubSearchReq) HasQ() bool`

HasQ returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


