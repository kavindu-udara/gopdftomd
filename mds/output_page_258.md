252 CHAPTER 6. INFORMATION COMMUNICATION

Figure 6.41

Problem 6.28: Communication System Design
RU Communication Systems has been asked to design a communication system that meets the following
requirements.

The baseband message signal has a bandwidth of 10 kHz.
The RUCS engineers nd that the entropy H of the sampled message signal depends on how many
bits b are used in the A/D converter (see table below).
The signal is to be sent through a noisy channel having a bandwidth of 25 kHz centered at 2 MHz and
a signal-to-noise ratio within that band of 10 dB.
Once received, the message signal must have a signal-to-noise ratio of at least 20 dB.

Can these speci cations be met? Justify your answer.

Problem 6.29: HDTV
As HDTV (high-de nition television) was being developed, the FCC restricted this digital system to use in
the same bandwidth (6 MHz) as its analog (AM) counterpart. HDTV video is sampled on a 1035 1840
raster at 30 images per second for each of the three colors. The least-acceptable picture received by television
sets located at an analog station's broadcast perimeter has a signal-to-noise ratio of about 10 dB.

(a) Using signal-to-noise ratio as the criterion, how many bits per sample must be used to guarantee that
a high-quality picture, which achieves a signal-to-noise ratio of 20 dB, can be received by any HDTV
set within the same broadcast region?
(b) Assuming the digital television channel has the same characteristics as an analog one, how much
compression must HDTV systems employ?

Problem 6.30: Digital Cellular Telephones
In designing a digital version of a wireless telephone, you must rst consider certain fundamentals. First
of all, the quality of the received signal, as measured by the signal-to-noise ratio, must be at least as good
as that provided by wireline telephones (30 dB) and the message bandwidth must be the same as wireline
telephone. The signal-to-noise ratio of the allocated wireless channel, which has a 5 kHz bandwidth, measured
100 meters from the tower is 70 dB. The desired range for a cell is 1 km. Can a digital cellphone system be
designed according to these criteria?

Problem 6.31: Optimal Ethernet Random Access Protocols
Assume a population of N computers want to transmit information on a random access channel. The access
algorithm works as follows.
