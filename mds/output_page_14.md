8 CHAPTER 1. INTRODUCTION

(a) What is the smallest transmission interval that makes sense with the frequency f ?
(b) Assuming that ten cycles of the sinusoid comprise a single bit's transmission interval, what is the
datarate of this transmission scheme?
(c) Now suppose instead of using on-o signaling, we allow one of several di erent values for the
amplitude during any transmission interval. If N amplitude values are used, what is the resulting
datarate?
(d) The classic communications block diagram applies to the modem. Discuss how the transmitter must
interface with the message source since the source is producing letters of the alphabet, not bits.

Problem 1.3: Advanced Modems
To transmit symbols, such as letters of the alphabet, RU computer modems use two frequencies (1600
and 1800 Hz) and several amplitude levels. A transmission is sent for a period of time T (known as the
transmission or baud interval) and equals the sum of two amplitude-weighted carriers.

x ( t ) = A sin (2 f t ) + A sin (2 f t ) , 0 t T

We send successive symbols by choosing an appropriate frequency and amplitude combination, and sending
them one after another.

(a) What is the smallest transmission interval that makes sense to use with the frequencies given above?
In other words, what should T be so that an integer number of cycles of the carrier occurs?
(b) Sketch (using Matlab) the signal that modem produces over several transmission intervals. Make sure
you axes are labeled.
(c) Using your signal transmission interval, how many amplitude levels are needed to transmit ASCII
characters at a datarate of 3,200 bits/s? Assume use of the extended (8-bit) ASCII code.

note: We use a discrete set of values for A and A . If we have N values for amplitude A , and N
values for A , we have N N possible symbols that can be sent during each T second interval. To
convert this number into bits (the fundamental unit of information engineers use to qualify things),
compute log ( N N ) .
