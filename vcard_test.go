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
	"io/ioutil"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	testFiles := []string{
		"testdata/rsdoiel.vcf",
		"testdata/john.doe.vcf",
		"testdata/forrest.gump-v2.vcf",
		"testdata/forrest.gump-v3.vcf",
		"testdata/forrest.gump-v4.vcf",
		"testdata/apple.vcf",
		"testdata/alphabet.vcf",
		"testdata/google.vcf",
	}

	fn := []string{
		"Robert Doiel",
		"John Doe",
		"Forrest Gump",
		"Forrest Gump",
		"Forrest Gump",
		"Apple Inc.",
		"Larry Page",
		"Google",
	}

	orgs := [][]string{
		[]string{"Caltech", "Library Services"},
		[]string{"Example Science Institute", ""},
		[]string{"Bubba Gump Shrimp Co."},
		[]string{"Bubba Gump Shrimp Co."},
		[]string{"Bubba Gump Shrimp Co."},
		[]string{"Apple Inc."},
		[]string{"Alphabet Inc"},
		[]string{"Google, Inc"},
	}

	for i, fname := range testFiles {
		src, err := ioutil.ReadFile(fname)
		if err != nil {
			t.Errorf("can't read test data %s, %s", fname, err)
			t.FailNow()
		}
		vcf := NewVCard()
		if err := vcf.Parse(src); err != nil {
			t.Errorf("parse failed for %s, %s", fname, err)
		}
		if strings.Compare(fn[i], vcf.FullName) != 0 {
			t.Errorf("expected %q, got %q", fn[i], vcf.FullName)
		}

		if len(orgs[i]) != len(vcf.Organization) {
			t.Errorf("expected length %d, got %d", len(orgs[i]), len(vcf.Organization))
		} else {
			for j, val := range orgs[i] {
				if strings.Compare(val, vcf.Organization[j]) != 0 {
					t.Errorf("expected %q, got %q", val, vcf.Organization[j])
				}
			}
		}
	}
}
