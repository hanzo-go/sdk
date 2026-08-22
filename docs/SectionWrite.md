# SectionWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **interface{}** |  | [optional] 
**Id** | Pointer to **string** | ID is the record&#39;s id. Omit it on a create and one is minted; the single-valued sections (profile, risk) hold one record whatever is named. | [optional] 
**Kind** | Pointer to **string** | Kind is the section being written. The URL is the authority. | [optional] 
**Ord** | Pointer to **int32** | Ord orders this record within its section, ascending, ties broken by id. It is the organization&#39;s own ordering — the page renders in it. | [optional] 

## Methods

### NewSectionWrite

`func NewSectionWrite() *SectionWrite`

NewSectionWrite instantiates a new SectionWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectionWriteWithDefaults

`func NewSectionWriteWithDefaults() *SectionWrite`

NewSectionWriteWithDefaults instantiates a new SectionWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SectionWrite) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SectionWrite) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SectionWrite) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *SectionWrite) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *SectionWrite) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *SectionWrite) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetId

`func (o *SectionWrite) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SectionWrite) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SectionWrite) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SectionWrite) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *SectionWrite) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SectionWrite) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SectionWrite) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SectionWrite) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetOrd

`func (o *SectionWrite) GetOrd() int32`

GetOrd returns the Ord field if non-nil, zero value otherwise.

### GetOrdOk

`func (o *SectionWrite) GetOrdOk() (*int32, bool)`

GetOrdOk returns a tuple with the Ord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrd

`func (o *SectionWrite) SetOrd(v int32)`

SetOrd sets Ord field to given value.

### HasOrd

`func (o *SectionWrite) HasOrd() bool`

HasOrd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


