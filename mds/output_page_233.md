227

Figure 6.22: In a (3,1) repetition code, only two of the possible eight three-bit data blocks are codewords.
We can represent these bit patterns geometrically with the axes being bit positions in the data block. In
the left plot, the lled circles represent the codewords [ 0 0 0 ] and [ 1 1 1 ], the only possible codewords.
The un lled ones correspond to the transmission. The center plot shows that the distance between
codewords is three. Because distance corresponds to ipping a bit, calculating the Hamming distance
geometrically means following the axes rather than going as the crow ies. The right plot shows the
data words that result when one error occurs as the codeword goes through the channel. The three data
words are unit distance from the original codeword. Note that the received data word groups do not
overlap, which means the code can correct all single-bit errors. Np (1 p ) . The number of errors
the channel introduces equals the number of ones in e ; the probability of any particular error vector decreases
with the number of errors.
To perform decoding when errors occur, we want to nd the codeword (one of the lled circles in Fig-
ure 6.22) that has the highest probability of occurring: the one closest to the one received. Note that if a data
word lies a distance of one from two codewords, it is impossible to determine which codeword was actually
sent. This criterion means that if any two codewords are two bits apart, then the code cannot correct the
channel-induced error. Thus, to have a code that can correct all single-bit errors, codewords must
have a minimum separation of three. Our repetition code has this property.
Introducing code bits increases the probability that any bit arrives in error (because bit interval durations
decrease). However, using a well-designed error-correcting code corrects bit reception errors. Do we win or
lose by using an error-correcting code? The answer is that we can win if the code is well-designed. The (3,1)
repetition code demonstrates that we can lose (Exercise 6.28). To develop good channel coding, we need to
develop rst a general framework for channel codes and discover what it takes for a code to be maximally
e cient: Correct as many errors as possible using the fewest error correction bits as possible (making the
e ciency K=N as large as possible.) We also need a systematic way of nding the codeword closest to any
received data word. A much better code than our (3,1) repetition code is the following (7,4) code.

c (1) = b (1)

c (2) = b (2)

c (3) = b (3)

c (4) = b (4)

c (5) = b (1) b (2) b (3)

c (6) = b (2) b (3) b (4)

c (7) = b (1) b (2) b (4)
This content is available online at http://cnx.org/content/m10283/2.29/ .
