150 CHAPTER 5. DIGITAL SIGNAL PROCESSING

### 5.4 Amplitude Quantization

The Sampling Theorem says that if we sample a bandlimited signal s ( t ) fast enough, it can be recovered with-
out error from its samples s ( nT ) , n 2 f : : : ; 1 ; 0 ; 1 ; : : : g . Sampling is only the rst phase of acquiring data
into a computer: Computational processing further requires that the samples be quantized : analog values
are converted into digital form. In short, we will have performed analog-to-digital (A/D) conversion .

(a)

(b)

Figure 5.5: A three-bit A/D converter assigns voltage in the range [ 1 ; 1] to one of eight integers
between 0 and 7. For example, all inputs having values lying between 0.5 and 0.75 are assigned the
integer value six and, upon conversion back to an analog value, they all become 0.625. The width of a

single quantization interval equals
2
. The bottom panel shows a signal going through the analog-

to-digital converter, where B is the number of bits used in the A/D conversion process (3 in the case
depicted here). First it is sampled, then amplitude-quantized to three bits. Note how the sampled signal
waveform becomes distorted after amplitude quantization. For example the two signal values between 0.5
and 0.75 become 0.625. This distortion is irreversible; it can be reduced (but not eliminated) by using
more bits in the A/D converter.

A phenomenon reminiscent of the errors incurred in representing numbers on a computer prevents signal
amplitudes from being converted with no error into a binary number representation. In analog-to-digital
conversion, the signal is assumed to lie within a prede ned range. Assuming we can scale the signal without
a ecting the information it expresses, we'll de ne this range to be [ 1 ; 1] . Furthermore, the A/D converter
assigns amplitude values in this range to a set of integers. A B -bit converter produces one of the integers
0 ; 1 ; : : : ; 2 1 for each sampled input. Figure 5.5 shows how a three-bit A/D converter assigns input
values to the integers. We de ne a quantization interval to be the range of values assigned to the same
integer. Thus, for our example three-bit A/D converter, the quantization interval is 0 : 25 ; in general, it is
2
.

Exercise 5.7 (Solution on p. 191.)
Recalling the plot of average daily highs in this frequency domain problem (Problem 4.5), why is
this plot so jagged? Interpret this e ect in terms of analog-to-digital conversion.

Because values lying anywhere within a quantization interval are assigned the same value for computer
processing, the original amplitude value cannot be recovered without error . Typically, the D/A
This content is available online at http://cnx.org/content/m0051/2.22/ .

![Im205](../output/images/test_156_Im205.R10.tif)
