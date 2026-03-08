152 CHAPTER 5. DIGITAL SIGNAL PROCESSING

Once we have acquired signals with an A/D converter, we can process them using digital hardware or software.
It can be shown that if the computer processing is linear, the result of sampling, computer processing, and
un-sampling is equivalent to some analog linear system. Why go to all the bother if the same function can
be accomplished using analog techniques? Knowing when digital processing excels and when it does not is
an important issue.

### 5.5 Discrete-Time Signals and Systems

Mathematically, analog signals are functions having as their independent variables continuous quantities,
such as space and time. Discrete-time signals are functions de ned on the integers; they are sequences. As
with analog signals, we seek ways of decomposing discrete-time signals into simpler components. Because
this approach leads to a better understanding of signal structure, we can exploit that structure to represent
information (create ways of representing information with signals) and to extract information (retrieve the
information thus represented). For symbolic-valued signals, the approach is di erent: We develop a common
representation of all symbolic-valued signals so that we can embody the information they contain in a
uni ed way. From an information representation perspective, the most important issue becomes, for both
real-valued and symbolic-valued signals, e ciency: what is the most parsimonious and compact way to
represent information so that it can be extracted later.

5.5.1 Real- and Complex-valued Signals

A discrete-time signal is represented symbolically as s ( n ) , where n = f : : : ; 1 ; 0 ; 1 ; : : : g .

Figure 5.7: The discrete-time cosine signal is plotted as a stem plot. Can you nd the formula for this
signal?

We usually draw discrete-time signals as stem plots to emphasize the fact they are functions de ned only on
the integers. We can delay a discrete-time signal by an integer just as with analog ones. A signal delayed
by m samples has the expression s ( n m ) .

5.5.2 Complex Exponentials

The most important signal is, of course, the complex exponential sequence .

s ( n ) = e (5.12)

Note that the frequency variable f is dimensionless and that adding an integer to the frequency of the
discrete-time complex exponential has no e ect on the signal's value.

e = e e

= e
(5.13)

This derivation follows because the complex exponential evaluated at an integer multiple of 2 equals one.
Thus, we need only consider frequency to have a value in some unit-length interval.
This content is available online at http://cnx.org/content/m10342/2.15/ .
