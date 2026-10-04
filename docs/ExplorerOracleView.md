# ExplorerOracleView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Feed** | Pointer to **string** | Feed is the trading pair this feed prices. | [optional] 
**Id** | Pointer to **string** | ID is the feed&#39;s own id, else its trading pair. | [optional] 
**Name** | Pointer to **string** | Name is the feed&#39;s display name — the trading pair, e.g. \&quot;LUX/USD\&quot;. | [optional] 
**Source** | Pointer to **string** | Source is the oracle network the feed originates from; \&quot;O-Chain\&quot; by default. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;active\&quot; for a listed feed. | [optional] 
**UpdatedAt** | Pointer to **string** | UpdatedAt is the feed&#39;s own timestamp, RFC 3339 UTC. | [optional] 
**Value** | Pointer to **string** | Value is the feed&#39;s price, verbatim as the registry carries it. | [optional] 

## Methods

### NewExplorerOracleView

`func NewExplorerOracleView() *ExplorerOracleView`

NewExplorerOracleView instantiates a new ExplorerOracleView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExplorerOracleViewWithDefaults

`func NewExplorerOracleViewWithDefaults() *ExplorerOracleView`

NewExplorerOracleViewWithDefaults instantiates a new ExplorerOracleView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeed

`func (o *ExplorerOracleView) GetFeed() string`

GetFeed returns the Feed field if non-nil, zero value otherwise.

### GetFeedOk

`func (o *ExplorerOracleView) GetFeedOk() (*string, bool)`

GetFeedOk returns a tuple with the Feed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeed

`func (o *ExplorerOracleView) SetFeed(v string)`

SetFeed sets Feed field to given value.

### HasFeed

`func (o *ExplorerOracleView) HasFeed() bool`

HasFeed returns a boolean if a field has been set.

### GetId

`func (o *ExplorerOracleView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExplorerOracleView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExplorerOracleView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExplorerOracleView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ExplorerOracleView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExplorerOracleView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExplorerOracleView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExplorerOracleView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSource

`func (o *ExplorerOracleView) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ExplorerOracleView) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ExplorerOracleView) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ExplorerOracleView) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *ExplorerOracleView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExplorerOracleView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExplorerOracleView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExplorerOracleView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ExplorerOracleView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ExplorerOracleView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ExplorerOracleView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ExplorerOracleView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetValue

`func (o *ExplorerOracleView) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ExplorerOracleView) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ExplorerOracleView) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ExplorerOracleView) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


