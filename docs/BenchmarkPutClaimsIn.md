# BenchmarkPutClaimsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]BenchmarkPublishedClaim**](BenchmarkPublishedClaim.md) | Data is the claims to record. One row is a correction; many is an import. There is no separate bulk endpoint because there is no separate operation: importing a leaderboard and fixing one number are the same write. | [optional] 
**Visibility** | Pointer to **string** | Visibility is who may read every row in Data: \&quot;private\&quot;, the default, keeps them to the caller&#39;s org; \&quot;public\&quot; shows them to anyone. Restating a claim with the other value moves it. | [optional] 

## Methods

### NewBenchmarkPutClaimsIn

`func NewBenchmarkPutClaimsIn() *BenchmarkPutClaimsIn`

NewBenchmarkPutClaimsIn instantiates a new BenchmarkPutClaimsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBenchmarkPutClaimsInWithDefaults

`func NewBenchmarkPutClaimsInWithDefaults() *BenchmarkPutClaimsIn`

NewBenchmarkPutClaimsInWithDefaults instantiates a new BenchmarkPutClaimsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BenchmarkPutClaimsIn) GetData() []BenchmarkPublishedClaim`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BenchmarkPutClaimsIn) GetDataOk() (*[]BenchmarkPublishedClaim, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BenchmarkPutClaimsIn) SetData(v []BenchmarkPublishedClaim)`

SetData sets Data field to given value.

### HasData

`func (o *BenchmarkPutClaimsIn) HasData() bool`

HasData returns a boolean if a field has been set.

### GetVisibility

`func (o *BenchmarkPutClaimsIn) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *BenchmarkPutClaimsIn) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *BenchmarkPutClaimsIn) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *BenchmarkPutClaimsIn) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


