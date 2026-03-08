179

### 5.16 Discrete-Time Filtering of Analog Signals

Because of the Sampling Theorem, we can process, in particular lter, analog signals with a computer by
constructing the system shown in Figure 5.24. To use this system, we are assuming that the input signal has
a lowpass spectrum and can be bandlimited without a ecting important signal aspects. Bandpass signals
can also be ltered digitally, but require a more complicated system. Highpass signals cannot be ltered
digitally. Note that the input and output lters must be analog lters; trying to operate without them can
lead to potentially very inaccurate digitization.

Figure 5.24: To process an analog signal digitally, the signal x ( t ) must be ltered with an anti-aliasing
lter (to ensure a bandlimited signal) before A/D conversion. This lowpass lter (LPF) has a cuto
frequency of W Hz, which determines allowable sampling intervals T . The greater the number of bits
in the amplitude quantization portion Q [ ] of the A/D converter, the greater the accuracy of the entire
system. The resulting digital signal x ( n ) can now be ltered in the time-domain with a di erence
equation or in the frequency domain with Fourier transforms. The resulting output y ( n ) then drives a
D/A converter and a second anti-aliasing lter (having the same bandwidth as the rst one).

Another implicit assumption is that the digital lter can operate in real time : The computer and the
ltering algorithm must be su ciently fast so that outputs are computed faster than input values arrive.
The sampling interval, which is determined by the analog signal's bandwidth, thus determines how long our
program has to compute each output y ( n ) . The computational complexity for calculating each output with
a di erence equation (5.42) is O ( p + q ) . Frequency domain implementation of the lter is also possible.
The idea begins by computing the Fourier transform of a length- N portion of the input x ( n ) , multiplying
it by the lter's transfer function, and computing the inverse transform of the result. This approach seems
overly complex and potentially ine cient. Detailing the complexity, however, we have O ( N log N ) for the two
transforms (computed using the FFT algorithm) and O ( N ) for the multiplication by the transfer function,
which makes the total complexity O ( N log N ) for N input values . A frequency domain implementation
thus requires O (log N ) computational complexity for each output value. The complexities of time-domain
and frequency-domain implementations depend on di erent aspects of the ltering: The time-domain imple-
mentation depends on the combined orders of the lter while the frequency-domain implementation depends
on the logarithm of the Fourier transform's length.
It could well be that in some problems the time-domain version is more e cient (more easily satis es the
real time requirement), while in others the frequency domain approach is faster. In the latter situations, it is
the FFT algorithm for computing the Fourier transforms that enables the superiority of frequency-domain
implementations. Because complexity considerations only express how algorithm running-time increases with
system parameter choices, we need to detail both implementations to determine which will be more suitable
for any given ltering problem. Filtering with a di erence equation is straightforward, and the number of
computations that must be made for each output value is 2 ( p + q ) .

Exercise 5.34 (Solution on p. 194.)
Derive this value for the number of computations for the general di erence equation (5.42).
This content is available online at http://cnx.org/content/m0511/2.21/ .
