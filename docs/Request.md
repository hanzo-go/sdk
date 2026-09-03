# Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Doctypes** | Pointer to **[]string** | DocTypes restricts the semantic leg to a subset of indexed knowledge types. | [optional] 
**Index** | Pointer to **string** | Index names the lexical index to query. Defaults to \&quot;kb\&quot;. | [optional] 
**Limit** | Pointer to **int64** | Limit bounds the FUSED result set (default 10, max 50). | [optional] 
**Mode** | Pointer to **string** | Mode selects the legs: auto (default) | text | semantic | hybrid. | [optional] 
**Offset** | Pointer to **int64** | Offset pages the fused result set. | [optional] 
**Project** | Pointer to **string** | Project narrows to one project scope within the org. | [optional] 
**Query** | Pointer to **string** | Query is the natural-language or keyword query. Required. | [optional] 

## Methods

### NewRequest

`func NewRequest() *Request`

NewRequest instantiates a new Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestWithDefaults

`func NewRequestWithDefaults() *Request`

NewRequestWithDefaults instantiates a new Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDoctypes

`func (o *Request) GetDoctypes() []string`

GetDoctypes returns the Doctypes field if non-nil, zero value otherwise.

### GetDoctypesOk

`func (o *Request) GetDoctypesOk() (*[]string, bool)`

GetDoctypesOk returns a tuple with the Doctypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoctypes

`func (o *Request) SetDoctypes(v []string)`

SetDoctypes sets Doctypes field to given value.

### HasDoctypes

`func (o *Request) HasDoctypes() bool`

HasDoctypes returns a boolean if a field has been set.

### GetIndex

`func (o *Request) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Request) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Request) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *Request) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetLimit

`func (o *Request) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Request) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Request) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Request) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMode

`func (o *Request) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Request) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Request) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *Request) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetOffset

`func (o *Request) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Request) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *Request) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *Request) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetProject

`func (o *Request) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Request) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Request) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *Request) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetQuery

`func (o *Request) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Request) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Request) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *Request) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


