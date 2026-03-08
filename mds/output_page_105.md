99

Thus, the Fourier series for the square wave is

sq ( t ) =
X
2
e
(4.6)

Consequently, the square wave equals a sum of complex exponentials, but only those having fre-
quencies equal to odd multiples of the fundamental frequency
. The coe cients decay slowly as
the frequency index k increases. This index corresponds to the k -th harmonic of the signal's period.

4.2.1 Fourier coe cients properties

A signal's Fourier series spectrum c has interesting properties.

Property 4.1:
If s ( t ) is real, c = c (real-valued periodic signals have conjugate-symmetric spectra).

This result follows from the integral that calculates the c from the signal. Furthermore, this result means
that Re [ c ] = Re [ c ] : The real part of the Fourier coe cients for real-valued signals is even. Similarly,
Im [ c ] = Im [ c ] : The imaginary parts of the Fourier coe cients have odd symmetry. Consequently, if
you are given the Fourier coe cients for positive indices and zero and are told the signal is real-valued, you
can nd the negative-indexed coe cients, hence the entire spectrum. This kind of symmetry, c = c , is
known as conjugate symmetry .

Property 4.2:
If s ( t ) = s ( t ) , which says the signal has even symmetry about the origin, c = c .

Given the previous property for real-valued signals, the Fourier coe cients of even signals are real-valued.
A real-valued Fourier expansion amounts to an expansion in terms of only cosines, which is the simplest
example of an even signal.

Property 4.3:
If s ( t ) = s ( t ) , which says the signal has odd symmetry, c = c .

Therefore, the Fourier coe cients are purely imaginary. The square wave is a great example of an odd-
symmetric signal.

Property 4.4:

The spectral coe cients for a periodic signal delayed by , s ( t ) , are c e
, where c denotes
the spectrum of s ( t ) . Delaying a signal by seconds results in a spectrum having a linear phase

shift of
2 k
in comparison to the spectrum of the un-delayed signal. Note that the spectral

magnitude is una ected. Showing this property is easy.

1

Z

s ( t ) e
dt =
1

Z

s ( t ) e
dt

=
1
e

Z

s ( t ) e
dt

(4.7)

Note that the range of integration extends over a period of the integrand. Consequently, it should

not matter how we integrate over a period, which means that
R
( ) dt =
R
( ) dt , and we have
our result.

By the way, you can easily show that no matter how you integrate a periodic signal over a period, you get
the same answer. Let x ( t ) be period with period T .

Z

x ( t ) dt =

Z

x ( t ) dt +

Z

x ( t ) dt
