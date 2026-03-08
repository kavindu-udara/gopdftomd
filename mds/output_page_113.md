107

Exercise 4.7 (Solution on p. 139.)
In high-end audio, deviation of a sine wave from the ideal is measured by the total harmonic
distortion , which equals the total power in the harmonics higher than the rst compared to the
signal's total power. Find an expression for the total harmonic distortion for any periodic signal.
Is this calculation most easily performed in the time or frequency domain?

### 4.5 Fourier Series Approximation of Signals

It is interesting to consider the sequence of signals that we obtain as we incorporate more terms into the
Fourier series approximation of the half-wave recti ed sine wave (Example 4.2). De ne s ( t ) to be the
signal containing only those terms up to and including the K harmonic.

s ( t ) =
X
a cos
2 kt
+
X
b sin
2 kt

=
X
c e

(4.23)

Figure 4.5 shows how this sequence of signals portrays the signal more accurately as more terms are added.
We need to assess quantitatively the accuracy of the Fourier series approximation so that we can judge
how rapidly the series approaches the signal. When we use a truncated series, the error the di erence
between the signal and the K -harmonic series corresponds to the unused terms from the series.

( t ) =
X
a cos
2 kt
+
X
b sin
2 kt

=
X
c e

(4.24)

To nd the rms error, we must square this expression and integrate it over a period. Again, the integral of
most cross-terms is zero, leaving

rms ( ) =
1

X
a + b

= 2
X
j c j

(4.25)

Figure 4.6 shows how the error in the Fourier series for the half-wave recti ed sinusoid decreases as more
terms are incorporated. In particular, the use of four terms, as shown in the bottom plot of Figure 4.5, has
a rms error (relative to the rms value of the signal) of about 3%. The Fourier series in this case converges
quickly to the signal.
We can look at Figure 4.7 to see the power spectrum and the rms approximation error for the square
wave. Because the Fourier coe cients decay more slowly here than for the half-wave recti ed sinusoid, the
rms error is not decreasing quickly. Said another way, the square-wave's spectrum contains more power at
higher frequencies than does the half-wave-recti ed sinusoid. This di erence between the two Fourier series

results because the half-wave recti ed sinusoid's Fourier coe cients are proportional to
1
while those of the

square wave are proportional to
1
. If fact, after 99 terms of the square wave's approximation, the error is

bigger than 10 terms of the approximation for the half-wave recti ed sinusoid. Mathematicians have shown
that no signal has an rms approximation error that decays more slowly than it does for the square wave.
This content is available online at http://cnx.org/content/m10687/2.9/ .
