247

(d) In the digital case, the recovered speech signal can be considered to have two noise sources added
to each sample's true value: One is the A/D amplitude quantization noise and the second is due to
channel errors. Because these are separate, the total noise power equals the sum of these two. What
is the signal-to-noise ratio of the received speech signal as a function of p ?
(e) Compute and plot the received signal's signal-to-noise ratio for the two transmission schemes for a few
values of channel signal-to-noise ratios.
(f) Compare and evaluate these systems.

Problem 6.16: Source Compression
Consider the following 5-letter source. : 5 : 25 : 125 : 0625 : 0625

(a) Find this source's entropy.
(b) Show that the simple binary coding is ine cient.
(c) Find an unequal-length codebook for this sequence that satis es the Source Coding Theorem. Does
your code achieve the entropy limit?
(d) How much more e cient is this code than the simple binary code?

Problem 6.17: Source Compression
Consider the following 5-letter source. : 4 : 2 : 15 : 15 : 1

(a) Find this source's entropy.
(b) Show that the simple binary coding is ine cient.
(c) Find the Hu man code for this source. What is its average code length?

Problem 6.18: Speech Compression
When we sample a signal, such as speech, we quantize the signal's amplitude to a set of integers. For a b -bit
converter, signal amplitudes are represented by 2 integers. Although these integers could be represented by
a binary code for digital transmission, we should consider whether a Hu man coding would be more e cient.

(a) Load into Matlab the segment of speech contained in y.mat . Its sampled values lie in the interval
(-1, 1). To simulate a 3-bit converter, we use Matlab's round function to create quantized amplitudes
corresponding to the integers [0 1 2 3 4 5 6 7] .
