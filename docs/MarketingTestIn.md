# MarketingTestIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the broadcast, from the path. | [optional] 
**To** | Pointer to **[]string** | To are at most five addresses to send the test to. | [optional] 

## Methods

### NewMarketingTestIn

`func NewMarketingTestIn() *MarketingTestIn`

NewMarketingTestIn instantiates a new MarketingTestIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingTestInWithDefaults

`func NewMarketingTestInWithDefaults() *MarketingTestIn`

NewMarketingTestInWithDefaults instantiates a new MarketingTestIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MarketingTestIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketingTestIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketingTestIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketingTestIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTo

`func (o *MarketingTestIn) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MarketingTestIn) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MarketingTestIn) SetTo(v []string)`

SetTo sets To field to given value.

### HasTo

`func (o *MarketingTestIn) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


