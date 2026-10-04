# ProviderGithubSearchReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int64** | Limit caps the answer; 0 takes the default and anything above the ceiling is clamped rather than refused. | [optional] 
**Q** | Pointer to **string** | Q is GitHub&#39;s own search syntax, passed through: \&quot;tetris language:go\&quot;, \&quot;org:hanzoai stars:&gt;10\&quot;. Passing it through rather than inventing a vocabulary means one thing to learn, and it is theirs. | [optional] 

## Methods

### NewProviderGithubSearchReq

`func NewProviderGithubSearchReq() *ProviderGithubSearchReq`

NewProviderGithubSearchReq instantiates a new ProviderGithubSearchReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderGithubSearchReqWithDefaults

`func NewProviderGithubSearchReqWithDefaults() *ProviderGithubSearchReq`

NewProviderGithubSearchReqWithDefaults instantiates a new ProviderGithubSearchReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *ProviderGithubSearchReq) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ProviderGithubSearchReq) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ProviderGithubSearchReq) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ProviderGithubSearchReq) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetQ

`func (o *ProviderGithubSearchReq) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *ProviderGithubSearchReq) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *ProviderGithubSearchReq) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *ProviderGithubSearchReq) HasQ() bool`

HasQ returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


