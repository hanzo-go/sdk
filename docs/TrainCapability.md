# TrainCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**By** | Pointer to **string** |  | [optional] 
**Corrections** | Pointer to **[]string** | Corrections are the corrections it is taught from. | [optional] 
**Created** | Pointer to **int64** |  | [optional] 
**Installed** | Pointer to **int64** | Installed is the version in the org&#39;s published set, 0 when none. | [optional] 
**Live** | Pointer to [**TrainLive**](TrainLive.md) | Live is what its installed version did since it went in. | [optional] 
**Name** | Pointer to **string** | Name is its name, one per org; it is the name its versions install under. | [optional] 
**Question** | Pointer to [**TrainQuestion**](TrainQuestion.md) | Question is what it answers: the route it is installed behind. | [optional] 
**State** | Pointer to **string** | State is where it is in its life: draft, trained, evaluated, publishable, installed, active, disabled or needs_validation. | [optional] 
**Updated** | Pointer to **int64** |  | [optional] 
**Version** | Pointer to **int64** | Version is its latest version, 0 in draft. | [optional] 
**Versions** | Pointer to [**[]TrainVersion**](TrainVersion.md) | Versions are its versions, newest first (one capability read). | [optional] 

## Methods

### NewTrainCapability

`func NewTrainCapability() *TrainCapability`

NewTrainCapability instantiates a new TrainCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainCapabilityWithDefaults

`func NewTrainCapabilityWithDefaults() *TrainCapability`

NewTrainCapabilityWithDefaults instantiates a new TrainCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBy

`func (o *TrainCapability) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *TrainCapability) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *TrainCapability) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *TrainCapability) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetCorrections

`func (o *TrainCapability) GetCorrections() []string`

GetCorrections returns the Corrections field if non-nil, zero value otherwise.

### GetCorrectionsOk

`func (o *TrainCapability) GetCorrectionsOk() (*[]string, bool)`

GetCorrectionsOk returns a tuple with the Corrections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrections

`func (o *TrainCapability) SetCorrections(v []string)`

SetCorrections sets Corrections field to given value.

### HasCorrections

`func (o *TrainCapability) HasCorrections() bool`

HasCorrections returns a boolean if a field has been set.

### GetCreated

`func (o *TrainCapability) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TrainCapability) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TrainCapability) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *TrainCapability) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetInstalled

`func (o *TrainCapability) GetInstalled() int64`

GetInstalled returns the Installed field if non-nil, zero value otherwise.

### GetInstalledOk

`func (o *TrainCapability) GetInstalledOk() (*int64, bool)`

GetInstalledOk returns a tuple with the Installed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalled

`func (o *TrainCapability) SetInstalled(v int64)`

SetInstalled sets Installed field to given value.

### HasInstalled

`func (o *TrainCapability) HasInstalled() bool`

HasInstalled returns a boolean if a field has been set.

### GetLive

`func (o *TrainCapability) GetLive() TrainLive`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *TrainCapability) GetLiveOk() (*TrainLive, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *TrainCapability) SetLive(v TrainLive)`

SetLive sets Live field to given value.

### HasLive

`func (o *TrainCapability) HasLive() bool`

HasLive returns a boolean if a field has been set.

### GetName

`func (o *TrainCapability) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrainCapability) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrainCapability) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TrainCapability) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuestion

`func (o *TrainCapability) GetQuestion() TrainQuestion`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *TrainCapability) GetQuestionOk() (*TrainQuestion, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *TrainCapability) SetQuestion(v TrainQuestion)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *TrainCapability) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetState

`func (o *TrainCapability) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TrainCapability) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TrainCapability) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TrainCapability) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdated

`func (o *TrainCapability) GetUpdated() int64`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *TrainCapability) GetUpdatedOk() (*int64, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *TrainCapability) SetUpdated(v int64)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *TrainCapability) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetVersion

`func (o *TrainCapability) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TrainCapability) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TrainCapability) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TrainCapability) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersions

`func (o *TrainCapability) GetVersions() []TrainVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *TrainCapability) GetVersionsOk() (*[]TrainVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *TrainCapability) SetVersions(v []TrainVersion)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *TrainCapability) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


