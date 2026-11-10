# GitWorkspaceView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | Pointer to **string** | Login is the caller&#39;s forge account in a made workspace. | [optional] 
**Made** | Pointer to **bool** | Made says the platform made it for the org; false is the estate&#39;s. | [optional] 
**Owner** | Pointer to **string** | Owner is the forge namespace. | [optional] 

## Methods

### NewGitWorkspaceView

`func NewGitWorkspaceView() *GitWorkspaceView`

NewGitWorkspaceView instantiates a new GitWorkspaceView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitWorkspaceViewWithDefaults

`func NewGitWorkspaceViewWithDefaults() *GitWorkspaceView`

NewGitWorkspaceViewWithDefaults instantiates a new GitWorkspaceView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *GitWorkspaceView) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *GitWorkspaceView) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *GitWorkspaceView) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *GitWorkspaceView) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetMade

`func (o *GitWorkspaceView) GetMade() bool`

GetMade returns the Made field if non-nil, zero value otherwise.

### GetMadeOk

`func (o *GitWorkspaceView) GetMadeOk() (*bool, bool)`

GetMadeOk returns a tuple with the Made field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMade

`func (o *GitWorkspaceView) SetMade(v bool)`

SetMade sets Made field to given value.

### HasMade

`func (o *GitWorkspaceView) HasMade() bool`

HasMade returns a boolean if a field has been set.

### GetOwner

`func (o *GitWorkspaceView) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GitWorkspaceView) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GitWorkspaceView) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GitWorkspaceView) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


