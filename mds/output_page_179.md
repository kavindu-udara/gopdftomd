173

Computing the inverse DFT yields a length- N signal no matter what the actual duration of the unit-
sample response might be . If the unit-sample response has a duration less than or equal to N (it's a FIR
lter), computing the inverse DFT of the sampled frequency response indeed yields the unit-sample response.
If, however, the duration exceeds N , errors are encountered. The nature of these errors is easily explained
by appealing to the Sampling Theorem. By sampling in the frequency domain, we have the potential for
aliasing in the time domain (sampling in one domain, be it time or frequency, can result in aliasing in the
other) unless we sample fast enough. Here, the duration of the unit-sample response determines the minimal
sampling rate that prevents aliasing. For FIR systems they by de nition have nite-duration unit sample
responses the number of required DFT samples equals the unit-sample response's duration: N q .

Exercise 5.29 (Solution on p. 193.)
Derive the minimal DFT length for a length- q unit-sample response using the Sampling Theorem.
Because sampling in the frequency domain causes repetitions of the unit-sample response in the
time domain, sketch the time-domain result for various choices of the DFT length N .

Exercise 5.30 (Solution on p. 193.)
Express the unit-sample response of a FIR lter in terms of di erence equation coe cients. Note
that the corresponding question for IIR lters is far more di cult to answer: Consider the example
(Example 5.5).

For IIR systems, we cannot use the DFT to nd the system's unit-sample response: aliasing of the unit-
sample response will always occur. Consequently, we can only implement an IIR lter accurately in the time
domain with the system's di erence equation. Frequency-domain implementations are restricted to
FIR lters.
Another issue arises in frequency-domain ltering that is related to time-domain aliasing, this time when
we consider the output. Assume we have an input signal having duration N that we pass through a FIR
lter having a length- q + 1 unit-sample response. What is the duration of the output signal? The di erence
equation for this lter is
y ( n ) = b x ( n ) + + b x ( n q ) (5.54)

This equation says that the output depends on current and past input values, with the input value q samples
previous de ning the extent of the lter's memory of past input values. For example, the output at index
N depends on x ( N ) (which equals zero), x ( N 1) , through x ( N q ) . Thus, the output returns to zero
only after the last input value passes through the lter's memory. As the input signal's last value occurs at
index N 1 , the last nonzero output value occurs when n q = N 1 or n = q + N 1 . Thus, the output
signal's duration equals q + N .

Exercise 5.31 (Solution on p. 193.)
In words, we express this result as The output's duration equals the input's duration plus the
lter's duration minus one. Demonstrate the accuracy of this statement.

The main theme of this result is that a lter's output extends longer than either its input or its unit-sample
response. Thus, to avoid aliasing when we use DFTs, the dominant factor is not the duration of input or
of the unit-sample response, but of the output. Thus, the number of values at which we must evaluate the
frequency response's DFT must be at least q + N and we must compute the same length DFT of the input.
To accommodate a shorter signal than DFT length, we simply zero-pad the input: Ensure that for indices
extending beyond the signal's duration that the signal is zero. Frequency-domain ltering, diagrammed in
Figure 5.20, is accomplished by storing the lter's frequency response as the DFT H ( k ) , computing the
input's DFT X ( k ) , multiplying them to create the output's DFT Y ( k ) = H ( k ) X ( k ) , and computing the
inverse DFT of the result to yield y ( n ) .
