# CloudBlueprintRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EstCentsPerMonth** | Pointer to **int32** | CentsPerMonth is the estimated compute cost of running the whole stack for one month, in USD cents, from the rate card GET /v1/blueprint/health echoes. | [optional] 
**Services** | Pointer to **int32** | Services is how many compose services the stack runs. | [optional] 
**TemplateId** | Pointer to **string** | TemplateID is the blueprint slug — the id GET /v1/blueprint/sbom takes as ?template&#x3D; and the path under templates.hanzo.ai/blueprints/&lt;id&gt;/. | [optional] 

## Methods

### NewCloudBlueprintRow

`func NewCloudBlueprintRow() *CloudBlueprintRow`

NewCloudBlueprintRow instantiates a new CloudBlueprintRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBlueprintRowWithDefaults

`func NewCloudBlueprintRowWithDefaults() *CloudBlueprintRow`

NewCloudBlueprintRowWithDefaults instantiates a new CloudBlueprintRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEstCentsPerMonth

`func (o *CloudBlueprintRow) GetEstCentsPerMonth() int32`

GetEstCentsPerMonth returns the EstCentsPerMonth field if non-nil, zero value otherwise.

### GetEstCentsPerMonthOk

`func (o *CloudBlueprintRow) GetEstCentsPerMonthOk() (*int32, bool)`

GetEstCentsPerMonthOk returns a tuple with the EstCentsPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstCentsPerMonth

`func (o *CloudBlueprintRow) SetEstCentsPerMonth(v int32)`

SetEstCentsPerMonth sets EstCentsPerMonth field to given value.

### HasEstCentsPerMonth

`func (o *CloudBlueprintRow) HasEstCentsPerMonth() bool`

HasEstCentsPerMonth returns a boolean if a field has been set.

### GetServices

`func (o *CloudBlueprintRow) GetServices() int32`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *CloudBlueprintRow) GetServicesOk() (*int32, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *CloudBlueprintRow) SetServices(v int32)`

SetServices sets Services field to given value.

### HasServices

`func (o *CloudBlueprintRow) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetTemplateId

`func (o *CloudBlueprintRow) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *CloudBlueprintRow) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *CloudBlueprintRow) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *CloudBlueprintRow) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


