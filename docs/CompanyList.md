# CompanyList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Company**](Company.md) | Data is the page of companies, most recently updated first. | [optional] 

## Methods

### NewCompanyList

`func NewCompanyList() *CompanyList`

NewCompanyList instantiates a new CompanyList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyListWithDefaults

`func NewCompanyListWithDefaults() *CompanyList`

NewCompanyListWithDefaults instantiates a new CompanyList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CompanyList) GetData() []Company`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CompanyList) GetDataOk() (*[]Company, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CompanyList) SetData(v []Company)`

SetData sets Data field to given value.

### HasData

`func (o *CompanyList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


