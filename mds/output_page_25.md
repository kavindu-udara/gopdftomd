19

systems instantiated as computer programs. Subsequent modules describe how virtually all analog signal
processing can be performed with software.
Discrete-time systems can act on discrete-time signals in ways similar to those found in analog signals
and systems. Because of the role of software in discrete-time systems, many more di erent systems can be
envisioned and constructed with programs than can be with analog signals. Consequently, discrete-time
systems can be easily produced in software, with equivalent analog realizations di cult, if not impossible,
to design.
As important as linking analog signals to discrete-time ones may be, discrete-time signals are more
general, encompassing signals derived from analog ones and signals that aren't. For example, the characters
forming a text le form a sequence, which is also a discrete-time signal. We must deal with such symbolic
valued (p. 153) signals and systems as well.
As with analog signals, we seek ways of decomposing real-valued discrete-time signals into simpler com-
ponents. With this approach leading to a better understanding of signal structure, we can exploit that
structure to represent information (create ways of representing information with signals) and to extract in-
formation (retrieve the information thus represented). For symbolic-valued signals, the approach is di erent:
We develop a common representation of all symbolic-valued signals so that we can embody the information
they contain in a uni ed way. From an information representation perspective, the most important issue
becomes, for both real-valued and symbolic-valued signals, e ciency; What is the most parsimonious and
compact way to represent information so that it can be extracted later.

2.4.1 Real- and Complex-valued Signals

A discrete-time signal is represented symbolically as s ( n ) , where n = f : : : ; 1 ; 0 ; 1 ; : : : g . We usually draw
discrete-time signals as stem plots to emphasize the fact they are functions de ned only on the integers.
We can delay a discrete-time signal by an integer just as with analog ones. A delayed unit sample has the
expression ( n m ) , and equals one when n = m .

Figure 2.7: The discrete-time cosine signal is plotted as a stem plot. Can you nd the formula for this
signal?

2.4.2 Complex Exponentials

The most important signal is, of course, the complex exponential sequence .

s ( n ) = e (2.26)

2.4.3 Sinusoids

Discrete-time sinusoids have the obvious form s ( n ) = A cos (2 fn + ) . As opposed to analog complex
exponentials and sinusoids that can have their frequencies be any real value, frequencies of their discrete-
time counterparts yield unique waveforms only when f lies in the interval
;
. This property can be
easily understood by noting that adding an integer to the frequency of the discrete-time complex exponential
has no e ect on the signal's value.

e = e e

= e
(2.27)
