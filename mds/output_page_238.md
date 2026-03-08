232 CHAPTER 6. INFORMATION COMMUNICATION

Figure 6.24: The capacity per transmission through a binary symmetric channel is plotted as a function
of the digital channel's error probability (upper) and as a function of the signal-to-noise ratio for a BPSK
signal set (lower).
troduces errors) and receive the information without error if the code is su ciently ine cient compared
to the channel's characteristics. Generally, a channel's capacity changes with the signal-to-noise ratio: As
one increases or decreases, so does the other. The capacity measures the overall error characteristics of a
channel the smaller the capacity the more frequently errors occur and an overly e cient error-correcting
code will not build in enough error correction capability to counteract channel errors.
This result astounded communication engineers when Shannon published it in 1948. Analog communi-
cation always yields a noisy version of the transmitted signal; in digital communication, error correction can
be powerful enough to correct all errors as the block length increases. The key for this capability to exist is
that the code's e ciency be less than the channel's capacity. For a binary symmetric channel, the capacity
is given by
C = 1 + p log p + (1 p ) log (1 p ) bits = transmission (6.60)

Figure 6.24 (Capacity of a Channel) shows how capacity varies with error probability. For example, our (7,4)
Hamming code has an e ciency of 0 : 57 , and codes having the same e ciency but longer block sizes can be
used on additive noise channels where the signal-to-noise ratio exceeds 0 dB .

### 6.31 Capacity of a Channel

In addition to the Noisy Channel Coding Theorem and its converse, Shannon also derived the capacity for a
bandlimited (to W Hz) additive white noise channel. For this case, the signal set is unrestricted, even to the
point that more than one bit can be transmitted each bit interval. Instead of constraining channel code
e ciency, the revised Noisy Channel Coding Theorem states that some error-correcting code exists such that
as the block length increases, error-free transmission is possible if the source coder's datarate, ( A ) R , is
less than capacity.
C = W log (1 + SNR) bits/s (6.61)

This result sets the maximum datarate of the source coder's output that can be transmitted through the
bandlimited channel with no error. Shannon's proof of his theorem was very clever, and did not indicate
This content is available online at http://cnx.org/content/m0098/2.13/ .
The bandwidth restriction arises not so much from channel properties, but from spectral regulation, especially for wireless
channels.

![Im259](../output/images/test_238_Im259.R10.tif)
