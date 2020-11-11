## Sign Extension

- Task
  - Given w-bit signed interger x
  - Convert it to w + k integer with same value

- Rule
  - Make `k` copies of sign bit
  - x‘ = x<sub>w-1</sub>, ..., x<sub>3</sub>, x<sub>2</sub>, x<sub>1</sub>

### Example
 Give 4-bit signed int convert it to 6-bit
```
// sign bit 0
    0 1 1 0
  0 0 1 1 0
0 0 0 1 1 0

// sign bit 1
    1 1 1 0 =>           -8 + 4 + 2 + 0 = -2
  1 1 1 1 0 =>      -16 + 8 + 4 + 2 + 0 = -2
1 1 1 1 1 0 => -32 + 16 + 8 + 4 + 2 + 0 = -2

```


## Sign Truncating

- Bits are droped
```
// usigned
// mod operation
1 1 0 1 1 => 27
  1 0 1 1 => 27 mod 16 => 11

// signed
1 1 0 1 1 => -5
  1 0 1 1 => -5

// can deal with it as usigned
1 0 0 1 1 => -13
  0 0 1 1 => 19 mod 16 => 3

// no rule
0 1 0 1 1 => 11
  1 0 1 1 => -5
```

## Unsigned Addition
w-bit, will drop w + 1 bit when overflow

## Signd Addition
normally addition, magiclly

## Multiply with Shift
Historically, the integer multiply instruction on many machines was fairly slow, requiring 10 or more clock cycles, whereas other integer operations—such as addition, subtraction, bit-level operations, and shifting—required only 1 clock cycle. As a consequence, one important optimization used by compilers is to attempt to replace multiplications by constant factors with combinations of shift and addition operations. 

B2U<sub>w+k</sub>([x<sub>w-1</sub>, x<sub>w-2</sub>, ..., x<sub>0</sub>, 0, ..., 0]) = x2<sup>k</sup>;

When shifting left by `k` for a fixed w-bits, the high order `k` bits are discard
[x<sub>w−k−1</sub>, x<sub>w−k−2</sub>, . . . , x<sub>0</sub>, 0,  . . . , 0]

combine with addition
> x * 14 = (x<<3) + (x<<2) + (x<<1); 14 = 2<sup>3</sup> + 2<sup>2</sub> + 2<sup>1</sub>

## Divide with Shift(for ints)
- unsigned
  Right shift and round 0
- signed
  Right shift with basis
Right shift

## Byte Ordering
Conventions:
> Big Endian: Sun, PPC Mac, internet
  - Least signification byte has highest address
> Little Endian: x86, ARM processors runing Android, IOS, and Windows
  - Least signification byte has lowest address



## Extra
  ### Representing Strings
  - Represented by array of characters
  - Each character encode in ASCII format
    - Standrad 7-bit encoding of character set
    - Character "0" has code 0x30
      - Digit i has code 0x30+i
  - String should be null-terminated
    - Final character = 0
  - Byter ordering not an issue