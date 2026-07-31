# CloudCaptableEquityPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoardApprovalDate** | Pointer to **string** | BoardApprovalDate is the ISO date the board approved the plan. | [optional] 
**Comments** | Pointer to **string** | Comments is free-form notes on the plan. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is when the plan was recorded, in unix milliseconds. | [optional] 
**DefaultCancellatonBehavior** | Pointer to **string** | DefaultCancellatonBehavior is what happens to cancelled grants, RETIRE or RETURN_TO_POOL. The key is spelled as the cap-table wire spells it. | [optional] 
**Id** | Pointer to **string** | ID is the equity plan id. | [optional] 
**InitialSharesReserved** | Pointer to **int32** | InitialSharesReserved is how many shares the plan reserves. | [optional] 
**Name** | Pointer to **string** | Name is the plan name, e.g. \&quot;2026 Stock Option Plan\&quot;. | [optional] 
**PlanEffectiveDate** | Pointer to **string** | PlanEffectiveDate is the ISO date the plan takes effect. | [optional] 
**ShareClassId** | Pointer to **string** | ShareClassID is the class the reserved shares come from. | [optional] 

## Methods

### NewCloudCaptableEquityPlan

`func NewCloudCaptableEquityPlan() *CloudCaptableEquityPlan`

NewCloudCaptableEquityPlan instantiates a new CloudCaptableEquityPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCaptableEquityPlanWithDefaults

`func NewCloudCaptableEquityPlanWithDefaults() *CloudCaptableEquityPlan`

NewCloudCaptableEquityPlanWithDefaults instantiates a new CloudCaptableEquityPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoardApprovalDate

`func (o *CloudCaptableEquityPlan) GetBoardApprovalDate() string`

GetBoardApprovalDate returns the BoardApprovalDate field if non-nil, zero value otherwise.

### GetBoardApprovalDateOk

`func (o *CloudCaptableEquityPlan) GetBoardApprovalDateOk() (*string, bool)`

GetBoardApprovalDateOk returns a tuple with the BoardApprovalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardApprovalDate

`func (o *CloudCaptableEquityPlan) SetBoardApprovalDate(v string)`

SetBoardApprovalDate sets BoardApprovalDate field to given value.

### HasBoardApprovalDate

`func (o *CloudCaptableEquityPlan) HasBoardApprovalDate() bool`

HasBoardApprovalDate returns a boolean if a field has been set.

### GetComments

`func (o *CloudCaptableEquityPlan) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *CloudCaptableEquityPlan) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *CloudCaptableEquityPlan) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *CloudCaptableEquityPlan) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudCaptableEquityPlan) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudCaptableEquityPlan) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudCaptableEquityPlan) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudCaptableEquityPlan) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefaultCancellatonBehavior

`func (o *CloudCaptableEquityPlan) GetDefaultCancellatonBehavior() string`

GetDefaultCancellatonBehavior returns the DefaultCancellatonBehavior field if non-nil, zero value otherwise.

### GetDefaultCancellatonBehaviorOk

`func (o *CloudCaptableEquityPlan) GetDefaultCancellatonBehaviorOk() (*string, bool)`

GetDefaultCancellatonBehaviorOk returns a tuple with the DefaultCancellatonBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCancellatonBehavior

`func (o *CloudCaptableEquityPlan) SetDefaultCancellatonBehavior(v string)`

SetDefaultCancellatonBehavior sets DefaultCancellatonBehavior field to given value.

### HasDefaultCancellatonBehavior

`func (o *CloudCaptableEquityPlan) HasDefaultCancellatonBehavior() bool`

HasDefaultCancellatonBehavior returns a boolean if a field has been set.

### GetId

`func (o *CloudCaptableEquityPlan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudCaptableEquityPlan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudCaptableEquityPlan) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudCaptableEquityPlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInitialSharesReserved

`func (o *CloudCaptableEquityPlan) GetInitialSharesReserved() int32`

GetInitialSharesReserved returns the InitialSharesReserved field if non-nil, zero value otherwise.

### GetInitialSharesReservedOk

`func (o *CloudCaptableEquityPlan) GetInitialSharesReservedOk() (*int32, bool)`

GetInitialSharesReservedOk returns a tuple with the InitialSharesReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialSharesReserved

`func (o *CloudCaptableEquityPlan) SetInitialSharesReserved(v int32)`

SetInitialSharesReserved sets InitialSharesReserved field to given value.

### HasInitialSharesReserved

`func (o *CloudCaptableEquityPlan) HasInitialSharesReserved() bool`

HasInitialSharesReserved returns a boolean if a field has been set.

### GetName

`func (o *CloudCaptableEquityPlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudCaptableEquityPlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudCaptableEquityPlan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudCaptableEquityPlan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlanEffectiveDate

`func (o *CloudCaptableEquityPlan) GetPlanEffectiveDate() string`

GetPlanEffectiveDate returns the PlanEffectiveDate field if non-nil, zero value otherwise.

### GetPlanEffectiveDateOk

`func (o *CloudCaptableEquityPlan) GetPlanEffectiveDateOk() (*string, bool)`

GetPlanEffectiveDateOk returns a tuple with the PlanEffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanEffectiveDate

`func (o *CloudCaptableEquityPlan) SetPlanEffectiveDate(v string)`

SetPlanEffectiveDate sets PlanEffectiveDate field to given value.

### HasPlanEffectiveDate

`func (o *CloudCaptableEquityPlan) HasPlanEffectiveDate() bool`

HasPlanEffectiveDate returns a boolean if a field has been set.

### GetShareClassId

`func (o *CloudCaptableEquityPlan) GetShareClassId() string`

GetShareClassId returns the ShareClassId field if non-nil, zero value otherwise.

### GetShareClassIdOk

`func (o *CloudCaptableEquityPlan) GetShareClassIdOk() (*string, bool)`

GetShareClassIdOk returns a tuple with the ShareClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareClassId

`func (o *CloudCaptableEquityPlan) SetShareClassId(v string)`

SetShareClassId sets ShareClassId field to given value.

### HasShareClassId

`func (o *CloudCaptableEquityPlan) HasShareClassId() bool`

HasShareClassId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


