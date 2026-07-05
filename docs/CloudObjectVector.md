# CloudObjectVector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **[]float32** |  | [optional] 
**Dimension** | Pointer to **int64** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**File** | Pointer to **string** |  | [optional] 
**Index** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Score** | Pointer to **float32** |  | [optional] 
**Store** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**TokenCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewCloudObjectVector

`func NewCloudObjectVector() *CloudObjectVector`

NewCloudObjectVector instantiates a new CloudObjectVector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudObjectVectorWithDefaults

`func NewCloudObjectVectorWithDefaults() *CloudObjectVector`

NewCloudObjectVectorWithDefaults instantiates a new CloudObjectVector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *CloudObjectVector) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *CloudObjectVector) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *CloudObjectVector) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *CloudObjectVector) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCurrency

`func (o *CloudObjectVector) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CloudObjectVector) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CloudObjectVector) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CloudObjectVector) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetData

`func (o *CloudObjectVector) GetData() []float32`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudObjectVector) GetDataOk() (*[]float32, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudObjectVector) SetData(v []float32)`

SetData sets Data field to given value.

### HasData

`func (o *CloudObjectVector) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDimension

`func (o *CloudObjectVector) GetDimension() int64`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *CloudObjectVector) GetDimensionOk() (*int64, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *CloudObjectVector) SetDimension(v int64)`

SetDimension sets Dimension field to given value.

### HasDimension

`func (o *CloudObjectVector) HasDimension() bool`

HasDimension returns a boolean if a field has been set.

### GetDisplayName

`func (o *CloudObjectVector) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CloudObjectVector) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CloudObjectVector) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CloudObjectVector) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFile

`func (o *CloudObjectVector) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *CloudObjectVector) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *CloudObjectVector) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *CloudObjectVector) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetIndex

`func (o *CloudObjectVector) GetIndex() int64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *CloudObjectVector) GetIndexOk() (*int64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *CloudObjectVector) SetIndex(v int64)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *CloudObjectVector) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetName

`func (o *CloudObjectVector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudObjectVector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudObjectVector) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudObjectVector) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *CloudObjectVector) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CloudObjectVector) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CloudObjectVector) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CloudObjectVector) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPrice

`func (o *CloudObjectVector) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CloudObjectVector) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CloudObjectVector) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CloudObjectVector) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetProvider

`func (o *CloudObjectVector) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudObjectVector) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudObjectVector) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudObjectVector) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetScore

`func (o *CloudObjectVector) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *CloudObjectVector) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *CloudObjectVector) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *CloudObjectVector) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetStore

`func (o *CloudObjectVector) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *CloudObjectVector) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *CloudObjectVector) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *CloudObjectVector) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetText

`func (o *CloudObjectVector) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CloudObjectVector) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CloudObjectVector) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *CloudObjectVector) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTokenCount

`func (o *CloudObjectVector) GetTokenCount() int64`

GetTokenCount returns the TokenCount field if non-nil, zero value otherwise.

### GetTokenCountOk

`func (o *CloudObjectVector) GetTokenCountOk() (*int64, bool)`

GetTokenCountOk returns a tuple with the TokenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenCount

`func (o *CloudObjectVector) SetTokenCount(v int64)`

SetTokenCount sets TokenCount field to given value.

### HasTokenCount

`func (o *CloudObjectVector) HasTokenCount() bool`

HasTokenCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


