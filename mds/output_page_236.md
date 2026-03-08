230 CHAPTER 6. INFORMATION COMMUNICATION The codeword lengths N and data-block lengths K that form a perfect single-bit error correcting code
are shown, along with the coding e ciency found by E = K=N .
( rst, second, and sixth bits). Such an error pattern cannot be detected by our coding strategy, but such
multiple error patterns are very unlikely to occur. Our receiver uses the principle of maximum probability:
An error-free transmission is much more likely than one with three errors if the bit-error probability p is
small enough.

Exercise 6.32 (Solution on p. 258.)
How small must p be so that a single-bit error is more likely to occur than a triple-bit error?

### 6.29 Error-Correcting Codes: Hamming Codes

For the (7,4) example, we have 2 1 = 7 error patterns that can be corrected. We start with single-bit
error patterns, and multiply them by the parity check matrix. If we obtain unique answers, we are done; if
two or more error patterns yield the same result, we can try double-bit error patterns. In our case, single-bit
error patterns give a unique result.
This corresponds to our decoding table: We associate the parity check matrix multiplication result with
the error pattern and add this to the received word. If more than one error occurs (unlikely though it may
be), this error correction strategy usually makes the error worse in the sense that more bits are changed
from what was transmitted.
As with the repetition code, we must question whether our (7,4) code's error correction capability com-
pensates for the increased error probability due to the necessitated reduction in bit energy. Figure 6.23 shows
that if the signal-to-noise ratio is large enough channel coding yields a smaller error probability. Because the
bit stream emerging from the source decoder is segmented into four-bit blocks, the fair way of comparing
coded and uncoded transmission is to compute the probability of block error: the probability that any bit
in a block remains in error despite error correction and regardless of whether the error occurs in the data
or in coding buts. Clearly, our (7,4) channel code does yield smaller error rates, and is worth the additional
systems required to make it work.
Note that our (7,4) code has the length and number of data bits that perfectly ts correcting single bit
errors. This pleasant property arises because the number of error patterns that can be corrected, 2 1 ,
equals the codeword length N . Codes that have 2 1 = N are known as Hamming codes , and the
following table (Table 6.4) provides the parameters of these codes. Hamming codes are the simplest single-bit
error correction codes, and the generator/parity check matrix formalism for channel coding and decoding
works for them.
Unfortunately, for such large blocks, the probability of multiple-bit errors can exceed the number of
single-bit errors unless the channel single-bit error probability p is very small. Consequently, we need to
enhance the code's error correcting capability by adding double as well as single-bit error correction.
This content is available online at http://cnx.org/content/m0097/2.25/ .
