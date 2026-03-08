131

(b) Calculate the Fourier transform of the two-pulse sequence shown below (Figure 4.24(b)).
(c) Calculate the Fourier transform for the ten -pulse sequence shown in below (Figure 4.24(c)). You
should look for a general expression that holds for sequences of any length.
(d) Using Matlab, plot the magnitudes of the three spectra. Describe how the spectra change as the
number of repeated pulses increases.

(a)

(b)

(c)

Figure 4.24

Problem 4.9: Spectra of Digital Communication Signals
One way to represent bits with signals is shown in Figure 4.25. If the value of a bit is a 1, it is represented
by a positive pulse of duration T . If it is a 0, it is represented by a negative pulse of the same duration.
To represent a sequence of bits, the appropriately chosen pulses are placed one after the other.

(a) What is the spectrum of the waveform that represents the alternating bit sequence : : : 01010101 : : : ?
(b) This signal's bandwidth is de ned to be the frequency range over which 90% of the power is contained.
What is this signal's bandwidth?
(c) Suppose the bit sequence becomes : : : 00110011 : : : . Now what is the bandwidth?

Problem 4.10: Lowpass Filtering a Square Wave
Let a square wave (period T ) serve as the input to a rst-order lowpass system constructed as a RC lter.
We want to derive an expression for the time-domain response of the lter to this input.

(a) First, consider the response of the lter to a simple pulse, having unit amplitude and width
. Derive
an expression for the lter's output to this pulse.
(b) Noting that the square wave is a superposition of a sequence of these pulses, what is the lter's response
to the square wave?
(c) The nature of this response should change as the relation between the square wave's period and the
lter's cuto frequency change. How long must the period be so that the response does not achieve

Figure 4.25
