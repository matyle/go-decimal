# go-decimal

## golang high performance decimal

MySQL's decimal uses an array buf of decimal_digit_t (int32) with a length of len to store the digits of the decimal. Each decimal_digit_t can store up to 9 digits. Intg is used to represent the number of digits in the integer part, and frac is used to represent the number of digits in the fractional part. sign represents the sign. The decimal part and the integer part need to be stored separately, and cannot be mixed in a decimal_digit_t. Both parts are aligned to the decimal point. This is because integers and decimals usually need to be calculated separately, so this format can make it easier to combine different decimal_t decimals Aligned with integers, which is convenient for addition and subtraction operations. len is always 9 in MySQL implementation, which represents the upper limit of storage, and the actual effective part of buf is determined by intg and frac together.

### benchmark
```bash
goos: darwin
goarch: amd64
pkg: github.com/matyle/go-decimal
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkShopspringDecimalCase1-8   	 2680896	       454.1 ns/op	     176 B/op	       7 allocs/op
BenchmarkGoDecimalCase1-8           	12947986	        93.14 ns/op	      48 B/op	       1 allocs/op
BenchmarkShopspringDecimalCase2-8   	  259377	      4650 ns/op	    2896 B/op	      95 allocs/op
BenchmarkGoDecimalCase2-8           	 1689056	       718.0 ns/op	     555 B/op	      13 allocs/op
PASS
ok  	github.com/matyle/go-decimal	6.774s
```

### Thanks 
thanks  @ingcap/tidb and mysql 
