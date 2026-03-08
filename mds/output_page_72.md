66 CHAPTER 3. ANALOG SIGNAL PROCESSING

### 3.19 Operational Ampli ers

Figure 3.40: The op-amp has four terminals to which connections can be made. Inputs attach to nodes
a and b , and the output is node c . As the circuit model on the right shows, the op-amp serves as an
ampli er for the di erence of the input node voltages.

Op-amps not only have the circuit model shown in Figure 3.40, but their element values are very special.

The input resistance , R , is typically large , on the order of 1 M .
The output resistance , R , is small , usually less than 100 .
The voltage gain , G , is large , exceeding 10 .

The large gain catches the eye; it suggests that an op-amp could turn a 1 mV input signal into a 100 V one.
If you were to build such a circuit attaching a voltage source to node a , attaching node b to the reference,
and looking at the output you would be disappointed. In dealing with electronic components, you cannot
forget the unrepresented but needed power supply.

Unmodeled limitations imposed by power supplies: It is impossible for electronic compo-
nents to yield voltages that exceed those provided by the power supply or for them to yield currents
that exceed the power supply's rating.

Typical power supply voltages required for op-amp circuits are 15 V. Attaching the 1 mv signal not
only would fail to produce a 100 V signal, the resulting waveform would be severely distorted. While a
desirable outcome if you are a rock & roll a cionado, high-quality stereos should not distort signals. Another
consideration in designing circuits with op-amps is that these element values are typical: Careful control of
the gain can only be obtained by choosing a circuit so that its element values dictate the resulting gain,
which must be smaller than that provided by the op-amp.
This content is available online at http://cnx.org/content/m0036/2.31/ .
