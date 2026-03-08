119

### 4.9 Linear Time Invariant Systems

When we apply a periodic input to a linear, time-invariant system, the output is periodic and has Fourier
series coe cients equal to the product of the system's frequency response and the input's Fourier coe cients
(Filtering Periodic Signals (4.27)). The way we derived the spectrum of non-periodic signal from periodic
ones makes it clear that the same kind of result works when the input is not periodic: If x ( t ) serves as
the input to a linear, time-invariant system having frequency response H ( f ) , the spectrum of
the output is X ( f ) H ( f ) .

Example 4.6
Let's use this frequency-domain input-output relationship for linear, time-invariant systems to nd
a formula for the RC -circuit's response to a pulse input. We have expressions for the input's
spectrum and the system's frequency response.

P ( f ) = e
sin ( f )
(4.38)

H ( f ) =
1 j 2 fRC
(4.39)

Thus, the output's Fourier transform equals

Y ( f ) = e
sin ( f )

1 j 2 fRC
(4.40)

You won't nd this Fourier transform in our table, and the required integral is di cult to evaluate
as the expression stands. This situation requires cleverness and an understanding of the Fourier
transform's properties. In particular, recall Euler's relation for the sinusoidal term and note the
fact that multiplication by a complex exponential in the frequency domain amounts to a time delay.
Let's momentarily make the expression for Y ( f ) .

e
sin ( f )
= e
e e
2 f

=
1 2 f
1 e

(4.41)

Consequently,

Y ( f ) =
1 2 f
1 e
1 j 2 fRC
(4.42)

The table of Fourier transform properties (Table 4.2) suggests thinking about this expression as a
product of terms.

Multiplication by
1 2 f
means integration.

Multiplication by the complex exponential e means delay by seconds in the time
domain.
The term 1 e means, in the time domain, subtract the time-delayed signal from its
original.

The inverse transform of the frequency response is
1
e u( t ) .

We can translate each of these frequency-domain products into time-domain operations in any
order we like because the order in which multiplications occur doesn't a ect the result. Let's
start with the product of 1 =j 2 f (integration in the time domain) and the transfer function:

1 2 f

1 j 2 fRC
! 1 e u( t ) (4.43)
This content is available online at http://cnx.org/content/m0048/2.18/ .
