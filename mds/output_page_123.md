117

Table 4.1 s ( t ) + a s ( t ) S ( f ) + a S ( f ) ( t ) 2 R ( f ) = S ( f )
( t ) = s ( t ) ( f ) = S ( f ) ( t ) = s ( t ) ( f ) = S ( f ) ( at ) a j
S
f

( t ) S ( f ) s ( t ) ( f f ) ( t ) cos (2 f t ) ( f f ) + S ( f + f ) ( t ) sin (2 f t ) ( f f ) S ( f + f ) j
s ( t ) 2 fS ( f )

s ( ) d 2 f
S ( f ) if S (0) = 0 t ( t )
1 2

d
S ( f )

s ( t ) dt (0) (0)

S ( f ) df

j s ( t ) j dt

j S ( f ) j df

Example 4.5
In communications, a very important operation on a signal s ( t ) is to amplitude modulate it.
Using this operation more as an example rather than elaborating the communications aspects here,
we want to compute the Fourier transform the spectrum of

1 + s ( t ) cos (2 f t )

Thus,
1 + s ( t ) cos (2 f t ) = cos (2 f t ) + s ( t ) cos (2 f t )

For the spectrum of cos (2 f t ) , we use the Fourier series. Its period is 1 =f , and its only nonzero
Fourier coe cients are c =
. The second term is not periodic unless s ( t ) has the same period
as the sinusoid. Using Euler's relation, the spectrum of the second term can be derived as

s ( t ) cos (2 f t ) =

Z

S ( f ) e df cos (2 f t )
