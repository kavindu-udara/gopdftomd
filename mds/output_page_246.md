240 CHAPTER 6. INFORMATION COMMUNICATION

Figure 6.30: Protocols are organized according to the level of detail required for information transmis-
sion. Protocols at the lower levels (shown toward the bottom) concern reliable bit transmission. Higher
level protocols concern how bits are organized to represent information, what kind of information is de-
ned by bit sequences, what software needs the information, and how the information is to be interpreted.
Bodies such as the IEEE (Institute for Electronics and Electrical Engineers) and the ISO (International
Standards Organization) de ne standards such as this. Despite being a standard, it does not constrain
protocol implementation so much that innovation and competitive individuality are ruled out.
an essentially stagnant Internet is but one example of the power of organizing how information transmission
occurs without overly constraining the details.

### 6.38 Information Communication Problems

Problem 6.1: Signals on Transmission Lines
A modulated signal needs to be sent over a transmission line having a characteristic impedance of Z = 50 .
So that the signal does not interfere with signals others may be transmitting, it must be bandpass ltered
so that its bandwidth is 1 MHz and centered at 3.5 MHz. The lter's gain should be one in magnitude. An
op-amp lter (Figure 6.31) is proposed.

(a) What is the transfer function between the input voltage and the voltage across the transmission line?
(b) Find values for the resistors and capacitors so that design goals are met.
This content is available online at http://cnx.org/content/m10352/2.24/ .
