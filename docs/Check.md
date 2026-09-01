# Check

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **time.Time** |  | [optional] 
**Job** | Pointer to **string** | Job is the job that decided State. A run reports one conclusion for however many jobs it holds, and the jobs are not interchangeable: the pipeline fails at &#x60;gate&#x60; before it builds anything and at &#x60;receipt&#x60; after it has already built, pinned and proved the release live. Both read &#x60;failure&#x60; on the run, and only the first one means nothing shipped. | [optional] 
**Number** | Pointer to **int32** |  | [optional] 
**State** | Pointer to **string** | State is success | failure | running | absent. &#x60;absent&#x60; is not a kind of failure and is kept apart from one: a failing run is a build that ran and said no, while an absent run is Hanzo Git never having constructed a run for the commit at all — a workflow it cannot parse or a reference it cannot resolve. There is no log to open for the second, so a page that draws them the same sends you looking for one that does not exist. | [optional] 
**Tested** | Pointer to **bool** | Tested reports that the run&#39;s tests executed; Verdict reports that the run said anything about tests at all. They are separate because the interesting case is a run that passed while its test step was skipped — a green build that proved nothing — and that is invisible if the two are one flag. | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Verdict** | Pointer to **bool** |  | [optional] 

## Methods

### NewCheck

`func NewCheck() *Check`

NewCheck instantiates a new Check object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckWithDefaults

`func NewCheckWithDefaults() *Check`

NewCheckWithDefaults instantiates a new Check object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *Check) GetAt() time.Time`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *Check) GetAtOk() (*time.Time, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *Check) SetAt(v time.Time)`

SetAt sets At field to given value.

### HasAt

`func (o *Check) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetJob

`func (o *Check) GetJob() string`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *Check) GetJobOk() (*string, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *Check) SetJob(v string)`

SetJob sets Job field to given value.

### HasJob

`func (o *Check) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetNumber

`func (o *Check) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Check) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Check) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Check) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetState

`func (o *Check) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Check) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Check) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Check) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTested

`func (o *Check) GetTested() bool`

GetTested returns the Tested field if non-nil, zero value otherwise.

### GetTestedOk

`func (o *Check) GetTestedOk() (*bool, bool)`

GetTestedOk returns a tuple with the Tested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTested

`func (o *Check) SetTested(v bool)`

SetTested sets Tested field to given value.

### HasTested

`func (o *Check) HasTested() bool`

HasTested returns a boolean if a field has been set.

### GetUrl

`func (o *Check) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Check) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Check) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Check) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVerdict

`func (o *Check) GetVerdict() bool`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *Check) GetVerdictOk() (*bool, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *Check) SetVerdict(v bool)`

SetVerdict sets Verdict field to given value.

### HasVerdict

`func (o *Check) HasVerdict() bool`

HasVerdict returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


