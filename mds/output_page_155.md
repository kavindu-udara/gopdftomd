149

Figure 5.4: The spectrum of some bandlimited (to W Hz) signal is shown in the top plot. If the
sampling interval T is chosen too large relative to the bandwidth W , aliasing will occur. In the bottom
plot, the sampling interval is chosen su ciently small to avoid aliasing. Note that if the signal were not
bandlimited, the component spectra would always overlap.
no frequencies above the Nyquist frequency (22.05 kHz in our example). If, however, the signal contains
frequencies beyond the sound card's Nyquist frequency, the resulting aliasing can be impossible to remove.

Exercise 5.5 (Solution on p. 191.)
To gain a better appreciation of aliasing, sketch the spectrum of a sampled square wave. For

simplicity consider only the spectral repetitions centered at
1
, 0 ,
1
. Let the sampling interval

T be 1; consider two values for the square wave's period: 3.5 and 4. Note in particular where the
spectral lines go as the period decreases; some will move to the left and some to the right. What
property characterizes the ones going the same direction?

If we satisfy the Sampling Theorem's conditions, the signal will change only slightly during each pulse. As
we narrow the pulse, making smaller and smaller, the nonzero values of the signal s ( t ) p ( t ) will simply
be s ( nT ) , the signal's samples . If indeed the Nyquist frequency equals the signal's highest frequency, at
least two samples will occur within the period of the signal's highest frequency sinusoid. In these ways, the
sampling signal captures the sampled signal's temporal variations in a way that leaves all the original signal's
structure intact.

Exercise 5.6 (Solution on p. 191.)
What is the simplest bandlimited signal? Using this signal, convince yourself that less than two

samples/period will not su ce to specify it. If the sampling rate
1
is not high enough, what signal

would your resulting under-sampled signal become?
