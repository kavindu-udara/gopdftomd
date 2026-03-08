243

Figure 6.33

Problem 6.7: Secret Communications
A system for hiding AM transmissions has the transmitter randomly switching between two carrier fre-
quencies f and f . Random switching means that one carrier frequency is used for some period of time,
switches to the other for some other period of time, back to the rst, etc. The receiver knows what the
carrier frequencies are but not when carrier frequency switches occur. Consequently, the receiver must be
designed to receive the transmissions regardless of which carrier frequency is used. Assume the message

signal has bandwidth W . The channel adds white noise of spectral height
N
.

(a) How di erent should the carrier frequencies be so that the message could be received?
(b) What receiver would you design?
(c) What signal-to-noise ratio for the demodulated signal does your receiver yield?

Problem 6.8: AM Stereo
Stereophonic radio transmits two signals simultaneously that correspond to what comes out of the left and
right speakers of the receiving radio. While FM stereo is commonplace, AM stereo is not, but is much simpler
to understand and analyze. An amazing aspect of AM stereo is that both signals are transmitted within the
same bandwidth as used to transmit just one. Assume the left and right signals are bandlimited to W Hz.

x ( t ) = A 1 + m ( t ) cos (2 f t ) + Am ( t ) sin (2 f t )

(a) Find the Fourier transform of x ( t ) . What is the transmission bandwidth and how does it compare
with that of standard AM?
(b) Let us use a coherent demodulator as the receiver, shown in Figure 6.34. Show that this receiver indeed
works: It produces the left and right signals separately.
(c) Assume the channel adds white noise to the transmitted signal. Find the signal-to-noise ratio of each
signal.

Figure 6.34
