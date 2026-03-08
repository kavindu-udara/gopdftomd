241

Figure 6.31

Problem 6.2: Noise in AM Systems
The signal b s ( t ) emerging from an AM communication system consists of two parts: the message signal, s ( t ) ,
and additive noise. The plot (Figure 6.32) shows the message spectrum S ( f ) and noise power spectrum
P ( f ) . The noise power spectrum lies completely within the signal's band, and has a constant value there

of
N
.

Figure 6.32

(a) What is the message signal's power? What is the signal-to-noise ratio?
(b) Because the power in the message decreases with frequency, the signal-to-noise ratio is not constant
within sub-bands. What is the signal-to-noise ratio in the upper half of the frequency band?
(c) A clever ELEC 241 student suggests ltering the message before the transmitter modulates it so that
the signal spectrum is balanced (constant) across frequency. Realizing that this ltering a ects the
message signal, the student realizes that the receiver must also compensate for the message to arrive
intact. Draw a block diagram of this communication system. How does this system's signal-to-noise
ratio compare with that of the usual AM radio?

Problem 6.3: Complementary Filters
Complementary lters usually have opposite ltering characteristics (like a lowpass and a highpass)
and have transfer functions that add to one. Mathematically, H ( f ) and H ( f ) are complementary if

H ( f ) + H ( f ) = 1

We can use complementary lters to separate a signal into two parts by passing it through each lter. Each
output can then be transmitted separately and the original signal reconstructed at the receiver. Let's assume

the message is bandlimited to W Hz and that H ( f ) =
a + j 2 f
.

(a) What circuits would be used to produce the complementary lters?
