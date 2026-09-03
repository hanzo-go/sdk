# CaptableEquityPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoardApprovalDate** | Pointer to **string** | BoardApprovalDate is the ISO date the board approved the plan. | [optional] 
**Comments** | Pointer to **string** | Comments is free-form notes on the plan. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is when the plan was recorded, in unix milliseconds. | [optional] 
**DefaultCancellatonBehavior** | Pointer to **string** | DefaultCancellatonBehavior is what happens to cancelled grants, RETIRE or RETURN_TO_POOL. The key is spelled as the cap-table wire spells it. | [optional] 
**Id** | Pointer to **string** | ID is the equity plan id. | [optional] 
**InitialSharesReserved** | Pointer to **int64** | InitialSharesReserved is how many shares the plan reserves. | [optional] 
**Name** | Pointer to **string** | Name is the plan name, e.g. \&quot;2026 Stock Option Plan\&quot;. | [optional] 
**PlanEffectiveDate** | Pointer to **string** | PlanEffectiveDate is the ISO date the plan takes effect. | [optional] 
**ShareClassId** | Pointer to **string** | ShareClassID is the class the reserved shares come from. | [optional] 

## Methods

### NewCaptableEquityPlan

`func NewCaptableEquityPlan() *CaptableEquityPlan`

NewCaptableEquityPlan instantiates a new CaptableEquityPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableEquityPlanWithDefaults

`func NewCaptableEquityPlanWithDefaults() *CaptableEquityPlan`

NewCaptableEquityPlanWithDefaults instantiates a new CaptableEquityPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoardApprovalDate

`func (o *CaptableEquityPlan) GetBoardApprovalDate() string`

GetBoardApprovalDate returns the BoardApprovalDate field if non-nil, zero value otherwise.

### GetBoardApprovalDateOk

`func (o *CaptableEquityPlan) GetBoardApprovalDateOk() (*string, bool)`

GetBoardApprovalDateOk returns a tuple with the BoardApprovalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardApprovalDate

`func (o *CaptableEquityPlan) SetBoardApprovalDate(v string)`

SetBoardApprovalDate sets BoardApprovalDate field to given value.

### HasBoardApprovalDate

`func (o *CaptableEquityPlan) HasBoardApprovalDate() bool`

HasBoardApprovalDate returns a boolean if a field has been set.

### GetComments

`func (o *CaptableEquityPlan) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *CaptableEquityPlan) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *CaptableEquityPlan) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *CaptableEquityPlan) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CaptableEquityPlan) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CaptableEquityPlan) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CaptableEquityPlan) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CaptableEquityPlan) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefaultCancellatonBehavior

`func (o *CaptableEquityPlan) GetDefaultCancellatonBehavior() string`

GetDefaultCancellatonBehavior returns the DefaultCancellatonBehavior field if non-nil, zero value otherwise.

### GetDefaultCancellatonBehaviorOk

`func (o *CaptableEquityPlan) GetDefaultCancellatonBehaviorOk() (*string, bool)`

GetDefaultCancellatonBehaviorOk returns a tuple with the DefaultCancellatonBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCancellatonBehavior

`func (o *CaptableEquityPlan) SetDefaultCancellatonBehavior(v string)`

SetDefaultCancellatonBehavior sets DefaultCancellatonBehavior field to given value.

### HasDefaultCancellatonBehavior

`func (o *CaptableEquityPlan) HasDefaultCancellatonBehavior() bool`

HasDefaultCancellatonBehavior returns a boolean if a field has been set.

### GetId

`func (o *CaptableEquityPlan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CaptableEquityPlan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CaptableEquityPlan) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CaptableEquityPlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInitialSharesReserved

`func (o *CaptableEquityPlan) GetInitialSharesReserved() int64`

GetInitialSharesReserved returns the InitialSharesReserved field if non-nil, zero value otherwise.

### GetInitialSharesReservedOk

`func (o *CaptableEquityPlan) GetInitialSharesReservedOk() (*int64, bool)`

GetInitialSharesReservedOk returns a tuple with the InitialSharesReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialSharesReserved

`func (o *CaptableEquityPlan) SetInitialSharesReserved(v int64)`

SetInitialSharesReserved sets InitialSharesReserved field to given value.

### HasInitialSharesReserved

`func (o *CaptableEquityPlan) HasInitialSharesReserved() bool`

HasInitialSharesReserved returns a boolean if a field has been set.

### GetName

`func (o *CaptableEquityPlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CaptableEquityPlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CaptableEquityPlan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CaptableEquityPlan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlanEffectiveDate

`func (o *CaptableEquityPlan) GetPlanEffectiveDate() string`

GetPlanEffectiveDate returns the PlanEffectiveDate field if non-nil, zero value otherwise.

### GetPlanEffectiveDateOk

`func (o *CaptableEquityPlan) GetPlanEffectiveDateOk() (*string, bool)`

GetPlanEffectiveDateOk returns a tuple with the PlanEffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanEffectiveDate

`func (o *CaptableEquityPlan) SetPlanEffectiveDate(v string)`

SetPlanEffectiveDate sets PlanEffectiveDate field to given value.

### HasPlanEffectiveDate

`func (o *CaptableEquityPlan) HasPlanEffectiveDate() bool`

HasPlanEffectiveDate returns a boolean if a field has been set.

### GetShareClassId

`func (o *CaptableEquityPlan) GetShareClassId() string`

GetShareClassId returns the ShareClassId field if non-nil, zero value otherwise.

### GetShareClassIdOk

`func (o *CaptableEquityPlan) GetShareClassIdOk() (*string, bool)`

GetShareClassIdOk returns a tuple with the ShareClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareClassId

`func (o *CaptableEquityPlan) SetShareClassId(v string)`

SetShareClassId sets ShareClassId field to given value.

### HasShareClassId

`func (o *CaptableEquityPlan) HasShareClassId() bool`

HasShareClassId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


