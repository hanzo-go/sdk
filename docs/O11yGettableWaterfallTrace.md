# O11yGettableWaterfallTrace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTimestampMillis** | Pointer to **int32** |  | [optional] 
**HasMissingSpans** | Pointer to **bool** |  | [optional] 
**HasMore** | Pointer to **bool** |  | [optional] 
**RootServiceEntryPoint** | Pointer to **string** |  | [optional] 
**RootServiceName** | Pointer to **string** |  | [optional] 
**Spans** | Pointer to [**[]O11yWaterfallSpan**](O11yWaterfallSpan.md) |  | [optional] 
**StartTimestampMillis** | Pointer to **int32** |  | [optional] 
**TotalErrorSpansCount** | Pointer to **int32** |  | [optional] 
**TotalSpansCount** | Pointer to **int32** |  | [optional] 
**UncollapsedSpans** | Pointer to **[]string** |  | [optional] 

## Methods

### NewO11yGettableWaterfallTrace

`func NewO11yGettableWaterfallTrace() *O11yGettableWaterfallTrace`

NewO11yGettableWaterfallTrace instantiates a new O11yGettableWaterfallTrace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yGettableWaterfallTraceWithDefaults

`func NewO11yGettableWaterfallTraceWithDefaults() *O11yGettableWaterfallTrace`

NewO11yGettableWaterfallTraceWithDefaults instantiates a new O11yGettableWaterfallTrace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTimestampMillis

`func (o *O11yGettableWaterfallTrace) GetEndTimestampMillis() int32`

GetEndTimestampMillis returns the EndTimestampMillis field if non-nil, zero value otherwise.

### GetEndTimestampMillisOk

`func (o *O11yGettableWaterfallTrace) GetEndTimestampMillisOk() (*int32, bool)`

GetEndTimestampMillisOk returns a tuple with the EndTimestampMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimestampMillis

`func (o *O11yGettableWaterfallTrace) SetEndTimestampMillis(v int32)`

SetEndTimestampMillis sets EndTimestampMillis field to given value.

### HasEndTimestampMillis

`func (o *O11yGettableWaterfallTrace) HasEndTimestampMillis() bool`

HasEndTimestampMillis returns a boolean if a field has been set.

### GetHasMissingSpans

`func (o *O11yGettableWaterfallTrace) GetHasMissingSpans() bool`

GetHasMissingSpans returns the HasMissingSpans field if non-nil, zero value otherwise.

### GetHasMissingSpansOk

`func (o *O11yGettableWaterfallTrace) GetHasMissingSpansOk() (*bool, bool)`

GetHasMissingSpansOk returns a tuple with the HasMissingSpans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMissingSpans

`func (o *O11yGettableWaterfallTrace) SetHasMissingSpans(v bool)`

SetHasMissingSpans sets HasMissingSpans field to given value.

### HasHasMissingSpans

`func (o *O11yGettableWaterfallTrace) HasHasMissingSpans() bool`

HasHasMissingSpans returns a boolean if a field has been set.

### GetHasMore

