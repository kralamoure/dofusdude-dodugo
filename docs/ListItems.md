# ListItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**PagedLinks**](PagedLinks.md) |  | [optional] 
**Items** | Pointer to [**[]ListItem**](ListItem.md) |  | [optional] 

## Methods

### NewListItems

`func NewListItems() *ListItems`

NewListItems instantiates a new ListItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListItemsWithDefaults

`func NewListItemsWithDefaults() *ListItems`

NewListItemsWithDefaults instantiates a new ListItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ListItems) GetLinks() PagedLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListItems) GetLinksOk() (*PagedLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListItems) SetLinks(v PagedLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListItems) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetItems

`func (o *ListItems) GetItems() []ListItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListItems) GetItemsOk() (*[]ListItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListItems) SetItems(v []ListItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListItems) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


