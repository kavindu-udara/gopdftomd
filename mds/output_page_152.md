146 CHAPTER 5. DIGITAL SIGNAL PROCESSING

the only oating point number having a zero fraction. The sign of the mantissa represents the sign of the
number and the exponent can be a signed integer.
A computer's representation of integers is either perfect or only approximate, the latter situation occurring
when the integer exceeds the range of numbers that a limited set of bytes can represent. Floating point
representations have similar representation problems: if the number x can be multiplied/divided by enough
powers of two to yield a fraction lying between 1 = 2 and 1 that has a nite binary-fraction representation, the
number is represented exactly in oating point. Otherwise, we can only represent the number approximately,
not catastrophically in error as with integers. For example, the number 2 : 5 equals 0 : 625 2 , the fractional
part of which has an exact binary representation. However, the number 2 : 6 does not have an exact binary
representation, and only be represented approximately in oating point. In single precision oating point
numbers, which require 32 bits (one byte for the exponent and the remaining 24 bits for the mantissa), the
number 2.6 will be represented as 2 : 600000079 : : : . Note that this approximation has a much longer decimal
expansion. This level of accuracy may not su ce in numerical calculations. Double precision oating
point numbers consume 8 bytes, and quadruple precision 16 bytes. The more bits used in the mantissa,
the greater the accuracy. This increasing accuracy means that more numbers can be represented exactly, but
there are always some that cannot. Such inexact numbers have an in nite binary representation. Realizing
that real numbers can be only represented approximately is quite important, and underlies the entire eld
of numerical analysis , which seeks to predict the numerical accuracy of any computation.

Exercise 5.2 (Solution on p. 191.)
What are the largest and smallest numbers that can be represented in 32-bit oating point? in
64-bit oating point that has sixteen bits allocated to the exponent? Note that both exponent and
mantissa require a sign bit.

So long as the integers aren't too large, they can be represented exactly in a computer using the binary
positional notation. Electronic circuits that make up the physical computer can add and subtract integers
without error. (This statement isn't quite true; when does addition cause problems?)

5.2.3 Computer Arithmetic and Logic

The binary addition and multiplication tables are

0 + 0 = 0 0 0 = 0

0 + 1 = 1 0 1 = 0

1 + 1 = 10 1 1 = 1

1 + 0 = 1 1 0 = 0

(5.4)

Note that if carries are ignored, subtraction of two single-digit binary numbers yields the same bit as
addition. Computers use high and low voltage values to express a bit, and an array of such voltages express
numbers akin to positional notation. Logic circuits perform arithmetic operations.

Exercise 5.3 (Solution on p. 191.)
Add twenty- ve and seven in base-2. Note the carries that might occur. Why is the result nice?

The variables of logic indicate truth or falsehood. A
T
B , the AND of A and B , represents a statement
that both A and B must be true for the statement to be true. You use this kind of statement to tell search
engines that you want to restrict hits to cases where both of the events A and B occur. A
S
B , the OR of
A and B , yields a value of truth if either is true. Note that if we represent truth by a 1 and falsehood
by a 0, binary multiplication corresponds to AND and addition (ignoring carries) to XOR .
XOR, the exclusive or operator, equals the union of A
S
B and A
T
B . The Irish mathematician George
See if you can nd this representation.
Note that there will always be numbers that have an in nite representation in any chosen positional system. The choice
of base de nes which do and which don't. If you were thinking that base-10 numbers would solve this inaccuracy, note that
1 = 3 = 0 : 333333 :::: has an in nite representation in decimal (and binary for that matter), but has nite representation in base-3.
A carry means that a computation performed at a given position a ects other positions as well. Here, 1 + 1 = 10 is an
example of a computation that involves a carry.
