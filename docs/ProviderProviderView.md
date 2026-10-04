# ProviderProviderView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]ProviderOpView**](ProviderOpView.md) | Actions are what the connector can be asked to do, by a flow step or an agent (POST /v1/provider/{provider}/run when Runnable). | [optional] 
**App** | Pointer to **string** | App names the OAuth app registration its sign-in uses (\&quot;GOOGLE\&quot;), for a connector whose sign-in awaits one. | [optional] 
**Available** | Pointer to **bool** | Available is whether THIS DEPLOYMENT has the provider&#39;s app credentials, so connect can succeed. False renders the card without a working Connect button. | [optional] 
**Category** | Pointer to **string** | Category groups the card (\&quot;Communication\&quot;, \&quot;Developer\&quot;, \&quot;Marketing\&quot;). | [optional] 
**Checks** | Pointer to **bool** | Checks is whether a pasted credential is checked with the provider before it is stored. False means it is sealed on its shape and checked on use. | [optional] 
**Connected** | Pointer to **bool** | Connected is whether this org has a live connection to the provider. | [optional] 
**Connection** | Pointer to [**ProviderConnectionView**](ProviderConnectionView.md) | Connection is the connected account&#39;s non-secret detail. Absent when the org has no connection; tokens NEVER appear here (they live only in KMS). | [optional] 
**Description** | Pointer to **string** | Description is the one-line pitch the console card shows. | [optional] 
**Fields** | Pointer to [**[]ProviderField**](ProviderField.md) | Fields is the form a credential connector asks for. Absent otherwise. | [optional] 
**Id** | Pointer to **string** | ID is the provider&#39;s registry id and the :provider path segment (\&quot;slack\&quot;). | [optional] 
**Kind** | Pointer to **string** | Kind is how a reader connects it NOW: \&quot;oauth\&quot; (send the browser to the provider), \&quot;credential\&quot; (fill in Fields), \&quot;device\&quot; (enter a code), or \&quot;none\&quot; (nothing to connect — ready, or not connectable here; Note says). | [optional] 
**Logo** | Pointer to **string** | Logo is where the connector&#39;s logo is served. Absent when it has none. | [optional] 
**Modes** | Pointer to **[]string** | Modes are the ways this provider can be connected, derived from what it actually declares — \&quot;managed\&quot;, \&quot;oauth\&quot;, \&quot;apikey\&quot;, \&quot;device\&quot;, \&quot;mcp\&quot;. One console page renders every card from this field rather than carrying a table of its own, which is how a provider added here shows up there without a second edit. | [optional] 
**Name** | Pointer to **string** | Name is the provider&#39;s display name (\&quot;Slack\&quot;). | [optional] 
**Note** | Pointer to **string** | Note says why the connector cannot be connected here, when it cannot. | [optional] 
**Runnable** | Pointer to **bool** | Runnable is whether its actions run on this deployment. | [optional] 
**Scopes** | Pointer to **[]string** | Scopes are the permissions a connection will ask for. Never null. | [optional] 
**Triggers** | Pointer to [**[]ProviderOpView**](ProviderOpView.md) | Triggers are the events a flow can start on. | [optional] 

## Methods

### NewProviderProviderView

`func NewProviderProviderView() *ProviderProviderView`

NewProviderProviderView instantiates a new ProviderProviderView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderProviderViewWithDefaults

`func NewProviderProviderViewWithDefaults() *ProviderProviderView`

NewProviderProviderViewWithDefaults instantiates a new ProviderProviderView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ProviderProviderView) GetActions() []ProviderOpView`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ProviderProviderView) GetActionsOk() (*[]ProviderOpView, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ProviderProviderView) SetActions(v []ProviderOpView)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ProviderProviderView) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetApp

`func (o *ProviderProviderView) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *ProviderProviderView) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *ProviderProviderView) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *ProviderProviderView) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetAvailable

`func (o *ProviderProviderView) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *ProviderProviderView) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *ProviderProviderView) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *ProviderProviderView) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCategory

`func (o *ProviderProviderView) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ProviderProviderView) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ProviderProviderView) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ProviderProviderView) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetChecks

`func (o *ProviderProviderView) GetChecks() bool`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *ProviderProviderView) GetChecksOk() (*bool, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *ProviderProviderView) SetChecks(v bool)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *ProviderProviderView) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetConnected

`func (o *ProviderProviderView) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *ProviderProviderView) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *ProviderProviderView) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *ProviderProviderView) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetConnection

`func (o *ProviderProviderView) GetConnection() ProviderConnectionView`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *ProviderProviderView) GetConnectionOk() (*ProviderConnectionView, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *ProviderProviderView) SetConnection(v ProviderConnectionView)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *ProviderProviderView) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetDescription

`func (o *ProviderProviderView) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProviderProviderView) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProviderProviderView) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProviderProviderView) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFields

`func (o *ProviderProviderView) GetFields() []ProviderField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ProviderProviderView) GetFieldsOk() (*[]ProviderField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ProviderProviderView) SetFields(v []ProviderField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ProviderProviderView) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetId

`func (o *ProviderProviderView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProviderProviderView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProviderProviderView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProviderProviderView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *ProviderProviderView) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ProviderProviderView) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ProviderProviderView) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ProviderProviderView) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLogo

`func (o *ProviderProviderView) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *ProviderProviderView) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *ProviderProviderView) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *ProviderProviderView) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetModes

`func (o *ProviderProviderView) GetModes() []string`

GetModes returns the Modes field if non-nil, zero value otherwise.

### GetModesOk

`func (o *ProviderProviderView) GetModesOk() (*[]string, bool)`

GetModesOk returns a tuple with the Modes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModes

`func (o *ProviderProviderView) SetModes(v []string)`

SetModes sets Modes field to given value.

### HasModes

`func (o *ProviderProviderView) HasModes() bool`

HasModes returns a boolean if a field has been set.

### GetName

`func (o *ProviderProviderView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderProviderView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderProviderView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProviderProviderView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNote

`func (o *ProviderProviderView) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ProviderProviderView) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ProviderProviderView) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *ProviderProviderView) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetRunnable

`func (o *ProviderProviderView) GetRunnable() bool`

GetRunnable returns the Runnable field if non-nil, zero value otherwise.

### GetRunnableOk

`func (o *ProviderProviderView) GetRunnableOk() (*bool, bool)`

GetRunnableOk returns a tuple with the Runnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunnable

`func (o *ProviderProviderView) SetRunnable(v bool)`

SetRunnable sets Runnable field to given value.

### HasRunnable

`func (o *ProviderProviderView) HasRunnable() bool`

HasRunnable returns a boolean if a field has been set.

### GetScopes

`func (o *ProviderProviderView) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ProviderProviderView) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ProviderProviderView) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ProviderProviderView) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetTriggers

`func (o *ProviderProviderView) GetTriggers() []ProviderOpView`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *ProviderProviderView) GetTriggersOk() (*[]ProviderOpView, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *ProviderProviderView) SetTriggers(v []ProviderOpView)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *ProviderProviderView) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


