220 CHAPTER 6. INFORMATION COMMUNICATION

Example 6.2
A four-symbol alphabet has the following probabilities.

Pr [ a ] =
1
Pr [ a ] =
1
Pr [ a ] =
1
Pr [ a ] =
1

and an entropy of 1.75 bits (Example 6.1). Let's see if we can nd a codebook for this four-letter
alphabet that satis es the Source Coding Theorem. The simplest code to try is known as the
simple binary code : convert the symbol's index into a binary number and use the same number
of bits for each symbol by including leading zeros where necessary.

( a $ 00) ( a $ 01) ( a $ 10) ( a $ 11) (6.53)

Whenever the number of symbols in the alphabet is a power of two (as in this case), the average
number of bits ( A ) equals log K , which equals 2 in this case. Because the entropy equals 1 : 75
bits, the simple binary code indeed satis es the Source Coding Theorem we are within one bit
of the entropy limit but you might wonder if you can do better. If we chose a codebook with
di ering number of bits for the symbols, a smaller average number of bits could indeed be obtained.
The idea is to use shorter bit sequences for the symbols that occur more often. One codebook like
this is
( a $ 0) ( a $ 10) ( a $ 110) ( a $ 111) (6.54)

Now ( A ) = 1
+ 2
+ 3
+ 3
= 1 : 75 . We can reach the entropy limit! The simple binary
code is, in this case, less e cient than the unequal-length code. Using the e cient code, we can
transmit the symbolic-valued signal having this alphabet 12.5% faster. Furthermore, we know that
no more e cient codebook can be found because of Shannon's Theorem.

### 6.22 Compression and the Hu man Code

Shannon's Source Coding Theorem (6.52) has additional applications in data compression . Here, we have
a symbolic-valued signal source, like a computer le or an image, that we want to represent with as few
bits as possible. Compression schemes that assign symbols to bit sequences are known as lossless if they
obey the Source Coding Theorem; they are lossy if they use fewer bits than the alphabet's entropy. Using
a lossy compression scheme means that you cannot recover a symbolic-valued signal from its compressed
version without incurring some error. You might be wondering why anyone would want to intentionally
create errors, but lossy compression schemes are frequently used where the e ciency gained in representing
the signal outweighs the signi cance of the errors.
Shannon's Source Coding Theorem states that symbolic-valued signals require on the average at least
H ( A ) number of bits to represent each of its values, which are symbols drawn from the alphabet A . In the
Section 6.21, we found that using a so-called xed rate source coder, one that produces a xed number of
bits/symbol, may not be the most e cient way of encoding symbols into bits. What is not discussed there is
a procedure for designing an e cient source coder: one guaranteed to produce the fewest bits/symbol on
the average. That source coder is not unique, and one approach that does achieve that limit is the Hu man
source coding algorithm .

Point of Interest: In the early years of information theory, the race was on to be the rst to nd
a provably maximally e cient source coding algorithm. The race was won by then MIT graduate
student David Hu man in 1954, who worked on the problem as a project in his information theory
course. We're pretty sure he received an A.

Create a vertical table for the symbols, the best ordering being in decreasing order of probability.
This content is available online at http://cnx.org/content/m0092/2.19/ .
https://en.wikipedia.org/wiki/David_A._Huffman
