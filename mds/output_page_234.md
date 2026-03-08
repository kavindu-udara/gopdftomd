228 CHAPTER 6. INFORMATION COMMUNICATION

where the generator matrix is

G =

2

6
6
6
6
6
6
6
6
6
6
6
6
6
4

1 0 0 0

0 1 0 0

0 0 1 0

0 0 0 1

1 1 1 0

0 1 1 1

1 1 0 1

3

7
7
7
7
7
7
7
7
7
7
7
7
7
5

In this (7,4) code, 2 = 16 of the 2 = 128 possible blocks at the channel decoder correspond to error-free
transmission and reception.
Error correction amounts to searching for the codeword c closest to the received block b c in terms of the
Hamming distance between the two. The error correction capability of a channel code is limited by how close
together any two error-free blocks are. Bad codes would produce blocks close together, which would result
in ambiguity when assigning a block of data bits to a received block. The quantity to examine, therefore, in
designing code error correction codes is the minimum distance between codewords.

d = min ( d ( c ; c )) ; c 6 = c (6.56)

To have a channel code that can correct all single-bit errors, d 3 .

Exercise 6.29 (Solution on p. 258.)
Suppose we want a channel code to have an error-correction capability of n bits. What must the
minimum Hamming distance between codewords d be?

How do we calculate the minimum distance between codewords? Because we have 2 codewords, the number
of possible unique pairs equals 2 2 1 , which can be a large number. Recall that our channel coding
procedure is linear, with c = G b . Therefore ( c c ) = G (( b b )) . Because b b always yields another
block of data bits, we nd that the di erence between any two codewords is another codeword! Thus, to
nd d we need only compute the number of ones that comprise all non-zero codewords. Finding these
codewords is easy once we examine the coder's generator matrix. Note that the columns of G are codewords
(why is this?), and that all codewords can be found by all possible pairwise sums of the columns. To nd
d , we need only count the number of bits in each column and sums of columns. For our example (7 ; 4) ,
G 's rst column has three ones, the next one four, and the last two three. Considering sums of column pairs
next, note that because the upper portion of G is an identity matrix, the corresponding upper portion of all
column sums must have exactly two bits. Because the bottom portion of each column di ers from the other
columns in at least one place, the bottom portion of a sum of columns must have at least one bit. Triple
sums will have at least three bits because the upper portion of G is an identity matrix. Thus, no sum of
columns has fewer than three bits, which means that d = 3 , and we have a channel coder that can correct
all occurrences of one error within a received 7 -bit block.

### 6.28 Error-Correcting Codes: Channel Decoding

Because the idea of channel coding has merit (so long as the code is e cient), let's develop a systematic
procedure for performing channel decoding. One way of checking for errors is to try recreating the error
correction bits from the data portion of the received block b c . Using matrix notation, we make this calculation
by multiplying the received block b c by the matrix H known as the parity check matrix . It is formed from
the generator matrix G by taking the bottom, error-correction portion of G and attaching to it an identity
This content is available online at http://cnx.org/content/m0072/2.20/ .
