140 CHAPTER 4. FREQUENCY DOMAIN

Solution to Exercise 4.7 (p. 107)

Total harmonic distortion equals

2
X
j c j

X
j c j

. The signal's average value, represented by c , is ignored in

harmonic distortion calculations. The factor of two in the numerator and denominator arises because we
need to sum over negative as well as positive frequency. Because for real-valued signals, c = c , the sums
over these frequency regions are the same. Clearly, the factors of two cancel. Total harmonic distortion
is most easily computed in the frequency domain. However, the numerator can be computed in the time
domain by noting that it equals equals the square of the signal's rms value (after subtracting the signal's
average value) minus the power in the rst harmonic.
Solution to Exercise 4.8 (p. 108)

Total harmonic distortion in the square wave is 1

= 20% .
Solution to Exercise 4.9 (p. 111)

N signals directly encoded require a bandwidth of N=T . Using a binary representation, we need
log N
.

For N = 128 , the binary-encoding scheme has a factor of
7
= 0 : 05 smaller bandwidth. Clearly, binary

encoding is superior.
Solution to Exercise 4.10 (p. 111)
We can use N di erent amplitude values at only one frequency to represent the various letters.
Solution to Exercise 4.11 (p. 113)
Because the lter's gain at zero frequency equals one, the average output values equal the respective average
input values.
Solution to Exercise 4.12 (p. 115)

F ( S ( f )) =

Z

S ( f ) e df =

Z

S ( f ) e df = s ( t )

Solution to Exercise 4.13 (p. 116)
F ( F ( F ( F ( s ( t ))))) = s ( t ) . We know that F ( S ( f )) =
R
S ( f ) e df =
R
S ( f ) e df =
s ( t ) . Therefore, two Fourier transforms applied to s ( t ) yields s ( t ) . We need two more to get us back
where we started.
Solution to Exercise 4.14 (p. 118)
The signal is the inverse Fourier transform of the triangularly shaped spectrum, and equals s ( t ) =

W
sin ( W t )
Solution to Exercise 4.15 (p. 118)
The result is most easily found in the spectrum's formula: the power in the signal-related part of x ( t ) is half
the power of the signal s ( t ) .
Solution to Exercise 4.16 (p. 120)
The inverse transform of the frequency response is
e u( t ) . Multiplying the frequency response
by 1 e means subtract from the original signal its time-delayed version. Delaying the frequency
response's time-domain version by results in
e u( t ) . Subtracting from the undelayed

signal yields
e u( t )
e u( t ) . Now we integrate this sum. Because the integral of
a sum equals the sum of the component integrals (integration is linear), we can consider each separately.
Because integration and signal-delay are linear, the integral of a delayed signal equals the delayed version of
the integral. The integral is provided in the example (4.44).
Solution to Exercise 4.17 (p. 122)
If the glottis were linear, a constant input (a zero-frequency sinusoid) should yield a constant output. The
periodic output indicates nonlinear behavior.
