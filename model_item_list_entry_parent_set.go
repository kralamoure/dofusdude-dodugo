/*
dofusdude

# A project for you - the developer. The all-in-one toolbelt for your next Ankama related project.  ## Client SDKs - [Javascript](https://github.com/dofusdude/dofusdude-js) `npm i dofusdude-js --save` - [Typescript](https://github.com/dofusdude/dofusdude-ts) `npm i dofusdude-ts --save` - [Go](https://github.com/dofusdude/dodugo) `go get -u github.com/dofusdude/dodugo` - [Python](https://github.com/dofusdude/dofusdude-py) `pip install dofusdude` - [PHP](https://github.com/dofusdude/dofusdude-php) - [Java](https://github.com/dofusdude/dofusdude-java) Maven with GitHub packages setup  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.  Your favorite language is missing? Please let me know!  # Main Features - 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 10 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 📨 **Discord Integration** Ankama related RSS and Almanax feeds to post to Discord servers with advanced features like filters or mentions. Use the endpoints as a dev or the official [Web Client](https://discord.dofusdude.com) as a user.  - 🩸 **Dofus 2 Beta** from stable to bleeding edge by replacing /dofus2 with /dofus2beta.  - 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_ including the dropped languages from the Dofus website _de_ and _it_.  - 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing.  - 🕵️ **Complete** actual data from the game including items invisible to the encyclopedia like quest items.  - 🖼️ **HD Images** rendering game assets to high-res images with up to 800x800 px.  ... and much more on the Roadmap on my [Discord](https://discord.gg/3EtHskZD8h). 

API version: 0.9.2
Contact: stelzo@steado.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dodugo

import (
	"encoding/json"
)

// checks if the ItemListEntryParentSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemListEntryParentSet{}

// ItemListEntryParentSet struct for ItemListEntryParentSet
type ItemListEntryParentSet struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewItemListEntryParentSet instantiates a new ItemListEntryParentSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemListEntryParentSet() *ItemListEntryParentSet {
	this := ItemListEntryParentSet{}
	return &this
}

// NewItemListEntryParentSetWithDefaults instantiates a new ItemListEntryParentSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemListEntryParentSetWithDefaults() *ItemListEntryParentSet {
	this := ItemListEntryParentSet{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ItemListEntryParentSet) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemListEntryParentSet) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ItemListEntryParentSet) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ItemListEntryParentSet) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ItemListEntryParentSet) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemListEntryParentSet) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ItemListEntryParentSet) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ItemListEntryParentSet) SetName(v string) {
	o.Name = &v
}

func (o ItemListEntryParentSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ItemListEntryParentSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableItemListEntryParentSet struct {
	value *ItemListEntryParentSet
	isSet bool
}

func (v NullableItemListEntryParentSet) Get() *ItemListEntryParentSet {
	return v.value
}

func (v *NullableItemListEntryParentSet) Set(val *ItemListEntryParentSet) {
	v.value = val
	v.isSet = true
}

func (v NullableItemListEntryParentSet) IsSet() bool {
	return v.isSet
}

func (v *NullableItemListEntryParentSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemListEntryParentSet(val *ItemListEntryParentSet) *NullableItemListEntryParentSet {
	return &NullableItemListEntryParentSet{value: val, isSet: true}
}

func (v NullableItemListEntryParentSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemListEntryParentSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


