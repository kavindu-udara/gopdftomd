256 CHAPTER 6. INFORMATION COMMUNICATION

Solution to Exercise 6.8 (p. 205)
The additive-noise channel is not linear because it does not have the zero-input-zero-output property (even
though we might transmit nothing, the receiver's input consists of noise).
Solution to Exercise 6.9 (p. 207)
The signal-related portion of the transmitted spectrum is given by X ( f ) =
M ( f f ) +
M ( f + f ) .
Multiplying at the receiver by the carrier shifts this spectrum to + f and to f , and scales the result by
half.
1
X ( f f ) +
1
X ( f + f ) =
1
( M ( f 2 f ) + M ( f )) +
1
( M ( f + 2 f ) + M ( f ))

=
1
M ( f 2 f ) +
1
M ( f ) +
1
M ( f + 2 f )

(6.67)

The signal components centered at twice the carrier frequency are removed by the lowpass lter, while the
baseband signal M ( f ) emerges.
Solution to Exercise 6.10 (p. 208)
The key here is that the two spectra M ( f f ) , M ( f + f ) do not overlap because we have assumed
that the carrier frequency f is much greater than the signal's highest frequency. Consequently, the term
M ( f f ) M ( f + f ) normally obtained in computing the magnitude-squared equals zero.
Solution to Exercise 6.11 (p. 209)
Separation is 2 W . Commercial AM signal bandwidth is 5kHz . Speech is well contained in this bandwidth,
much better than in the telephone!
Solution to Exercise 6.12 (p. 210)
x ( t ) =
P
s ( t nT ) .
Solution to Exercise 6.13 (p. 211)
k = 4 .
Solution to Exercise 6.14 (p. 211)

x ( t ) =
X
( 1) Ap ( t nT ) sin
2 kt

Solution to Exercise 6.15 (p. 212)
The harmonic distortion is 10%.
Solution to Exercise 6.16 (p. 212)
Twice the baseband bandwidth because both positive and negative frequencies are shifted to the carrier by
the modulation: 3 R .
Solution to Exercise 6.17 (p. 214)
In BPSK, the signals are negatives of each other: s ( t ) = s ( t ) . Consequently, the output of each
multiplier-integrator combination is the negative of the other. Choosing the largest therefore amounts to
choosing which one is positive. We only need to calculate one of these. If it is positive, we are done. If it is
negative, we choose the other signal.
Solution to Exercise 6.18 (p. 214)

The matched lter outputs are
A T
because the sinusoid has less power than a pulse having the same

amplitude.
Solution to Exercise 6.19 (p. 216)
The noise-free integrator outputs di er by A T , the factor of two smaller value than in the baseband case
arising because the sinusoidal signals have less energy for the same amplitude. Stated in terms of E , the
di erence equals 2 E just as in the baseband case.
Solution to Exercise 6.20 (p. 217)

The noise-free integrator output di erence now equals A T =
E
. The noise power remains the same as

in the BPSK case, which from the probability of error equation (6.46) yields p = Q

s E

!

.
