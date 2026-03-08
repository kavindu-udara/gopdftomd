47

you buy a battery, you get a voltage source if its equivalent resistance is much smaller than the equivalent
resistance of the circuit to which you attach it. On the other hand, if you attach it to a circuit having a
small equivalent resistance, you bought a current source.

Léon Charles Thévenin: He was an engineer with France's Postes, Télégraphe et Téléphone. In
1883, he published (twice!) a proof of what is now called the Thévenin equivalent while developing
ways of teaching electrical engineering concepts at the École Polytechnique. He did not realize that
the same result had been published by Hermann Helmholtz, the renowned nineteenth century
physicist, thirty years earlier.

Hans Ferdinand Mayer: After earning his doctorate in physics in 1920, he turned to com-
munications engineering when he joined Siemens & Halske in 1922. In 1926, he published in a
German technical journal the Mayer-Norton equivalent. During his interesting career, he rose to
lead Siemens's Central Laboratory in 1936, surreptitiously leaked to the British all he knew of
German warfare capabilities a month after the Nazis invaded Poland, was arrested by the Gestapo
in 1943 for listening to BBC radio broadcasts, spent two years in Nazi concentration camps, and
went to the United States for four years working for the Air Force and Cornell University before
returning to Siemens in 1950. He rose to a position on Siemens's Board of Directors before retiring.

Edward L. Norton: Edward Norton was an electrical engineer who worked at Bell Laboratory
from its inception in 1922. In the same month when Mayer's paper appeared, Norton wrote in an
internal technical memorandum a paragraph describing the current-source equivalent. No evidence
suggests Norton knew of Mayer's publication.

### 3.8 Circuits with Capacitors and Inductors

Let's consider a circuit having something other than resistors and sources, such as shown in Figure 3.21.
Because of KVL, we know that v = v + v . The current through the capacitor is given by i = C
v ,
and this current equals that passing through the resistor. Substituting v = Ri into the KVL equation and
using the v-i relation for the capacitor, we arrive at

RC
d
v + v = v (3.11)

The input-output relation for circuits involving energy storage elements takes the form of an ordinary di er-
ential equation, which we must solve to determine what the output voltage is for a given input. In contrast
to resistive circuits, where we obtain an explicit input-output relation, we now have an implicit relation
that requires more work to obtain answers.
At this point, we could learn how to solve di erential equations. Note rst that even nding the di erential
equation relating an output variable to a source is often very tedious. The parallel and series combination
rules that apply to resistors don't directly apply when capacitors and inductors occur. We would have to slog
our way through the circuit equations, simplifying them until we nally found the equation that related the
source(s) to the output. At the turn of the twentieth century, a method was discovered that not only made
nding the di erential equation easy, but also simpli ed the solution process in the most common situation.
Although not original with him, Charles Steinmetz presented the key paper describing the impedance
approach in 1893. It allows circuits containing capacitors and inductors to be solved with the same methods
we have learned to solved resistor circuits. To use impedances, we must master complex numbers . Though
the arithmetic of complex numbers is mathematically more complicated than with real numbers, the increased
insight into circuit behavior and the ease with which circuits are solved with impedances is well worth the
diversion. But more importantly, the impedance concept is central to engineering and physics, having a
reach far beyond just circuits.
http://www-gap.dcs.st-and.ac.uk/~history/Biographies/Helmholtz.html
http://www.ece.rice.edu/~dhj/norton
This content is available online at http://cnx.org/content/m0023/2.12/ .
http://www.edisontechcenter.org/CharlesProteusSteinmetz.html
