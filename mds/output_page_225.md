219

Exercise 6.21 (Solution on p. 256.)
Derive the maximum-entropy results, both the numeric aspect (entropy equals log K ) and the
theoretical one (equally likely symbols maximize entropy). Derive the value of the minimum entropy
alphabet.

Example 6.1
A four-symbol alphabet has the following probabilities.

Pr [ a ] =
1
Pr [ a ] =
1
Pr [ a ] =
1
Pr [ a ] =
1

Note that these probabilities sum to one as they should. As
= 2 , log
= 1 . The entropy
of this alphabet equals

H ( A ) =
1
log
1
+
1
log
1
+
1
log
1
+
1
log
1

=
1
( 1) +
1
( 2) +
1
( 3) +
1
( 3)

= 1 : 75 bits

(6.51)

### 6.21 Source Coding Theorem

The signi cance of an alphabet's entropy rests in how we can represent it with a sequence of bits . Bit
sequences form the coin of the realm in digital communications: they are the universal way of representing
symbolic-valued signals. We convert back and forth between symbols to bit-sequences with what is known
as a codebook : a table that associates symbols to bit sequences. In creating this table, we must be able to
assign a unique bit sequence to each symbol so that we can go between symbol and bit sequences without
error.

note: You may be conjuring the notion of hiding information from others when we use the
name codebook for the symbol-to-bit-sequence table. There is no relation to cryptology, which
comprises mathematically provable methods of securing information. The codebook terminology
was developed during the beginnings of information theory just after World War II.

As we shall explore in some detail elsewhere, digital communication is the transmission of symbolic-valued
signals from one place to another. When faced with the problem, for example, of sending a le across the
Internet, we must rst represent each character by a bit sequence. Because we want to send the le quickly,
we want to use as few bits as possible. However, we don't want to use so few bits that the receiver cannot
determine what each character was from the bit sequence. For example, we could use one bit for every
character: File transmission would be fast but useless because the codebook creates errors. Shannon proved
in his monumental work what we call today the Source Coding Theorem . Let B ( a ) denote the number
of bits used to represent the symbol a . The average number of bits ( A ) required to represent the entire

alphabet equals
P
B ( a ) Pr [ a ] . The Source Coding Theorem states that the average number
of bits needed to accurately represent the alphabet need only to satisfy

H ( A ) ( A ) < H ( A ) + 1 (6.52)

Thus, the alphabet's entropy speci es to within one bit how many bits on the average need to be used to
send the alphabet. The smaller an alphabet's entropy, the fewer bits required for digital transmission of les
expressed in that alphabet.
This content is available online at http://cnx.org/content/m0091/2.13/ .
