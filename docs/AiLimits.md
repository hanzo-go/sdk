# AiLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]AiAction**](AiAction.md) | Actions are the ways on: upgrade, add prepaid credit. | [optional] 
**Classes** | Pointer to [**map[string]AiClass**](AiClass.md) | Classes are premium (third-party frontier models) and ours (Hanzo&#39;s), each present when the plan includes it. | [optional] 
**CreditsAfterAllowance** | Pointer to **bool** | CreditsAfterAllowance is the payer&#39;s choice to keep using a model on credits once the plan&#39;s included usage of it is spent (PUT /v1/ai/limits sets it). | [optional] 
**Day** | Pointer to [**AiWindow**](AiWindow.md) |  | [optional] 
**Limited** | Pointer to [**AiLimited**](AiLimited.md) | Limited is present in limited mode. | [optional] 
**Paused** | Pointer to [**[]AiPaused**](AiPaused.md) | Paused are the models whose share of the plan is used for now: each is answered by its fallback in chat until its share resets. | [optional] 
**PeriodEnd** | Pointer to **string** |  | [optional] 
**PeriodStart** | Pointer to **string** | PeriodStart and PeriodEnd bound the billing period (RFC3339). | [optional] 
**Plan** | Pointer to **string** | Plan is the plan the caller is served as (\&quot;dev\&quot;, \&quot;max-5x\&quot;, \&quot;max-20x\&quot;, \&quot;team\&quot;, \&quot;team-annual\&quot;, \&quot;agency\&quot;, \&quot;advisory\&quot;, \&quot;dedicated\&quot;, ...), \&quot;free\&quot; when none counts. | [optional] 
**Session** | Pointer to [**AiWindow**](AiWindow.md) | Session and Day are the request windows over every request the plan covers. | [optional] 
**State** | Pointer to **string** | State is the worst class&#39;s: ok, near or limited. | [optional] 
**Upgrade** | Pointer to **string** | Upgrade is the plan that raises these limits, absent at the top. | [optional] 

## Methods

### NewAiLimits

`func NewAiLimits() *AiLimits`

NewAiLimits instantiates a new AiLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiLimitsWithDefaults

`func NewAiLimitsWithDefaults() *AiLimits`

NewAiLimitsWithDefaults instantiates a new AiLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *AiLimits) GetActions() []AiAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AiLimits) GetActionsOk() (*[]AiAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AiLimits) SetActions(v []AiAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *AiLimits) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetClasses

`func (o *AiLimits) GetClasses() map[string]AiClass`

GetClasses returns the Classes field if non-nil, zero value otherwise.

### GetClassesOk

`func (o *AiLimits) GetClassesOk() (*map[string]AiClass, bool)`

GetClassesOk returns a tuple with the Classes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClasses

`func (o *AiLimits) SetClasses(v map[string]AiClass)`

SetClasses sets Classes field to given value.

### HasClasses

`func (o *AiLimits) HasClasses() bool`

HasClasses returns a boolean if a field has been set.

### GetCreditsAfterAllowance

`func (o *AiLimits) GetCreditsAfterAllowance() bool`

GetCreditsAfterAllowance returns the CreditsAfterAllowance field if non-nil, zero value otherwise.

### GetCreditsAfterAllowanceOk

`func (o *AiLimits) GetCreditsAfterAllowanceOk() (*bool, bool)`

GetCreditsAfterAllowanceOk returns a tuple with the CreditsAfterAllowance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsAfterAllowance

`func (o *AiLimits) SetCreditsAfterAllowance(v bool)`

SetCreditsAfterAllowance sets CreditsAfterAllowance field to given value.

### HasCreditsAfterAllowance

`func (o *AiLimits) HasCreditsAfterAllowance() bool`

HasCreditsAfterAllowance returns a boolean if a field has been set.

### GetDay

`func (o *AiLimits) GetDay() AiWindow`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *AiLimits) GetDayOk() (*AiWindow, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *AiLimits) SetDay(v AiWindow)`

SetDay sets Day field to given value.

### HasDay

`func (o *AiLimits) HasDay() bool`

HasDay returns a boolean if a field has been set.

### GetLimited

`func (o *AiLimits) GetLimited() AiLimited`

GetLimited returns the Limited field if non-nil, zero value otherwise.

### GetLimitedOk

`func (o *AiLimits) GetLimitedOk() (*AiLimited, bool)`

GetLimitedOk returns a tuple with the Limited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimited

`func (o *AiLimits) SetLimited(v AiLimited)`

SetLimited sets Limited field to given value.

### HasLimited

`func (o *AiLimits) HasLimited() bool`

HasLimited returns a boolean if a field has been set.

### GetPaused

`func (o *AiLimits) GetPaused() []AiPaused`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *AiLimits) GetPausedOk() (*[]AiPaused, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *AiLimits) SetPaused(v []AiPaused)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *AiLimits) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *AiLimits) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *AiLimits) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *AiLimits) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *AiLimits) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetPeriodStart

`func (o *AiLimits) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *AiLimits) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *AiLimits) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *AiLimits) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetPlan

`func (o *AiLimits) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *AiLimits) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *AiLimits) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *AiLimits) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetSession

`func (o *AiLimits) GetSession() AiWindow`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *AiLimits) GetSessionOk() (*AiWindow, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *AiLimits) SetSession(v AiWindow)`

SetSession sets Session field to given value.

### HasSession

`func (o *AiLimits) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetState

`func (o *AiLimits) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AiLimits) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AiLimits) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AiLimits) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpgrade

`func (o *AiLimits) GetUpgrade() string`

GetUpgrade returns the Upgrade field if non-nil, zero value otherwise.

### GetUpgradeOk

`func (o *AiLimits) GetUpgradeOk() (*string, bool)`

GetUpgradeOk returns a tuple with the Upgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrade

`func (o *AiLimits) SetUpgrade(v string)`

SetUpgrade sets Upgrade field to given value.

### HasUpgrade

`func (o *AiLimits) HasUpgrade() bool`

HasUpgrade returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


