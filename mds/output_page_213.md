207

Figure 6.6: The AM coherent receiver along with the spectra of key signals is shown for the case of a
triangular-shaped signal spectrum. The dashed line indicates the white noise level. Note that the lters'
characteristics cuto frequency and center frequency for the bandpass lter must be match to the
modulation and message parameters. carrier
sinusoid. Frequency modulation (FM) and less frequently used phase modulation (PM) are not discussed
here; we focus on amplitude modulation (AM). The amplitude modulated message signal has the form

x ( t ) = A 1 + m ( t ) cos (2 f t ) (6.29)

where f is the carrier frequency and A the carrier amplitude . Also, the signal's amplitude is assumed
to be less than one: j m ( t ) j < 1 . From our previous exposure to amplitude modulation (see the Fourier
Transform example (Example 4.5)), we know that the transmitted signal's spectrum occupies the frequency
range [ f W; f + W ] , assuming the signal's bandwidth is W Hz (see Figure 6.6). The carrier frequency
is usually much larger than the signal's highest frequency: ( f W ) , which means that the transmitter
antenna and carrier frequency are chosen jointly during the design process.
Ignoring the attenuation and noise introduced by the channel for the moment, reception of an amplitude
modulated signal is quite easy (see Problem 4.20). The so-called coherent receiver multiplies the input
signal by a sinusoid and lowpass- lters the result (Figure 6.6).

b m ( t ) = LPF [ x ( t ) cos (2 f t )]

= LPF A 1 + m ( t ) cos (2 f t )
(6.30)

Because of our trigonometric identities, we know that

cos (2 f t ) =
1
1 + cos (2 2 f t ) (6.31)

At this point, the message signal is multiplied by a constant and a sinusoid at twice the carrier frequency.
Multiplication by the constant term returns the message signal to baseband (where we want it to be!) while
multiplication by the double-frequency term yields a very high frequency signal. The lowpass lter removes
this high-frequency signal, leaving only the baseband signal. Thus, the received signal is

b m ( t ) =
A
1 + m ( t ) (6.32)

Exercise 6.9 (Solution on p. 256.)
This derivation relies solely on the time domain; derive the same result in the frequency domain.
You won't need the trigonometric identity with this approach.

Because it is so easy to remove the constant term by electrical means we insert a capacitor in series with
the receiver's output we typically ignore it and concentrate on the signal portion of the receiver's output
when calculating signal-to-noise ratio.
