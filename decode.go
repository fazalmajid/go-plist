package plist

import (
	"io"
	"reflect"
)

type plistValueDecoder interface {
	decodeDocument() (*plistValue, error)
}

type Decoder struct {
	valueDecoder plistValueDecoder
}

func (p *Decoder) Decode(v interface{}) error {
	pval, err := p.valueDecoder.decodeDocument()
	if err != nil {
		return err
	}

	return p.unmarshal(pval, reflect.ValueOf(v))
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		valueDecoder: newXMLPlistValueDecoder(r),
	}
}
