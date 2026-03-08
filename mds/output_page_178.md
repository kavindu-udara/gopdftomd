172 CHAPTER 5. DIGITAL SIGNAL PROCESSING

is positive, we have a lowpass lter; negative a results in a highpass lter. The larger the coe cient
in magnitude, the more pronounced the lowpass or highpass ltering.

Example 5.6
The length- q boxcar lter (di erence equation found in a previous example (Example 5.4)) has the
frequency response

H e =
1

X
e (5.51)

This expression amounts to the Fourier transform of the boxcar signal (Figure 5.19). There we
found that this frequency response has a magnitude equal to the absolute value of dsinc ( f ) ; see
the length-10 lter's frequency response (Figure 5.11). We see that boxcar lters length- q signal
averagers have a lowpass behavior, having a cuto frequency of
.

Exercise 5.27 (Solution on p. 193.)
Suppose we multiply the boxcar lter's coe cients by a sinusoid: b =
cos (2 f m ) . Use Fourier
transform properties to determine the transfer function. How would you characterize this system:
Does it act like a lter? If so, what kind of lter and how do you control its characteristics with
the lter's coe cients?

These examples illustrate the point that systems described (and implemented) by di erence equations serve
as lters for discrete-time signals. The lter's order is given by the number p of denominator coe cients
in the transfer function (if the system is IIR) or by the number q of numerator coe cients if the lter is
FIR. When a system's transfer function has both terms, the system is usually IIR, and its order equals p
regardless of q . By selecting the coe cients and lter type, lters having virtually any frequency response
desired can be designed. This design exibility can't be found in analog systems. In the next section, we
detail how analog signals can be ltered by computers, o ering a much greater range of ltering possibilities
than is possible with circuits.

### 5.14 Filtering in the Frequency Domain

Because we are interested in actual computations rather than analytic calculations, we must consider the
details of the discrete Fourier transform. To compute the length- N DFT, we assume that the signal has a
duration less than or equal to N . Because frequency responses have an explicit frequency-domain speci cation
(5.47) in terms of lter coe cients, we don't have a direct handle on which signal has a Fourier transform
equaling a given frequency response. Finding this signal is quite easy. First of all, note that the discrete-
time Fourier transform of a unit sample equals one for all frequencies. Because the input and output of
linear, shift-invariant systems are related to each other by Y e = H e X e , a unit-sample

input, which has X e = 1 , results in the output's Fourier transform equaling the system's
transfer function .

Exercise 5.28 (Solution on p. 193.)
This statement is a very important result. Derive it yourself.

In the time-domain, the output for a unit-sample input is known as the system's unit-sample response ,
and is denoted by h ( n ) . Combining the frequency-domain and time-domain interpretations of a linear, shift-
invariant system's unit-sample response, we have that h ( n ) and the transfer function are Fourier transform
pairs in terms of the discrete-time Fourier transform .
h ( n ) ! H e (5.52)

Returning to the issue of how to use the DFT to perform ltering, we can analytically specify the frequency
response, and derive the corresponding length- N DFT by sampling the frequency response.

H ( k ) = H e ; k = f 0 ; : : : ; N 1 g (5.53)
This content is available online at http://cnx.org/content/m10257/2.17/ .
