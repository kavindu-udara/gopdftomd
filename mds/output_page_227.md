221

Form a binary tree to the right of the table. A binary tree always has two branches at each node.
Build the tree by merging the two lowest probability symbols at each level, making the probability of
the node equal to the sum of the merged nodes' probabilities. If more than two nodes/symbols share
the lowest probability at a given level, pick any two; your choice won't a ect ( A ) .
At each node, label each of the emanating branches with a binary number. The bit sequence obtained
from passing from the tree's root to the symbol is its Hu man code.

Example 6.3
The simple four-symbol alphabet used in the Entropy (Example 6.1) and Source Coding (Exam-
ple 6.2) modules has a four-symbol alphabet with the following probabilities,

Pr [ a ] =
1
Pr [ a ] =
1
Pr [ a ] =
1
Pr [ a ] =
1

and an entropy of 1.75 bits (Example 6.1). This alphabet has the Hu man coding tree shown in
Figure 6.18.

Figure 6.18: We form a Hu man code for a four-letter alphabet having the indicated probabilities of
occurrence. The binary tree created by the algorithm extends to the right, with the root node (the one
at which the tree begins) de ning the codewords. The bit sequence obtained by traversing the tree from
the root to the symbol de nes that symbol's binary code.

The code thus obtained is not unique as we could have labeled the branches coming out of
each node di erently. The average number of bits required to represent this alphabet equals
1 : 75 bits, which is the Shannon entropy limit for this source alphabet. If we had the symbolic-
valued signal s ( m ) = f a ; a ; a ; a ; a ; a ; : : : g , our Hu man code would produce the bitstream
b ( n ) = 101100111010 : : : .
If the alphabet probabilities were di erent, clearly a di erent tree, and therefore di erent code,
could well result. Furthermore, we may not be able to achieve the entropy limit. If our symbols
had the probabilities Pr [ a ] =
, Pr [ a ] =
, Pr [ a ] =
, and Pr [ a ] =
, the average number
of bits/symbol resulting from the Hu man coding algorithm would equal 1 : 75 bits. However, the
entropy limit is 1.68 bits. The Hu man code does satisfy the Source Coding Theorem its average
length is within one bit of the alphabet's entropy but you might wonder if a better code existed.
David Hu man showed mathematically that no other code could achieve a shorter average code
than his. We can't do better.

Exercise 6.22 (Solution on p. 257.)
Derive the Hu man code for this second set of probabilities, and verify the claimed average code
length and alphabet entropy.
