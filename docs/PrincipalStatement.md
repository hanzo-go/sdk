# PrincipalStatement

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

### NewPrincipalStatement

`func NewPrincipalStatement() *PrincipalStatement`

NewPrincipalStatement instantiates a new PrincipalStatement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalStatementWithDefaults

`func NewPrincipalStatementWithDefaults() *PrincipalStatement`

NewPrincipalStatementWithDefaults instantiates a new PrincipalStatement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *PrincipalStatement) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *PrincipalStatement) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *PrincipalStatement) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *PrincipalStatement) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetExpires

`func (o *PrincipalStatement) GetExpires() int64`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *PrincipalStatement) GetExpiresOk() (*int64, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *PrincipalStatement) SetExpires(v int64)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *PrincipalStatement) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetForm

`func (o *PrincipalStatement) GetForm() string`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *PrincipalStatement) GetFormOk() (*string, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *PrincipalStatement) SetForm(v string)`

SetForm sets Form field to given value.

### HasForm

`func (o *PrincipalStatement) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetId

`func (o *PrincipalStatement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrincipalStatement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrincipalStatement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PrincipalStatement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssued

`func (o *PrincipalStatement) GetIssued() int64`

GetIssued returns the Issued field if non-nil, zero value otherwise.

### GetIssuedOk

`func (o *PrincipalStatement) GetIssuedOk() (*int64, bool)`

GetIssuedOk returns a tuple with the Issued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssued

`func (o *PrincipalStatement) SetIssued(v int64)`

SetIssued sets Issued field to given value.

### HasIssued

`func (o *PrincipalStatement) HasIssued() bool`

HasIssued returns a boolean if a field has been set.

### GetJws

`func (o *PrincipalStatement) GetJws() string`

GetJws returns the Jws field if non-nil, zero value otherwise.

### GetJwsOk

`func (o *PrincipalStatement) GetJwsOk() (*string, bool)`

GetJwsOk returns a tuple with the Jws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJws

`func (o *PrincipalStatement) SetJws(v string)`

SetJws sets Jws field to given value.

### HasJws

`func (o *PrincipalStatement) HasJws() bool`

HasJws returns a boolean if a field has been set.

### GetKind

`func (o *PrincipalStatement) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PrincipalStatement) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PrincipalStatement) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PrincipalStatement) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetOrg

`func (o *PrincipalStatement) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *PrincipalStatement) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *PrincipalStatement) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *PrincipalStatement) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetRevoked

`func (o *PrincipalStatement) GetRevoked() int64`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *PrincipalStatement) GetRevokedOk() (*int64, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *PrincipalStatement) SetRevoked(v int64)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *PrincipalStatement) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.

### GetStatus

`func (o *PrincipalStatement) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrincipalStatement) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrincipalStatement) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PrincipalStatement) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


