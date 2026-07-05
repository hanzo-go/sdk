# SearchFacetSearch200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FacetHits** | Pointer to [**[]SearchFacetSearch200ResponseFacetHitsInner**](SearchFacetSearch200ResponseFacetHitsInner.md) |  | [optional] 
**FacetQuery** | Pointer to **string** |  | [optional] 
**ProcessingTimeMs** | Pointer to **int32** |  | [optional] 

## Methods

### NewSearchFacetSearch200Response

`func NewSearchFacetSearch200Response() *SearchFacetSearch200Response`

NewSearchFacetSearch200Response instantiates a new SearchFacetSearch200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchFacetSearch200ResponseWithDefaults

`func NewSearchFacetSearch200ResponseWithDefaults() *SearchFacetSearch200Response`

NewSearchFacetSearch200ResponseWithDefaults instantiates a new SearchFacetSearch200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFacetHits

`func (o *SearchFacetSearch200Response) GetFacetHits() []SearchFacetSearch200ResponseFacetHitsInner`

GetFacetHits returns the FacetHits field if non-nil, zero value otherwise.

### GetFacetHitsOk

`func (o *SearchFacetSearch200Response) GetFacetHitsOk() (*[]SearchFacetSearch200ResponseFacetHitsInner, bool)`

GetFacetHitsOk returns a tuple with the FacetHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacetHits

`func (o *SearchFacetSearch200Response) SetFacetHits(v []SearchFacetSearch200ResponseFacetHitsInner)`

SetFacetHits sets FacetHits field to given value.

### HasFacetHits

`func (o *SearchFacetSearch200Response) HasFacetHits() bool`

HasFacetHits returns a boolean if a field has been set.

### GetFacetQuery

`func (o *SearchFacetSearch200Response) GetFacetQuery() string`

GetFacetQuery returns the FacetQuery field if non-nil, zero value otherwise.

### GetFacetQueryOk

`func (o *SearchFacetSearch200Response) GetFacetQueryOk() (*string, bool)`

GetFacetQueryOk returns a tuple with the FacetQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacetQuery

`func (o *SearchFacetSearch200Response) SetFacetQuery(v string)`

SetFacetQuery sets FacetQuery field to given value.

### HasFacetQuery

`func (o *SearchFacetSearch200Response) HasFacetQuery() bool`

HasFacetQuery returns a boolean if a field has been set.

### GetProcessingTimeMs

`func (o *SearchFacetSearch200Response) GetProcessingTimeMs() int32`

GetProcessingTimeMs returns the ProcessingTimeMs field if non-nil, zero value otherwise.

### GetProcessingTimeMsOk

`func (o *SearchFacetSearch200Response) GetProcessingTimeMsOk() (*int32, bool)`

GetProcessingTimeMsOk returns a tuple with the ProcessingTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeMs

`func (o *SearchFacetSearch200Response) SetProcessingTimeMs(v int32)`

SetProcessingTimeMs sets ProcessingTimeMs field to given value.

### HasProcessingTimeMs

`func (o *SearchFacetSearch200Response) HasProcessingTimeMs() bool`

HasProcessingTimeMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


