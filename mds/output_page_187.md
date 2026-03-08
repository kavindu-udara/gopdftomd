181

Figure 5.26

(a) Find the Fourier spectrum of this signal.
(b) Will this scheme work? If so, how should T be related to the signal's bandwidth? If not, why not?

Problem 5.4: Bandpass Sampling
The signal s ( t ) has the spectrum shown in Figure 5.27.

Figure 5.27

(a) What is the minimum sampling rate for this signal suggested by the Sampling Theorem?
(b) Because of the particular structure of this spectrum, one wonders whether a lower sampling rate could
be used. Show that this is indeed the case, and nd the system that reconstructs s ( t ) from its samples.

Problem 5.5: Sampling Signals

If a signal is bandlimited to W Hz, we can sample it at any rate
1
> 2 W and recover the waveform exactly.

This statement of the Sampling Theorem can be taken to mean that all information about the original
signal can be extracted from the samples. While true in principle, you do have to be careful how you do
so. In addition to the rms value of a signal, an important aspect of a signal is its peak value, which equals
max fj s ( t ) jg .

(a) Let s ( t ) be a sinusoid having frequency W Hz. If we sample it at precisely the Nyquist rate, how
accurately do the samples convey the sinusoid's amplitude? In other words, nd the worst case example.
(b) How fast would you need to sample for the amplitude estimate to be within 5% of the true value?
(c) Another issue in sampling is the inherent amplitude quantization produced by A/D converters. Assume
the maximum voltage allowed by the converter is V volts and that it quantizes amplitudes to
b bits. We can express the quantized sample Q ( s ( nT )) as s ( nT ) + ( t ) , where ( t ) represents
the quantization error at the n sample. Assuming the converter rounds, how large is maximum
quantization error?
