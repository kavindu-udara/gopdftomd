178 CHAPTER 5. DIGITAL SIGNAL PROCESSING

Figure 5.23: The gure shows the unit-sample response of a length-17 Hanning lter on the left and
the frequency response on the right. This lter functions as a lowpass lter having a cuto frequency of
about 0.1. N so that N + q = 2 .
Implementing the digital lter shown in the A/D block diagram (Figure 5.24) with a frequency-domain
implementation requires some additional signal management not required by time-domain implementations.
Conceptually, a real-time, time-domain lter could accept each sample as it becomes available, calculate
the di erence equation, and produce the output value, all in less than the sampling interval T . Frequency-
domain approaches don't operate on a sample-by-sample basis; instead, they operate on sections. They
lter in real time by producing N outputs for the same number of inputs faster than N T . Because they
generally take longer to produce an output section than the sampling interval duration, we must lter one
section while accepting into memory the next section to be ltered. In programming, the operation of
building up sections while computing on previous ones is known as bu ering . Bu ering can also be used
in time-domain lters as well but isn't required.

Example 5.8
We want to lowpass lter a signal that contains a sinusoid and a signi cant amount of noise. The
example shown in Figure 5.22 shows a portion of the noisy signal's waveform. If it weren't for the
overlaid sinusoid, discerning the sine wave in the signal is virtually impossible. One of the primary
applications of linear lters is noise removal : preserve the signal by matching lter's passband
with the signal's spectrum and greatly reduce all other frequency components that may be present
in the noisy signal.
A smart Rice engineer has selected a FIR lter having a unit-sample response corresponding a
period-17 sinusoid: h ( n ) =
1 cos (2 n= 17) , n = f 0 ; : : : ; 16 g , which makes q = 16 . Its fre-
quency response (determined by computing the discrete Fourier transform) is shown in Figure 5.23.
To apply, we can select the length of each section so that the frequency-domain ltering approach
is maximally e cient: Choose the section length N so that N + q is a power of two. To use
a length-64 FFT, each section must be 48 samples long. Filtering with the di erence equation
would require 33 computations per output while the frequency domain requires a little over 16; this
frequency-domain implementation is over twice as fast! Figure 5.22 shows how frequency-domain
ltering works.
We note that the noise has been dramatically reduced, with a sinusoid now clearly visible in the
ltered output. Some residual noise remains because noise components within the lter's passband
appear in the output as well as the signal.

Exercise 5.33 (Solution on p. 193.)
Note that when compared to the input signal's sinusoidal component, the output's sinusoidal com-
ponent seems to be delayed. What is the source of this delay? Can it be removed?
