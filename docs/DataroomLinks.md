# DataroomLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]DataroomLink**](DataroomLink.md) | Links is every non-archived link, newest first. | [optional] 

## Methods

### NewDataroomLinks

`func NewDataroomLinks() *DataroomLinks`

NewDataroomLinks instantiates a new DataroomLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataroomLinksWithDefaults

`func NewDataroomLinksWithDefaults() *DataroomLinks`

NewDataroomLinksWithDefaults instantiates a new DataroomLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *DataroomLinks) GetLinks() []DataroomLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DataroomLinks) GetLinksOk() (*[]DataroomLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DataroomLinks) SetLinks(v []DataroomLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DataroomLinks) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


