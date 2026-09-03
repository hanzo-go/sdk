# SeoDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | Domain is the competitor. | [optional] 
**Keywords** | Pointer to **int64** | Keywords is how many of the phrases it places for. | [optional] 
**Position** | Pointer to **float64** | Position is its average rank across the phrases. | [optional] 
**Traffic** | Pointer to **float64** | Traffic is the estimated monthly visits those placements earn. | [optional] 
**Visibility** | Pointer to **float64** | Visibility is its share of the possible attention across those phrases. | [optional] 

## Methods

### NewSeoDomain

`func NewSeoDomain() *SeoDomain`

NewSeoDomain instantiates a new SeoDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeoDomainWithDefaults

`func NewSeoDomainWithDefaults() *SeoDomain`

NewSeoDomainWithDefaults instantiates a new SeoDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *SeoDomain) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *SeoDomain) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *SeoDomain) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *SeoDomain) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetKeywords

`func (o *SeoDomain) GetKeywords() int64`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *SeoDomain) GetKeywordsOk() (*int64, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *SeoDomain) SetKeywords(v int64)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *SeoDomain) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetPosition

`func (o *SeoDomain) GetPosition() float64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *SeoDomain) GetPositionOk() (*float64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *SeoDomain) SetPosition(v float64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *SeoDomain) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetTraffic

`func (o *SeoDomain) GetTraffic() float64`

GetTraffic returns the Traffic field if non-nil, zero value otherwise.

### GetTrafficOk

`func (o *SeoDomain) GetTrafficOk() (*float64, bool)`

GetTrafficOk returns a tuple with the Traffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffic

`func (o *SeoDomain) SetTraffic(v float64)`

SetTraffic sets Traffic field to given value.

### HasTraffic

`func (o *SeoDomain) HasTraffic() bool`

HasTraffic returns a boolean if a field has been set.

### GetVisibility

`func (o *SeoDomain) GetVisibility() float64`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *SeoDomain) GetVisibilityOk() (*float64, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *SeoDomain) SetVisibility(v float64)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *SeoDomain) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


