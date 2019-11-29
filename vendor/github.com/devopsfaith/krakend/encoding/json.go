package encoding

import (
	"encoding/json"
	"encoding/xml"
	"io"
)

// JSON is the key for the json encoding
const JSON = "json"

// NewJSONDecoder return the right JSON decoder
func NewJSONDecoder(isCollection bool) func(io.Reader, *map[string]interface{}) error {
	if isCollection {
		return JSONCollectionDecoder
	}
	return JSONDecoder
}

// JSONDecoder implements the Decoder interface
func JSONDecoder(r io.Reader, v *map[string]interface{}) error {
	d := json.NewDecoder(r)
	d.UseNumber()
	return d.Decode(v)
}

// JSONCollectionDecoder implements the Decoder interface over a collection
func JSONCollectionDecoder(r io.Reader, v *map[string]interface{}) error {
	var collection []interface{}
	d := json.NewDecoder(r)
	d.UseNumber()
	if err := d.Decode(&collection); err != nil {
		return err
	}
	*(v) = map[string]interface{}{"collection": collection}
	return nil
}

const XML = "xml"

func NewXMLDecoder(isCollection bool) func(io.Reader, *map[string]interface{}) error {
	if isCollection {
		return XMLCollectionDecoder
	}
	return XMLDecoder
}

func XMLDecoder(r io.Reader, v *map[string]interface{}) error {
	d := xml.NewDecoder(r)
	//d.UseNumber()
	return d.Decode(v)
}

func XMLCollectionDecoder(r io.Reader, v *map[string]interface{}) error {
	var collection []interface{}
	d := xml.NewDecoder(r)
	//d.UseNumber()
	if err := d.Decode(&collection); err != nil {
		return err
	}
	*(v) = map[string]interface{}{"collection": collection}
	return nil
}
