208 CHAPTER 6. INFORMATION COMMUNICATION

### 6.12 Signal-to-Noise Ratio of an Amplitude-Modulated Signal

When we consider the much more realistic situation when we have a channel that introduces attenuation
and noise, we can make use of the just-described receiver's linear nature to directly derive the receiver's
output. The attenuation a ects the output in the same way as the transmitted signal: It scales the output
signal by the same amount. The white noise, on the other hand, should be ltered from the received signal
before demodulation. We must thus insert a bandpass lter having bandwidth 2 W and center frequency f :
This lter has no e ect on the received signal-related component, but does remove out-of-band noise power.
As shown in the triangular-shaped signal spectrum (Figure 6.6), we apply coherent receiver to this ltered
signal, with the result that the demodulated output contains noise that cannot be removed: It lies in the
same spectral band as the signal.
As we derive the signal-to-noise ratio in the demodulated signal, let's also calculate the signal-to-noise
ratio of the bandpass lter's output ~ r ( t ) . The signal component of ~ r ( t ) equals A m ( t ) cos (2 f t ) . This
signal's Fourier transform equals
A
M ( f + f ) + M ( f f ) (6.33)

making the power spectrum,
A
j M ( f + f ) j + j M ( f f ) j (6.34)

Exercise 6.10 (Solution on p. 256.)
If you calculate the magnitude-squared of the rst equation, you don't obtain the second unless
you make an assumption. What is it?

Thus, the total signal-related power in ~ r ( t ) is
A
power ( m ) . The noise power equals the integral of the

noise power spectrum; because the power spectrum is constant over the transmission band, this integral
equals the noise amplitude N times the lter's bandwidth 2 W . The so-called received signal-to-noise
ratio the signal-to-noise ratio after the de rigeur front-end bandpass lter and before demodulation
equals

SNR =
A power ( m ) N W
(6.35)

The demodulated signal b m ( t ) =
A m ( t )
+ n ( t ) . Clearly, the signal power equals
A power ( m )
.

To determine the noise power, we must understand how the coherent demodulator a ects the bandpass
noise found in ~ r ( t ) . Because we are concerned with noise, we must deal with the power spectrum since we
don't have the Fourier transform available to us. Letting P ( f ) denote the power spectrum of ~ r ( t ) 's noise
component, the power spectrum after multiplication by the carrier has the form

P ( f + f ) + P ( f f )
(6.36)

The delay and advance in frequency indicated here results in two spectral noise bands falling in the low-
frequency region of lowpass lter's passband. Thus, the total noise power in this lter's output equals

2
W 2
=
N W
. The signal-to-noise ratio of the receiver's output thus equals

SNR =
A power ( m ) N W

= 2SNR

(6.37)

Let's break down the components of this signal-to-noise ratio to better appreciate how the channel and
the transmitter parameters a ect communications performance. Better performance, as measured by the
SNR , occurs as it increases.
This content is available online at http://cnx.org/content/m0541/2.18/ .
