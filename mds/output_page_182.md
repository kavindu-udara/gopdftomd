176 CHAPTER 5. DIGITAL SIGNAL PROCESSING

The colored plane plots ( N + q ) (2 q + 1) , the computations required to implement directly an FIR lter
with a di erence equation. The red plane displays the number of computations needed by the frequency
domain approach: 6 ( N + q ) + 5 ( N + q ) log ( N + q ) .

To determine for what signal and lter durations a time- or frequency-domain implementation would be
the most e cient, we need only count the computations required by each. For the time-domain, di erence-
equation approach, we need ( N + q ) (2 q + 1) . As Figure 5.20 shows, the frequency-domain approach requires
three Fourier transforms but we need to only compute two of them: typically, H ( k ) is computed once and

stored. Each requires
5 K
log K computations for a length- K FFT. We must multiply two complex-valued

length- K spectra; that requires 6 K computations. Because the output-signal-duration-determined length
must be at least N + q , we must compare

( N + q ) (2 q + 1) ! 6 ( N + q ) + 2
5
( N + q ) log ( N + q )

Exact analytic evaluation of this comparison is quite di cult (we have a transcendental equation to solve).
Insight into this comparison is best obtained by dividing by N + q .

2 q + 1 ! 6 + 5log ( N + q )

With this manipulation, we are evaluating the number of computations per sample. For any given value
of the lter's order q , the right side, the number of frequency-domain computations, will exceed the left if
the signal's duration is long enough. However, for lter durations greater than about 10, as long as the
input is at least 10 samples, the frequency-domain approach is faster so long as the FFT's power-of-two
constraint is advantageous . Graphically, the parameter range over which the frequency-domain approach
results in fewer computations is shown in the gure. As the analysis and the plot shows, only for short lter
durations does the di erence equation approach require fewer computations.
The frequency-domain approach is not yet viable; what will we do when the input signal is in nitely long?
The di erence equation scenario ts perfectly with the envisioned digital ltering structure (Figure 5.24), but
so far we have required the input to have limited duration (so that we could calculate its Fourier transform).
The solution to this problem is quite simple: Section the input into frames, lter each, and add the results
together. To section a signal means expressing it as a linear combination of length- N non-overlapping
This content is available online at http://cnx.org/content/m10279/2.16/ .

![Im223](../output/images/test_182_Im223.Im1.png)
