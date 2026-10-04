# AgentCodingPull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mergeable** | Pointer to **bool** | Mergeable is the forge&#39;s verdict that the base takes the branch without conflict. | [optional] 
**MoreReviews** | Pointer to **bool** | MoreReviews says it has more reviews than the first 50 that Reviews lists. | [optional] 
**Number** | Pointer to **int64** | Number is the pull request&#39;s number in its repository. | [optional] 
**Reviews** | Pointer to [**[]AgentCodingReview**](AgentCodingReview.md) | Reviews are its reviews, oldest first: the first 50. | [optional] 
**State** | Pointer to **string** | State is open, closed or merged. | [optional] 
**Title** | Pointer to **string** | Title is its title. | [optional] 
**Url** | Pointer to **string** | URL is where a person reads and merges it. | [optional] 

## Methods

### NewAgentCodingPull

`func NewAgentCodingPull() *AgentCodingPull`

NewAgentCodingPull instantiates a new AgentCodingPull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentCodingPullWithDefaults

`func NewAgentCodingPullWithDefaults() *AgentCodingPull`

NewAgentCodingPullWithDefaults instantiates a new AgentCodingPull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMergeable

`func (o *AgentCodingPull) GetMergeable() bool`

GetMergeable returns the Mergeable field if non-nil, zero value otherwise.

### GetMergeableOk

`func (o *AgentCodingPull) GetMergeableOk() (*bool, bool)`

GetMergeableOk returns a tuple with the Mergeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeable

`func (o *AgentCodingPull) SetMergeable(v bool)`

SetMergeable sets Mergeable field to given value.

### HasMergeable

`func (o *AgentCodingPull) HasMergeable() bool`

HasMergeable returns a boolean if a field has been set.

### GetMoreReviews

`func (o *AgentCodingPull) GetMoreReviews() bool`

GetMoreReviews returns the MoreReviews field if non-nil, zero value otherwise.

### GetMoreReviewsOk

`func (o *AgentCodingPull) GetMoreReviewsOk() (*bool, bool)`

GetMoreReviewsOk returns a tuple with the MoreReviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoreReviews

`func (o *AgentCodingPull) SetMoreReviews(v bool)`

SetMoreReviews sets MoreReviews field to given value.

### HasMoreReviews

`func (o *AgentCodingPull) HasMoreReviews() bool`

HasMoreReviews returns a boolean if a field has been set.

### GetNumber

`func (o *AgentCodingPull) GetNumber() int64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *AgentCodingPull) GetNumberOk() (*int64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *AgentCodingPull) SetNumber(v int64)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *AgentCodingPull) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetReviews

`func (o *AgentCodingPull) GetReviews() []AgentCodingReview`

GetReviews returns the Reviews field if non-nil, zero value otherwise.

### GetReviewsOk

`func (o *AgentCodingPull) GetReviewsOk() (*[]AgentCodingReview, bool)`

GetReviewsOk returns a tuple with the Reviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviews

`func (o *AgentCodingPull) SetReviews(v []AgentCodingReview)`

SetReviews sets Reviews field to given value.

### HasReviews

`func (o *AgentCodingPull) HasReviews() bool`

HasReviews returns a boolean if a field has been set.

### GetState

`func (o *AgentCodingPull) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AgentCodingPull) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AgentCodingPull) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AgentCodingPull) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTitle

`func (o *AgentCodingPull) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AgentCodingPull) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AgentCodingPull) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AgentCodingPull) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *AgentCodingPull) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AgentCodingPull) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AgentCodingPull) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AgentCodingPull) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


