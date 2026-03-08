180 CHAPTER 5. DIGITAL SIGNAL PROCESSING

### 5.17 Digital Signal Processing Problems

Problem 5.1: Sampling and Filtering
The signal s ( t ) is bandlimited to 4 kHz. We want to sample it, but it has been subjected to various signal
processing manipulations.

(a) What sampling frequency (if any works) can be used to sample the result of passing s ( t ) through an
RC highpass lter with R = 10k and C = 8nF ?
(b) What sampling frequency (if any works) can be used to sample the derivative of s ( t ) ?
(c) The signal s ( t ) has been modulated by an 8 kHz sinusoid having an unknown phase: the resulting
signal is s ( t ) sin (2 f t + ) , with f = 8 kHz and =? Can the modulated signal be sampled so that
the original signal can be recovered from the modulated signal regardless of the phase value ? If so,
show how and nd the smallest sampling rate that can be used; if not, show why not.

Problem 5.2: Non-Standard Sampling
Using the properties of the Fourier series can ease nding a signal's spectrum.

(a) Suppose a signal s ( t ) is periodic with period T . If c represents the signal's Fourier series coe cients,
what are the Fourier series coe cients of s t
?
(b) Find the Fourier series of the signal p ( t ) shown in Figure 5.25.

(c) Suppose this signal is used to sample a signal bandlimited to
1
Hz . Find an expression for and sketch

the spectrum of the sampled signal.
(d) Does aliasing occur? If so, can a change in sampling rate prevent aliasing; if not, show how the signal
can be recovered from these samples.

Figure 5.25

Problem 5.3: A Di erent Sampling Scheme
A signal processing engineer from Texas A&M claims to have developed an improved sampling scheme. He
multiplies the bandlimited signal by the depicted periodic pulse signal to perform sampling (Figure 5.26).
This content is available online at http://cnx.org/content/m10351/2.38/ .
