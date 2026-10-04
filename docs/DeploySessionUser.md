# DeploySessionUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to **[]string** | Groups is the caller&#39;s group list, always empty here: this console authorizes on the platform SuperAdmin fact alone, not on argocd RBAC groups. Absent for an anonymous caller. | [optional] 
**Iss** | Pointer to **string** | Iss is the token issuer as the SPA expects to see it — the literal \&quot;argocd\&quot;, so the UI never triggers an SSO redirect of its own. Absent for an anonymous caller. | [optional] 
**LoggedIn** | Pointer to **bool** | LoggedIn reports whether this browser holds a session this console accepts. | [optional] 
**LoginUrl** | Pointer to **string** | LoginURL is where an anonymous caller signs in. Absent once signed in. | [optional] 
**LogoutUrl** | Pointer to **string** | LogoutURL is where a signed-in caller ends the session. Absent when anonymous. | [optional] 
**Username** | Pointer to **string** | Username is the validated principal&#39;s user ID — the opaque gateway id, which is what argocd&#39;s UI renders as the signed-in user here — or \&quot;admin\&quot; when the principal carries none. Absent when anonymous. | [optional] 

## Methods

### NewDeploySessionUser

`func NewDeploySessionUser() *DeploySessionUser`

NewDeploySessionUser instantiates a new DeploySessionUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploySessionUserWithDefaults

`func NewDeploySessionUserWithDefaults() *DeploySessionUser`

NewDeploySessionUserWithDefaults instantiates a new DeploySessionUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *DeploySessionUser) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *DeploySessionUser) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *DeploySessionUser) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *DeploySessionUser) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetIss

`func (o *DeploySessionUser) GetIss() string`

GetIss returns the Iss field if non-nil, zero value otherwise.

### GetIssOk

`func (o *DeploySessionUser) GetIssOk() (*string, bool)`

GetIssOk returns a tuple with the Iss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIss

`func (o *DeploySessionUser) SetIss(v string)`

SetIss sets Iss field to given value.

### HasIss

`func (o *DeploySessionUser) HasIss() bool`

HasIss returns a boolean if a field has been set.

### GetLoggedIn

`func (o *DeploySessionUser) GetLoggedIn() bool`

GetLoggedIn returns the LoggedIn field if non-nil, zero value otherwise.

### GetLoggedInOk

`func (o *DeploySessionUser) GetLoggedInOk() (*bool, bool)`

GetLoggedInOk returns a tuple with the LoggedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggedIn

`func (o *DeploySessionUser) SetLoggedIn(v bool)`

SetLoggedIn sets LoggedIn field to given value.

### HasLoggedIn

`func (o *DeploySessionUser) HasLoggedIn() bool`

HasLoggedIn returns a boolean if a field has been set.

### GetLoginUrl

`func (o *DeploySessionUser) GetLoginUrl() string`

GetLoginUrl returns the LoginUrl field if non-nil, zero value otherwise.

### GetLoginUrlOk

`func (o *DeploySessionUser) GetLoginUrlOk() (*string, bool)`

GetLoginUrlOk returns a tuple with the LoginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginUrl

`func (o *DeploySessionUser) SetLoginUrl(v string)`

SetLoginUrl sets LoginUrl field to given value.

### HasLoginUrl

`func (o *DeploySessionUser) HasLoginUrl() bool`

HasLoginUrl returns a boolean if a field has been set.

### GetLogoutUrl

`func (o *DeploySessionUser) GetLogoutUrl() string`

GetLogoutUrl returns the LogoutUrl field if non-nil, zero value otherwise.

### GetLogoutUrlOk

`func (o *DeploySessionUser) GetLogoutUrlOk() (*string, bool)`

GetLogoutUrlOk returns a tuple with the LogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrl

`func (o *DeploySessionUser) SetLogoutUrl(v string)`

SetLogoutUrl sets LogoutUrl field to given value.

### HasLogoutUrl

`func (o *DeploySessionUser) HasLogoutUrl() bool`

HasLogoutUrl returns a boolean if a field has been set.

### GetUsername

`func (o *DeploySessionUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DeploySessionUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DeploySessionUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DeploySessionUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


