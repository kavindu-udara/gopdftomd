2 CHAPTER 1. INTRODUCTION

Telegraphy represents the earliest electrical information system, and it dates from 1837. At that time,
electrical science was largely empirical, and only those with experience and intuition could develop telegraph
systems. Electrical science came of age when James Clerk Maxwell proclaimed in 1864 a set of equations
that he claimed governed all electrical phenomena. These equations predicted that light was an electro-
magnetic wave, and that energy could propagate. Because of the complexity of Maxwell's presentation, the
development of the telephone in 1876 was due largely to empirical work. Once Heinrich Hertz con rmed
Maxwell's prediction of what we now call radio waves in about 1882, Maxwell's equations were simpli ed
by Oliver Heaviside and others, and were widely read. This understanding of fundamentals led to a quick
succession of inventions the wireless telegraph (1899), the vacuum tube (1905), and radio broadcasting that
marked the true emergence of the communications age.
During the rst part of the twentieth century, circuit theory and electromagnetic theory were all an
electrical engineer needed to know to be quali ed and produce rst-rate designs. Consequently, circuit theory
served as the foundation and the framework of all of electrical engineering education. At mid-century, three
inventions changed the ground rules. These were the rst public demonstration of the rst electronic
computer (1946), the invention of the transistor (1947), and the publication of A Mathematical Theory
of Communication by Claude Shannon (1948). Although conceived separately, these creations gave birth
to the information age, in which digital and analog communication systems interact and compete for design
preferences. About twenty years later, the laser was invented, which opened even more design possibilities.
Thus, the primary focus shifted from how to build communication systems (the circuit theory era) to what
communications systems were intended to accomplish. Only once the intended system is speci ed can an
implementation be selected. Today's electrical engineer must be mindful of the system's ultimate goal,
and understand the tradeo s between digital and analog alternatives, and between hardware and software
con gurations in designing information systems.

note: Thanks to the translation e orts of Rice University's Disability Support Services , this
collection is now available in a Braille-printable version. Please click here to download a .zip le
containing all the necessary .dxb and image les.

### 1.2 Signals Represent Information

Whether analog or digital, information is represented by the fundamental quantity in electrical engineering:
the signal . Stated in mathematical terms, a signal is merely a function . Analog signals are continuous-
valued; digital signals are discrete-valued. The independent variable of the signal could be time (speech, for
example), space (images), or the integers (denoting the sequencing of letters and numbers in the football
score).

1.2.1 Analog Signals

Analog signals are usually signals de ned over continuous independent variable(s) . Speech, as
described in Section 4.10, is produced by your vocal cords exciting acoustic resonances in your vocal tract.
The result is pressure waves propagating in the air, and the speech signal thus corresponds to a function
having independent variables of space and time and a value corresponding to air pressure: s ( x ; t ) (Here we
use vector notation x to denote spatial coordinates). When you record someone talking, you are evaluating
the speech signal at a particular spatial location, x say. An example of the resulting waveform s ( x ; t ) is
shown in Figure 1.1.
Photographs are static, and are continuous-valued signals de ned over space. Black-and-white images
have only one value at each point in space, which amounts to its optical re ection properties. In Figure 1.2,
an image is shown, demonstrating that it (and all other images as well) are functions of two independent
spatial variables.
http://www-groups.dcs.st-andrews.ac.uk/~history/Biographies/Maxwell.html
http://www.dss.rice.edu/
http://cnx.org/content/m0000/latest/FundElecEngBraille.zip
This content is available online at http://cnx.org/content/m0001/2.27/ .
