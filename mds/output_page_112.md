106 CHAPTER 4. FREQUENCY DOMAIN

Figure 4.4: Power spectrum of a half-wave recti ed sinusoid. uniqueness : You can unambiguously
nd the spectrum from the signal (decomposition (4.15)) and the signal from the spectrum (composition).
Thus, any aspect of the signal can be found from the spectrum and vice versa. A signal's frequency
domain expression is its spectrum . A periodic signal can be de ned either in the time domain (as a
function) or in the frequency domain (as a spectrum).
A fundamental aspect of solving electrical engineering problems is whether the time or frequency domain
provides the most understanding of a signal's properties and the simplest way of manipulating it. The
uniqueness property says that either domain can provide the right answer. As a simple example, suppose
we want to know the (periodic) signal's maximum value. Clearly the time domain provides the answer
directly. To use a frequency domain approach would require us to nd the spectrum, form the signal from
the spectrum and calculate the maximum; we're back in the time domain!
Another feature of a signal is its average power . A signal's instantaneous power is de ned to be its
square. The average power is the average of the instantaneous power over some time interval. For a periodic
signal, the natural time interval is clearly its period; for non-periodic signals, a better choice would be entire
time or time from onset. For a periodic signal, the average power is the square of its root-mean-squared
(rms) value. We de ne the rms value of a periodic signal to be

rms ( s ) =

s

Z

s ( t ) dt (4.20)

and thus its average power is

power ( s ) = rms ( s ) =
1

Z

s ( t ) dt (4.21)

Exercise 4.6 (Solution on p. 139.)
What is the rms value of the half-wave recti ed sinusoid?

To nd the average power in the frequency domain, we only need Parseval's Theorem.

1

Z

s ( t ) dt =
X
j c j = a +
1

X
a + b (4.22)

It could well be that computing the sum is easier than integrating the signal's square. Furthermore,
the contribution of each term in the Fourier series toward representing the signal can be measured by its
contribution to the signal's average power. Thus, the power contained in a signal at its k th harmonic is
a + b
or j c j . The power spectrum , P ( k ) , such as shown in Figure 4.4, plots each harmonic's

contribution to the total power.
