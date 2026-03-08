50 CHAPTER 3. ANALOG SIGNAL PROCESSING

Figure 3.23: The time and frequency domains are linked by assuming signals are complex exponentials.
In the time domain, signals can have any form. Passing into the frequency domain work room, signals
are represented entirely by complex amplitudes.

As we unfold the impedance story, we'll see that the powerful use of impedances suggested by Steinmetz
greatly simpli es solving circuits, alleviates us from solving di erential equations, and suggests a general way
of thinking about circuits. Because of the importance of this approach, let's go over how it works.

1. Even though it's not, pretend the source is a complex exponential. We do this because the impedance
approach simpli es nding how input and output are related. If it were a voltage source having voltage
v = p ( t ) (a pulse), still let v = V e . We'll learn how to get the pulse back later.
2. With a source equaling a complex exponential, all variables in a linear circuit will also be complex
exponentials having the same frequency. The circuit's only remaining mystery is what each variable's
complex amplitude might be. To nd these, we consider the source to be a complex number ( V here)
and the elements to be impedances.
3. We can now solve using series and parallel combination rules how the complex amplitude of any variable
relates to the sources complex amplitude.

Example 3.3

To illustrate the impedance approach, we refer to the RC circuit shown in Figure 3.24, and we
assume that v = V e .
http://www.edisontechcenter.org/CharlesProteusSteinmetz.html

![Im61](../output/images/test_056_Im61.R11.tif)
