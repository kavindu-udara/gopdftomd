113

(a)

(b)

Figure 4.10: A periodic pulse signal, such as shown on the left part ( =T = 0 : 2) , serves as the input to
an RC lowpass lter. The input's period was 1 ms (millisecond). The lter's cuto frequency was set to
the various values indicated in the top row, which display the output signal's spectrum and the lter's
transfer function. The bottom row shows the output signal derived from the Fourier series coe cients
shown in the top row. (a) Periodic pulse signal (b) Top plots show the pulse signal's spectrum for various
cuto frequencies. Bottom plots show the lter's output signals.
The periodic pulse signal shown on the left above serves as the input to a RC -circuit shown in
Figure 3.28 that has the transfer function (calculated elsewhere)

H ( f ) =
1 j 2 fRC
(4.28)

Figure 4.10 shows the output changes as we vary the lter's cuto frequency. Note how the signal's
spectrum extends well above its fundamental frequency. Having a cuto frequency ten times higher
than the fundamental does perceptibly change the output waveform, rounding the leading and
trailing edges. As the cuto frequency decreases (center, then left), the rounding becomes more
prominent, with the leftmost waveform showing a small ripple.

Exercise 4.11 (Solution on p. 140.)
What is the average value of each output waveform? The correct answer may surprise you.

This example also illustrates the impact a lowpass lter can have on a waveform. The simple RC lter used
here has a rather gradual frequency response, which means that higher harmonics are smoothly suppressed.
Later, we will describe lters that have much more rapidly varying frequency responses, allowing a much
more dramatic selection of the input's Fourier coe cients.
More importantly, we have calculated the output of a circuit to a periodic input without writing,
much less solving, the di erential equation governing the circuit's behavior. Furthermore, we made these

![Im165](../output/images/test_119_Im165.R12.tif)
