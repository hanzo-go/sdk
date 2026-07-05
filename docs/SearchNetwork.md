# SearchNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to **string** | This instance&#39;s URL | [optional] 
**Remotes** | Pointer to [**map[string]SearchNetworkRemotesValue**](SearchNetworkRemotesValue.md) |  | [optional] 

## Methods

### NewSearchNetwork

`func NewSearchNetwork() *SearchNetwork`

NewSearchNetwork instantiates a new SearchNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchNetworkWithDefaults

`func NewSearchNetworkWithDefaults() *SearchNetwork`

NewSearchNetworkWithDefaults instantiates a new SearchNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *SearchNetwork) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *SearchNetwork) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *SearchNetwork) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *SearchNetwork) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetRemotes

`func (o *SearchNetwork) GetRemotes() map[string]SearchNetworkRemotesValue`

GetRemotes returns the Remotes field if non-nil, zero value otherwise.

### GetRemotesOk

`func (o *SearchNetwork) GetRemotesOk() (*map[string]SearchNetworkRemotesValue, bool)`

GetRemotesOk returns a tuple with the Remotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotes

`func (o *SearchNetwork) SetRemotes(v map[string]SearchNetworkRemotesValue)`

SetRemotes sets Remotes field to given value.

### HasRemotes

`func (o *SearchNetwork) HasRemotes() bool`

HasRemotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


