251

a (7,4) Hamming code having the generator matrix

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

1 0 1 1

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

This code corrects all single-bit error, but if a double bit error occurs, it corrects using a single-bit error
correction approach.

(a) How many double-bit errors can occur in a codeword?
(b) For each double-bit error pattern, what is the result of channel decoding? Express your result as a
binary error sequence for the data bits.

Problem 6.26: Selective Error Correction
We have found that digital transmission errors occur with a probability that remains constant no matter
how important the bit may be. For example, in transmitting digitized signals, errors occur as frequently
for the most signi cant bit as they do for the least signi cant bit. Yet, the former errors have a much larger
impact on the overall signal-to-noise ratio than the latter. Rather than applying error correction to each
sample value, why not concentrate the error correction on the most important bits? Assume that we sample
an 8 kHz signal with an 8-bit A/D converter. We use single-bit error correction on the most signi cant four
bits and none on the least signi cant four. Bits are transmitted using a modulated BPSK signal set over an
additive white noise channel.

(a) How many error correction bits must be added to provide single-bit error correction on the most
signi cant bits?
(b) How large must the signal-to-noise ratio of the received signal be to insure reliable communication?
(c) Assume that once error correction is applied, only the least signi cant 4 bits can be received in error.
How much would the output signal-to-noise ratio improve using this error correction scheme?

Problem 6.27: Compact Disk
Errors occur in reading audio compact disks. Very few errors are due to noise in the compact disk player;
most occur because of dust and scratches on the disk surface. Because scratches span several bits, a single-
bit error is rare; several consecutive bits in error are much more common. Assume that scratch and
dust-induced errors are four or fewer consecutive bits long. The audio CD standard requires 16-bit, 44.1 kHz
analog-to-digital conversion of each channel of the stereo analog signal.

(a) How many error-correction bits are required to correct scratch-induced errors for each 16-bit sample?
(b) Rather than use a code that can correct several errors in a codeword, a clever ELEC 241 engineer
proposes interleaving consecutive coded samples. As the cartoon (Figure 6.41) shows, the bits repre-
senting coded samples are interspersed before they are written on the CD. The CD player de-interleaves
the coded data, then performs error-correction. Now, evaluate this proposed scheme with respect to
the non-interleaved one.
