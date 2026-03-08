147

Boole discovered this equivalence in the mid-nineteenth century. It laid the foundation for what we now
call Boolean algebra, which expresses as equations logical statements. More importantly, any computer
using base-2 representations and arithmetic can also easily evaluate logical statements. This fact makes an
integer-based computational device much more powerful than might be apparent.

### 5.3 The Sampling Theorem

5.3.1 Analog-to-Digital Conversion

Because of the way computers are organized, signal must be represented by a nite number of bytes. This
restriction means that both the time axis and the amplitude axis must be quantized : They must each be
a multiple of the integers. Quite surprisingly, the Sampling Theorem allows us to quantize the time axis
without error for some signals. The signals that can be sampled without introducing error are interesting,
and as described in the next section, we can make a signal samplable by ltering. In contrast, no one
has found a way of performing the amplitude quantization step without introducing an unrecoverable error.
Thus, a signal's value can no longer be any real number. Signals processed by digital computers must
be discrete-valued : their values must be proportional to the integers. Consequently, analog-to-digital
conversion introduces error .

5.3.2 The Sampling Theorem

Digital transmission of information and digital signal processing all require signals to rst be acquired by
a computer. One of the most amazing and useful results in electrical engineering is that signals can be
converted from a function of time into a sequence of numbers without error : We can convert the numbers
back into the signal with (theoretically) no error. Harold Nyquist, a Bell Laboratories engineer, rst derived
this result, known as the Sampling Theorem, in the 1920s. It found no real application back then. Claude
Shannon, also at Bell Laboratories, revived the result once computers were made public after World War II.
The sampled version of the analog signal s ( t ) is s ( nT ) , with T known as the sampling interval .
Clearly, the value of the original signal at the sampling times is preserved; the issue is how the signal
values between the samples can be reconstructed since they are lost in the sampling process. To char-
acterize sampling, we approximate it as the product x ( t ) = s ( t ) P ( t ) , with P ( t ) being the periodic
pulse signal. The resulting signal, as shown in Figure 5.3, has nonzero values only during the time intervals
nT
; nT +
, n 2 f : : : ; 1 ; 0 ; 1 ; : : : g . For our purposes here, we center the periodic pulse signal about
the origin so that its Fourier series coe cients are real (the signal is even).

p ( t ) =
X
c e
(5.5)

where

c =
sin

(5.6)

If the properties of s ( t ) and the periodic pulse signal are chosen properly, we can recover s ( t ) from x ( t ) by
ltering.
To understand how signal values between the samples can be lled in, we need to calculate the sampled
signal's spectrum. Using the Fourier series representation of the periodic sampling signal,

x ( t ) =
X
c e
s ( t ) (5.7)
http://www-groups.dcs.st-and.ac.uk/~history/Biographies/Boole.html
This content is available online at http://cnx.org/content/m0050/2.19/ .
We assume that we do not use oating-point A/D converters.
