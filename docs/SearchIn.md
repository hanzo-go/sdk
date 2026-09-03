# SearchIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Doctypes** | Pointer to **[]string** | DocTypes restricts retrieval to a subset of the indexed knowledge doctypes (kb.page, kb.memory, kb.source). An empty or foreign list reads all of them. | [optional] 
**Limit** | Pointer to **int64** | Limit bounds the hits returned. Default 10, maximum 50. | [optional] 
**Project** | Pointer to **string** | Project narrows retrieval to one project scope. | [optional] 
**Query** | Pointer to **string** | Query is the natural-language question. Required. | [optional] 

## Methods

### NewSearchIn

`func NewSearchIn() *SearchIn`

NewSearchIn instantiates a new SearchIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchInWithDefaults

`func NewSearchInWithDefaults() *SearchIn`

NewSearchInWithDefaults instantiates a new SearchIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDoctypes

`func (o *SearchIn) GetDoctypes() []string`

GetDoctypes returns the Doctypes field if non-nil, zero value otherwise.

### GetDoctypesOk

`func (o *SearchIn) GetDoctypesOk() (*[]string, bool)`

GetDoctypesOk returns a tuple with the Doctypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoctypes

`func (o *SearchIn) SetDoctypes(v []string)`

SetDoctypes sets Doctypes field to given value.

### HasDoctypes

`func (o *SearchIn) HasDoctypes() bool`

HasDoctypes returns a boolean if a field has been set.

### GetLimit

`func (o *SearchIn) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SearchIn) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SearchIn) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SearchIn) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetProject

`func (o *SearchIn) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *SearchIn) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *SearchIn) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *SearchIn) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetQuery

`func (o *SearchIn) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SearchIn) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SearchIn) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *SearchIn) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


