# GithubClaimIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to **[]string** | Accounts names GitHub logins the App is installed on (\&quot;hanzoai\&quot;). Matched case-insensitively, since GitHub logins are. Ignored when all is true. | [optional] 
**All** | Pointer to **bool** | All binds every account the App holds, instead of naming them. | [optional] 

## Methods

### NewGithubClaimIn

`func NewGithubClaimIn() *GithubClaimIn`

NewGithubClaimIn instantiates a new GithubClaimIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubClaimInWithDefaults

`func NewGithubClaimInWithDefaults() *GithubClaimIn`

NewGithubClaimInWithDefaults instantiates a new GithubClaimIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *GithubClaimIn) GetAccounts() []string`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *GithubClaimIn) GetAccountsOk() (*[]string, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *GithubClaimIn) SetAccounts(v []string)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *GithubClaimIn) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetAll

`func (o *GithubClaimIn) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *GithubClaimIn) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *GithubClaimIn) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *GithubClaimIn) HasAll() bool`

HasAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


