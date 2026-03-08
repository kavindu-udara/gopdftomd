189

(a) What answer should Samantha obtain?
(b) As a check, her group partner Sammy says that he computed the inverse DFT of her answer and got
( n + 1) + ( n 1) . Does Sammy's result mean that Samantha's answer is wrong?
(c) The homework problem says to lowpass- lter the sequence by multiplying its DFT by

H ( k ) =

(
1 k = f 0 ; 1 ; 7 g

0 otherwise

and then computing the inverse DFT. Will this ltering algorithm work? If so, nd the ltered output;
if not, why not?

Problem 5.31: Stock Market Data Processing
Because a trading week lasts ve days, stock markets frequently compute running averages each day over the
previous ve trading days to smooth price uctuations. The technical stock analyst at the Buy-Lo Sell-Hi
brokerage rm has heard that FFT ltering techniques work better than any others (in terms of producing
more accurate averages).

(a) What is the di erence equation governing the ve-day averager for daily stock prices?
(b) Design an e cient FFT-based ltering algorithm for the broker. How much data should be processed
at once to produce an e cient algorithm? What length transform should be used?
(c) Is the analyst's information correct that FFT techniques produce more accurate averages than any
others? Why or why not?

Problem 5.32: Echoes
Echoes not only occur in canyons, but also in auditoriums and telephone circuits. In one situation where
the echoed signal has been sampled, the input signal x ( n ) emerges as x ( n ) + a x ( n n ) + a x ( n n ) .

(a) Find the di erence equation of the system that models the production of echoes.
(b) To simulate this echo system, ELEC 241 students are asked to write the most e cient (quickest)
program that has the same input-output relationship. Suppose the duration of x ( n ) is 1,000 and that
a =
, n = 10 , a =
, and n = 25 . Half the class votes to just program the di erence equation
while the other half votes to program a frequency domain approach that exploits the speed of the FFT.
Because of the undecided vote, you must break the tie. Which approach is more e cient and why?
(c) Find the transfer function and di erence equation of the system that suppresses the echoes. In other
words, with the echoed signal as the input, what system's output is the signal x ( n ) ?

Problem 5.33: Digital Filtering of Analog Signals
RU Electronics wants to develop a lter that would be used in analog applications, but that is implemented
digitally. The lter is to operate on signals that have a 10 kHz bandwidth, and will serve as a lowpass lter.

(a) What is the block diagram for your lter implementation? Explicitly denote which components are
analog, which are digital (a computer performs the task), and which interface between analog and
digital worlds.
(b) What sampling rate must be used and how many bits must be used in the A/D converter for the
acquired signal's signal-to-noise ratio to be at least 60 dB? For this calculation, assume the signal is a
sinusoid.
(c) If the lter is a length-128 FIR lter (the duration of the lter's unit-sample response equals 128),
should it be implemented in the time or frequency domain?
(d) Assuming H e is the transfer function of the digital lter, what is the transfer function of your
system?

Problem 5.34: Signal Compression
Because of the slowness of the Internet, lossy signal compression becomes important if you want signals to
be received quickly. An enterprising ELEC 241 student has proposed a scheme based on frequency-domain
