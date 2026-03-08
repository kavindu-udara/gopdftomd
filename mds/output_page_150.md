144 CHAPTER 5. DIGITAL SIGNAL PROCESSING

Figure 5.1: Generic computer hardware organization. input devices (keyboard, mouse, A/D (analog-to-digital) converter, etc.),
a computational unit , and output devices (monitors, printers, D/A converters). The computational unit
is the computer's heart, and usually consists of a central processing unit (CPU), a memory , and an
input/output (I/O) interface. What I/O devices might be present on a given computer vary greatly.

A simple computer operates fundamentally in discrete time. Computers are clocked devices,
in which computational steps occur periodically according to ticks of a clock. This description belies
clock speed: When you say I have a 1 GHz computer, you mean that your computer takes 1 nanosec-
ond to perform each step. That is incredibly fast! A step does not, unfortunately, necessarily mean
a computation like an addition; computers break such computations down into several stages, which
means that the clock speed need not express the computational speed. Computational speed is ex-
pressed in units of millions of instructions/second (Mips). Your 1 GHz computer (clock speed) may
have a computational speed of 200 Mips.
Computers perform integer (discrete-valued) computations. Computer calculations can be
numeric (obeying the laws of arithmetic), logical (obeying the laws of an algebra), or symbolic (obeying
any law you like). Each computer instruction that performs an elementary numeric calculation
an addition, a multiplication, or a division does so only for integers. The sum or product of two
integers is also an integer, but the quotient of two integers is likely to not be an integer. How does
a computer deal with numbers that have digits to the right of the decimal point? This problem is
addressed by using the so-called oating-point representation of real numbers. At its heart, however,
this representation relies on integer-valued computations.

5.2.2 Representing Numbers

Focusing on numbers, all numbers can represented by the positional notation system . The b -ary po-
sitional representation system uses the position of digits ranging from 0 to b -1 to denote a number. The
quantity b is known as the base of the number system. Mathematically, positional systems represent the
positive integer n as

n =
X
d b ; d 2 f 0 ; : : : ; b 1 g (5.1)

and we succinctly express n in base- b as n = d d : : : d . The number 25 in base-10 equals 2 10 +5 10 ,
so that the digits representing this number are d = 5 , d = 2 , and all other d equal zero. This same
number in binary (base-2) equals 11001 ( 1 2 + 1 2 + 0 2 + 0 2 + 1 2 ) and 19 in hexadecimal
An example of a symbolic computation is sorting a list of names.
Alternative number representation systems exist. For example, we could use stick gure counting or Roman numerals.
These were useful in ancient times, but very limiting when it comes to arithmetic calculations: ever tried to divide two Roman
numerals?
