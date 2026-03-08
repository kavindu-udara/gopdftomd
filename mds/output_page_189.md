183

The voltage at time t = 4 T should diminish by a factor of two the further the pulse occurs from this
time. In other words, the voltage due to a pulse at 3 T should be twice that of a pulse produced at 2 T ,
which in turn is twice that of a pulse at T , etc.
The 4-bit D/A converter must support a 10 kHz sampling rate.

Show the circuit that works. How do the converter's parameters change with sampling rate and number of
bits in the converter?

Problem 5.8: Discrete-Time Fourier Transforms
Find the Fourier transforms of the following sequences, where s ( n ) is some sequence having Fourier transform
S e .

(a) ( 1) s ( n ) (b) s ( n ) cos (2 f n )

(c) x ( n ) =

(
s
n even

0 n odd
(d) ns ( n )

Problem 5.9: Spectra of Finite-Duration Signals
Find the indicated spectra for the following signals.

(a) The discrete-time Fourier transform of s ( n ) =

(
cos
n n = f 1 ; 0 ; 1 g

0 otherwise

(b) The discrete-time Fourier transform of s ( n ) =

(
n n = f 2 ; 1 ; 0 ; 1 ; 2 g

0 otherwise

(c) The discrete-time Fourier transform of s ( n ) =

(
sin
n n = f 0 ; : : : ; 7 g

0 otherwise

(d) The length-8 DFT of the previous signal.

Problem 5.10: Just Whistlin'
Sammy loves to whistle and decides to record and analyze his whistling in lab. He is a very good whistler; his
whistle is a pure sinusoid that can be described by s ( t ) = sin (4000 t ) . To analyze the spectrum, he samples
his recorded whistle with a sampling interval of T = 2 : 5 10 to obtain s ( n ) = s ( nT ) . Sammy (wisely)
decides to analyze a few samples at a time, so he grabs 30 consecutive, but arbitrarily chosen, samples. He
calls this sequence x ( n ) and realizes he can write it as

x ( n ) = sin (4000 nT + ) , n = f 0 ; : : : ; 29 g

(a) Did Sammy under- or over-sample his whistle?
(b) What is the discrete-time Fourier transform of x ( n ) and how does it depend on ?
(c) How does the 32-point DFT of x ( n ) depend on ?

Problem 5.11: Discrete-Time Filtering
We can nd the input-output relation for a discrete-time lter much more easily than for analog lters. The
key idea is that a sequence can be written as a weighted linear combination of unit samples.

(a) Show that x ( n ) =
P
x ( i ) ( n i ) where ( n ) is the unit-sample.

( n ) =

(
1 n = 0

0 otherwise

(b) If h ( n ) denotes the unit-sample response the output of a discrete-time linear, shift-invariant lter
to a unit-sample input nd an expression for the output.
