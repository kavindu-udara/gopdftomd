148 CHAPTER 5. DIGITAL SIGNAL PROCESSING

Figure 5.3: The waveform of an example signal is shown in the top plot and its sampled version in the
bottom.
exponential and the signal. Evaluating this transform directly is quite easy.

Z

s ( t ) e
e dt =

Z

s ( t ) e
dt = S f
k
(5.8)

Thus, the spectrum of the sampled signal consists of weighted (by the coe cients c ) and delayed versions
of the signal's spectrum (Figure 5.4).

X ( f ) =
X
c S f
k
(5.9)

In general, the terms in this sum overlap each other in the frequency domain, rendering recovery of the
original signal impossible. This unpleasant phenomenon is known as aliasing . If, however, we satisfy two
conditions:

The signal s ( t ) is bandlimited has power in a restricted frequency range to W Hz, and
the sampling interval T is small enough so that the individual components in the sum do not overlap
T < 1 = 2 W ,

aliasing will not occur. In this delightful case, we can recover the original signal by lowpass ltering x ( t )
with a lter having a cuto frequency equal to W Hz. These two conditions ensure the ability to recover a
bandlimited signal from its sampled version: We thus have the Sampling Theorem .

Exercise 5.4 (Solution on p. 191.)
The Sampling Theorem (as stated) does not mention the pulse width . What is the e ect of this
parameter on our ability to recover a signal from its samples (assuming the Sampling Theorem's
two conditions are met)?

The frequency
1 T
, known today as the Nyquist frequency and the Shannon sampling frequency ,

corresponds to the highest frequency at which a signal can contain energy and remain compatible with
the Sampling Theorem. High-quality sampling systems ensure that no aliasing occurs by unceremoniously
lowpass ltering the signal (cuto frequency being slightly lower than the Nyquist frequency) before sampling.
Such systems therefore vary the anti-aliasing lter's cuto frequency as the sampling rate varies. Because
such quality features cost money, many sound cards do not have anti-aliasing lters or, for that matter,
