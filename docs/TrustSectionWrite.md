# TrustSectionWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **interface{}** |  | [optional] 
**Id** | Pointer to **string** | ID is the record&#39;s id. Omit it on a create and one is minted; the single-valued sections (profile, risk) hold one record whatever is named. | [optional] 
**Kind** | Pointer to **string** | Kind is the section being written. The URL is the authority. | [optional] 
**Ord** | Pointer to **int64** | Ord orders this record within its section, ascending, ties broken by id. It is the organization&#39;s own ordering — the page renders in it. | [optional] 

## Methods

### NewTrustSectionWrite

`func NewTrustSectionWrite() *TrustSectionWrite`

NewTrustSectionWrite instantiates a new TrustSectionWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustSectionWriteWithDefaults

`func NewTrustSectionWriteWithDefaults() *TrustSectionWrite`

NewTrustSectionWriteWithDefaults instantiates a new TrustSectionWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TrustSectionWrite) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TrustSectionWrite) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TrustSectionWrite) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *TrustSectionWrite) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *TrustSectionWrite) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *TrustSectionWrite) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetId

`func (o *TrustSectionWrite) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrustSectionWrite) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrustSectionWrite) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrustSectionWrite) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *TrustSectionWrite) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TrustSectionWrite) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TrustSectionWrite) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *TrustSectionWrite) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetOrd

`func (o *TrustSectionWrite) GetOrd() int64`

GetOrd returns the Ord field if non-nil, zero value otherwise.

### GetOrdOk

`func (o *TrustSectionWrite) GetOrdOk() (*int64, bool)`

GetOrdOk returns a tuple with the Ord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrd

`func (o *TrustSectionWrite) SetOrd(v int64)`

SetOrd sets Ord field to given value.

### HasOrd

`func (o *TrustSectionWrite) HasOrd() bool`

HasOrd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


