145

(base-16). Fractions between zero and one are represented the same way.

f =
X
d b ; d 2 f 0 ; : : : ; b 1 g (5.2)

All numbers can be represented by their sign, integer and fractional parts. Complex numbers can be thought
of as two real numbers that obey special rules to manipulate them.
Humans use base-10, commonly assumed to be due to us having ten ngers. Digital computers use the
base-2 or binary number representation, each digit of which is known as a bit ( b inary dig it ).

Figure 5.2: The various ways numbers are represented in binary are illustrated. The number of bytes
for the exponent and mantissa components of oating point numbers varies.

Here, each bit is represented as a voltage that is either high or low, thereby representing 1 or 0,
respectively. To represent signed values, we tack on a special bit the sign bit to express the sign. The
computer's memory consists of an ordered sequence of bytes , a collection of eight bits. A byte can therefore
represent an unsigned number ranging from 0 to 255 . If we take one of the bits and make it the sign bit, we
can make the same byte to represent numbers ranging from 128 to 127 . But a computer cannot represent
all possible real numbers. The fault is not with the binary number system; rather having only a nite number
of bytes is the problem. While a gigabyte of memory may seem to be a lot, it takes an in nite number of bits
to represent . Since we want to store many numbers in a computer's memory, we are restricted to those
that have a nite binary representation. Large integers can be represented by an ordered sequence of bytes.
Common lengths, usually expressed in terms of the number of bits, are 16, 32, and 64. Thus, an unsigned
32-bit number can represent integers ranging between 0 and 2 1 (4,294,967,295), a number almost big
enough to enumerate every human in the world!

Exercise 5.1 (Solution on p. 191.)
For both 32-bit and 64-bit integer representations, what are the largest numbers that can be rep-
resented if a sign bit must also be included.

While this system represents integers well, how about numbers having nonzero digits to the right of the
decimal point? In other words, how are numbers that have fractional parts represented? For such numbers,
the binary representation system is used, but with a little more complexity. The oating-point system
uses a number of bytes typically 4 or 8 to represent the number, but with one byte (sometimes two
bytes) reserved to represent the exponent e of a power-of-two multiplier for the number the mantissa
m expressed by the remaining bytes.
x = m 2 (5.3)

The mantissa is usually taken to be a binary fraction having a magnitude in the range
; 1 , which means
that the binary representation is such that d = 1 . The number zero is an exception to this rule, and it is
You need one more bit to do that.
In some computers, this normalization is taken to an extreme: the leading binary digit is not explicitly expressed, providing
an extra bit to represent the mantissa a little more accurately. This convention is known as the hidden-ones notation .
