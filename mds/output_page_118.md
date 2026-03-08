112 CHAPTER 4. FREQUENCY DOMAIN

Figure 4.9: The encoding of signals via the Fourier spectrum is shown over three periods. In this ex-
ample, only the third and fourth harmonics are used, as shown by the spectral magnitudes corresponding
to each T -second interval plotted below the waveforms. Can you determine the phase of the harmonics
from the waveform?
dently speci ed and that they can be uniquely recovered from the time-domain signal over one period. Do
note that the signal representing the entire document is no longer periodic. By understanding the Fourier se-
ries' properties (in particular that coe cients are determined only over a T -second interval, we can construct
a communications system. This approach represents a simpli cation of how modern modems represent text
that they transmit over telephone lines.

### 4.7 Filtering Periodic Signals

The Fourier series representation of a periodic signal makes it easy to determine how a linear, time-invariant
lter reshapes such signals in general . The fundamental property of a linear system is that its input-output
relation obeys superposition: S [ a s ( t ) + a s ( t )] = a S [ s ( t )] + a S [ s ( t )] . Because the Fourier series
represents a periodic signal as a linear combination of complex exponentials, we can exploit the superposition
property. Furthermore, we found for linear circuits that their output to a complex exponential input is just the
frequency response evaluated at the signal's frequency times the complex exponential. Said mathematically,

if x ( t ) = e
, then the output y ( t ) = H
k
e
because f =
k
. Thus, if x ( t ) is periodic thereby

having a Fourier series, a linear circuit's output to this signal will be the superposition of the output to each
component.

y ( t ) =
X
c H
k
e
(4.27)

Thus, the output has a Fourier series, which means that it too is periodic. Its Fourier coe cients equal

c H
k
. To obtain the spectrum of the output, we simply multiply the input spectrum by

the frequency response . The circuit modi es the magnitude and phase of each Fourier coe cient. Note
especially that while the Fourier coe cients do not depend on the signal's period, the circuit's transfer
function does depend on frequency, which means that the circuit's output will di er as the period varies.
This content is available online at http://cnx.org/content/m0044/2.10/ .
