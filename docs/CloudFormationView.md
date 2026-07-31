# CloudFormationView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Formation** | Pointer to [**CloudFormation**](CloudFormation.md) | Formation is the org&#39;s one incorporation record. | [optional] 
**NextStages** | Pointer to **[]string** | NextStages are the stages reachable from the formation&#39;s current stage, whether or not their guards are satisfied yet. | [optional] 

## Methods

### NewCloudFormationView

`func NewCloudFormationView() *CloudFormationView`

NewCloudFormationView instantiates a new CloudFormationView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudFormationViewWithDefaults

`func NewCloudFormationViewWithDefaults() *CloudFormationView`

NewCloudFormationViewWithDefaults instantiates a new CloudFormationView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormation

`func (o *CloudFormationView) GetFormation() CloudFormation`

GetFormation returns the Formation field if non-nil, zero value otherwise.

### GetFormationOk

`func (o *CloudFormationView) GetFormationOk() (*CloudFormation, bool)`

GetFormationOk returns a tuple with the Formation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormation

`func (o *CloudFormationView) SetFormation(v CloudFormation)`

SetFormation sets Formation field to given value.

### HasFormation

`func (o *CloudFormationView) HasFormation() bool`

HasFormation returns a boolean if a field has been set.

### GetNextStages

`func (o *CloudFormationView) GetNextStages() []string`

GetNextStages returns the NextStages field if non-nil, zero value otherwise.

### GetNextStagesOk

`func (o *CloudFormationView) GetNextStagesOk() (*[]string, bool)`

GetNextStagesOk returns a tuple with the NextStages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextStages

`func (o *CloudFormationView) SetNextStages(v []string)`

SetNextStages sets NextStages field to given value.

### HasNextStages

`func (o *CloudFormationView) HasNextStages() bool`

HasNextStages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


