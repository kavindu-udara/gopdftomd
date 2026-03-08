122 CHAPTER 4. FREQUENCY DOMAIN

Figure 4.14: The systems model for the vocal tract. The signals l ( t ) , p ( t ) , and s ( t ) are the air
pressure provided by the lungs, the periodic pulse output provided by the vocal cords, and the speech
output respectively. Control signals from the brain are shown as entering the systems from the top.
Clearly, these come from the same source, but for modeling purposes we describe them separately since
they control di erent aspects of the speech signal.
is linear, we can use linear systems in our model with a fair amount of accuracy. The naturalness of linear
system models for speech does not extend to other situations. In many cases, the underlying mathematics
governed by the physics, biology, and/or chemistry of the problem are nonlinear, leaving linear systems
models as approximations. Nonlinear models are far more di cult at the current state of knowledge to
understand, and information engineers frequently prefer linear models because they provide a greater level
of comfort, but not necessarily a su cient level of accuracy.
Figure 4.13 shows the actual speech production system and Figure 4.14 the model speech production
system. The characteristics of the model depends on whether you are saying a vowel or a consonant. We
concentrate rst on the vowel production mechanism. When the vocal cords are placed under tension by
the surrounding musculature, air pressure from the lungs causes the vocal cords to vibrate. To visualize this
e ect, take a rubber band and hold it in front of your lips. If held open when you blow through it, the air
passes through more or less freely; this situation corresponds to breathing mode. If held tautly and close
together, blowing through the opening causes the sides of the rubber band to vibrate. This e ect works best
with a wide rubber band. You can imagine what the air ow is like on the opposite side of the rubber band
or the vocal cords. Your lung power is the simple source referred to earlier; it can be modeled as a constant
supply of air pressure. The vocal cords respond to this input by vibrating, which means the output of this
system is some periodic function.

Exercise 4.17 (Solution on p. 140.)
Note that the vocal cord system takes a constant input and produces a periodic air ow that corre-
sponds to its output signal. Is this system linear or nonlinear? Justify your answer.

Singers modify vocal cord tension to change the pitch to produce the desired musical note. Vocal cord
tension is governed by a control input to the musculature; in system's models we represent control inputs as
signals coming into the top or bottom of the system. Certainly in the case of speech and in many other cases
as well, it is the control input that carries information, impressing it on the system's output. The change of
signal structure resulting from varying the control input enables information to be conveyed by the signal,
a process generically known as modulation . In singing, musicality is largely conveyed by pitch; in western
speech, pitch is much less important. A sentence can be read in a monotone fashion without completely
destroying the information expressed by the sentence. However, the di erence between a statement and a
question is frequently expressed by pitch changes. For example, note the sound di erences between Let's
go to the park and Let's go to the park? ;
For some consonants, the vocal cords vibrate just as in vowels. For example, the so-called nasal sounds
n and m have this property. For others, the vocal cords do not produce a periodic output. Going back to
mechanism, when consonants such as f are produced, the vocal cords are placed under much less tension,
which results in turbulent ow. The resulting output air ow is quite erratic, so much so that we describe it
as being noise . We de ne noise carefully later when we delve into communication problems.
The vocal cords' periodic output can be well described by the periodic pulse train p ( t ) as shown in
the periodic pulse signal (Figure 4.1), with T denoting the pitch period. The spectrum of this signal (4.9)
contains harmonics of the frequency 1 =T , what is known as the pitch frequency or the fundamental
