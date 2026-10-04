# BlueprintBlueprintRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EstCentsPerMonth** | Pointer to **int64** | CentsPerMonth is the estimated compute cost of running the whole stack for one month, in USD cents, from the rate card GET /v1/blueprint/health echoes. | [optional] 
**Services** | Pointer to **int64** | Services is how many compose services the stack runs. | [optional] 
**TemplateId** | Pointer to **string** | TemplateID is the blueprint slug — the id GET /v1/blueprint/sbom takes as ?template&#x3D; and the path under templates.hanzo.ai/blueprints/&lt;id&gt;/. | [optional] 

## Methods

### NewBlueprintBlueprintRow

`func NewBlueprintBlueprintRow() *BlueprintBlueprintRow`

NewBlueprintBlueprintRow instantiates a new BlueprintBlueprintRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintBlueprintRowWithDefaults

`func NewBlueprintBlueprintRowWithDefaults() *BlueprintBlueprintRow`

NewBlueprintBlueprintRowWithDefaults instantiates a new BlueprintBlueprintRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEstCentsPerMonth

`func (o *BlueprintBlueprintRow) GetEstCentsPerMonth() int64`

GetEstCentsPerMonth returns the EstCentsPerMonth field if non-nil, zero value otherwise.

### GetEstCentsPerMonthOk

`func (o *BlueprintBlueprintRow) GetEstCentsPerMonthOk() (*int64, bool)`

GetEstCentsPerMonthOk returns a tuple with the EstCentsPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstCentsPerMonth

`func (o *BlueprintBlueprintRow) SetEstCentsPerMonth(v int64)`

SetEstCentsPerMonth sets EstCentsPerMonth field to given value.

### HasEstCentsPerMonth

`func (o *BlueprintBlueprintRow) HasEstCentsPerMonth() bool`

HasEstCentsPerMonth returns a boolean if a field has been set.

### GetServices

`func (o *BlueprintBlueprintRow) GetServices() int64`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *BlueprintBlueprintRow) GetServicesOk() (*int64, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *BlueprintBlueprintRow) SetServices(v int64)`

SetServices sets Services field to given value.

### HasServices

`func (o *BlueprintBlueprintRow) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetTemplateId

`func (o *BlueprintBlueprintRow) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *BlueprintBlueprintRow) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *BlueprintBlueprintRow) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *BlueprintBlueprintRow) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


