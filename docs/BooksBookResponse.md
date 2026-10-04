# BooksBookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Posted** | Pointer to **bool** | Posted is true when this call wrote the voucher, false when the same scan had already booked and nothing was written. | [optional] 
**ScanId** | Pointer to **string** | ScanID echoes the scan that was booked. | [optional] 

## Methods

### NewBooksBookResponse

`func NewBooksBookResponse() *BooksBookResponse`

NewBooksBookResponse instantiates a new BooksBookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBooksBookResponseWithDefaults

`func NewBooksBookResponseWithDefaults() *BooksBookResponse`

NewBooksBookResponseWithDefaults instantiates a new BooksBookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPosted

`func (o *BooksBookResponse) GetPosted() bool`

GetPosted returns the Posted field if non-nil, zero value otherwise.

### GetPostedOk

`func (o *BooksBookResponse) GetPostedOk() (*bool, bool)`

GetPostedOk returns a tuple with the Posted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosted

`func (o *BooksBookResponse) SetPosted(v bool)`

SetPosted sets Posted field to given value.

### HasPosted

`func (o *BooksBookResponse) HasPosted() bool`

HasPosted returns a boolean if a field has been set.

### GetScanId

`func (o *BooksBookResponse) GetScanId() string`

GetScanId returns the ScanId field if non-nil, zero value otherwise.

### GetScanIdOk

`func (o *BooksBookResponse) GetScanIdOk() (*string, bool)`

GetScanIdOk returns a tuple with the ScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanId

`func (o *BooksBookResponse) SetScanId(v string)`

SetScanId sets ScanId field to given value.

### HasScanId

`func (o *BooksBookResponse) HasScanId() bool`

HasScanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


