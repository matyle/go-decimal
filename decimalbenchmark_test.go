// Copyright 2019 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package decimal

import (
	"log"
	"testing"

	"github.com/shopspring/decimal"
)

var bench1Str = "135.348"
var bench1Bytes = []byte(bench1Str)
var bench2String = "121.34"
var bench2Bytes = []byte(bench2String)

func shopspringDecimalCase() decimal.Decimal {
	dec1, err := decimal.NewFromString(bench1Str)
	if err != nil {
		log.Fatal(err)
	}
	dec2 := decimal.NewFromFloat(100)
	dec3 := dec1.Mul(dec2)
	return dec3
}

func goDecimalCase() *Decimal {
	dec1 := new(Decimal)
	err := dec1.FromString(bench1Bytes)
	if err != nil {
		log.Fatal(err)
	}
	dec2 := new(Decimal).FromInt(100)
	dec3 := new(Decimal)
	err = DecimalMul(dec1, dec2, dec3)
	if err != nil {
		log.Fatal(err)
	}
	return dec3
}

func shopspringDecimalCaseB() decimal.Decimal {
	dec1, err := decimal.NewFromString(bench2String)
	if err != nil {
		log.Fatal(err)
	}
	dec2 := decimal.NewFromFloat(0.1)
	for i := 0; i < 10; i++ {
		dec1 = dec1.Add(dec2)
	}
	return dec1
}

func goDecimalCaseB() *Decimal {
	dec1 := new(Decimal)
	dec1.FromString(bench2Bytes)
	dec2 := new(Decimal)
	dec2.FromFloat64(0.1)
	for i := 0; i < 10; i++ {
		DecimalAdd(dec1, dec2, dec1)
	}
	return dec1

}

func BenchmarkShopspringDecimalCase1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		shopspringDecimalCase()
	}
}

func BenchmarkGoDecimalCase1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goDecimalCase()
	}
}

func BenchmarkShopspringDecimalCase2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		shopspringDecimalCaseB()
	}
}

func BenchmarkGoDecimalCase2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goDecimalCaseB()
	}
}
