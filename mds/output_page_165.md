159 s ( n ) + a s ( n ) S e + a S e
s ( n ) e = S e
( n ) 2 R e = S e
( n ) = s ( n ) e = S e
( n ) = s ( n ) e = S e
( n ) =

8
<

:

s ( n=M ) n = : : : ; M; 0 ; M; 2 M; : : :

0 otherwise e = S e
( nM )
X
S e
( n n ) S e
s ( n ) e
f n ) s ( n ) e + S e
f n ) s ( n ) e S e
j n ( n )
1 2

d
S e s ( n ) e
(0)

S e df
s ( n )

S e df
continuous-time spectral analysis.
The formula for the DTFT (5.17) is a sum, which conceptually can be easily computed save for two
issues.

Signal duration . The sum extends over the signal's duration, which must be nite to compute the
signal's spectrum. It is exceedingly di cult to store an in nite-length signal in any case, so we'll
assume that the signal extends over [0 ; N 1] .
Continuous frequency . Subtler than the signal duration issue is the fact that the frequency variable
is continuous: It may only need to span one period, like
;
or [0 ; 1] , but the DTFT formula as it
stands requires evaluating the spectra at all frequencies within a period. Let's compute the spectrum

at a few frequencies; the most obvious ones are the equally spaced ones f =
k
, k 2 f 0 ; : : : ; K 1 g .
