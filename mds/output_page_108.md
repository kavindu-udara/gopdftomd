102 CHAPTER 4. FREQUENCY DOMAIN

Figure 4.2: The magnitude and phase of the periodic pulse sequence's spectrum is shown for positive-
frequency indices. Here
= 0 : 2 and A = 1 .

Also note the presence of a linear phase term (the rst term in \ c is proportional to frequency k=T ).
Comparing this term with that predicted from delaying a signal, a delay of = 2 is present in our signal.
Advancing the signal by this amount centers the pulse about the origin, leaving an even signal, which in
turn means that its spectrum is real-valued. Thus, our calculated spectrum is consistent with the properties
of the Fourier spectrum.

Exercise 4.2 (Solution on p. 139.)
What is the value of c ? Recalling that this spectral coe cient corresponds to the signal's average
value, does your answer make sense?

The phase plot shown in Figure 4.2 (Periodic Pulse Sequence) requires some explanation as it does not seem
to agree with what (4.10) suggests. There, the phase has a linear component, with a jump of every time
the sinusoidal term changes sign. We must realize that any integer multiple of 2 can be added to a phase
at each frequency without a ecting the value of the complex spectrum. We see that at frequency index 4,
the phase is nearly . The phase at index 5 is unde ned because the magnitude is zero in this example. At
index 6, the formula suggests that the phase of the linear term should be less than (more negative than) .
In addition, we expect a shift of in the phase between indices 4 and 6. Thus, the phase value predicted
by the formula is a little less than 2 . Because we can add 2 without a ecting the value of the spectrum
at index 6, the result is a slightly negative number as shown. Thus, the formula and the plot do agree. In
phase calculations like those made in MATLAB, values are usually con ned to the range [ ; ) by adding
some (possibly negative) multiple of 2 to each phase value.
Another aspect of the spectrum needs to be emphasized and this example demonstrates it well. Note
that the period T enters into the spectrum only through the ratio of the pulse width and the period. If
the period is changed in such a way that this ratio remains constant, the spectral plot in Figure (4.2) is
unchanged when considered as a function of coe cient index k . However, when we plot a spectrum as a
function of frequency k=T , the plot stretches or compresses horizontally as we change the period. We will
nd that what is important in applications is what frequencies, not what harmonic indices, are contained in
a periodic signal .
