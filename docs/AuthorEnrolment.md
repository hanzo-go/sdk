# AuthorEnrolment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **bool** | Created reports whether this call enrolled the org (201) or found an existing enrolment (200). | [optional] 
**GithubLogin** | Pointer to **string** | GithubLogin is the linked forge account. | [optional] 
**Id** | Pointer to **string** | ID is the author record&#39;s server-minted handle, \&quot;aut_\&quot;-prefixed. | [optional] 
**ShareBps** | Pointer to **int64** | ShareBps is this author&#39;s royalty share in basis points of the spend their deployed work generates. | [optional] 
**Status** | Pointer to **string** | Status is connected, approved or suspended. Only an approved author earns. | [optional] 
**Verified** | Pointer to **bool** | Verified reports whether any repository or owner claim has been proven yet. | [optional] 
**VerifyCode** | Pointer to **string** | VerifyCode is this author&#39;s stable proof token — the value a repository&#39;s verify file must carry. | [optional] 
**VerifyFile** | Pointer to **string** | VerifyFile is the repo-root file the file method reads, on the default branch. | [optional] 
**VerifySnippet** | Pointer to **string** | VerifySnippet is that file&#39;s exact contents, ready to commit. | [optional] 

## Methods

### NewAuthorEnrolment

`func NewAuthorEnrolment() *AuthorEnrolment`

NewAuthorEnrolment instantiates a new AuthorEnrolment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorEnrolmentWithDefaults

`func NewAuthorEnrolmentWithDefaults() *AuthorEnrolment`

NewAuthorEnrolmentWithDefaults instantiates a new AuthorEnrolment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *AuthorEnrolment) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AuthorEnrolment) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AuthorEnrolment) SetCreated(v bool)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AuthorEnrolment) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetGithubLogin

`func (o *AuthorEnrolment) GetGithubLogin() string`

GetGithubLogin returns the GithubLogin field if non-nil, zero value otherwise.

### GetGithubLoginOk

`func (o *AuthorEnrolment) GetGithubLoginOk() (*string, bool)`

GetGithubLoginOk returns a tuple with the GithubLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubLogin

`func (o *AuthorEnrolment) SetGithubLogin(v string)`

SetGithubLogin sets GithubLogin field to given value.

### HasGithubLogin

`func (o *AuthorEnrolment) HasGithubLogin() bool`

HasGithubLogin returns a boolean if a field has been set.

### GetId

`func (o *AuthorEnrolment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthorEnrolment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthorEnrolment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthorEnrolment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetShareBps

`func (o *AuthorEnrolment) GetShareBps() int64`

GetShareBps returns the ShareBps field if non-nil, zero value otherwise.

### GetShareBpsOk

`func (o *AuthorEnrolment) GetShareBpsOk() (*int64, bool)`

GetShareBpsOk returns a tuple with the ShareBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareBps

`func (o *AuthorEnrolment) SetShareBps(v int64)`

SetShareBps sets ShareBps field to given value.

### HasShareBps

`func (o *AuthorEnrolment) HasShareBps() bool`

HasShareBps returns a boolean if a field has been set.

### GetStatus

`func (o *AuthorEnrolment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthorEnrolment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthorEnrolment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthorEnrolment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVerified

`func (o *AuthorEnrolment) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *AuthorEnrolment) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *AuthorEnrolment) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *AuthorEnrolment) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetVerifyCode

`func (o *AuthorEnrolment) GetVerifyCode() string`

GetVerifyCode returns the VerifyCode field if non-nil, zero value otherwise.

### GetVerifyCodeOk

`func (o *AuthorEnrolment) GetVerifyCodeOk() (*string, bool)`

GetVerifyCodeOk returns a tuple with the VerifyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCode

`func (o *AuthorEnrolment) SetVerifyCode(v string)`

SetVerifyCode sets VerifyCode field to given value.

### HasVerifyCode

`func (o *AuthorEnrolment) HasVerifyCode() bool`

HasVerifyCode returns a boolean if a field has been set.

### GetVerifyFile

`func (o *AuthorEnrolment) GetVerifyFile() string`

GetVerifyFile returns the VerifyFile field if non-nil, zero value otherwise.

### GetVerifyFileOk

`func (o *AuthorEnrolment) GetVerifyFileOk() (*string, bool)`

GetVerifyFileOk returns a tuple with the VerifyFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyFile

`func (o *AuthorEnrolment) SetVerifyFile(v string)`

SetVerifyFile sets VerifyFile field to given value.

### HasVerifyFile

`func (o *AuthorEnrolment) HasVerifyFile() bool`

HasVerifyFile returns a boolean if a field has been set.

### GetVerifySnippet

`func (o *AuthorEnrolment) GetVerifySnippet() string`

GetVerifySnippet returns the VerifySnippet field if non-nil, zero value otherwise.

### GetVerifySnippetOk

`func (o *AuthorEnrolment) GetVerifySnippetOk() (*string, bool)`

GetVerifySnippetOk returns a tuple with the VerifySnippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySnippet

`func (o *AuthorEnrolment) SetVerifySnippet(v string)`

SetVerifySnippet sets VerifySnippet field to given value.

### HasVerifySnippet

`func (o *AuthorEnrolment) HasVerifySnippet() bool`

HasVerifySnippet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


