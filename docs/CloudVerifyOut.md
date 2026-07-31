# CloudVerifyOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the account label the provider reported. Present only when active. | [optional] 
**Active** | Pointer to **bool** | Active is whether the stored credential verified live against the provider. | [optional] 
**ExternalId** | Pointer to **string** | ExternalID is the provider&#39;s account id. Present only when active. | [optional] 
**Provider** | Pointer to **string** | Provider is the connector&#39;s registry id. | [optional] 
**Reason** | Pointer to **string** | Reason is why the check failed. Present only when active is false. | [optional] 
**Scopes** | Pointer to **[]string** | Scopes are the permissions the credential carries. Present only when active. | [optional] 

## Methods

### NewCloudVerifyOut

`func NewCloudVerifyOut() *CloudVerifyOut`

NewCloudVerifyOut instantiates a new CloudVerifyOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudVerifyOutWithDefaults

`func NewCloudVerifyOutWithDefaults() *CloudVerifyOut`

NewCloudVerifyOutWithDefaults instantiates a new CloudVerifyOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CloudVerifyOut) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CloudVerifyOut) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CloudVerifyOut) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CloudVerifyOut) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetActive

`func (o *CloudVerifyOut) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CloudVerifyOut) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CloudVerifyOut) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CloudVerifyOut) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetExternalId

`func (o *CloudVerifyOut) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CloudVerifyOut) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CloudVerifyOut) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CloudVerifyOut) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetProvider

`func (o *CloudVerifyOut) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudVerifyOut) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudVerifyOut) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudVerifyOut) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetReason

`func (o *CloudVerifyOut) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CloudVerifyOut) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CloudVerifyOut) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CloudVerifyOut) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetScopes

`func (o *CloudVerifyOut) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *CloudVerifyOut) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *CloudVerifyOut) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *CloudVerifyOut) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


