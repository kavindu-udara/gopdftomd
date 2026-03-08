250 CHAPTER 6. INFORMATION COMMUNICATION

(a) What is this code's e ciency?
(b) Find the generator matrix G and parity-check matrix H for this code.
(c) Give the decoding table for this code. How many patterns of one, two, and three errors are correctly
decoded?
(d) What is the block error probability (the probability of any number of errors occurring in the decoded
codeword)?

Problem 6.23: Digital Communication
A digital source produces sequences of nine letters with the following probabilities.

(a) Find a Hu man code that compresses this source. How does the resulting code compare with the best
possible code?
(b) A clever engineer proposes the following (6 ; 3) code to correct errors after transmission through a digital
channel.

c = d c = ( d d d )

c = d c = ( d d )

c = d c = d

What is the error correction capability of this code?
(c) The channel's bit error probability is 1/8. What kind of code should be used to transmit data over
this channel?

Problem 6.24: Overly Designed Error Correction Codes
An Aggie engineer wants not only to have codewords for his data, but also to hide the information from Rice
engineers (no fear of the UT engineers). He decides to represent 3-bit data with 6-bit codewords in which
none of the data bits appear explicitly.

c = ( d d ) c = ( d d d )

c = ( d d ) c = ( d d )

c = ( d d ) c = ( d d d )

(a) Find the generator matrix G and parity-check matrix H for this code.
(b) Find a 3 6 matrix that recovers the data bits from the codeword.
(c) What is the error correcting capability of the code?

Problem 6.25: Error Correction?
It is important to realize that when more transmission errors than can be corrected, error correction algo-
rithms believe that a smaller number of errors have occurred and correct accordingly. For example, consider
