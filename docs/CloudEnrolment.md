# CloudEnrolment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **bool** | Created reports whether this call enrolled the org (201) or found an existing enrolment (200). | [optional] 
**GithubLogin** | Pointer to **string** | GithubLogin is the linked forge account. | [optional] 
**Id** | Pointer to **string** | ID is the author record&#39;s server-minted handle, \&quot;aut_\&quot;-prefixed. | [optional] 
**ShareBps** | Pointer to **int32** | ShareBps is this author&#39;s royalty share in basis points of the spend their deployed work generates. | [optional] 
**Status** | Pointer to **string** | Status is connected, approved or suspended. Only an approved author earns. | [optional] 
**Verified** | Pointer to **bool** | Verified reports whether any repository or owner claim has been proven yet. | [optional] 
**VerifyCode** | Pointer to **string** | VerifyCode is this author&#39;s stable proof token — the value a repository&#39;s verify file must carry. | [optional] 
**VerifyFile** | Pointer to **string** | VerifyFile is the repo-root file the file method reads, on the default branch. | [optional] 
**VerifySnippet** | Pointer to **string** | VerifySnippet is that file&#39;s exact contents, ready to commit. | [optional] 

## Methods

### NewCloudEnrolment

`func NewCloudEnrolment() *CloudEnrolment`

NewCloudEnrolment instantiates a new CloudEnrolment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudEnrolmentWithDefaults

`func NewCloudEnrolmentWithDefaults() *CloudEnrolment`

NewCloudEnrolmentWithDefaults instantiates a new CloudEnrolment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *CloudEnrolment) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CloudEnrolment) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CloudEnrolment) SetCreated(v bool)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CloudEnrolment) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetGithubLogin

`func (o *CloudEnrolment) GetGithubLogin() string`

GetGithubLogin returns the GithubLogin field if non-nil, zero value otherwise.

### GetGithubLoginOk

`func (o *CloudEnrolment) GetGithubLoginOk() (*string, bool)`

GetGithubLoginOk returns a tuple with the GithubLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubLogin

`func (o *CloudEnrolment) SetGithubLogin(v string)`

SetGithubLogin sets GithubLogin field to given value.

### HasGithubLogin

`func (o *CloudEnrolment) HasGithubLogin() bool`

HasGithubLogin returns a boolean if a field has been set.

### GetId

`func (o *CloudEnrolment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudEnrolment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudEnrolment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudEnrolment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetShareBps

`func (o *CloudEnrolment) GetShareBps() int32`

GetShareBps returns the ShareBps field if non-nil, zero value otherwise.

### GetShareBpsOk

`func (o *CloudEnrolment) GetShareBpsOk() (*int32, bool)`

GetShareBpsOk returns a tuple with the ShareBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareBps

`func (o *CloudEnrolment) SetShareBps(v int32)`

SetShareBps sets ShareBps field to given value.

### HasShareBps

`func (o *CloudEnrolment) HasShareBps() bool`

HasShareBps returns a boolean if a field has been set.

### GetStatus

`func (o *CloudEnrolment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudEnrolment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudEnrolment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudEnrolment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVerified

`func (o *CloudEnrolment) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *CloudEnrolment) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *CloudEnrolment) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *CloudEnrolment) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetVerifyCode

`func (o *CloudEnrolment) GetVerifyCode() string`

GetVerifyCode returns the VerifyCode field if non-nil, zero value otherwise.

### GetVerifyCodeOk

`func (o *CloudEnrolment) GetVerifyCodeOk() (*string, bool)`

GetVerifyCodeOk returns a tuple with the VerifyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCode

`func (o *CloudEnrolment) SetVerifyCode(v string)`

SetVerifyCode sets VerifyCode field to given value.

### HasVerifyCode

`func (o *CloudEnrolment) HasVerifyCode() bool`

HasVerifyCode returns a boolean if a field has been set.

### GetVerifyFile

`func (o *CloudEnrolment) GetVerifyFile() string`

GetVerifyFile returns the VerifyFile field if non-nil, zero value otherwise.

### GetVerifyFileOk

`func (o *CloudEnrolment) GetVerifyFileOk() (*string, bool)`

GetVerifyFileOk returns a tuple with the VerifyFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyFile

`func (o *CloudEnrolment) SetVerifyFile(v string)`

SetVerifyFile sets VerifyFile field to given value.

### HasVerifyFile

`func (o *CloudEnrolment) HasVerifyFile() bool`

HasVerifyFile returns a boolean if a field has been set.

### GetVerifySnippet

`func (o *CloudEnrolment) GetVerifySnippet() string`

GetVerifySnippet returns the VerifySnippet field if non-nil, zero value otherwise.

### GetVerifySnippetOk

`func (o *CloudEnrolment) GetVerifySnippetOk() (*string, bool)`

GetVerifySnippetOk returns a tuple with the VerifySnippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySnippet

`func (o *CloudEnrolment) SetVerifySnippet(v string)`

SetVerifySnippet sets VerifySnippet field to given value.

### HasVerifySnippet

`func (o *CloudEnrolment) HasVerifySnippet() bool`

HasVerifySnippet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


