# AgentCodingReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** | At is when it was submitted, RFC 3339 in UTC; empty for one that has not been. | [optional] 
**Author** | Pointer to **string** | Author is the reviewer&#39;s forge login. | [optional] 
**Body** | Pointer to **string** | Body is what the reviewer wrote; empty for a bare verdict. At most 16 KiB of it, and 256 KiB across the reviews listed. | [optional] 
**State** | Pointer to **string** | State is the verdict: approved, request_changes, comment, pending or request_review. | [optional] 
**Truncated** | Pointer to **bool** | Truncated marks a Body cut at either bound — or left empty past the second, where the review keeps its author, verdict and time. | [optional] 

## Methods

### NewAgentCodingReview

`func NewAgentCodingReview() *AgentCodingReview`

NewAgentCodingReview instantiates a new AgentCodingReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentCodingReviewWithDefaults

`func NewAgentCodingReviewWithDefaults() *AgentCodingReview`

NewAgentCodingReviewWithDefaults instantiates a new AgentCodingReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *AgentCodingReview) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *AgentCodingReview) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *AgentCodingReview) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *AgentCodingReview) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetAuthor

`func (o *AgentCodingReview) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *AgentCodingReview) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *AgentCodingReview) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *AgentCodingReview) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetBody

`func (o *AgentCodingReview) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *AgentCodingReview) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *AgentCodingReview) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *AgentCodingReview) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetState

`func (o *AgentCodingReview) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AgentCodingReview) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AgentCodingReview) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AgentCodingReview) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTruncated

`func (o *AgentCodingReview) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *AgentCodingReview) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *AgentCodingReview) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *AgentCodingReview) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


