154 CHAPTER 5. DIGITAL SIGNAL PROCESSING

5.5.7 Discrete-Time Systems

Discrete-time systems can act on discrete-time signals in ways similar to those found in analog signals and
systems. Because of the role of software in discrete-time systems, many more di erent systems can be
envisioned and constructed with programs than can be with analog signals. In fact, a special class of
analog signals can be converted into discrete-time signals, processed with software, and converted back into
an analog signal, all without the incursion of error. For such signals, systems can be easily produced in
software, with equivalent analog realizations di cult, if not impossible, to design.

### 5.6 Discrete-Time Fourier Transform (DTFT)

The Fourier transform of the discrete-time signal s ( n ) is de ned to be

S e =
X
s ( n ) e (5.17)

Frequency here has no units. As should be expected, this de nition is linear, with the transform of a
sum of signals equaling the sum of their transforms. Real-valued signals have conjugate-symmetric spectra:
S e = S e .

Exercise 5.11 (Solution on p. 192.)
A special property of the discrete-time Fourier transform is that it is periodic with period one:
S e = S e . Derive this property from the de nition of the DTFT.

Because of this periodicity, we need only plot the spectrum over one period to understand completely the
spectrum's structure; typically, we plot the spectrum over the frequency range
;
. When the signal

is real-valued, we can further simplify our plotting chores by showing the spectrum only over 0 ;
; the
spectrum at negative frequencies can be derived from positive-frequency spectral values.
When we obtain the discrete-time signal via sampling an analog signal, the Nyquist frequency (p. 148)
corresponds to the discrete-time frequency
. To show this, note that a sinusoid having a frequency equal

to the Nyquist frequency
1 T
has a sampled waveform that equals

cos 2
1 T s
nT s = cos ( n ) = ( 1)

The exponential in the DTFT at frequency
equals e
= e = ( 1) , meaning that discrete-time
frequency equals analog frequency multiplied by the sampling interval

f = f T (5.18)

f and f represent discrete-time and analog frequency variables, respectively. Figure 5.4 inspires another
way of deriving this result. As the duration of each pulse in the periodic sampling signal p ( t ) narrows, the
amplitudes of the signal's spectral repetitions, which are governed by the Fourier series coe cients (4.10) of
p ( t ) , become increasingly equal. Examination of the periodic pulse signal (Figure 4.1) reveals that as

decreases, the value of c , the largest Fourier coe cient, decreases to zero: j c j =
A
. Thus, to maintain

a mathematically viable Sampling Theorem, the amplitude A must increase as
1
, becoming in nitely large

as the pulse duration decreases. Practical systems use a small value of , say 0 : 1 T and use ampli ers to

rescale the signal. Thus, the sampled signal's spectrum becomes periodic with period
1
. Thus, the Nyquist

frequency
1 T
corresponds to the frequency
1
.
This content is available online at http://cnx.org/content/m10247/2.31/ .
