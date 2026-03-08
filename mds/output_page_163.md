157

The ratio of sine functions has the generic form of
sin ( Nx ) x )
, which is known as the discrete-time sinc

function dsinc ( x ) . Thus, our transform can be concisely expressed as S e = e dsinc ( f ) .
The discrete-time pulse's spectrum contains many ripples, the number of which increase with N , the pulse's
duration.

Figure 5.11: The spectrum of a length-ten pulse is shown. Can you explain the rather complicated
appearance of the phase?

The inverse discrete-time Fourier transform is easily derived from the following relationship:

Z
e e df =

(
1 ; if m = n

0 ; if m 6 = n

= ( m n ) (5.28)

Therefore, we nd that

Z
S e e df =

Z
X
s ( m ) e e df

=
X
s ( m )

Z
e df

= s ( n )

(5.29)

The Fourier transform pairs in discrete-time are

S e =
X
s ( n ) e

s ( n ) =

Z
S e e df

(5.30)