`func (o *O11yGettableWaterfallTrace) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *O11yGettableWaterfallTrace) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *O11yGettableWaterfallTrace) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *O11yGettableWaterfallTrace) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetRootServiceEntryPoint

`func (o *O11yGettableWaterfallTrace) GetRootServiceEntryPoint() string`

GetRootServiceEntryPoint returns the RootServiceEntryPoint field if non-nil, zero value otherwise.

### GetRootServiceEntryPointOk

`func (o *O11yGettableWaterfallTrace) GetRootServiceEntryPointOk() (*string, bool)`

GetRootServiceEntryPointOk returns a tuple with the RootServiceEntryPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootServiceEntryPoint

`func (o *O11yGettableWaterfallTrace) SetRootServiceEntryPoint(v string)`

SetRootServiceEntryPoint sets RootServiceEntryPoint field to given value.

### HasRootServiceEntryPoint

`func (o *O11yGettableWaterfallTrace) HasRootServiceEntryPoint() bool`

HasRootServiceEntryPoint returns a boolean if a field has been set.

### GetRootServiceName

`func (o *O11yGettableWaterfallTrace) GetRootServiceName() string`

GetRootServiceName returns the RootServiceName field if non-nil, zero value otherwise.

### GetRootServiceNameOk

`func (o *O11yGettableWaterfallTrace) GetRootServiceNameOk() (*string, bool)`

GetRootServiceNameOk returns a tuple with the RootServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootServiceName

`func (o *O11yGettableWaterfallTrace) SetRootServiceName(v string)`

SetRootServiceName sets RootServiceName field to given value.

### HasRootServiceName

`func (o *O11yGettableWaterfallTrace) HasRootServiceName() bool`

HasRootServiceName returns a boolean if a field has been set.

### GetSpans

`func (o *O11yGettableWaterfallTrace) GetSpans() []O11yWaterfallSpan`

GetSpans returns the Spans field if non-nil, zero value otherwise.

### GetSpansOk

`func (o *O11yGettableWaterfallTrace) GetSpansOk() (*[]O11yWaterfallSpan, bool)`

GetSpansOk returns a tuple with the Spans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpans

`func (o *O11yGettableWaterfallTrace) SetSpans(v []O11yWaterfallSpan)`

SetSpans sets Spans field to given value.

### HasSpans

`func (o *O11yGettableWaterfallTrace) HasSpans() bool`

HasSpans returns a boolean if a field has been set.

### GetStartTimestampMillis

`func (o *O11yGettableWaterfallTrace) GetStartTimestampMillis() int32`

GetStartTimestampMillis returns the StartTimestampMillis field if non-nil, zero value otherwise.

### GetStartTimestampMillisOk

`func (o *O11yGettableWaterfallTrace) GetStartTimestampMillisOk() (*int32, bool)`

GetStartTimestampMillisOk returns a tuple with the StartTimestampMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestampMillis

`func (o *O11yGettableWaterfallTrace) SetStartTimestampMillis(v int32)`

SetStartTimestampMillis sets StartTimestampMillis field to given value.

### HasStartTimestampMillis

`func (o *O11yGettableWaterfallTrace) HasStartTimestampMillis() bool`

HasStartTimestampMillis returns a boolean if a field has been set.

### GetTotalErrorSpansCount

`func (o *O11yGettableWaterfallTrace) GetTotalErrorSpansCount() int32`

GetTotalErrorSpansCount returns the TotalErrorSpansCount field if non-nil, zero value otherwise.

### GetTotalErrorSpansCountOk

`func (o *O11yGettableWaterfallTrace) GetTotalErrorSpansCountOk() (*int32, bool)`

GetTotalErrorSpansCountOk returns a tuple with the TotalErrorSpansCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalErrorSpansCount

`func (o *O11yGettableWaterfallTrace) SetTotalErrorSpansCount(v int32)`

SetTotalErrorSpansCount sets TotalErrorSpansCount field to given value.

### HasTotalErrorSpansCount

`func (o *O11yGettableWaterfallTrace) HasTotalErrorSpansCount() bool`

HasTotalErrorSpansCount returns a boolean if a field has been set.

### GetTotalSpansCount

`func (o *O11yGettableWaterfallTrace) GetTotalSpansCount() int32`

GetTotalSpansCount returns the TotalSpansCount field if non-nil, zero value otherwise.

### GetTotalSpansCountOk

`func (o *O11yGettableWaterfallTrace) GetTotalSpansCountOk() (*int32, bool)`

GetTotalSpansCountOk returns a tuple with the TotalSpansCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSpansCount

`func (o *O11yGettableWaterfallTrace) SetTotalSpansCount(v int32)`

SetTotalSpansCount sets TotalSpansCount field to given value.

### HasTotalSpansCount

`func (o *O11yGettableWaterfallTrace) HasTotalSpansCount() bool`

HasTotalSpansCount returns a boolean if a field has been set.

### GetUncollapsedSpans

`func (o *O11yGettableWaterfallTrace) GetUncollapsedSpans() []string`

GetUncollapsedSpans returns the UncollapsedSpans field if non-nil, zero value otherwise.

### GetUncollapsedSpansOk

`func (o *O11yGettableWaterfallTrace) GetUncollapsedSpansOk() (*[]string, bool)`

GetUncollapsedSpansOk returns a tuple with the UncollapsedSpans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUncollapsedSpans

`func (o *O11yGettableWaterfallTrace) SetUncollapsedSpans(v []string)`

SetUncollapsedSpans sets UncollapsedSpans field to given value.

### HasUncollapsedSpans

`func (o *O11yGettableWaterfallTrace) HasUncollapsedSpans() bool`

HasUncollapsedSpans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


