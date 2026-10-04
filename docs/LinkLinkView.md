# LinkLinkView

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

### NewLinkLinkView

`func NewLinkLinkView() *LinkLinkView`

NewLinkLinkView instantiates a new LinkLinkView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkLinkViewWithDefaults

`func NewLinkLinkViewWithDefaults() *LinkLinkView`

NewLinkLinkViewWithDefaults instantiates a new LinkLinkView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *LinkLinkView) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *LinkLinkView) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *LinkLinkView) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *LinkLinkView) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetBilling

`func (o *LinkLinkView) GetBilling() string`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *LinkLinkView) GetBillingOk() (*string, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *LinkLinkView) SetBilling(v string)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *LinkLinkView) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LinkLinkView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LinkLinkView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LinkLinkView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LinkLinkView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHost

`func (o *LinkLinkView) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *LinkLinkView) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *LinkLinkView) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *LinkLinkView) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *LinkLinkView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LinkLinkView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LinkLinkView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LinkLinkView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *LinkLinkView) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *LinkLinkView) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *LinkLinkView) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *LinkLinkView) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLastSeen

`func (o *LinkLinkView) GetLastSeen() string`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *LinkLinkView) GetLastSeenOk() (*string, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *LinkLinkView) SetLastSeen(v string)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *LinkLinkView) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetMachine

`func (o *LinkLinkView) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *LinkLinkView) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *LinkLinkView) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *LinkLinkView) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetOs

`func (o *LinkLinkView) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *LinkLinkView) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *LinkLinkView) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *LinkLinkView) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetPlan

`func (o *LinkLinkView) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *LinkLinkView) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *LinkLinkView) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *LinkLinkView) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetProvider

`func (o *LinkLinkView) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *LinkLinkView) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *LinkLinkView) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *LinkLinkView) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetStatus

`func (o *LinkLinkView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LinkLinkView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LinkLinkView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LinkLinkView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *LinkLinkView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LinkLinkView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LinkLinkView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LinkLinkView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsage

`func (o *LinkLinkView) GetUsage() interface{}`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *LinkLinkView) GetUsageOk() (*interface{}, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *LinkLinkView) SetUsage(v interface{})`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *LinkLinkView) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### SetUsageNil

`func (o *LinkLinkView) SetUsageNil(b bool)`

 SetUsageNil sets the value for Usage to be an explicit nil

### UnsetUsage
`func (o *LinkLinkView) UnsetUsage()`

UnsetUsage ensures that no value is present for Usage, not even an explicit nil
### GetUser

`func (o *LinkLinkView) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LinkLinkView) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LinkLinkView) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *LinkLinkView) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


