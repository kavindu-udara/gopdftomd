100 CHAPTER 4. FREQUENCY DOMAIN

Work on the rst integral

=

Z

x ( u T ) du +

Z

x ( t ) dt [ u = t + T ]

=

Z

x ( u ) du +

Z

x ( t ) dt [ x ( u T ) = x ( t )]

Reversing the order of the intergrals and combining them, we get

Z

x ( t ) dt =

Z

x ( t ) dt

Property 4.5:
The Fourier series obeys Parseval's Theorem , one of the most important results in signal anal-
ysis. Parseval's Theorem states that you can calculate a signal's power in either the time domain
or the frequency domain with very similar formulas.

Theorem 4.1: Parseval's Theorem
Average power calculated in the time domain equals the power calculated in the frequency domain.

1

Z

s ( t ) dt =
X
j c j (4.8)

Proof:
Proving Parseval's Theorem is easy when we use the orthogonality property of harmonically related
complex exponentials (equation 4.2). First of all, re-write the square of the signal as a magnitude-
squared and substitute the Fourier series of the signal into the integral.

1

Z

s ( t ) dt =
1

Z
X
c e

!
X
c e

!

dt

=
1

X X
c c

Z

e
e
dt

Using the fact that

Z

e
e
dt = T ( k + l ) , we nd that

1

Z

j s ( t ) j dt =
X
c c

=
X
j c j [ For real x ( t ) ; c = c ]

4.2.2 A signal's spectrum

The Fourier series expression of (4.2) has a very deep interpretation. The time-domain signal s ( t ) is there
expressed as a weighted sum of complex exponentials that di er only in their frequency k=T . In this way,
we have a frequency-domain representation of the signal. The deep part is that either representation can
be used to characterize a signal: we can specify a signal in the time-domain or in the frequency domain.
Because of the Fourier series, we can nd how to represent the signal in the other domain. How the Fourier
Parseval's biography can be found at http://www-history.mcs.st-andrews.ac.uk/Biographies/Parseval.html .
