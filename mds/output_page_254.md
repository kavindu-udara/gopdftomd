248 CHAPTER 6. INFORMATION COMMUNICATION

y_quant = round(3.5*y + 3.5);

Find the relative frequency of occurrence of quantized amplitude values. The following Matlab program
computes the number of times each quantized value occurs.

for n=0:7; count(n+1) = sum(y_quant == n); end;

Find the entropy of this source.
(b) Find the Hu man code for this source. How would you characterize this source code in words?
(c) How many fewer bits would be used in transmitting this speech segment with your Hu man code in
comparison to simple binary coding?

Problem 6.19: Digital Communication
In a digital cellular system, a signal bandlimited to 5 kHz is sampled with a two-bit A/D converter at its
Nyquist frequency. The sample values are found to have the shown relative frequencies.

We send the bit stream consisting of Hu man-coded samples using one of the two depicted signal sets
(Figure 6.39).

Figure 6.39

(a) What is the datarate of the compressed source?
(b) Which choice of signal set maximizes the communication system's performance?
(c) With no error-correcting coding, what signal-to-noise ratio would be needed for your chosen signal set
to guarantee that the bit error probability will not exceed 10 ? If the receiver moves twice as far
from the transmitter (relative to the distance at which the 10 error rate was obtained), how does
the performance change?

Problem 6.20: Signal Compression
Letters drawn from a four-symbol alphabet have the indicated probabilities.
