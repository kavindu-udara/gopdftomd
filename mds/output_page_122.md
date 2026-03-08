116 CHAPTER 4. FREQUENCY DOMAIN

Properties of the Fourier transform and some useful transform pairs are provided in the accompanying tables
(Table 4.1: Short Table of Fourier Transform Pairs and Table 4.2: Fourier Transform Properties). Especially
important among these properties is Parseval's Theorem , which states that power computed in either
domain equals the power in the other.

Z

s ( t ) dt =

Z

j S ( f ) j df (4.36)

Of practical importance is the conjugate symmetry property: When s ( t ) is real-valued, the spectrum at
negative frequencies equals the complex conjugate of the spectrum at the corresponding positive frequencies.
Consequently, we need only plot the positive frequency portion of the spectrum (we can easily determine the
remainder of the spectrum).

Exercise 4.13 (Solution on p. 140.)
How many Fourier transform operations need to be applied to get the original signal back:
F ( ( F ( s ))) = s ( t ) ?

Note that the mathematical relationships between the time domain and frequency domain versions of the
same signal are termed transforms . We are transforming (in the nontechnical meaning of the word) a
signal from one representation to another. We express Fourier transform pairs as ( s ( t ) ! S ( f )) . A
signal's time and frequency domain representations are uniquely related to each other. A signal thus exists
in both the time and frequency domains, with the Fourier transform bridging between the two. We can
de ne an information carrying signal in either the time or frequency domains; it behooves the wise engineer
to use the simpler of the two.
A common misunderstanding is that while a signal exists in both the time and frequency domains, a single
formula expressing a signal must contain only time or frequency: Both cannot be present simultaneously.
This situation mirrors what happens with complex amplitudes in circuits: As we reveal how communications
systems work and are designed, we will de ne signals entirely in the frequency domain without explicitly nd-
ing their time domain variants. This idea is shown in Section 4.6, where we de ne Fourier series coe cients
according to letter to be transmitted. Thus, a signal, though most familiarly de ned in the time-domain,
really can be de ned equally as well (and sometimes more easily) in the frequency domain. For example,
impedances depend on frequency and the time variable cannot appear.
We will learn that nding a linear, time-invariant system's output in the time domain can be most easily
calculated by determining the input signal's spectrum, performing a simple calculation in the frequency
domain, and inverse transforming the result. Furthermore, understanding communications and information
processing systems requires a thorough understanding of signal structure and of how systems work in both
the time and frequency domains.
The only di culty in calculating the Fourier transform of any signal occurs when we have periodic signals
(in either domain). Realizing that the Fourier series is a special case of the Fourier transform, we simply
calculate the Fourier series coe cients instead, and plot them along with the spectra of non-periodic signals
on the same frequency axis. ( t ) ( f ) u( t ) 2 f + a
a f + a
( t ) =

(
1 ; j t j <
0 ; j t j >
f ) W t ) ( f ) =

(
1 ; j f j < W

0 ; j f j > W
