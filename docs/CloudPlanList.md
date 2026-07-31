# CloudPlanList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plans** | Pointer to **[]interface{}** | Plans are the plans in this section, each an opaque object exactly as the @hanzo/plans catalog emits it — typically id, name, description, priceMonthly, category, a feature list, a limits block and a price_ref. | [optional] 

## Methods

### NewCloudPlanList

`func NewCloudPlanList() *CloudPlanList`

NewCloudPlanList instantiates a new CloudPlanList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPlanListWithDefaults

`func NewCloudPlanListWithDefaults() *CloudPlanList`

NewCloudPlanListWithDefaults instantiates a new CloudPlanList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlans

`func (o *CloudPlanList) GetPlans() []interface{}`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *CloudPlanList) GetPlansOk() (*[]interface{}, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *CloudPlanList) SetPlans(v []interface{})`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *CloudPlanList) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


