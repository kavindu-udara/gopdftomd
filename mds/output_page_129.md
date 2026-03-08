123

Figure 4.15: The ideal frequency response of the vocal tract as it produces the sounds oh and ee
are shown on the top left and top right, respectively. The spectral peaks are known as formants, and
are numbered consecutively from low to high frequency. The bottom plots show speech waveforms
corresponding to these sounds. F0 . The primary di erence between adult male and female/prepubescent speech is pitch. Before
puberty, pitch frequency for normal speech ranges between 150 400 Hz for both males and females. After
puberty, the vocal cords of males undergo a physical change, which has the e ect of lowering their pitch
frequency to the range 80 160 Hz. If we could examine the vocal cord output, we could probably discern
whether the speaker was male or female. This di erence is also readily apparent in the speech signal itself.
To simplify our speech modeling e ort, we shall assume that the pitch period is constant. With this
simpli cation, we collapse the vocal-cord-lung system as a simple source that produces the periodic pulse
signal (Figure 4.14). The sound pressure signal thus produced enters the mouth behind the tongue, creates
acoustic disturbances, and exits primarily through the lips and to some extent through the nose. Speech
specialists tend to name the mouth, tongue, teeth, lips, and nasal cavity the vocal tract . The physics
governing the sound disturbances produced in the vocal tract and those of an organ pipe are quite similar.
Whereas the organ pipe has the simple physical structure of a straight tube, the cross-section of the vocal
tract tube varies along its length because of the positions of the tongue, teeth, and lips. It is these positions
that are controlled by the brain to produce the vowel sounds. Spreading the lips, bringing the teeth together,
and bringing the tongue toward the front portion of the roof of the mouth produces the sound ee. Rounding
the lips, spreading the teeth, and positioning the tongue toward the back of the oral cavity produces the
sound oh. These variations result in a linear, time-invariant system that has a frequency response typi ed
by several peaks, as shown in Figure 4.15.
These peaks are known as formants . Thus, speech signal processors would say that the sound oh has
a higher rst formant frequency than the sound ee, with F2 being much higher during ee. F2 and F3
(the second and third formants) have more energy in ee than in oh. Rather than serving as a lter,
