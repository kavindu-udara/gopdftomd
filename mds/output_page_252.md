246 CHAPTER 6. INFORMATION COMMUNICATION

Problem 6.13: Mixed Analog and Digital Transmission
A signal m ( t ) is transmitted using amplitude modulation in the usual way. The signal has bandwidth W
Hz, and the carrier frequency is f . In addition to sending this analog signal, the transmitter also wants to
send ASCII text in an auxiliary band that lies slightly above the analog transmission band. Using an 8-bit
representation of the characters and a simple baseband BPSK signal set (the constant signal +1 corresponds
to a 0 , the constant 1 to a 1 ), the data signal d ( t ) representing the text is transmitted as the same time
as the analog signal m ( t ) . The transmission signal spectrum is as shown (Figure 6.38), and has a total
bandwidth B .

Figure 6.38

(a) Write an expression for the time-domain version of the transmitted signal in terms of m ( t ) and the
digital signal d ( t ) .
(b) What is the maximum datarate the scheme can provide in terms of the available bandwidth?
(c) Find a receiver that yields both the analog signal and the bit stream.

Problem 6.14: Digital Stereo
Just as with analog communication, it should be possible to send two signals simultaneously over a digital
channel. Assume you have two CD-quality signals (each sampled at 44.1 kHz with 16 bits/sample). One
suggested transmission scheme is to use a quadrature BPSK scheme. If b ( n ) and b ( n ) each represent a
bit stream, the transmitted signal has the form

x ( t ) = A
X
b ( n ) sin (2 f ( t nT )) p ( t nT ) + b ( n ) cos (2 f ( t nT )) p ( t nT )

where p ( t ) is a unit-amplitude pulse having duration T and b ( n ) , b ( n ) equal either +1 or -1 according
to the bit being transmitted for each signal. The channel adds white noise and attenuates the transmitted
signal.

(a) What value would you choose for the carrier frequency f ?
(b) What is the transmission bandwidth?
(c) What receiver would you design that would yield both bit streams?

Problem 6.15: Digital and Analog Speech Communication
Suppose we transmit speech signals over comparable digital and analog channels. We want to compare the
resulting quality of the received signals. Assume the transmitters use the same power, and the channels
introduce the same attenuation and additive white noise. Assume the speech signal has a 4 kHz bandwidth
and, in the digital case, is sampled at an 8 kHz rate with eight-bit A/D conversion. Assume simple binary
source coding and a modulated BPSK transmission scheme.

(a) What is the transmission bandwidth of the analog (AM) and digital schemes?
(b) Assume the speech signal's amplitude has a magnitude less than one. What is maximum amplitude
quantization error introduced by the A/D converter?
(c) In the digital case, each bit in quantized speech sample is received in error with probability p that
depends on signal-to-noise ratio
. However, errors in each bit have a di erent impact on the error in
the reconstructed speech sample. Find the mean-squared error between the transmitted and received
amplitude.
