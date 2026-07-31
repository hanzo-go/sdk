# CloudSearchIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Doctypes** | Pointer to **[]string** | DocTypes restricts retrieval to a subset of the indexed knowledge doctypes (kb-page, kb-memory, kb-source). An empty or foreign list reads all of them. | [optional] 
**Limit** | Pointer to **int32** | Limit bounds the hits returned. Default 10, maximum 50. | [optional] 
**Project** | Pointer to **string** | Project narrows retrieval to one project scope. | [optional] 
**Query** | Pointer to **string** | Query is the natural-language question. Required. | [optional] 

## Methods

### NewCloudSearchIn

`func NewCloudSearchIn() *CloudSearchIn`

NewCloudSearchIn instantiates a new CloudSearchIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSearchInWithDefaults

`func NewCloudSearchInWithDefaults() *CloudSearchIn`

NewCloudSearchInWithDefaults instantiates a new CloudSearchIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDoctypes

`func (o *CloudSearchIn) GetDoctypes() []string`

GetDoctypes returns the Doctypes field if non-nil, zero value otherwise.

### GetDoctypesOk

`func (o *CloudSearchIn) GetDoctypesOk() (*[]string, bool)`

GetDoctypesOk returns a tuple with the Doctypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoctypes

`func (o *CloudSearchIn) SetDoctypes(v []string)`

SetDoctypes sets Doctypes field to given value.

### HasDoctypes

`func (o *CloudSearchIn) HasDoctypes() bool`

HasDoctypes returns a boolean if a field has been set.

### GetLimit

`func (o *CloudSearchIn) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CloudSearchIn) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CloudSearchIn) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CloudSearchIn) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetProject

`func (o *CloudSearchIn) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CloudSearchIn) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CloudSearchIn) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *CloudSearchIn) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetQuery

`func (o *CloudSearchIn) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CloudSearchIn) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CloudSearchIn) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *CloudSearchIn) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


