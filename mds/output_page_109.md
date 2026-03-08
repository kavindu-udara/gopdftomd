103

### 4.3 Classic Fourier Series

The classic Fourier series as derived originally by Euler and Fourier expressed a periodic signal (period T )
in terms of harmonically related sines and cosines.

s ( t ) = a +
X
a cos
2 kt
+
X
b sin
2 kt
(4.11)

The Fourier series and the sine-cosine series are di erent versions of the same idea , each
representing a signal's spectrum. The classic Fourier coe cients, a and b , express the real and imaginary
parts respectively of the spectrum while the coe cients c of the Fourier series express the spectrum as a
magnitude and phase. Equating the classic Fourier series (4.11) to the Fourier series (4.1), an extra factor
of two and a complex conjugate become necessary to relate the Fourier coe cients in each.

c =

(
a k = 0
( a jb ) k 6 = 0

Exercise 4.3 (Solution on p. 139.)
Derive this relationship between the coe cients of the two Fourier series.

Just as with the Fourier series, we can nd the sine-cosine Fourier coe cients using the orthogonality
properties of sinusoids. Note that the cosine and sine of harmonically related frequencies, even the same
frequency, are orthogonal.

Z

sin
2 kt
cos
2 lt
dt = 0 ; k 2 Z ; l 2 Z

Z

sin
2 kt
sin
2 lt
dt =

8
<

:

T
; if k = l and k 6 = 0 and l 6 = 0

0 ; k 6 = l or k = 0 = l

Z

cos
2 kt
cos
2 lt
dt =

8
>
>
<

>
>
:

T
; k = l and k 6 = 0 and l 6 = 0

T; k = 0 = l

0 ; k 6 = l

(4.12)

These orthogonality relations follow from the following important trigonometric identities.

sin ( ) sin ( ) =
1
cos ( ) cos ( + )

cos ( ) cos ( ) =
1
cos ( + ) + cos ( )

sin ( ) cos ( ) =
1
sin ( + ) + sin ( )

(4.13)

These identities allow you to substitute a sum of sines and/or cosines for a product of them. Each term in
the sum can be integrated by noticing one of two important properties of sinusoids.

The integral of a sinusoid over an integer number of periods equals zero.
The integral of the square of a unit-amplitude sinusoid over a period T equals T= 2 .

To use these, let's, for example, multiply the Fourier series for a signal by the cosine of the l harmonic

cos
2 lt
and integrate. The idea is that, because integration is linear, the integration will sift out all but
This content is available online at http://cnx.org/content/m0039/2.22/ .
