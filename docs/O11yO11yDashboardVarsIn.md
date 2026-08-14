# O11yO11yDashboardVarsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | Query is the variable query to evaluate. Required. | 
**Variables** | Pointer to **map[string]map[string]interface{}** | Variables are the current values of the other dashboard variables, for queries that reference them. | [optional] 

## Methods

### NewO11yO11yDashboardVarsIn

`func NewO11yO11yDashboardVarsIn(query string, ) *O11yO11yDashboardVarsIn`

NewO11yO11yDashboardVarsIn instantiates a new O11yO11yDashboardVarsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yDashboardVarsInWithDefaults

`func NewO11yO11yDashboardVarsInWithDefaults() *O11yO11yDashboardVarsIn`

NewO11yO11yDashboardVarsInWithDefaults instantiates a new O11yO11yDashboardVarsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *O11yO11yDashboardVarsIn) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *O11yO11yDashboardVarsIn) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *O11yO11yDashboardVarsIn) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetVariables

`func (o *O11yO11yDashboardVarsIn) GetVariables() map[string]map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *O11yO11yDashboardVarsIn) GetVariablesOk() (*map[string]map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *O11yO11yDashboardVarsIn) SetVariables(v map[string]map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *O11yO11yDashboardVarsIn) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


