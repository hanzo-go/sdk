# GitLinkAll

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | Pointer to **string** | Owner is the GitHub account a repository named n is looked for under, as github.com/&lt;owner&gt;/&lt;n&gt;. Empty is the account the estate&#39;s table names for the org; an org it does not name must give one. | [optional] 

## Methods

### NewGitLinkAll

`func NewGitLinkAll() *GitLinkAll`

NewGitLinkAll instantiates a new GitLinkAll object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitLinkAllWithDefaults

`func NewGitLinkAllWithDefaults() *GitLinkAll`

NewGitLinkAllWithDefaults instantiates a new GitLinkAll object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *GitLinkAll) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GitLinkAll) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GitLinkAll) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GitLinkAll) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


