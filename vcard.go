//
// vcard package is a light weight implementation of the VCard v4 spec implemented
// to make it easy to harvest our public directory.
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
//
// Copyright (c) 2017, Caltech
// All rights not granted herein are expressly reserved by Caltech.
//
// Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package vcard

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	//"time"
)

const (
	Version = `v0.0.1`
	CRLF    = "\r\n"
	LF      = "\n"
	CR      = "\r"
)

// VCard is a struct based on VCard V4.0 example
type VCard struct {
	Version      string   `json:"version"`
	Name         []string `json:"name"`
	FullName     string   `json:"full_name"`
	Organization []string `json:"organization,omitempty"`
	Title        string   `json:"title,omitempty"`
	EMail        []string `json:"email,omitempty"`
	Telephone    []string `json:"telephone,omitempty"`
	Source       string   `json:"source"`
	Revision     string   `json:"revision"`
}

// NewVCard creates a new vcard
func NewVCard() *VCard {
	v := new(VCard)
	return v
}

// Parse takes source of a vcard as a byte array and populates the
// struct with what it finds.
func (vcard *VCard) Parse(src []byte) error {
	var (
		err          error
		field, value string
		inVCard      bool
	)
	// Break out text into lines
	lines := bytes.Split(src, []byte(LF))
	for i, line := range lines {
		if bytes.Compare(line, []byte("BEGIN:VCARD")) == 0 {
			if inVCard == true {
				err = fmt.Errorf("line %d, unexpected BEGIN:VCARD", i)
				break
			}
			inVCard = true
			field = ""
			value = ""
		}
		if inVCard == true {
			if bytes.Compare(line, []byte("END:VCARD")) == 0 {
				inVCard = false
				break
			} else {
				// Are we starting a field or are we in another field?
				if bytes.Contains(line, []byte(":")) == true {
					parts := bytes.SplitN(line, []byte(":"), 2)
					field = fmt.Sprintf("%s", parts[0])
					value = fmt.Sprintf("%s", parts[1])
					switch field {
					case "BEGIN":
					case "N":
						vcard.Name = strings.Split(value, ";")
					case "FN":
						vcard.FullName = value
					case "ORG":
						vcard.Organization = strings.Split(value, ";")
					case "TITLE":
						vcard.Title = value
					case "EMAIL":
						vcard.EMail = strings.Split(value, ";")
					case "TEL":
						vcard.Telephone = strings.Split(value, ";")
					case "SOURCE":
						vcard.Source = value
					case "REV":
						vcard.Revision = value
					}
				}
			}
		}
	}
	if inVCard == true && err == nil {
		err = fmt.Errorf("line %d unexpected end of vCard", len(lines))
	}
	return err
}

// String renders a text view of vcard contents (should really render in VCard v4 format...)
func (vcf *VCard) String() string {
	//FIXME: this should render text in VCard v4 format...
	return "vcard.String() Not implemented"
}

// AsJSON takes a vcard and returns it as a JSON doc suitable for using with datatools or dataset
func (vcf *VCard) AsJSON() ([]byte, error) {
	return json.Marshal(vcf)
}
