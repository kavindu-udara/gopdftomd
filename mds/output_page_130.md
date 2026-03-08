124 CHAPTER 4. FREQUENCY DOMAIN

rejecting high or low frequencies, the vocal tract serves to shape the spectrum of the vocal cords. In the
time domain, we have a periodic signal, the pitch, serving as the input to a linear system. We know that
the output the speech signal we utter and that is heard by others and ourselves will also be periodic.
Example time-domain speech signals are shown in Figure 4.15, where the periodicity is quite apparent.

Exercise 4.18 (Solution on p. 140.)
From the waveform plots shown in Figure 4.15, determine the pitch period and the pitch frequency.

Since speech signals are periodic, speech has a Fourier series representation given by a linear circuit's response
to a periodic signal (4.27). Because the acoustics of the vocal tract are linear, we know that the spectrum of
the output equals the product of the pitch signal's spectrum and the vocal tract's frequency response. We
thus obtain the fundamental model of speech production .

S ( f ) = P ( f ) H ( f ) (4.45)

Here, H ( f ) is the transfer function of the vocal tract system. The Fourier series for the vocal cords' output,
derived in this equation (p. 100), is

c = Ae

sin
k

(4.46)

and is plotted on the top in Figure 4.16. If we had, for example, a male speaker with about a 110 Hz
pitch ( T 9 : 1ms ) saying the vowel oh, the spectrum of his speech predicted by our model is shown in
Figure 4.16(b).
The model spectrum idealizes the measured spectrum, and captures all the important features. The
measured spectrum certainly demonstrates what are known as pitch lines , and we realize from our model
that they are due to the vocal cord's periodic excitation of the vocal tract. The vocal tract's shaping of the
line spectrum is clearly evident, but di cult to discern exactly, especially at the higher frequencies. The
model transfer function for the vocal tract makes the formants much more readily evident.

Exercise 4.19 (Solution on p. 141.)
The Fourier series coe cients for speech are related to the vocal tract's transfer function only at
the frequencies k=T , k 2 f 1 ; 2 ; : : : g ; see previous result (4.9). Would male or female speech tend to
have a more clearly identi able formant structure when its spectrum is computed? Consider, for
example, how the spectrum shown on the right in Figure 4.16 would change if the pitch were twice
as high ( 300Hz ).

When we speak, pitch and the vocal tract's transfer function are not static; they change according to their
control signals to produce speech. Engineers typically display how the speech spectrum changes over time
with what is known as a spectrogram (detailed in Section 5.10), an example of which is shown in Figure 4.17.
Note how the line spectrum, which indicates how the pitch changes, is visible during the vowels, but not
during the consonants (like the ce in Rice ).
The fundamental model for speech indicates how engineers use the physics underlying the signal gen-
eration process and exploit its structure to produce a systems model that suppresses the physics while
emphasizing how the signal is constructed. From everyday life, we know that speech contains a wealth of
information. We want to determine how to transmit and receive it. E cient and e ective speech transmis-
sion requires us to know the signal's properties and its structure (as expressed by the fundamental model of
speech production). We see from Figure 4.17, for example, that speech contains signi cant energy from zero
frequency up to around 5 kHz.
E ective speech transmission systems must be able to cope with signals having this bandwidth. It
is interesting that one system that does not support this 5 kHz bandwidth is the telephone: Telephone
systems act like a bandpass lter passing energy between about 200 Hz and 3.2 kHz. The most important
consequence of this ltering is the removal of high frequency energy. In our sample utterance, the ce sound
in Rice contains most of its energy above 3.2 kHz; this ltering e ect is why it is extremely di cult to
distinguish the sounds s and f over the telephone. Try this yourself: Call a friend and determine if they
can distinguish between the words six and x. If you say these words in isolation so that no context
