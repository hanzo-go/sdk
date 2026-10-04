# LinkRouteCandidate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the provider-side account identifier. | [optional] 
**Available** | Pointer to **bool** | Available reports whether the candidate is routable right now. | [optional] 
**Billing** | Pointer to **string** | Billing is the cost consequence of dialing this candidate: plan (the user&#39;s own subscription) or commerce (the metered gateway path). | [optional] 
**HeadroomPct** | Pointer to **float64** | HeadroomPct is the remaining rate-limit capacity, 0..100. A link with no snapshot counts as full headroom. | [optional] 
**Host** | Pointer to **string** | Host is that machine&#39;s hostname label. | [optional] 
**Kind** | Pointer to **string** | Kind is how the account authenticates: subscription or apikey. | [optional] 
**LinkId** | Pointer to **string** | LinkID is the underlying link&#39;s opaque handle. | [optional] 
**Machine** | Pointer to **string** | Machine is the machine the account is signed in on. | [optional] 
**Plan** | Pointer to **string** | Plan is the provider plan label the account is on. | [optional] 
**Provider** | Pointer to **string** | Provider is the AI provider the candidate account belongs to. | [optional] 
**Reason** | Pointer to **string** | Reason says why the candidate is not routable, when Available is false. | [optional] 

## Methods

### NewLinkRouteCandidate

`func NewLinkRouteCandidate() *LinkRouteCandidate`

NewLinkRouteCandidate instantiates a new LinkRouteCandidate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkRouteCandidateWithDefaults

`func NewLinkRouteCandidateWithDefaults() *LinkRouteCandidate`

NewLinkRouteCandidateWithDefaults instantiates a new LinkRouteCandidate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *LinkRouteCandidate) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *LinkRouteCandidate) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *LinkRouteCandidate) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *LinkRouteCandidate) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAvailable

`func (o *LinkRouteCandidate) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *LinkRouteCandidate) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *LinkRouteCandidate) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *LinkRouteCandidate) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetBilling

`func (o *LinkRouteCandidate) GetBilling() string`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *LinkRouteCandidate) GetBillingOk() (*string, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *LinkRouteCandidate) SetBilling(v string)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *LinkRouteCandidate) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetHeadroomPct

`func (o *LinkRouteCandidate) GetHeadroomPct() float64`

GetHeadroomPct returns the HeadroomPct field if non-nil, zero value otherwise.

### GetHeadroomPctOk

`func (o *LinkRouteCandidate) GetHeadroomPctOk() (*float64, bool)`

GetHeadroomPctOk returns a tuple with the HeadroomPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadroomPct

`func (o *LinkRouteCandidate) SetHeadroomPct(v float64)`

SetHeadroomPct sets HeadroomPct field to given value.

### HasHeadroomPct

`func (o *LinkRouteCandidate) HasHeadroomPct() bool`

HasHeadroomPct returns a boolean if a field has been set.

### GetHost

`func (o *LinkRouteCandidate) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *LinkRouteCandidate) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *LinkRouteCandidate) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *LinkRouteCandidate) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetKind

`func (o *LinkRouteCandidate) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *LinkRouteCandidate) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *LinkRouteCandidate) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *LinkRouteCandidate) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLinkId

`func (o *LinkRouteCandidate) GetLinkId() string`

GetLinkId returns the LinkId field if non-nil, zero value otherwise.

### GetLinkIdOk

`func (o *LinkRouteCandidate) GetLinkIdOk() (*string, bool)`

GetLinkIdOk returns a tuple with the LinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkId

`func (o *LinkRouteCandidate) SetLinkId(v string)`

SetLinkId sets LinkId field to given value.

### HasLinkId

`func (o *LinkRouteCandidate) HasLinkId() bool`

HasLinkId returns a boolean if a field has been set.

### GetMachine

`func (o *LinkRouteCandidate) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *LinkRouteCandidate) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *LinkRouteCandidate) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *LinkRouteCandidate) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetPlan

`func (o *LinkRouteCandidate) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *LinkRouteCandidate) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *LinkRouteCandidate) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *LinkRouteCandidate) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetProvider

`func (o *LinkRouteCandidate) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *LinkRouteCandidate) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *LinkRouteCandidate) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *LinkRouteCandidate) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetReason

`func (o *LinkRouteCandidate) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *LinkRouteCandidate) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *LinkRouteCandidate) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *LinkRouteCandidate) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


