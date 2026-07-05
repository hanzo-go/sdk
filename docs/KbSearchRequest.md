# KbSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | Natural-language question | 
**Limit** | Pointer to **int32** |  | [optional] [default to 10]
**Project** | Pointer to **string** | Optional project scope | [optional] 
**Doctypes** | Pointer to **[]string** | Optional restriction to a subset of indexed knowledge doctypes | [optional] 

## Methods

### NewKbSearchRequest

`func NewKbSearchRequest(query string, ) *KbSearchRequest`

NewKbSearchRequest instantiates a new KbSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKbSearchRequestWithDefaults

`func NewKbSearchRequestWithDefaults() *KbSearchRequest`

NewKbSearchRequestWithDefaults instantiates a new KbSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *KbSearchRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *KbSearchRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *KbSearchRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetLimit

`func (o *KbSearchRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *KbSearchRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *KbSearchRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *KbSearchRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetProject

`func (o *KbSearchRequest) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *KbSearchRequest) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *KbSearchRequest) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *KbSearchRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetDoctypes

`func (o *KbSearchRequest) GetDoctypes() []string`

GetDoctypes returns the Doctypes field if non-nil, zero value otherwise.

### GetDoctypesOk

`func (o *KbSearchRequest) GetDoctypesOk() (*[]string, bool)`

GetDoctypesOk returns a tuple with the Doctypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoctypes

`func (o *KbSearchRequest) SetDoctypes(v []string)`

SetDoctypes sets Doctypes field to given value.

### HasDoctypes

`func (o *KbSearchRequest) HasDoctypes() bool`

HasDoctypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


