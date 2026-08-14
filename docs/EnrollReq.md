# EnrollReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the provider-side account identifier. | [optional] 
**Host** | Pointer to **string** | Host is the machine&#39;s human hostname label. | [optional] 
**Kind** | Pointer to **string** | Kind decides how the account&#39;s inference BILLS and defaults to subscription: a subscription account bills the user&#39;s own monthly plan and is metered here for visibility only, while an apikey account bills through commerce on the gateway path. | [optional] 
**Machine** | Pointer to **string** | Machine is the stable machine identifier. Required, length-bounded. | [optional] 
**Os** | Pointer to **string** | OS is the machine&#39;s operating system label. | [optional] 
**Plan** | Pointer to **string** | Plan is the provider plan label (e.g. \&quot;Claude Max\&quot;). | [optional] 
**Provider** | Pointer to **string** | Provider is the AI provider the account belongs to. Required, length-bounded. | [optional] 
**Usage** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewEnrollReq

`func NewEnrollReq() *EnrollReq`

NewEnrollReq instantiates a new EnrollReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollReqWithDefaults

`func NewEnrollReqWithDefaults() *EnrollReq`

NewEnrollReqWithDefaults instantiates a new EnrollReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *EnrollReq) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *EnrollReq) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *EnrollReq) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *EnrollReq) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetHost

`func (o *EnrollReq) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *EnrollReq) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *EnrollReq) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *EnrollReq) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetKind

`func (o *EnrollReq) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EnrollReq) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EnrollReq) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *EnrollReq) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMachine

`func (o *EnrollReq) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *EnrollReq) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *EnrollReq) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *EnrollReq) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetOs

`func (o *EnrollReq) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *EnrollReq) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *EnrollReq) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *EnrollReq) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetPlan

`func (o *EnrollReq) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *EnrollReq) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *EnrollReq) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *EnrollReq) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetProvider

`func (o *EnrollReq) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *EnrollReq) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *EnrollReq) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *EnrollReq) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetUsage

`func (o *EnrollReq) GetUsage() interface{}`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *EnrollReq) GetUsageOk() (*interface{}, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *EnrollReq) SetUsage(v interface{})`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *EnrollReq) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### SetUsageNil

`func (o *EnrollReq) SetUsageNil(b bool)`

 SetUsageNil sets the value for Usage to be an explicit nil

### UnsetUsage
`func (o *EnrollReq) UnsetUsage()`

UnsetUsage ensures that no value is present for Usage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


