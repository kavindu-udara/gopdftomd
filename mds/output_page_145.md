139

### Solutions to Exercises in Chapter 4

Solution to Exercise 4.1 (p. 98)
Because of Euler's relation,

sin (2 ft ) =
1 j
e
1 j
e (4.47)

Thus, c =
1 j
, c =
1 j
, and the other coe cients are zero.

Solution to Exercise 4.2 (p. 102)

c =
A
. This quantity clearly corresponds to the periodic pulse signal's average value.

Solution to Exercise 4.3 (p. 103)
Write the coe cients of the Fourier series in Cartesian form as c = A + jB and substitute into the
expression for the Fourier series.

X
c e
=
X
( A + jB ) e

Simplifying each term in the sum using Euler's formula,

( A + jB ) e
= ( A + jB ) cos
2 kt
+ j sin
2 kt

= A cos
2 kt
B sin
2 kt
+ j A sin
2 kt
+ B cos
2 kt

We now combine terms that have the same frequency index in magnitude . Because the signal is real-valued,
the coe cients of the Fourier series have conjugate symmetry: c = c or A = A and B = B .
After we add the positive-indexed and negative-indexed terms, each term in the Fourier series becomes

2 A cos
2 kt
2 B sin
2 kt
. To obtain the classic Fourier series (4.11), we must have 2 A = a and

2 B = b .
Solution to Exercise 4.4 (p. 104)
The average of a set of numbers is the sum divided by the number of terms. Viewing signal integration as
the limit of a Riemann sum, the integral corresponds to the average.
Solution to Exercise 4.5 (p. 104)
We found that the Fourier series coe cients are given by c =
. The coe cients are pure imaginary,

which means a = 0 . The coe cients of the sine terms are given by b = 2Im [ c ] so that

b =

8
<

:

4
; k odd

0 ; k even

Thus, the Fourier series for the square wave is

sq ( t ) =
X
4
sin
2 kt
(4.48)

Solution to Exercise 4.6 (p. 106)
The rms value of a sinusoid equals its amplitude divided by
p . As a half-wave recti ed sine wave is zero
during half of the period, its rms value is A= 2 since the integral of the squared half-wave recti ed sine wave
equals half that of a squared sinusoid.
