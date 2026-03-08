258 CHAPTER 6. INFORMATION COMMUNICATION

Figure 6.44

Solution to Exercise 6.28 (p. 226)
In binary arithmetic (see Table 6.2), adding 0 to a binary value results in that binary value while adding 1
results in the opposite binary value.
Solution to Exercise 6.29 (p. 228)

d = 2 n + 1

Solution to Exercise 6.30 (p. 229)
When we multiply the parity-check matrix times any codeword equal to a column of G , the result consists of
the sum of an entry from the lower portion of G and itself that, by the laws of binary arithmetic, is always
zero.
Because the code is linear sum of any two codewords is a codeword we can generate all codewords as
sums of columns of G . Since multiplying by H is also linear, H c = 0 .
Solution to Exercise 6.31 (p. 229)
In binary arithmetic, adding 0 to a binary value results in that binary value while adding 1 results in the
opposite binary value.
Solution to Exercise 6.32 (p. 230)

The probability of a single-bit error in a length- N block is Np (1 p ) and a triple-bit error has prob-

ability
N

3
p (1 p ) . For the rst to be greater than the second, we must have

p <
1 N 1) ( N 2)
+ 1
