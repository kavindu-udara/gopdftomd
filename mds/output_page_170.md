164 CHAPTER 5. DIGITAL SIGNAL PROCESSING

Figure 5.13: The basic computational element of the fast Fourier transform is the butter y. It takes
two complex numbers, represented by a and b , and forms the quantities shown. Each butter y requires
one complex multiplication and two complex additions.

By considering together the computations involving common output frequencies from the two half-length
DFTs, we see that the two complex multiplies are related to each other, and we can reduce our computational
work even further. By further decomposing the length-4 DFTs into two length-2 DFTs and combining their
outputs, we arrive at the diagram summarizing the length-8 fast Fourier transform (Figure 5.12). Although
most of the complex multiplies are quite simple (multiplying by e means negating real and imaginary

parts), let's count those for purposes of evaluating the complexity as full complex multiplies. We have
N
= 4

complex multiplies and N = 8 additions for each stage and log N = 3 stages, making the number of basic

computations
3 N
log N as predicted.

Exercise 5.19 (Solution on p. 192.)
Note that the ordering of the input sequence in the two parts of Figure 5.12 aren't quite the same.
Why not? How is the ordering determined?

Other fast algorithms were discovered, all of which make use of how many common factors the transform
length N has. In number theory, the number of prime factors a given integer has measures how composite
it is. The numbers 16 and 81 are highly composite (equaling 2 and 3 respectively), the number 18 is less so
( 2 3 ), and 17 not at all (it's prime). In over thirty years of Fourier transform algorithm development, the
original Cooley-Tukey algorithm is far and away the most frequently used. It is so computationally e cient
that power-of-two transform lengths are frequently used regardless of what the actual length of the data.

Exercise 5.20 (Solution on p. 192.)
Suppose the length of the signal were 500 ? How would you compute the spectrum of this signal
using the Cooley-Tukey algorithm? What would the length N of the transform be?

### 5.10 Spectrograms

We know how to acquire analog signals for digital processing by A/D conversion (Section 5.4), consisting
of pre- ltering, sampling and amplitude quantization, and to compute spectra of discrete-time signals using
the FFT algorithm (Section 5.9). Let's put these various components together to learn how the spectrogram
shown in Figure 5.14, which is used to analyze speech, is calculated. The speech was sampled at a rate of
11.025 kHz and passed through a 16-bit A/D converter.

Point of interest: Music compact discs (CDs) encode their signals at a sampling rate of
44.1 kHz. We'll learn the rationale for this number later. The 11.025 kHz sampling rate for the
speech is 1/4 of the CD sampling rate, and was the lowest available sampling rate commensurate
with speech signal bandwidths available on my computer.

Exercise 5.21 (Solution on p. 192.)
Looking at Figure 5.14 the signal lasted a little over 1.2 seconds. How long was the sampled signal
(in terms of samples)? What was the datarate during the sampling process in bps (bits per second)?
This content is available online at http://cnx.org/content/m0505/2.19/ .
