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