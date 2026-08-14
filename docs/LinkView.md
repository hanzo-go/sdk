# LinkView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the provider-side account identifier, when the collector knows it. | [optional] 
**Billing** | Pointer to **string** | Billing is how this account&#39;s inference bills — plan (the user&#39;s own subscription, metered here for visibility only) or commerce (the gateway path). Derived from Kind, never stored. | [optional] 
**CreatedAt** | Pointer to **string** | CreatedAt is when the link was first registered, RFC 3339 UTC. | [optional] 
**Host** | Pointer to **string** | Host is the machine&#39;s human hostname label, from its most recent report. | [optional] 
**Id** | Pointer to **string** | ID is the link&#39;s opaque handle (\&quot;link_\&quot; + 32 hex chars). | [optional] 
**Kind** | Pointer to **string** | Kind is how the account authenticates: subscription or apikey. | [optional] 
**LastSeen** | Pointer to **string** | LastSeen is when the account last reported, RFC 3339 UTC. | [optional] 
**Machine** | Pointer to **string** | Machine is the stable machine identifier the collector reports. | [optional] 
**Os** | Pointer to **string** | OS is the machine&#39;s operating system label. | [optional] 
**Plan** | Pointer to **string** | Plan is the provider plan label (e.g. \&quot;Claude Max\&quot;). | [optional] 
**Provider** | Pointer to **string** | Provider is the AI provider this account belongs to (claude, openai, hanzo…). | [optional] 
**Status** | Pointer to **string** | Status is linked or revoked. Revoked rows are retained for history. | [optional] 
**UpdatedAt** | Pointer to **string** | UpdatedAt is when the link was last refreshed, RFC 3339 UTC. | [optional] 
**Usage** | Pointer to **interface{}** |  | [optional] 
**User** | Pointer to **string** | User is the owning subject — the validated caller who registered the link. | [optional] 

## Methods

### NewLinkView

`func NewLinkView() *LinkView`

NewLinkView instantiates a new LinkView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkViewWithDefaults

`func NewLinkViewWithDefaults() *LinkView`

NewLinkViewWithDefaults instantiates a new LinkView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *LinkView) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *LinkView) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *LinkView) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *LinkView) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetBilling

`func (o *LinkView) GetBilling() string`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *LinkView) GetBillingOk() (*string, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *LinkView) SetBilling(v string)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *LinkView) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LinkView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LinkView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LinkView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LinkView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHost

`func (o *LinkView) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *LinkView) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *LinkView) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *LinkView) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *LinkView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LinkView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LinkView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LinkView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *LinkView) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *LinkView) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *LinkView) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *LinkView) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLastSeen

`func (o *LinkView) GetLastSeen() string`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *LinkView) GetLastSeenOk() (*string, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *LinkView) SetLastSeen(v string)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *LinkView) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetMachine

`func (o *LinkView) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *LinkView) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *LinkView) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *LinkView) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetOs

`func (o *LinkView) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *LinkView) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *LinkView) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *LinkView) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetPlan

`func (o *LinkView) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *LinkView) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *LinkView) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *LinkView) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetProvider

`func (o *LinkView) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *LinkView) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *LinkView) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *LinkView) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetStatus

`func (o *LinkView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LinkView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LinkView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LinkView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *LinkView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LinkView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LinkView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LinkView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsage

`func (o *LinkView) GetUsage() interface{}`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *LinkView) GetUsageOk() (*interface{}, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *LinkView) SetUsage(v interface{})`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *LinkView) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### SetUsageNil

`func (o *LinkView) SetUsageNil(b bool)`

 SetUsageNil sets the value for Usage to be an explicit nil

### UnsetUsage
`func (o *LinkView) UnsetUsage()`

UnsetUsage ensures that no value is present for Usage, not even an explicit nil
### GetUser

`func (o *LinkView) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LinkView) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LinkView) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *LinkView) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


