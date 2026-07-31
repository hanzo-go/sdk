# GraphOracle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Trading pair (e.g. LUX/USD) | [optional] 
**Feed** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** | Feed price verbatim | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGraphOracle

`func NewGraphOracle() *GraphOracle`

NewGraphOracle instantiates a new GraphOracle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphOracleWithDefaults

`func NewGraphOracleWithDefaults() *GraphOracle`

NewGraphOracleWithDefaults instantiates a new GraphOracle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GraphOracle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GraphOracle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GraphOracle) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GraphOracle) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GraphOracle) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GraphOracle) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GraphOracle) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GraphOracle) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFeed

`func (o *GraphOracle) GetFeed() string`

GetFeed returns the Feed field if non-nil, zero value otherwise.

### GetFeedOk

`func (o *GraphOracle) GetFeedOk() (*string, bool)`

GetFeedOk returns a tuple with the Feed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeed

`func (o *GraphOracle) SetFeed(v string)`

SetFeed sets Feed field to given value.

### HasFeed

`func (o *GraphOracle) HasFeed() bool`

HasFeed returns a boolean if a field has been set.

### GetValue

`func (o *GraphOracle) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GraphOracle) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GraphOracle) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *GraphOracle) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetSource

`func (o *GraphOracle) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GraphOracle) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GraphOracle) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GraphOracle) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *GraphOracle) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GraphOracle) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GraphOracle) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GraphOracle) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GraphOracle) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GraphOracle) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GraphOracle) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GraphOracle) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


