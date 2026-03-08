125

(a) pulse

(b) voice spectrum

Figure 4.16: The vocal tract's transfer function, shown as the thin, smooth line, is superimposed on
the spectrum of actual male speech corresponding to the sound oh. The pitch lines corresponding to
harmonics of the pitch frequency are indicated. (a) The vocal cords' output spectrum P ( f ) . (b) The
vocal tract's transfer function, H ( f ) and the speech spectrum.
support this bandwidth (see more about AM and FM radio systems in Section 6.11).
E cient speech transmission systems exploit the speech signal's special structure: What makes speech
speech? You can conjure many signals that span the same frequencies as speech car engine sounds, violin
music, dog barks but don't sound at all like speech. We shall learn later that transmission of any 5 kHz
bandwidth signal requires about 80 kbps (thousands of bits per second) to transmit digitally. Speech signals
can be transmitted using less than 1 kbps because of its special structure. To reduce the digital bandwidth
so drastically means that engineers spent many years to develop signal processing and coding methods that
could capture the special characteristics of speech without destroying how it sounds. If you used a speech
transmission system to send a violin sound, it would arrive horribly distorted; speech transmitted the same
way would sound ne.
Exploiting the special structure of speech requires going beyond the capabilities of analog signal processing
systems. Many speech transmission systems work by nding the speaker's pitch and the formant frequencies.
Fundamentally, we need to do more than ltering to determine the speech signal's structure; we need to
manipulate signals in more ways than are possible with analog systems. Such exibility is achievable (but
not without some loss) with programmable digital systems.
