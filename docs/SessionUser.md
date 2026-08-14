# SessionUser

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

### NewSessionUser

`func NewSessionUser() *SessionUser`

NewSessionUser instantiates a new SessionUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionUserWithDefaults

`func NewSessionUserWithDefaults() *SessionUser`

NewSessionUserWithDefaults instantiates a new SessionUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *SessionUser) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *SessionUser) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *SessionUser) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *SessionUser) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetIss

`func (o *SessionUser) GetIss() string`

GetIss returns the Iss field if non-nil, zero value otherwise.

### GetIssOk

`func (o *SessionUser) GetIssOk() (*string, bool)`

GetIssOk returns a tuple with the Iss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIss

`func (o *SessionUser) SetIss(v string)`

SetIss sets Iss field to given value.

### HasIss

`func (o *SessionUser) HasIss() bool`

HasIss returns a boolean if a field has been set.

### GetLoggedIn

`func (o *SessionUser) GetLoggedIn() bool`

GetLoggedIn returns the LoggedIn field if non-nil, zero value otherwise.

### GetLoggedInOk

`func (o *SessionUser) GetLoggedInOk() (*bool, bool)`

GetLoggedInOk returns a tuple with the LoggedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggedIn

`func (o *SessionUser) SetLoggedIn(v bool)`

SetLoggedIn sets LoggedIn field to given value.

### HasLoggedIn

`func (o *SessionUser) HasLoggedIn() bool`

HasLoggedIn returns a boolean if a field has been set.

### GetLoginUrl

`func (o *SessionUser) GetLoginUrl() string`

GetLoginUrl returns the LoginUrl field if non-nil, zero value otherwise.

### GetLoginUrlOk

`func (o *SessionUser) GetLoginUrlOk() (*string, bool)`

GetLoginUrlOk returns a tuple with the LoginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginUrl

`func (o *SessionUser) SetLoginUrl(v string)`

SetLoginUrl sets LoginUrl field to given value.

### HasLoginUrl

`func (o *SessionUser) HasLoginUrl() bool`

HasLoginUrl returns a boolean if a field has been set.

### GetLogoutUrl

`func (o *SessionUser) GetLogoutUrl() string`

GetLogoutUrl returns the LogoutUrl field if non-nil, zero value otherwise.

### GetLogoutUrlOk

`func (o *SessionUser) GetLogoutUrlOk() (*string, bool)`

GetLogoutUrlOk returns a tuple with the LogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrl

`func (o *SessionUser) SetLogoutUrl(v string)`

SetLogoutUrl sets LogoutUrl field to given value.

### HasLogoutUrl

`func (o *SessionUser) HasLogoutUrl() bool`

HasLogoutUrl returns a boolean if a field has been set.

### GetUsername

`func (o *SessionUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SessionUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SessionUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SessionUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


