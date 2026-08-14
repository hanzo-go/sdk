# GithubInstallationView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connected** | Pointer to **bool** | Connected reports whether THIS org has already bound this account. | [optional] 
**Grant** | Pointer to **string** | Grant is \&quot;all\&quot; or \&quot;selected\&quot; — how many of the account&#39;s repositories the install covers. A reader deciding what to import needs the reach, not just the name. | [optional] 
**HtmlUrl** | Pointer to **string** | HTMLURL is the account&#39;s page on GitHub. | [optional] 
**Login** | Pointer to **string** | Login is the GitHub account name — the org or user the App is installed on. | [optional] 
**Type** | Pointer to **string** | Type is \&quot;Organization\&quot; or \&quot;User\&quot;. | [optional] 

## Methods

### NewGithubInstallationView

`func NewGithubInstallationView() *GithubInstallationView`

NewGithubInstallationView instantiates a new GithubInstallationView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubInstallationViewWithDefaults

`func NewGithubInstallationViewWithDefaults() *GithubInstallationView`

NewGithubInstallationViewWithDefaults instantiates a new GithubInstallationView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnected

`func (o *GithubInstallationView) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *GithubInstallationView) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *GithubInstallationView) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *GithubInstallationView) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetGrant

`func (o *GithubInstallationView) GetGrant() string`

GetGrant returns the Grant field if non-nil, zero value otherwise.

### GetGrantOk

`func (o *GithubInstallationView) GetGrantOk() (*string, bool)`

GetGrantOk returns a tuple with the Grant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrant

`func (o *GithubInstallationView) SetGrant(v string)`

SetGrant sets Grant field to given value.

### HasGrant

`func (o *GithubInstallationView) HasGrant() bool`

HasGrant returns a boolean if a field has been set.

### GetHtmlUrl

`func (o *GithubInstallationView) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *GithubInstallationView) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *GithubInstallationView) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.

### HasHtmlUrl

`func (o *GithubInstallationView) HasHtmlUrl() bool`

HasHtmlUrl returns a boolean if a field has been set.

### GetLogin

`func (o *GithubInstallationView) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *GithubInstallationView) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *GithubInstallationView) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *GithubInstallationView) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetType

`func (o *GithubInstallationView) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GithubInstallationView) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GithubInstallationView) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GithubInstallationView) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


