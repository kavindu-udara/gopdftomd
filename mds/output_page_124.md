118 CHAPTER 4. FREQUENCY DOMAIN

Using Euler's relation for the cosine,

s ( t ) cos (2 f t ) =
1

Z

S ( f ) e df +
1

Z

S ( f ) e df

=
1

Z

S ( f f ) e df +
1

Z

S ( f + f ) e df

=

Z
S ( f f ) + S ( f + f )
e df

Exploiting the uniqueness property of the Fourier transform, we have

F s ( t ) cos (2 f t ) =
S ( f f ) + S ( f + f )
(4.37)

This component of the spectrum consists of the original signal's spectrum delayed and advanced in
frequency . The spectrum of the amplitude modulated signal is shown in Figure 4.12.

Figure 4.12: A signal which has a triangular shaped spectrum is shown in the top plot. Its highest
frequency the largest frequency containing power is W Hz. Once amplitude modulated, the resulting
spectrum has lines corresponding to the Fourier series components at f and the original triangular
spectrum shifted to components at f and scaled by
.

Note how in this gure the signal s ( t ) is de ned in the frequency domain. To nd its time domain
representation, we simply use the inverse Fourier transform.

Exercise 4.14 (Solution on p. 140.)
What is the signal s ( t ) that corresponds to the spectrum shown in the upper panel of Figure 4.12?

Exercise 4.15 (Solution on p. 140.)
What is the power in x ( t ) , the amplitude-modulated signal? Try the calculation in both the time
and frequency domains.

In this example, we call the signal s ( t ) a baseband signal because its power is contained at low frequencies.
Signals such as speech and the Dow Jones averages are baseband signals. The baseband signal's bandwidth
equals W , the highest frequency at which it has power. Since x ( t ) 's spectrum is con ned to a frequency band
not close to the origin (we assume ( f W ) ), we have a bandpass signal . The bandwidth of a bandpass
signal is not its highest frequency, but the range of positive frequencies where the signal has power. Thus,
in this example, the bandwidth is 2 W Hz . Why a signal's bandwidth should depend on its spectral shape
will become clear once we develop communications systems.
