# RiskDatasetSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cuts** | Pointer to **[]string** | Cuts are the two RFC 3339 instants dividing train | val | test. Omit them to take 70% and 85% of the window by time. Splitting is TEMPORAL and then grouped by subject — a random split puts one device on both sides of the line and the model memorises the entity instead of the behaviour. | [optional] 
**Dims** | Pointer to **[]string** | Dims are the coordinates to carry, by published name. Empty takes the whole surface. They are stored in the plane&#39;s own order, never the order given, so two requests naming the same dims produce identical rows. | [optional] 
**From** | Pointer to **string** | From is where the event window opens, RFC 3339, INCLUSIVE. The window may not be longer than the source&#39;s own retention: past that, its older half is already gone and the dataset would silently be shorter than it says. | [optional] 
**Horizon** | Pointer to **int32** | Horizon is how many days a row must have aged before it may be admitted. It is what keeps a fact that was not yet knowable at scoring time out of a training set: a chargeback lands 30 to 120 days after the transaction it condemns, so 120 for the payment lane and 14 for signup abuse. Zero admits the whole window and is honest only where the outcome is immediate. | [optional] 
**Kind** | Pointer to **string** | Kind narrows to one subject kind — person, session or account. Empty takes every kind. | [optional] 
**Name** | Pointer to **string** | Name identifies the dataset across its versions: lower-case letters, digits and hyphens, starting with a letter. | [optional] 
**Rows** | Pointer to **int32** | Rows caps the materialisation. Zero takes the plane&#39;s own bound. | [optional] 
**Seed** | Pointer to **string** | Seed decides WHICH subjects are admitted when the window holds more rows than the cap allows. It is recorded on the version, so a capped dataset is reproducible rather than being whichever rows the store returned first. Omit it to seed from the dataset&#39;s name. | [optional] 
**To** | Pointer to **string** | To is where the window ends, EXCLUSIVE, so two datasets meeting at one instant share no row. A materialisation reads less than this — the end is pulled back by Horizon, and the lineage reports the window it actually read. | [optional] 

## Methods

### NewRiskDatasetSpec

`func NewRiskDatasetSpec() *RiskDatasetSpec`

NewRiskDatasetSpec instantiates a new RiskDatasetSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskDatasetSpecWithDefaults

`func NewRiskDatasetSpecWithDefaults() *RiskDatasetSpec`

NewRiskDatasetSpecWithDefaults instantiates a new RiskDatasetSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCuts

`func (o *RiskDatasetSpec) GetCuts() []string`

GetCuts returns the Cuts field if non-nil, zero value otherwise.

### GetCutsOk

`func (o *RiskDatasetSpec) GetCutsOk() (*[]string, bool)`

GetCutsOk returns a tuple with the Cuts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCuts

`func (o *RiskDatasetSpec) SetCuts(v []string)`

SetCuts sets Cuts field to given value.

### HasCuts

`func (o *RiskDatasetSpec) HasCuts() bool`

HasCuts returns a boolean if a field has been set.

### GetDims

`func (o *RiskDatasetSpec) GetDims() []string`

GetDims returns the Dims field if non-nil, zero value otherwise.

### GetDimsOk

`func (o *RiskDatasetSpec) GetDimsOk() (*[]string, bool)`

GetDimsOk returns a tuple with the Dims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDims

`func (o *RiskDatasetSpec) SetDims(v []string)`

SetDims sets Dims field to given value.

### HasDims

`func (o *RiskDatasetSpec) HasDims() bool`

HasDims returns a boolean if a field has been set.

### GetFrom

`func (o *RiskDatasetSpec) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *RiskDatasetSpec) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *RiskDatasetSpec) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *RiskDatasetSpec) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetHorizon

`func (o *RiskDatasetSpec) GetHorizon() int32`

GetHorizon returns the Horizon field if non-nil, zero value otherwise.

### GetHorizonOk

`func (o *RiskDatasetSpec) GetHorizonOk() (*int32, bool)`

GetHorizonOk returns a tuple with the Horizon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizon

`func (o *RiskDatasetSpec) SetHorizon(v int32)`

SetHorizon sets Horizon field to given value.

### HasHorizon

`func (o *RiskDatasetSpec) HasHorizon() bool`

HasHorizon returns a boolean if a field has been set.

### GetKind

`func (o *RiskDatasetSpec) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RiskDatasetSpec) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RiskDatasetSpec) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RiskDatasetSpec) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *RiskDatasetSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RiskDatasetSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RiskDatasetSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RiskDatasetSpec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRows

`func (o *RiskDatasetSpec) GetRows() int32`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *RiskDatasetSpec) GetRowsOk() (*int32, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *RiskDatasetSpec) SetRows(v int32)`

SetRows sets Rows field to given value.

### HasRows

`func (o *RiskDatasetSpec) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetSeed

`func (o *RiskDatasetSpec) GetSeed() string`

GetSeed returns the Seed field if non-nil, zero value otherwise.

### GetSeedOk

`func (o *RiskDatasetSpec) GetSeedOk() (*string, bool)`

GetSeedOk returns a tuple with the Seed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeed

`func (o *RiskDatasetSpec) SetSeed(v string)`

SetSeed sets Seed field to given value.

### HasSeed

`func (o *RiskDatasetSpec) HasSeed() bool`

HasSeed returns a boolean if a field has been set.

### GetTo

`func (o *RiskDatasetSpec) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *RiskDatasetSpec) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *RiskDatasetSpec) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *RiskDatasetSpec) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


