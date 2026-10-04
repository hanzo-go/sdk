# MarketplaceStatement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | Pointer to **string** | Agent is the agent it is about. | [optional] 
**Expires** | Pointer to **int64** | Expires is when it stops standing, unix seconds; zero for a W-9, which stands until the form changes. | [optional] 
**Form** | Pointer to **string** | Form is the class of tax form it rests on: w9, w8ben or w8bene. | [optional] 
**Id** | Pointer to **string** | ID is the statement&#39;s jti. | [optional] 
**Issued** | Pointer to **int64** | Issued is when it was signed, unix seconds. | [optional] 
**Jws** | Pointer to **string** | JWS is the signed statement itself — the credential an agent presents — answered only to the org it is about. | [optional] 
**Kind** | Pointer to **string** | Kind is what it states: tax. | [optional] 
**Org** | Pointer to **string** | Org is the principal the agent answers to. | [optional] 
**Revoked** | Pointer to **int64** | Revoked is when it was withdrawn, unix seconds; zero while it stands. | [optional] 
**Status** | Pointer to **string** | Status is live, expired or revoked. | [optional] 

## Methods

### NewMarketplaceStatement

`func NewMarketplaceStatement() *MarketplaceStatement`

NewMarketplaceStatement instantiates a new MarketplaceStatement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceStatementWithDefaults

`func NewMarketplaceStatementWithDefaults() *MarketplaceStatement`

NewMarketplaceStatementWithDefaults instantiates a new MarketplaceStatement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *MarketplaceStatement) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *MarketplaceStatement) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *MarketplaceStatement) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *MarketplaceStatement) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetExpires

`func (o *MarketplaceStatement) GetExpires() int64`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *MarketplaceStatement) GetExpiresOk() (*int64, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *MarketplaceStatement) SetExpires(v int64)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *MarketplaceStatement) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetForm

`func (o *MarketplaceStatement) GetForm() string`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *MarketplaceStatement) GetFormOk() (*string, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *MarketplaceStatement) SetForm(v string)`

SetForm sets Form field to given value.

### HasForm

`func (o *MarketplaceStatement) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetId

`func (o *MarketplaceStatement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketplaceStatement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketplaceStatement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketplaceStatement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssued

`func (o *MarketplaceStatement) GetIssued() int64`

GetIssued returns the Issued field if non-nil, zero value otherwise.

### GetIssuedOk

`func (o *MarketplaceStatement) GetIssuedOk() (*int64, bool)`

GetIssuedOk returns a tuple with the Issued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssued

`func (o *MarketplaceStatement) SetIssued(v int64)`

SetIssued sets Issued field to given value.

### HasIssued

`func (o *MarketplaceStatement) HasIssued() bool`

HasIssued returns a boolean if a field has been set.

### GetJws

`func (o *MarketplaceStatement) GetJws() string`

GetJws returns the Jws field if non-nil, zero value otherwise.

### GetJwsOk

`func (o *MarketplaceStatement) GetJwsOk() (*string, bool)`

GetJwsOk returns a tuple with the Jws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJws

`func (o *MarketplaceStatement) SetJws(v string)`

SetJws sets Jws field to given value.

### HasJws

`func (o *MarketplaceStatement) HasJws() bool`

HasJws returns a boolean if a field has been set.

### GetKind

`func (o *MarketplaceStatement) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MarketplaceStatement) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MarketplaceStatement) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *MarketplaceStatement) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetOrg

`func (o *MarketplaceStatement) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *MarketplaceStatement) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *MarketplaceStatement) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *MarketplaceStatement) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetRevoked

`func (o *MarketplaceStatement) GetRevoked() int64`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *MarketplaceStatement) GetRevokedOk() (*int64, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *MarketplaceStatement) SetRevoked(v int64)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *MarketplaceStatement) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.

### GetStatus

`func (o *MarketplaceStatement) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MarketplaceStatement) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MarketplaceStatement) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MarketplaceStatement) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


