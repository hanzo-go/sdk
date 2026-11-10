# AutoJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Automation** | Pointer to **string** |  | [optional] 
**Cost** | Pointer to **int64** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Finished** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Origin** | Pointer to **string** |  | [optional] 
**Outputs** | Pointer to **[]string** |  | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**Progress** | Pointer to **float64** |  | [optional] 
**Ref** | Pointer to **string** |  | [optional] 
**Session** | Pointer to **string** |  | [optional] 
**Started** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Updated** | Pointer to **string** |  | [optional] 

## Methods

### NewAutoJob

`func NewAutoJob() *AutoJob`

NewAutoJob instantiates a new AutoJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoJobWithDefaults

`func NewAutoJobWithDefaults() *AutoJob`

NewAutoJobWithDefaults instantiates a new AutoJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomation

`func (o *AutoJob) GetAutomation() string`

GetAutomation returns the Automation field if non-nil, zero value otherwise.

### GetAutomationOk

`func (o *AutoJob) GetAutomationOk() (*string, bool)`

GetAutomationOk returns a tuple with the Automation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomation

`func (o *AutoJob) SetAutomation(v string)`

SetAutomation sets Automation field to given value.

### HasAutomation

`func (o *AutoJob) HasAutomation() bool`

HasAutomation returns a boolean if a field has been set.

### GetCost

`func (o *AutoJob) GetCost() int64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *AutoJob) GetCostOk() (*int64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *AutoJob) SetCost(v int64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *AutoJob) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetError

`func (o *AutoJob) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AutoJob) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AutoJob) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AutoJob) HasError() bool`

HasError returns a boolean if a field has been set.

### GetFinished

`func (o *AutoJob) GetFinished() string`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *AutoJob) GetFinishedOk() (*string, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *AutoJob) SetFinished(v string)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *AutoJob) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### GetId

`func (o *AutoJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoJob) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AutoJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *AutoJob) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AutoJob) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AutoJob) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AutoJob) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetOrigin

`func (o *AutoJob) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *AutoJob) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *AutoJob) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *AutoJob) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetOutputs

`func (o *AutoJob) GetOutputs() []string`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *AutoJob) GetOutputsOk() (*[]string, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *AutoJob) SetOutputs(v []string)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *AutoJob) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetProduct

`func (o *AutoJob) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AutoJob) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AutoJob) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AutoJob) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetProgress

`func (o *AutoJob) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *AutoJob) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *AutoJob) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *AutoJob) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetRef

`func (o *AutoJob) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *AutoJob) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *AutoJob) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *AutoJob) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetSession

`func (o *AutoJob) GetSession() string`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *AutoJob) GetSessionOk() (*string, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *AutoJob) SetSession(v string)`

SetSession sets Session field to given value.

### HasSession

`func (o *AutoJob) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetStarted

`func (o *AutoJob) GetStarted() string`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *AutoJob) GetStartedOk() (*string, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *AutoJob) SetStarted(v string)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *AutoJob) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetState

`func (o *AutoJob) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AutoJob) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AutoJob) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AutoJob) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatus

`func (o *AutoJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AutoJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AutoJob) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AutoJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSummary

`func (o *AutoJob) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AutoJob) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AutoJob) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AutoJob) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *AutoJob) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AutoJob) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AutoJob) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AutoJob) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdated

`func (o *AutoJob) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AutoJob) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *AutoJob) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *AutoJob) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


