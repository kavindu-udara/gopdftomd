182 CHAPTER 5. DIGITAL SIGNAL PROCESSING

(d) We can describe the quantization error as noise, with a power proportional to the square of the
maximum error. What is the signal-to-noise ratio of the quantization error for a full-range sinusoid?
Express your result in decibels.

Problem 5.6: Hardware Error
An A/D converter has a curious hardware problem: Every other sampling pulse is half its normal amplitude
(Figure 5.28).

Figure 5.28

(a) Find the Fourier series for this signal.

(b) Can this signal be used to sample a bandlimited signal having highest frequency W =
1 T
?

Problem 5.7: Simple D/A Converter
Commercial digital-to-analog converters don't work this way, but a simple circuit illustrates how they work.
Let's assume we have a B -bit converter. Thus, we want to convert numbers having a B -bit representation
into a voltage proportional to that number. The rst step taken by our simple converter is to represent
the number by a sequence of B pulses occurring at multiples of a time interval T . The presence of a pulse
indicates a 1 in the corresponding bit position, and pulse absence means a 0 occurred. For a 4-bit
converter, the number 13 has the binary representation 1101 ( 13 = 1 2 + 1 2 + 0 2 + 1 2 ) and
would be represented by the depicted pulse sequence. Note that the pulse sequence is backwards from the
binary representation. We'll see why that is.

Figure 5.29

This signal (Figure 5.29) serves as the input to a rst-order RC lowpass lter. We want to design the lter
and the parameters and T so that the output voltage at time 4 T (for a 4-bit converter) is proportional
to the number. This combination of pulse creation and ltering constitutes our simple D/A converter. The
requirements are
