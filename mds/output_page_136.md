130 CHAPTER 4. FREQUENCY DOMAIN

the integral needed to calculate Fourier series coe cients by a Riemann sum as follows.

c =
1

X
s ( n ) e ; k = 0 ; : : : ; 4

(c) What is the harmonic distortion in the two signals? Exclude c from this calculation.
(d) Because the harmonic distortion is small, let's concentrate only on the rst harmonic. What is the
phase shift between input and output signals, expressed both in degrees and weeks?

This phase shift is very important in developing a simple circuit model for the system that describes how
daily temperatures are related to the number of daylight hours. For example, because there is a phase shift,
we know that a simple gain (a resistive circuit) won't work.

(e) Find the transfer function of the simplest possible linear model (a rst-order lter, either lowpass or
highpass) that would describe the data. Note that the phase shift is negative: the output lags (occurs
later) than the input. Based on this observation, which rst-order model ts the data and what is its
transfer function?
(f) Characterize and interpret the structure of this model. Here, let the input (length of day) be a voltage
source and the output (daily high) a voltage. Give a physical explanation for the phase shift.

Problem 4.6: Fourier Transform Pairs
Find the Fourier or inverse Fourier transform of the following.

(a) x ( t ) = e (b) x ( t ) = te u( t )

(c) X ( f ) =

(
1 j f j < W

0 j f j > W
(d) x ( t ) = e cos (2 f t ) u( t )

Problem 4.7: Duality in Fourier Transforms
Duality means that the Fourier transform and the inverse Fourier transform are very similar. Consequently,
the waveform s ( t ) in the time domain and the spectrum s ( f ) have a Fourier transform and an inverse Fourier
transform, respectively, that are very similar.

(a) Calculate the Fourier transform of the signal shown below (Figure 4.23(a)).
(b) Calculate the inverse Fourier transform of the spectrum shown below (Figure 4.23(b)).
(c) How are these answers related? What is the general relationship between the Fourier transform of s ( t )
and the inverse transform of s ( f ) ?

(a)

(b)

Figure 4.23

Problem 4.8: Spectra of Pulse Sequences
Pulse sequences occur often in digital communication and in other elds as well. What are their spectral
properties?

(a) Calculate the Fourier transform of the single pulse shown below (Figure 4.24(a)).
