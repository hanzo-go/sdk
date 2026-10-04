# MarketplaceSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to **string** | App is the app. | [optional] 
**Status** | Pointer to **string** | Status is ok, absent or failed. | [optional] 

## Methods

### NewMarketplaceSource

`func NewMarketplaceSource() *MarketplaceSource`

NewMarketplaceSource instantiates a new MarketplaceSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceSourceWithDefaults

`func NewMarketplaceSourceWithDefaults() *MarketplaceSource`

NewMarketplaceSourceWithDefaults instantiates a new MarketplaceSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *MarketplaceSource) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *MarketplaceSource) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *MarketplaceSource) SetApp(v string)`

SetApp sets App field to given value.

### HasApp

`func (o *MarketplaceSource) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetStatus

`func (o *MarketplaceSource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MarketplaceSource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MarketplaceSource) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MarketplaceSource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


