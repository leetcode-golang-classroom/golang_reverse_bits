# golang_reverse_bits

Reverse bits of a given 32 bits unsigned integer.

**Note:**

- Note that in some languages, such as Java, there is no unsigned integer type. In this case, both input and output will be given as a signed integer type. They should not affect your implementation, as the integer's internal binary representation is the same, whether it is signed or unsigned.
- In Java, the compiler represents the signed integers using [2's complement notation](https://en.wikipedia.org/wiki/Two%27s_complement). Therefore, in **Example 2** above, the input represents the signed integer `3` and the output represents the signed integer `1073741825`.

## Examples

**Example 1:**

```
Input: n = 00000010100101000001111010011100
Output:    964176192 (00111001011110000010100101000000)
Explanation:The input binary string00000010100101000001111010011100 represents the unsigned integer 43261596, so return 964176192 which its binary representation is00111001011110000010100101000000.

```

**Example 2:**

```
Input: n = 11111111111111111111111111111101
Output:   3221225471 (10111111111111111111111111111111)
Explanation:The input binary string11111111111111111111111111111101 represents the unsigned integer 4294967293, so return 3221225471 which its binary representation is10111111111111111111111111111111.

```

**Constraints:**

- The input must be a **binary string** of length `32`

**Follow up:** If this function is called many times, how would you optimize it?

## 解析

給定一個 32-bit 的 unsigned integer  num

要求寫一個演算法來反轉 bit

建立一個 unsigned int res

逐步針對每個 bit 做以下操作

先把 res << 1 (unsigned left shift 1 bit)

每次先對 num & 1

然後把這個值 加到一個 unsigned int res

如下圖

![](https://i.imgur.com/v03ehe1.png)

等 loop 結束後

會剛好把原本的 num 做 bit 反轉

因為總共是 32 bit 所以時間複雜度是 O(1)

## 程式碼
```go
package sol

func reverseBits(num uint32) uint32 {
	var ans uint32 = 0
	for pos := 0; pos < 32; pos++ {
		ans = ans << 1
		if (num & 1) != 0 {
			ans += 1
		}
		num = num >> 1
	}
	return ans
}

```
## 困難點

1. 需要想出左移的作法來搬 copy 出來的 bit
2. 要注意要做 unsigned shift 對於左移跟右移

## Solve Point

- [x]  建立一個 unsigned integer res
- [x]  從 i = 0.. 31 做以下操作
- [x]  res = res << 1
- [x]  if (num&1) ≠ 0 res += 1
- [x]  num = num >> 1
- [x]  回傳 res