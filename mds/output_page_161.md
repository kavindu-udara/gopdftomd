155

Example 5.1
Let's compute the discrete-time Fourier transform of the exponentially decaying sequence s ( n ) =
a u( n ) , where u( n ) is the unit-step sequence. Simply plugging the signal's expression into the
Fourier transform formula,

S e =
X
a u( n ) e

=
X
ae

(5.19)

This sum is a special case of the geometric series .
X
=
1
; j j < 1 (5.20)

Let = ae . As long as j a j < 1 , j j < 1 . We nd our Fourier transform to be

S e =
1 ae
; j a j < 1 (5.21)

Using Euler's relation, we can express the magnitude and phase of this spectrum.

j S e j =
1 a cos (2 f )) + a sin (2 f )

(5.22)

\ S e = tan
a sin (2 f ) a cos (2 f )
(5.23)

No matter what value of a we choose, the above formulae clearly demonstrate the periodic nature
of the spectra of discrete-time signals. Figure 5.9 shows indeed that the spectrum is a periodic
function. We need only consider the spectrum between
and
to unambiguously de ne it.
When a > 0 , we have a lowpass spectrum the spectrum diminishes as frequency increases from 0
to
with increasing a leading to a greater low frequency content; for a < 0 , we have a highpass
spectrum (Figure 5.10).

Example 5.2
Analogous to the analog pulse signal, let's nd the spectrum of the length- N pulse sequence.

s ( n ) =

(
1 ; if 0 n N 1

0 ; otherwise
(5.24)

The Fourier transform of this sequence has the form of a truncated geometric series.

S e =
X
e (5.25)

For the so-called nite geometric series, we know that

X
=
1
(5.26)

for all values of .

Exercise 5.12 (Solution on p. 192.)
Derive this formula for the nite geometric series sum. The trick is to consider the di erence
between the series' sum and the sum of the series multiplied by .
