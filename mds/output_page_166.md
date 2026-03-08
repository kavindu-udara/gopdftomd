160 CHAPTER 5. DIGITAL SIGNAL PROCESSING

We thus de ne the discrete Fourier transform (DFT) to be

S ( k ) =
X
s ( n ) e
; k 2 f 0 ; : : : ; K 1 g (5.33)

Here, S ( k ) is shorthand for S e
.

By choosing the value of K , we can compute the spectrum at as many equally spaced frequencies as
we like. Note that you can think about this computationally motivated choice as sampling the spectrum;
more about this interpretation later. The issue now is how many frequencies are enough to capture how
the spectrum changes with frequency. One way of answering this question is determining an inverse discrete
Fourier transform formula: given S ( k ) , k = f 0 ; : : : ; K 1 g how do we nd s ( n ) , n = f 0 ; : : : ; N 1 g ?

Presumably, the formula will be of the form s ( n ) =
X
S ( k ) e
. Substituting the DFT formula in this

prototype inverse transform yields

s ( n ) =
X X
s ( m ) e

!

e
(5.34)

Note that the orthogonality relation we use so often has a di erent character now.

X
e
e
=

(
K; if m = f n; ( n K ) ; ( n 2 K ) ; : : : g

0 ; otherwise
(5.35)

We obtain nonzero value whenever the two indices di er by multiples of K . We can express this result as
K
P
( ( m n lK )) . Thus, our formula becomes

s ( n ) =
X
s ( m ) K
X
( m n lK ) (5.36)

The integers n and m both range over f 0 ; : : : ; N 1 g . To have an inverse transform, we need the sum to be a
single unit sample for m , n in this range. If it did not, then s ( n ) would equal a sum of values, and we would
not have a valid transform: Once going into the frequency domain, we could not get back unambiguously!
Clearly, the term l = 0 always provides a unit sample (we'll take care of the factor of K soon). If we evaluate
the spectrum at fewer frequencies than the signal's duration, the term corresponding to m = n + K will
also appear for some values of m , n = f 0 ; : : : ; N 1 g . This situation means that our prototype transform
equals s ( n ) + s ( n + K ) for some values of n . The only way to eliminate this problem is to require K N :
We must have at least as many frequency samples as the signal's duration. In this way, we can return from
the frequency domain we entered via the DFT.

Exercise 5.15 (Solution on p. 192.)
When we have fewer frequency samples than the signal's duration, some discrete-time signal values
equal the sum of the original signal values. Given the sampling interpretation of the spectrum,
characterize this e ect a di erent way.

Another way to understand this requirement is to use the theory of linear equations. If we write out the
expression for the DFT as a set of linear equations,

s (0) + s (1) + + s ( N 1) = S (0)

s (0) + s (1) e
+ + s ( N 1) e
= S (1)

.
.
.

s (0) + s (1) e
+ + s ( N 1) e
= S ( K 1)

(5.37)
