192 CHAPTER 5. DIGITAL SIGNAL PROCESSING

Solution to Exercise 5.10 (p. 151)
A 16-bit A/D converter yields a SNR of 6 16 + 10log 1 : 5 = 97 : 8 dB.
Solution to Exercise 5.11 (p. 154)

S e =
X
s ( n ) e

=
X
e s ( n ) e

=
X
s ( n ) e

= S e

(5.56)

Solution to Exercise 5.12 (p. 155)

X X
=

which, after manipulation, yields the geometric sum formula.
Solution to Exercise 5.13 (p. 158)
If the sampling frequency exceeds the Nyquist frequency, the spectrum of the samples equals the analog
spectrum, but over the normalized analog frequency fT . Thus, the energy in the sampled signal equals the
original signal's energy multiplied by T .
Solution to Exercise 5.14 (p. 158)
After lowpass ltering, the interspersed zeros at the odd indices would be lled in with signal values that
correspond to the original analog signal sampled at twice the rate. Consequently, stretching the signal by
a factor of M then lowpass ltering with a cuto frequency of 1 = 2 M corresponds to up-sampling (a virtual
increase of the sampling rate) by a factor of M .
Solution to Exercise 5.15 (p. 160)
This situation amounts to aliasing in the time-domain.
Solution to Exercise 5.16 (p. 161)
Since conjugate symmetry requires S ( N= 2) = S ( N= 2) , the spectrum at the halfway index must be real-
valued.
Solution to Exercise 5.17 (p. 162)
When the signal is real-valued, we may only need half the spectral values, but the complexity remains
unchanged. If the data are complex-valued, which demands retaining all frequency values, the complexity is
again the same. When only K frequencies are needed, the complexity is O ( KN ) .
Solution to Exercise 5.18 (p. 162)
If a DFT required 1 ms to compute, and signal having ten times the duration would require 100ms to
compute. Using the FFT, a 1ms computing time would increase by a factor of about 10 log 10 = 33 , a
factor of three less than the DFT would have needed.
Solution to Exercise 5.19 (p. 164)
The upper panel has not used the FFT algorithm to compute the length-4 DFTs while the lower one has.
The ordering is determined by the algorithm.
Solution to Exercise 5.20 (p. 164)
The transform can have any greater than or equal to the actual duration of the signal. We simply pad the
signal with zero-valued samples until a computationally advantageous signal length results. Recall that the
FFT is an algorithm to compute the DFT. Extending the length of the signal this way merely means we
are sampling the frequency axis more nely than required. To use the Cooley-Tukey algorithm, the length
of the resulting zero-padded signal can be 512, 1024, etc. samples long.
