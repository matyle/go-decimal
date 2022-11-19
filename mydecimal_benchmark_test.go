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

var bench1Str = "135.902348"
var bench1Bytes = []byte(bench1Str)
var bench2String = "121.234234"
var bench2Bytes = []byte(bench2String)

func ShopspringDecimalCase1() decimal.Decimal {
	dec1, err := decimal.NewFromString(bench1Str)
	if err != nil {
		log.Fatal(err)
	}
	dec2 := decimal.NewFromFloat(100)
	dec3 := dec1.Mul(dec2)
	return dec3
}

func TidbDecimalCase1() *Decimal {
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

func ShopspringDecimalCase2() decimal.Decimal {
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

func TidbDecimalCase2() *Decimal {
	dec1 := new(Decimal)
	dec1.FromString(bench2Bytes)
	dec2 := new(Decimal)
	dec2.FromFloat64(0.1)
	for i := 0; i < 10; i++ {
		DecimalAdd(dec1, dec2, dec1)
	}
	return dec1

}

// case1: 135.90*100 ====
func BenchmarkShopspringDecimalCase1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShopspringDecimalCase1()
	}
	// b.Log(ShopspringDecimalCase1()) // output: 13590
}

func BenchmarkTidbDecimalCase1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TidbDecimalCase1()
	}
	// b.Log(TidbDecimalCase1()) //  output: 13590.00
}

// case2: 0.1 add 10 times ===
func BenchmarkShopspringDecimalCase2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShopspringDecimalCase2()
	}
	// b.Log(ShopspringDecimalCase2()) //  output: 1
}

func BenchmarkTidbDecimalCase2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TidbDecimalCase2()
	}
	// b.Log(TidbDecimalCase2()) //  output: 1.0
}
