257

Solution to Exercise 6.21 (p. 218)

Equally likely symbols each have a probability of
1
. Thus, H ( A ) =
P 1
log
1
= log K . To prove

that this is the maximum-entropy probability assignment, we must explicitly take into account that prob-
abilities sum to one. Focus on a particular symbol, say the rst. Pr [ a ] appears twice in the entropy for-
mula: the terms Pr [ a ] log (Pr [ a ]) and (1 Pr [ a ] + + Pr [ a ]) log (1 Pr [ a ] + + Pr [ a ]) .
The derivative with respect to this probability (and all the others) must be zero. The derivative equals
log (Pr [ a ]) log (1 Pr [ a ] + + Pr [ a ]) , and all other derivatives have the same form (just substi-
tute your letter's index). Thus, each probability must equal the others, and we are done. For the minimum
entropy answer, one term is 1log 1 = 0 , and the others are 0log 0 , which we de ne to be zero also. The
minimum value of entropy is zero.
Solution to Exercise 6.22 (p. 221)
The Hu man coding tree for the second set of probabilities is identical to that for the rst (Figure 6.18).
The average code length is
1 +
2 +
3 +
3 = 1 : 75 bits. The entropy calculation is straightforward:
H ( A ) =
log
+
log
+
log
+
log
, which equals 1 : 68 bits.
Solution to Exercise 6.23 (p. 222)

T =
1 ( A ) R
.

Solution to Exercise 6.24 (p. 222)
Because no codeword begins with another's codeword, the rst codeword encountered in a bit stream must
be the right one. Note that we must start at the beginning of the bit stream; jumping into the middle does
not guarantee perfect decoding. The end of one codeword and the beginning of another could be a codeword,
and we would get lost.
Solution to Exercise 6.25 (p. 222)
Consider the bitstream : : : 0110111 : : : taken from the bitstream 0|10|110|110|111| : : : . We would decode the
initial part incorrectly, then would synchronize. If we had a xed-length code (say 00,01,10,11), the situation
is much worse. Jumping into the middle leads to no synchronization at all!
Solution to Exercise 6.26 (p. 224)
This question is equivalent to 3 p (1 p )+ p 1 or 2 p +( 3) p +1 0 . Because this is an upward-going
parabola, we need only check where its roots are. Using the quadratic formula, we nd that they are located
at
and 1 . Consequently in the range 0 p
the error rate produced by coding is smaller.
Solution to Exercise 6.27 (p. 225)
With no coding, the average bit-error probability p is given by the probability of error equation (6.47):

p = Q

s E

!

. With a threefold repetition code, the bit-error probability is given by 3 p (1 p )+ p ,

where p = Q

s E N

!

. Plotting this reveals that the increase in bit-error probability out of the channel

because of the energy reduction is not compensated by the repetition coding.
