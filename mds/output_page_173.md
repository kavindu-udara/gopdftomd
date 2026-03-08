167

Figure 5.17: The original speech segment and the sequence of overlapping Hanning windows applied
to it are shown in the upper portion. Frames were 256 samples long and a Hanning window was applied
with a half-frame overlap. A length-512 FFT of each frame was computed, with the magnitude of the
rst 257 FFT values displayed vertically, with spectral amplitude values color-coded.
are much better behaved and spectral changes are much better captured.
The speech signal, such as shown in Figure 5.14, is sectioned into overlapping, equal-length frames, with
a Hanning window applied to each frame. The spectra of each of these is calculated, and displayed in
spectrograms with frequency extending vertically, window time location running horizontally, and spectral
magnitude color-coded. Figure 5.17 illustrates these computations.
Exercise 5.23 (Solution on p. 193.)
Why the speci c values of 256 for N and 512 for K ? Another issue is how was the length-512
transform of each length-256 windowed frame computed?

### 5.11 Discrete-Time Systems

When we developed analog systems, interconnecting the circuit elements provided a natural starting place for
constructing useful devices. In discrete-time signal processing, we are not limited by hardware considerations
but by what can be constructed in software.

Exercise 5.24 (Solution on p. 193.)
One of the rst analog systems we described was the ampli er (Section 2.6.2). We found that
implementing an ampli er was di cult in analog systems, requiring an op-amp at least. What is
the discrete-time implementation of an ampli er? Is this especially hard or easy?

In fact, we will discover that frequency-domain implementation of systems, wherein we multiply the input
signal's Fourier transform by a frequency response, is not only a viable alternative, but also a computationally
e cient one. We begin with discussing the underlying mathematical structure of linear, shift-invariant
systems, and devise how software lters can be constructed.
This content is available online at http://cnx.org/content/m0507/2.5/ .
