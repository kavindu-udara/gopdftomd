193

Solution to Exercise 5.21 (p. 164)
Number of samples equals 1 : 2 11025 = 13230 . The datarate is 11025 16 = 176 : 4 kbps. The storage
required would be 26460 bytes.
Solution to Exercise 5.22 (p. 165)
The oscillations are due to the boxcar window's Fourier transform, which equals the sinc function.
Solution to Exercise 5.23 (p. 167)
These numbers are powers-of-two, and the FFT algorithm can be exploited with these lengths. To compute
a longer transform than the input signal's duration, we simply zero-pad the signal.
Solution to Exercise 5.24 (p. 167)
In discrete-time signal processing, an ampli er amounts to a multiplication, a very easy operation to perform.
Solution to Exercise 5.25 (p. 169)
The indices can be negative, and this condition is not allowed in MATLAB. To x it, we must start the
signals later in the array.
Solution to Exercise 5.26 (p. 170)
Such terms would require the system to know what future input or output values would be before the current
value was computed. Thus, such terms can cause di culties.
Solution to Exercise 5.27 (p. 172)
It now acts like a bandpass lter with a center frequency of f and a bandwidth equal to twice of the original
lowpass lter.
Solution to Exercise 5.28 (p. 172)
The DTFT of the unit sample equals a constant (equaling 1). Thus, the Fourier transform of the output
equals the transfer function.
Solution to Exercise 5.29 (p. 173)
In sampling a discrete-time signal's Fourier transform L times equally over [0 ; 2 ) to form the DFT, the
corresponding signal equals the periodic repetition of the original signal.

S ( k ) !
X
s ( n iL ) (5.57)

To avoid aliasing (in the time domain), the transform length must equal or exceed the signal's duration.
Solution to Exercise 5.30 (p. 173)
The di erence equation for an FIR lter has the form

y ( n ) =
X
b x ( n m ) (5.58)

The unit-sample response equals

h ( n ) =
X
b ( n m ) (5.59)

which corresponds to the representation described in a problem (Example 5.6) of a length- q boxcar lter.
Solution to Exercise 5.31 (p. 173)
The unit-sample response's duration is q + 1 and the signal's N . Thus the statement is correct.
Solution to Exercise 5.32 (p. 177)
Let N denote the input's total duration. The time-domain implementation requires a total of N (2 q + 1)

computations, or 2 q + 1 computations per input value. In the frequency domain, we split the input into
N

sections, each of which requires log ( N + q )+6+
q + q
per input in the section. Because we divide again

by N to nd the number of computations per input value in the entire input, this quantity decreases as
N increases. For the time-domain implementation, it stays constant.
