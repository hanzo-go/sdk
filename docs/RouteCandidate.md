# RouteCandidate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the provider-side account identifier. | [optional] 
**Available** | Pointer to **bool** | Available reports whether the candidate is routable right now. | [optional] 
**Billing** | Pointer to **string** | Billing is the cost consequence of dialing this candidate: plan (the user&#39;s own subscription) or commerce (the metered gateway path). | [optional] 
**HeadroomPct** | Pointer to **float32** | HeadroomPct is the remaining rate-limit capacity, 0..100. A link with no snapshot counts as full headroom. | [optional] 
**Host** | Pointer to **string** | Host is that machine&#39;s hostname label. | [optional] 
**Kind** | Pointer to **string** | Kind is how the account authenticates: subscription or apikey. | [optional] 
**LinkId** | Pointer to **string** | LinkID is the underlying link&#39;s opaque handle. | [optional] 
**Machine** | Pointer to **string** | Machine is the machine the account is signed in on. | [optional] 
**Plan** | Pointer to **string** | Plan is the provider plan label the account is on. | [optional] 
**Provider** | Pointer to **string** | Provider is the AI provider the candidate account belongs to. | [optional] 
**Reason** | Pointer to **string** | Reason says why the candidate is not routable, when Available is false. | [optional] 

## Methods

### NewRouteCandidate

`func NewRouteCandidate() *RouteCandidate`

NewRouteCandidate instantiates a new RouteCandidate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteCandidateWithDefaults

`func NewRouteCandidateWithDefaults() *RouteCandidate`

NewRouteCandidateWithDefaults instantiates a new RouteCandidate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *RouteCandidate) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *RouteCandidate) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *RouteCandidate) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *RouteCandidate) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAvailable

`func (o *RouteCandidate) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *RouteCandidate) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *RouteCandidate) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *RouteCandidate) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetBilling

`func (o *RouteCandidate) GetBilling() string`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *RouteCandidate) GetBillingOk() (*string, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *RouteCandidate) SetBilling(v string)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *RouteCandidate) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetHeadroomPct

`func (o *RouteCandidate) GetHeadroomPct() float32`

GetHeadroomPct returns the HeadroomPct field if non-nil, zero value otherwise.

### GetHeadroomPctOk

`func (o *RouteCandidate) GetHeadroomPctOk() (*float32, bool)`

GetHeadroomPctOk returns a tuple with the HeadroomPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadroomPct

`func (o *RouteCandidate) SetHeadroomPct(v float32)`

SetHeadroomPct sets HeadroomPct field to given value.

### HasHeadroomPct

`func (o *RouteCandidate) HasHeadroomPct() bool`

HasHeadroomPct returns a boolean if a field has been set.

### GetHost

`func (o *RouteCandidate) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RouteCandidate) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RouteCandidate) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *RouteCandidate) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetKind

`func (o *RouteCandidate) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RouteCandidate) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RouteCandidate) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RouteCandidate) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLinkId

`func (o *RouteCandidate) GetLinkId() string`

GetLinkId returns the LinkId field if non-nil, zero value otherwise.

### GetLinkIdOk

`func (o *RouteCandidate) GetLinkIdOk() (*string, bool)`

GetLinkIdOk returns a tuple with the LinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkId

`func (o *RouteCandidate) SetLinkId(v string)`

SetLinkId sets LinkId field to given value.

### HasLinkId

`func (o *RouteCandidate) HasLinkId() bool`

HasLinkId returns a boolean if a field has been set.

### GetMachine

`func (o *RouteCandidate) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *RouteCandidate) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *RouteCandidate) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *RouteCandidate) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetPlan

`func (o *RouteCandidate) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *RouteCandidate) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *RouteCandidate) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *RouteCandidate) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetProvider

`func (o *RouteCandidate) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *RouteCandidate) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *RouteCandidate) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *RouteCandidate) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetReason

`func (o *RouteCandidate) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RouteCandidate) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RouteCandidate) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RouteCandidate) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


